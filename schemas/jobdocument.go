//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #JobDocument.v1_0_0.JobDocument

package schemas

import (
	"encoding/json"
)

type DataType string

const (
	// BooleanDataType is a boolean.
	BooleanDataType DataType = "Boolean"
	// NumberDataType is a number.
	NumberDataType DataType = "Number"
	// StringDataType is a string.
	StringDataType DataType = "String"
)

// JobDocument shall represent a job document for a Redfish implementation.
type JobDocument struct {
	Entity
	// CreationTime shall contain the date and time when this job document resource
	// was created.
	CreationTime string
	// DocumentData shall contain a Base64-encoded string of the job document data.
	// This property shall not be present if 'DocumentDataURI' is present.
	DocumentData string
	// DocumentDataHash shall contain the hash of the job document data as a
	// hex-encoded string.
	DocumentDataHash string
	// DocumentDataURI shall contain the URI at which to access the job document
	// data. This property shall not be present if 'DocumentData' is present.
	DocumentDataURI string
	// DocumentType shall contain the type of job document data associated with
	// this job document.
	DocumentType string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ParameterMetadata shall contain the metadata for each of the parameters
	// supported by this job document for the 'SubmitJob' action.
	ParameterMetadata []ParameterMetadata
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Version shall contain the version of this job document.
	Version string
	// submitJobTarget is the URL to send SubmitJob requests.
	submitJobTarget string
	// supportedExecutors are the URIs for SupportedExecutors.
	supportedExecutors []string
}

// UnmarshalJSON unmarshals a JobDocument object from the raw JSON.
func (j *JobDocument) UnmarshalJSON(b []byte) error {
	type temp JobDocument
	type jActions struct {
		SubmitJob ActionTarget `json:"#JobDocument.SubmitJob"`
	}
	type jLinks struct {
		SupportedExecutors Links `json:"SupportedExecutors"`
	}
	var tmp struct {
		temp
		Actions jActions
		Links   jLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*j = JobDocument(tmp.temp)

	// Extract the links to other entities for later
	j.submitJobTarget = tmp.Actions.SubmitJob.Target
	j.supportedExecutors = tmp.Links.SupportedExecutors.ToStrings()

	return nil
}

// GetJobDocument will get a JobDocument instance from the service.
func GetJobDocument(c Client, uri string) (*JobDocument, error) {
	return GetObject[JobDocument](c, uri)
}

// ListReferencedJobDocuments gets the collection of JobDocument from
// a provided reference.
func ListReferencedJobDocuments(c Client, link string) ([]*JobDocument, error) {
	return GetCollectionObjects[JobDocument](c, link)
}

// JobDocumentSubmitJobParameters holds the parameters for the SubmitJob action.
type JobDocumentSubmitJobParameters struct {
	// HidePayload shall indicate whether the contents of the parameters should be
	// hidden from view after the job has been created. If 'true', responses shall
	// not return the 'Parameters' property. If 'false', responses shall return the
	// 'Parameters' property. If this parameter is not present when the job is
	// created, the default is 'false'.
	HidePayload bool `json:"HidePayload,omitempty"`
	// JobCreator shall contain a link to a resource of type 'Job' that represents
	// the job that is submitting this job.
	JobCreator string `json:"JobCreator,omitempty"`
	// Parameters shall contain the list of parameters for the new job that are
	// specific to this job document. Services shall reject requests containing
	// parameters that do not meet the requirements specified by the
	// 'ParameterMetadata' property.
	Parameters []any `json:"Parameters,omitempty"`
	// PreferredExecutors shall contain an array of links to resources of type
	// 'JobExecutor' that represent the preferred executors to run this job.
	PreferredExecutors []string `json:"PreferredExecutors,omitempty"`
	// StartTime shall contain the date and time when the job is scheduled to
	// start.
	StartTime string `json:"StartTime,omitempty"`
}

// This action shall create a new 'Job' resource based on the contents of this
// job document and additional parameters.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (j *JobDocument) SubmitJob(params *JobDocumentSubmitJobParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(j.client,
		j.submitJobTarget, params, j.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// SupportedExecutors gets the SupportedExecutors linked resources.
func (j *JobDocument) SupportedExecutors() ([]*JobExecutor, error) {
	return GetObjects[JobExecutor](j.client, j.supportedExecutors)
}

// ParameterMetadata shall contain the metadata a parameter supported by a job
// document for the 'SubmitJob' action.
type ParameterMetadata struct {
	// AllowableNumbers shall indicate the allowable numeric values, inclusive
	// ranges of values, and incremental step values for this parameter, as defined
	// in the 'Allowable values for numbers and durations' clause of the Redfish
	// Specification. This property shall only be present for numeric parameters or
	// string parameters that specify a duration.
	AllowableNumbers []string
	// AllowablePattern shall contain a regular expression that describes the
	// allowable values for this parameter. This property shall only be present for
	// string parameters.
	AllowablePattern string
	// AllowableValueDescriptions shall contain the descriptions of allowable
	// values for this parameter. The descriptions shall appear in the same array
	// order as the 'AllowableValues' property.
	AllowableValueDescriptions []string
	// AllowableValues shall indicate the allowable values for this parameter.
	AllowableValues []string
	// DataType shall contain the JSON property type for this parameter.
	DataType DataType
	// Description provides a description of this resource.
	Description string
	// MaximumValue shall contain the maximum value that this service supports.
	// This property shall not be present for non-integer or number parameters.
	MaximumValue float64
	// MinimumValue shall contain the minimum value that this service supports.
	// This property shall not be present for non-integer or number parameters.
	MinimumValue float64
	// Name is the name of the resource or array element.
	Name string
	// Required shall indicate whether the parameter is required.
	Required bool
	// ValueHint shall contain a hint value for the parameter.
	ValueHint string
}
