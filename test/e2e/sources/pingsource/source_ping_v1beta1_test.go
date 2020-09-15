// +build e2e

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

package pingsource

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/util/uuid"

	"knative.dev/eventing/pkg/test"
)

func TestPingSourceV1Beta1NG(t *testing.T) {
	test.NewTest(t).Run(func(ctx test.TestContext) {
		// Create an event recorder where to send events to
		recorder := ctx.NewEventRecorderOrFail()

		// Create our PingSource
		ctx.CreateFromYAMLOrFail(fmt.Sprintf(pingSourceTemplate, "e2e-ping-source", uuid.NewUUID(), recorder.Name()))

		//ctx.Eventually(ctx.Get()).Should(matchers.BeReady())
		//Eventually(client.Logs(eventrecorder.Name)).Should(SatisfyAll(
		//	HaveEventCount(1))))
	})
}
