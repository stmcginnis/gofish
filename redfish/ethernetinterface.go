//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DHCPv6OperatingMode is the IPv6 DHCP mode.
type DHCPv6OperatingMode string

const (
	// StatefulDHCPv6OperatingMode shall operate in stateful mode on this
	// interface. DHCPv6 stateful mode is used to configure addresses, and
	// when it is enabled, stateless mode is also implicitly enabled.
	StatefulDHCPv6OperatingMode DHCPv6OperatingMode = "Stateful"
	// StatelessDHCPv6OperatingMode shall operate in  stateless mode on this
	// interface.  DHCPv6 stateless mode allows configuring the interface
	// using DHCP options but does not configure addresses. It is always
	// enabled by default whenever DHCPv6 Stateful mode is also enabled.
	StatelessDHCPv6OperatingMode DHCPv6OperatingMode = "Stateless"
	// DisabledDHCPv6OperatingMode shall be disabled for this interface.
	DisabledDHCPv6OperatingMode DHCPv6OperatingMode = "Disabled"
)

// EthernetDeviceType is the device type.
type EthernetDeviceType string

const (

	// PhysicalEthernetDeviceType shall indicate a physical traditional
	// network interface.
	PhysicalEthernetDeviceType EthernetDeviceType = "Physical"
	// VirtualEthernetDeviceType shall indicate a network device function has
	// multiple VLANs and is representing one of them as a virtual Ethernet
	// interface.  The NetworkDeviceFunction property within Links shall
	// contain the locator for the parent network device function.
	VirtualEthernetDeviceType EthernetDeviceType = "Virtual"
)

// LinkStatus is the interface link status.
type LinkStatus string

const (
	// LinkUpLinkStatus The link is available for communication on this
	// interface.
	LinkUpLinkStatus LinkStatus = "LinkUp"
	// NoLinkLinkStatus There is no link or connection detected on this
	// interface.
	NoLinkLinkStatus LinkStatus = "NoLink"
	// LinkDownLinkStatus There is no link on this interface, but the
	// interface is connected.
	LinkDownLinkStatus LinkStatus = "LinkDown"
)

type RoutingScope string

const (
	// ExternalRoutingScope shall indicate this interface is externally accessible as if it were physically on the same
	// network as the hosting system.
	ExternalRoutingScope RoutingScope = "External"
	// HostOnlyRoutingScope shall indicate this interface is only accessible to a dedicated interface on the hosting
	// system.
	HostOnlyRoutingScope RoutingScope = "HostOnly"
	// InternalRoutingScope shall indicate this interface is only accessible to internal networking on the host, such
	// as when virtual machines or containers are allowed to communicate with each other on the same host system as
	// well as a dedicated interface on the hosting system.
	InternalRoutingScope RoutingScope = "Internal"
	// LimitedRoutingScope shall indicate this interface is accessible through IP translation provided by the hosting
	// system, such as a NAT (network address translation).
	LimitedRoutingScope RoutingScope = "Limited"
)

type TeamMode string

const (
	// NoneTeamMode No teaming.
	NoneTeamMode TeamMode = "None"
	// RoundRobinTeamMode Packets are transmitted in sequential order from the teamed interfaces.
	RoundRobinTeamMode TeamMode = "RoundRobin"
	// ActiveBackupTeamMode One interface in the team is active and the others are kept in standby until a failure
	// occurs.
	ActiveBackupTeamMode TeamMode = "ActiveBackup"
	// XORTeamMode Transmitting is determined based upon a hash policy.
	XORTeamMode TeamMode = "XOR"
	// BroadcastTeamMode Packets are transmitted on all interfaces in the team.
	BroadcastTeamMode TeamMode = "Broadcast"
	// IEEE8023adTeamMode The interfaces in the team create an IEEE802.3ad link aggregation group.
	IEEE8023adTeamMode TeamMode = "IEEE802_3ad"
	// AdaptiveTransmitLoadBalancingTeamMode Packets are transmitted based upon the current load of each interface in
	// the team.
	AdaptiveTransmitLoadBalancingTeamMode TeamMode = "AdaptiveTransmitLoadBalancing"
	// AdaptiveLoadBalancingTeamMode Packets are transmitted and received based upon the current load of each interface
	// in the team.
	AdaptiveLoadBalancingTeamMode TeamMode = "AdaptiveLoadBalancing"
)

