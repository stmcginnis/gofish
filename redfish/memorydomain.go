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
	// memoryChunks shall be a link to a collection of type MemoryChunkCollection.
	memoryChunks string
}

// UnmarshalJSON unmarshals a MemoryDomain object from the raw JSON.
func (memorydomain *MemoryDomain) UnmarshalJSON(b []byte) error {
	type temp MemoryDomain
	var t struct {
		temp
		MemoryChunks common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*memorydomain = MemoryDomain(t.temp)
	memorydomain.memoryChunks = string(t.MemoryChunks)

	return nil
}

// GetMemoryDomain will get a MemoryDomain instance from the service.
func GetMemoryDomain(c common.Client, uri string) (*MemoryDomain, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memorydomain MemoryDomain
	err = json.NewDecoder(resp.Body).Decode(&memorydomain)
	if err != nil {
		return nil, err
	}

	memorydomain.SetClient(c)
	return &memorydomain, nil
}

// ListReferencedMemoryDomains gets the collection of MemoryDomain from
// a provided reference.
func ListReferencedMemoryDomains(c common.Client, link string) ([]*MemoryDomain, error) {
	var result []*MemoryDomain
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, memorydomainLink := range links.ItemLinks {
		memorydomain, err := GetMemoryDomain(c, memorydomainLink)
		if err != nil {
			return result, err
		}
		result = append(result, memorydomain)
	}

	return result, nil
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
