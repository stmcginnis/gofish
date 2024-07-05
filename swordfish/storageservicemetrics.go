//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// StorageServiceMetrics shall contain the usage and health statistics for a storage service in a Redfish
// implementation.
type StorageServiceMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this storage service.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
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
