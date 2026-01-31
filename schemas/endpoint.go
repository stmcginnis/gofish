//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #Endpoint.v1_8_2.Endpoint

package schemas

import (
	"encoding/json"
	"strconv"
)

type EntityRole string

const (
	// InitiatorEntityRole The entity sends commands, messages, or other types of
	// requests to other entities on the fabric, but cannot receive commands from
	// other entities.
	InitiatorEntityRole EntityRole = "Initiator"
	// TargetEntityRole The entity receives commands, messages, or other types of
	// requests from other entities on the fabric, but cannot send commands to
	// other entities.
	TargetEntityRole EntityRole = "Target"
	// BothEntityRole The entity can both send and receive commands, messages, and
	// other requests to or from other entities on the fabric.
	BothEntityRole EntityRole = "Both"
)

type EntityType string

const (
	// StorageInitiatorEntityType shall indicate the entity this endpoint
	// represents is a storage initiator. The 'EntityLink' property, if present,
	// should be of type 'StorageController'.
	StorageInitiatorEntityType EntityType = "StorageInitiator"
	// RootComplexEntityType shall indicate the entity this endpoint represents is
	// a PCIe root complex. The 'EntityLink' property, if present, should be of
	// type 'ComputerSystem'.
	RootComplexEntityType EntityType = "RootComplex"
	// NetworkControllerEntityType shall indicate the entity this endpoint
	// represents is a network controller. The 'EntityLink' property, if present,
	// should be of type 'NetworkDeviceFunction' or EthernetInterface.
	NetworkControllerEntityType EntityType = "NetworkController"
	// DriveEntityType shall indicate the entity this endpoint represents is a
	// drive. The 'EntityLink' property, if present, should be of type 'Drive'.
	DriveEntityType EntityType = "Drive"
	// StorageExpanderEntityType shall indicate the entity this endpoint represents
	// is a storage expander. The 'EntityLink' property, if present, should be of
	// type 'Chassis'.
	StorageExpanderEntityType EntityType = "StorageExpander"
	// DisplayControllerEntityType shall indicate the entity this endpoint
	// represents is a display controller.
	DisplayControllerEntityType EntityType = "DisplayController"
	// BridgeEntityType shall indicate the entity this endpoint represents is a
	// PCIe bridge.
	BridgeEntityType EntityType = "Bridge"
	// ProcessorEntityType shall indicate the entity this endpoint represents is a
	// processor. The 'EntityLink' property, if present, should be of type
	// 'Processor'.
	ProcessorEntityType EntityType = "Processor"
	// VolumeEntityType shall indicate the entity this endpoint represents is a
	// volume. The 'EntityLink' property, if present, should be of type 'Volume'.
	VolumeEntityType EntityType = "Volume"
	// AccelerationFunctionEntityType shall indicate the entity this endpoint
	// represents is an acceleration function. The 'EntityLink' property, if
	// present, should be of type 'AccelerationFunction'.
	AccelerationFunctionEntityType EntityType = "AccelerationFunction"
	// MediaControllerEntityType shall indicate the entity this endpoint represents
	// is a media controller. The 'EntityLink' property, if present, should be of
	// type 'MediaController'.
	MediaControllerEntityType EntityType = "MediaController"
	// MemoryChunkEntityType shall indicate the entity this endpoint represents is
	// a memory chunk. The 'EntityLink' property, if present, should be of type
	// 'MemoryChunk'.
	MemoryChunkEntityType EntityType = "MemoryChunk"
	// SwitchEntityType shall indicate the entity this endpoint represents is a
	// switch and not an expander. The 'EntityLink' property, if present, should be
	// of type 'Switch'.
	SwitchEntityType EntityType = "Switch"
	// FabricBridgeEntityType shall indicate the entity this endpoint represents is
	// a fabric bridge. The 'EntityLink' property, if present, should be of type
	// 'FabricAdapter'.
	FabricBridgeEntityType EntityType = "FabricBridge"
	// ManagerEntityType shall indicate the entity this endpoint represents is a
	// manager. The 'EntityLink' property, if present, should be of type 'Manager'.
	ManagerEntityType EntityType = "Manager"
	// StorageSubsystemEntityType shall indicate the entity this endpoint
	// represents is a storage subsystem. The 'EntityLink' property, if present,
	// should be of type 'Storage'.
	StorageSubsystemEntityType EntityType = "StorageSubsystem"
	// MemoryEntityType shall indicate the entity this endpoint represents is a
	// memory device. The 'EntityLink' property, if present, should be of type
	// 'Memory'.
	MemoryEntityType EntityType = "Memory"
	// CXLDeviceEntityType shall indicate the entity this endpoint represents is a
	// CXL logical device. The 'EntityLink' property, if present, should be of type
	// 'CXLLogicalDevice'.
	CXLDeviceEntityType EntityType = "CXLDevice"
)

