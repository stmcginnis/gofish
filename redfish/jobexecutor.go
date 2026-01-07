//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #JobExecutor.v1_0_0.JobExecutor

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// JobExecutor shall represent a job executor for a Redfish implementation.
type JobExecutor struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// chassis are the URIs for Chassis.
	chassis []string
	// computerSystem is the URI for ComputerSystem.
	computerSystem string
	// executingJobs are the URIs for ExecutingJobs.
	executingJobs []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a JobExecutor object from the raw JSON.
func (j *JobExecutor) UnmarshalJSON(b []byte) error {
	type temp JobExecutor
	type jLinks struct {
		Chassis        common.Links `json:"Chassis"`
		ComputerSystem common.Link  `json:"ComputerSystem"`
		ExecutingJobs  common.Links `json:"ExecutingJobs"`
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

	// This is a read/write object, so we need to save the raw object data for later
	j.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (j *JobExecutor) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return j.UpdateFromRawData(j, j.rawData, readWriteFields)
}

// GetJobExecutor will get a JobExecutor instance from the service.
func GetJobExecutor(c common.Client, uri string) (*JobExecutor, error) {
	return common.GetObject[JobExecutor](c, uri)
}

// ListReferencedJobExecutors gets the collection of JobExecutor from
// a provided reference.
func ListReferencedJobExecutors(c common.Client, link string) ([]*JobExecutor, error) {
	return common.GetCollectionObjects[JobExecutor](c, link)
}

// Chassis gets the Chassis linked resources.
func (j *JobExecutor) Chassis(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, j.chassis)
}

// ComputerSystem gets the ComputerSystem linked resource.
func (j *JobExecutor) ComputerSystem(client common.Client) (*ComputerSystem, error) {
	if j.computerSystem == "" {
		return nil, nil
	}
	return common.GetObject[ComputerSystem](client, j.computerSystem)
}

// ExecutingJobs gets the ExecutingJobs linked resources.
func (j *JobExecutor) ExecutingJobs(client common.Client) ([]*Job, error) {
	return common.GetObjects[Job](client, j.executingJobs)
}
