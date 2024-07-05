//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type RegionType string

const (
	// StaticRegionType Static memory region. The whole address range is always covered by memory.
	StaticRegionType RegionType = "Static"
	// DynamicRegionType Dynamic memory region. The address range coverage may be changed in the runtime.
	DynamicRegionType RegionType = "Dynamic"
)

// MemoryChunk shall contain the definition of a memory chunk providing capacity for memory region.
type MemoryChunk struct {
	// ChunkLink shall contain a link to a resource of type MemoryChunks that provides capacity to the memory region.
	chunkLink string
	// ChunkOffsetMiB shall be the offset of the memory chunk within the memory region in mebibytes (MiB).
	ChunkOffsetMiB int
}

// UnmarshalJSON unmarshals a MemoryChunk object from the raw JSON.
func (memorychunk *MemoryChunk) UnmarshalJSON(b []byte) error {
	type temp MemoryChunk
	var t struct {
		temp
		ChunkLink common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorychunk = MemoryChunk(t.temp)

	// Extract the links to other entities for later
	memorychunk.chunkLink = t.ChunkLink.String()

	return nil
}

// MemoryChunks gets the memory chunks providing capacity.
func (memorychunk *MemoryChunk) MemoryChunks(c common.Client) (*MemoryChunks, error) {
	if memorychunk.chunkLink == "" {
		return nil, nil
	}

	return GetMemoryChunks(c, memorychunk.chunkLink)
}

// MemoryExtent shall contain the definition of a memory extent identifying an available address range in the
// dynamic capacity memory region.
type MemoryExtent struct {
	// ExtentOffsetMiB shall be the offset of the memory extent within the memory region in mebibytes (MiB).
	ExtentOffsetMiB int
	// ExtentSizeMiB shall contain the size of the memory extent in MiB.
	ExtentSizeMiB int
	// SequenceNumber shall contain the sequence number instructing host on the relative order the extents have to be
	// placed in the host adjacent virtual address space.
	SequenceNumber int
	// Tag shall contain an opaque context attached to each extent to track usage of each extent or map extent to
	// specific processes, transactions, or workloads on the host.
	Tag string
}

// MemoryRegion shall represent memory region in a Redfish implementation.
type MemoryRegion struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BlockSizeMiB shall contain the memory region block size in mebibytes (MiB). The region size, base offset, all
	// extent sizes, and all extent base offsets shall be aligned to this block size.
	BlockSizeMiB int
	// Description provides a description of this resource.
	Description string
	// ExtentsCount shall contain the number of extents defined for this memory region.
	ExtentsCount int
	// HardwareManagedCoherencyRegion shall indicate whether the device manages the cache coherency across hosts and
	// thereby ensures that each host has a consistent view of this memory region content as defined in the 'Flags'
	// field of 'Device Scoped Memory Affinity Structure' defined in the Coherent Device Attribute Table (CDAT)
	// Specification.
	HardwareManagedCoherencyRegion bool
	// MemoryChunks shall contain the set of memory chunks providing capacity for this memory region.
	MemoryChunks []MemoryChunk
	// MemoryExtents shall contain the set of memory extents defining address ranges available for an access in dynamic
	// capacity memory regions.
	MemoryExtents []MemoryExtent
	// NonVolatileRegion shall indicate whether this memory region represents non-volatile memory as defined in the
	// 'Flags' field of 'Device Scoped Memory Affinity Structure' defined in the Coherent Device Attribute Table (CDAT)
	// Specification.
	NonVolatileRegion bool
	// RegionBaseOffsetMiB shall contain the offset of the memory region in the device address range in mebibytes
	// (MiB).
	RegionBaseOffsetMiB int
	// RegionNumber shall contain the memory region number.
	RegionNumber int
	// RegionSizeMiB shall contain the size of the memory region in mebibytes (MiB).
	RegionSizeMiB int
	// RegionType shall contain the type of memory region.
	RegionType RegionType
	// SanitizeOnRelease shall indicate whether the device has been configured such that capacity released from this
	// memory region will be sanitized before it is made available to any host.
	SanitizeOnRelease bool
	// ShareableRegion shall indicate whether this memory region can be shared across multiple hosts as defined in the
	// 'Flags' field of 'Device Scoped Memory Affinity Structure' defined in the Coherent Device Attribute Table (CDAT)
	// Specification.
	ShareableRegion bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MemoryRegion object from the raw JSON.
func (memoryregion *MemoryRegion) UnmarshalJSON(b []byte) error {
	type temp MemoryRegion
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memoryregion = MemoryRegion(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	memoryregion.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (memoryregion *MemoryRegion) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(MemoryRegion)
	original.UnmarshalJSON(memoryregion.rawData)

	readWriteFields := []string{
		"BlockSizeMiB",
		"SanitizeOnRelease",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memoryregion).Elem()

	return memoryregion.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemoryRegion will get a MemoryRegion instance from the service.
func GetMemoryRegion(c common.Client, uri string) (*MemoryRegion, error) {
	return common.GetObject[MemoryRegion](c, uri)
}

// ListReferencedMemoryRegions gets the collection of MemoryRegion from
// a provided reference.
func ListReferencedMemoryRegions(c common.Client, link string) ([]*MemoryRegion, error) {
	return common.GetCollectionObjects[MemoryRegion](c, link)
}
