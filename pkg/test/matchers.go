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
	"github.com/onsi/gomega/types"
	matchers "knative.dev/eventing/pkg/test/internal/matchers"
)

// BeReady succeeds if a resource is ready
func BeReady() types.GomegaMatcher {
	return &matchers.BeReady{}
}

// HaveEventCount succeeds actual is a EventRecorder log and it contains
// the number of events
func HaveEventCount(expected int) types.GomegaMatcher {
	return &matchers.HaveEventCount{Count: expected}
}
