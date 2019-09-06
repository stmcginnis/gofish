//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// IOAccessPattern is used to specify an IO access pattern.
type IOAccessPattern string

const (
	// ReadWriteIOAccessPattern shall indicate a Uniform distribution of
	// reads and writes.
	ReadWriteIOAccessPattern IOAccessPattern = "ReadWrite"
	// SequentialReadIOAccessPattern shall indicate a sequential read pattern
	// of access.
	SequentialReadIOAccessPattern IOAccessPattern = "SequentialRead"
	// SequentialWriteIOAccessPattern shall indicate a sequential write
	// pattern of access.
	SequentialWriteIOAccessPattern IOAccessPattern = "SequentialWrite"
	// RandomReadNewIOAccessPattern shall indicate an access pattern of
	// random reads of uncached data.
	RandomReadNewIOAccessPattern IOAccessPattern = "RandomReadNew"
	// RandomReadAgainIOAccessPattern shall indicate an access pattern of
	// random reads of cached data.
	RandomReadAgainIOAccessPattern IOAccessPattern = "RandomReadAgain"
)

// IOPerformanceLoSCapabilities shall describe the capabilities of the
// system to support various IO performance service options.
type IOPerformanceLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IOLimitingIsSupported if true, the system should limit IOPS to
	// MaxIOOperationsPerSecondPerTerabyte * (Volume Size in Terabytes).
	// Otherwise, the system shall not enforce a limit. The default value for
	// this property is false.
	IOLimitingIsSupported bool
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// MaxSamplePeriod shall be an ISO 8601 duration specifying the maximum
	// sampling period over which average values are calculated.
	MaxSamplePeriod string
	// MinSamplePeriod shall be an ISO 8601 duration specifying the minimum
	// sampling period over which average values are calculated.
	MinSamplePeriod string
	// MinSupportedIoOperationLatencyMicroseconds shall be the minimum supported
	// average IO latency in microseconds calculated over the SamplePeriod
	MinSupportedIoOperationLatencyMicroseconds int
	// SupportedIOWorkloads shall be a collection of supported workloads.
	SupportedIOWorkloads []IOWorkload
	// SupportedLinesOfService shall be a collection ofsupported IO performance
	// service options.
	SupportedLinesOfService []IOPerformanceLineOfService
	// SupportedLinesOfServiceCount is
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
}

// GetIOPerformanceLoSCapabilities will get a IOPerformanceLoSCapabilities instance from the service.
func GetIOPerformanceLoSCapabilities(c common.Client, uri string) (*IOPerformanceLoSCapabilities, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ioperformanceloscapabilities IOPerformanceLoSCapabilities
	err = json.NewDecoder(resp.Body).Decode(&ioperformanceloscapabilities)
	if err != nil {
		return nil, err
	}

	ioperformanceloscapabilities.SetClient(c)
	return &ioperformanceloscapabilities, nil
}

// ListReferencedIOPerformanceLoSCapabilitiess gets the collection of IOPerformanceLoSCapabilities from
// a provided reference.
func ListReferencedIOPerformanceLoSCapabilitiess(c common.Client, link string) ([]*IOPerformanceLoSCapabilities, error) {
	var result []*IOPerformanceLoSCapabilities
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, ioperformanceloscapabilitiesLink := range links.ItemLinks {
		ioperformanceloscapabilities, err := GetIOPerformanceLoSCapabilities(c, ioperformanceloscapabilitiesLink)
		if err != nil {
			return result, err
		}
		result = append(result, ioperformanceloscapabilities)
	}

	return result, nil
}

// IOWorkload is used to describe an IO Workload.
type IOWorkload struct {
	// Components shall be an array of IO workload component
	// descriptions.
	Components []IOWorkloadComponent
}

// IOWorkloadComponent is used to describe a component of an IO workload.
type IOWorkloadComponent struct {
	// AverageIOBytes shall be the expected average I/O size.
	AverageIOBytes int
	// Duration is The value of each entry shall be an ISO 8601 duration that
	// shall specify the expected length of time that this component is
	// applied to the workload. This attribute shall be specified if a
	// schedule is specified and otherwise shall not be specified.
	Duration string
	// IOAccessPattern is The enumeration literal shall be the expected
	// access pattern.
	IOAccessPattern IOAccessPattern
	// PercentOfData shall be the expected percent of the data
	// referenced by the workload that is covered by this component.
	PercentOfData int
	// PercentOfIOPS shall be the expected percent of the total
	// IOPS for this workload that is covered by this component
	PercentOfIOPS int
	// Schedule shall specifies when this workload component is
	// applied to the overall workload.
	Schedule common.Schedule
}
