//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// DeviceClass is the device class.
type DeviceClass string

const (
	// UnclassifiedDeviceDeviceClass An unclassified device.
	UnclassifiedDeviceDeviceClass DeviceClass = "UnclassifiedDevice"
	// MassStorageControllerDeviceClass A mass storage controller.
	MassStorageControllerDeviceClass DeviceClass = "MassStorageController"
	// NetworkControllerDeviceClass A network controller.
	NetworkControllerDeviceClass DeviceClass = "NetworkController"
	// DisplayControllerDeviceClass A display controller.
	DisplayControllerDeviceClass DeviceClass = "DisplayController"
	// MultimediaControllerDeviceClass A multimedia controller.
	MultimediaControllerDeviceClass DeviceClass = "MultimediaController"
	// MemoryControllerDeviceClass A memory controller.
	MemoryControllerDeviceClass DeviceClass = "MemoryController"
	// BridgeDeviceClass A bridge.
	BridgeDeviceClass DeviceClass = "Bridge"
	// CommunicationControllerDeviceClass A communication controller.
	CommunicationControllerDeviceClass DeviceClass = "CommunicationController"
	// GenericSystemPeripheralDeviceClass A generic system peripheral.
	GenericSystemPeripheralDeviceClass DeviceClass = "GenericSystemPeripheral"
	// InputDeviceControllerDeviceClass An input device controller.
	InputDeviceControllerDeviceClass DeviceClass = "InputDeviceController"
	// DockingStationDeviceClass A docking station.
	DockingStationDeviceClass DeviceClass = "DockingStation"
	// ProcessorDeviceClass A processor.
	ProcessorDeviceClass DeviceClass = "Processor"
	// SerialBusControllerDeviceClass A serial bus controller.
	SerialBusControllerDeviceClass DeviceClass = "SerialBusController"
	// WirelessControllerDeviceClass A wireless controller.
	WirelessControllerDeviceClass DeviceClass = "WirelessController"
	// IntelligentControllerDeviceClass An intelligent controller.
	IntelligentControllerDeviceClass DeviceClass = "IntelligentController"
	// SatelliteCommunicationsControllerDeviceClass A satellite
	// communications controller.
	SatelliteCommunicationsControllerDeviceClass DeviceClass = "SatelliteCommunicationsController"
	// EncryptionControllerDeviceClass An encryption controller.
	EncryptionControllerDeviceClass DeviceClass = "EncryptionController"
	// SignalProcessingControllerDeviceClass A signal processing controller.
	SignalProcessingControllerDeviceClass DeviceClass = "SignalProcessingController"
	// ProcessingAcceleratorsDeviceClass A processing accelerators.
	ProcessingAcceleratorsDeviceClass DeviceClass = "ProcessingAccelerators"
	// NonEssentialInstrumentationDeviceClass A non-essential
	// instrumentation.
	NonEssentialInstrumentationDeviceClass DeviceClass = "NonEssentialInstrumentation"
	// CoprocessorDeviceClass A coprocessor.
	CoprocessorDeviceClass DeviceClass = "Coprocessor"
	// UnassignedClassDeviceClass An unassigned class.
	UnassignedClassDeviceClass DeviceClass = "UnassignedClass"
	// OtherDeviceClass A other class. The function Device Class Id needs to
	// be verified.
	OtherDeviceClass DeviceClass = "Other"
)

type FunctionProtocol string

const (
	// PCIeFunctionProtocol A standard PCIe function.
	PCIeFunctionProtocol FunctionProtocol = "PCIe"
	// CXLFunctionProtocol A PCIe function supporting CXL extensions.
	CXLFunctionProtocol FunctionProtocol = "CXL"
)

// FunctionType is the function type.
type FunctionType string

const (
	// PhysicalFunctionType A physical PCie function.
	PhysicalFunctionType FunctionType = "Physical"
	// VirtualFunctionType A virtual PCIe function.
	VirtualFunctionType FunctionType = "Virtual"
)

// PCIeFunction is used to represent a PCIeFunction attached to a System.
type PCIeFunction struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClassCode shall be the PCI Class Code of the PCIe device function.
	ClassCode string
	// Description provides a description of this resource.
	Description string
	// DeviceClass shall be the device class of the PCIe device function such as
	// Storage, Network, Memory etc.
	DeviceClass DeviceClass
	// DeviceID shall be the PCI Device ID of the PCIe device function.
	DeviceID string
	// FunctionID shall the PCIe device function number within a given PCIe
	// device.
	FunctionID int
	// FunctionProtocol shall contain the protocol supported by this PCIe function.
	FunctionProtocol FunctionProtocol
	// FunctionType shall be the function type of the PCIe device function such
	// as Physical or Virtual.
	FunctionType FunctionType
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// RevisionID shall be the PCI Revision ID of the PCIe device function.
	RevisionID string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SubsystemID shall be the PCI Subsystem ID of the PCIe device function.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall be the PCI Subsystem Vendor ID of the PCIe device
	// function.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall be the PCI Vendor ID of the PCIe device function.
	VendorID string `json:"VendorId"`

	cxlLogicalDevice string
	// Drives shall reference a resource of type Drive that represents the
	// storage drives associated with this resource.
	drives []string
	// DrivesCount is the number of drives.
	DrivesCount int
	// EthernetInterfaces shall reference a resource of type EthernetInterface
	// that represents the network interfaces associated with this resource.
	ethernetInterfaces []string
	// EthernetInterfacesCount is the number of ethernet interfaces.
	EthernetInterfacesCount int
	// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
	// domains associated with this PCIe function.
	memoryDomains []string
	// MemoryDomainsCount is the number of memory domains associated with this PCIe function.
	MemoryDomainsCount int
	// NetworkDeviceFunctions shall be an array of references to resources of
	// type NetworkDeviceFunction that represents the network device functions
	// associated with this resource.
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount is the number of network device functions.
	NetworkDeviceFunctionsCount int
	// PCIeDevice shall be a reference to the resource that this function is a
	// part of and shall reference a resource of type PCIeDevice.
	pcieDevice string
	// Processor shall link to a resource of type Processor that represents the processor that is hosted on this PCIe
	// function.
	processor string
	// StorageControllers shall reference a resource of type StorageController
	// that represents the storage controllers associated with this resource.
	storageControllers []string
	// StorageControllersCount is the number of storage controllers.
	StorageControllersCount int
}

