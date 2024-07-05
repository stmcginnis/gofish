//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"time"

	"github.com/stmcginnis/gofish/common"
)

// JobServiceCapabilities shall contain properties that describe the
// capabilities or supported features of this implementation of a job service.
type JobServiceCapabilities struct {
	// MaxJobs shall contain the maximum number of jobs supported by the
	// implementation.
	MaxJobs int
	// MaxSteps shall contain the maximum number of steps supported by a
	// single job instance.
	MaxSteps int
	// Scheduling shall indicate whether the Schedule property within the job
	// supports scheduling of jobs.
	Scheduling bool
}

// JobService shall represent a job service for a Redfish implementation.
type JobService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// DateTime shall contain the current date and time setting for the job
	// service.
	DateTime time.Time
	// Description provides a description of this resource.
	Description string
	// jobs shall contain a link to a resource collection of type JobCollection.
	jobs string
	// log shall contain a link to a resource of type LogService that this
	// job service uses.
	log string
	// ServiceCapabilities shall contain properties that describe the
	// capabilities or supported features of this implementation of a job
	// service.
	ServiceCapabilities JobServiceCapabilities
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a JobService object from the raw JSON.
func (jobservice *JobService) UnmarshalJSON(b []byte) error {
	type temp JobService
	var t struct {
		temp
		Jobs common.Link
		Log  common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*jobservice = JobService(t.temp)

	// Extract the links to other entities for later
	jobservice.jobs = t.Jobs.String()
	jobservice.log = t.Log.String()
	return nil
}

// GetJobService will get a JobService instance from the service.
func GetJobService(c common.Client, uri string) (*JobService, error) {
	return common.GetObject[JobService](c, uri)
}

// Jobs gets the collection of jobs of this job service
func (jobservice *JobService) Jobs() ([]*Job, error) {
	return ListReferencedJobs(jobservice.GetClient(), jobservice.jobs)
}

// Log gets the LogService instance for this job service
func (jobservice *JobService) Log() (*LogService, error) {
	return GetLogService(jobservice.GetClient(), jobservice.log)
}
