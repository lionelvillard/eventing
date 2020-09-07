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
package matchers

import (
	"fmt"

	"github.com/onsi/gomega/format"

	"knative.dev/pkg/apis"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

type BeReady struct{}

func (matcher *BeReady) Match(actual interface{}) (success bool, err error) {
	switch obj := actual.(type) {
	case duckv1beta1.KResource:
		ready := obj.Status.GetCondition(apis.ConditionReady)
		return ready != nil && ready.IsTrue(), nil
		// TODO: k8s object
	default:
		return false, fmt.Errorf("BeReady matcher does not support %v", actual)
	}
}

func (matcher *BeReady) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n%s\nto be ready", format.Object(actual, 1))

}

func (matcher *BeReady) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n%s\nnot to be ready", format.Object(actual, 1))
}
