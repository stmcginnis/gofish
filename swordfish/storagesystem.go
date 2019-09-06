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
func ListReferencedStorageSystems(c common.Client, link string) ([]*StorageSystem, error) {
	var result []*StorageSystem
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storageSystemLink := range links.ItemLinks {
		storageSystem, err := GetStorageSystem(c, storageSystemLink)
		if err != nil {
			return result, err
		}
		result = append(result, storageSystem)
	}

	return result, nil
}
