//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #PCIeFunction.v1_7_0.PCIeFunction

package schemas

import (
	"encoding/json"
)

type DeviceClass string

const (
	// UnclassifiedDeviceDeviceClass is an unclassified device.
	UnclassifiedDeviceDeviceClass DeviceClass = "UnclassifiedDevice"
	// MassStorageControllerDeviceClass is a mass storage controller.
	MassStorageControllerDeviceClass DeviceClass = "MassStorageController"
	// NetworkControllerDeviceClass is a network controller.
	NetworkControllerDeviceClass DeviceClass = "NetworkController"
	// DisplayControllerDeviceClass is a display controller.
	DisplayControllerDeviceClass DeviceClass = "DisplayController"
	// MultimediaControllerDeviceClass is a multimedia controller.
	MultimediaControllerDeviceClass DeviceClass = "MultimediaController"
	// MemoryControllerDeviceClass is a memory controller.
	MemoryControllerDeviceClass DeviceClass = "MemoryController"
	// BridgeDeviceClass is a bridge.
	BridgeDeviceClass DeviceClass = "Bridge"
	// CommunicationControllerDeviceClass is a communication controller.
	CommunicationControllerDeviceClass DeviceClass = "CommunicationController"
	// GenericSystemPeripheralDeviceClass is a generic system peripheral.
	GenericSystemPeripheralDeviceClass DeviceClass = "GenericSystemPeripheral"
	// InputDeviceControllerDeviceClass is an input device controller.
	InputDeviceControllerDeviceClass DeviceClass = "InputDeviceController"
	// DockingStationDeviceClass is a docking station.
	DockingStationDeviceClass DeviceClass = "DockingStation"
	// ProcessorDeviceClass is a processor.
	ProcessorDeviceClass DeviceClass = "Processor"
	// SerialBusControllerDeviceClass is a serial bus controller.
	SerialBusControllerDeviceClass DeviceClass = "SerialBusController"
	// WirelessControllerDeviceClass is a wireless controller.
	WirelessControllerDeviceClass DeviceClass = "WirelessController"
	// IntelligentControllerDeviceClass is an intelligent controller.
	IntelligentControllerDeviceClass DeviceClass = "IntelligentController"
	// SatelliteCommunicationsControllerDeviceClass is a satellite communications
	// controller.
	SatelliteCommunicationsControllerDeviceClass DeviceClass = "SatelliteCommunicationsController"
	// EncryptionControllerDeviceClass is an encryption controller.
	EncryptionControllerDeviceClass DeviceClass = "EncryptionController"
	// SignalProcessingControllerDeviceClass is a signal processing controller.
	SignalProcessingControllerDeviceClass DeviceClass = "SignalProcessingController"
	// ProcessingAcceleratorsDeviceClass is a processing accelerators.
	ProcessingAcceleratorsDeviceClass DeviceClass = "ProcessingAccelerators"
	// NonEssentialInstrumentationDeviceClass is a non-essential instrumentation.
	NonEssentialInstrumentationDeviceClass DeviceClass = "NonEssentialInstrumentation"
	// CoprocessorDeviceClass is a coprocessor.
	CoprocessorDeviceClass DeviceClass = "Coprocessor"
	// UnassignedClassDeviceClass is an unassigned class.
	UnassignedClassDeviceClass DeviceClass = "UnassignedClass"
	// OtherDeviceClass Other class. The function Class Code needs to be verified.
	OtherDeviceClass DeviceClass = "Other"
)

type FunctionProtocol string

const (
	// PCIeFunctionProtocol is a standard PCIe function.
	PCIeFunctionProtocol FunctionProtocol = "PCIe"
	// CXLFunctionProtocol is a PCIe function supporting CXL extensions.
	CXLFunctionProtocol FunctionProtocol = "CXL"
)

type FunctionType string

const (
	// PhysicalFunctionType is a physical PCIe function.
	PhysicalFunctionType FunctionType = "Physical"
	// VirtualFunctionType is a virtual PCIe function.
	VirtualFunctionType FunctionType = "Virtual"
)

