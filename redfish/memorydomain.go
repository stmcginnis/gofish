//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// MemoryDomain is used to represent Memory Domains.
type MemoryDomain struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowsBlockProvisioning shall indicate if this Memory Domain supports the
	// creation of Blocks of memory.
	AllowsBlockProvisioning bool
	// AllowsMemoryChunkCreation shall indicate if this Memory Domain supports
	// the creation of Memory Chunks.
	AllowsMemoryChunkCreation bool
	// AllowsMirroring shall indicate if this Memory Domain supports the
	// creation of Memory Chunks with mirroring enabled.
	AllowsMirroring bool
	// AllowsSparing shall indicate if this Memory Domain supports the creation
	// of Memory Chunks with sparing enabled.
	AllowsSparing bool
	// Description provides a description of this resource.
	Description string
	// InterleavableMemorySets shall represent the interleave sets for the
	// memory chunk.
	InterleavableMemorySets []MemorySet
	// MemoryChunkIncrementMiB shall contain the incremental size, from MemoryChunkIncrementMiB, allowed for a memory
	// chunk within this domain in mebibytes (MiB).
	MemoryChunkIncrementMiB int
	// memoryChunks shall be a link to a collection of type MemoryChunkCollection.
	memoryChunks string
	// MemorySizeMiB shall contain the total size of the memory domain in mebibytes (MiB).
	MemorySizeMiB int
	// MinMemoryChunkSizeMiB shall contain the minimum size allowed for a memory chunk within this domain in mebibytes
	// (MiB).
	MinMemoryChunkSizeMiB int
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	cxlLogicalDevices []string
	// CXLLogicalDevicesCount is the number of CXL logical devices that are associated with this memory domain.
	CXLLogicalDevicesCount int
	fabricAdapters         []string
	// FabricAdaptersCount is the number of fabric adapters that present this memory domain to a fabric.
	FabricAdaptersCount int
	mediaControllers    []string
	// MediaControllersCount is the number of media controllers for this memory domain.
	// This property has been deprecated in favor of the FabricAdapters property.
	MediaControllersCount int
	pcieFunctions         []string
	// PCIeFunctionsCount is the number of PCIe functions representing this memory domain.
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a MemoryDomain object from the raw JSON.
func (memorydomain *MemoryDomain) UnmarshalJSON(b []byte) error {
	type temp MemoryDomain
	type Links struct {
		// CXLLogicalDevices shall contain an array of links to resources of type CXLLogicalDevice that represent the CXL
		// logical devices that are associated with this memory domain.
		CXLLogicalDevices      common.Links
		CXLLogicalDevicesCount int `json:"CXLLogicalDevices@odata.count"`
		// FabricAdapters shall contain an array of links to resources of type FabricAdapter that represent the fabric
		// adapters that present this memory domain to a fabric.
		FabricAdapters      common.Links
		FabricAdaptersCount int `json:"FabricAdapters@odata.count"`
		// MediaControllers is array of links to the media controllers for this memory domain.
		// This property has been deprecated in favor of the FabricAdapters property.
		MediaControllers      common.Links
		MediaControllersCount int `json:"MediaControllers@odata.count"`
		// PCIeFunctions shall contain an array of links to resources of type PCIeFunction that represent the PCIe
		// functions representing this memory domain.
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	}
	var t struct {
		temp
		MemoryChunks common.Link
		Links        Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*memorydomain = MemoryDomain(t.temp)
	memorydomain.memoryChunks = t.MemoryChunks.String()

	memorydomain.cxlLogicalDevices = t.Links.CXLLogicalDevices.ToStrings()
	memorydomain.CXLLogicalDevicesCount = t.Links.CXLLogicalDevicesCount
	memorydomain.fabricAdapters = t.Links.FabricAdapters.ToStrings()
	memorydomain.FabricAdaptersCount = t.Links.FabricAdaptersCount
	memorydomain.mediaControllers = t.Links.MediaControllers.ToStrings()
	memorydomain.MediaControllersCount = t.Links.MediaControllersCount
	memorydomain.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	memorydomain.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	return nil
}

// CXLLogicalDevices gets the CXLLogicalDevice that represent the CXL logical devices
// that are associated with this memory domain.
func (memorydomain *MemoryDomain) CXLLogicalDevices() ([]*CXLLogicalDevice, error) {
	return common.GetObjects[CXLLogicalDevice](memorydomain.GetClient(), memorydomain.cxlLogicalDevices)
}

// FabricAdapters gets the fabric adapters that present this memory domain to a fabric.
func (memorydomain *MemoryDomain) FabricAdapters() ([]*FabricAdapter, error) {
	return common.GetObjects[FabricAdapter](memorydomain.GetClient(), memorydomain.fabricAdapters)
}

// MediaControllers gets the media controllers for this memory domain.
// This property has been deprecated in favor of the FabricAdapters property.
func (memorydomain *MemoryDomain) MediaControllers() ([]*MediaController, error) {
	return common.GetObjects[MediaController](memorydomain.GetClient(), memorydomain.mediaControllers)
}

// PCIeFunctions gets the PCIe functions representing this memory domain.
func (memorydomain *MemoryDomain) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](memorydomain.GetClient(), memorydomain.pcieFunctions)
}

// GetMemoryDomain will get a MemoryDomain instance from the service.
func GetMemoryDomain(c common.Client, uri string) (*MemoryDomain, error) {
	return common.GetObject[MemoryDomain](c, uri)
}

// ListReferencedMemoryDomains gets the collection of MemoryDomain from
// a provided reference.
func ListReferencedMemoryDomains(c common.Client, link string) ([]*MemoryDomain, error) {
	return common.GetCollectionObjects[MemoryDomain](c, link)
}

// MemorySet shall represent the interleave sets for a memory chunk.
type MemorySet struct {
	// MemorySet shall be links to objects of type Memory.
	memorySet []string
	// MemorySetCount is the number of memory sets.
	MemorySetCount int `json:"MemorySet@odata.count"`
}

// UnmarshalJSON unmarshals a MemorySet object from the raw JSON.
func (memoryset *MemorySet) UnmarshalJSON(b []byte) error {
	type temp MemorySet
	var t struct {
		temp
		MemorySet common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*memoryset = MemorySet(t.temp)
	memoryset.memorySet = t.MemorySet.ToStrings()

	return nil
}

// MemorySet gets the Memory objects that are part of this set.
func (memoryset *MemorySet) MemorySet(c common.Client) ([]*Memory, error) {
	return common.GetObjects[Memory](c, memoryset.memorySet)
}
