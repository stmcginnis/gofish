//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var simpleTaskServiceBody = `{
   "@odata.context":"/redfish/v1/$metadata#TaskService.TaskService",
   "@odata.etag":"\"1614826331\"",
   "@odata.id":"/redfish/v1/TaskService",
   "@odata.type":"#TaskService.v1_1_4.TaskService",
   "CompletedTaskOverWritePolicy":"Oldest",
   "DateTime":"2022-10-10T00:04:51-05:00",
   "Description":"Task Service",
   "Id":"TaskService",
   "LifeCycleEventOnTaskStateChange":true,
   "Name":"Task Service",
   "ServiceEnabled":true,
   "Status":{
      "Health":"OK",
      "State":"Enabled"
   },
   "Tasks":{
      "@odata.id":"/redfish/v1/TaskService/Tasks"
   }
}`

func TestTaskService(t *testing.T) {
	var result TaskService
	err := json.NewDecoder(strings.NewReader(simpleTaskServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "TaskService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Task Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.tasks != "/redfish/v1/TaskService/Tasks" {
		t.Errorf("Received invalid tasks link: %s", result.tasks)
	}

	if result.CompletedTaskOverWritePolicy != "Oldest" {
		t.Errorf("Received %s completed task overwrite policy", result.CompletedTaskOverWritePolicy)
	}
}