// DHCPv4Configuration describes the configuration of DHCP v4.
type DHCPv4Configuration struct {
	// DHCPEnabled shall indicate whether DHCP v4 is enabled for this
	// EthernetInterface.
	DHCPEnabled bool
	// UseDNSServers shall indicate whether the interface will use
	// DHCPv4-supplied DNS servers.
	UseDNSServers bool
	// UseDomainName shall indicate whether the interface will use a
	// DHCPv4-supplied domain name.
	UseDomainName bool
	// UseGateway shall indicate whether the interface will use a
	// DHCPv4-supplied gateway.
	UseGateway bool
	// UseNTPServers shall indicate whether the interface will use
	// DHCPv4-supplied NTP servers.
	UseNTPServers bool
	// UseStaticRoutes shall indicate whether the interface will use a
	// DHCPv4-supplied static routes.
	UseStaticRoutes bool
}

// DHCPv6Configuration describes the configuration of DHCP v6.
type DHCPv6Configuration struct {
	// OperatingMode is used to configure addresses, and when it is enabled,
	// stateless mode is also implicitly enabled.
	OperatingMode DHCPv6OperatingMode
	// UseDNSServers shall indicate whether the interface will use
	// DHCPv6-supplied DNS servers.
	UseDNSServers bool
	// UseDomainName shall indicate whether the interface will use a domain name
	// supplied through  DHCPv6 stateless mode.
	UseDomainName bool
	// UseNTPServers shall indicate whether the interface will use
	// DHCPv6-supplied NTP servers.
	UseNTPServers bool
	// UseRapidCommit shall indicate whether the interface will use DHCPv6 rapid
	// commit mode for stateful mode address assignments.
	UseRapidCommit bool
}

