//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #Job.v1_3_0.Job

package schemas

import (
	"encoding/json"
)

type JobState string

const (
	// NewJobState shall indicate that this job is newly created but the operation
	// has not yet started. This shall be the initial state for document-based
	// jobs. Upon receiving the 'Validate' action, or if the value of
	// 'ValidationPolicy' in the 'JobService' resource contains 'Automatic', the
	// document-based job's 'JobState' shall transition to 'Validating'. If the
	// 'ValidationPolicy' property in the 'JobService' resource contains 'Bypass',
	// the 'JobState' for a document-based job shall transition to 'Pending'.
	NewJobState JobState = "New"
	// StartingJobState shall indicate that the operation is starting.
	StartingJobState JobState = "Starting"
	// RunningJobState shall indicate that the operation is executing. Jobs that
	// complete successfully shall transition from this state to the 'Completed'
	// state. Jobs that do not complete successfully shall transition from this
	// state to the 'Exception' state.
	RunningJobState JobState = "Running"
	// SuspendedJobState shall indicate that the operation has been suspended but
	// is expected to restart and is therefore not complete. To resume a job,
	// perform the 'Resume' action.
	SuspendedJobState JobState = "Suspended"
	// InterruptedJobState shall indicate that the operation has been interrupted
	// but is expected to restart and is therefore not complete.
	InterruptedJobState JobState = "Interrupted"
	// PendingJobState shall indicate that the operation is pending some condition
	// and has not yet begun to execute.
	PendingJobState JobState = "Pending"
	// StoppingJobState shall indicate that the operation is stopping but is not
	// yet complete.
	StoppingJobState JobState = "Stopping"
	// CompletedJobState shall indicate that the operation completed successfully
	// or with warnings. The job may restart in the future based on the scheduling
	// configuration of the job or operations performed by a user.
	CompletedJobState JobState = "Completed"
	// CancelledJobState shall indicate that the operation completed because the
	// job was cancelled by an operator. The job may restart in the future based on
	// the scheduling configuration of the job or operations performed by a user.
	CancelledJobState JobState = "Cancelled"
	// ExceptionJobState shall indicate that the operation completed with errors.
	// The job may restart in the future based on the scheduling configuration of
	// the job or operations performed by a user.
	ExceptionJobState JobState = "Exception"
	// ServiceJobState shall indicate that the operation is now running as a
	// service and expected to continue operation until stopped or killed.
	ServiceJobState JobState = "Service"
	// UserInterventionJobState shall indicate that the operation is waiting for a
	// user to intervene and needs to be manually continued, stopped, or cancelled.
	UserInterventionJobState JobState = "UserIntervention"
	// ContinueJobState shall indicate that the operation has been resumed from a
	// paused condition and should return to a Running state.
	ContinueJobState JobState = "Continue"
	// ValidatingJobState shall indicate that the document-based job is validating
	// the state of the system to determine if it can run. For example, a job that
	// runs on factory equipment might check to ensure that the equipment is
	// properly configured and has sufficient ingredients to run the job. If the
	// validation checks fail, the job shall transition to the 'Invalid' state. If
	// the validation checks are successful, the job shall transition to the
	// 'Pending' state.
	ValidatingJobState JobState = "Validating"
	// InvalidJobState shall indicate that validation has determined that the
	// system is not properly configured to run the document-based job. To perform
	// validation checks again, perform the 'Validate' action.
	InvalidJobState JobState = "Invalid"
)

type JobType string

const (
	// DocumentBasedJobType shall indicate a job that was created by performing the
	// 'SubmitJob' action on a 'JobDocument' resource.
	DocumentBasedJobType JobType = "DocumentBased"
	// UserSpecifiedJobType shall indicate a job that was created by performing an
	// HTTP 'POST' operation on a 'JobCollection' resource.
	UserSpecifiedJobType JobType = "UserSpecified"
	// ServiceGeneratedJobType shall indicate a job that was created automatically
	// by the service as part of its internal policies.
	ServiceGeneratedJobType JobType = "ServiceGenerated"
)

