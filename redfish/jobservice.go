//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #JobService.v1_1_0.JobService

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type ValidationPolicy string

const (
	// AutomaticValidationPolicy shall indicate that jobs are validated
	// automatically. The 'JobState' property of a new job shall contain
	// 'Validating' when created with the 'SubmitJob' action from the 'JobDocument'
	// resource.
	AutomaticValidationPolicy ValidationPolicy = "Automatic"
	// ManualValidationPolicy shall indicate jobs are validated manually. The
	// 'JobState' property of a new job shall contain 'New' when created with the
	// 'SubmitJob' action from the 'JobDocument' resource and wait for a user to
	// perform the 'Validate' action in the 'Job' resource.
	ManualValidationPolicy ValidationPolicy = "Manual"
	// BypassValidationPolicy shall indicate jobs are not validated. The 'JobState'
	// property of a new job shall contain 'Pending' when created with the
	// 'SubmitJob' action from the 'JobDocument' resource.
	BypassValidationPolicy ValidationPolicy = "Bypass"
)

// JobService shall represent a job service for a Redfish implementation.
type JobService struct {
	common.Entity
	// DateTime shall contain the current date and time setting for the job
	// service.
	DateTime string
	// JobDocuments shall contain a link to a resource collection of type
	// 'JobDocumentCollection'. This property shall only be present if the service
	// supports document-based jobs.
	//
	// Version added: v1.1.0
	jobDocuments string
	// JobExecutors shall contain a link to a resource collection of type
	// 'JobExecutorCollection'. This property shall only be present if the service
	// supports document-based jobs.
	//
	// Version added: v1.1.0
	jobExecutors string
	// Jobs shall contain a link to a resource collection of type 'JobCollection'.
	jobs string
	// Log shall contain a link to a resource of type 'LogService' that this job
	// service uses.
	log string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceCapabilities shall contain properties that describe the capabilities
	// or supported features of this implementation of a job service.
	ServiceCapabilities JobServiceCapabilities
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ValidationPolicy shall contain policy for how document-based jobs are
	// validated.
	//
	// Version added: v1.1.0
	ValidationPolicy ValidationPolicy
	// cancelAllJobsTarget is the URL to send CancelAllJobs requests.
	cancelAllJobsTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a JobService object from the raw JSON.
func (j *JobService) UnmarshalJSON(b []byte) error {
	type temp JobService
	type jActions struct {
		CancelAllJobs common.ActionTarget `json:"#JobService.CancelAllJobs"`
	}
	var tmp struct {
		temp
		Actions      jActions
		JobDocuments common.Link `json:"jobDocuments"`
		JobExecutors common.Link `json:"jobExecutors"`
		Jobs         common.Link `json:"jobs"`
		Log          common.Link `json:"log"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*j = JobService(tmp.temp)

	// Extract the links to other entities for later
	j.cancelAllJobsTarget = tmp.Actions.CancelAllJobs.Target
	j.jobDocuments = tmp.JobDocuments.String()
	j.jobExecutors = tmp.JobExecutors.String()
	j.jobs = tmp.Jobs.String()
	j.log = tmp.Log.String()

	// This is a read/write object, so we need to save the raw object data for later
	j.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (j *JobService) Update() error {
	readWriteFields := []string{
		"ServiceCapabilities",
		"ServiceEnabled",
		"Status",
	}

	return j.UpdateFromRawData(j, j.rawData, readWriteFields)
}

// GetJobService will get a JobService instance from the service.
func GetJobService(c common.Client, uri string) (*JobService, error) {
	return common.GetObject[JobService](c, uri)
}

// ListReferencedJobServices gets the collection of JobService from
// a provided reference.
func ListReferencedJobServices(c common.Client, link string) ([]*JobService, error) {
	return common.GetCollectionObjects[JobService](c, link)
}

// CancelAllJobs shall cancel all jobs. The service shall transition all jobs to
// the 'Cancelled' state.
func (j *JobService) CancelAllJobs() error {
	payload := make(map[string]any)
	return j.Post(j.cancelAllJobsTarget, payload)
}

// JobDocuments gets the JobDocuments collection.
func (j *JobService) JobDocuments(client common.Client) ([]*JobDocument, error) {
	if j.jobDocuments == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[JobDocument](client, j.jobDocuments)
}

// JobExecutors gets the JobExecutors collection.
func (j *JobService) JobExecutors(client common.Client) ([]*JobExecutor, error) {
	if j.jobExecutors == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[JobExecutor](client, j.jobExecutors)
}

// Jobs gets the Jobs collection.
func (j *JobService) Jobs(client common.Client) ([]*Job, error) {
	if j.jobs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Job](client, j.jobs)
}

// Log gets the Log linked resource.
func (j *JobService) Log(client common.Client) (*LogService, error) {
	if j.log == "" {
		return nil, nil
	}
	return common.GetObject[LogService](client, j.log)
}

// JobServiceCapabilities shall contain properties that describe the
// capabilities or supported features of this implementation of a job service.
type JobServiceCapabilities struct {
	// DocumentBasedJobs shall indicate whether document-based jobs are supported.
	// Document-based jobs are jobs that are created by performing the 'SubmitJob'
	// action on a 'JobDocument' resource in the 'JobDocumentCollection' resource
	// referenced by the 'JobDocuments' property.
	//
	// Version added: v1.1.0
	DocumentBasedJobs bool
	// MaxJobs shall contain the maximum number of jobs supported by the
	// implementation.
	MaxJobs *int `json:",omitempty"`
	// MaxSteps shall contain the maximum number of steps supported by a single job
	// instance.
	MaxSteps *int `json:",omitempty"`
	// Scheduling shall indicate whether the 'Schedule' property within the job
	// supports scheduling of jobs.
	Scheduling bool
	// UserSpecifiedJobs shall indicate whether user-specified jobs are supported.
	// User-specified jobs are jobs that are created by performing an HTTP 'POST'
	// operation on the 'JobCollection' resource referenced by the 'Jobs' property.
	//
	// Version added: v1.1.0
	UserSpecifiedJobs bool
}
