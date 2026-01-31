//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #NetworkDeviceFunction.v1_11_1.NetworkDeviceFunction

package schemas

import (
	"encoding/json"
)

type BootMode string

const (
	// DisabledBootMode Do not indicate to UEFI/BIOS that this device is bootable.
	DisabledBootMode BootMode = "Disabled"
	// PXEBootMode Boot this device by using the embedded PXE support. Only
	// applicable if the 'NetDevFuncType' is 'Ethernet' or 'InfiniBand'.
	PXEBootMode BootMode = "PXE"
	// iSCSIBootMode Boot this device by using the embedded iSCSI boot support and
	// configuration. Only applicable if the 'NetDevFuncType' is 'iSCSI' or
	// 'Ethernet'.
	ISCSIBootMode BootMode = "iSCSI"
	// FibreChannelBootMode Boot this device by using the embedded Fibre Channel
	// support and configuration. Only applicable if the 'NetDevFuncType' is
	// 'FibreChannel'.
	FibreChannelBootMode BootMode = "FibreChannel"
	// FibreChannelOverEthernetBootMode Boot this device by using the embedded
	// Fibre Channel over Ethernet (FCoE) boot support and configuration. Only
	// applicable if the 'NetDevFuncType' is 'FibreChannelOverEthernet'.
	FibreChannelOverEthernetBootMode BootMode = "FibreChannelOverEthernet"
	// HTTPBootMode Boot this device by using the embedded HTTP/HTTPS support. Only
	// applicable if the 'NetDevFuncType' is 'Ethernet'.
	HTTPBootMode BootMode = "HTTP"
)

type IPAddressType string

const (
	// IPv4IPAddressType IPv4 addressing is used for all IP-fields in this object.
	IPv4IPAddressType IPAddressType = "IPv4"
	// IPv6IPAddressType IPv6 addressing is used for all IP-fields in this object.
	IPv6IPAddressType IPAddressType = "IPv6"
)

type NetworkDeviceTechnology string

const (
	// DisabledNetworkDeviceTechnology Neither enumerated nor visible to the
	// operating system.
	DisabledNetworkDeviceTechnology NetworkDeviceTechnology = "Disabled"
	// EthernetNetworkDeviceTechnology Appears to the operating system as an
	// Ethernet device.
	EthernetNetworkDeviceTechnology NetworkDeviceTechnology = "Ethernet"
	// FibreChannelNetworkDeviceTechnology Appears to the operating system as a
	// Fibre Channel device.
	FibreChannelNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannel"
	// iSCSINetworkDeviceTechnology Appears to the operating system as an iSCSI
	// device.
	ISCSINetworkDeviceTechnology NetworkDeviceTechnology = "iSCSI"
	// FibreChannelOverEthernetNetworkDeviceTechnology Appears to the operating
	// system as an FCoE device.
	FibreChannelOverEthernetNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannelOverEthernet"
	// InfiniBandNetworkDeviceTechnology Appears to the operating system as an
	// InfiniBand device.
	InfiniBandNetworkDeviceTechnology NetworkDeviceTechnology = "InfiniBand"
)

type WWNSource string

const (
	// ConfiguredLocallyWWNSource The set of FC/FCoE boot targets was applied
	// locally through API or UI.
	ConfiguredLocallyWWNSource WWNSource = "ConfiguredLocally"
	// ProvidedByFabricWWNSource The set of FC/FCoE boot targets was applied by the
	// Fibre Channel fabric.
	ProvidedByFabricWWNSource WWNSource = "ProvidedByFabric"
)

