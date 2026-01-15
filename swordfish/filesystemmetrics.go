//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.8 - #FileSystemMetrics.v1_1_0.FileSystemMetrics

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// FileSystemMetrics shall contain the usage and health statistics for a file
// system in a Redfish implementation.
type FileSystemMetrics struct {
	common.Entity
	// CompressionSavingsBytes shall represent the current compression savings in
	// the file system in Bytes.
	//
	// Version added: v1.1.0
	CompressionSavingsBytes *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// in the file system in Bytes.
	//
	// Version added: v1.1.0
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall represent IO statistics for this file system.
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
	// savings in the file system in Bytes.
	//
	// Version added: v1.1.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a FileSystemMetrics object from the raw JSON.
func (f *FileSystemMetrics) UnmarshalJSON(b []byte) error {
	type temp FileSystemMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FileSystemMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	f.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *FileSystemMetrics) Update() error {
	readWriteFields := []string{
		"IOStatistics",
	}

	return f.UpdateFromRawData(f, f.rawData, readWriteFields)
}

// GetFileSystemMetrics will get a FileSystemMetrics instance from the service.
func GetFileSystemMetrics(c common.Client, uri string) (*FileSystemMetrics, error) {
	return common.GetObject[FileSystemMetrics](c, uri)
}

// ListReferencedFileSystemMetricss gets the collection of FileSystemMetrics from
// a provided reference.
func ListReferencedFileSystemMetricss(c common.Client, link string) ([]*FileSystemMetrics, error) {
	return common.GetCollectionObjects[FileSystemMetrics](c, link)
}
