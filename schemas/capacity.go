//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.2 - #Capacity.v1_2_2.CapacitySource

package schemas

import (
	"encoding/json"
)

// Capacity shall be equal to the total capacity for the data store.
type Capacity struct {
	// Data shall be capacity information relating to provisioned user data.
	Data CapacityInfo
	// IsThinProvisioned shall be fully allocated. The default value shall be
	// false.
	IsThinProvisioned bool
	// Metadata shall be capacity information relating to provisioned system
	// (non-user accessible) data.
	Metadata CapacityInfo
	// Snapshot shall be capacity information relating to provisioned snapshot or
	// backup data.
	Snapshot CapacityInfo
}

// CapacityInfo This composition may be used to represent the utilization of
// storage capacity.
type CapacityInfo struct {
	// AllocatedBytes shall be the number of bytes currently allocated by the
	// storage system in this data store for this data type.
	AllocatedBytes *int `json:",omitempty"`
	// ConsumedBytes shall be the number of logical bytes currently consumed in
	// this data store for this data type.
	ConsumedBytes *int `json:",omitempty"`
	// GuaranteedBytes shall be the number of bytes the storage system guarantees
	// can be allocated in this data store for this data type.
	GuaranteedBytes *int `json:",omitempty"`
	// ProvisionedBytes shall be the maximum number of bytes that can be allocated
	// in this data store for this data type.
	ProvisionedBytes *int `json:",omitempty"`
}

// CapacitySource This composition may be used to represent the source and type
// of storage capacity. At most one of the ProvidingDrives, ProvidingVolumes,
// ProvidingMemoryChunks, ProvidingMemory or ProvidingPools properties may have
// a value. If any of ProvidingDrives, ProvidingVolumes, ProvidingMemory or
// ProvidingPools reference more than one resource, allocation of capacity
// across those resources is implementation dependent.
type CapacitySource struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProvidedCapacity shall be the amount of space that has been provided from
	// the ProvidingDrives, ProvidingVolumes, ProvidingMemory or ProvidingPools.
	ProvidedCapacity Capacity
	// ProvidedClassOfService shall reference the provided ClassOfService from the
	// ProvidingDrives, ProvidingVolumes, ProvidingMemoryChunks, ProvidingMemory or
	// ProvidingPools.
	ProvidedClassOfService ClassOfService
	// ProvidingDrives shall be a reference to a contributing drive or drives.
	providingDrives string
	// ProvidingMemory shall be a reference to the contributing memory.
	//
	// Version added: v1.1.0
	providingMemory string
	// ProvidingMemoryChunks shall be a reference to the contributing memory
	// chunks.
	//
	// Version added: v1.1.0
	providingMemoryChunks string
	// ProvidingPools shall be a reference to a contributing storage pool or
	// storage pools.
	providingPools string
	// ProvidingVolumes shall be a reference to a contributing volume or volumes.
	providingVolumes string
}

// UnmarshalJSON unmarshals a CapacitySource object from the raw JSON.
func (c *CapacitySource) UnmarshalJSON(b []byte) error {
	type temp CapacitySource
	var tmp struct {
		temp
		ProvidingDrives       Link `json:"ProvidingDrives"`
		ProvidingMemory       Link `json:"ProvidingMemory"`
		ProvidingMemoryChunks Link `json:"ProvidingMemoryChunks"`
		ProvidingPools        Link `json:"ProvidingPools"`
		ProvidingVolumes      Link `json:"ProvidingVolumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CapacitySource(tmp.temp)

	// Extract the links to other entities for later
	c.providingDrives = tmp.ProvidingDrives.String()
	c.providingMemory = tmp.ProvidingMemory.String()
	c.providingMemoryChunks = tmp.ProvidingMemoryChunks.String()
	c.providingPools = tmp.ProvidingPools.String()
	c.providingVolumes = tmp.ProvidingVolumes.String()

	return nil
}

// GetCapacitySource will get a CapacitySource instance from the service.
func GetCapacitySource(c Client, uri string) (*CapacitySource, error) {
	return GetObject[CapacitySource](c, uri)
}

// ListReferencedCapacitySources gets the collection of CapacitySource from
// a provided reference.
func ListReferencedCapacitySources(c Client, link string) ([]*CapacitySource, error) {
	return GetCollectionObjects[CapacitySource](c, link)
}

// ProvidingDrives gets the ProvidingDrives collection.
func (c *CapacitySource) ProvidingDrives() ([]*Drive, error) {
	if c.providingDrives == "" {
		return nil, nil
	}
	return GetCollectionObjects[Drive](c.client, c.providingDrives)
}

// ProvidingMemory gets the ProvidingMemory collection.
func (c *CapacitySource) ProvidingMemory() ([]*Memory, error) {
	if c.providingMemory == "" {
		return nil, nil
	}
	return GetCollectionObjects[Memory](c.client, c.providingMemory)
}

// ProvidingMemoryChunks gets the ProvidingMemoryChunks collection.
func (c *CapacitySource) ProvidingMemoryChunks() ([]*MemoryChunks, error) {
	if c.providingMemoryChunks == "" {
		return nil, nil
	}
	return GetCollectionObjects[MemoryChunks](c.client, c.providingMemoryChunks)
}

// ProvidingPools gets the ProvidingPools collection.
func (c *CapacitySource) ProvidingPools() ([]*StoragePool, error) {
	if c.providingPools == "" {
		return nil, nil
	}
	return GetCollectionObjects[StoragePool](c.client, c.providingPools)
}

// ProvidingVolumes gets the ProvidingVolumes collection.
func (c *CapacitySource) ProvidingVolumes() ([]*Volume, error) {
	if c.providingVolumes == "" {
		return nil, nil
	}
	return GetCollectionObjects[Volume](c.client, c.providingVolumes)
}
