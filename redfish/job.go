//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// JobState indicates the state of a job.
type JobState string

const (
	// NewJobState shall represent that this job is newly created but the
	// operation has not yet started.
	NewJobState JobState = "New"
	// StartingJobState shall represent that the operation is starting.
	StartingJobState JobState = "Starting"
	// RunningJobState shall represent that the operation is executing.
	RunningJobState JobState = "Running"
	// SuspendedJobState shall represent that the operation has been
	// suspended but is expected to restart and is therefore not complete.
	SuspendedJobState JobState = "Suspended"
	// InterruptedJobState shall represent that the operation has been
	// interrupted but is expected to restart and is therefore not complete.
	InterruptedJobState JobState = "Interrupted"
	// PendingJobState shall represent that the operation is pending some
	// condition and has not yet begun to execute.
	PendingJobState JobState = "Pending"
	// StoppingJobState shall represent that the operation is stopping but is
	// not yet complete.
	StoppingJobState JobState = "Stopping"
	// CompletedJobState shall represent that the operation completed
	// successfully or with warnings.
	CompletedJobState JobState = "Completed"
	// CancelledJobState shall represent that the operation completed because
	// the job was cancelled by an operator.
	CancelledJobState JobState = "Cancelled"
	// ExceptionJobState shall represent that the operation completed with
	// errors.
	ExceptionJobState JobState = "Exception"
	// ServiceJobState shall represent that the operation is now running as a
	// service and expected to continue operation until stopped or killed.
	ServiceJobState JobState = "Service"
	// UserInterventionJobState shall represent that the operation is waiting
	// for a user to intervene and needs to be manually continued, stopped,
	// or cancelled.
	UserInterventionJobState JobState = "UserIntervention"
	// ContinueJobState shall represent that the operation has been resumed
	// from a paused condition and should return to a Running state.
	ContinueJobState JobState = "Continue"
)

// JobPayload shall contain information detailing the HTTP and JSON payload
// information for executing this Job.
type JobPayload struct {
	// HTTPHeaders shall contain an array of HTTP headers in this Job.
	HTTPHeaders []string `json:"HttpHeaders"`
	// HTTPOperation shall contain the HTTP operation that executes this Job.
	HTTPOperation string `json:"HttpOperation"`
	// JsonBody shall contain JSON-formatted payload for this Job.
	JSONBody string `json:"JsonBody"`
	// TargetUri shall contain link to a target location for an HTTP operation.
	TargetURI string `json:"TargetUri"`
}

// Job shall contain a job in a Redfish implementation.
type Job struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CreatedBy shall contain the user name, software program name,
	// or other identifier indicating the creator of this job.
	CreatedBy string
	// Description provides a description of this resource.
	Description string
	// EndTime shall indicate the date and time when the job was completed.
	// This property shall not appear if the job is running or was not
	// completed.  This property shall appear only if the JobState is
	// Completed, Cancelled, or Exception.
	EndTime string
	// EstimatedDuration shall indicate the estimated total time needed to
	// complete the job. The value is not expected to change while the job is
	// in progress, but the service may update the value if it obtains new
	// information that significantly changes the expected duration.
	// Services should be conservative in the reported estimate and clients
	// should treat this value as an estimate.
	EstimatedDuration string
	// HidePayload shall indicate whether the contents of the payload should
	// be hidden from view after the job has been created. If 'true',
	// responses shall not return the Payload property. If 'false',
	// responses shall return the Payload property.
	// If this property is not present when the job is created,
	// the default is 'false'.
	HidePayload bool
	// JobState shall indicate the state of the job.
	JobState JobState
	// JobStatus shall indicate the health status of the job.
	// This property should contain 'Critical' if one or more messages in the
	// Messages array contains the severity 'Critical'.
	// This property should contain 'Warning' if one or more messages in the
	// Messages array contains the severity 'Warning' and no messages contain
	// the severity 'Critical'.
	// This property should contain 'OK' if all messages in the Messages
	// array contain the severity 'OK' or the array is empty.
	JobStatus common.Health
	// MaxExecutionTime shall be an ISO 8601 conformant duration describing
	// the maximum duration the job is allowed to execute before being
	// stopped by the service.
	MaxExecutionTime string
	// Messages shall contain an array of messages associated with the job.
	Messages []common.Message
	// Payload shall contain the HTTP and JSON payload information for
	// executing this job. This property shall not be included in the
	// response if the HidePayload property is 'true'.
	Payload JobPayload
	// PercentComplete shall indicate the completion progress of the job,
	// reported in percent of completion. If the job has not been started,
	// the value shall be zero.
	PercentComplete int
	// Schedule shall contain the scheduling details for this job and the
	// recurrence frequency for future instances of this job.
	Schedule common.Schedule
	// StartTime shall indicate the date and time when the job was last
	// started or is scheduled to start.
	StartTime string
	// StepOrder shall contain an array of IDs for the job steps in the order
	// that they shall be executed.
	// Each step shall be completed prior to the execution of the next step
	// in array order. An incomplete list of steps shall be considered an
	// invalid configuration. If this property is not present or contains an
	// empty array it shall indicate that the step execution order is omitted
	// and may occur in parallel or in series as determined by the service.
	StepOrder []string
	// Steps shall contain the link to a resource collection of type
	// JobCollection. This property shall not be present if this resource
	// represents a step for a job.
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

	*job = Job(t.temp)

	// Extract the links to other entities for later
	job.steps = t.Steps.String()
	return nil
}

// Steps gets the collection of steps for this job.
func (job *Job) Steps() ([]*Job, error) {
	return ListReferencedJobs(job.GetClient(), job.steps)
}

// GetJob will get a Job instance from the service.
func GetJob(c common.Client, uri string) (*Job, error) {
	return common.GetObject[Job](c, uri)
}

// ListReferencedJobs gets the collection of Job from
// a provided reference.
func ListReferencedJobs(c common.Client, link string) ([]*Job, error) {
	return common.GetCollectionObjects[Job](c, link)
}