// Job shall contain a job in a Redfish implementation.
type Job struct {
	Entity
	// CreatedBy shall contain the username, software program name, or other
	// identifier indicating the creator of this job.
	CreatedBy string
	// CreationTime shall contain the date and time when the job was created.
	//
	// Version added: v1.3.0
	CreationTime string
	// EndTime shall contain the date and time when the job was completed. This
	// property shall not appear if the job is running or was not completed. This
	// property shall appear only if the 'JobState' is 'Completed', 'Cancelled', or
	// 'Exception'.
	EndTime string
	// EstimatedCompletionTime shall contain the date and time when the job is
	// expected to complete. If the 'EstimatedDuration' property is supported, the
	// value of this property shall contain the summation of the 'StartTime'
	// property and the 'EstimatedDuration' property.
	//
	// Version added: v1.3.0
	EstimatedCompletionTime string
	// EstimatedDuration shall contain the estimated total time needed to complete
	// the job. The value is not expected to change while the job is in progress,
	// but the service may update the value if it obtains new information that
	// significantly changes the expected duration. Services should be conservative
	// in the reported estimate and clients should treat this value as an estimate.
	//
	// Version added: v1.1.0
	EstimatedDuration string
	// HidePayload shall indicate whether the contents of the payload should be
	// hidden from view after the job has been created. If 'true', responses shall
	// not return the 'Payload' or 'Parameters' properties. If 'false', responses
	// shall return the 'Payload' or 'Parameters' properties. If this property is
	// not present when the job is created, the default is 'false'.
	HidePayload bool
	// JobPriority shall contain the requested priority of this job. The value '0'
	// shall indicate the highest priority. Increasing values shall represent
	// decreasing priority.
	//
	// Version added: v1.3.0
	JobPriority *uint `json:",omitempty"`
	// JobState shall contain the state of the job.
	JobState JobState
	// JobStatus shall contain the health status of the job. This property should
	// contain 'Critical' if one or more messages in the 'Messages' array contains
	// the severity 'Critical'. This property should contain 'Warning' if one or
	// more messages in the 'Messages' array contains the severity 'Warning' and if
	// no messages contain the severity 'Critical'. This property should contain
	// 'OK' if all messages in the 'Messages' array contain the severity 'OK' or if
	// the array is empty.
	JobStatus Health
	// JobType shall contain the type of this job.
	//
	// Version added: v1.3.0
	JobType JobType
	// MaxExecutionTime shall contain the maximum duration the job is allowed to
	// execute before being stopped by the service.
	MaxExecutionTime string
	// Messages shall contain an array of messages associated with the job.
	Messages []Message
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Parameters shall contain the parameters specified for running this
	// document-based job. This property shall only be present for document-based
	// jobs and if the 'HidePayload' property is 'false'.
	//
	// Version added: v1.3.0
	Parameters map[string]any
	// Payload shall contain the HTTP and JSON request payload information for
	// executing this user-specified job. This property shall only be present for
	// user-specified jobs and if the 'HidePayload' property is 'false'.
	Payload JobPayload
	// PercentComplete shall contain the completion progress of the job, reported
	// in percent of completion, '0' to '100'. If the job has not been started, the
	// value shall be zero.
	PercentComplete *uint `json:",omitempty"`
	// Schedule shall contain the scheduling details for this job and the
	// recurrence frequency for future instances of this job. This property shall
	// not be present for document-based jobs.
	Schedule Schedule
	// StartTime shall contain the date and time when the job was last started or
	// is scheduled to start.
	StartTime string
	// StepOrder shall contain an array of 'Id' property values for the job steps
	// in the order that they shall be executed. Each step shall be completed prior
	// to the execution of the next step in array order. An incomplete list of
	// steps shall be considered an invalid configuration. If this property is not
	// present or contains an empty array it shall indicate that the step execution
	// order is omitted and may occur in parallel or in series as determined by the
	// service. This property shall not be present for document-based jobs.
	StepOrder []string
	// Steps shall contain the link to a resource collection of type
	// 'JobCollection'. This property shall not be present if this resource
	// represents a step for a job. This property shall not be present for
	// document-based jobs.
	steps string
	// cancelTarget is the URL to send Cancel requests.
	cancelTarget string
	// forceStartTarget is the URL to send ForceStart requests.
	forceStartTarget string
	// invalidateTarget is the URL to send Invalidate requests.
	invalidateTarget string
	// resubmitTarget is the URL to send Resubmit requests.
	resubmitTarget string
	// resumeTarget is the URL to send Resume requests.
	resumeTarget string
	// suspendTarget is the URL to send Suspend requests.
	suspendTarget string
	// validateTarget is the URL to send Validate requests.
	validateTarget string
	// createdResources are the URIs for CreatedResources.
	createdResources []string
	// executor is the URI for Executor.
	executor string
	// jobDocument is the URI for JobDocument.
	jobDocument string
	// parentJob is the URI for ParentJob.
	parentJob string
	// preferredExecutors are the URIs for PreferredExecutors.
	preferredExecutors []string
	// subsidiaryJobs are the URIs for SubsidiaryJobs.
	subsidiaryJobs []string
	// validatedExecutors are the URIs for ValidatedExecutors.
	validatedExecutors []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Job object from the raw JSON.
func (j *Job) UnmarshalJSON(b []byte) error {
	type temp Job
	type jActions struct {
		Cancel     ActionTarget `json:"#Job.Cancel"`
		ForceStart ActionTarget `json:"#Job.ForceStart"`
		Invalidate ActionTarget `json:"#Job.Invalidate"`
		Resubmit   ActionTarget `json:"#Job.Resubmit"`
		Resume     ActionTarget `json:"#Job.Resume"`
		Suspend    ActionTarget `json:"#Job.Suspend"`
		Validate   ActionTarget `json:"#Job.Validate"`
	}
	type jLinks struct {
		CreatedResources   Links `json:"CreatedResources"`
		Executor           Link  `json:"Executor"`
		JobDocument        Link  `json:"JobDocument"`
		ParentJob          Link  `json:"ParentJob"`
		PreferredExecutors Links `json:"PreferredExecutors"`
		SubsidiaryJobs     Links `json:"SubsidiaryJobs"`
		ValidatedExecutors Links `json:"ValidatedExecutors"`
	}
	var tmp struct {
		temp
		Actions jActions
		Links   jLinks
		Steps   Link `json:"Steps"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*j = Job(tmp.temp)

	// Extract the links to other entities for later
	j.cancelTarget = tmp.Actions.Cancel.Target
	j.forceStartTarget = tmp.Actions.ForceStart.Target
	j.invalidateTarget = tmp.Actions.Invalidate.Target
	j.resubmitTarget = tmp.Actions.Resubmit.Target
	j.resumeTarget = tmp.Actions.Resume.Target
	j.suspendTarget = tmp.Actions.Suspend.Target
	j.validateTarget = tmp.Actions.Validate.Target
	j.createdResources = tmp.Links.CreatedResources.ToStrings()
	j.executor = tmp.Links.Executor.String()
	j.jobDocument = tmp.Links.JobDocument.String()
	j.parentJob = tmp.Links.ParentJob.String()
	j.preferredExecutors = tmp.Links.PreferredExecutors.ToStrings()
	j.subsidiaryJobs = tmp.Links.SubsidiaryJobs.ToStrings()
	j.validatedExecutors = tmp.Links.ValidatedExecutors.ToStrings()
	j.steps = tmp.Steps.String()

	// This is a read/write object, so we need to save the raw object data for later
	j.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (j *Job) Update() error {
	readWriteFields := []string{
		"JobState",
		"MaxExecutionTime",
	}

	return j.UpdateFromRawData(j, j.RawData, readWriteFields)
}

// GetJob will get a Job instance from the service.
func GetJob(c Client, uri string) (*Job, error) {
	return GetObject[Job](c, uri)
}

// ListReferencedJobs gets the collection of Job from
// a provided reference.
func ListReferencedJobs(c Client, link string) ([]*Job, error) {
	return GetCollectionObjects[Job](c, link)
}

// This action shall cancel the job if it is currently in the 'Running',
// 'Invalid', or 'Suspended' states. The job shall transition to 'Cancelled' if
// the action is successful.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Cancel() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.cancelTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall force the job to start running if it is in the 'Pending'
// state. The job shall transition to 'Running' if the action is successful.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) ForceStart() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.forceStartTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall invalidate the job if it is in the 'Pending' state. The
// job shall transition to 'Invalid' if the action is successful.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Invalidate() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.invalidateTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall resubmit a job to the job service. The new job shall be
// based on the job document, and properties of the resource associated with
// this action. Services shall take appropriate measures to make sure that
// appropriate security is maintained - for instance, only allowing the same
// user that created the job to resubmit it.
// startTime - This parameter shall contain the time to start the job.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Resubmit(startTime string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["StartTime"] = startTime
	resp, taskInfo, err := PostWithTask(j.client,
		j.resubmitTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall resume the job if it is in the 'Suspended' state. The job
// shall transition to 'Running' if the action is successful.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Resume() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.resumeTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall suspend the job if it is in the 'Running' state. The job
// shall transition to 'Suspended' if the action is successful.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Suspend() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.suspendTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall request the validation the job if it is in the 'New'
// state. The job shall transition to 'Pending' if the action is successful. If
// the job is not valid, it shall transition to 'Invalid'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *Job) Validate() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(j.client,
		j.validateTarget, payload, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// CreatedResources gets the CreatedResources linked resources.
func (j *Job) CreatedResources() ([]*Entity, error) {
	return GetObjects[Entity](j.client, j.createdResources)
}

// Executor gets the Executor linked resource.
func (j *Job) Executor() (*JobExecutor, error) {
	if j.executor == "" {
		return nil, nil
	}
	return GetObject[JobExecutor](j.client, j.executor)
}

// JobDocument gets the JobDocument linked resource.
func (j *Job) JobDocument() (*JobDocument, error) {
	if j.jobDocument == "" {
		return nil, nil
	}
	return GetObject[JobDocument](j.client, j.jobDocument)
}

// ParentJob gets the ParentJob linked resource.
func (j *Job) ParentJob() (*Job, error) {
	if j.parentJob == "" {
		return nil, nil
	}
	return GetObject[Job](j.client, j.parentJob)
}

// PreferredExecutors gets the PreferredExecutors linked resources.
func (j *Job) PreferredExecutors() ([]*JobExecutor, error) {
	return GetObjects[JobExecutor](j.client, j.preferredExecutors)
}

// SubsidiaryJobs gets the SubsidiaryJobs linked resources.
func (j *Job) SubsidiaryJobs() ([]*Job, error) {
	return GetObjects[Job](j.client, j.subsidiaryJobs)
}

// ValidatedExecutors gets the ValidatedExecutors linked resources.
func (j *Job) ValidatedExecutors() ([]*JobExecutor, error) {
	return GetObjects[JobExecutor](j.client, j.validatedExecutors)
}

// Steps gets the Steps collection.
func (j *Job) Steps() ([]*Job, error) {
	if j.steps == "" {
		return nil, nil
	}
	return GetCollectionObjects[Job](j.client, j.steps)
}

// JobExcerpt shall contain a job in a Redfish implementation.
type JobExcerpt struct {
	// EstimatedCompletionTime shall contain the date and time when the job is
	// expected to complete. If the 'EstimatedDuration' property is supported, the
	// value of this property shall contain the summation of the 'StartTime'
	// property and the 'EstimatedDuration' property.
	//
	// Version added: v1.3.0
	EstimatedCompletionTime string
	// JobState shall contain the state of the job.
	JobState JobState
	// PercentComplete shall contain the completion progress of the job, reported
	// in percent of completion, '0' to '100'. If the job has not been started, the
	// value shall be zero.
	PercentComplete *uint `json:",omitempty"`
}

// JobPayload shall contain information detailing the HTTP and JSON payload
// information for executing this job.
type JobPayload struct {
	// HTTPHeaders shall contain an array of HTTP headers in this job.
	HTTPHeaders []string `json:"HttpHeaders"`
	// HTTPOperation shall contain the HTTP operation that executes this job.
	HTTPOperation string `json:"HttpOperation"`
	// JSONBody shall contain JSON-formatted payload for this job.
	JSONBody string `json:"JsonBody"`
	// TargetURI shall contain link to a target location for an HTTP operation.
	TargetURI string `json:"TargetUri"`
}
