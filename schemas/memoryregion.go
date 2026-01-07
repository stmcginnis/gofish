//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #MemoryRegion.v1_0_3.MemoryRegion

package schemas

import (
	"encoding/json"
)

type RegionType string

const (
	// StaticRegionType Static memory region. The whole address range is always
	// covered by memory.
	StaticRegionType RegionType = "Static"
	// DynamicRegionType Dynamic memory region. The address range coverage may be
	// changed in the runtime.
	DynamicRegionType RegionType = "Dynamic"
)

// MemoryRegion shall represent a memory region in a Redfish implementation.
type MemoryRegion struct {
	Entity
	// BlockSizeMiB shall contain the memory region block size in mebibytes (MiB).
	// The region size, base offset, all extent sizes, and all extent base offsets
	// shall be aligned to this block size.
	BlockSizeMiB int
	// ExtentsCount shall contain the number of extents defined for this memory
	// region.
	ExtentsCount *int `json:",omitempty"`
	// HardwareManagedCoherencyRegion shall indicate whether the device manages the
	// cache coherency across hosts and thereby ensures that each host has a
	// consistent view of this memory region content as defined in the 'Flags'
	// field of 'Device Scoped Memory Affinity Structure' defined in the Coherent
	// Device Attribute Table (CDAT) Specification.
	HardwareManagedCoherencyRegion bool
	// MemoryChunks shall contain the set of memory chunks providing capacity for
	// this memory region.
	MemoryChunks []MemoryChunk
	// MemoryExtents shall contain the set of memory extents defining address
	// ranges available for an access in dynamic capacity memory regions.
	MemoryExtents []MemoryExtent
	// NonVolatileRegion shall indicate whether this memory region represents
	// non-volatile memory as defined in the 'Flags' field of 'Device Scoped Memory
	// Affinity Structure' defined in the Coherent Device Attribute Table (CDAT)
	// Specification.
	NonVolatileRegion bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RegionBaseOffsetMiB shall contain the offset of the memory region in the
	// device address range in mebibytes (MiB).
	RegionBaseOffsetMiB int
	// RegionNumber shall contain the memory region number.
	RegionNumber int
	// RegionSizeMiB shall contain the size of the memory region in mebibytes
	// (MiB).
	RegionSizeMiB int
	// RegionType shall contain the type of memory region.
	RegionType RegionType
	// SanitizeOnRelease shall indicate whether the device has been configured such
	// that capacity released from this memory region will be sanitized before it
	// is made available to any host.
	SanitizeOnRelease bool
	// ShareableRegion shall indicate whether this memory region can be shared
	// across multiple hosts as defined in the 'Flags' field of 'Device Scoped
	// Memory Affinity Structure' defined in the Coherent Device Attribute Table
	// (CDAT) Specification.
	ShareableRegion bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MemoryRegion object from the raw JSON.
func (m *MemoryRegion) UnmarshalJSON(b []byte) error {
	type temp MemoryRegion
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryRegion(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MemoryRegion) Update() error {
	readWriteFields := []string{
		"BlockSizeMiB",
		"SanitizeOnRelease",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetMemoryRegion will get a MemoryRegion instance from the service.
func GetMemoryRegion(c Client, uri string) (*MemoryRegion, error) {
	return GetObject[MemoryRegion](c, uri)
}

// ListReferencedMemoryRegions gets the collection of MemoryRegion from
// a provided reference.
func ListReferencedMemoryRegions(c Client, link string) ([]*MemoryRegion, error) {
	return GetCollectionObjects[MemoryRegion](c, link)
}

// MemoryChunk shall contain the definition of a memory chunk providing capacity
// for memory region.
type MemoryChunk struct {
	// ChunkLink shall contain a link to a resource of type 'MemoryChunks' that
	// provides capacity to the memory region.
	chunkLink string
	// ChunkOffsetMiB shall be the offset of the memory chunk within the memory
	// region in mebibytes (MiB).
	ChunkOffsetMiB int
}

// UnmarshalJSON unmarshals a MemoryChunk object from the raw JSON.
func (m *MemoryChunk) UnmarshalJSON(b []byte) error {
	type temp MemoryChunk
	var tmp struct {
		temp
		ChunkLink Link `json:"ChunkLink"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryChunk(tmp.temp)

	// Extract the links to other entities for later
	m.chunkLink = tmp.ChunkLink.String()

	return nil
}

// ChunkLink gets the ChunkLink linked resource.
func (m *MemoryChunk) ChunkLink(client Client) (*MemoryChunks, error) {
	if m.chunkLink == "" {
		return nil, nil
	}
	return GetObject[MemoryChunks](client, m.chunkLink)
}

// MemoryExtent shall contain the definition of a memory extent identifying an
// available address range in the dynamic capacity memory region.
type MemoryExtent struct {
	// ExtentOffsetMiB shall be the offset of the memory extent within the memory
	// region in mebibytes (MiB).
	ExtentOffsetMiB int
	// ExtentSizeMiB shall contain the size of the memory extent in MiB.
	ExtentSizeMiB int
	// SequenceNumber shall contain the sequence number instructing host on the
	// relative order the extents have to be placed in the host adjacent virtual
	// address space.
	SequenceNumber *int `json:",omitempty"`
	// Tag shall contain an opaque context attached to each extent to track usage
	// of each extent or map extent to specific processes, transactions, or
	// workloads on the host.
	Tag string
}
