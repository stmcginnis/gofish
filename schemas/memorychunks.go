//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/MemoryChunks.v1_6_2.json
// 2023.2 - #MemoryChunks.v1_6_2.MemoryChunks

package schemas

import (
	"encoding/json"
)

type AddressRangeType string

const (
	// VolatileAddressRangeType Volatile memory.
	VolatileAddressRangeType AddressRangeType = "Volatile"
	// PMEMAddressRangeType Byte accessible persistent memory.
	PMEMAddressRangeType AddressRangeType = "PMEM"
	// BlockAddressRangeType Block accessible memory.
	BlockAddressRangeType AddressRangeType = "Block"
)

type MediaLocation string

const (
	// LocalMediaLocation The memory chunk was created using local media.
	LocalMediaLocation MediaLocation = "Local"
	// RemoteMediaLocation The memory chunk was created using remote media
	// accessible through a fabric.
	RemoteMediaLocation MediaLocation = "Remote"
	// MixedMediaLocation The memory chunk was created using both local media and
	// remote media accessible through a fabric.
	MixedMediaLocation MediaLocation = "Mixed"
)

type OperationalState string

const (
	// OnlineOperationalState Memory chunk can be used.
	OnlineOperationalState OperationalState = "Online"
	// OfflineOperationalState Memory chunk cannot be used. Consumers of this
	// memory chunk should perform cleanup operations as needed to prepare for the
	// removal of this memory chunk.
	OfflineOperationalState OperationalState = "Offline"
)

// MemoryChunks shall represent memory chunks and interleave sets in a Redfish
// implementation.
type MemoryChunks struct {
	Entity
	// AddressRangeOffsetMiB shall be the offset of the memory chunk in the address
	// range in MiB.
	//
	// Version added: v1.3.0
	AddressRangeOffsetMiB *int `json:",omitempty"`
	// AddressRangeType shall contain the type of memory chunk.
	AddressRangeType AddressRangeType
	// DisplayName shall contain a user-configurable string to name the memory
	// chunk.
	//
	// Version added: v1.4.0
	DisplayName string
	// InterleaveSets shall represent the interleave sets for the memory chunk. If
	// not specified by the client during a create operation, the memory chunk
	// shall be created across all available memory within the memory domain.
	InterleaveSets []InterleaveSet
	// IsMirrorEnabled shall indicate whether memory mirroring is enabled for this
	// memory chunk.
	IsMirrorEnabled bool
	// IsSpare shall indicate whether sparing is enabled for this memory chunk.
	IsSpare bool
	// MediaLocation shall contain the location of the memory media for this memory
	// chunk.
	//
	// Version added: v1.5.0
	MediaLocation MediaLocation
	// MemoryChunkSizeMiB shall contain the size of the memory chunk in MiB.
	MemoryChunkSizeMiB *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RequestedOperationalState shall contain the requested operational state of
	// this memory chunk.
	//
	// Version added: v1.5.0
	RequestedOperationalState OperationalState
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.2.0
	Status Status
	// cXLLogicalDevices are the URIs for CXLLogicalDevices.
	cXLLogicalDevices []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// memoryRegions are the URIs for MemoryRegions.
	memoryRegions []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MemoryChunks object from the raw JSON.
func (m *MemoryChunks) UnmarshalJSON(b []byte) error {
	type temp MemoryChunks
	type mLinks struct {
		CXLLogicalDevices Links `json:"CXLLogicalDevices"`
		Endpoints         Links `json:"Endpoints"`
		MemoryRegions     Links `json:"MemoryRegions"`
	}
	var tmp struct {
		temp
		Links mLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryChunks(tmp.temp)

	// Extract the links to other entities for later
	m.cXLLogicalDevices = tmp.Links.CXLLogicalDevices.ToStrings()
	m.endpoints = tmp.Links.Endpoints.ToStrings()
	m.memoryRegions = tmp.Links.MemoryRegions.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MemoryChunks) Update() error {
	readWriteFields := []string{
		"DisplayName",
		"MediaLocation",
		"RequestedOperationalState",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetMemoryChunks will get a MemoryChunks instance from the service.
func GetMemoryChunks(c Client, uri string) (*MemoryChunks, error) {
	return GetObject[MemoryChunks](c, uri)
}

// ListReferencedMemoryChunkss gets the collection of MemoryChunks from
// a provided reference.
func ListReferencedMemoryChunkss(c Client, link string) ([]*MemoryChunks, error) {
	return GetCollectionObjects[MemoryChunks](c, link)
}

// CXLLogicalDevices gets the CXLLogicalDevices linked resources.
func (m *MemoryChunks) CXLLogicalDevices() ([]*CXLLogicalDevice, error) {
	return GetObjects[CXLLogicalDevice](m.client, m.cXLLogicalDevices)
}

// Endpoints gets the Endpoints linked resources.
func (m *MemoryChunks) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](m.client, m.endpoints)
}

// MemoryRegions gets the MemoryRegions linked resources.
func (m *MemoryChunks) MemoryRegions() ([]*MemoryRegion, error) {
	return GetObjects[MemoryRegion](m.client, m.memoryRegions)
}

// InterleaveSet shall describe an interleave set of which the memory chunk is a
// part.
type InterleaveSet struct {
	// Memory shall contain the memory device to which these settings apply.
	memory string
	// MemoryLevel shall contain the level of this interleave set for multi-level
	// tiered memory.
	MemoryLevel *int `json:",omitempty"`
	// OffsetMiB shall contain the offset within the DIMM that corresponds to the
	// start of this memory region, with units in MiB.
	OffsetMiB *int `json:",omitempty"`
	// RegionID shall contain the DIMM region identifier.
	RegionID string `json:"RegionId"`
	// SizeMiB shall contain the size of this memory region, with units in MiB.
	SizeMiB *int `json:",omitempty"`
}

// UnmarshalJSON unmarshals a InterleaveSet object from the raw JSON.
func (i *InterleaveSet) UnmarshalJSON(b []byte) error {
	type temp InterleaveSet
	var tmp struct {
		temp
		Memory Link `json:"Memory"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = InterleaveSet(tmp.temp)

	// Extract the links to other entities for later
	i.memory = tmp.Memory.String()

	return nil
}

// Memory gets the Memory linked resource.
func (i *InterleaveSet) Memory(client Client) (*Entity, error) {
	if i.memory == "" {
		return nil, nil
	}
	return GetObject[Entity](client, i.memory)
}
