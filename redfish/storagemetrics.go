//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.4 - #StorageMetrics.v1_0_0.StorageMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// StorageMetrics shall contain the usage and health statistics for a storage
// subsystem in a Redfish implementation.
type StorageMetrics struct {
	common.Entity
	// CompressionSavingsBytes shall represent the current compression savings on
	// the storage system in bytes.
	CompressionSavingsBytes *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// on the storage system in bytes.
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall contain the I/O statistics for this storage system.
	IOStatistics common.IOStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// StateChangeCount shall contain the number of times the 'State' property
	// within the 'Status' property of the parent 'Storage' resource changed.
	StateChangeCount *int `json:",omitempty"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the storage system in bytes.
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StorageMetrics object from the raw JSON.
func (s *StorageMetrics) UnmarshalJSON(b []byte) error {
	type temp StorageMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageMetrics) Update() error {
	readWriteFields := []string{
		"IOStatistics",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStorageMetrics will get a StorageMetrics instance from the service.
func GetStorageMetrics(c common.Client, uri string) (*StorageMetrics, error) {
	return common.GetObject[StorageMetrics](c, uri)
}

// ListReferencedStorageMetricss gets the collection of StorageMetrics from
// a provided reference.
func ListReferencedStorageMetricss(c common.Client, link string) ([]*StorageMetrics, error) {
	return common.GetCollectionObjects[StorageMetrics](c, link)
}
