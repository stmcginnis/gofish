//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// StorageServiceMetrics shall contain the usage and health statistics for a storage service in a Redfish
// implementation.
type StorageServiceMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this storage service.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetStorageServiceMetrics will get a StorageServiceMetrics instance from the service.
func GetStorageServiceMetrics(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageServiceMetrics, error) {
	return GetStorageServiceMetricsWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetStorageServiceMetricsWithContext will get a StorageServiceMetrics instance from the service.
func GetStorageServiceMetricsWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageServiceMetrics, error) {
	return common.GetObjectWithContext[StorageServiceMetrics](ctx, c, uri, queryOpts...)
}

// ListReferencedStorageServiceMetricss gets the collection of StorageServiceMetrics from
// a provided reference.
func ListReferencedStorageServiceMetricss(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageServiceMetrics, error) {
	return ListReferencedStorageServiceMetricssWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedStorageServiceMetricssWithContext gets the collection of StorageServiceMetrics from
// a provided reference.
func ListReferencedStorageServiceMetricssWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageServiceMetrics, error) {
	return common.GetCollectionObjectsWithContext[StorageServiceMetrics](ctx, c, link, queryOpts...)
}
