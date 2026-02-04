//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/IOPerformanceLoSCapabilities.v1_3_0.json
// 1.2.1c - #IOPerformanceLoSCapabilities.v1_3_0.IOPerformanceLoSCapabilities

package schemas

import (
	"encoding/json"
)

// IOAccessPattern is The enumeration literals may be used to specify an IO
// access pattern.
type IOAccessPattern string

const (
	// ReadWriteIOAccessPattern shall indicate a Uniform distribution of reads and
	// writes.
	ReadWriteIOAccessPattern IOAccessPattern = "ReadWrite"
	// SequentialReadIOAccessPattern shall indicate a sequential read pattern of
	// access.
	SequentialReadIOAccessPattern IOAccessPattern = "SequentialRead"
	// SequentialWriteIOAccessPattern shall indicate a sequential write pattern of
	// access.
	SequentialWriteIOAccessPattern IOAccessPattern = "SequentialWrite"
	// RandomReadNewIOAccessPattern shall indicate an access pattern of random
	// reads of uncached data.
	RandomReadNewIOAccessPattern IOAccessPattern = "RandomReadNew"
	// RandomReadAgainIOAccessPattern shall indicate an access pattern of random
	// reads of cached data.
	RandomReadAgainIOAccessPattern IOAccessPattern = "RandomReadAgain"
)

// IOPerformanceLoSCapabilities shall describe the capabilities of the system to
// support various IO performance service options.
type IOPerformanceLoSCapabilities struct {
	Entity
	// IOLimitingIsSupported shall not inforce a limit. The default value for this
	// property is false.
	IOLimitingIsSupported bool
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// MaxSamplePeriod shall be an ISO 8601 duration specifying the maximum
	// sampling period over which average values are calculated.
	MaxSamplePeriod string
	// MinSamplePeriod shall be an ISO 8601 duration specifying the minimum
	// sampling period over which average values are calculated.
	MinSamplePeriod string
	// MinSupportedIoOperationLatencyMicroseconds shall be the minimum supported
	// average IO latency in microseconds calculated over the SamplePeriod.
	MinSupportedIoOperationLatencyMicroseconds *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupportedIOWorkloads shall be a collection of supported workloads.
	SupportedIOWorkloads []IOWorkload
	// SupportedLinesOfService shall be a collection supported IO performance
	// service options.
	supportedLinesOfService []string
	// SupportedLinesOfServiceCount
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IOPerformanceLoSCapabilities object from the raw JSON.
func (i *IOPerformanceLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp IOPerformanceLoSCapabilities
	var tmp struct {
		temp
		SupportedLinesOfService Links `json:"SupportedLinesOfService"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = IOPerformanceLoSCapabilities(tmp.temp)

	// Extract the links to other entities for later
	i.supportedLinesOfService = tmp.SupportedLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IOPerformanceLoSCapabilities) Update() error {
	readWriteFields := []string{
		"IOLimitingIsSupported",
		"MaxSamplePeriod",
		"MinSamplePeriod",
		"MinSupportedIoOperationLatencyMicroseconds",
		"SupportedLinesOfService",
	}

	return i.UpdateFromRawData(i, i.RawData, readWriteFields)
}

// GetIOPerformanceLoSCapabilities will get a IOPerformanceLoSCapabilities instance from the service.
func GetIOPerformanceLoSCapabilities(c Client, uri string) (*IOPerformanceLoSCapabilities, error) {
	return GetObject[IOPerformanceLoSCapabilities](c, uri)
}

// ListReferencedIOPerformanceLoSCapabilitiess gets the collection of IOPerformanceLoSCapabilities from
// a provided reference.
func ListReferencedIOPerformanceLoSCapabilitiess(c Client, link string) ([]*IOPerformanceLoSCapabilities, error) {
	return GetCollectionObjects[IOPerformanceLoSCapabilities](c, link)
}

// SupportedLinesOfService gets the SupportedLinesOfService linked resources.
func (i *IOPerformanceLoSCapabilities) SupportedLinesOfService() ([]*IOPerformanceLineOfService, error) {
	return GetObjects[IOPerformanceLineOfService](i.client, i.supportedLinesOfService)
}

// IOWorkload This structure may be used to describe an IO Workload.
type IOWorkload struct {
	// Components shall be an array of IO workload component descriptions.
	Components []IOWorkloadComponent
	// Name is the name of the resource or array element.
	Name string
}

// IOWorkloadComponent This structure may be used to describe a component of an
// IO workload.
type IOWorkloadComponent struct {
	// AverageIOBytes shall be the expected average I/O size.
	AverageIOBytes *int `json:",omitempty"`
	// Duration shall be an ISO 8601 duration that shall specify the expected
	// length of time that this component is applied to the workload. This
	// attribute shall be specified if a schedule is specified and otherwise shall
	// not be specified.
	Duration string
	// IOAccessPattern shall be the expected access pattern.
	IOAccessPattern IOAccessPattern
	// PercentOfData shall be the expected percent of the data referenced by the
	// workload that is covered by this component.
	PercentOfData *int `json:",omitempty"`
	// PercentOfIOPS shall be the expected percent of the total IOPS for this
	// workload that is covered by this component.
	PercentOfIOPS *int `json:",omitempty"`
	// Schedule shall specifies when this workload component is applied to the
	// overall workload.
	Schedule Schedule
}
