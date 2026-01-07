//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #TaskService.v1_3_0.TaskService

package schemas

import (
	"encoding/json"
)

type TaskServiceOverWritePolicy string

const (
	// ManualTaskServiceOverWritePolicy Completed tasks are not automatically overwritten.
	ManualTaskServiceOverWritePolicy TaskServiceOverWritePolicy = "Manual"
	// OldestTaskServiceOverWritePolicy Oldest completed tasks are overwritten.
	OldestTaskServiceOverWritePolicy TaskServiceOverWritePolicy = "Oldest"
)

// TaskService This resource contains a task service for a Redfish
// implementation.
type TaskService struct {
	Entity
	// CompletedTaskOverWritePolicy shall contain the overwrite policy for
	// completed tasks. This property shall indicate if the task service overwrites
	// completed task information.
	CompletedTaskOverWritePolicy TaskServiceOverWritePolicy
	// DateTime shall contain the current date and time for the task service, with
	// UTC offset.
	DateTime string
	// LifeCycleEventOnTaskStateChange shall indicate whether a task state change
	// sends an event. Services should send an event containing a message defined
	// in the Task Event Message Registry when the state of a task changes.
	LifeCycleEventOnTaskStateChange bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// TaskAutoDeleteTimeoutMinutes shall contain the number of minutes after which
	// a completed task, where 'TaskState' contains the value 'Completed',
	// 'Killed', 'Cancelled', or 'Exception', is deleted by the service.
	//
	// Version added: v1.2.0
	TaskAutoDeleteTimeoutMinutes uint
	// TaskMonitorAutoExpirySeconds shall contain the number of seconds after
	// reading a task monitor for a completed task until the service deletes the
	// task monitor. If the task is cancelled before it completes the task monitor
	// shall be removed at that time.
	//
	// Version added: v1.3.0
	TaskMonitorAutoExpirySeconds *uint `json:",omitempty"`
	// Tasks shall contain a link to a resource collection of type
	// 'TaskCollection'.
	tasks string
	// deleteAllCompletedTasksTarget is the URL to send DeleteAllCompletedTasks requests.
	deleteAllCompletedTasksTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a TaskService object from the raw JSON.
func (t *TaskService) UnmarshalJSON(b []byte) error {
	type temp TaskService
	type tActions struct {
		DeleteAllCompletedTasks ActionTarget `json:"#TaskService.DeleteAllCompletedTasks"`
	}
	var tmp struct {
		temp
		Actions tActions
		Tasks   Link `json:"Tasks"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TaskService(tmp.temp)

	// Extract the links to other entities for later
	t.deleteAllCompletedTasksTarget = tmp.Actions.DeleteAllCompletedTasks.Target
	t.tasks = tmp.Tasks.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *TaskService) Update() error {
	readWriteFields := []string{
		"ServiceEnabled",
		"TaskAutoDeleteTimeoutMinutes",
		"TaskMonitorAutoExpirySeconds",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetTaskService will get a TaskService instance from the service.
func GetTaskService(c Client, uri string) (*TaskService, error) {
	return GetObject[TaskService](c, uri)
}

// ListReferencedTaskServices gets the collection of TaskService from
// a provided reference.
func ListReferencedTaskServices(c Client, link string) ([]*TaskService, error) {
	return GetCollectionObjects[TaskService](c, link)
}

// This action shall delete all 'Task' resources whose 'TaskState' property
// contains 'Completed', 'Killed', 'Cancelled', or 'Exception'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TaskService) DeleteAllCompletedTasks() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.deleteAllCompletedTasksTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Tasks gets the Tasks collection.
func (t *TaskService) Tasks() ([]*Task, error) {
	if t.tasks == "" {
		return nil, nil
	}
	return GetCollectionObjects[Task](t.client, t.tasks)
}
