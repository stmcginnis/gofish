//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AuthenticationMethod is the method used for authentication.
type AuthenticationMethod string

const (
	// NoneAuthenticationMethod No iSCSI authentication is used.
	NoneAuthenticationMethod AuthenticationMethod = "None"
	// CHAPAuthenticationMethod iSCSI Challenge Handshake Authentication
	// Protocol (CHAP) authentication is used.
	CHAPAuthenticationMethod AuthenticationMethod = "CHAP"
	// MutualCHAPAuthenticationMethod iSCSI Mutual Challenge Handshake
	// Authentication Protocol (CHAP) authentication is used.
	MutualCHAPAuthenticationMethod AuthenticationMethod = "MutualCHAP"
)

// BootMode is the boot operation mode.
type BootMode string

const (
	// DisabledBootMode Do not indicate to UEFI/BIOS that this device is
	// bootable.
	DisabledBootMode BootMode = "Disabled"
	// PXEBootMode Boot this device using the embedded PXE support. Only
	// applicable if the NetworkDeviceFunctionType is set to Ethernet.
	PXEBootMode BootMode = "PXE"
	// ISCSIBootMode Boot this device using the embedded iSCSI boot support
	// and configuration. Only applicable if the NetworkDeviceFunctionType
	// is set to iSCSI.
	ISCSIBootMode BootMode = "iSCSI"
	// FibreChannelBootMode Boot this device using the embedded Fibre Channel
	// support and configuration. Only applicable if the
	// NetworkDeviceFunctionType is set to FibreChannel.
	FibreChannelBootMode BootMode = "FibreChannel"
	// FibreChannelOverEthernetBootMode Boot this device using the embedded
	// Fibre Channel over Ethernet (FCoE) boot support and configuration.
	// Only applicable if the NetworkDeviceFunctionType is set to
	// FibreChannelOverEthernet.
	FibreChannelOverEthernetBootMode BootMode = "FibreChannelOverEthernet"
	// HTTPBootMode Boot this device by using the embedded HTTP/HTTPS support. Only applicable if the NetDevFuncType is
	// 'Ethernet'.
	HTTPBootMode BootMode = "HTTP"
)

// IPAddressType is the version of IP protocol.
type IPAddressType string

const (
	// IPv4IPAddressType IPv4 addressing is used for all IP-fields in this
	// object.
	IPv4IPAddressType IPAddressType = "IPv4"
	// IPv6IPAddressType IPv6 addressing is used for all IP-fields in this
	// object.
	IPv6IPAddressType IPAddressType = "IPv6"
)

// NetworkDeviceTechnology is the technology type of the network device.
type NetworkDeviceTechnology string

const (
	// DisabledNetworkDeviceTechnology Neither enumerated nor visible to the
	// operating system.
	DisabledNetworkDeviceTechnology NetworkDeviceTechnology = "Disabled"
	// EthernetNetworkDeviceTechnology Appears to the operating system as an
	// Ethernet device.
	EthernetNetworkDeviceTechnology NetworkDeviceTechnology = "Ethernet"
	// FibreChannelNetworkDeviceTechnology Appears to the operating system as
	// a Fibre Channel device.
	FibreChannelNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannel"
	// ISCSINetworkDeviceTechnology Appears to the operating system as an
	// iSCSI device.
	ISCSINetworkDeviceTechnology NetworkDeviceTechnology = "iSCSI"
	// FibreChannelOverEthernetNetworkDeviceTechnology Appears to the
	// operating system as an FCoE device.
	FibreChannelOverEthernetNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannelOverEthernet"
	// InfiniBandNetworkDeviceTechnology Appears to the operating system as an InfiniBand device.
	InfiniBandNetworkDeviceTechnology NetworkDeviceTechnology = "InfiniBand"
)

// WWNSource is the source of the world wide name.
type WWNSource string