// Endpoint This resource contains a fabric endpoint for a Redfish
// implementation.
type Endpoint struct {
	Entity
	// ConnectedEntities shall contain all entities to which this endpoint allows
	// access.
	ConnectedEntities []ConnectedEntity
	// EndpointProtocol shall contain the protocol this endpoint uses to
	// communicate with other endpoints on this fabric.
	EndpointProtocol Protocol
	// HostReservationMemoryBytes shall contain the amount of memory in bytes that
	// the host should allocate to connect to this endpoint.
	HostReservationMemoryBytes *int `json:",omitempty"`
	// IPTransportDetails shall contain the details for each IP transport supported
	// by this endpoint.
	//
	// Version added: v1.1.0
	IPTransportDetails []IPTransportDetails
	// Identifiers shall be unique in the context of other endpoints that can
	// reached over the connected network.
	Identifiers []Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PciID shall contain the PCI ID of the endpoint.
	PciID PciID `json:"PciId"`
	// Redundancy shall show how this endpoint is grouped with other endpoints for
	// form redundancy sets.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// addressPools are the URIs for AddressPools.
	addressPools []string
	// connectedPorts are the URIs for ConnectedPorts.
	connectedPorts []string
	// connections are the URIs for Connections.
	connections []string
	// localPorts are the URIs for LocalPorts.
	localPorts []string
	// mutuallyExclusiveEndpoints are the URIs for MutuallyExclusiveEndpoints.
	mutuallyExclusiveEndpoints []string
	// networkDeviceFunction are the URIs for NetworkDeviceFunction.
	networkDeviceFunction []string
	// ports are the URIs for Ports.
	ports []string
	// zones are the URIs for Zones.
	zones []string
}

