//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
)

// StorageSystem is a Swordfish storage system instance.
type StorageSystem struct {
	redfish.ComputerSystem
}

// GetStorageSystem will get a StorageSystem instance from the Swordfish service.
func GetStorageSystem(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageSystem, error) {
	return common.GetObject[StorageSystem](c, uri, queryOpts...)
}

// ListReferencedStorageSystems gets the collection of StorageSystems.
func ListReferencedStorageSystems(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageSystem, error) {
	return common.GetCollectionObjects[StorageSystem](c, link, queryOpts...)
}