const (
	// ConfiguredLocallyWWNSource The set of FC/FCoE boot targets was applied
	// locally through API or UI.
	ConfiguredLocallyWWNSource WWNSource = "ConfiguredLocally"
	// ProvidedByFabricWWNSource The set of FC/FCoE boot targets was applied
	// by the Fibre Channel fabric.
	ProvidedByFabricWWNSource WWNSource = "ProvidedByFabric"
)

// BootTargets shall describe a Fibre Channel boot target configured for a
// network device function.
type BootTargets struct {
	// BootPriority shall be the relative priority for this entry in the boot
	// targets array. Lower numbers shall represent higher priority, with zero being the highest priority.
	// The BootPriority shall be unique for all entries of the BootTargets
	// array.
	BootPriority int
	// LUNID shall contain the logical unit number (LUN) ID from which to boot on the device to which the corresponding
	// WWPN refers.
	LUNID string
	// WWPN shall be World-Wide Port Name (WWPN) to boot from.
	WWPN string
}

// Ethernet shall describe the Ethernet capabilities, status, and configuration
// values for a network device function.
type Ethernet struct {
	// EthernetInterfaces shall contain a link to a collection of type EthernetInterfaceCollection that represent the
	// Ethernet interfaces present on this network device function. This property shall not be present if this network
	// device function is not referenced by a NetworkInterface resource.
	ethernetInterfaces []string
	// MACAddress shall contain the effective current MAC address of this network device function. If an assignable MAC
	// address is not supported, this is a read-only alias of the PermanentMACAddress.
	MACAddress string
	// MTUSize The maximum transmission unit (MTU) configured for this network device function. This value serves as a
	// default for the OS driver when booting. The value only takes effect on boot.
	MTUSize int
	// MTUSizeMaximum shall contain the largest maximum transmission unit (MTU) size supported for this network device
	// function.
	MTUSizeMaximum int
	// PermanentMACAddress shall contain the permanent MAC Address of this function. Typically, this value is
	// programmed during manufacturing. This address is not assignable.
	PermanentMACAddress string
	// VLAN shall contain the VLAN for this interface. If this interface supports more than one VLAN, this property is
	// not present.
	VLAN VLAN
	// VLANs is used, the VLANEnabled and VLANId property shall not be used.
	// This property has been deprecated in favor of representing multiple VLANs as EthernetInterface resources.
	vlans []string
}

