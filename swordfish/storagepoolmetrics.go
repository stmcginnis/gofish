//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.8 - #StoragePoolMetrics.v1_1_0.StoragePoolMetrics

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// StoragePoolMetrics shall contain the usage and health statistics for a
// storage pool in a Redfish implementation.
type StoragePoolMetrics struct {
	common.Entity
	// CompressionSavingsBytes shall represent the current compression savings on
	// the storage pool in Bytes.
	//
	// Version added: v1.1.0
	CompressionSavingsBytes *int `json:",omitempty"`
	// ConsistencyCheckErrorCount shall contain the number of consistency check
	// errors over the lifetime of the storage pool.
	ConsistencyCheckErrorCount *int `json:",omitempty"`
	// CorrectableIOReadErrorCount shall contain the number of the correctable read
	// errors for the lifetime of the storage pool.
	CorrectableIOReadErrorCount *int `json:",omitempty"`
	// CorrectableIOWriteErrorCount shall contain the number of the correctable
	// write errors for the lifetime of the storage pool.
	CorrectableIOWriteErrorCount *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// on the storage pool in Bytes.
	//
	// Version added: v1.1.0
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall represent IO statistics for this storage pool.
	IOStatistics IOStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RebuildErrorCount shall contain the number of rebuild errors over the
	// lifetime of the storage pool.
	RebuildErrorCount *int `json:",omitempty"`
	// StateChangeCount shall contain the number of state changes (changes in
	// Status.State) for this storage pool.
	StateChangeCount *int `json:",omitempty"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the storage pool in Bytes.
	//
	// Version added: v1.1.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
	// UncorrectableIOReadErrorCount shall contain the number of the uncorrectable
	// read errors for the lifetime of the storage pool.
	UncorrectableIOReadErrorCount *int `json:",omitempty"`
	// UncorrectableIOWriteErrorCount shall contain the number of the uncorrectable
	// write errors for the lifetime of the storage pool.
	UncorrectableIOWriteErrorCount *int `json:",omitempty"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StoragePoolMetrics object from the raw JSON.
func (s *StoragePoolMetrics) UnmarshalJSON(b []byte) error {
	type temp StoragePoolMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StoragePoolMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StoragePoolMetrics) Update() error {
	readWriteFields := []string{
		"IOStatistics",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStoragePoolMetrics will get a StoragePoolMetrics instance from the service.
func GetStoragePoolMetrics(c common.Client, uri string) (*StoragePoolMetrics, error) {
	return common.GetObject[StoragePoolMetrics](c, uri)
}

// ListReferencedStoragePoolMetricss gets the collection of StoragePoolMetrics from
// a provided reference.
func ListReferencedStoragePoolMetricss(c common.Client, link string) ([]*StoragePoolMetrics, error) {
	return common.GetCollectionObjects[StoragePoolMetrics](c, link)
}
