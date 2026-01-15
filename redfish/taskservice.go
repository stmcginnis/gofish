//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #TaskService.v1_2_1.TaskService

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type OverWritePolicy string

const (
	// ManualOverWritePolicy Completed tasks are not automatically overwritten.
	ManualOverWritePolicy OverWritePolicy = "Manual"
	// OldestOverWritePolicy Oldest completed tasks are overwritten.
	OldestOverWritePolicy OverWritePolicy = "Oldest"
)

// TaskService This resource contains a task service for a Redfish
// implementation.
type TaskService struct {
	common.Entity
	// CompletedTaskOverWritePolicy shall contain the overwrite policy for
	// completed tasks. This property shall indicate if the task service overwrites
	// completed task information.
	CompletedTaskOverWritePolicy OverWritePolicy
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TaskAutoDeleteTimeoutMinutes shall contain the number of minutes after which
	// a completed task, where 'TaskState' contains the value 'Completed',
	// 'Killed', 'Cancelled', or 'Exception', is deleted by the service.
	//
	// Version added: v1.2.0
	TaskAutoDeleteTimeoutMinutes uint
	// Tasks shall contain a link to a resource collection of type
	// 'TaskCollection'.
	tasks string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a TaskService object from the raw JSON.
func (t *TaskService) UnmarshalJSON(b []byte) error {
	type temp TaskService
	var tmp struct {
		temp
		Tasks common.Link `json:"tasks"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TaskService(tmp.temp)

	// Extract the links to other entities for later
	t.tasks = tmp.Tasks.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *TaskService) Update() error {
	readWriteFields := []string{
		"ServiceEnabled",
		"Status",
		"TaskAutoDeleteTimeoutMinutes",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
}

// GetTaskService will get a TaskService instance from the service.
func GetTaskService(c common.Client, uri string) (*TaskService, error) {
	return common.GetObject[TaskService](c, uri)
}

// ListReferencedTaskServices gets the collection of TaskService from
// a provided reference.
func ListReferencedTaskServices(c common.Client, link string) ([]*TaskService, error) {
	return common.GetCollectionObjects[TaskService](c, link)
}

// Tasks gets the Tasks collection.
func (t *TaskService) Tasks(client common.Client) ([]*Task, error) {
	if t.tasks == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Task](client, t.tasks)
}