// UnmarshalJSON unmarshals a Ethernet object from the raw JSON.
func (ethernet *Ethernet) UnmarshalJSON(b []byte) error {
	type temp Ethernet
	var t struct {
		temp
		EthernetInterfaces common.LinksCollection
		VLANs              common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernet = Ethernet(t.temp)

	// Extract the links to other entities for later
	ethernet.ethernetInterfaces = t.EthernetInterfaces.ToStrings()
	ethernet.vlans = t.VLANs.ToStrings()

	return nil
}

// FibreChannel shall describe the Fibre Channel capabilities, status, and
// configuration values for a network device function.
type FibreChannel struct {
	// AllowFIPVLANDiscovery is used to determine the FCoE VLAN ID selected
	// by the network device function for the FCoE connection. If true, and
	// the FIP VLAN Discovery succeeds, the FCoEActiveVLANId property shall
	// reflect the FCoE VLAN ID to be used for all FCoE traffic. If false,
	// or if the FIP VLAN Discovery protocol fails, the FCoELocalVLANId shall
	// be used for all FCoE traffic and the FCoEActiveVLANId shall reflect
	// the FCoELocalVLANId.
	AllowFIPVLANDiscovery bool
	// BootTargets shall be an array of Fibre
	// Channel boot targets configured for this network device function.
	BootTargets []BootTargets
	// FCoEActiveVLANID is used for FCoE traffic. When the FCoE link is down
	// this value shall be null. When the FCoE link is up this value shall
	// be either the FCoELocalVLANId property or a VLAN discovered via the
	// FIP protocol.
	FCoEActiveVLANID int
	// FCoELocalVLANID is used for FCoE traffic to this network device
	// function during boot unless AllowFIPVLANDiscovery is true and a valid
	// FCoE VLAN ID is found via the FIP VLAN Discovery Protocol.
	FCoELocalVLANID int
	// FibreChannelID shall indicate the Fibre Channel Id assigned by the switch
	// for this interface.
	FibreChannelID string
	// PermanentWWNN shall be the permanent World-Wide Node Name (WWNN) of this
	// network device function (physical function). This value is typically
	// programmed during the manufacturing time. This address is not assignable.
	PermanentWWNN string
	// PermanentWWPN shall be the permanent World-Wide Port Name (WWPN) of this
	// network device function (physical function). This value is typically
	// programmed during the manufacturing time. This address is not assignable.
	PermanentWWPN string
	// WWNN shall be the effective current World-Wide Node Name (WWNN) of this
	// network device function (physical function). If an assignable WWNN is not
	// supported, this is a read only alias of the PermanentWWNN.
	WWNN string
	// WWNSource shall be the configuration source of the World-Wide Names
	// (WWNs) for this connection (WWPN and WWNN).
	WWNSource WWNSource
	// WWPN shall be the effective current World-Wide Port Name (WWPN) of this
	// network device function (physical function). If an assignable WWPN is not
	// supported, this is a read only alias of the PermanentWWPN.
	WWPN string
}

// HTTPBoot shall describe the HTTP and HTTPS boot capabilities, status, and configuration values for a network
// device function.
type HTTPBoot struct {
	// BootMediaURI shall contain the URI of the boot media loaded with this network device function. An empty string
	// shall indicate no boot media is configured. All other values shall begin with 'http://' or 'https://'.
	BootMediaURI string
}

// InfiniBandNetworkDeviceFunction shall describe the InfiniBand capabilities, status, and configuration values for a network device
// function.
type InfiniBandNetworkDeviceFunction struct {
	// MTUSize The maximum transmission unit (MTU) configured for this network device function.
	MTUSize int
	// NodeGUID shall contain the effective current node GUID of this virtual port of this network device function. If
	// an assignable node GUID is not supported, this is a read-only alias of the PermanentNodeGUID.
	NodeGUID string
	// PermanentNodeGUID shall contain the permanent node GUID of this network device function. Typically, this value
	// is programmed during manufacturing. This address is not assignable.
	PermanentNodeGUID string
	// PermanentPortGUID shall contain the permanent port GUID of this network device function. Typically, this value
	// is programmed during manufacturing. This address is not assignable.
	PermanentPortGUID string
	// PermanentSystemGUID shall contain the permanent system GUID of this network device function. Typically, this
	// value is programmed during manufacturing. This address is not assignable.
	PermanentSystemGUID string
	// PortGUID shall contain the effective current virtual port GUID of this network device function. If an assignable
	// port GUID is not supported, this is a read-only alias of the PermanentPortGUID.
	PortGUID string
	// SupportedMTUSizes shall contain an array of the maximum transmission unit (MTU) sizes supported for this network
	// device function.
	SupportedMTUSizes []int
	// SystemGUID shall contain the effective current system GUID of this virtual port of this network device function.
	// If an assignable system GUID is not supported, this is a read-only alias of the PermanentSystemGUID.
	SystemGUID string
}

// Limit shall describe a single array element of the packet and byte limits of a network device function.
type Limit struct {
	// BurstBytesPerSecond shall contain the maximum number of bytes per second in a burst allowed for this network
	// device function.
	BurstBytesPerSecond int
	// BurstPacketsPerSecond shall contain the maximum number of packets per second in a burst allowed for this network
	// device function.
	BurstPacketsPerSecond int
	// Direction shall indicate the direction of the data to which this limit applies for this network device function.
	Direction DataDirection
	// SustainedBytesPerSecond shall contain the maximum number of sustained bytes per second allowed for this network
	// device function.
	SustainedBytesPerSecond int
	// SustainedPacketsPerSecond shall contain the maximum number of sustained packets per second allowed for this
	// network device function.
	SustainedPacketsPerSecond int
}

// NetworkDeviceFunction is A Network Device Function represents a
// logical interface exposed by the network adapter.
type NetworkDeviceFunction struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowDeny shall contain a link to a resource collection of type AllowDenyCollection that contains the
	// permissions for packets leaving and arriving to this network device function.
	allowDeny []string
	// AssignablePhysicalNetworkPorts shall contain an array of links to resources of type Port that are the physical
	// ports to which this network device function can be assigned.
	assignablePhysicalNetworkPorts []string
	// AssignablePhysicalNetworkPortsCount gets the number of physical
	// ports to which this network device function can be assigned.
	AssignablePhysicalNetworkPortsCount int `json:"AssignablePhysicalNetworkPorts@odata.count"`
	// AssignablePhysicalPorts shall be an array of physical port references
	// that this network device function may be assigned to.
	// This property has been deprecated in favor of the AssignablePhysicalNetworkPorts property.
	assignablePhysicalPorts []string
	// AssignablePhysicalPortsCount is the number of assignable physical ports.
	// This property has been deprecated in favor of the AssignablePhysicalNetworkPorts property.
	AssignablePhysicalPortsCount int `json:"AssignablePhysicalPorts@odata.count"`
	// BootMode shall be the boot mode configured for this network device
	// function. If the value is not Disabled", this network device function
	// shall be configured for boot using the specified technology.
	BootMode BootMode
	// Description provides a description of this resource.
	Description string
	// DeviceEnabled shall be a boolean indicating whether the network device
	// function is enabled. Disabled network device functions shall not be
	// enumerated or seen by the operating system.
	DeviceEnabled bool
	// Ethernet shall contain Ethernet capabilities, status, and configuration
	// values for this network device function.
	Ethernet Ethernet
	// FibreChannel shall contain Fibre Channel capabilities, status, and
	// configuration values for this network device function.
	FibreChannel FibreChannel
	// HTTPBoot shall contain HTTP and HTTPS boot capabilities, status, and configuration values for this network
	// device function.
	HTTPBoot HTTPBoot
	// InfiniBand shall contain InfiniBand capabilities, status, and configuration values for this network device
	// function.
	InfiniBand InfiniBandNetworkDeviceFunction
	// Limits shall contain an array of byte and packet limits for this network device function.
	Limits []Limit
	// MaxVirtualFunctions shall be the number of virtual functions (VFs) that
	// are available for this Network Device Function.
	MaxVirtualFunctions int
	// Metrics shall contain a link to a resource of type NetworkDeviceFunctionMetrics that contains the metrics
	// associated with this network function.
	metrics string
	// NetDevFuncCapabilities shall contain an array of capabilities of this
	// network device function.
	NetDevFuncCapabilities []NetworkDeviceTechnology
	// NetDevFuncType shall be the configured capability of this network device
	// function.
	NetDevFuncType NetworkDeviceTechnology
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SAVIEnabled shall indicate if the RFC7039-defined Source Address Validation Improvement (SAVI) is enabled for
	// this network device function.
	SAVIEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VirtualFunctionsEnabled shall be a boolean indicating whether Single Root
	// I/O Virtualization (SR-IOV) Virtual Functions (VFs) are enabled for this
	// Network Device Function.
	VirtualFunctionsEnabled bool
	// iSCSIBoot shall contain iSCSI boot capabilities, status, and
	// configuration values for this network device function.
	ISCSIBoot ISCSIBoot
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// Endpoints shall contain an array property who's members reference
	// resources, of type Endpoint, which are associated with this network
	// device function.
	endpoints []string
	// EndpointsCount is the number of Endpoints.
	EndpointsCount int
	// EthernetInterfaces shall contain an array of links to resources of type EthernetInterface that represent the
	// virtual interfaces that were created when one of the network device function VLANs is represented as a virtual
	// NIC for the purpose of showing the IP address associated with that VLAN.
	ethernetInterfaces []string
	// EthernetInterfacesCount is the number of virtual interfaces that were created when one of the network device
	// function VLANs is represented as a virtual NIC for the purpose of showing the IP address associated with that VLAN.
	EthernetInterfacesCount int
	// OffloadProcessors shall contain an array of links to resources of type Processor that represent the processors
	// that performs offload computation for this network function, such as with a SmartNIC. This property shall not be
	// present if OffloadSystem is present.
	offloadProcessors []string
	// OffloadProcessorsCount is the number of processors that performs offload computation for this network function, such
	// as with a SmartNIC.
	OffloadProcessorsCount int
	offloadSystem          string
	// PCIeFunction shall be a references of type PCIeFunction that represents
	// the PCI-e Function associated with this Network Device Function.
	pcieFunction string
	// Deprecated (v1.5): PhysicalPortAssignment shall be the physical port that this network
	// device function is currently assigned to. This value shall be one of the
	// AssignablePhysicalPorts array members.
	physicalPortAssignment string
	// (v1.5+) The physical port to which this network device function is currently assigned.
	physicalNetworkPortAssignment string
}

