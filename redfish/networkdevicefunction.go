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
	// LUNID shall be the Logical Unit Number
	// (LUN) ID to boot from on the device referred to by the corresponding
	// WWPN.
	LUNID string
	// WWPN shall be World-Wide Port Name
	// (WWPN) to boot from.
	WWPN string
}

// UnmarshalJSON unmarshals a BootTargets object from the raw JSON.
func (boottargets *BootTargets) UnmarshalJSON(b []byte) error {
	type temp BootTargets
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boottargets = BootTargets(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Ethernet shall describe the Ethernet capabilities, status, and configuration
// values for a network device function.
type Ethernet struct {
	// MACAddress shall be the effective
	// current MAC Address of this network device function. If an assignable
	// MAC address is not supported, this is a read only alias of the
	// PermanentMACAddress.
	MACAddress string
	// MTUSize is The Maximum Transmission Unit (MTU) configured for this
	// Network Device Function. This value serves as a default for the OS
	// driver when booting. The value only takes-effect on boot.
	MTUSize int
	// PermanentMACAddress shall be the Permanent MAC Address of this network
	// device function (physical function). This value is typically programmed
	// during the manufacturing time. This address is not assignable.
	PermanentMACAddress string
	// VLAN shall be the VLAN for this interface. If this interface supports
	// more than one VLAN, the VLAN property shall not be present and the VLANS
	// collection link shall be present instead.
	vlan string
	// VLANs is used, the VLANEnabled and VLANId property shall not be used.
	vlans string
}

// UnmarshalJSON unmarshals a Ethernet object from the raw JSON.
func (ethernet *Ethernet) UnmarshalJSON(b []byte) error {
	type temp Ethernet
	var t struct {
		temp
		VLAN  common.Link
		VLANs common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernet = Ethernet(t.temp)

	// Extract the links to other entities for later
	ethernet.vlan = string(t.VLAN)
	ethernet.vlans = string(t.VLANs)

	return nil
}

// TODO: Add functions to get VLAN information.

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
	// FCoEActiveVLANId is used for FCoE traffic. When the FCoE link is down
	// this value shall be null. When the FCoE link is up this value shall
	// be either the FCoELocalVLANId property or a VLAN discovered via the
	// FIP protocol.
	FCoEActiveVLANId int
	// FCoELocalVLANId is used for FCoE traffic to this network device
	// function during boot unless AllowFIPVLANDiscovery is true and a valid
	// FCoE VLAN ID is found via the FIP VLAN Discovery Protocol.
	FCoELocalVLANId int
	// FibreChannelID shall indicate the Fibre Channel Id assigned by the switch
	// for this interface.
	FibreChannelID string `json:"FibreChannelId"`
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

// NetworkDeviceFunction is A Network Device Function represents a
// logical interface exposed by the network adapter.
type NetworkDeviceFunction struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AssignablePhysicalPorts shall be an array of physical port references
	// that this network device function may be assigned to.
	// assignablePhysicalPorts []string
	// AssignablePhysicalPortsCount is the number of assignable physical ports.
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
	// MaxVirtualFunctions shall be the number of virtual functions (VFs) that
	// are available for this Network Device Function.
	MaxVirtualFunctions int
	// NetDevFuncCapabilities shall contain an array of capabilities of this
	// network device function.
	NetDevFuncCapabilities []NetworkDeviceTechnology
	// NetDevFuncType shall be the configured capability of this network device
	// function.
	NetDevFuncType NetworkDeviceTechnology
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VirtualFunctionsEnabled shall be a boolean indicating whether Single Root
	// I/O Virtualization (SR-IOV) Virtual Functions (VFs) are enabled for this
	// Network Device Function.
	VirtualFunctionsEnabled bool
	// iSCSIBoot shall contain iSCSI boot capabilities, status, and
	// configuration values for this network device function.
	ISCSIBoot ISCSIBoot `json:"iSCSIBoot"`
	// Endpoints shall contain an array property who's members reference
	// resources, of type Endpoint, which are associated with this network
	// device function.
	endpoints []string
	// EndpointsCount is the number of Endpoints.
	EndpointsCount int
	// PCIeFunction shall be a references of type PCIeFunction that represents
	// the PCI-e Function associated with this Network Device Function.
	pcieFunction string
	// PhysicalPortAssignment shall be the physical port that this network
	// device function is currently assigned to. This value shall be one of the
	// AssignablePhysicalPorts array members.
	physicalPortAssignment string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NetworkDeviceFunction object from the raw JSON.
func (networkdevicefunction *NetworkDeviceFunction) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunction
	type links struct {
		Endpoints              common.Links
		EndpointsCount         int `json:"Endpoints@odata.count"`
		PCIeFunction           common.Link
		PhysicalPortAssignment common.Link
	}
	var t struct {
		temp
		Links                   links
		AssignablePhysicalPorts common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkdevicefunction = NetworkDeviceFunction(t.temp)

	// Extract the links to other entities for later
	networkdevicefunction.endpoints = t.Links.Endpoints.ToStrings()
	networkdevicefunction.EndpointsCount = t.Links.EndpointsCount
	networkdevicefunction.pcieFunction = string(t.Links.PCIeFunction)
	networkdevicefunction.physicalPortAssignment = string(t.Links.PhysicalPortAssignment)

	// This is a read/write object, so we need to save the raw object data for later
	networkdevicefunction.rawData = b

	return nil
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
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkdevicefunction).Elem()

	return networkdevicefunction.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkDeviceFunction will get a NetworkDeviceFunction instance from the service.
func GetNetworkDeviceFunction(c common.Client, uri string) (*NetworkDeviceFunction, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkdevicefunction NetworkDeviceFunction
	err = json.NewDecoder(resp.Body).Decode(&networkdevicefunction)
	if err != nil {
		return nil, err
	}

	networkdevicefunction.SetClient(c)
	return &networkdevicefunction, nil
}

// ListReferencedNetworkDeviceFunctions gets the collection of NetworkDeviceFunction from
// a provided reference.
func ListReferencedNetworkDeviceFunctions(c common.Client, link string) ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, networkdevicefunctionLink := range links.ItemLinks {
		networkdevicefunction, err := GetNetworkDeviceFunction(c, networkdevicefunctionLink)
		if err != nil {
			return result, err
		}
		result = append(result, networkdevicefunction)
	}

	return result, nil
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
	// PrimaryVLANId is used if PrimaryVLANEnable is true.
	PrimaryVLANId int
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
	// SecondaryVLANId is used if SecondaryVLANEnable is true.
	SecondaryVLANId int
	// TargetInfoViaDHCP shall be a boolean
	// indicating whether the iSCSI boot target name, LUN, IP address, and
	// netmask should be obtained from DHCP.
	TargetInfoViaDHCP bool
}
