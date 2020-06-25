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
	DeviceID string `json:"DeviceId"`
	// FunctionID shall the PCIe device function number within a given PCIe
	// device.
	FunctionID int `json:"FunctionId"`
	// FunctionType shall be the function type of the PCIe device function such
	// as Physical or Virtual.
	FunctionType FunctionType
	// RevisionID shall be the PCI Revision ID of the PCIe device function.
	RevisionID string `json:"RevisionID"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SubsystemID shall be the PCI Subsystem ID of the PCIe device function.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall be the PCI Subsystem Vendor ID of the PCIe device
	// function.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall be the PCI Vendor ID of the PCIe device function.
	VendorID string `json:"VendorId"`
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
	// NetworkDeviceFunctions shall be an array of references to resources of
	// type NetworkDeviceFunction that represents the network device functions
	// associated with this resource.
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount is the number of network device functions.
	NetworkDeviceFunctionsCount int
	// PCIeDevice shall be a reference to the resource that this function is a
	// part of and shall reference a resource of type PCIeDevice.
	pcieDevice string
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
		Drives                      common.Links
		DrivesCount                 int `json:"Drives@odata.count"`
		EthernetInterfaces          common.Links
		EthernetInterfacesCount     int `json:"EthernetInterfaces@odata.count"`
		NetworkDeviceFunctions      common.Links
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		PCIeDevice                  common.Link
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
	pciefunction.drives = t.Links.Drives.ToStrings()
	pciefunction.DrivesCount = t.Links.DrivesCount
	pciefunction.ethernetInterfaces = t.Links.EthernetInterfaces.ToStrings()
	pciefunction.EthernetInterfacesCount = t.Links.EthernetInterfacesCount
	pciefunction.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	pciefunction.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	pciefunction.pcieDevice = string(t.Links.PCIeDevice)
	pciefunction.storageControllers = t.Links.StorageControllers.ToStrings()
	pciefunction.StorageControllersCount = t.Links.StorageControllersCount

	return nil
}

// GetPCIeFunction will get a PCIeFunction instance from the service.
func GetPCIeFunction(c common.Client, uri string) (*PCIeFunction, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pciefunction PCIeFunction
	err = json.NewDecoder(resp.Body).Decode(&pciefunction)
	if err != nil {
		return nil, err
	}

	pciefunction.SetClient(c)
	return &pciefunction, nil
}

// ListReferencedPCIeFunctions gets the collection of PCIeFunction from
// a provided reference.
func ListReferencedPCIeFunctions(c common.Client, link string) ([]*PCIeFunction, error) {
	var result []*PCIeFunction
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, pciefunctionLink := range links.ItemLinks {
		pciefunction, err := GetPCIeFunction(c, pciefunctionLink)
		if err != nil {
			return result, err
		}
		result = append(result, pciefunction)
	}

	return result, nil
}

// Drives gets the PCIe function's drives.
func (pciefunction *PCIeFunction) Drives() ([]*Drive, error) {
	var result []*Drive
	for _, driveLink := range pciefunction.drives {
		drive, err := GetDrive(pciefunction.Client, driveLink)
		if err != nil {
			return result, err
		}
		result = append(result, drive)
	}
	return result, nil
}

// EthernetInterfaces gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	var result []*EthernetInterface
	for _, ethLink := range pciefunction.ethernetInterfaces {
		eth, err := GetEthernetInterface(pciefunction.Client, ethLink)
		if err != nil {
			return result, err
		}
		result = append(result, eth)
	}
	return result, nil
}

// NetworkDeviceFunctions gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction
	for _, netLink := range pciefunction.networkDeviceFunctions {
		net, err := GetNetworkDeviceFunction(pciefunction.Client, netLink)
		if err != nil {
			return result, err
		}
		result = append(result, net)
	}
	return result, nil
}

// PCIeDevice gets the associated PCIe device for this function.
func (pciefunction *PCIeFunction) PCIeDevice() (*PCIeDevice, error) {
	if pciefunction.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(pciefunction.Client, pciefunction.pcieDevice)
}

// StorageControllers gets the associated storage controllers.
func (pciefunction *PCIeFunction) StorageControllers() ([]*StorageController, error) {
	var result []*StorageController
	for _, scLink := range pciefunction.storageControllers {
		sc, err := GetStorageController(pciefunction.Client, scLink)
		if err != nil {
			return result, err
		}
		result = append(result, sc)
	}
	return result, nil
}
