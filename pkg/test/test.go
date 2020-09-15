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
	"github.com/onsi/gomega"

	"knative.dev/eventing/test/lib"
)

type testImpl struct {
	lc        *lib.Client
	logPrefix string
}

func (t *testImpl) Run(fn func(TestContext)) {
	tc := &testContextImpl{
		client:    t.lc.Dynamic,
		t:         t.lc.T,
		namespace: t.lc.Namespace,
		WithT:     gomega.NewGomegaWithT(t.lc.T),
	}

	t.log("=== Running Test")

	fn(tc)

	lib.TearDown(t.lc)
}

func (t *testImpl) log(args ...interface{}) {
	a := append([]interface{}{t.logPrefix}, args...)
	t.lc.T.Log(a...)
}

func (t *testImpl) ConformanceIt() {
	t.logPrefix = t.logPrefix + "[conformance] "
}
