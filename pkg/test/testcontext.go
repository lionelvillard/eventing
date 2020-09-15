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

package test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"knative.dev/eventing/pkg/test/component/recorder"
)

type testContextImpl struct {
	client    dynamic.Interface
	t         *testing.T
	namespace string
	*gomega.WithT
}

func (c *testContextImpl) CreateOrFail(obj runtime.Object) {
	gvr, _ := meta.UnsafeGuessKindToResource(obj.GetObjectKind().GroupVersionKind())
	u, err := toUnstructured(obj)
	if err != nil {
		c.t.Fatal(err)
	}
	_, err = c.client.Resource(gvr).Namespace(c.namespace).Create(context.TODO(), u, v1.CreateOptions{})
	if err != nil {
		c.t.Fatal(err)
	}
}

func (c *testContextImpl) CreateFromYAMLOrFail(yamlSpec string) {
	decoder := yaml.NewYAMLToJSONDecoder(strings.NewReader(yamlSpec))

	out := unstructured.Unstructured{}
	if err := decoder.Decode(out); err != nil {
		c.t.Fatal(err)
	}

	gvr, _ := meta.UnsafeGuessKindToResource(out.GroupVersionKind())
	_, err := c.client.Resource(gvr).Namespace(c.namespace).Create(context.TODO(), &out, v1.CreateOptions{})
	if err != nil {
		c.t.Fatal(err)
	}
}

func (c *testContextImpl) NewEventRecorderOrFail() recorder.EventRecorder {
	return recorder.NewOrFail(c)
}

func toUnstructured(desired runtime.Object) (*unstructured.Unstructured, error) {
	// Convert desired to unstructured.Unstructured
	b, err := json.Marshal(desired)
	if err != nil {
		return nil, err
	}
	ud := &unstructured.Unstructured{}
	if err := json.Unmarshal(b, ud); err != nil {
		return nil, err
	}
	return ud, nil
}

// --- testing.T wrapper

func (c *testContextImpl) Error(args ...interface{}) {
	c.t.Error(args...)
}

func (c *testContextImpl) Errorf(format string, args ...interface{}) {
	c.t.Errorf(format, args...)
}

func (c *testContextImpl) Fail() {
	c.t.Fail()
}

func (c *testContextImpl) FailNow() {
	c.t.FailNow()
}

func (c *testContextImpl) Failed() bool {
	return c.t.Failed()
}

func (c *testContextImpl) Fatal(args ...interface{}) {
	c.t.Fatal(args...)
}

func (c *testContextImpl) Fatalf(format string, args ...interface{}) {
	c.t.Fatalf(format, args)
}

func (c *testContextImpl) Helper() {
	c.t.Helper()
}

func (c *testContextImpl) Log(args ...interface{}) {
	c.t.Log(args...)
}

func (c *testContextImpl) Logf(format string, args ...interface{}) {
	c.t.Logf(format, args...)
}

func (c *testContextImpl) Name() string {
	return c.t.Name()
}

func (c *testContextImpl) Skip(args ...interface{}) {
	c.t.Skip(args...)
}

func (c *testContextImpl) SkipNow() {
	c.t.SkipNow()
}

func (c *testContextImpl) Skipf(format string, args ...interface{}) {
	c.t.Skipf(format, args...)
}

func (c *testContextImpl) Skipped() bool {
	return c.t.Skipped()
}
