//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.8 - #VolumeMetrics.v1_2_0.VolumeMetrics

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// VolumeMetrics shall contain the usage and health statistics for a volume in a
// Redfish implementation.
type VolumeMetrics struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VolumeMetrics object from the raw JSON.
func (v *VolumeMetrics) UnmarshalJSON(b []byte) error {
	type temp VolumeMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VolumeMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	v.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *VolumeMetrics) Update() error {
	readWriteFields := []string{
		"IOStatistics",
	}

	return v.UpdateFromRawData(v, v.rawData, readWriteFields)
}

// GetVolumeMetrics will get a VolumeMetrics instance from the service.
func GetVolumeMetrics(c common.Client, uri string) (*VolumeMetrics, error) {
	return common.GetObject[VolumeMetrics](c, uri)
}

// ListReferencedVolumeMetricss gets the collection of VolumeMetrics from
// a provided reference.
func ListReferencedVolumeMetricss(c common.Client, link string) ([]*VolumeMetrics, error) {
	return common.GetCollectionObjects[VolumeMetrics](c, link)
}
