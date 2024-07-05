//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// StoragePoolMetrics shall contain the usage and health statistics for a storage pool in a Redfish implementation.
type StoragePoolMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConsistencyCheckErrorCount shall contain the number of consistency check errors over the lifetime of the storage
	// pool.
	ConsistencyCheckErrorCount int
	// CorrectableIOReadErrorCount shall contain the number of the correctable read errors for the lifetime of the
	// storage pool.
	CorrectableIOReadErrorCount int
	// CorrectableIOWriteErrorCount shall contain the number of the correctable write errors for the lifetime of the
	// storage pool.
	CorrectableIOWriteErrorCount int
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this storage pool.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RebuildErrorCount shall contain the number of rebuild errors over the lifetime of the storage pool.
	RebuildErrorCount int
	// StateChangeCount shall contain the number of state changes (changes in Status.State) for this storage pool.
	StateChangeCount int
	// UncorrectableIOReadErrorCount shall contain the number of the uncorrectable read errors for the lifetime of the
	// storage pool.
	UncorrectableIOReadErrorCount int
	// UncorrectableIOWriteErrorCount shall contain the number of the uncorrectable write errors for the lifetime of
	// the storage pool.
	UncorrectableIOWriteErrorCount int
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
