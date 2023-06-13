//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

// JobState indicates the state of a job.
type JobState string

const (

	// CancelledJobState This value shall represent that the operation completed because the job was cancelled by an operator.
	CancelledJobState JobState = "Cancelled"
	// CompletedJobState This value shall represent that the operation completed successfully or with warnings.
	CompletedJobState JobState = "Completed"
	// ContinueJobState This value shall represent that the operation has been resumed from a paused condition and should return to a Running state.
	ContinueJobState JobState = "Continue"
	// ExceptionJobState This value shall represent that the operation completed with errors.
	ExceptionJobState JobState = "Exception"
	// InterruptedJobState This value shall represent that the operation has been interrupted but is expected to restart and is therefore not complete.
	InterruptedJobState JobState = "Interrupted"
	// NewJobState This value shall represent that this job is newly created but the operation has not yet started.
	NewJobState JobState = "New"
	// PendingJobState This value shall represent that the operation is pending some condition and has not yet begun to execute.
	PendingJobState JobState = "Pending"
	// RunningJobState This value shall represent that the operation is executing.
	RunningJobState JobState = "Running"
	// ServiceJobState This value shall represent that the operation is now running as a service and expected to continue operation until stopped or killed.
	ServiceJobState JobState = "Service"
	// StartingJobState This value shall represent that the operation is starting.
	StartingJobState JobState = "Starting"
	// StoppingJobState This value shall represent that the operation is stopping but is not yet complete.
	StoppingJobState JobState = "Stopping"
	// SuspendedJobState This value shall represent that the operation has been suspended but is expected to restart and is therefore not complete.
	SuspendedJobState JobState = "Suspended"
	// UserInterventionJobState This value shall represent that the operation is waiting for a user to intervene and needs to be manually continued, stopped, or cancelled.
	UserInterventionJobState JobState = "UserIntervention"
)

// JobPayload shall contain information detailing the HTTP
// and JSON payload information for executing this Job.
type JobPayload struct {
	// HTTPHeaders is used in the execution of this Job.
	HTTPHeaders []string `json:"HttpHeaders"`
	// HTTPOperation shall contain the HTTP operation to
	// execute for this Job.
	HTTPOperation string `json:"HttpOperation"`
	// JSONBody is used for this Job.
	JSONBody string `json:"JsonBody"`
	// TargetURI is used as the target for an HTTP operation.
	TargetURI string `json:"TargetUri"`
}

type Job struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// CreatedBy The person or program that created this job entry.
	CreatedBy string
	// EndTime The date and time when the job was completed.
	EndTime string
	// EstimatedDuration The estimated total time required to complete the job.
	EstimatedDuration string
	// HidePayload An indication of whether the contents of the payload should be hidden from view after the job has been created.
	// If true, responses do not return the payload. If false, responses return the payload.
	// If this property is not present when the job is created, the default is false.
	HidePayload bool
	// JobState This property shall indicate the state of the job.
	JobState JobState
	// JobStatus This property shall indicate the state of the job.
	JobStatus common.Health
	// MaxExecutionTime The maximum amount of time the job is allowed to execute.
	MaxExecutionTime string
	// Messages shall be an array of messages associated with the job.
	Messages []common.Message
	// Payload shall contain information detailing the HTTP and JSON payload information for executing this job. This object shall not be included in the response if the HidePayload property is set to True.
	Payload JobPayload
	// PercentComplete This property shall indicate the completion progress of the job, reported in percent of completion. If the job has not been started, the value shall be zero.
	PercentComplete int
	// Schedule This object shall contain the scheduling details for this job and the recurrence frequency for future instances of this job.
	Schedule common.Schedule
	// StartTime This property shall indicate the date and time when the job was last started or is scheduled to start.
	StartTime string
	// StepOrder This property shall contain an array of IDs for the job steps in the order that they shall be executed. Each step shall be completed prior to the execution of the next step in array order. An incomplete list of steps shall be considered an invalid configuration. If this property is not present or contains an empty array it shall indicate that the step execution order is omitted and may occur in parallel or in series as determined by the service.
	StepOrder []string
	// Steps This property shall contain the link to a resource collection of type JobCollection. This property shall not be present if this resource represents a step for a job.
	steps string
}

// UnmarshalJSON unmarshals a Job object from the raw JSON.
func (job *Job) UnmarshalJSON(b []byte) error {
	type temp Job
	var t struct {
		temp
		Steps common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*job = Job(t.temp)
	job.steps = t.Steps.String()

	return nil
}

// Delete deletes a specific job using the job service.
func (job *Job) Delete() error {
	return DeleteJob(job.GetClient(), job.ODataID)
}

// DeleteJob will delete a Job.
func DeleteJob(c common.Client, uri string) error {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}

	resp, err := c.Delete(uri)
	if err == nil {
		defer resp.Body.Close()
	}

	return err
}
