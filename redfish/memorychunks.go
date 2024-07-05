//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
	// RemoteMediaLocation The memory chunk was created using remote media accessible through a fabric.
	RemoteMediaLocation MediaLocation = "Remote"
	// MixedMediaLocation The memory chunk was created using both local media and remote media accessible through a
	// fabric.
	MixedMediaLocation MediaLocation = "Mixed"
)

type OperationalState string

const (
	// OnlineOperationalState Memory chunk can be used.
	OnlineOperationalState OperationalState = "Online"
	// OfflineOperationalState Memory chunk cannot be used. Consumers of this memory chunk should perform cleanup
	// operations as needed to prepare for the removal of this memory chunk.
	OfflineOperationalState OperationalState = "Offline"
)

// InterleaveSet shall describe an interleave set of which the memory chunk is a part.
type InterleaveSet struct {
	// Memory shall contain the memory device to which these settings apply.
	memory string
	// MemoryLevel shall contain the level of this interleave set for multi-level tiered memory.
	MemoryLevel int
	// OffsetMiB shall contain the offset within the DIMM that corresponds to the start of this memory region, with
	// units in MiB.
	OffsetMiB int
	// RegionID shall contain the DIMM region identifier.
	RegionID string
	// SizeMiB shall contain the size of this memory region, with units in MiB.
	SizeMiB int
}

// UnmarshalJSON unmarshals a InterleaveSet object from the raw JSON.
func (interleaveset *InterleaveSet) UnmarshalJSON(b []byte) error {
	type temp InterleaveSet
	var t struct {
		temp
		Memory common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*interleaveset = InterleaveSet(t.temp)

	// Extract the links to other entities for later
	interleaveset.memory = t.Memory.String()

	return nil
}

// Memory gets the associated memory device.
func (interleaveset *InterleaveSet) Memory(c common.Client) (*Memory, error) {
	if interleaveset.memory == "" {
		return nil, nil
	}
	return GetMemory(c, interleaveset.memory)
}

// MemoryChunks shall represent memory chunks and interleave sets in a Redfish implementation.
type MemoryChunks struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AddressRangeOffsetMiB shall be the offset of the memory chunk in the address range in MiB.
	AddressRangeOffsetMiB int
	// AddressRangeType shall contain the type of memory chunk.
	AddressRangeType AddressRangeType
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-configurable string to name the memory chunk.
	DisplayName string
	// InterleaveSets shall represent the interleave sets for the memory chunk. If not specified by the client during a
	// create operation, the memory chunk shall be created across all available memory within the memory domain.
	InterleaveSets []InterleaveSet
	// IsMirrorEnabled shall indicate whether memory mirroring is enabled for this memory chunk.
	IsMirrorEnabled bool
	// IsSpare shall indicate whether sparing is enabled for this memory chunk.
	IsSpare bool
	// MediaLocation shall contain the location of the memory media for this memory chunk.
	MediaLocation MediaLocation
	// MemoryChunkSizeMiB shall contain the size of the memory chunk in MiB.
	MemoryChunkSizeMiB int
	// RequestedOperationalState shall contain the requested operational state of this memory chunk.
	RequestedOperationalState OperationalState
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	cxlLogicalDevices []string
	// CXLLogicalDevicesCount is the number of CXL logical devices associated with this memory chunk.
	CXLLogicalDevicesCount int
	endpoints              []string
	// EndpointsCount is the number of endpoints with which this memory chunk is associated.
	EndpointsCount int
	memoryRegions  []string
	// MemoryRegionsCount is the number of memory regions for which this memory chunk provides capacity.
	MemoryRegionsCount int
}

// UnmarshalJSON unmarshals a MemoryChunks object from the raw JSON.
func (memorychunks *MemoryChunks) UnmarshalJSON(b []byte) error {
	type temp MemoryChunks
	type Links struct {
		// CXLLogicalDevices shall contain an array of links to resources of type CXLLogicalDevice that represent the CXL
		// logical devices associated with this memory chunk.
		CXLLogicalDevices      common.Links
		CXLLogicalDevicesCount int `json:"CXLLogicalDevices@odata.count"`
		// Endpoints shall contain a link to the resources of type Endpoint with which this memory chunk is associated.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
		// MemoryRegions shall contain an array of links to resources of type MemoryRegion that represent the memory
		// regions for which this memory chunk provides capacity.
		MemoryRegions      common.Links
		MemoryRegionsCount int `json:"MemoryRegions@odata.count"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorychunks = MemoryChunks(t.temp)

	// Extract the links to other entities for later
	memorychunks.cxlLogicalDevices = t.Links.CXLLogicalDevices.ToStrings()
	memorychunks.CXLLogicalDevicesCount = t.Links.CXLLogicalDevicesCount
	memorychunks.endpoints = t.Links.Endpoints.ToStrings()
	memorychunks.EndpointsCount = t.Links.EndpointsCount
	memorychunks.memoryRegions = t.Links.MemoryRegions.ToStrings()
	memorychunks.MemoryRegionsCount = t.Links.MemoryRegionsCount

	// This is a read/write object, so we need to save the raw object data for later
	memorychunks.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (memorychunks *MemoryChunks) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(MemoryChunks)
	original.UnmarshalJSON(memorychunks.rawData)

	readWriteFields := []string{
		"DisplayName",
		"MediaLocation",
		"RequestedOperationalState",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memorychunks).Elem()

	return memorychunks.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemoryChunks will get a MemoryChunks instance from the service.
func GetMemoryChunks(c common.Client, uri string) (*MemoryChunks, error) {
	return common.GetObject[MemoryChunks](c, uri)
}

// ListReferencedMemoryChunks gets the collection of MemoryChunks from
// a provided reference.
func ListReferencedMemoryChunks(c common.Client, link string) ([]*MemoryChunks, error) {
	return common.GetCollectionObjects[MemoryChunks](c, link)
}
