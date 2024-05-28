//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// StorageSystem is a Swordfish storage system instance.
type StorageSystem struct {
	redfish.ComputerSystem
}

// GetStorageSystem will get a StorageSystem instance from the Swordfish service.
func GetStorageSystem(c common.Client, uri string) (*StorageSystem, error) {
	var storageSystem StorageSystem
	return &storageSystem, storageSystem.Get(c, uri, &storageSystem)
}

// ListReferencedStorageSystems gets the collection of StorageSystems.
func ListReferencedStorageSystems(c common.Client, link string) ([]*StorageSystem, error) { //nolint:dupl
	if link == "" {
		return nil, nil
	}

	type GetResult struct {
		Item  *StorageSystem
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		storagesystem, err := GetStorageSystem(c, link)
		ch <- GetResult{Item: storagesystem, Link: link, Error: err}
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

	// Save unordered results into link-to-StorageSystem helper map.
	unorderedResults := map[string]*StorageSystem{}
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
	results := make([]*StorageSystem, len(links))
	for i, link := range links {
		results[i] = unorderedResults[link]
	}

	return results, nil
}
