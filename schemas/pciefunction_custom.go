//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// StorageControllers gets the StorageControllers linked resources.
//
// The schema links these to Storage.StorageController, the embedded element of
// the deprecated Storage.StorageControllers array. Services link to standalone
// StorageController resources (Storage.Controllers), so this returns
// StorageController.
func (p *PCIeFunction) StorageControllers() ([]*StorageController, error) {
	return GetObjects[StorageController](p.client, p.storageControllers)
}
