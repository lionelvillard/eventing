/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package yaml

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

// ParseYAML parses YAML files in the given directory
func ParseYAMLs(t *testing.T, pathname string) []runtime.Object {
	info, err := os.Stat(pathname)
	if err != nil {
		t.Fatal(err)
	}

	if info.IsDir() {
		return readDir(t, pathname)
	}

	return []runtime.Object{readFile(t, pathname)}
}

// ParseYAML parses a single YAML file
func ParseYAML(t *testing.T, pathname string) runtime.Object {
	return readFile(t, pathname)
}

// readDir parses all files in a single directory and it's descendant directories
// if the recursive flag is set to true.
func readDir(t *testing.T, pathname string) []runtime.Object {
	list, err := ioutil.ReadDir(pathname)
	if err != nil {
		t.Fatal(err)
	}

	var aggregated []runtime.Object
	for _, f := range list {
		name := path.Join(pathname, f.Name())
		if !f.IsDir() {
			aggregated = append(aggregated, readFile(t, name))
		}
	}
	return aggregated
}

// readFile parses a single YAML file into the given obj
func readFile(t *testing.T, pathname string) runtime.Object {
	file, err := os.Open(pathname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewYAMLToJSONDecoder(file)

	out := unstructured.Unstructured{}
	if err := decoder.Decode(out); err != nil {
		t.Fatal(err)
	}

	obj, err := scheme.Scheme.New(out.GroupVersionKind())
	if err != nil {
		t.Fatal(err)
	}

	if err := decoder.Decode(obj); err != nil {
		t.Fatal(err)
	}

	return obj
}

func ObjectsPath(name string) string {
	return "./testdata/" + name + "/objects"
}

func StatusUpdatesPath(name string) string {
	return "./testdata/" + name + "/statusupdates"
}

func EventsPath(name string) string {
	return "./testdata/" + name + "/events"
}