// UnmarshalJSON unmarshals a NetworkDeviceFunction object from the raw JSON.
func (networkdevicefunction *NetworkDeviceFunction) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunction
	type links struct {
		Endpoints                     common.Links
		EndpointsCount                int `json:"Endpoints@odata.count"`
		EthernetInterfaces            common.Links
		EthernetInterfacesCount       int `json:"EthernetInterfaces@odata.count"`
		OffloadProcessors             common.Links
		OffloadProcessorsCount        int `json:"OffloadProcessors@odata.count"`
		OffloadSystem                 common.Link
		PCIeFunction                  common.Link
		PhysicalPortAssignment        common.Link
		PhysicalNetworkPortAssignment common.Link
	}
	var t struct {
		temp
		Links                          links
		AllowDeny                      common.LinksCollection
		AssignablePhysicalNetworkPorts common.Links
		AssignablePhysicalPorts        common.Links
		Metrics                        common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkdevicefunction = NetworkDeviceFunction(t.temp)

	// Extract the links to other entities for later
	networkdevicefunction.allowDeny = t.AllowDeny.ToStrings()
	networkdevicefunction.assignablePhysicalNetworkPorts = t.AssignablePhysicalNetworkPorts.ToStrings()
	networkdevicefunction.assignablePhysicalPorts = t.AssignablePhysicalPorts.ToStrings()
	networkdevicefunction.metrics = t.Metrics.String()

	networkdevicefunction.endpoints = t.Links.Endpoints.ToStrings()
	networkdevicefunction.EndpointsCount = t.Links.EndpointsCount
	networkdevicefunction.ethernetInterfaces = t.Links.EthernetInterfaces.ToStrings()
	networkdevicefunction.EthernetInterfacesCount = t.Links.EthernetInterfacesCount
	networkdevicefunction.offloadProcessors = t.Links.OffloadProcessors.ToStrings()
	networkdevicefunction.OffloadProcessorsCount = t.Links.OffloadProcessorsCount
	networkdevicefunction.offloadSystem = t.Links.OffloadSystem.String()
	networkdevicefunction.pcieFunction = t.Links.PCIeFunction.String()
	networkdevicefunction.physicalPortAssignment = t.Links.PhysicalPortAssignment.String()
	networkdevicefunction.physicalNetworkPortAssignment = t.Links.PhysicalNetworkPortAssignment.String()

	// This is a read/write object, so we need to save the raw object data for later
	networkdevicefunction.rawData = b

	return nil
}

