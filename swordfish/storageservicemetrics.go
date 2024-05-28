//
// SPDX-License-Identifier: BSD-3-Clause
//

//nolint:dupl
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
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storageservicemetrics StorageServiceMetrics
	err = json.NewDecoder(resp.Body).Decode(&storageservicemetrics)
	if err != nil {
		return nil, err
	}

	storageservicemetrics.SetClient(c)
	return &storageservicemetrics, nil
}

// ListReferencedStorageServiceMetricss gets the collection of StorageServiceMetrics from
// a provided reference.
func ListReferencedStorageServiceMetricss(c common.Client, link string) ([]*StorageServiceMetrics, error) { //nolint:dupl
	if link == "" {
		return nil, nil
	}

	type GetResult struct {
		Item  *StorageServiceMetrics
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		storageservicemetrics, err := GetStorageServiceMetrics(c, link)
		ch <- GetResult{Item: storageservicemetrics, Link: link, Error: err}
	}

	var links []string
	var err error
	go func() {
		links, err = common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	// Save unordered results into link-to-StorageServiceMetrics helper map.
	unorderedResults := map[string]*StorageServiceMetrics{}
	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			unorderedResults[r.Link] = r.Item
		}
	}

	if !collectionError.Empty() {
		return nil, collectionError
	}
	// Build the final ordered slice based on the original order from the links list.
	results := make([]*StorageServiceMetrics, len(links))
	for i, link := range links {
		results[i] = unorderedResults[link]
	}

	return results, nil
}