// EthernetInterface is used to represent NIC resources.
type EthernetInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AutoNeg shall be true if auto negotiation of speed and duplex is enabled
	// on this interface and false if it is disabled.
	AutoNeg bool
	// DHCPv4 shall contain the configuration of DHCP v4.
	DHCPv4 DHCPv4Configuration
	// DHCPv6 shall contain the configuration of DHCP v6.
	DHCPv6 DHCPv6Configuration
	// Description provides a description of this resource.
	Description string
	// EthernetInterfaceType shall contain the type of interface.
	EthernetInterfaceType EthernetDeviceType
	// FQDN shall be the fully qualified domain name for this interface.
	FQDN string
	// FullDuplex shall represent the duplex status of the Ethernet connection
	// on this interface.
	FullDuplex bool
	// HostName shall be host name for this interface.
	HostName string
	// IPv4Addresses is used to represent the IPv4 connection characteristics
	// for this interface. It is recommended that this property be regarded as
	// read-only, with configuration of static addresses performed by
	// updating the values within IPv4StaticAddresses. Services may reject
	// updates to this array for this reason.
	IPv4Addresses []IPv4Address
	// IPv4StaticAddresses is used to represent all IPv4 static addresses
	// assigned (but not necessarily in use) to this interface. Addresses in
	// use by this interface shall also appear in the IPv4Addresses property.
	IPv4StaticAddresses []IPv4Address
	// IPv6AddressPolicyTable is used to represent the Address Selection
	// Policy Table as defined in RFC 6724.
	IPv6AddressPolicyTable []IPv6AddressPolicyEntry
	// IPv6Addresses is used to represent the IPv6 connection characteristics
	// for this interface.
	IPv6Addresses []IPv6Address
	// IPv6DefaultGateway shall be the current
	// IPv6 default gateway address that is in use on this interface.
	IPv6DefaultGateway string
	// IPv6Enabled shall indicate whether IPv6 is enabled on this interface. If this property contains 'false', the
	// interface shall not contain any assigned IPv6 addresses, shall not initiate DHCPv6 requests, and shall not send
	// or process ICMPv6 packets. If this property is not present, but this interface contains other IPv6 properties,
	// the value shall be assumed to be 'true'.
	IPv6Enabled bool
	// IPv6StaticAddresses is used to represent the IPv6 static connection
	// characteristics for this interface.
	IPv6StaticAddresses []IPv6StaticAddress
	// IPv6StaticDefaultGateways is The values in this array shall represent
	// the IPv6 static default gateway addresses for this interface.
	IPv6StaticDefaultGateways []IPv6GatewayStaticAddress
	// InterfaceEnabled shall be a boolean
	// indicating whether this interface is enabled.
	InterfaceEnabled bool
	// LinkStatus shall be the link status of this interface (port).
	LinkStatus LinkStatus
	// MACAddress shall be the effective
	// current MAC Address of this interface. If an assignable MAC address is
	// not supported, this is a read only alias of the PermanentMACAddress.
	MACAddress string
	// MTUSize shall be the size in bytes of largest Protocol Data Unit (PDU)
	// that can be passed in an Ethernet (MAC) frame on this interface.
	MTUSize int
	// MaxIPv6StaticAddresses shall indicate the number of array items supported
	// by IPv6StaticAddresses.
	MaxIPv6StaticAddresses int
	// NameServers used on this interface.
	NameServers []string
	// Oem object used on this interface.
	Oem json.RawMessage
	// PermanentMACAddress shall be the Permanent MAC Address of this interface
	// (port). This value is typically programmed during the manufacturing time.
	// This address is not assignable.
	PermanentMACAddress string
	// RoutingScope shall contain the routing scope for this interface. This property shall only be present if this
	// interface belongs to a virtual machine or container.
	RoutingScope RoutingScope
	// SpeedMbps shall be the link speed of the interface in Mbps.
	SpeedMbps int
	// StatelessAddressAutoConfig is This object shall contain the IPv4 and
	// IPv6 Stateless Address Automatic Configuration (SLAAC) properties for
	// this interface.
	StatelessAddressAutoConfig StatelessAddressAutoConfiguration
	// StaticNameServers is used when DHCP provisioning is not in enabled for
	// name server configuration. As an implementation option they may also
	// be used in addition to DHCP provided addresses, or in cases where the
	// DHCP server provides no DNS assignments.
	StaticNameServers []string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// TeamMode shall contain the team mode for this interface. If this property is not present, the value shall be
	// assumed to be 'None'.
	TeamMode TeamMode
	// UefiDevicePath shall be the UEFI device path to the device which
	// implements this interface (port).
	UefiDevicePath string
	// VLAN shall contain the VLAN for this interface. If this interface supports more than one VLAN, the VLAN property
	// shall be absent and, instead, the VLANs call should be used instead.
	VLAN VLAN
	// VLANs is a collection of VLANs and is only used if the interface supports
	// more than one VLANs, applies only if the interface supports more than one VLAN. If this property
	// is present, the VLANEnabled and VLANId properties shall not be present.
	// This property has been deprecated in favor of newer methods indicating multiple VLANs.
	vlans string

	affiliatedInterfaces []string
	// AffiliatedInterfacesCount is the number of affiliated interfaces.
	AffiliatedInterfacesCount int
	// Chassis shall be a reference to a resource of type Chassis that represent
	// the physical container associated with this Ethernet Interface.
	chassis string
	// Endpoints shall be a reference to the resources that this ethernet
	// interface is associated with and shall reference a resource of type
	// Endpoint.
	endpoints []string
	// EndpointsCount is the number of endpoints.
	EndpointsCount int
	// HostInterface is used by a host to communicate with a Manager.
	hostInterface          string
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount are the number of network device functions associated with this interface.
	NetworkDeviceFunctionsCount int
	ports                       []string
	// PortsCount is the number of ports associated with this interface.
	PortsCount        int
	relatedInterfaces []string
	// RelatedInterfacesCount is the number of related interfaces.
	RelatedInterfacesCount int

	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EthernetInterface object from the raw JSON.
