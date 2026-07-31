//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
)

// StorageSystem is a Swordfish storage system instance.
type StorageSystem struct {
	redfish.ComputerSystem
}

// GetStorageSystem will get a StorageSystem instance from the Swordfish service.
func GetStorageSystem(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageSystem, error) {
	return GetStorageSystemWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetStorageSystemWithContext will get a StorageSystem instance from the Swordfish service.
func GetStorageSystemWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageSystem, error) {
	return common.GetObjectWithContext[StorageSystem](ctx, c, uri, queryOpts...)
}

// ListReferencedStorageSystems gets the collection of StorageSystems.
func ListReferencedStorageSystems(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageSystem, error) {
	return ListReferencedStorageSystemsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedStorageSystemsWithContext gets the collection of StorageSystems.
func ListReferencedStorageSystemsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageSystem, error) {
	return common.GetCollectionObjectsWithContext[StorageSystem](ctx, c, link, queryOpts...)
}
