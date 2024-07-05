//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// EntityRole is the role of the endpoint.
type EntityRole string

const (
	// InitiatorEntityRole means the entity is acting as an initiator.
	InitiatorEntityRole EntityRole = "Initiator"
	// TargetEntityRole means the entity is acting as a target.
	TargetEntityRole EntityRole = "Target"
	// BothEntityRole means the entity is acting as both an initiator and a target.
	BothEntityRole EntityRole = "Both"
)

// EntityType is the type of endpoint.
type EntityType string

const (
	// StorageInitiatorEntityType shall indicate the entity this endpoint represents is a storage initiator. The
	// EntityLink property, if present, should be of type StorageController.
	StorageInitiatorEntityType EntityType = "StorageInitiator"
	// RootComplexEntityType shall indicate the entity this endpoint represents is a PCIe root complex. The EntityLink
	// property, if present, should be of type ComputerSystem.
	RootComplexEntityType EntityType = "RootComplex"
	// NetworkControllerEntityType shall indicate the entity this endpoint represents is a network controller. The
	// EntityLink property, if present, should be of type NetworkDeviceFunction or EthernetInterface.
	NetworkControllerEntityType EntityType = "NetworkController"
	// DriveEntityType shall indicate the entity this endpoint represents is a drive. The EntityLink property, if
	// present, should be of type Drive.
	DriveEntityType EntityType = "Drive"
	// StorageExpanderEntityType shall indicate the entity this endpoint represents is a storage expander. The
	// EntityLink property, if present, should be of type Chassis.
	StorageExpanderEntityType EntityType = "StorageExpander"
	// DisplayControllerEntityType shall indicate the entity this endpoint represents is a display controller.
	DisplayControllerEntityType EntityType = "DisplayController"
	// BridgeEntityType shall indicate the entity this endpoint represents is a PCIe bridge.
	BridgeEntityType EntityType = "Bridge"
	// ProcessorEntityType shall indicate the entity this endpoint represents is a processor. The EntityLink property,
	// if present, should be of type Processor.
	ProcessorEntityType EntityType = "Processor"
	// VolumeEntityType shall indicate the entity this endpoint represents is a volume. The EntityLink property, if
	// present, should be of type Volume.
	VolumeEntityType EntityType = "Volume"
	// AccelerationFunctionEntityType shall indicate the entity this endpoint represents is an acceleration function.
	// The EntityLink property, if present, should be of type AccelerationFunction.
	AccelerationFunctionEntityType EntityType = "AccelerationFunction"
	// MediaControllerEntityType shall indicate the entity this endpoint represents is a media controller. The
	// EntityLink property, if present, should be of type MediaController.
	MediaControllerEntityType EntityType = "MediaController"
	// MemoryChunkEntityType shall indicate the entity this endpoint represents is a memory chunk. The EntityLink
	// property, if present, should be of type MemoryChunk.
	MemoryChunkEntityType EntityType = "MemoryChunk"
	// SwitchEntityType shall indicate the entity this endpoint represents is a switch and not an expander. The
	// EntityLink property, if present, should be of type Switch.
	SwitchEntityType EntityType = "Switch"
	// FabricBridgeEntityType shall indicate the entity this endpoint represents is a fabric bridge. The EntityLink
	// property, if present, should be of type FabricAdapter.
	FabricBridgeEntityType EntityType = "FabricBridge"
	// ManagerEntityType shall indicate the entity this endpoint represents is a manager. The EntityLink property, if
	// present, should be of type Manager.
	ManagerEntityType EntityType = "Manager"
	// StorageSubsystemEntityType shall indicate the entity this endpoint represents is a storage subsystem. The
	// EntityLink property, if present, should be of type Storage.
	StorageSubsystemEntityType EntityType = "StorageSubsystem"
	// MemoryEntityType shall indicate the entity this endpoint represents is a memory device. The EntityLink property,
	// if present, should be of type Memory.
	MemoryEntityType EntityType = "Memory"
	// CXLDeviceEntityType shall indicate the entity this endpoint represents is a CXL logical device. The EntityLink
	// property, if present, should be of type CXLLogicalDevice.
	CXLDeviceEntityType EntityType = "CXLDevice"
)

