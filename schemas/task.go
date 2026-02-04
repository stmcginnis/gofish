//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Task.v1_7_4.json
// 2022.3 - #Task.v1_7_4.Task

package schemas

import (
	"encoding/json"
)

type TaskState string

const (
	// NewTaskState shall represent that the task is newly created but has not
	// started.
	NewTaskState TaskState = "New"
	// StartingTaskState shall represent that the task is starting.
	StartingTaskState TaskState = "Starting"
	// RunningTaskState shall represent that the task is executing.
	RunningTaskState TaskState = "Running"
	// SuspendedTaskState shall represent that the task has been suspended but is
	// expected to restart and is therefore not complete.
	SuspendedTaskState TaskState = "Suspended"
	// InterruptedTaskState shall represent that the task has been interrupted but
	// is expected to restart and is therefore not complete.
	InterruptedTaskState TaskState = "Interrupted"
	// PendingTaskState shall represent that the task is pending some condition and
	// has not yet begun to execute.
	PendingTaskState TaskState = "Pending"
	// StoppingTaskState shall represent that the task is stopping but is not yet
	// complete.
	StoppingTaskState TaskState = "Stopping"
	// CompletedTaskState shall represent that the task completed successfully or
	// with warnings.
	CompletedTaskState TaskState = "Completed"
	// KilledTaskState shall represent that the task is complete because an
	// operator killed it.
	KilledTaskState TaskState = "Killed"
	// ExceptionTaskState shall represent that the task completed with errors.
	ExceptionTaskState TaskState = "Exception"
	// ServiceTaskState shall represent that the task is now running as a service
	// and expected to continue operation until stopped or killed.
	ServiceTaskState TaskState = "Service"
	// CancellingTaskState shall represent that the task is in the process of being
	// cancelled.
	CancellingTaskState TaskState = "Cancelling"
	// CancelledTaskState shall represent that either a 'DELETE' operation on a
	// task monitor or 'Task' resource or by an internal process cancelled the
	// task.
	CancelledTaskState TaskState = "Cancelled"
)

// Task This resource contains a task for a Redfish implementation.
type Task struct {
	Entity
	// EndTime shall indicate the date and time when the task was completed. This
	// property shall not appear if the task is running or otherwise has not been
	// completed. This property shall appear only if the 'TaskState' is
	// 'Completed', 'Killed', 'Cancelled', or 'Exception'.
	EndTime string
	// EstimatedDuration shall indicate the estimated total time needed to complete
	// the task. The value is not expected to change while the task is in progress,
	// but the service may update the value if it obtains new information that
	// significantly changes the expected duration. Services should be conservative
	// in the reported estimate and clients should treat this value as an estimate.
	//
	// Version added: v1.6.0
	EstimatedDuration string
	// HidePayload shall indicate whether the contents of the payload should be
	// hidden from view after the task has been created. If 'true', responses shall
	// not return the 'Payload' property. If 'false', responses shall return the
	// 'Payload' property. If this property is not present when the task is
	// created, the default is 'false'. This property shall be supported if the
	// 'Payload' property is supported.
	//
	// Version added: v1.3.0
	HidePayload bool
	// Messages shall contain an array of messages associated with the task.
	Messages []Message
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Payload shall contain information detailing the HTTP and JSON request
	// payload information for executing this task. This property shall not be
	// included in the response if the 'HidePayload' property is 'true'.
	//
	// Version added: v1.3.0
	Payload Payload
	// PercentComplete shall indicate the completion progress of the task, reported
	// in percent of completion, '0' to '100'. If the task has not been started,
	// the value shall be zero.
	//
	// Version added: v1.4.0
	PercentComplete *uint `json:",omitempty"`
	// StartTime shall indicate the date and time when the task was started.
	StartTime string
	// SubTasks shall contain a link to a resource collection of type
	// 'TaskCollection'. This property shall not be present if this resource
	// represents a sub-task for a task.
	//
	// Version added: v1.5.0
	subTasks string
	// TaskMonitor shall contain a URI to task monitor as defined in the Redfish
	// Specification.
	//
	// Version added: v1.2.0
	TaskMonitor string
	// TaskState shall indicate the state of the task.
	TaskState TaskState
	// TaskStatus shall contain the completion status of the task and shall not be
	// set until the task completes. This property should contain 'Critical' if one
	// or more messages in the 'Messages' array contains the severity 'Critical'.
	// This property should contain 'Warning' if one or more messages in the
	// 'Messages' array contains the severity 'Warning' and if no messages contain
	// the severity 'Critical'. This property should contain 'OK' if all messages
	// in the 'Messages' array contain the severity 'OK' or if the array is empty.
	TaskStatus Health
	// createdResources are the URIs for CreatedResources.
	createdResources []string
}

// UnmarshalJSON unmarshals a Task object from the raw JSON.
func (t *Task) UnmarshalJSON(b []byte) error {
	type temp Task
	type tLinks struct {
		CreatedResources Links `json:"CreatedResources"`
	}
	var tmp struct {
		temp
		Links    tLinks
		SubTasks Link `json:"SubTasks"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = Task(tmp.temp)

	// Extract the links to other entities for later
	t.createdResources = tmp.Links.CreatedResources.ToStrings()
	t.subTasks = tmp.SubTasks.String()

	return nil
}

// GetTask will get a Task instance from the service.
func GetTask(c Client, uri string) (*Task, error) {
	return GetObject[Task](c, uri)
}

// ListReferencedTasks gets the collection of Task from
// a provided reference.
func ListReferencedTasks(c Client, link string) ([]*Task, error) {
	return GetCollectionObjects[Task](c, link)
}

// CreatedResources gets the CreatedResources linked resources.
func (t *Task) CreatedResources() ([]*Entity, error) {
	return GetObjects[Entity](t.client, t.createdResources)
}

// SubTasks gets the SubTasks collection.
func (t *Task) SubTasks() ([]*Task, error) {
	if t.subTasks == "" {
		return nil, nil
	}
	return GetCollectionObjects[Task](t.client, t.subTasks)
}

// Payload shall contain information detailing the HTTP and JSON payload
// information for executing this task.
type Payload struct {
	// HTTPHeaders shall contain an array of HTTP headers that this task includes.
	//
	// Version added: v1.3.0
	HTTPHeaders []string `json:"HttpHeaders"`
	// HTTPOperation shall contain the HTTP operation to execute for this task.
	//
	// Version added: v1.3.0
	HTTPOperation string `json:"HttpOperation"`
	// JSONBody shall contain the JSON-formatted payload used for this task.
	//
	// Version added: v1.3.0
	JSONBody string `json:"JsonBody"`
	// TargetURI shall contain a link to the location to use as the target for an
	// HTTP operation.
	//
	// Version added: v1.3.0
	TargetURI string `json:"TargetUri"`
}