func (ethernetinterface *EthernetInterface) UnmarshalJSON(b []byte) error {
	type temp EthernetInterface

	type links struct {
		// AffiliatedInterfaces shall contain an array of links to resources of type EthernetInterface that represent the
		// Ethernet interfaces that are affiliated with this interface. EthernetInterface resources referenced in this
		// property shall reference this resource in their RelatedInterfaces property.
		AffiliatedInterfaces      common.Links
		AffiliatedInterfacesCount int `json:"AffiliatedInterfaces@odata.count"`
		Chassis                   common.Link
		Endpoints                 common.Links
		EndpointsCount            int `json:"Endpoints@odata.count"`
		HostInterface             common.Link
		// NetworkDeviceFunctions shall contain an array of links to resources of type NetworkDeviceFunction.
		NetworkDeviceFunctions common.Links
		// NetworkDeviceFunctions@odata.count
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		// Ports shall contain an array of links to resources of type Port that represent the ports providing this Ethernet
		// interface. This property shall not be present if the Ethernet interface is not directly associated to a physical
		// port.
		Ports common.Links
		// Ports@odata.count
		PortsCount int `json:"Ports@odata.count"`
		// RelatedInterfaces shall contain an array of links to resources of type EthernetInterface. If TeamMode contains
		// 'None', this property shall contain one member that represents the parent interface for the VLAN. For other
		// values of TeamMode, this property shall contain the members of the team.
		RelatedInterfaces common.Links
		// RelatedInterfaces@odata.count
		RelatedInterfacesCount int `json:"RelatedInterfaces@odata.count"`
	}

	var t struct {
		temp
		Links links
		VLANs common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernetinterface = EthernetInterface(t.temp)

	// Extract the links to other entities for later
	ethernetinterface.affiliatedInterfaces = t.Links.AffiliatedInterfaces.ToStrings()
	ethernetinterface.AffiliatedInterfacesCount = t.Links.AffiliatedInterfacesCount
	ethernetinterface.chassis = t.Links.Chassis.String()
	ethernetinterface.endpoints = t.Links.Endpoints.ToStrings()
	ethernetinterface.EndpointsCount = t.Links.EndpointsCount
	ethernetinterface.hostInterface = t.Links.HostInterface.String()
	ethernetinterface.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	ethernetinterface.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	ethernetinterface.ports = t.Links.Ports.ToStrings()
	ethernetinterface.PortsCount = t.Links.PortsCount
	ethernetinterface.relatedInterfaces = t.Links.RelatedInterfaces.ToStrings()
	ethernetinterface.RelatedInterfacesCount = t.Links.RelatedInterfacesCount

	ethernetinterface.vlans = t.VLANs.String()

	// This is a read/write object, so we need to save the raw object data for later
	ethernetinterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ethernetinterface *EthernetInterface) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EthernetInterface)
	err := original.UnmarshalJSON(ethernetinterface.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AutoNeg",
		"FQDN",
		"FullDuplex",
		"HostName",
		"IPv6Enabled",
		"InterfaceEnabled",
		"MACAddress",
		"MTUSize",
		"SpeedMbps",
		"StaticNameServers",
		"TeamMode",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ethernetinterface).Elem()

	return ethernetinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEthernetInterface will get a EthernetInterface instance from the service.
func GetEthernetInterface(c common.Client, uri string) (*EthernetInterface, error) {
	var ethernetInterface EthernetInterface
	return &ethernetInterface, ethernetInterface.Get(c, uri, &ethernetInterface)
}

