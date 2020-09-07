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

package yaml

import (
	"k8s.io/apimachinery/pkg/runtime"
	k8stesting "k8s.io/client-go/testing"
)

// ToUpdateActions convets list of objects to list of update actions
func ToUpdateActions(objs []runtime.Object) []k8stesting.UpdateActionImpl {
	actions := make([]k8stesting.UpdateActionImpl, len(objs))
	for i, obj := range objs {
		actions[i].Object = obj
	}
	return actions
}

func ToStrings(objs []runtime.Object) []string {
	// TODO
	return nil
}
