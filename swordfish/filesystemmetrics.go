//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// FileSystemMetrics shall contain the usage and health statistics for a file system in a Redfish implementation.
type FileSystemMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this file system.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetFileSystemMetrics will get a FileSystemMetrics instance from the service.
func GetFileSystemMetrics(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*FileSystemMetrics, error) {
	return GetFileSystemMetricsWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetFileSystemMetricsWithContext will get a FileSystemMetrics instance from the service.
func GetFileSystemMetricsWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*FileSystemMetrics, error) {
	return common.GetObjectWithContext[FileSystemMetrics](ctx, c, uri, queryOpts...)
}

// ListReferencedFileSystemMetricses gets the collection of FileSystemMetrics from
// a provided reference.
func ListReferencedFileSystemMetricses(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*FileSystemMetrics, error) {
	return ListReferencedFileSystemMetricsesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedFileSystemMetricsesWithContext gets the collection of FileSystemMetrics from
// a provided reference.
func ListReferencedFileSystemMetricsesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*FileSystemMetrics, error) {
	return common.GetCollectionObjectsWithContext[FileSystemMetrics](ctx, c, link, queryOpts...)
}
