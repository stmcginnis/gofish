//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/StorageServiceMetrics.v1_1_0.json
// 1.2.8 - #StorageServiceMetrics.v1_1_0.StorageServiceMetrics

package schemas

import (
	"encoding/json"
)

// StorageServiceMetrics shall contain the usage and health statistics for a
// storage service in a Redfish implementation.
type StorageServiceMetrics struct {
	Entity
	// CompressionSavingsBytes shall represent the current compression savings on
	// the storage service in Bytes.
	//
	// Version added: v1.1.0
	CompressionSavingsBytes *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// on the storage service in Bytes.
	//
	// Version added: v1.1.0
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall represent IO statistics for this storage service.
	IOStatistics IOStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the storage service in Bytes.
	//
	// Version added: v1.1.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
}

// GetStorageServiceMetrics will get a StorageServiceMetrics instance from the service.
func GetStorageServiceMetrics(c Client, uri string) (*StorageServiceMetrics, error) {
	return GetObject[StorageServiceMetrics](c, uri)
}

// ListReferencedStorageServiceMetricss gets the collection of StorageServiceMetrics from
// a provided reference.
func ListReferencedStorageServiceMetricss(c Client, link string) ([]*StorageServiceMetrics, error) {
	return GetCollectionObjects[StorageServiceMetrics](c, link)
}
