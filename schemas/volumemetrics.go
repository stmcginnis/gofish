//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/VolumeMetrics.v1_2_0.json
// 1.2.8 - #VolumeMetrics.v1_2_0.VolumeMetrics

package schemas

import (
	"encoding/json"
)

// VolumeMetrics shall contain the usage and health statistics for a volume in a
// Redfish implementation.
type VolumeMetrics struct {
	Entity
	// CompressionSavingsBytes shall represent the current compression savings on
	// the volume in Bytes.
	//
	// Version added: v1.2.0
	CompressionSavingsBytes *int `json:",omitempty"`
	// ConsistencyCheckCount shall contain the number of consistency checks
	// completed over the lifetime of the volume.
	//
	// Version added: v1.1.0
	ConsistencyCheckCount *int `json:",omitempty"`
	// ConsistencyCheckErrorCount shall contain the number of consistency check
	// errors over the lifetime of the volume.
	ConsistencyCheckErrorCount *int `json:",omitempty"`
	// CorrectableIOReadErrorCount shall contain the number of the correctable read
	// errors for the lifetime of the volume.
	CorrectableIOReadErrorCount *int `json:",omitempty"`
	// CorrectableIOWriteErrorCount shall contain the number of the correctable
	// write errors for the lifetime of the volume.
	CorrectableIOWriteErrorCount *int `json:",omitempty"`
	// DeduplicationSavingsBytes shall represent the current deduplication savings
	// on the volume in Bytes.
	//
	// Version added: v1.2.0
	DeduplicationSavingsBytes *int `json:",omitempty"`
	// IOStatistics shall represent IO statistics for this volume.
	//
	// Version added: v1.1.0
	IOStatistics IOStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RebuildErrorCount shall contain the number of rebuild errors over the
	// lifetime of the volume.
	RebuildErrorCount *int `json:",omitempty"`
	// StateChangeCount shall contain the number of state changes (changes in
	// Status.State) for this volume.
	StateChangeCount *int `json:",omitempty"`
	// ThinProvisioningSavingsBytes shall represent the current thin provisioning
	// savings on the volume in Bytes.
	//
	// Version added: v1.2.0
	ThinProvisioningSavingsBytes *int `json:",omitempty"`
	// UncorrectableIOReadErrorCount shall contain the number of the uncorrectable
	// read errors for the lifetime of the volume.
	UncorrectableIOReadErrorCount *int `json:",omitempty"`
	// UncorrectableIOWriteErrorCount shall contain the number of the uncorrectable
	// write errors for the lifetime of the volume.
	UncorrectableIOWriteErrorCount *int `json:",omitempty"`
}

// GetVolumeMetrics will get a VolumeMetrics instance from the service.
func GetVolumeMetrics(c Client, uri string) (*VolumeMetrics, error) {
	return GetObject[VolumeMetrics](c, uri)
}

// ListReferencedVolumeMetricss gets the collection of VolumeMetrics from
// a provided reference.
func ListReferencedVolumeMetricss(c Client, link string) ([]*VolumeMetrics, error) {
	return GetCollectionObjects[VolumeMetrics](c, link)
}