// UnmarshalJSON unmarshals a Endpoint object from the raw JSON.
func (e *Endpoint) UnmarshalJSON(b []byte) error {
	type temp Endpoint
	type eLinks struct {
		AddressPools               Links `json:"AddressPools"`
		ConnectedPorts             Links `json:"ConnectedPorts"`
		Connections                Links `json:"Connections"`
		LocalPorts                 Links `json:"LocalPorts"`
		MutuallyExclusiveEndpoints Links `json:"MutuallyExclusiveEndpoints"`
		NetworkDeviceFunction      Links `json:"NetworkDeviceFunction"`
		Ports                      Links `json:"Ports"`
		Zones                      Links `json:"Zones"`
	}
	var tmp struct {
		temp
		Links      eLinks
		Redundancy Link `json:"Redundancy"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = Endpoint(tmp.temp)

	// Extract the links to other entities for later
	e.addressPools = tmp.Links.AddressPools.ToStrings()
	e.connectedPorts = tmp.Links.ConnectedPorts.ToStrings()
	e.connections = tmp.Links.Connections.ToStrings()
	e.localPorts = tmp.Links.LocalPorts.ToStrings()
	e.mutuallyExclusiveEndpoints = tmp.Links.MutuallyExclusiveEndpoints.ToStrings()
	e.networkDeviceFunction = tmp.Links.NetworkDeviceFunction.ToStrings()
	e.ports = tmp.Links.Ports.ToStrings()
	e.zones = tmp.Links.Zones.ToStrings()
	e.redundancy = tmp.Redundancy.String()

	return nil
}

// GetEndpoint will get a Endpoint instance from the service.
func GetEndpoint(c Client, uri string) (*Endpoint, error) {
	return GetObject[Endpoint](c, uri)
}

// ListReferencedEndpoints gets the collection of Endpoint from
// a provided reference.
func ListReferencedEndpoints(c Client, link string) ([]*Endpoint, error) {
	return GetCollectionObjects[Endpoint](c, link)
}

// AddressPools gets the AddressPools linked resources.
func (e *Endpoint) AddressPools() ([]*AddressPool, error) {
	return GetObjects[AddressPool](e.client, e.addressPools)
}

// ConnectedPorts gets the ConnectedPorts linked resources.
func (e *Endpoint) ConnectedPorts() ([]*Port, error) {
	return GetObjects[Port](e.client, e.connectedPorts)
}

// Connections gets the Connections linked resources.
func (e *Endpoint) Connections() ([]*Connection, error) {
	return GetObjects[Connection](e.client, e.connections)
}

// LocalPorts gets the LocalPorts linked resources.
func (e *Endpoint) LocalPorts() ([]*Port, error) {
	return GetObjects[Port](e.client, e.localPorts)
}

// MutuallyExclusiveEndpoints gets the MutuallyExclusiveEndpoints linked resources.
func (e *Endpoint) MutuallyExclusiveEndpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](e.client, e.mutuallyExclusiveEndpoints)
}

// NetworkDeviceFunction gets the NetworkDeviceFunction linked resources.
func (e *Endpoint) NetworkDeviceFunction() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](e.client, e.networkDeviceFunction)
}

// Ports gets the Ports linked resources.
func (e *Endpoint) Ports() ([]*Port, error) {
	return GetObjects[Port](e.client, e.ports)
}

// Zones gets the Zones linked resources.
func (e *Endpoint) Zones() ([]*Zone, error) {
	return GetObjects[Zone](e.client, e.zones)
}

// Redundancy gets the Redundancy linked resource.
func (e *Endpoint) Redundancy() (*Redundancy, error) {
	if e.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](e.client, e.redundancy)
}

// ConnectedEntity shall represent a remote resource that is connected to a
// network accessible to an endpoint.
type ConnectedEntity struct {
	// EntityLink shall contain a link to an entity of the type specified by the
	// description of the 'EntityType' property value.
	entityLink string
	// EntityPciID shall contain the PCI ID of the connected PCIe entity.
	EntityPciID PciID `json:"EntityPciId"`
	// EntityRole shall indicate if the specified entity is an initiator, target,
	// or both.
	EntityRole EntityRole
	// EntityType shall indicate if type of connected entity.
	EntityType EntityType
	// GenZ shall contain the Gen-Z related properties for the entity.
	//
	// Version added: v1.4.0
	GenZ GenZ
	// Identifiers shall be unique in the context of other resources that can
	// reached over the connected network.
	Identifiers []Identifier
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PciClassCode shall contain the PCI Class Code, Subclass, and Programming
	// Interface of the PCIe device function.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'ClassCode' property
	// inside the 'EntityPciId' object.
	PciClassCode string
	// PciFunctionNumber shall contain the PCI Function Number of the connected
	// PCIe entity.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'FunctionNumber' property
	// inside the 'EntityPciId' object.
	PciFunctionNumber *int `json:",omitempty"`
}

// UnmarshalJSON unmarshals a ConnectedEntity object from the raw JSON.
func (c *ConnectedEntity) UnmarshalJSON(b []byte) error {
	type temp ConnectedEntity
	var tmp struct {
		temp
		EntityLink Link `json:"EntityLink"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ConnectedEntity(tmp.temp)

	// Extract the links to other entities for later
	c.entityLink = tmp.EntityLink.String()

	return nil
}

// EntityLink gets the EntityLink linked resource.
func (c *ConnectedEntity) EntityLink(client Client) (*Resource, error) {
	if c.entityLink == "" {
		return nil, nil
	}
	return GetObject[Resource](client, c.entityLink)
}

// GCID shall contain the Gen-Z Core Specification-defined Global Component ID.
type GCID struct {
	// CID shall contain the 12 bit component identifier portion of the GCID of the
	// entity.
	//
	// Version added: v1.4.0
	CID string
	// SID shall contain the 16 bit subnet identifier portion of the GCID of the
	// entity.
	//
	// Version added: v1.4.0
	SID string
}

// GenZ shall contain the Gen-Z related properties for an entity.
type GenZ struct {
	// AccessKey shall contain the Gen-Z Core Specification-defined 6 bit Access
	// Key for the entity.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the 'ConnectionKeys' property
	// in the 'Connection' resource.
	AccessKey string
	// GCID shall contain the Gen-Z Core Specification-defined Global Component ID
	// for the entity.
	//
	// Version added: v1.4.0
	GCID GCID
	// RegionKey shall contain the Gen-Z Core Specification-defined 32 bit Region
	// Key for the entity.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the 'ConnectionKeys' property
	// in the 'Connection' resource.
	RegionKey string
}

// IPTransportDetails shall contain properties that specify the details of the
// transport supported by the endpoint.
type IPTransportDetails struct {
	// IPv4Address shall contain the IPv4 address.
	//
	// Version added: v1.1.0
	IPv4Address IPv4Address
	// IPv6Address shall contain the IPv6 address.
	//
	// Version added: v1.1.0
	IPv6Address IPv6Address
	// Port shall contain a specified UDP or TCP port number used for communication
	// with the endpoint.
	//
	// Version added: v1.1.0
	Port uint
	// TransportProtocol shall contain the protocol used by the connection entity.
	//
	// Version added: v1.1.0
	TransportProtocol Protocol
}

// PciID shall describe a PCI ID.
type PciID struct {
	// ClassCode shall contain the PCI Class Code, Subclass, and Programming
	// Interface of the PCIe device function.
	//
	// Version added: v1.2.0
	ClassCode string
	// DeviceID shall contain the PCI Device ID of the PCIe device function.
	DeviceID string `json:"DeviceId"`
	// FunctionNumber shall contain the PCI Function Number of the connected PCIe
	// entity.
	//
	// Version added: v1.2.0
	FunctionNumber *int `json:",omitempty"`
	// SubsystemID shall contain the PCI Subsystem ID of the PCIe device function.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall contain the PCI Subsystem Vendor ID of the PCIe
	// device function.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall contain the PCI Vendor ID of the PCIe device function.
	VendorID string `json:"VendorId"`
}

// UnmarshalJSON unmarshals a ConnectedEntity object from the raw JSON.
func (p *PciID) UnmarshalJSON(b []byte) error {
	type temp PciID
	var tmp struct {
		temp
		FunctionNumber any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PciID(tmp.temp)

	// Extract the links to other entities for later
	if tmp.FunctionNumber != nil {
		switch val := tmp.FunctionNumber.(type) {
		case string:
			fn, err := strconv.Atoi(val)
			if err == nil {
				p.FunctionNumber = &fn
			}
		case int:
			fn := val
			p.FunctionNumber = &fn
		}
	}

	return nil
}
