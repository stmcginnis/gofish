//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/StorageMetrics.v1_0_0.json
// 2024.4 - #StorageMetrics.v1_0_0.StorageMetrics

package schemas

import (
	"encoding/json"
)

// StorageMetrics shall contain the usage and health statistics for a storage
// subsystem in a Redfish implementation.
type StorageMetrics struct {
	Entity
	// CompressionSavingsBytes shall represent the current compression savings on
	// the storage system in bytes.
	CompressionSavingsBytes *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// on the storage system in bytes.
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall contain the I/O statistics for this storage system.
	IOStatistics IOStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// StateChangeCount shall contain the number of times the 'State' property
	// within the 'Status' property of the parent 'Storage' resource changed.
	StateChangeCount *int `json:",omitempty"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the storage system in bytes.
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
}

// GetStorageMetrics will get a StorageMetrics instance from the service.
func GetStorageMetrics(c Client, uri string) (*StorageMetrics, error) {
	return GetObject[StorageMetrics](c, uri)
}

// ListReferencedStorageMetricss gets the collection of StorageMetrics from
// a provided reference.
func ListReferencedStorageMetricss(c Client, link string) ([]*StorageMetrics, error) {
	return GetCollectionObjects[StorageMetrics](c, link)
}