// Endpoints gets the endpoints that are associated with this network device function.
func (networkdevicefunction *NetworkDeviceFunction) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](networkdevicefunction.GetClient(), networkdevicefunction.endpoints)
}

// EthernetInterfaces gets the virtual interfaces that were created when one of the network device function VLANs is
// represented as a virtual NIC for the purpose of showing the IP address associated with that VLAN.
func (networkdevicefunction *NetworkDeviceFunction) EthernetInterfaces() ([]*EthernetInterface, error) {
	return common.GetObjects[EthernetInterface](networkdevicefunction.GetClient(), networkdevicefunction.ethernetInterfaces)
}

// OffloadProcessors gets the processors that performs offload computation for this network function, such as
// with a SmartNIC. This property shall not be present if OffloadSystem is present.
func (networkdevicefunction *NetworkDeviceFunction) OffloadProcessors() ([]*Processor, error) {
	return common.GetObjects[Processor](networkdevicefunction.GetClient(), networkdevicefunction.offloadProcessors)
}

// OffloadSystem shall contain a link to a resource of type ComputerSystem that represents the system that performs
// offload computation for this network function, such as with a SmartNIC. The SystemType property contained in the
// referenced ComputerSystem resource should contain the value 'DPU'. This property shall not be present if
// OffloadProcessors is present.
func (networkdevicefunction *NetworkDeviceFunction) OffloadSystem() (*ComputerSystem, error) {
	if networkdevicefunction.offloadSystem == "" {
		return nil, nil
	}
	return GetComputerSystem(networkdevicefunction.GetClient(), networkdevicefunction.offloadSystem)
}

