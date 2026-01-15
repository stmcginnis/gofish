//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.8 - #StorageServiceMetrics.v1_1_0.StorageServiceMetrics

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// StorageServiceMetrics shall contain the usage and health statistics for a
// storage service in a Redfish implementation.
type StorageServiceMetrics struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the storage service in Bytes.
	//
	// Version added: v1.1.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StorageServiceMetrics object from the raw JSON.
func (s *StorageServiceMetrics) UnmarshalJSON(b []byte) error {
	type temp StorageServiceMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageServiceMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageServiceMetrics) Update() error {
	readWriteFields := []string{
		"IOStatistics",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStorageServiceMetrics will get a StorageServiceMetrics instance from the service.
func GetStorageServiceMetrics(c common.Client, uri string) (*StorageServiceMetrics, error) {
	return common.GetObject[StorageServiceMetrics](c, uri)
}

// ListReferencedStorageServiceMetricss gets the collection of StorageServiceMetrics from
// a provided reference.
func ListReferencedStorageServiceMetricss(c common.Client, link string) ([]*StorageServiceMetrics, error) {
	return common.GetCollectionObjects[StorageServiceMetrics](c, link)
}
