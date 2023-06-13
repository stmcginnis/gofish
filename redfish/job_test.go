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

var jobBody = strings.NewReader(
	`{
		"@odata.type": "#Job.v1_2_1.Job",
		"Id": "RebootRack",
		"Name": "Scheduled Nightly Reboot of the rack",
		"JobStatus": "OK",
		"JobState": "Running",
		"StartTime": "2018-04-01T00:01+6:00",
		"PercentComplete": 24,
		"Messages": [
			{
				"Message": "Successfully imported and applied Server Configuration Profile.",
				"MessageId": "SYS053",
				"MessageArgs": [],
				"MessageArgs@odata.count": 0
			}
		],
		"Schedule": {
			"Lifetime": "P4Y",
			"InitialStartTime": "2018-01-01T01:00:00+06:00",
			"RecurrenceInterval": "P1D",
			"EnabledDaysOfWeek": [
				"Monday",
				"Tuesday",
				"Wednesday",
				"Thursday",
				"Friday"
			]
		},
		"Steps": {
			"@odata.id": "/redfish/v1/JobService/Jobs/RebootRack/Steps"
		},
		"StepOrder": [
			"Red"			
		],
		"@odata.id": "/redfish/v1/JobService/Jobs/RebootRack"
	}`)

// TestJob tests the parsing of Job objects.
func TestJob(t *testing.T) {
	var result Job
	err := json.NewDecoder(jobBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "RebootRack" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Scheduled Nightly Reboot of the rack" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.JobStatus != common.OKHealth {
		t.Errorf("Invalid JobStatus: %s", result.JobStatus)
	}

	if result.JobState != RunningJobState {
		t.Errorf("Invalid JobState: %s", result.JobState)
	}

	if result.PercentComplete != 24 {
		t.Errorf("Received invalid PercentComplete: %d", result.PercentComplete)
	}
}
