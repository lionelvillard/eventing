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
package recorder

import (
	"fmt"

	"github.com/google/uuid"
	"knative.dev/pkg/test/helpers"

	"knative.dev/eventing/pkg/test/resource"
	pkgTest "knative.dev/pkg/test"
)

type EventRecorder interface {
	// Name return the name of the event recorder service
	Name() string

	// TODO: assertions....
}

func NewOrFail(rc resource.ResourceContext) EventRecorder {
	name := helpers.AppendRandomString("event-recorder-")
	uuid, err := uuid.NewUUID()
	if err != nil {
		rc.Fatal(err)
	}

	rc.CreateFromYAMLOrFail(eventRecordPod(name, uuid.String()))
	rc.CreateFromYAMLOrFail(eventRecordService(name, uuid.String()))

	return &eventRecorderImpl{
		name: name,
	}
}

type eventRecorderImpl struct {
	name string
}

func (e eventRecorderImpl) Name() string {
	return e.name
}

func eventRecordPod(name, uuid string) string {
	image := pkgTest.ImagePath("recordevents")
	return fmt.Sprintf(eventLoggerPodTemplate, name, uuid, "recordevents", image)
}

func eventRecordService(name, uuid string) string {
	return fmt.Sprintf(eventLoggerServiceTemplate, name, uuid, uuid)
}
