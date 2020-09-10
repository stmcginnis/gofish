//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// IOPerformanceLineOfService is used to define a service option related
// to IO performance.
type IOPerformanceLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AverageIOOperationLatencyMicroseconds shall be the expected average IO
	// latency in microseconds calculated over sample periods (see
	// SamplePeriodSeconds).
	AverageIOOperationLatencyMicroseconds int
	// Description provides a description of this resource.
	Description string
	// IOOperationsPerSecondIsLimited means if true, the system should not allow
	// IOPS to exceed MaxIoOperationsPerSecondPerTerabyte * VolumeSize.
	// Otherwise, the system shall not enforce a limit. The default value
	// for this property is false.
	IOOperationsPerSecondIsLimited bool
	// IOWorkload shall be a description of the expected
	// workload. The workload provides the context in which the values of
	// MaxIOOperationsPerSecondPerTerabyte and
	// AverageIOOperationLatencyMicroseconds are expected to be achievable.
	IOWorkload IOWorkload
	// MaxIOOperationsPerSecondPerTerabyte shall be the amount
	// of IOPS a volume of a given committed size in Terabytes can support.
	// This IOPS density value is useful as a metric that is independent of
	// capacity. Cost is a function of this value and the
	// AverageIOOperationLatencyMicroseconds.
	MaxIOOperationsPerSecondPerTerabyte int
	// SamplePeriod shall be an ISO 8601 duration specifying the
	// sampling period over which average values are calculated.
	SamplePeriod string
}

// GetIOPerformanceLineOfService will get a IOPerformanceLineOfService instance from the service.
func GetIOPerformanceLineOfService(c common.Client, uri string) (*IOPerformanceLineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ioperformancelineofservice IOPerformanceLineOfService
	err = json.NewDecoder(resp.Body).Decode(&ioperformancelineofservice)
	if err != nil {
		return nil, err
	}

	ioperformancelineofservice.SetClient(c)
	return &ioperformancelineofservice, nil
}

// ListReferencedIOPerformanceLineOfServices gets the collection of IOPerformanceLineOfService from
// a provided reference.
func ListReferencedIOPerformanceLineOfServices(c common.Client, link string) ([]*IOPerformanceLineOfService, error) {
	var result []*IOPerformanceLineOfService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, ioperformancelineofserviceLink := range links.ItemLinks {
		ioperformancelineofservice, err := GetIOPerformanceLineOfService(c, ioperformancelineofserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, ioperformancelineofservice)
	}

	return result, nil
}
