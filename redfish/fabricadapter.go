//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #FabricAdapter.v1_5_3.FabricAdapter

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// FabricAdapter shall represent a physical fabric adapter capable of connecting
// to an interconnect fabric.
type FabricAdapter struct {
	common.Entity
	// ASICManufacturer shall contain the manufacturer name of the ASIC for the
	// fabric adapter as defined by the manufacturer.
	ASICManufacturer string
	// ASICPartNumber shall contain the part number of the ASIC for the fabric
	// adapter as defined by the manufacturer.
	ASICPartNumber string
	// ASICRevisionIdentifier shall contain the revision identifier of the ASIC for
	// the fabric adapter as defined by the manufacturer.
	ASICRevisionIdentifier string
	// FabricType shall contain the configured fabric type of this fabric adapter.
	//
	// Version added: v1.3.0
	FabricType common.Protocol
	// FabricTypeCapabilities shall contain an array of fabric types supported by
	// this fabric adapter.
	//
	// Version added: v1.3.0
	FabricTypeCapabilities []common.Protocol
	// FirmwareVersion shall contain the firmware version for the fabric adapter as
	// defined by the manufacturer.
	FirmwareVersion string
	// GenZ shall contain the Gen-Z specific properties for this fabric adapter.
	GenZ FabricAdapterGenZ
	// Location shall contain the location information of the fabric adapter.
	//
	// Version added: v1.1.0
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	//
	// Version added: v1.4.0
	LocationIndicatorActive bool
	// Manufacturer shall contain a value that represents the manufacturer of the
	// fabric adapter.
	Manufacturer string
	// Model shall contain the information about how the manufacturer refers to
	// this fabric adapter.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this
	// PCIe-based fabric adapter to its host.
	PCIeInterface PCIeInterface
	// PartNumber shall contain the part number for the fabric adapter as defined
	// by the manufacturer.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// SKU shall contain the SKU for the fabric adapter.
	SKU string
	// SerialNumber shall contain the serial number for the fabric adapter.
	SerialNumber string
	// SparePartNumber shall contain the spare part number for the fabric adapter
	// as defined by the manufacturer.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain a universally unique identifier number for the fabric
	// adapter.
	UUID string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// memoryDomains are the URIs for MemoryDomains.
	memoryDomains []string
	// pCIeDevices are the URIs for PCIeDevices.
	pCIeDevices []string
	// processors are the URIs for Processors.
	processors []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a FabricAdapter object from the raw JSON.
func (f *FabricAdapter) UnmarshalJSON(b []byte) error {
	type temp FabricAdapter
	type fLinks struct {
		Endpoints     common.Links `json:"Endpoints"`
		MemoryDomains common.Links `json:"MemoryDomains"`
		PCIeDevices   common.Links `json:"PCIeDevices"`
		Processors    common.Links `json:"Processors"`
	}
	var tmp struct {
		temp
		Links fLinks
		Ports common.Link `json:"ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FabricAdapter(tmp.temp)

	// Extract the links to other entities for later
	f.endpoints = tmp.Links.Endpoints.ToStrings()
	f.memoryDomains = tmp.Links.MemoryDomains.ToStrings()
	f.pCIeDevices = tmp.Links.PCIeDevices.ToStrings()
	f.processors = tmp.Links.Processors.ToStrings()
	f.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *FabricAdapter) Update() error {
	readWriteFields := []string{
		"FabricType",
		"GenZ",
		"Location",
		"LocationIndicatorActive",
		"PCIeInterface",
		"Status",
	}

	return f.UpdateFromRawData(f, f.rawData, readWriteFields)
}

// GetFabricAdapter will get a FabricAdapter instance from the service.
func GetFabricAdapter(c common.Client, uri string) (*FabricAdapter, error) {
	return common.GetObject[FabricAdapter](c, uri)
}

// ListReferencedFabricAdapters gets the collection of FabricAdapter from
// a provided reference.
func ListReferencedFabricAdapters(c common.Client, link string) ([]*FabricAdapter, error) {
	return common.GetCollectionObjects[FabricAdapter](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (f *FabricAdapter) Endpoints(client common.Client) ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](client, f.endpoints)
}

// MemoryDomains gets the MemoryDomains linked resources.
func (f *FabricAdapter) MemoryDomains(client common.Client) ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](client, f.memoryDomains)
}

// PCIeDevices gets the PCIeDevices linked resources.
func (f *FabricAdapter) PCIeDevices(client common.Client) ([]*PCIeDevice, error) {
	return common.GetObjects[PCIeDevice](client, f.pCIeDevices)
}

// Processors gets the Processors linked resources.
func (f *FabricAdapter) Processors(client common.Client) ([]*Processor, error) {
	return common.GetObjects[Processor](client, f.processors)
}

// Ports gets the Ports collection.
func (f *FabricAdapter) Ports(client common.Client) ([]*Port, error) {
	if f.ports == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, f.ports)
}

// FabricAdapterGenZ shall contain Gen-Z related properties for a fabric adapter.
type FabricAdapterGenZ struct {
	// MSDT shall contain a link to a resource collection of type
	// 'RouteEntryCollection' that represents the Gen-Z Core Specification-defined
	// MSDT structure.
	mSDT string
	// PIDT shall contain an array of table entry values for the Gen-Z Core
	// Specification-defined Packet Injection Delay Table for the component.
	PIDT []string
	// RITable shall contain an array of table entry values for the Gen-Z Core
	// Specification-defined Responder Interface Table for the component.
	RITable []string
	// RequestorVCAT shall contain a link to a resource collection of type
	// 'VCATEntryCollection' that represents the Gen-Z Core Specification-defined
	// REQ-VCAT structure.
	requestorVCAT string
	// ResponderVCAT shall contain a link to a resource collection of type
	// 'VCATEntryCollection' that represents the Gen-Z Core Specification-defined
	// RSP-VCAT structure.
	responderVCAT string
	// SSDT shall contain a link to a resource collection of type
	// 'RouteEntryCollection' that represents the Gen-Z Core Specification-defined
	// SSDT structure.
	sSDT string
}

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (g *FabricAdapterGenZ) UnmarshalJSON(b []byte) error {
	type temp FabricAdapterGenZ
	var tmp struct {
		temp
		MSDT          common.Link `json:"mSDT"`
		RequestorVCAT common.Link `json:"requestorVCAT"`
		ResponderVCAT common.Link `json:"responderVCAT"`
		SSDT          common.Link `json:"sSDT"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = FabricAdapterGenZ(tmp.temp)

	// Extract the links to other entities for later
	g.mSDT = tmp.MSDT.String()
	g.requestorVCAT = tmp.RequestorVCAT.String()
	g.responderVCAT = tmp.ResponderVCAT.String()
	g.sSDT = tmp.SSDT.String()

	return nil
}

// MSDT gets the MSDT collection.
func (g *FabricAdapterGenZ) MSDT(client common.Client) ([]*RouteEntry, error) {
	if g.mSDT == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[RouteEntry](client, g.mSDT)
}

// RequestorVCAT gets the RequestorVCAT collection.
func (g *FabricAdapterGenZ) RequestorVCAT(client common.Client) ([]*VCATEntry, error) {
	if g.requestorVCAT == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[VCATEntry](client, g.requestorVCAT)
}

// ResponderVCAT gets the ResponderVCAT collection.
func (g *FabricAdapterGenZ) ResponderVCAT(client common.Client) ([]*VCATEntry, error) {
	if g.responderVCAT == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[VCATEntry](client, g.responderVCAT)
}

// SSDT gets the SSDT collection.
func (g *FabricAdapterGenZ) SSDT(client common.Client) ([]*RouteEntry, error) {
	if g.sSDT == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[RouteEntry](client, g.sSDT)
}
