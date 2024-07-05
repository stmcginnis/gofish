//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// VolumeMetrics shall contain the usage and health statistics for a volume in a Redfish implementation.
type VolumeMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConsistencyCheckCount shall contain the number of consistency checks completed over the lifetime of the volume.
	ConsistencyCheckCount int64
	// ConsistencyCheckErrorCount shall contain the number of consistency check errors over the lifetime of the volume.
	ConsistencyCheckErrorCount int64
	// CorrectableIOReadErrorCount shall contain the number of the correctable read errors for the lifetime of the
	// volume.
	CorrectableIOReadErrorCount int64
	// CorrectableIOWriteErrorCount shall contain the number of the correctable write errors for the lifetime of the
	// volume.
	CorrectableIOWriteErrorCount int64
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this volume.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RebuildErrorCount shall contain the number of rebuild errors over the lifetime of the volume.
	RebuildErrorCount int64
	// StateChangeCount shall contain the number of state changes (changes in Status.State) for this volume.
	StateChangeCount int64
	// UncorrectableIOReadErrorCount shall contain the number of the uncorrectable read errors for the lifetime of the
	// volume.
	UncorrectableIOReadErrorCount int64
	// UncorrectableIOWriteErrorCount shall contain the number of the uncorrectable write errors for the lifetime of
	// the volume.
	UncorrectableIOWriteErrorCount int64
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
