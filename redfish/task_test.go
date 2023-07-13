//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var taskBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Task.Task",
		"@odata.type": "#Task.v1_4_1.Task",
		"@odata.id": "/redfish/v1/Task",
		"Id": "Task-1",
		"Name": "TaskOne",
		"Description": "Task One",
		"EndTime": "2012-03-07T14:44+06:00",
		"HidePayload": false,
		"Messages": [{
            "Resolution": "None",
            "@odata.type": "#Message.v1_1_2.Message",
            "Message": "An update is in progress.",
            "MessageArgs": [],
            "MessageId": "Update.1.0.UpdateInProgress",
            "MessageSeverity": "OK"
        }],
		"Payload": {
			"HttpHeaders": ["User-Agent: Tadpole"],
			"HttpOperation": "POST",
			"JsonBody": "{}",
			"TargetUri": "http://example.com/API"
		},
		"PercentComplete": 60,
		"StartTime": "2012-03-07T14:04+06:00",
		"TaskMonitor": "http://example.com/API/Tasks/1",
		"TaskState": "Running",
		"TaskStatus": "OK"
	}`)

// TestTask tests the parsing of Task objects.
func TestTask(t *testing.T) {
	var result Task
	err := json.NewDecoder(taskBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Task-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "TaskOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Payload.HTTPOperation != "POST" {
		t.Errorf("Invalid HTTP operation: %s", result.Payload.HTTPOperation)
	}

	if result.TaskState != RunningTaskState {
		t.Errorf("Invalid TaskState: %s", result.TaskState)
	}

	if result.TaskStatus != common.OKHealth {
		t.Errorf("Invalid TaskStatus: %s", result.TaskStatus)
	}

	if len(result.Messages) != 1 {
		t.Errorf("Incorrect number of task messages: %d", len(result.Messages))
	}
}
