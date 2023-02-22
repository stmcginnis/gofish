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
	assertMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Check default redfish fields", func(t *testing.T) {
		c := &common.TestClient{}
		result.Client = c

		err := json.NewDecoder(strings.NewReader(simpleTaskServiceBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.tasks, "/redfish/v1/TaskService/Tasks")
		assertMessage(t, result.CompletedTaskOverWritePolicy, "Oldest")
	})
}
