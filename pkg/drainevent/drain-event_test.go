// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package drainevent_test

import (
	"testing"
	"time"

	"github.com/aws/aws-node-termination-handler/pkg/drainevent"
	h "github.com/aws/aws-node-termination-handler/pkg/test"
)

func TestTimeUntilEvent(t *testing.T) {
	startTime := time.Now().Add(time.Second * 10)
	expected := startTime.Sub(time.Now()).Round(time.Second)

	event := &drainevent.DrainEvent{
		StartTime: startTime,
	}

	result := event.TimeUntilEvent()
	h.Equals(t, expected, result.Round(time.Second))
}
