//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var jobServiceBody = `{
   "@odata.type": "#JobService.v1_0_4.JobService",
    "Id": "JobService",
    "Name": "Job Service",
    "DateTime": "2023-06-13T06:41:27-05:00",
    "Status": {
        "State": "Enabled",
        "Health": "OK"
    },
    "ServiceEnabled": true,
    "ServiceCapabilities": {
        "MaxJobs": 100,
        "MaxSteps": 50,
        "Scheduling": true
    },
    "Jobs": {
        "@odata.id": "/redfish/v1/JobService/Jobs"
    },
    "Log": {
        "@odata.id": "/redfish/v1/JobService/Log"
    },
    "Actions": {
        "Oem": {
            "#Contoso.EasyButton": {
                "target": "/redfish/v1/JobService/Contoso.EasyButton",
                "@Redfish.ActionInfo": "/redfish/v1/JobService/EasyButtonActionInfo"
            }
        }
    },
    "@odata.id": "/redfish/v1/JobService"
}`

func TestJobService(t *testing.T) {
	var result JobService
	err := json.NewDecoder(strings.NewReader(jobServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "JobService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Job Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ServiceEnabled != true {
		t.Errorf("Received invalid serviceEnabled: %t", result.ServiceEnabled)
	}

	if result.jobs != "/redfish/v1/JobService/Jobs" {
		t.Errorf("Received invalid jobs: %s", result.jobs)
	}
}
