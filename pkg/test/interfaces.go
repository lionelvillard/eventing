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

	"github.com/onsi/gomega"
	"knative.dev/eventing/pkg/test/component/recorder"
	"knative.dev/eventing/pkg/test/resource"
	"knative.dev/eventing/test/lib"
)

// Test defines functions for configuring and running a single test case
type Test interface {
	// ConformanceIt marks this test as a conformance test
	ConformanceIt()

	// Run the test, supplied as a lambda
	Run(fn func(ctx TestContext))
}

// TestContext is the context when running a test case
type TestContext interface {
	resource.ResourceContext

	// --- Built-in components

	// NewEventRecorderOrFail returns an instantiated event recorder component (or fail)
	NewEventRecorderOrFail() recorder.EventRecorder

	// --- Assertion

	// Gomega assertion
	gomega.Gomega
}

// NewTest creates a single test case
func NewTest(t *testing.T) Test {
	lc := lib.Setup(t, true)

	return &testImpl{
		lc:        lc,
		logPrefix: "",
	}
}
