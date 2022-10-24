//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// StorageSystem is a Swordfish storage system instance.
type StorageSystem struct {
	redfish.ComputerSystem
}

// GetStorageSystem will get a StorageSystem instance from the Swordfish service.
func GetStorageSystem(c common.Client, uri string) (*StorageSystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storageSystem StorageSystem
	err = json.NewDecoder(resp.Body).Decode(&storageSystem)
	if err != nil {
		return nil, err
	}

	storageSystem.SetClient(c)
	return &storageSystem, nil
}

// ListReferencedStorageSystems gets the collection of StorageSystems.
func ListReferencedStorageSystems(c common.Client, link string) ([]*StorageSystem, error) { //nolint:dupl
	var result []*StorageSystem
	if link == "" {
		return result, nil
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

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
