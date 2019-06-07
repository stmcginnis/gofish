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
)

// Capacity is used to represent storage capacity.  The sum of the values
// in Data, Metadata, and Snapshot shall be equal to the total capacity
// for the data store.
type Capacity struct {
	// Data shall be capacity information relating to
	// provisioned user data.
	Data CapacityInfo
	// IsThinProvisioned is If the value is false, the capacity shall be
	// fully allocated.  The default value shall be false.
	IsThinProvisioned bool
	// Metadata shall be capacity information relating to
	// provisioned system (non-user accessible) data.
	Metadata CapacityInfo
	// Snapshot shall be capacity information relating to
	// provisioned snapshot or backup data.
	Snapshot CapacityInfo
}

// CapacityInfo is used to represent the utilization of storage capacity.
type CapacityInfo struct {
	// AllocatedBytes shall be the number of bytes currently
	// allocated by the storage system in this data store for this data type.
	AllocatedBytes int
	// ConsumedBytes shall be the number of logical bytes
	// currently consumed in this data store for this data type.
	ConsumedBytes int
	// GuaranteedBytes shall be the number of bytes the storage
	// system guarantees can be allocated in this data store for this data
	// type.
	GuaranteedBytes int
	// ProvisionedBytes shall be the maximum number of bytes
	// that can be allocated in this data store for this data type.
	ProvisionedBytes int
}

// CapacitySource is used to represent the source and type of storage
// capacity.  At most one of the ProvidingDrives, ProvidingVolumes,
// ProvidingMemoryChunks, ProvidingMemory or ProvidingPools properties
// may have a value.  If any of ProvidingDrives, ProvidingVolumes,
// ProvidingMemory or ProvidingPools reference more than one resource,
// allocation of capacity across those resources is implementation
// dependent.
type CapacitySource struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// ProvidedCapacity shall be the amount of space that has
	// been provided from the ProvidingDrives, ProvidingVolumes,
	// ProvidingMemory or ProvidingPools.
	ProvidedCapacity Capacity
	// ProvidedClassOfService shall reference the provided
	// ClassOfService from the ProvidingDrives, ProvidingVolumes,
	// ProvidingMemoryChunks, ProvidingMemory or ProvidingPools.
	ProvidedClassOfService ClassesOfService
	// ProvidingDrives is If present, the value shall be a reference to a
	// contributing drive or drives.
	// providingDrives DriveCollection
	// ProvidingMemory is If present, the value shall be a reference to the
	// contributing memory.
	// providingMemory MemoryCollection
	// ProvidingMemoryChunks is If present, the value shall be a reference to
	// the contributing memory chunks.
	// providingMemoryChunks MemoryChunksCollection
	// ProvidingPools is If present, the value shall be a reference to a
	// contributing storage pool or storage pools.
	// providingPools StoragePoolCollection
	// ProvidingVolumes is If present, the value shall be a reference to a
	// contributing volume or volumes.
	// providingVolumes VolumeCollection
}

// UnmarshalJSON unmarshals a CapacitySource object from the raw JSON.
func (capacitysource *CapacitySource) UnmarshalJSON(b []byte) error {
	type temp CapacitySource
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*capacitysource = CapacitySource(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetCapacitySource will get a CapacitySource instance from the service.
func GetCapacitySource(c common.Client, uri string) (*CapacitySource, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var capacitysource CapacitySource
	err = json.NewDecoder(resp.Body).Decode(&capacitysource)
	if err != nil {
		return nil, err
	}

	capacitysource.SetClient(c)
	return &capacitysource, nil
}

// ListReferencedCapacitySources gets the collection of CapacitySources from
// a provided reference.
func ListReferencedCapacitySources(c common.Client, link string) ([]*CapacitySource, error) {
	var result []*CapacitySource
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, capSourceLink := range links.ItemLinks {
		capSource, err := GetCapacitySource(c, capSourceLink)
		if err != nil {
			return result, err
		}
		result = append(result, capSource)
	}

	return result, nil
}
