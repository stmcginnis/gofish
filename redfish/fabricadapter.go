//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FabricAdapter shall represent a physical fabric adapter capable of connecting to an interconnect fabric.
type FabricAdapter struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ASICManufacturer shall contain the manufacturer name of the ASIC for the fabric adapter as defined by the
	// manufacturer.
	ASICManufacturer string
	// ASICPartNumber shall contain the part number of the ASIC for the fabric adapter as defined by the manufacturer.
	ASICPartNumber string
	// ASICRevisionIdentifier shall contain the revision identifier of the ASIC for the fabric adapter as defined by
	// the manufacturer.
	ASICRevisionIdentifier string
	// Description provides a description of this resource.
	Description string
	// FabricType shall contain the configured fabric type of this fabric adapter.
	FabricType common.Protocol
	// FabricTypeCapabilities shall contain an array of fabric types supported by this fabric adapter.
	FabricTypeCapabilities []common.Protocol
	// FirmwareVersion shall contain the firmware version for the fabric adapter as defined by the manufacturer.
	FirmwareVersion string
	// GenZ shall contain the Gen-Z specific properties for this fabric adapter.
	GenZ FabricAdapterGenZ
	// Location shall contain the location information of the fabric adapter.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain a value that represents the manufacturer of the fabric adapter.
	Manufacturer string
	// Model shall contain the information about how the manufacturer refers to this fabric adapter.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this PCIe-based fabric adapter to its
	// host.
	PCIeInterface PCIeInterface
	// PartNumber shall contain the part number for the fabric adapter as defined by the manufacturer.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// SKU shall contain the SKU for the fabric adapter.
	SKU string
	// SerialNumber shall contain the serial number for the fabric adapter.
	SerialNumber string
	// SparePartNumber shall contain the spare part number for the fabric adapter as defined by the manufacturer.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain a universally unique identifier number for the fabric adapter.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	endpoints []string
	// EndpointsCount gets the number of logical fabric connections associated with this fabric adapter.
	EndpointsCount int
	memoryDomains  []string
	// MemoryDomainsCount gets the number of memory domains associated with this fabric adapter.
	MemoryDomainsCount int
	pcieDevices        []string
	// PCIeDevicesCount gets the number of PCIe devices associated with this fabric adapter.
	PCIeDevicesCount int
	processors       []string
	// ProcessorsCount gets the number of processors that this fabric adapter provides to a fabric.
	ProcessorsCount int
}

// UnmarshalJSON unmarshals a FabricAdapter object from the raw JSON.
func (fabricadapter *FabricAdapter) UnmarshalJSON(b []byte) error {
	type temp FabricAdapter
	type Links struct {
		// Endpoints shall contain an array of links to resources of type Endpoint that represent the logical fabric
		// connections associated with this fabric adapter.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
		// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
		// domains associated with this fabric adapter.
		MemoryDomains      common.Links
		MemoryDomainsCount int `json:"MemoryDomains@odata.count"`
		// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
		// Redfish Specification-described requirements.
		OEM json.RawMessage `json:"Oem"`
		// PCIeDevices shall contain an array of links to resources of type PCIeDevice that represent the PCIe devices
		// associated with this fabric adapter.
		PCIeDevices      common.Links
		PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
		// Processors shall contain an array of links to resources of type Processor that represent the processors that
		// this fabric adapter provides to a fabric.
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
	}
	var t struct {
		temp
		Ports common.Link
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fabricadapter = FabricAdapter(t.temp)

	// Extract the links to other entities for later
	fabricadapter.ports = t.Ports.String()

	fabricadapter.endpoints = t.Links.Endpoints.ToStrings()
	fabricadapter.EndpointsCount = t.Links.EndpointsCount
	fabricadapter.memoryDomains = t.Links.MemoryDomains.ToStrings()
	fabricadapter.MemoryDomainsCount = t.Links.MemoryDomainsCount
	fabricadapter.pcieDevices = t.Links.PCIeDevices.ToStrings()
	fabricadapter.PCIeDevicesCount = t.Links.PCIeDevicesCount
	fabricadapter.processors = t.Links.Processors.ToStrings()
	fabricadapter.ProcessorsCount = t.Links.ProcessorsCount

	// This is a read/write object, so we need to save the raw object data for later
	fabricadapter.rawData = b

	return nil
}

// Ports gets any ports associated with this interface.
func (fabricadapter *FabricAdapter) Ports() ([]*Port, error) {
	if fabricadapter.ports == "" {
		return []*Port{}, nil
	}
	return ListReferencedPorts(fabricadapter.GetClient(), fabricadapter.ports)
}

// Endpoints gets the endpoints connected to this interface.
func (fabricadapter *FabricAdapter) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](fabricadapter.GetClient(), fabricadapter.endpoints)
}

// MemoryDomains gets the MemoryDomains associated to this interface.
func (fabricadapter *FabricAdapter) MemoryDomains() ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](fabricadapter.GetClient(), fabricadapter.memoryDomains)
}

// PCIeDevices gets the PCIe devices associated to this interface.
func (fabricadapter *FabricAdapter) PCIeDevices() ([]*PCIeDevice, error) {
	return common.GetObjects[PCIeDevice](fabricadapter.GetClient(), fabricadapter.pcieDevices)
}

// Processors gets the processors associated to this interface.
func (fabricadapter *FabricAdapter) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](fabricadapter.GetClient(), fabricadapter.processors)
}

// Update commits updates to this object's properties to the running system.
func (fabricadapter *FabricAdapter) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(FabricAdapter)
	original.UnmarshalJSON(fabricadapter.rawData)

	readWriteFields := []string{
		"FabricType",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fabricadapter).Elem()

	return fabricadapter.Entity.Update(originalElement, currentElement, readWriteFields)
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

// FabricAdapterGenZ shall contain Gen-Z related properties for a fabric adapter.
type FabricAdapterGenZ struct {
	// MSDT shall contain a link to a resource collection of type RouteEntryCollection that represents the Gen-Z Core
	// Specification-defined MSDT structure.
	msdt string
	// PIDT shall contain an array of table entry values for the Gen-Z Core Specification-defined Packet Injection
	// Delay Table for the component.
	PIDT []string
	// RITable shall contain an array of table entry values for the Gen-Z Core Specification-defined Responder
	// Interface Table for the component.
	RITable []string
	// RequestorVCAT shall contain a link to a resource collection of type VCATEntryCollection that represents the
	// Gen-Z Core Specification-defined REQ-VCAT structure.
	requestorVCAT string
	// ResponderVCAT shall contain a link to a resource collection of type VCATEntryCollection that represents the
	// Gen-Z Core Specification-defined RSP-VCAT structure.
	responderVCAT string
	// SSDT shall contain a link to a resource collection of type RouteEntryCollection that represents the Gen-Z Core
	// Specification-defined SSDT structure.
	ssdt string
}

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (genz *FabricAdapterGenZ) UnmarshalJSON(b []byte) error {
	type temp FabricAdapterGenZ
	var t struct {
		temp
		MSDT          common.Link
		RequestorVCAT common.Link
		ResponderVCAT common.Link
		SSDT          common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*genz = FabricAdapterGenZ(t.temp)

	// Extract the links to other entities for later
	genz.msdt = t.MSDT.String()
	genz.requestorVCAT = t.RequestorVCAT.String()
	genz.responderVCAT = t.ResponderVCAT.String()
	genz.ssdt = t.SSDT.String()

	return nil
}
