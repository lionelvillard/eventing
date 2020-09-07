/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"testing"

	v1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis/duck"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/client-go/dynamic"
	"knative.dev/pkg/test"
	"knative.dev/pkg/test/helpers"
)

type client struct {
	namespace string
	dynamic   dynamic.Interface
}

// NewClient instantiates a client
func NewClient(t *testing.T) *client {
	// Create a new namespace to run this test case.
	namespace := makeK8sNamespace(t.Name())
	t.Logf("namespace is : %q", namespace)

	config, err := test.BuildClientConfig(test.Flags.Kubeconfig, test.Flags.Cluster)
	if err != nil {
		t.Fatal(err)
	}

	d, err := dynamic.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	//
	//// Clean up resources if the test is interrupted in the middle.
	//test.CleanupOnInterrupt(func() { TearDown(client) }, t.Logf)

	return &client{
		namespace: namespace,
		dynamic:   d,
	}
}

func makeK8sNamespace(baseFuncName string) string {
	base := helpers.MakeK8sNamePrefix(baseFuncName)
	return names.SimpleNameGenerator.GenerateName(base + "-")
}

func (c *client) Create(obj runtime.Object) error {
	u, err := duck.ToUnstructured(obj.(duck.OneOfOurs)) // TODO: not OneOfOurs
	if err != nil {
		return err
	}

	gvr, _ := meta.UnsafeGuessKindToResource(obj.GetObjectKind().GroupVersionKind())

	// TODO: RetryWebHookError
	_, err = c.dynamic.Resource(gvr).Namespace(c.namespace).Create(u, metav1.CreateOptions{})
	return err
}

func (c *client) Get(name string, objType runtime.Object) (runtime.Object, error) {
	gvr, _ := meta.UnsafeGuessKindToResource(objType.GetObjectKind().GroupVersionKind())
	obj, err := c.dynamic.Resource(gvr).Namespace(c.namespace).Get(name, metav1.GetOptions{})
	return obj, err
}

func (c *client) Logs(name string) (string, error) {
	gvr := v1.SchemeGroupVersion.WithResource("pods")
	_, err := c.dynamic.Resource(gvr).Namespace(c.namespace).Get(name, metav1.GetOptions{}, "log")

	// TODO: convert to string

	return "", err
}
