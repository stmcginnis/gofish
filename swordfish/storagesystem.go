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
func ListReferencedStorageSystems(c common.Client, link string) ([]*StorageSystem, error) {
	return common.GetCollectionObjects(c, link, GetStorageSystem)
}
