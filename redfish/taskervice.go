//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/stmcginnis/gofish/common"
)

type TaskOverWritePolicy string

const (
	// ManualTaskOverWritePolicy Completed tasks are not automatically overwritten.
	ManualTaskOverWritePolicy TaskOverWritePolicy = "Manual"
	// OldestTaskOverWritePolicy Oldest completed tasks are overwritten.
	OldestTaskOverWritePolicy TaskOverWritePolicy = "Oldest"
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
	CompletedTaskOverWritePolicy TaskOverWritePolicy
	// DateTime system time.
	DateTime time.Time
	// LifeCycleEventOnTaskStateChange whether an event is reported when the task status is changed.
	LifeCycleEventOnTaskStateChange bool
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled indicates whether this service isenabled.
	ServiceEnabled bool
	// Status describes the status and health of a resource and its children.
	Status common.Status
	// TaskAutoDeleteTimeoutMinutes shall contain the number of minutes after which a completed task, where TaskState
	// contains the value 'Completed', 'Killed', 'Cancelled', or 'Exception', is deleted by the service.
	TaskAutoDeleteTimeoutMinutes int
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

	// This is a read/write object, so we need to save the raw object data for later
	taskService.rawData = b

	return nil
}

// Tasks gets the collection of tasks of this task service
func (taskService *TaskService) Tasks() ([]*Task, error) {
	return ListReferencedTasks(taskService.GetClient(), taskService.tasks)
}

// Update commits updates to this object's properties to the running system.
func (taskService *TaskService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(TaskService)
	original.UnmarshalJSON(taskService.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
		"TaskAutoDeleteTimeoutMinutes",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(taskService).Elem()

	return taskService.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetTaskService will get a TaskService instance from the service.
func GetTaskService(c common.Client, uri string) (*TaskService, error) {
	return common.GetObject[TaskService](c, uri)
}
