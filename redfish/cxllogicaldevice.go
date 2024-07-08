//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type CXLSemantic string

const (
	// CXLioCXLSemantic shall indicate the device conforms with the CXL Specification-defined 'CXL.io' semantic.
	CXLioCXLSemantic CXLSemantic = "CXLio"
	// CXLcacheCXLSemantic shall indicate the device conforms with the CXL Specification-defined 'CXL.cache' semantic.
	CXLcacheCXLSemantic CXLSemantic = "CXLcache"
	// CXLmemCXLSemantic shall indicate the device conforms with the CXL Specification-defined 'CXL.mem' semantic.
	CXLmemCXLSemantic CXLSemantic = "CXLmem"
)

// CXLLogicalDevice shall represent a CXL logical device that is a part of a PCIe device.
type CXLLogicalDevice struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifiers shall contain a list of all known durable names for the associated CXL logical device.
	Identifiers []common.Identifier
	// Log shall contain a link to a resource of type LogService.
	log string
	// MemoryRegions shall contain a link to a resource collection of type MemoryRegionCollection that represents the
	// memory regions associated with this CXL logical device.
	memoryRegions string
	// MemorySizeMiB shall contain the memory region size defined in this CXL logical device in mebibytes (MiB).
	MemorySizeMiB int
	// QoS shall contain the quality of service configuration for this CXL logical device.
	QoS QoS
	// QoSTelemetryCapabilities shall contain the quality of service telemetry capabilities for this CXL logical
	// device.
	QoSTelemetryCapabilities QoSTelemetryCapabilities
	// SemanticsSupported shall contain the CXL Specification-defined semantics that are supported by this CXL logical
	// device.
	SemanticsSupported []CXLSemantic
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	endpoints []string
	// EndpointsCount is the number of endpoints associated with this CXL logical device.
	EndpointsCount int
	memoryChunks   []string
	// MemoryChunksCount is the number of memory chunks owned by this CXL logical device.
	MemoryChunksCount int
	memoryDomains     []string
	// MemoryDomainsCount is the number of memory domains associated with this CXL logical device.
	MemoryDomainsCount int
	pcieFunctions      []string
	// PCIeFunctionsCount is the number of PCIe functions assigned to this CXL logical device.
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a CXLLogicalDevice object from the raw JSON.
func (cxllogicaldevice *CXLLogicalDevice) UnmarshalJSON(b []byte) error {
	type temp CXLLogicalDevice
	type Links struct {
		// Endpoints shall contain an array of links to resources of type Endpoint that represent the endpoints associated
		// with this CXL logical device.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
		// MemoryChunks shall contain an array of links to resources of type MemoryChunks that represent the memory chunks
		// owned by this CXL logical device.
		MemoryChunks      common.Links
		MemoryChunksCount int `json:"MemoryChunks@odata.count"`
		// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
		// domains associated with this CXL logical device.
		MemoryDomains      common.Links
		MemoryDomainsCount int `json:"MemoryDomains@odata.count"`
		// PCIeFunctions shall contain an array of links to resources of type PCIeFunction that represent the PCIe
		// functions assigned to this CXL logical device.
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	}
	var t struct {
		temp
		Log           common.Link
		MemoryRegions common.Link
		Links         Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cxllogicaldevice = CXLLogicalDevice(t.temp)

	// Extract the links to other entities for later
	cxllogicaldevice.log = t.Log.String()
	cxllogicaldevice.memoryRegions = t.MemoryRegions.String()

	cxllogicaldevice.endpoints = t.Links.Endpoints.ToStrings()
	cxllogicaldevice.EndpointsCount = t.Links.EndpointsCount
	cxllogicaldevice.memoryChunks = t.Links.MemoryChunks.ToStrings()
	cxllogicaldevice.MemoryChunksCount = t.Links.MemoryChunksCount
	cxllogicaldevice.memoryDomains = t.Links.MemoryDomains.ToStrings()
	cxllogicaldevice.MemoryDomainsCount = t.Links.MemoryDomainsCount
	cxllogicaldevice.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	cxllogicaldevice.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	return nil
}

// Log gets the LogService for this device.
func (cxllogicaldevice *CXLLogicalDevice) Log() (*LogService, error) {
	if cxllogicaldevice.log == "" {
		return nil, nil
	}

	return GetLogService(cxllogicaldevice.GetClient(), cxllogicaldevice.log)
}

// Endpoints get the endpoints associated with this CXL logical device.
func (cxllogicaldevice *CXLLogicalDevice) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](cxllogicaldevice.GetClient(), cxllogicaldevice.endpoints)
}

// MemoryChunks get the memory chunks associated with this CXL logical device.
func (cxllogicaldevice *CXLLogicalDevice) MemoryChunks() ([]*MemoryChunks, error) {
	return common.GetObjects[MemoryChunks](cxllogicaldevice.GetClient(), cxllogicaldevice.memoryChunks)
}

// MemoryDomains get the memory domains associated with this CXL logical device.
func (cxllogicaldevice *CXLLogicalDevice) MemoryDomains() ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](cxllogicaldevice.GetClient(), cxllogicaldevice.memoryDomains)
}

// PCIeFunctions get the PCIe functions associated with this CXL logical device.
func (cxllogicaldevice *CXLLogicalDevice) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](cxllogicaldevice.GetClient(), cxllogicaldevice.pcieFunctions)
}

// GetCXLLogicalDevice will get a CXLLogicalDevice instance from the service.
func GetCXLLogicalDevice(c common.Client, uri string) (*CXLLogicalDevice, error) {
	return common.GetObject[CXLLogicalDevice](c, uri)
}

// ListReferencedCXLLogicalDevices gets the collection of CXLLogicalDevice from
// a provided reference.
func ListReferencedCXLLogicalDevices(c common.Client, link string) ([]*CXLLogicalDevice, error) {
	return common.GetCollectionObjects[CXLLogicalDevice](c, link)
}

// QoS shall contain the quality of service properties of this CXL logical device.
type QoS struct {
	// AllocatedBandwidth shall contain the bandwidth allocated for this CXL logical device in multiples of 256.
	AllocatedBandwidth int
	// LimitPercent shall contain the bandwidth limit, '0' to '100', for this CXL logical device as a percentage.
	LimitPercent int
}
