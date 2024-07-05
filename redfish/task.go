//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// TaskState indicates the state of a task.
type TaskState string

const (

	// NewTaskState shall represent that this task is newly created but the
	// operation has not yet started.
	NewTaskState TaskState = "New"
	// StartingTaskState shall represent that the operation is starting.
	StartingTaskState TaskState = "Starting"
	// RunningTaskState shall represent that the operation is executing.
	RunningTaskState TaskState = "Running"
	// SuspendedTaskState shall represent that the operation has been
	// suspended but is expected to restart and is therefore not complete.
	SuspendedTaskState TaskState = "Suspended"
	// InterruptedTaskState shall represent that the operation has been
	// interrupted but is expected to restart and is therefore not complete.
	InterruptedTaskState TaskState = "Interrupted"
	// PendingTaskState shall represent that the operation is pending some
	// condition and has not yet begun to execute.
	PendingTaskState TaskState = "Pending"
	// StoppingTaskState shall represent that the operation is stopping but
	// is not yet complete.
	StoppingTaskState TaskState = "Stopping"
	// CompletedTaskState shall represent that the operation is complete and
	// completed successfully or with warnings.
	CompletedTaskState TaskState = "Completed"
	// KilledTaskState shall represent that the operation is complete because
	// the task was killed by an operator.
	KilledTaskState TaskState = "Killed"
	// ExceptionTaskState shall represent that the operation is complete and
	// completed with errors.
	ExceptionTaskState TaskState = "Exception"
	// ServiceTaskState shall represent that the operation is now running as
	// a service and expected to continue operation until stopped or killed.
	ServiceTaskState TaskState = "Service"
	// CancellingTaskState shall represent that the operation is in the
	// process of being cancelled.
	CancellingTaskState TaskState = "Cancelling"
	// CancelledTaskState shall represent that the operation was cancelled
	// either through a Delete on a Task Monitor or Task Resource or by an
	// internal process.
	CancelledTaskState TaskState = "Cancelled"
)

// Payload shall contain information detailing the HTTP
// and JSON payload information for executing this Task.
type Payload struct {
	// HTTPHeaders is used in the execution of this Task.
	HTTPHeaders []string `json:"HttpHeaders"`
	// HTTPOperation shall contain the HTTP operation to
	// execute for this Task.
	HTTPOperation string `json:"HttpOperation"`
	// JSONBody shall contain the JSON-formatted payload used for this task.
	JSONBody string `json:"JsonBody"`
	// TargetURI is used as the target for an HTTP operation.
	TargetURI string `json:"TargetUri"`
}

// Task is used to represent a Task for a Redfish implementation.
type Task struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// EndTime shall indicate the time the task was completed.
	EndTime string
	// HidePayload shall be set to True if the Payload object shall not be
	// returned on GET operations, and set to False if the contents can be
	// returned normally. If this property is not specified when the Task is
	// created, the default value shall be False.
	HidePayload bool
	Links       struct {
		// CreatedResources are an array of resource IDs created by this task.
		CreatedResources []string
		// CreatedResourcesCount is the number of created resources.
		CreatedResourcesCount int `json:"CreatedResources@odata.count"`
	}
	// Messages shall be an array of messages associated with the task.
	Messages []common.Message
	// Payload shall contain information detailing the HTTP and JSON payload
	// information for executing this task. This object shall not be included in
	// the response if the HidePayload property is set to True.
	Payload Payload
	// PercentComplete shall indicate the completion progress of the task,
	// reported in percent of completion. If the task has not been started, the
	// value shall be zero.
	PercentComplete int
	// StartTime shall indicate the time the task was started.
	StartTime string
	// SubTasks shall contain a link to a resource collection of type TaskCollection. This property shall not be
	// present if this resource represents a sub-task for a task.
	subTasks []string
	// TaskMonitor shall contain a URI to Task Monitor as defined in the Redfish
	// Specification.
	TaskMonitor string
	// TaskState is used to indicate that the task is a new task which has
	// just been instantiated and is in the initial state and indicates it
	// has never been started.  Starting shall be used to indicate that the
	// task is moving from the New, Suspended, or Service states into the
	// Running state.  Running shall be used to indicate that the Task is
	// running.  Suspended shall be used to indicate  that the Task is
	// stopped (e.g., by a user), but can be restarted in a seamless manner.
	// Interrupted shall be used to indicate  that the Task was interrupted
	// (e.g., by a server crash) in the middle of processing, and the user
	// should either re-run/restart the Task.  Pending shall be used to
	// indicate  that the Task has been queued and will be scheduled for
	// processing as soon as resources are available to handle the request.
	// Stopping shall be used to indicate that the Task is in the process of
	// moving to a Completed, Killed, or Exception state.  Completed shall be
	// used to indicate that the task has completed normally.  Killed shall
	// be used to indicate  that the task has been stopped by a Kill state
	// change request (non-graceful shutdown).  Exception shall be used to
	// indicate  that the Task is in an abnormal state that might be
	// indicative of an error condition.  Service shall be used to indicate
	// that the Task is in a state that supports problem discovery, or
	// resolution, or both.  This state is used when a corrective action is
	// possible.
	TaskState TaskState
	// TaskStatus shall be the completion status of the task, as defined in the
	// Status section of the Redfish specification and shall not be set until
	// the task has completed.
	TaskStatus common.Health
	// Oem property contains OEM specific task information
	Oem json.RawMessage `json:"Oem,omitempty"`
}

// UnmarshalJSON unmarshals a Task object from the raw JSON.
func (task *Task) UnmarshalJSON(b []byte) error {
	type temp Task
	var t struct {
		temp
		SubTasks common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*task = Task(t.temp)
	task.subTasks = t.SubTasks.ToStrings()

	return nil
}

// SubTasks gets the sub-tasks for this task.
// This property shall not be present if this resource represents a sub-task for a task.
func (task *Task) SubTasks() ([]*Task, error) {
	var result []*Task

	collectionError := common.NewCollectionError()
	for _, uri := range task.subTasks {
		item, err := GetTask(task.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// GetTask will get a Task instance from the service.
func GetTask(c common.Client, uri string) (*Task, error) {
	return common.GetObject[Task](c, uri)
}

// ListReferencedTasks gets the collection of Task from
// a provided reference.
func ListReferencedTasks(c common.Client, link string) ([]*Task, error) {
	return common.GetCollectionObjects[Task](c, link)
}