// ConnectedEntity shall represent a remote resource that is
// connected to a network accessible to an endpoint.
type ConnectedEntity struct {
	// entityLink shall be a reference to an entity of the
	// type specified by the description of the value of the EntityType
	// property.
	// entityLink common.Link
	// EntityPciID shall be the PCI ID of the connected PCIe entity.
	EntityPciID PciID `json:"entityPciId"`
	// entityRole shall indicate if the specified entity is an initiator,
	// target, or both.
	EntityRole EntityRole
	// entityType shall indicate if type of connected entity.
	EntityType EntityType
	// GenZ shall contain the Gen-Z related properties for the entity.
	GenZ GenZ
	// Identifiers for the remote entity shall be unique in
	// the context of other resources that can reached over the connected
	// network.
	Identifiers []common.Identifier
}

// Endpoint is used to represent a fabric endpoint for a Redfish
// implementation.
type Endpoint struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConnectedEntities shall contain all the entities which this endpoint
	// allows access to.
	ConnectedEntities []ConnectedEntity
	// Description provides a description of this resource.
	Description string
	// EndpointProtocol shall contain the protocol this endpoint uses to
	// communicate with other endpoints on this fabric.
	EndpointProtocol common.Protocol
	// HostReservationMemoryBytes shall be the amount of memory in Bytes that
	// the Host should allocate to connect to this endpoint.
	HostReservationMemoryBytes int64
	// IPTransportDetails shall contain the details for each IP transport
	// supported by this endpoint.
	IPTransportDetails []IPTransportDetails
	// Identifiers shall be unique in the context of other endpoints that can
	// reached over the connected network.
	Identifiers []common.Identifier
	// PciID shall be the PCI ID of the endpoint.
	PciID PciID `json:"PciId"`
	// Redundancy is used to show how this endpoint is grouped with other
	// endpoints to form redundancy sets.
	Redundancy []Redundancy
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status

	// MutuallyExclusiveEndpoints shall be an array of references of type
	// Endpoint that cannot be used in a zone if this endpoint is used in a zone.
	mutuallyExclusiveEndpoints []string
	// MutuallyExclusiveEndpointsCount is the number of MutuallyExclusiveEndpoints.
	MutuallyExclusiveEndpointsCount int
	// NetworkDeviceFunction shall be a reference to a NetworkDeviceFunction
	// resource, with which this endpoint is associated.
	networkDeviceFunction []string
	// NetworkDeviceFunctionCount is the number of NetworkDeviceFunctions.
	NetworkDeviceFunctionCount int
	// Ports shall be an array of references of type Port that are utilized by
	// this endpoint.
	ports []string
	// PortsCount is the number of Ports.
	PortsCount int
	// addressPools shall contain an array of links to
	// resources of type AddressPool with which this endpoint is associated.
	addressPools []string
	// AddressPoolsCount is the number of AddressPools.
	AddressPoolsCount int
	// connectedPorts shall contain an array of links to
	// resources of type Port that represent ports associated with this
	// endpoint.
	connectedPorts []string
	// ConnectedPortCount is the number of ConnectedPorts.
	ConnectedPortsCount int
}