// PCIeFunction gets the PCIe function associated with this network device function.
func (networkdevicefunction *NetworkDeviceFunction) PCIeFunction() (*PCIeFunction, error) {
	if networkdevicefunction.pcieFunction == "" {
		return nil, nil
	}
	return GetPCIeFunction(networkdevicefunction.GetClient(), networkdevicefunction.pcieFunction)
}

// PhysicalNetworkPortAssignment gets the physical port to which this network device function is currently assigned.
func (networkdevicefunction *NetworkDeviceFunction) PhysicalNetworkPortAssignment() (*Port, error) {
	if networkdevicefunction.physicalNetworkPortAssignment == "" {
		return nil, nil
	}
	return GetPort(networkdevicefunction.GetClient(), networkdevicefunction.physicalNetworkPortAssignment)
}

// PhysicalPortAssignment gets the physical port to which this network device function is currently assigned.
// This property has been deprecated in favor of the PhysicalNetworkPortAssignment property.
func (networkdevicefunction *NetworkDeviceFunction) PhysicalPortAssignment() (*NetworkPort, error) {
	if networkdevicefunction.physicalPortAssignment == "" {
		return nil, nil
	}
	return GetNetworkPort(networkdevicefunction.GetClient(), networkdevicefunction.physicalPortAssignment)
}

