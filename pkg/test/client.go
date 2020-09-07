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
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/eventing/pkg/test/internal/client"
)

type Client interface {
	Create(obj runtime.Object) error
	Get(name string, objType runtime.Object) (runtime.Object, error)
	Logs(name string) (string, error)
}

// NewClient instantiates and returns several clientsets required for making request to the
// cluster specified by the combination of clusterName and configPath.
func NewClient(t *testing.T) Client {
	return client.NewClient(t)
}
