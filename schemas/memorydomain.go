//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #MemoryDomain.v1_5_1.MemoryDomain

package schemas

import (
	"encoding/json"
)

// MemoryDomain shall represent memory domains in a Redfish implementation.
type MemoryDomain struct {
	Entity
	// AllowsBlockProvisioning shall indicate whether this memory domain supports
	// the creation of blocks of memory.
	AllowsBlockProvisioning bool
	// AllowsMemoryChunkCreation shall indicate whether this memory domain supports
	// the creation of memory chunks.
	AllowsMemoryChunkCreation bool
	// AllowsMirroring shall indicate whether this memory domain supports the
	// creation of memory chunks with mirroring enabled.
	//
	// Version added: v1.1.0
	AllowsMirroring bool
	// AllowsSparing shall indicate whether this memory domain supports the
	// creation of memory chunks with sparing enabled.
	//
	// Version added: v1.1.0
	AllowsSparing bool
	// InterleavableMemorySets shall represent the interleave sets for the memory
	// chunk.
	InterleavableMemorySets []MemorySet
	// MemoryChunkIncrementMiB shall contain the incremental size, from
	// 'MemoryChunkIncrementMiB', allowed for a memory chunk within this domain in
	// mebibytes (MiB).
	//
	// Version added: v1.5.0
	MemoryChunkIncrementMiB *int `json:",omitempty"`
	// MemoryChunks shall contain a link to a resource collection of type
	// 'MemoryChunksCollection'.
	memoryChunks string
	// MemorySizeMiB shall contain the total size of the memory domain in mebibytes
	// (MiB).
	//
	// Version added: v1.5.0
	MemorySizeMiB *int `json:",omitempty"`
	// MinMemoryChunkSizeMiB shall contain the minimum size allowed for a memory
	// chunk within this domain in mebibytes (MiB).
	//
	// Version added: v1.5.0
	MinMemoryChunkSizeMiB *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.5.0
	Status Status
	// cXLLogicalDevices are the URIs for CXLLogicalDevices.
	cXLLogicalDevices []string
	// fabricAdapters are the URIs for FabricAdapters.
	fabricAdapters []string
	// mediaControllers are the URIs for MediaControllers.
	mediaControllers []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
}

// UnmarshalJSON unmarshals a MemoryDomain object from the raw JSON.
func (m *MemoryDomain) UnmarshalJSON(b []byte) error {
	type temp MemoryDomain
	type mLinks struct {
		CXLLogicalDevices Links `json:"CXLLogicalDevices"`
		FabricAdapters    Links `json:"FabricAdapters"`
		MediaControllers  Links `json:"MediaControllers"`
		PCIeFunctions     Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Links        mLinks
		MemoryChunks Link `json:"MemoryChunks"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryDomain(tmp.temp)

	// Extract the links to other entities for later
	m.cXLLogicalDevices = tmp.Links.CXLLogicalDevices.ToStrings()
	m.fabricAdapters = tmp.Links.FabricAdapters.ToStrings()
	m.mediaControllers = tmp.Links.MediaControllers.ToStrings()
	m.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	m.memoryChunks = tmp.MemoryChunks.String()

	return nil
}

// GetMemoryDomain will get a MemoryDomain instance from the service.
func GetMemoryDomain(c Client, uri string) (*MemoryDomain, error) {
	return GetObject[MemoryDomain](c, uri)
}

// ListReferencedMemoryDomains gets the collection of MemoryDomain from
// a provided reference.
func ListReferencedMemoryDomains(c Client, link string) ([]*MemoryDomain, error) {
	return GetCollectionObjects[MemoryDomain](c, link)
}

// CXLLogicalDevices gets the CXLLogicalDevices linked resources.
func (m *MemoryDomain) CXLLogicalDevices() ([]*CXLLogicalDevice, error) {
	return GetObjects[CXLLogicalDevice](m.client, m.cXLLogicalDevices)
}

// FabricAdapters gets the FabricAdapters linked resources.
func (m *MemoryDomain) FabricAdapters() ([]*FabricAdapter, error) {
	return GetObjects[FabricAdapter](m.client, m.fabricAdapters)
}

// MediaControllers gets the MediaControllers linked resources.
func (m *MemoryDomain) MediaControllers() ([]*MediaController, error) {
	return GetObjects[MediaController](m.client, m.mediaControllers)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (m *MemoryDomain) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](m.client, m.pCIeFunctions)
}

// MemoryChunks gets the MemoryChunks collection.
func (m *MemoryDomain) MemoryChunks() ([]*MemoryChunks, error) {
	if m.memoryChunks == "" {
		return nil, nil
	}
	return GetCollectionObjects[MemoryChunks](m.client, m.memoryChunks)
}

// MemorySet shall represent the interleave sets for a memory chunk.
type MemorySet struct {
	// MemorySet shall contain an array of links to resources of type 'Memory'.
	memorySet []string
	// MemorySetCount
	MemorySetCount int `json:"MemorySet@odata.count"`
}

// UnmarshalJSON unmarshals a MemorySet object from the raw JSON.
func (m *MemorySet) UnmarshalJSON(b []byte) error {
	type temp MemorySet
	var tmp struct {
		temp
		MemorySet Links `json:"MemorySet"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemorySet(tmp.temp)

	// Extract the links to other entities for later
	m.memorySet = tmp.MemorySet.ToStrings()

	return nil
}

// MemorySet gets the MemorySet linked resources.
func (m *MemorySet) MemorySet(client Client) ([]*Memory, error) {
	return GetObjects[Memory](client, m.memorySet)
}
