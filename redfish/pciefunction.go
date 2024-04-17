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
	pciefunction.pcieDevice = t.Links.PCIeDevice.String()
	pciefunction.storageControllers = t.Links.StorageControllers.ToStrings()
	pciefunction.StorageControllersCount = t.Links.StorageControllersCount

	return nil
}

// GetPCIeFunction will get a PCIeFunction instance from the service.
func GetPCIeFunction(c common.Client, uri string) (*PCIeFunction, error) {
	var pcieFunction PCIeFunction
	return &pcieFunction, pcieFunction.Get(c, uri, &pcieFunction)
}

// ListReferencedPCIeFunctions gets the collection of PCIeFunction from
// a provided reference.
func ListReferencedPCIeFunctions(c common.Client, link string) ([]*PCIeFunction, error) {
	var result []*PCIeFunction
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *PCIeFunction
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		pciefunction, err := GetPCIeFunction(c, link)
		ch <- GetResult{Item: pciefunction, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Drives gets the PCIe function's drives.
func (pciefunction *PCIeFunction) Drives() ([]*Drive, error) {
	var result []*Drive

	collectionError := common.NewCollectionError()
	for _, driveLink := range pciefunction.drives {
		drive, err := GetDrive(pciefunction.GetClient(), driveLink)
		if err != nil {
			collectionError.Failures[driveLink] = err
		} else {
			result = append(result, drive)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// EthernetInterfaces gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	var result []*EthernetInterface

	collectionError := common.NewCollectionError()
	for _, ethLink := range pciefunction.ethernetInterfaces {
		eth, err := GetEthernetInterface(pciefunction.GetClient(), ethLink)
		if err != nil {
			collectionError.Failures[ethLink] = err
		} else {
			result = append(result, eth)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// NetworkDeviceFunctions gets the PCIe function's ethernet interfaces.
func (pciefunction *PCIeFunction) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction

	collectionError := common.NewCollectionError()
	for _, netLink := range pciefunction.networkDeviceFunctions {
		net, err := GetNetworkDeviceFunction(pciefunction.GetClient(), netLink)
		if err != nil {
			collectionError.Failures[netLink] = err
		} else {
			result = append(result, net)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// PCIeDevice gets the associated PCIe device for this function.
func (pciefunction *PCIeFunction) PCIeDevice() (*PCIeDevice, error) {
	if pciefunction.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(pciefunction.GetClient(), pciefunction.pcieDevice)
}

// StorageControllers gets the associated storage controllers.
func (pciefunction *PCIeFunction) StorageControllers() ([]*StorageController, error) {
	var result []*StorageController

	collectionError := common.NewCollectionError()
	for _, scLink := range pciefunction.storageControllers {
		sc, err := GetStorageController(pciefunction.GetClient(), scLink)
		if err != nil {
			collectionError.Failures[scLink] = err
		} else {
			result = append(result, sc)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
