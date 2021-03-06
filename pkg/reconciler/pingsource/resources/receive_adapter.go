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

package resources

import (
	"strconv"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/eventing/pkg/adapter/mtping"
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/system"
)

var (
	mtlabels = map[string]string{
		"sources.knative.dev/role":    "adapter",
		"eventing.knative.dev/source": controllerAgentName,
	}
)

type Args struct {
	ServiceAccountName string
	AdapterName        string
	Image              string
	MetricsConfig      string
	LoggingConfig      string
	LeConfig           string
	NoShutdownAfter    int
}

// MakeReceiveAdapter generates the mtping deployment for pingsources
func MakeReceiveAdapter(args Args) *v1.Deployment {
	replicas := int32(1)

	return &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployments",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: system.Namespace(),
			Name:      args.AdapterName,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: mtlabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: mtlabels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: args.ServiceAccountName,
					Containers: []corev1.Container{
						{
							Name:  "dispatcher",
							Image: args.Image,
							Env:   makeEnv(args),

							// Set low resource requests and limits.
							// This should be configurable.
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("125m"),
									corev1.ResourceMemory: resource.MustParse("64Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("1000m"),
									corev1.ResourceMemory: resource.MustParse("2048Mi"),
								},
							},
							Ports: []corev1.ContainerPort{{
								Name:          "metrics",
								ContainerPort: 9090,
							}},
						},
					},
				},
			},
		},
	}
}

func makeEnv(args Args) []corev1.EnvVar {
	return []corev1.EnvVar{{
		Name:  system.NamespaceEnvKey,
		Value: system.Namespace(),
	}, {
		Name: adapter.EnvConfigNamespace,
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: "metadata.namespace",
			},
		},
	}, {
		Name:  adapter.EnvConfigMetricsConfig,
		Value: args.MetricsConfig,
	}, {
		Name:  adapter.EnvConfigLoggingConfig,
		Value: args.LoggingConfig,
	}, {
		Name:  adapter.EnvConfigLeaderElectionConfig,
		Value: args.LeConfig,
	}, {
		Name:  mtping.EnvNoShutdownAfter,
		Value: strconv.Itoa(args.NoShutdownAfter),
	}}
}
