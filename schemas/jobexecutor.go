//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #JobExecutor.v1_0_0.JobExecutor

package schemas

import (
	"encoding/json"
)

// JobExecutor shall represent a job executor for a Redfish implementation.
type JobExecutor struct {
	Entity
	// ExecutorType shall contain the primary type of job executed by this
	// resource.
	ExecutorType string
	// MaximumConcurrentJobs shall contain the maximum concurrent jobs this
	// executor can process.
	MaximumConcurrentJobs int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// chassis are the URIs for Chassis.
	chassis []string
	// computerSystem is the URI for ComputerSystem.
	computerSystem string
	// executingJobs are the URIs for ExecutingJobs.
	executingJobs []string
}

// UnmarshalJSON unmarshals a JobExecutor object from the raw JSON.
func (j *JobExecutor) UnmarshalJSON(b []byte) error {
	type temp JobExecutor
	type jLinks struct {
		Chassis        Links `json:"Chassis"`
		ComputerSystem Link  `json:"ComputerSystem"`
		ExecutingJobs  Links `json:"ExecutingJobs"`
	}
	var tmp struct {
		temp
		Links jLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*j = JobExecutor(tmp.temp)

	// Extract the links to other entities for later
	j.chassis = tmp.Links.Chassis.ToStrings()
	j.computerSystem = tmp.Links.ComputerSystem.String()
	j.executingJobs = tmp.Links.ExecutingJobs.ToStrings()

	return nil
}

// GetJobExecutor will get a JobExecutor instance from the service.
func GetJobExecutor(c Client, uri string) (*JobExecutor, error) {
	return GetObject[JobExecutor](c, uri)
}

// ListReferencedJobExecutors gets the collection of JobExecutor from
// a provided reference.
func ListReferencedJobExecutors(c Client, link string) ([]*JobExecutor, error) {
	return GetCollectionObjects[JobExecutor](c, link)
}

// Chassis gets the Chassis linked resources.
func (j *JobExecutor) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](j.client, j.chassis)
}

// ComputerSystem gets the ComputerSystem linked resource.
func (j *JobExecutor) ComputerSystem() (*ComputerSystem, error) {
	if j.computerSystem == "" {
		return nil, nil
	}
	return GetObject[ComputerSystem](j.client, j.computerSystem)
}

// ExecutingJobs gets the ExecutingJobs linked resources.
func (j *JobExecutor) ExecutingJobs() ([]*Job, error) {
	return GetObjects[Job](j.client, j.executingJobs)
}
