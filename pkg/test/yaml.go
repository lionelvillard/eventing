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
package test

import (
	"testing"

	k8stesting "k8s.io/client-go/testing"

	"k8s.io/apimachinery/pkg/runtime"

	"knative.dev/eventing/pkg/test/internal/yaml"
)

// ParseYAML parses YAML files in the given directory
func ParseYAMLs(t *testing.T, pathname string) []runtime.Object {
	return yaml.ParseYAMLs(t, pathname)
}

// ParseYAML parses a single YAML file
func ParseYAML(t *testing.T, pathname string) runtime.Object {
	return yaml.ParseYAML(t, pathname)
}

// ToUpdateActions convets list of objects to list of update actions
func ToUpdateActions(objs []runtime.Object) []k8stesting.UpdateActionImpl {
	return yaml.ToUpdateActions(objs)
}

func ToStrings(objs []runtime.Object) []string {
	return yaml.ToStrings(objs)
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
