//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.8 - #FileSystemMetrics.v1_1_0.FileSystemMetrics

package schemas

import (
	"encoding/json"
)

// FileSystemMetrics shall contain the usage and health statistics for a file
// system in a Redfish implementation.
type FileSystemMetrics struct {
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings in the file system in Bytes.
	//
	// Version added: v1.1.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
}

// GetFileSystemMetrics will get a FileSystemMetrics instance from the service.
func GetFileSystemMetrics(c Client, uri string) (*FileSystemMetrics, error) {
	return GetObject[FileSystemMetrics](c, uri)
}

// ListReferencedFileSystemMetricss gets the collection of FileSystemMetrics from
// a provided reference.
func ListReferencedFileSystemMetricss(c Client, link string) ([]*FileSystemMetrics, error) {
	return GetCollectionObjects[FileSystemMetrics](c, link)
}
