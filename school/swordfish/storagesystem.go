// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
	"github.com/stmcginnis/gofish/school/redfish"
)

// DefaultStorageSystemPath is the default URI for StorageSystem collections.
const DefaultStorageSystemPath = "/redfish/v1/StorageSystems"

// StorageSystem is a Swordfish storage system instance.
type StorageSystem struct {
	redfish.ComputerSystem
}

// GetStorageSystem will get a StorageSystem instance from the Swordfish service.
func GetStorageSystem(c common.Client, uri string) (*StorageSystem, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var StorageSystem StorageSystem
	err = json.NewDecoder(resp.Body).Decode(&StorageSystem)
	if err != nil {
		return nil, err
	}

	StorageSystem.SetClient(c)
	return &StorageSystem, nil
}

// ListReferencedStorageSystems gets the collection of StorageSystems.
func ListReferencedStorageSystems(c common.Client, link string) ([]*StorageSystem, error) {
	var result []*StorageSystem
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, StorageSystemLink := range links.ItemLinks {
		StorageSystem, err := GetStorageSystem(c, StorageSystemLink)
		if err != nil {
			return result, err
		}
		result = append(result, StorageSystem)
	}

	return result, nil
}

// ListStorageSystems gets all StorageSystem in the system.
func ListStorageSystems(c common.Client) ([]*StorageSystem, error) {
	return ListReferencedStorageSystems(c, DefaultStorageSystemPath)
}
