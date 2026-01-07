//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var jobBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Job.Job",
		"@odata.type": "#Job.v1_0_7.Job",
		"@odata.id": "/redfish/v1/JobService/Jobs/Job-1",
		"Id": "Job-1",
		"Name": "JobOne",
		"Description": "Job One",
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
		"StepOrder": [
        	"Step-1"
    	],		
		"Steps": {
			"@odata.id": "/redfish/v1/JobService/Jobs/Job-1/Steps"
		},
		"JobState": "Running",
		"JobStatus": "OK"
	}`)

// TestJob tests the parsing of Job objects.
func TestJob(t *testing.T) {
	var result Job
	err := json.NewDecoder(jobBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Job-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "JobOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Payload.HTTPOperation != "POST" {
		t.Errorf("Invalid HTTP operation: %s", result.Payload.HTTPOperation)
	}

	if result.JobState != RunningJobState {
		t.Errorf("Invalid JobState: %s", result.JobState)
	}

	if result.JobStatus != OKHealth {
		t.Errorf("Invalid JobStatus: %s", result.JobStatus)
	}

	if len(result.Messages) != 1 {
		t.Errorf("Incorrect number of job messages: %d", len(result.Messages))
	}

	if result.steps != "/redfish/v1/JobService/Jobs/Job-1/Steps" {
		t.Errorf("Invalid steps reference")
	}
}
