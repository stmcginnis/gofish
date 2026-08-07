//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/coreweave/gofish/common"
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
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SamplePeriod shall be an ISO 8601 duration specifying the
	// sampling period over which average values are calculated.
	SamplePeriod string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a IOPerformanceLineOfService object from the raw JSON.
func (ioperformancelineofservice *IOPerformanceLineOfService) UnmarshalJSON(b []byte) error {
	type temp IOPerformanceLineOfService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ioperformancelineofservice = IOPerformanceLineOfService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	ioperformancelineofservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ioperformancelineofservice *IOPerformanceLineOfService) Update() error {
	return ioperformancelineofservice.UpdateWithContext(common.ContextOf(ioperformancelineofservice.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (ioperformancelineofservice *IOPerformanceLineOfService) UpdateWithContext(ctx context.Context) error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(IOPerformanceLineOfService)
	original.UnmarshalJSON(ioperformancelineofservice.rawData)

	readWriteFields := []string{
		"AverageIOOperationLatencyMicroseconds",
		"IOOperationsPerSecondIsLimited",
		"MaxIOOperationsPerSecondPerTerabyte",
		"SamplePeriod",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ioperformancelineofservice).Elem()

	return ioperformancelineofservice.Entity.UpdateWithContext(ctx, originalElement, currentElement, readWriteFields)
}

// GetIOPerformanceLineOfService will get a IOPerformanceLineOfService instance from the service.
func GetIOPerformanceLineOfService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IOPerformanceLineOfService, error) {
	return GetIOPerformanceLineOfServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetIOPerformanceLineOfServiceWithContext will get a IOPerformanceLineOfService instance from the service.
func GetIOPerformanceLineOfServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IOPerformanceLineOfService, error) {
	return common.GetObjectWithContext[IOPerformanceLineOfService](ctx, c, uri, queryOpts...)
}

// ListReferencedIOPerformanceLineOfServices gets the collection of IOPerformanceLineOfService from
// a provided reference.
func ListReferencedIOPerformanceLineOfServices(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*IOPerformanceLineOfService, error) {
	return ListReferencedIOPerformanceLineOfServicesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedIOPerformanceLineOfServicesWithContext gets the collection of IOPerformanceLineOfService from
// a provided reference.
func ListReferencedIOPerformanceLineOfServicesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*IOPerformanceLineOfService, error) {
	return common.GetCollectionObjectsWithContext[IOPerformanceLineOfService](ctx, c, link, queryOpts...)
}