// UnmarshalJSON unmarshals a Endpoint object from the raw JSON.
func (endpoint *Endpoint) UnmarshalJSON(b []byte) error {
	type temp Endpoint
	type links struct {
		// AddressPools shall contain an array of links to
		// resources of type AddressPool with which this endpoint is associated.
		AddressPools common.Links
		// AddressPools@odata.count is
		AddressPoolsCount int `json:"AddressPools@odata.count"`
		// ConnectedPorts shall contain an array of links to
		// resources of type Port that represent ports associated with this
		// endpoint.
		ConnectedPorts common.Links
		// ConnectedPorts@odata.count is
		ConnectedPortsCount int `json:"ConnectedPorts@odata.count"`
		// MutuallyExclusiveEndpoints shall be an array of references of type
		// Endpoint that cannot be used in a zone if this endpoint is used in a zone.
		MutuallyExclusiveEndpoints common.Links
		// MutuallyExclusiveEndpointsCount is the number of MutuallyExclusiveEndpoints.
		MutuallyExclusiveEndpointsCount int `json:"MutuallyExclusiveEndpoints@odata.count"`
		// NetworkDeviceFunction shall be a reference to a NetworkDeviceFunction
		// resource, with which this endpoint is associated.
		NetworkDeviceFunction common.Links
		// NetworkDeviceFunctionCount is the number of NetworkDeviceFunctions.
		NetworkDeviceFunctionCount int `json:"NetworkDeviceFunction@odata.count"`
		// Ports shall be an array of references of type Port that are utilized by
		// this endpoint.
		Ports common.Links
		// PortsCount is the number of Ports.
		PortsCount int `json:"Ports@odata.count"`
	}
	var t struct {
		temp
		Links      links
		Redundancy common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*endpoint = Endpoint(t.temp)
	endpoint.addressPools = t.Links.AddressPools.ToStrings()
	endpoint.AddressPoolsCount = t.Links.AddressPoolsCount
	endpoint.connectedPorts = t.Links.ConnectedPorts.ToStrings()
	endpoint.ConnectedPortsCount = t.Links.ConnectedPortsCount
	endpoint.mutuallyExclusiveEndpoints = t.Links.MutuallyExclusiveEndpoints.ToStrings()
	endpoint.MutuallyExclusiveEndpointsCount = t.Links.MutuallyExclusiveEndpointsCount
	endpoint.networkDeviceFunction = t.Links.NetworkDeviceFunction.ToStrings()
	endpoint.NetworkDeviceFunctionCount = t.Links.NetworkDeviceFunctionCount
	endpoint.ports = t.Links.Ports.ToStrings()
	endpoint.PortsCount = t.Links.PortsCount

	return nil
}

// GetEndpoint will get a Endpoint instance from the service.
func GetEndpoint(c common.Client, uri string) (*Endpoint, error) {
	return common.GetObject[Endpoint](c, uri)
}

// ListReferencedEndpoints gets the collection of Endpoint from
// a provided reference.
func ListReferencedEndpoints(c common.Client, link string) ([]*Endpoint, error) {
	return common.GetCollectionObjects[Endpoint](c, link)
}

// GCID shall contain the Gen-Z Core Specification-defined Global
// Component ID.
type GCID struct {
	// CID shall contain the 12 bit component identifier
	// portion of the GCID of the entity.
	CID string
	// SID shall contain the 16 bit subnet identifier
	// portion of the GCID of the entity.
	SID string
}

// GenZ shall contain the Gen-Z related properties for an entity.
type GenZ struct {
	// AccessKey shall contain the Gen-Z Core Specification-
	// defined 6 bit Access Key for the entity.
	AccessKey string
	// GCID shall contain the Gen-Z Core Specification-
	// defined Global Component ID for the entity.
	GCID GCID
	// RegionKey shall contain the Gen-Z Core Specification-
	// defined 32 bit Region Key for the entity.
	RegionKey string
}

// IPTransportDetails shall contain properties which specify
// the details of the transport supported by the endpoint.
type IPTransportDetails struct {
	// IPv4Address shall specify the IPv4Address.
	IPv4Address string
	// IPv6Address shall specify the IPv6Address.
	IPv6Address string
	// Port shall be an specify UDP or TCP port number used for communication
	// with the Endpoint.
	Port int
	// TransportProtocol is used by the connection entity.
	TransportProtocol common.Protocol
}

// PciID shall describe a PCI ID.
type PciID struct {
	// ClassCode shall be the PCI Class Code,
	// Subclass code, and Programming Interface code of the PCIe device
	// function.
	ClassCode string
	// DeviceID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	DeviceID string `json:"DeviceId"`
	// FunctionNumber shall be the PCI Function Number of the connected PCIe entity.
	FunctionNumber string
	// SubsystemID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall be the PCI Vendor ID of the PCIe device function.
	VendorID string `json:"VendorId"`
}
