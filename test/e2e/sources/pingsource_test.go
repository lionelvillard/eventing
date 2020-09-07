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
package sources_test

import (
	corev1 "k8s.io/api/core/v1"

	sourcesv1beta1 "knative.dev/eventing/pkg/apis/sources/v1beta1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "knative.dev/eventing/pkg/test"
)

var _ = Describe("PingSource", func() {

	It("can create a PingSource", func() {
		// Load configs
		objs := ParseYAMLs(T, "./testdata")
		source := objs[0].(*sourcesv1beta1.PingSource)
		eventrecorder := objs[1].(*corev1.Pod)

		// Deploy and wait
		err := client.Create(source)
		Expect(err).NotTo(HaveOccurred())

		err = client.Create(eventrecorder)
		Expect(err).NotTo(HaveOccurred())

		// Assert
		Eventually(client.Get(source.Name, source)).Should(BeReady())
		Eventually(client.Logs(eventrecorder.Name)).Should(SatisfyAll(
			HaveEventCount(1)))
		// HaveEventSource(sourcesv1beta1.PingSourceSource(client.Namespace(), source.Name))))
	})

})
