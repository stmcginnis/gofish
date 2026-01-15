//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #MemoryDomain.v1_5_1.MemoryDomain

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// MemoryDomain shall represent memory domains in a Redfish implementation.
type MemoryDomain struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.5.0
	Status common.Status
	// cXLLogicalDevices are the URIs for CXLLogicalDevices.
	cXLLogicalDevices []string
	// fabricAdapters are the URIs for FabricAdapters.
	fabricAdapters []string
	// mediaControllers are the URIs for MediaControllers.
	mediaControllers []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MemoryDomain object from the raw JSON.
func (m *MemoryDomain) UnmarshalJSON(b []byte) error {
	type temp MemoryDomain
	type mLinks struct {
		CXLLogicalDevices common.Links `json:"CXLLogicalDevices"`
		FabricAdapters    common.Links `json:"FabricAdapters"`
		MediaControllers  common.Links `json:"MediaControllers"`
		PCIeFunctions     common.Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Links        mLinks
		MemoryChunks common.Link `json:"memoryChunks"`
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

	// This is a read/write object, so we need to save the raw object data for later
	m.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MemoryDomain) Update() error {
	readWriteFields := []string{
		"InterleavableMemorySets",
		"Status",
	}

	return m.UpdateFromRawData(m, m.rawData, readWriteFields)
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

// CXLLogicalDevices gets the CXLLogicalDevices linked resources.
func (m *MemoryDomain) CXLLogicalDevices(client common.Client) ([]*CXLLogicalDevice, error) {
	return common.GetObjects[CXLLogicalDevice](client, m.cXLLogicalDevices)
}

// FabricAdapters gets the FabricAdapters linked resources.
func (m *MemoryDomain) FabricAdapters(client common.Client) ([]*FabricAdapter, error) {
	return common.GetObjects[FabricAdapter](client, m.fabricAdapters)
}

// MediaControllers gets the MediaControllers linked resources.
func (m *MemoryDomain) MediaControllers(client common.Client) ([]*MediaController, error) {
	return common.GetObjects[MediaController](client, m.mediaControllers)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (m *MemoryDomain) PCIeFunctions(client common.Client) ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](client, m.pCIeFunctions)
}

// MemoryChunks gets the MemoryChunks collection.
func (m *MemoryDomain) MemoryChunks(client common.Client) ([]*MemoryChunks, error) {
	if m.memoryChunks == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[MemoryChunks](client, m.memoryChunks)
}

// MemorySet shall represent the interleave sets for a memory chunk.
type MemorySet struct {
	// MemorySet shall contain an array of links to resources of type 'Memory'.
	MemorySet []Memory
	// MemorySet@odata.count
	MemorySetCount int `json:"MemorySet@odata.count"`
}