// PCIeFunction shall represent a PCIe function in a Redfish implementation.
type PCIeFunction struct {
	Entity
	// ARIFunctionNumber shall contain the PCIe ARI (Alternative Routing-ID
	// Interpretation) function number of the PCIe device function. This property
	// shall not be present if 'ARIModeEnabled' contains 'false' or is not
	// supported. This property shall not be present if the PCIe device function is
	// fabric-attached or is shared with multiple systems.
	//
	// Version added: v1.7.0
	ARIFunctionNumber string
	// ARIModeEnabled shall indicate whether PCIe ARI (Alternative Routing-ID
	// Interpretation) is enabled for this PCIe function. This property shall not
	// be present if the PCIe device function is fabric-attached or is shared with
	// multiple systems.
	//
	// Version added: v1.7.0
	ARIModeEnabled bool
	// BusNumber shall contain the PCIe bus number of the PCIe device function.
	// This property shall not be present if the PCIe device function is
	// fabric-attached or is shared with multiple systems.
	//
	// Version added: v1.6.0
	BusNumber string
	// ClassCode shall contain the PCI Class Code, Subclass, and Programming
	// Interface of the PCIe device function in the order listed.
	ClassCode string
	// DeviceClass shall contain the device class of the PCIe device function, such
	// as storage, network, or memory.
	DeviceClass DeviceClass
	// DeviceID shall contain the PCI Device ID of the PCIe device function with
	// the most significant byte shown first.
	DeviceID string `json:"DeviceId"`
	// DeviceNumber shall contain the PCIe device number of the PCIe device
	// function. If 'ARIModeEnabled' contains 'true', this property shall contain
	// the value of the 'ARIFunctionNumber' property shifted right three bits. This
	// property shall not be present if the PCIe device function is fabric-attached
	// or is shared with multiple systems.
	//
	// Version added: v1.6.0
	DeviceNumber string
	// Enabled shall indicate if this PCIe device function is enabled.
	//
	// Version added: v1.3.0
	Enabled bool
	// FunctionID shall contain the PCIe function number within a given PCIe
	// device.
	FunctionID *int `json:"FunctionId,omitempty"`
	// FunctionNumber shall contain the PCIe function number of the PCIe device
	// function. If 'ARIModeEnabled' contains 'true', this property shall contain
	// the lower three bits of the value of the 'ARIFunctionNumber' property. This
	// property shall not be present if the PCIe device function is fabric-attached
	// or is shared with multiple systems.
	//
	// Version added: v1.6.0
	FunctionNumber string
	// FunctionProtocol shall contain the protocol supported by this PCIe function.
	//
	// Version added: v1.5.0
	FunctionProtocol FunctionProtocol
	// FunctionType shall contain the function type of the PCIe device function
	// such as physical or virtual.
	FunctionType FunctionType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RevisionID shall contain the PCI Revision ID of the PCIe device function.
	RevisionID string `json:"RevisionId"`
	// SegmentNumber shall contain the PCIe segment number of the PCIe device
	// function. This property shall not be present if the PCIe device function is
	// fabric-attached or is shared with multiple systems.
	//
	// Version added: v1.6.0
	SegmentNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SubsystemID shall contain the PCI Subsystem ID of the PCIe device function
	// with the most significant byte shown first.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall contain the PCI Subsystem Vendor ID of the PCIe
	// device function with the most significant byte shown first.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall contain the PCI Vendor ID of the PCIe device function with
	// the most significant byte shown first.
	VendorID string `json:"VendorId"`
	// cXLLogicalDevice is the URI for CXLLogicalDevice.
	cXLLogicalDevice string
	// drives are the URIs for Drives.
	drives []string
	// ethernetInterfaces are the URIs for EthernetInterfaces.
	ethernetInterfaces []string
	// memoryDomains are the URIs for MemoryDomains.
	memoryDomains []string
	// networkDeviceFunctions are the URIs for NetworkDeviceFunctions.
	networkDeviceFunctions []string
	// pCIeDevice is the URI for PCIeDevice.
	pCIeDevice string
	// processor is the URI for Processor.
	processor string
	// storageControllers are the URIs for StorageControllers.
	storageControllers []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PCIeFunction object from the raw JSON.
func (p *PCIeFunction) UnmarshalJSON(b []byte) error {
	type temp PCIeFunction
	type pLinks struct {
		CXLLogicalDevice       Link  `json:"CXLLogicalDevice"`
		Drives                 Links `json:"Drives"`
		EthernetInterfaces     Links `json:"EthernetInterfaces"`
		MemoryDomains          Links `json:"MemoryDomains"`
		NetworkDeviceFunctions Links `json:"NetworkDeviceFunctions"`
		PCIeDevice             Link  `json:"PCIeDevice"`
		Processor              Link  `json:"Processor"`
		StorageControllers     Links `json:"StorageControllers"`
	}
	var tmp struct {
		temp
		Links pLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PCIeFunction(tmp.temp)

	// Extract the links to other entities for later
	p.cXLLogicalDevice = tmp.Links.CXLLogicalDevice.String()
	p.drives = tmp.Links.Drives.ToStrings()
	p.ethernetInterfaces = tmp.Links.EthernetInterfaces.ToStrings()
	p.memoryDomains = tmp.Links.MemoryDomains.ToStrings()
	p.networkDeviceFunctions = tmp.Links.NetworkDeviceFunctions.ToStrings()
	p.pCIeDevice = tmp.Links.PCIeDevice.String()
	p.processor = tmp.Links.Processor.String()
	p.storageControllers = tmp.Links.StorageControllers.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PCIeFunction) Update() error {
	readWriteFields := []string{
		"Enabled",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPCIeFunction will get a PCIeFunction instance from the service.
func GetPCIeFunction(c Client, uri string) (*PCIeFunction, error) {
	return GetObject[PCIeFunction](c, uri)
}

// ListReferencedPCIeFunctions gets the collection of PCIeFunction from
// a provided reference.
func ListReferencedPCIeFunctions(c Client, link string) ([]*PCIeFunction, error) {
	return GetCollectionObjects[PCIeFunction](c, link)
}

// CXLLogicalDevice gets the CXLLogicalDevice linked resource.
func (p *PCIeFunction) CXLLogicalDevice() (*CXLLogicalDevice, error) {
	if p.cXLLogicalDevice == "" {
		return nil, nil
	}
	return GetObject[CXLLogicalDevice](p.client, p.cXLLogicalDevice)
}

// Drives gets the Drives linked resources.
func (p *PCIeFunction) Drives() ([]*Drive, error) {
	return GetObjects[Drive](p.client, p.drives)
}

// EthernetInterfaces gets the EthernetInterfaces linked resources.
func (p *PCIeFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	return GetObjects[EthernetInterface](p.client, p.ethernetInterfaces)
}

// MemoryDomains gets the MemoryDomains linked resources.
func (p *PCIeFunction) MemoryDomains() ([]*MemoryDomain, error) {
	return GetObjects[MemoryDomain](p.client, p.memoryDomains)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions linked resources.
func (p *PCIeFunction) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](p.client, p.networkDeviceFunctions)
}

// PCIeDevice gets the PCIeDevice linked resource.
func (p *PCIeFunction) PCIeDevice() (*PCIeDevice, error) {
	if p.pCIeDevice == "" {
		return nil, nil
	}
	return GetObject[PCIeDevice](p.client, p.pCIeDevice)
}

// Processor gets the Processor linked resource.
func (p *PCIeFunction) Processor() (*Processor, error) {
	if p.processor == "" {
		return nil, nil
	}
	return GetObject[Processor](p.client, p.processor)
}

// StorageControllers gets the StorageControllers linked resources.
func (p *PCIeFunction) StorageControllers() ([]*StorageController, error) {
	return GetObjects[StorageController](p.client, p.storageControllers)
}
