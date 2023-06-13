//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"time"

	"github.com/stmcginnis/gofish/common"
)

// ServiceCapabilities The supported capabilities of this job service implementation
type ServiceCapabilities struct {
	// MaxJobs This property shall contain the maximum number of jobs supported by the implementation.
	MaxJobs int
	// MaxSteps This property shall contain the maximum number of steps supported by a single job instance.
	MaxSteps int
	// ServiceEnabled This property shall indicate whether the Schedule property within the job supports scheduling of jobs.
	ServiceEnabled bool
}

// JobService is used to represent the job service offered by the redfish API
type JobService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// DateTime system time.
	DateTime time.Time
	// jobs points towards the jobs collection
	jobs string
	// log points towards the log collection
	// TODO: Add log collection fetching
	log string

	// ServiceCapabilities This type shall contain properties that describe the capabilities or supported features of this implementation of a job service.
	ServiceCapabilities ServiceCapabilities

	// ServiceEnabled indicates whether this service isenabled.
	ServiceEnabled bool
	// Status describes the status and health of a resource and its children.
	Status common.Status
}

// GetJob will get a Job instance from the service.
func GetJob(c common.Client, uri string) (*Job, error) {
	var job Job
	return &job, job.Get(c, uri, &job)
}

// Jobs will get collection of Job from the service.
func (jobService *JobService) Jobs() ([]*Job, error) {
	return jobService.listReferencedJobs(jobService.GetClient(), jobService.jobs)
}

// ListReferencedJobs gets the collection of Job from a provided reference.
func (jobService *JobService) listReferencedJobs(c common.Client, link string) ([]*Job, error) { //nolint:dupl
	var result []*Job
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Job
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		job, err := GetJob(c, link)
		ch <- GetResult{Item: job, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// UnmarshalJSON unmarshals a JobService object from the raw JSON.
func (jobService *JobService) UnmarshalJSON(b []byte) error {
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

	// Extract the links to other entities for later
	*jobService = JobService(t.temp)
	jobService.jobs = t.Jobs.String()
	jobService.log = t.Log.String()

	return nil
}

// GetJobService will get a JobService instance from the service.
func GetJobService(c common.Client, uri string) (*JobService, error) {
	var jobService JobService
	return &jobService, jobService.Get(c, uri, &jobService)
}