// ListReferencedEthernetInterfaces gets the collection of EthernetInterface from
// a provided reference.
func ListReferencedEthernetInterfaces(c common.Client, link string) ([]*EthernetInterface, error) {
	var result []*EthernetInterface
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *EthernetInterface
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		ethernetinterface, err := GetEthernetInterface(c, link)
		ch <- GetResult{Item: ethernetinterface, Link: link, Error: err}
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

// IPv6AddressPolicyEntry describes and entry in the Address Selection Policy
// Table as defined in RFC 6724.
type IPv6AddressPolicyEntry struct {
	// Label shall contain the IPv6 Label value for this table entry as defined
	// in RFC 6724 section 2.1.
	Label int
	// Precedence shall contain the IPv6 Precedence value for this table entry
	// as defined in RFC 6724 section 2.1.
	Precedence int
	// Prefix shall contain the IPv6 Address Prefix for this table entry as
	// defined in RFC 6724 section 2.1.
	Prefix string
}

// StatelessAddressAutoConfiguration describes the IPv4 and IPv6 Stateless
// Address Automatic Configuration (SLAAC) for this interface.
type StatelessAddressAutoConfiguration struct {
	// IPv4AutoConfigEnabled shall indicate whether IPv4 Stateless Address
	// Auto-Configuration (SLAAC) is enabled for this interface.
	IPv4AutoConfigEnabled bool
	// IPv6AutoConfigEnabled shall indicate whether IPv6 Stateless Address
	// Auto-Configuration (SLAAC) is enabled for this interface.
	IPv6AutoConfigEnabled bool
}

// VLAN gets the VLAN for this interface. If this interface supports more than one VLAN, the VLAN call
// will return nil and the VLANs call should be used instead.
func (ethernetinterface *EthernetInterface) VLANs() ([]*VLanNetworkInterface, error) {
	return ListReferencedVLanNetworkInterfaces(ethernetinterface.GetClient(), ethernetinterface.vlans)
}

// AffiliatedInterfaces gets any ethernet interfaces that are affiliated with this interface.
func (ethernetinterface *EthernetInterface) AffiliatedInterfaces() ([]*EthernetInterface, error) {
	var result []*EthernetInterface
	if len(ethernetinterface.affiliatedInterfaces) == 0 {
		return result, nil
	}

	collectionError := common.NewCollectionError()
	for _, uri := range ethernetinterface.affiliatedInterfaces {
		rb, err := GetEthernetInterface(ethernetinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, rb)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Chassis gets the containing chassis.
func (ethernetinterface *EthernetInterface) Chassis() (*Chassis, error) {
	if ethernetinterface.chassis == "" {
		return nil, nil
	}

	return GetChassis(ethernetinterface.GetClient(), ethernetinterface.chassis)
}

// Endpoints gets any endpoints associated with this interface.
func (ethernetinterface *EthernetInterface) Endpoints() ([]*Endpoint, error) {
	var result []*Endpoint
	if len(ethernetinterface.endpoints) == 0 {
		return result, nil
	}

	collectionError := common.NewCollectionError()
	for _, uri := range ethernetinterface.endpoints {
		rb, err := GetEndpoint(ethernetinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, rb)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// HostInterface gets the associated host interface.
func (ethernetinterface *EthernetInterface) HostInterface() (*HostInterface, error) {
	if ethernetinterface.hostInterface == "" {
		return nil, nil
	}

	return GetHostInterface(ethernetinterface.GetClient(), ethernetinterface.hostInterface)
}

// NetworkDeviceFunctions gets any device functions associated with this interface.
func (ethernetinterface *EthernetInterface) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction
	if len(ethernetinterface.networkDeviceFunctions) == 0 {
		return result, nil
	}

	collectionError := common.NewCollectionError()
	for _, uri := range ethernetinterface.networkDeviceFunctions {
		rb, err := GetNetworkDeviceFunction(ethernetinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, rb)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Ports gets any ports associated with this interface.
func (ethernetinterface *EthernetInterface) Ports() ([]*Port, error) {
	var result []*Port
	if len(ethernetinterface.ports) == 0 {
		return result, nil
	}

	collectionError := common.NewCollectionError()
	for _, uri := range ethernetinterface.ports {
		rb, err := GetPort(ethernetinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, rb)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
