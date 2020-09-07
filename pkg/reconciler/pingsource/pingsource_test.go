/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Veroute.on 2.0 (the "License");
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
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	fakeeventingclient "knative.dev/eventing/pkg/client/injection/client/fake"
	"knative.dev/eventing/pkg/client/injection/reconciler/sources/v1beta1/pingsource"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/client/injection/ducks/duck/v1/addressable"
	fakekubeclient "knative.dev/pkg/client/injection/kube/client/fake"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	logtesting "knative.dev/pkg/logging/testing"
	"knative.dev/pkg/resolver"
	"knative.dev/pkg/tracker"

	. "knative.dev/eventing/pkg/reconciler/testing"
	. "knative.dev/eventing/pkg/test"
	_ "knative.dev/pkg/client/injection/ducks/duck/v1beta1/addressable/fake"
	. "knative.dev/pkg/reconciler/testing"
)

const (
	sourceName = "test-ping-source"
	testNS     = "testnamespace"
)

func init() {
	// Add types to scheme
	_ = appsv1.AddToScheme(scheme.Scheme)
	_ = corev1.AddToScheme(scheme.Scheme)
	_ = duckv1.AddToScheme(scheme.Scheme)
}

func TestAllCases(t *testing.T) {
	const (
		missingSink = "missingsink"
		valid       = "valid"
	)

	table := TableTest{
		{
			Name: "bad workqueue key",
			// Make sure Reconcile handles bad keys.
			Key: "too/many/parts",
		}, {
			Name: "key not found",
			// Make sure Reconcile handles good keys that don't exist.
			Key: "foo/not-found",
		}, {
			Name:              missingSink,
			Objects:           ParseYAMLs(t, ObjectsPath(missingSink)),
			Key:               testNS + "/" + sourceName,
			WantStatusUpdates: ToUpdateActions(ParseYAMLs(t, StatusUpdatesPath(missingSink))),
			WantEvents:        ToStrings(ParseYAMLs(t, EventsPath(missingSink))),
		}, {
			Name:              valid,
			Objects:           ParseYAMLs(t, ObjectsPath(valid)),
			Key:               testNS + "/" + sourceName,
			WantStatusUpdates: ToUpdateActions(ParseYAMLs(t, StatusUpdatesPath(valid))),
		},
	}

	logger := logtesting.TestLogger(t)
	table.Test(t, MakeFactory(func(ctx context.Context, listers *Listers, cmw configmap.Watcher) controller.Reconciler {
		ctx = addressable.WithDuck(ctx)
		r := &Reconciler{
			kubeClientSet:    fakekubeclient.Get(ctx),
			pingLister:       listers.GetPingSourceV1beta1Lister(),
			deploymentLister: listers.GetDeploymentLister(),
			tracker:          tracker.New(func(types.NamespacedName) {}, 0),
		}
		r.sinkResolver = resolver.NewURIResolver(ctx, func(types.NamespacedName) {})

		return pingsource.NewReconciler(ctx, logging.FromContext(ctx),
			fakeeventingclient.Get(ctx), listers.GetPingSourceV1beta1Lister(),
			controller.GetEventRecorder(ctx), r)
	},
		true,
		logger,
	))
}
