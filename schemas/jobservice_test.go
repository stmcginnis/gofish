//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var simpleJobServiceBody = `{
   "@odata.context":"/redfish/v1/$metadata#JobService.JobService",
   "@odata.etag":"\"1614826331\"",
   "@odata.id":"/redfish/v1/JobService",
   "@odata.type":"#JobService.v1_0_4.JobService",
   "DateTime":"2023-04-25T01:15:32+00:00",
   "Description":"Job Service",
   "Id":"JobService",
   "Name":"Job Service",
   "ServiceEnabled":true,
   "ServiceCapabilities": {
      "Scheduling": true,
      "MaxSteps": 50,
      "MaxJobs": 50
   },
   "Status":{
      "Health":"OK",
      "HealthRollup": "OK",
      "State":"Enabled"
   },
   "Jobs":{
      "@odata.id":"/redfish/v1/JobService/Jobs"
   },
   "Log":{
      "@odata.id":"/redfish/v1/JobService/Log"
   }
}`

func TestJobService(t *testing.T) {
	var result JobService
	err := json.NewDecoder(strings.NewReader(simpleJobServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "JobService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Job Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.jobs != "/redfish/v1/JobService/Jobs" {
		t.Errorf("Received invalid jobs link: %s", result.jobs)
	}

	if result.log != "/redfish/v1/JobService/Log" {
		t.Errorf("Received invalid log link: %s", result.log)
	}

	if !result.ServiceCapabilities.Scheduling ||
		*result.ServiceCapabilities.MaxJobs != 50 ||
		*result.ServiceCapabilities.MaxSteps != 50 {
		t.Errorf("Received invalid service capabilities")
	}
}