// NetworkDeviceFunction shall represent a logical interface that a network
// adapter exposes in a Redfish implementation.
type NetworkDeviceFunction struct {
	Entity
	// AllowDeny shall contain a link to a resource collection of type
	// 'AllowDenyCollection' that contains the permissions for packets leaving and
	// arriving to this network device function.
	//
	// Version added: v1.7.0
	allowDeny string
	// AssignablePhysicalNetworkPorts shall contain an array of links to resources
	// of type 'Port' that are the physical ports or logical aggregated ports to
	// which this network device function can be assigned.
	//
	// Version added: v1.5.0
	assignablePhysicalNetworkPorts []string
	// AssignablePhysicalNetworkPortsCount
	AssignablePhysicalNetworkPortsCount int `json:"AssignablePhysicalNetworkPorts@odata.count"`
	// AssignablePhysicalPorts shall contain an array of links to resources of type
	// 'NetworkPort' that are the physical ports to which this network device
	// function can be assigned.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the
	// 'AssignablePhysicalNetworkPorts' property.
	assignablePhysicalPorts []string
	// AssignablePhysicalPortsCount
	AssignablePhysicalPortsCount int `json:"AssignablePhysicalPorts@odata.count"`
	// BootMode shall contain the boot mode configured for this network device
	// function. If the value is not 'Disabled', this network device function shall
	// be configured for boot by using the specified technology.
	BootMode BootMode
	// DeviceEnabled shall indicate whether the network device function is enabled.
	// The operating system shall not enumerate or see disabled network device
	// functions.
	DeviceEnabled bool
	// Ethernet shall contain Ethernet capabilities, status, and configuration
	// values for this network device function.
	Ethernet Ethernet
	// FibreChannel shall contain Fibre Channel capabilities, status, and
	// configuration values for this network device function.
	FibreChannel FibreChannel
	// HTTPBoot shall contain HTTP and HTTPS boot capabilities, status, and
	// configuration values for this network device function.
	//
	// Version added: v1.9.0
	HTTPBoot HTTPBoot
	// InfiniBand shall contain InfiniBand capabilities, status, and configuration
	// values for this network device function.
	//
	// Version added: v1.5.0
	InfiniBand InfiniBand
	// Limits shall contain an array of byte and packet limits for this network
	// device function.
	//
	// Version added: v1.7.0
	Limits []Limit
	// MaxVirtualFunctions shall contain the number of virtual functions that are
	// available for this network device function.
	MaxVirtualFunctions *int `json:",omitempty"`
	// Metrics shall contain a link to a resource of type
	// 'NetworkDeviceFunctionMetrics' that contains the metrics associated with
	// this network function.
	//
	// Version added: v1.6.0
	metrics string
	// NetDevFuncCapabilities shall contain an array of capabilities for this
	// network device function.
	NetDevFuncCapabilities []NetworkDeviceTechnology
	// NetDevFuncType shall contain the configured capability of this network
	// device function.
	NetDevFuncType NetworkDeviceTechnology
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalPortAssignment shall contain a link to a resource of type
	// 'NetworkPort' that is the physical port to which this network device
	// function is currently assigned. This value shall be one of the
	// 'AssignablePhysicalPorts' array members.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated and moved to the 'Links' property to avoid
	// loops on expand.
	physicalPortAssignment string
	// SAVIEnabled shall indicate if the RFC7039-defined Source Address Validation
	// Improvement (SAVI) is enabled for this network device function.
	//
	// Version added: v1.7.0
	SAVIEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// VirtualFunctionAllocation shall contain the number virtual functions
	// allocated to this device. This property should contain a value that is a
	// multiple of the value contained by the 'MinAssignmentGroupSize' property of
	// the corresponding 'Controllers' array member within the parent
	// 'NetworkAdapter' resource. The value shall not exceed the value contained in
	// the 'MaxVirtualFunctions' property.
	//
	// Version added: v1.11.0
	VirtualFunctionAllocation int
	// VirtualFunctionsEnabled shall indicate whether single root input/output
	// virtualization (SR-IOV) virtual functions are enabled for this network
	// device function.
	VirtualFunctionsEnabled bool
	// ISCSIBoot shall contain iSCSI boot capabilities, status, and configuration
	// values for this network device function.
	ISCSIBoot ISCSIBoot
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// ethernetInterface is the URI for EthernetInterface.
	ethernetInterface string
	// ethernetInterfaces are the URIs for EthernetInterfaces.
	ethernetInterfaces []string
	// offloadProcessors are the URIs for OffloadProcessors.
	offloadProcessors []string
	// offloadSystem is the URI for OffloadSystem.
	offloadSystem string
	// pCIeFunction is the URI for PCIeFunction.
	pCIeFunction string
	// physicalNetworkPortAssignment is the URI for PhysicalNetworkPortAssignment.
	physicalNetworkPortAssignment string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a NetworkDeviceFunction object from the raw JSON.
func (n *NetworkDeviceFunction) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunction
	type nLinks struct {
		Endpoints                     Links `json:"Endpoints"`
		EthernetInterface             Link  `json:"EthernetInterface"`
		EthernetInterfaces            Links `json:"EthernetInterfaces"`
		OffloadProcessors             Links `json:"OffloadProcessors"`
		OffloadSystem                 Link  `json:"OffloadSystem"`
		PCIeFunction                  Link  `json:"PCIeFunction"`
		PhysicalNetworkPortAssignment Link  `json:"PhysicalNetworkPortAssignment"`
		PhysicalPortAssignment        Link  `json:"PhysicalPortAssignment"`
	}
	var tmp struct {
		temp
		Links                          nLinks
		AllowDeny                      Link  `json:"AllowDeny"`
		AssignablePhysicalNetworkPorts Links `json:"AssignablePhysicalNetworkPorts"`
		AssignablePhysicalPorts        Links `json:"AssignablePhysicalPorts"`
		Metrics                        Link  `json:"Metrics"`
		PhysicalPortAssignment         Link  `json:"PhysicalPortAssignment"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkDeviceFunction(tmp.temp)

	// Extract the links to other entities for later
	n.endpoints = tmp.Links.Endpoints.ToStrings()
	n.ethernetInterface = tmp.Links.EthernetInterface.String()
	n.ethernetInterfaces = tmp.Links.EthernetInterfaces.ToStrings()
	n.offloadProcessors = tmp.Links.OffloadProcessors.ToStrings()
	n.offloadSystem = tmp.Links.OffloadSystem.String()
	n.pCIeFunction = tmp.Links.PCIeFunction.String()
	n.physicalNetworkPortAssignment = tmp.Links.PhysicalNetworkPortAssignment.String()
	n.physicalPortAssignment = tmp.Links.PhysicalPortAssignment.String()
	n.allowDeny = tmp.AllowDeny.String()
	n.assignablePhysicalNetworkPorts = tmp.AssignablePhysicalNetworkPorts.ToStrings()
	n.assignablePhysicalPorts = tmp.AssignablePhysicalPorts.ToStrings()
	n.metrics = tmp.Metrics.String()
	if tmp.Links.PhysicalPortAssignment.String() == "" {
		n.physicalPortAssignment = tmp.PhysicalPortAssignment.String()
	}

	// This is a read/write object, so we need to save the raw object data for later
	n.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (n *NetworkDeviceFunction) Update() error {
	readWriteFields := []string{
		"BootMode",
		"DeviceEnabled",
		"NetDevFuncType",
		"SAVIEnabled",
		"VirtualFunctionAllocation",
	}

	return n.UpdateFromRawData(n, n.RawData, readWriteFields)
}

// GetNetworkDeviceFunction will get a NetworkDeviceFunction instance from the service.
func GetNetworkDeviceFunction(c Client, uri string) (*NetworkDeviceFunction, error) {
	return GetObject[NetworkDeviceFunction](c, uri)
}

// ListReferencedNetworkDeviceFunctions gets the collection of NetworkDeviceFunction from
// a provided reference.
func ListReferencedNetworkDeviceFunctions(c Client, link string) ([]*NetworkDeviceFunction, error) {
	return GetCollectionObjects[NetworkDeviceFunction](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (n *NetworkDeviceFunction) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](n.client, n.endpoints)
}

// EthernetInterface gets the EthernetInterface linked resource.
func (n *NetworkDeviceFunction) EthernetInterface() (*EthernetInterface, error) {
	if n.ethernetInterface == "" {
		return nil, nil
	}
	return GetObject[EthernetInterface](n.client, n.ethernetInterface)
}

// EthernetInterfaces gets the EthernetInterfaces linked resources.
func (n *NetworkDeviceFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	return GetObjects[EthernetInterface](n.client, n.ethernetInterfaces)
}

// OffloadProcessors gets the OffloadProcessors linked resources.
func (n *NetworkDeviceFunction) OffloadProcessors() ([]*Processor, error) {
	return GetObjects[Processor](n.client, n.offloadProcessors)
}

// OffloadSystem gets the OffloadSystem linked resource.
func (n *NetworkDeviceFunction) OffloadSystem() (*ComputerSystem, error) {
	if n.offloadSystem == "" {
		return nil, nil
	}
	return GetObject[ComputerSystem](n.client, n.offloadSystem)
}

// PCIeFunction gets the PCIeFunction linked resource.
func (n *NetworkDeviceFunction) PCIeFunction() (*PCIeFunction, error) {
	if n.pCIeFunction == "" {
		return nil, nil
	}
	return GetObject[PCIeFunction](n.client, n.pCIeFunction)
}

// PhysicalNetworkPortAssignment gets the PhysicalNetworkPortAssignment linked resource.
func (n *NetworkDeviceFunction) PhysicalNetworkPortAssignment() (*Port, error) {
	if n.physicalNetworkPortAssignment == "" {
		return nil, nil
	}
	return GetObject[Port](n.client, n.physicalNetworkPortAssignment)
}

// AllowDeny gets the AllowDeny collection.
func (n *NetworkDeviceFunction) AllowDeny() ([]*AllowDeny, error) {
	if n.allowDeny == "" {
		return nil, nil
	}
	return GetCollectionObjects[AllowDeny](n.client, n.allowDeny)
}

// AssignablePhysicalNetworkPorts gets the AssignablePhysicalNetworkPorts linked resources.
func (n *NetworkDeviceFunction) AssignablePhysicalNetworkPorts() ([]*Port, error) {
	return GetObjects[Port](n.client, n.assignablePhysicalNetworkPorts)
}

// AssignablePhysicalPorts gets the AssignablePhysicalPorts linked resources.
func (n *NetworkDeviceFunction) AssignablePhysicalPorts() ([]*NetworkPort, error) {
	return GetObjects[NetworkPort](n.client, n.assignablePhysicalPorts)
}

// Metrics gets the Metrics linked resource.
func (n *NetworkDeviceFunction) Metrics() (*NetworkDeviceFunctionMetrics, error) {
	if n.metrics == "" {
		return nil, nil
	}
	return GetObject[NetworkDeviceFunctionMetrics](n.client, n.metrics)
}

// PhysicalPortAssignment gets the PhysicalPortAssignment linked resource.
func (n *NetworkDeviceFunction) PhysicalPortAssignment() (*NetworkPort, error) {
	if n.physicalPortAssignment == "" {
		return nil, nil
	}
	return GetObject[NetworkPort](n.client, n.physicalPortAssignment)
}

// BootTargets shall describe a Fibre Channel boot target configured for a
// network device function.
type BootTargets struct {
	// BootPriority shall contain the relative priority for this entry in the boot
	// targets array. Lower numbers shall represent higher priority, with zero
	// being the highest priority. The 'BootPriority' shall be unique for all
	// entries of the 'BootTargets' array.
	BootPriority *int `json:",omitempty"`
	// LUNID shall contain the logical unit number (LUN) ID from which to boot on
	// the device to which the corresponding WWPN refers.
	LUNID string
	// WWPN shall contain World Wide Port Name (WWPN) from which to boot.
	WWPN string
}

// Ethernet shall describe the Ethernet capabilities, status, and configuration
// values for a network device function.
type Ethernet struct {
	// AdditionalProtocols shall contain the list of protocols supported by the
	// hardware or firmware on the device.
	//
	// Version added: v1.10.0
	AdditionalProtocols []Protocol
	// EthernetInterfaces shall contain a link to a collection of type
	// 'EthernetInterfaceCollection' that represent the Ethernet interfaces present
	// on this network device function. This property shall only be present if this
	// NetworkDeviceFunction is not associated with a ComputerSystem, such as when
	// in a ResourcePool or representing an Ethernet based storage device.
	//
	// Version added: v1.7.0
	ethernetInterfaces string
	// MACAddress shall contain the effective current MAC address of this network
	// device function. If an assignable MAC address is not supported, this is a
	// read-only alias of the 'PermanentMACAddress'.
	MACAddress string
	// MTUSize The hardware maximum transmission unit (MTU) configured for this
	// network device function. This value serves as a default for the OS driver
	// when booting, but may be overridden by the OS. After the OS boots and while
	// the driver is loaded, the effective MTU size may be found in the associated
	// 'EthernetInterface' resource.
	MTUSize *int `json:",omitempty"`
	// MTUSizeMaximum shall contain the largest maximum transmission unit (MTU)
	// size supported for this network device function.
	//
	// Version added: v1.5.0
	MTUSizeMaximum *int `json:",omitempty"`
	// PermanentMACAddress shall contain the permanent MAC Address of this
	// function. Typically, this value is programmed during manufacturing. This
	// address is not assignable.
	PermanentMACAddress string
	// VLAN shall contain the VLAN for this interface. If this interface supports
	// more than one VLAN, the 'VLAN' property shall not be present and the 'VLANs'
	// property shall be present instead.
	//
	// Version added: v1.3.0
	VLAN VLAN
	// VLANs shall contain a link to a resource collection of type
	// 'VLanNetworkInterfaceCollection'. If this property is used, the VLANEnabled
	// and VLAN Id property shall not be used.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.7.0
	// This property has been deprecated in favor of representing multiple VLANs as
	// 'EthernetInterface' resources.
	vLANs string
}

// UnmarshalJSON unmarshals a Ethernet object from the raw JSON.
func (e *Ethernet) UnmarshalJSON(b []byte) error {
	type temp Ethernet
	var tmp struct {
		temp
		EthernetInterfaces Link `json:"EthernetInterfaces"`
		VLANs              Link `json:"VLANs"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = Ethernet(tmp.temp)

	// Extract the links to other entities for later
	e.ethernetInterfaces = tmp.EthernetInterfaces.String()
	e.vLANs = tmp.VLANs.String()

	return nil
}

// EthernetInterfaces gets the EthernetInterfaces linked resource.
func (e *Ethernet) EthernetInterfaces(client Client) (*EthernetInterface, error) {
	if e.ethernetInterfaces == "" {
		return nil, nil
	}
	return GetObject[EthernetInterface](client, e.ethernetInterfaces)
}

// VLANs gets the VLANs collection.
func (e *Ethernet) VLANs(client Client) ([]*VLanNetworkInterface, error) {
	if e.vLANs == "" {
		return nil, nil
	}
	return GetCollectionObjects[VLanNetworkInterface](client, e.vLANs)
}

// FibreChannel shall describe the Fibre Channel capabilities, status, and
// configuration values for a network device function.
type FibreChannel struct {
	// AllowFIPVLANDiscovery shall indicate whether the FIP VLAN Discovery Protocol
	// determines the FCoE VLAN ID selected by the network device function for the
	// FCoE connection. If 'true' and the FIP VLAN discovery succeeds, the
	// 'FCoEActiveVLANId' property shall reflect the FCoE VLAN ID to use for all
	// FCoE traffic. If 'false' or if the FIP VLAN Discovery protocol fails, the
	// 'FCoELocalVLANId' shall be used for all FCoE traffic and the
	// 'FCoEActiveVLANId' shall reflect the 'FCoELocalVLANId'.
	AllowFIPVLANDiscovery bool
	// BootTargets shall contain an array of Fibre Channel boot targets configured
	// for this network device function.
	BootTargets []BootTargets
	// FCoEActiveVLANID shall contain 'null' or a VLAN ID currently being used for
	// FCoE traffic. When the FCoE link is down this value shall be 'null'. When
	// the FCoE link is up this value shall be either the 'FCoELocalVLANId'
	// property or a VLAN discovered through the FIP protocol.
	FCoEActiveVLANID *uint `json:"FCoEActiveVLANId,omitempty"`
	// FCoELocalVLANID shall contain the VLAN ID configured locally by setting this
	// property. This value shall be used for FCoE traffic to this network device
	// function during boot unless AllowFIPVLANDiscovery is 'true' and a valid FCoE
	// VLAN ID is found through the FIP VLAN Discovery Protocol.
	FCoELocalVLANID *uint `json:"FCoELocalVLANId,omitempty"`
	// FibreChannelID shall indicate the Fibre Channel ID that the switch assigns
	// for this interface.
	//
	// Version added: v1.3.0
	FibreChannelID string `json:"FibreChannelId"`
	// PermanentWWNN shall contain the permanent World Wide Node Name (WWNN) of
	// this function. Typically, this value is programmed during manufacturing.
	// This address is not assignable.
	PermanentWWNN string
	// PermanentWWPN shall contain the permanent World Wide Port Name (WWPN) of
	// this function. Typically, this value is programmed during manufacturing.
	// This address is not assignable.
	PermanentWWPN string
	// WWNN shall contain the effective current World Wide Node Name (WWNN) of this
	// function. If an assignable WWNN is not supported, this is a read-only alias
	// of the permanent WWNN.
	WWNN string
	// WWNSource shall contain the configuration source of the World Wide Name
	// (WWN) for this World Wide Node Name (WWNN) and World Wide Port Name (WWPN)
	// connection.
	WWNSource WWNSource
	// WWPN shall contain the effective current World Wide Port Name (WWPN) of this
	// function. If an assignable WWPN is not supported, this is a read-only alias
	// of the permanent WWPN.
	WWPN string
}

// HTTPBoot shall describe the HTTP and HTTPS boot capabilities, status, and
// configuration values for a network device function.
type HTTPBoot struct {
	// BootMediaURI shall contain the URI of the boot media loaded with this
	// network device function. An empty string shall indicate no boot media is
	// configured. All other values shall begin with 'http://' or 'https://'.
	//
	// Version added: v1.9.0
	BootMediaURI string
}

// InfiniBand shall describe the InfiniBand capabilities, status, and
// configuration values for a network device function.
type InfiniBand struct {
	// MTUSize The maximum transmission unit (MTU) configured for this network
	// device function.
	//
	// Version added: v1.5.0
	MTUSize *int `json:",omitempty"`
	// NodeGUID shall contain the effective current node GUID of this virtual port
	// of this network device function. If an assignable node GUID is not
	// supported, this is a read-only alias of the PermanentNodeGUID.
	//
	// Version added: v1.5.0
	NodeGUID string
	// PermanentNodeGUID shall contain the permanent node GUID of this network
	// device function. Typically, this value is programmed during manufacturing.
	// This address is not assignable.
	//
	// Version added: v1.5.0
	PermanentNodeGUID string
	// PermanentPortGUID shall contain the permanent port GUID of this network
	// device function. Typically, this value is programmed during manufacturing.
	// This address is not assignable.
	//
	// Version added: v1.5.0
	PermanentPortGUID string
	// PermanentSystemGUID shall contain the permanent system GUID of this network
	// device function. Typically, this value is programmed during manufacturing.
	// This address is not assignable.
	//
	// Version added: v1.5.0
	PermanentSystemGUID string
	// PortGUID shall contain the effective current virtual port GUID of this
	// network device function. If an assignable port GUID is not supported, this
	// is a read-only alias of the PermanentPortGUID.
	//
	// Version added: v1.5.0
	PortGUID string
	// SupportedMTUSizes shall contain an array of the maximum transmission unit
	// (MTU) sizes supported for this network device function.
	//
	// Version added: v1.5.0
	SupportedMTUSizes []*int
	// SystemGUID shall contain the effective current system GUID of this virtual
	// port of this network device function. If an assignable system GUID is not
	// supported, this is a read-only alias of the PermanentSystemGUID.
	//
	// Version added: v1.5.0
	SystemGUID string
}

// Limit shall describe a single array element of the packet and byte limits of
// a network device function.
type Limit struct {
	// BurstBytesPerSecond shall contain the maximum number of bytes per second in
	// a burst allowed for this network device function.
	//
	// Version added: v1.7.0
	BurstBytesPerSecond *int `json:",omitempty"`
	// BurstPacketsPerSecond shall contain the maximum number of packets per second
	// in a burst allowed for this network device function.
	//
	// Version added: v1.7.0
	BurstPacketsPerSecond *int `json:",omitempty"`
	// Direction shall indicate the direction of the data to which this limit
	// applies for this network device function.
	//
	// Version added: v1.7.0
	Direction DataDirection
	// SustainedBytesPerSecond shall contain the maximum number of sustained bytes
	// per second allowed for this network device function.
	//
	// Version added: v1.7.0
	SustainedBytesPerSecond *int `json:",omitempty"`
	// SustainedPacketsPerSecond shall contain the maximum number of sustained
	// packets per second allowed for this network device function.
	//
	// Version added: v1.7.0
	SustainedPacketsPerSecond *int `json:",omitempty"`
}

// ISCSIBoot shall describe the iSCSI boot capabilities, status, and
// configuration values for a network device function.
type ISCSIBoot struct {
	// AuthenticationMethod shall contain the iSCSI boot authentication method for
	// this network device function.
	AuthenticationMethod AuthenticationMethod
	// CHAPSecret shall contain the shared secret for CHAP authentication.
	CHAPSecret string
	// CHAPUsername shall contain the username for CHAP authentication.
	CHAPUsername string
	// IPAddressType shall contain the type of IP address being populated in the
	// iSCSIBoot IP address fields. Mixing IPv6 and IPv4 addresses on the same
	// network device function shall not be permissible.
	IPAddressType IPAddressType
	// IPMaskDNSViaDHCP shall indicate whether the iSCSI boot initiator uses DHCP
	// to obtain the initiator name, IP address, and netmask.
	IPMaskDNSViaDHCP bool
	// InitiatorDefaultGateway shall contain the IPv6 or IPv4 iSCSI boot default
	// gateway.
	InitiatorDefaultGateway string
	// InitiatorIPAddress shall contain the IPv6 or IPv4 address of the iSCSI boot
	// initiator.
	InitiatorIPAddress string
	// InitiatorName shall contain the iSCSI boot initiator name. This property
	// should match formats defined in RFC3720 or RFC3721.
	InitiatorName string
	// InitiatorNetmask shall contain the IPv6 or IPv4 netmask of the iSCSI boot
	// initiator.
	InitiatorNetmask string
	// MutualCHAPSecret shall contain the CHAP secret for two-way CHAP
	// authentication.
	MutualCHAPSecret string
	// MutualCHAPUsername shall contain the CHAP username for two-way CHAP
	// authentication.
	MutualCHAPUsername string
	// PrimaryDNS shall contain the IPv6 or IPv4 address of the primary DNS server
	// for the iSCSI boot initiator.
	PrimaryDNS string
	// PrimaryLUN shall contain the logical unit number (LUN) for the primary iSCSI
	// boot target.
	PrimaryLUN *int `json:",omitempty"`
	// PrimaryTargetIPAddress shall contain the IPv4 or IPv6 address for the
	// primary iSCSI boot target.
	PrimaryTargetIPAddress string
	// PrimaryTargetName shall contain the name of the primary iSCSI boot target.
	// This property should match formats defined in RFC3720 or RFC3721.
	PrimaryTargetName string
	// PrimaryTargetTCPPort shall contain the TCP port for the primary iSCSI boot
	// target.
	PrimaryTargetTCPPort *int `json:",omitempty"`
	// PrimaryVLANEnable shall indicate whether this VLAN is enabled for the
	// primary iSCSI boot target.
	PrimaryVLANEnable bool
	// PrimaryVLANID shall contain the 802.1q VLAN ID to use for iSCSI boot from
	// the primary target. This VLAN ID is only used if 'PrimaryVLANEnable' is
	// 'true'.
	PrimaryVLANID *uint `json:"PrimaryVLANId,omitempty"`
	// RouterAdvertisementEnabled shall indicate whether IPv6 router advertisement
	// is enabled for the iSCSI boot target. This setting shall apply to only IPv6
	// configurations.
	RouterAdvertisementEnabled bool
	// SecondaryDNS shall contain the IPv6 or IPv4 address of the secondary DNS
	// server for the iSCSI boot initiator.
	SecondaryDNS string
	// SecondaryLUN shall contain the logical unit number (LUN) for the secondary
	// iSCSI boot target.
	SecondaryLUN *int `json:",omitempty"`
	// SecondaryTargetIPAddress shall contain the IPv4 or IPv6 address for the
	// secondary iSCSI boot target.
	SecondaryTargetIPAddress string
	// SecondaryTargetName shall contain the name of the secondary iSCSI boot
	// target. This property should match formats defined in RFC3720 or RFC3721.
	SecondaryTargetName string
	// SecondaryTargetTCPPort shall contain the TCP port for the secondary iSCSI
	// boot target.
	SecondaryTargetTCPPort *int `json:",omitempty"`
	// SecondaryVLANEnable shall indicate whether this VLAN is enabled for the
	// secondary iSCSI boot target.
	SecondaryVLANEnable bool
	// SecondaryVLANID shall contain the 802.1q VLAN ID to use for iSCSI boot from
	// the secondary target. This VLAN ID is only used if 'SecondaryVLANEnable' is
	// 'true'.
	SecondaryVLANID *uint `json:"SecondaryVLANId,omitempty"`
	// TargetInfoViaDHCP shall indicate whether the iSCSI boot target name, LUN, IP
	// address, and netmask should be obtained from DHCP.
	TargetInfoViaDHCP bool
}