// Update commits updates to this object's properties to the running system.
func (networkdevicefunction *NetworkDeviceFunction) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkDeviceFunction)
	err := original.UnmarshalJSON(networkdevicefunction.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"BootMode",
		"DeviceEnabled",
		"NetDevFuncType",
		"SAVIEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkdevicefunction).Elem()

	return networkdevicefunction.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkDeviceFunction will get a NetworkDeviceFunction instance from the service.
func GetNetworkDeviceFunction(c common.Client, uri string) (*NetworkDeviceFunction, error) {
	return common.GetObject[NetworkDeviceFunction](c, uri)
}

// ListReferencedNetworkDeviceFunctions gets the collection of NetworkDeviceFunction from
// a provided reference.
func ListReferencedNetworkDeviceFunctions(c common.Client, link string) ([]*NetworkDeviceFunction, error) {
	return common.GetCollectionObjects[NetworkDeviceFunction](c, link)
}

// ISCSIBoot shall describe the iSCSI boot capabilities, status, and
// configuration values for a network device function.
type ISCSIBoot struct {
	// AuthenticationMethod shall be the iSCSI
	// boot authentication method for this network device function.
	AuthenticationMethod AuthenticationMethod
	// CHAPSecret shall be the shared secret
	// for CHAP authentication.
	CHAPSecret string
	// CHAPUsername shall be the username for
	// CHAP authentication.
	CHAPUsername string
	// IPAddressType shall be the type of IP
	// address (IPv6 or IPv4) being populated in the iSCSIBoot IP address
	// fields. Mixing of IPv6 and IPv4 addresses on the same network device
	// function shall not be permissible.
	IPAddressType IPAddressType
	// IPMaskDNSViaDHCP shall be a boolean
	// indicating whether the iSCSI boot initiator uses DHCP to obtain the
	// initiator name, IP address, and netmask.
	IPMaskDNSViaDHCP bool
	// InitiatorDefaultGateway shall be the
	// IPv6 or IPv4 iSCSI boot default gateway.
	InitiatorDefaultGateway string
	// InitiatorIPAddress shall be the IPv6 or
	// IPv4 address of the iSCSI boot initiator.
	InitiatorIPAddress string
	// InitiatorName shall be the iSCSI boot
	// initiator name. The value of this property should match formats
	// defined in RFC3720 or RFC3721.
	InitiatorName string
	// InitiatorNetmask shall be the IPv6 or
	// IPv4 netmask of the iSCSI boot initiator.
	InitiatorNetmask string
	// MutualCHAPSecret shall be the CHAP
	// Secret for 2-way CHAP authentication.
	MutualCHAPSecret string
	// MutualCHAPUsername shall be the CHAP
	// Username for 2-way CHAP authentication.
	MutualCHAPUsername string
	// PrimaryDNS shall be the IPv6 or IPv4
	// address of the primary DNS server for the iSCSI boot initiator.
	PrimaryDNS string
	// PrimaryLUN shall be the logical unit
	// number (LUN) for the primary iSCSI boot target.
	PrimaryLUN int
	// PrimaryTargetIPAddress shall be the IP
	// address (IPv6 or IPv4) for the primary iSCSI boot target.
	PrimaryTargetIPAddress string
	// PrimaryTargetName shall be the name of
	// the primary iSCSI boot target. The value of this property should
	// match formats defined in RFC3720 or RFC3721.
	PrimaryTargetName string
	// PrimaryTargetTCPPort shall be the TCP
	// port for the primary iSCSI boot target.
	PrimaryTargetTCPPort int
	// PrimaryVLANEnable is used to indicate if this VLAN is enabled for the
	// primary iSCSI boot target.
	PrimaryVLANEnable bool
	// PrimaryVLANID is used if PrimaryVLANEnable is true.
	PrimaryVLANID int
	// RouterAdvertisementEnabled shall be a
	// boolean indicating whether IPv6 router advertisement is enabled for
	// the iSCSI boot target. This setting shall only apply to IPv6
	// configurations.
	RouterAdvertisementEnabled bool
	// SecondaryDNS shall be the IPv6 or IPv4
	// address of the secondary DNS server for the iSCSI boot initiator.
	SecondaryDNS string
	// SecondaryLUN shall be the logical unit
	// number (LUN) for the secondary iSCSI boot target.
	SecondaryLUN int
	// SecondaryTargetIPAddress shall be the IP
	// address (IPv6 or IPv4) for the secondary iSCSI boot target.
	SecondaryTargetIPAddress string
	// SecondaryTargetName shall be the name of
	// the secondary iSCSI boot target. The value of this property should
	// match formats defined in RFC3720 or RFC3721.
	SecondaryTargetName string
	// SecondaryTargetTCPPort shall be the TCP
	// port for the secondary iSCSI boot target.
	SecondaryTargetTCPPort int
	// SecondaryVLANEnable is used to indicate if this VLAN is enabled for
	// the secondary iSCSI boot target.
	SecondaryVLANEnable bool
	// SecondaryVLANID is used if SecondaryVLANEnable is true.
	SecondaryVLANID int
	// TargetInfoViaDHCP shall be a boolean
	// indicating whether the iSCSI boot target name, LUN, IP address, and
	// netmask should be obtained from DHCP.
	TargetInfoViaDHCP bool
}
