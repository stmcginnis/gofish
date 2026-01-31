//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.1.0 - #IOPerformanceLineOfService.v1_1_1.IOPerformanceLineOfService

package schemas

import (
	"encoding/json"
)

// IOPerformanceLineOfService This structure may be used to define a service
// option related to IO performance.
type IOPerformanceLineOfService struct {
	Entity
	// AverageIOOperationLatencyMicroseconds shall be the expected average IO
	// latency in microseconds calculated over sample periods (see
	// SamplePeriodSeconds).
	AverageIOOperationLatencyMicroseconds *int `json:",omitempty"`
	// IOOperationsPerSecondIsLimited shall not enforce a limit. The default value
	// for this property is false.
	IOOperationsPerSecondIsLimited bool
	// IOWorkload shall be a description of the expected workload. The workload
	// provides the context in which the values of
	// MaxIOOperationsPerSecondPerTerabyte and
	// AverageIOOperationLatencyMicroseconds are expected to be achievable.
	IOWorkload IOWorkload
	// MaxIOOperationsPerSecondPerTerabyte shall be the amount of IOPS a volume of
	// a given committed size in Terabytes can support. This IOPS density value is
	// useful as a metric that is independent of capacity. Cost is a function of
	// this value and the AverageIOOperationLatencyMicroseconds.
	MaxIOOperationsPerSecondPerTerabyte *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SamplePeriod shall be an ISO 8601 duration specifying the sampling period
	// over which average values are calculated.
	SamplePeriod string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IOPerformanceLineOfService object from the raw JSON.
func (i *IOPerformanceLineOfService) UnmarshalJSON(b []byte) error {
	type temp IOPerformanceLineOfService
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = IOPerformanceLineOfService(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IOPerformanceLineOfService) Update() error {
	readWriteFields := []string{
		"AverageIOOperationLatencyMicroseconds",
		"IOOperationsPerSecondIsLimited",
		"MaxIOOperationsPerSecondPerTerabyte",
		"SamplePeriod",
	}

	return i.UpdateFromRawData(i, i.RawData, readWriteFields)
}

// GetIOPerformanceLineOfService will get a IOPerformanceLineOfService instance from the service.
func GetIOPerformanceLineOfService(c Client, uri string) (*IOPerformanceLineOfService, error) {
	return GetObject[IOPerformanceLineOfService](c, uri)
}

// ListReferencedIOPerformanceLineOfServices gets the collection of IOPerformanceLineOfService from
// a provided reference.
func ListReferencedIOPerformanceLineOfServices(c Client, link string) ([]*IOPerformanceLineOfService, error) {
	return GetCollectionObjects[IOPerformanceLineOfService](c, link)
}
