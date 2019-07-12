// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// TaskState indicates the state of a task
type TaskState string

const (
	// CancelledTaskState indicates a task has been cancelled by an operator or
	// internal process.
	CancelledTaskState TaskState = "Cancelled"
	// CancellingTaskState indicates a task is in the process of being cancelled.
	CancellingTaskState TaskState = "Cancelling"
	// CompletedTaskState indicates a task has completed.
	CompletedTaskState TaskState = "Completed"
	// ExceptionTaskState indicates a task has stopped due to an exception condition.
	ExceptionTaskState TaskState = "Exception"
	// InterruptedTaskState indicates a task has been interrupted.
	InterruptedTaskState TaskState = "Interrupted"
	// KilledTaskState indicates a task was terminated.
	KilledTaskState TaskState = "Killed"
	// NewTaskState indicates it is a new task.
	NewTaskState TaskState = "New"
	// PendingTaskState indicates a task is pending and has not started.
	PendingTaskState TaskState = "Pending"
	// RunningTaskState indicates a task is running normally.
	RunningTaskState TaskState = "Running"
	// ServiceTaskState indicates a task is running as a service.
	ServiceTaskState TaskState = "Service"
	// StartingTaskState indicates a task is starting.
	StartingTaskState TaskState = "Starting"
	// StoppingTaskState indicates a task is in the process of stopping.
	StoppingTaskState TaskState = "Stopping"
	// SuspendedTaskState indicates a task has been suspended.
	SuspendedTaskState TaskState = "Suspended"
)

// Task contains information about a specific Task scheduled by or being
// executed by a Redfish service's Task Service.
type Task struct {
	common.Entity
	Modified   string
	TaskState  TaskState
	StartTime  string
	EndTime    string
	TaskStatus common.Health
}

// GetTask will get a Task instance from the Redfish service.
func GetTask(c common.Client, uri string) (*Task, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t Task
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedTasks gets the collection of Tasks
func ListReferencedTasks(c common.Client, link string) ([]*Task, error) {
	var result []*Task
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, tLink := range links.ItemLinks {
		t, err := GetTask(c, tLink)
		if err != nil {
			return result, err
		}
		result = append(result, t)
	}

	return result, nil
}