// UnmarshalJSON unmarshals a PCIeFunction object from the raw JSON.
func (pciefunction *PCIeFunction) UnmarshalJSON(b []byte) error {
	type temp PCIeFunction

	type links struct {
		CXLLogicalDevice            common.Link
		Drives                      common.Links
		DrivesCount                 int `json:"Drives@odata.count"`
		EthernetInterfaces          common.Links
		EthernetInterfacesCount     int `json:"EthernetInterfaces@odata.count"`
		MemoryDomains               common.Links
		MemoryDomainsCount          int `json:"MemoryDomains@odata.count"`
		NetworkDeviceFunctions      common.Links
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		PCIeDevice                  common.Link
		Processor                   common.Link
		StorageControllers          common.Links
		StorageControllersCount     int `json:"StorageControllers@odata.count"`
	}

	var t struct {
		temp
		Links links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciefunction = PCIeFunction(t.temp)

	// Extract the links to other entities for later
	pciefunction.cxlLogicalDevice = t.Links.CXLLogicalDevice.String()
	pciefunction.drives = t.Links.Drives.ToStrings()
	pciefunction.DrivesCount = t.Links.DrivesCount
	pciefunction.ethernetInterfaces = t.Links.EthernetInterfaces.ToStrings()
	pciefunction.EthernetInterfacesCount = t.Links.EthernetInterfacesCount
	pciefunction.memoryDomains = t.Links.MemoryDomains.ToStrings()
	pciefunction.MemoryDomainsCount = t.Links.MemoryDomainsCount
	pciefunction.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	pciefunction.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	pciefunction.pcieDevice = t.Links.PCIeDevice.String()
	pciefunction.processor = t.Links.Processor.String()
	pciefunction.storageControllers = t.Links.StorageControllers.ToStrings()
	pciefunction.StorageControllersCount = t.Links.StorageControllersCount

	return nil
}

// GetPCIeFunction will get a PCIeFunction instance from the service.
func GetPCIeFunction(c common.Client, uri string) (*PCIeFunction, error) {
	return common.GetObject[PCIeFunction](c, uri)
}

// ListReferencedPCIeFunctions gets the collection of PCIeFunction from
// a provided reference.
func ListReferencedPCIeFunctions(c common.Client, link string) ([]*PCIeFunction, error) {
	return common.GetCollectionObjects[PCIeFunction](c, link)
}

// CXLLogicalDevice gets the CXL logical device to which this PCIe function is assigned.
func (pciefunction *PCIeFunction) CXLLogicalDevice() (*CXLLogicalDevice, error) {
	if pciefunction.cxlLogicalDevice == "" {
		return nil, nil
	}
	return GetCXLLogicalDevice(pciefunction.GetClient(), pciefunction.cxlLogicalDevice)
}

// Drives gets the PCIe function's drives.
func (pciefunction *PCIeFunction) Drives() ([]*Drive, error) {
	return common.GetObjects[Drive](pciefunction.GetClient(), pciefunction.drives)
}

// EthernetInterfaces gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	return common.GetObjects[EthernetInterface](pciefunction.GetClient(), pciefunction.ethernetInterfaces)
}

// MemoryDomains gets the memory domains associated with this PCIe function.
func (pciefunction *PCIeFunction) MemoryDomains() ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](pciefunction.GetClient(), pciefunction.memoryDomains)
}

// NetworkDeviceFunctions gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return common.GetObjects[NetworkDeviceFunction](pciefunction.GetClient(), pciefunction.networkDeviceFunctions)
}

// PCIeDevice gets the associated PCIe device for this function.
func (pciefunction *PCIeFunction) PCIeDevice() (*PCIeDevice, error) {
	if pciefunction.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(pciefunction.GetClient(), pciefunction.pcieDevice)
}

// Processor gets the processor that is hosted on this PCIe function.
func (pciefunction *PCIeFunction) Processor() (*Processor, error) {
	if pciefunction.processor == "" {
		return nil, nil
	}
	return GetProcessor(pciefunction.GetClient(), pciefunction.processor)
}

// StorageControllers gets the associated storage controllers.
func (pciefunction *PCIeFunction) StorageControllers() ([]*StorageController, error) {
	return common.GetObjects[StorageController](pciefunction.GetClient(), pciefunction.storageControllers)
}
