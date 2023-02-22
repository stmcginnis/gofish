//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"time"

	"github.com/stmcginnis/gofish/common"
)

// TaskService is used to represent the task service offered by the redfish API
type TaskService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// CompletedTaskOverWritePolicy how to handle the completed tasks.
	CompletedTaskOverWritePolicy string
	// DateTime system time.
	DateTime time.Time
	// LifeCycleEventOnTaskStateChange whether an event is reported when the task status is changed.
	LifeCycleEventOnTaskStateChange bool
	// ServiceEnabled indicates whether this service isenabled.
	ServiceEnabled bool
	// Status describes the status and health of a resource and its children.
	Status common.Status
	// tasks points towards the tasks store endpoint
	tasks string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a TaskService object from the raw JSON.
func (taskService *TaskService) UnmarshalJSON(b []byte) error {
	type temp TaskService
	var t struct {
		temp
		Tasks common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*taskService = TaskService(t.temp)
	taskService.tasks = t.Tasks.String()
	taskService.rawData = b

	return nil
}

// GetTaskService will get a TaskService instance from the service.
func GetTaskService(c common.Client, uri string) (*TaskService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var taskService TaskService
	err = json.NewDecoder(resp.Body).Decode(&taskService)
	if err != nil {
		return nil, err
	}
	taskService.SetClient(c)
	return &taskService, nil
}

// Tasks gets the collection of tasks of this task service
func (taskService *TaskService) Tasks() ([]*Task, error) {
	return ListReferencedTasks(taskService.Client, taskService.tasks)
}
