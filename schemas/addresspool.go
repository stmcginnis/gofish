//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.1 - #AddressPool.v1_3_0.AddressPool

package schemas

import (
	"encoding/json"
)

// AddressPool shall represent an address pool in a Redfish implementation.
type AddressPool struct {
	Entity
	// Ethernet shall contain the Ethernet-related properties for this address
	// pool.
	//
	// Version added: v1.1.0
	Ethernet APEthernet
	// GenZ shall contain the Gen-Z related properties for this address pool.
	GenZ APGenZ
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// zones are the URIs for Zones.
	zones []string
}

// UnmarshalJSON unmarshals a AddressPool object from the raw JSON.
func (a *AddressPool) UnmarshalJSON(b []byte) error {
	type temp AddressPool
	type aLinks struct {
		Endpoints Links `json:"Endpoints"`
		Zones     Links `json:"Zones"`
	}
	var tmp struct {
		temp
		Links aLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AddressPool(tmp.temp)

	// Extract the links to other entities for later
	a.endpoints = tmp.Links.Endpoints.ToStrings()
	a.zones = tmp.Links.Zones.ToStrings()

	return nil
}

// GetAddressPool will get a AddressPool instance from the service.
func GetAddressPool(c Client, uri string) (*AddressPool, error) {
	return GetObject[AddressPool](c, uri)
}

// ListReferencedAddressPools gets the collection of AddressPool from
// a provided reference.
func ListReferencedAddressPools(c Client, link string) ([]*AddressPool, error) {
	return GetCollectionObjects[AddressPool](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (a *AddressPool) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](a.client, a.endpoints)
}

// Zones gets the Zones linked resources.
func (a *AddressPool) Zones() ([]*Zone, error) {
	return GetObjects[Zone](a.client, a.zones)
}

// ASNumberRange shall contain the Autonomous System (AS) number range.
type ASNumberRange struct {
	// Lower shall contain the lower Autonomous System (AS) number to be used as
	// part of a range of ASN values.
	//
	// Version added: v1.1.0
	Lower uint
	// Upper shall contain the upper Autonomous System (AS) number to be used as
	// part of a range of ASN values.
	//
	// Version added: v1.1.0
	Upper uint
}

// BFDSingleHopOnly shall contain the BFD-related properties for an Ethernet
// fabric that uses Bidirectional Forwarding Detection (BFD) for link fault
// detection.
type BFDSingleHopOnly struct {
	// DemandModeEnabled shall indicate if Bidirectional Forwarding Detection (BFD)
	// Demand Mode is enabled. In Demand mode, no periodic BFD Control packets will
	// flow in either direction.
	//
	// Version added: v1.1.0
	DemandModeEnabled bool
	// DesiredMinTxIntervalMilliseconds shall contain the minimum interval, in
	// milliseconds, that the local system would like to use when transmitting
	// Bidirectional Forwarding Detection (BFD) Control packets, less any jitter
	// applied.
	//
	// Version added: v1.1.0
	DesiredMinTxIntervalMilliseconds *int `json:",omitempty"`
	// KeyChain shall contain the name of the Bidirectional Forwarding Detection
	// (BFD) Key Chain.
	//
	// Version added: v1.1.0
	KeyChain string
	// LocalMultiplier shall contain the Bidirectional Forwarding Detection (BFD)
	// multiplier value. A BFD multiplier consists of the number of consecutive BFD
	// packets that shall be missed from a BFD peer before declaring that peer
	// unavailable and informing the higher-layer protocols of the failure.
	//
	// Version added: v1.1.0
	LocalMultiplier *int `json:",omitempty"`
	// MeticulousModeEnabled shall indicate whether the keyed MD5 sequence number
	// is updated with every packet. If 'true', the keyed MD5 sequence number is
	// updated with every packet. If 'false', it is updated periodically.
	//
	// Version added: v1.1.0
	MeticulousModeEnabled bool
	// RequiredMinRxIntervalMilliseconds shall contain the Bidirectional Forwarding
	// Detection (BFD) receive value. The BFD receive value determines how
	// frequently (in milliseconds) BFD packets will be expected to be received
	// from BFD peers.
	//
	// Version added: v1.1.0
	RequiredMinRxIntervalMilliseconds *int `json:",omitempty"`
	// SourcePort shall contain the Bidirectional Forwarding Detection (BFD) source
	// port.
	//
	// Version added: v1.1.0
	SourcePort *uint `json:",omitempty"`
}

// BGPEvpn shall contain the EVPN-related properties for an Ethernet fabric that
// uses an IETF-defined Ethernet Virtual Private Network (EVPN) based control
// plane specification based on RFC7432.
type BGPEvpn struct {
	// ARPProxyEnabled shall indicate whether proxy Address Resolution Protocol
	// (ARP) is enabled.
	//
	// Version added: v1.1.0
	ARPProxyEnabled bool
	// ARPSuppressionEnabled shall indicate whether Address Resolution Protocol
	// (ARP) suppression is enabled.
	//
	// Version added: v1.3.0
	ARPSuppressionEnabled bool
	// ARPSupressionEnabled shall indicate whether Address Resolution Protocol
	// (ARP) suppression is enabled.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'ARPSuppressionEnabled'
	// property.
	ARPSupressionEnabled bool
	// AnycastGatewayIPAddress shall contain the anycast gateway IPv4 address for a
	// host subnet.
	//
	// Version added: v1.1.0
	AnycastGatewayIPAddress string
	// AnycastGatewayMACAddress shall contain the anycast gateway MAC address for a
	// host subnet.
	//
	// Version added: v1.1.0
	AnycastGatewayMACAddress string
	// ESINumberRange shall contain Ethernet Segment Identifier (ESI) number ranges
	// for allocation in supporting functions such as multihoming.
	//
	// Version added: v1.1.0
	ESINumberRange ESINumberRange
	// EVINumberRange shall contain the Ethernet Virtual Private Network (EVPN)
	// Instance number (EVI) range for EVPN-based fabrics.
	//
	// Version added: v1.1.0
	EVINumberRange EVINumberRange
	// GatewayIPAddress shall contain the Gateway IPv4 address for a host subnet.
	//
	// Version added: v1.1.0
	GatewayIPAddress string
	// GatewayIPAddressRange shall contain the IPv4 address range for gateway nodes
	// on this subnet.
	//
	// Version added: v1.2.0
	GatewayIPAddressRange GatewayIPAddressRange
	// NDPProxyEnabled shall indicate whether Network Discovery Protocol (NDP)
	// proxy is enabled.
	//
	// Version added: v1.1.0
	NDPProxyEnabled bool
	// NDPSuppressionEnabled shall indicate whether Network Discovery Protocol
	// (NDP) suppression is enabled.
	//
	// Version added: v1.3.0
	NDPSuppressionEnabled bool
	// NDPSupressionEnabled shall indicate whether Network Discovery Protocol (NDP)
	// suppression is enabled.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'NDPSuppressionEnabled'
	// property.
	NDPSupressionEnabled bool
	// RouteDistinguisherAdministratorSubfield shall contain the RFC4364-defined
	// Route Distinguisher (RD) Administrator subfield.
	//
	// Version added: v1.2.0
	RouteDistinguisherAdministratorSubfield string
	// RouteDistinguisherRange shall contain the Route Distinguisher (RD) Instance
	// number range for Ethernet Virtual Private Network (EVPN) based fabrics.
	//
	// Version added: v1.1.0
	RouteDistinguisherRange RouteDistinguisherRange
	// RouteTargetAdministratorSubfield shall contain the RFC4364-defined Route
	// Target (RT) Administrator subfield.
	//
	// Version added: v1.2.0
	RouteTargetAdministratorSubfield string
	// RouteTargetRange shall contain the Route Target (RT) Instance number range
	// for Ethernet Virtual Private Network (EVPN) based fabrics.
	//
	// Version added: v1.1.0
	RouteTargetRange RouteTargetRange
	// UnderlayMulticastEnabled shall indicate whether multicast is enabled on the
	// Ethernet fabric underlay.
	//
	// Version added: v1.1.0
	UnderlayMulticastEnabled bool
	// UnknownUnicastSuppressionEnabled shall indicate whether unknown unicast
	// packets should be suppressed.
	//
	// Version added: v1.1.0
	UnknownUnicastSuppressionEnabled bool
	// VLANIdentifierAddressRange shall contain the Virtual LAN (VLAN) tag related
	// address range applicable to this Ethernet fabric or for endpoint host
	// subnets. VLAN tags can be used for the purpose of identifying packets
	// belonging to different networks.
	//
	// Version added: v1.1.0
	VLANIdentifierAddressRange VLANIdentifierAddressRange
}

// BGPNeighbor shall contain all Border Gateway Protocol (BGP) neighbor related
// properties.
type BGPNeighbor struct {
	// Address shall contain the IPv4 address assigned to a Border Gateway Protocol
	// (BGP) neighbor.
	//
	// Version added: v1.1.0
	Address string
	// AllowOwnASEnabled shall indicate whether the Autonomous System (AS) of the
	// receiving router is permitted in a Border Gateway Protocol (BGP) update. If
	// 'true', routes should be received and processed even if the router detects
	// its own ASN in the AS-Path. If 'false', they should be dropped.
	//
	// Version added: v1.1.0
	AllowOwnASEnabled bool
	// CIDR shall contain the RFC4271-defined Classless Inter-Domain Routing (CIDR)
	// value.
	//
	// Version added: v1.2.0
	CIDR int
	// ConnectRetrySeconds shall contain the Border Gateway Protocol (BGP) Retry
	// Timer. The BGP Retry Timer allows the administrator to set the amount of
	// time in seconds between retries to establish a connection to configured
	// peers that have gone down.
	//
	// Version added: v1.1.0
	ConnectRetrySeconds *int `json:",omitempty"`
	// Enabled shall indicate whether BGP neighbor communication is enabled.
	//
	// Version added: v1.2.0
	Enabled bool
	// HoldTimeSeconds shall contain the Border Gateway Protocol (BGP) Hold Timer
	// agreed upon between peers.
	//
	// Version added: v1.1.0
	HoldTimeSeconds *int `json:",omitempty"`
	// KeepaliveIntervalSeconds shall contain the Keepalive timer in seconds. It is
	// used in conjunction with the Border Gateway Protocol (BGP) hold timer.
	//
	// Version added: v1.1.0
	KeepaliveIntervalSeconds *int `json:",omitempty"`
	// LocalAS shall contain the Autonomous System (AS) number of the local Border
	// Gateway Protocol (BGP) peer.
	//
	// Version added: v1.1.0
	LocalAS *uint `json:",omitempty"`
	// LogStateChangesEnabled shall indicate whether Border Gateway Protocol (BGP)
	// neighbor state changes are logged.
	//
	// Version added: v1.1.0
	LogStateChangesEnabled bool
	// MaxPrefix These properties are applicable to configuring Border Gateway
	// Protocol (BGP) max prefix related properties.
	//
	// Version added: v1.1.0
	MaxPrefix MaxPrefix
	// MinimumAdvertisementIntervalSeconds shall contain the minimum time between
	// Border Gateway Protocol (BGP) route advertisements in seconds.
	//
	// Version added: v1.1.0
	MinimumAdvertisementIntervalSeconds *int `json:",omitempty"`
	// PassiveModeEnabled shall indicate whether Border Gateway Protocol (BGP)
	// passive mode is enabled.
	//
	// Version added: v1.1.0
	PassiveModeEnabled bool
	// PathMTUDiscoveryEnabled shall indicate whether MTU discovery is permitted.
	//
	// Version added: v1.1.0
	PathMTUDiscoveryEnabled bool
	// PeerAS shall contain the Autonomous System (AS) number of the external
	// Border Gateway Protocol (BGP) peer.
	//
	// Version added: v1.1.0
	PeerAS *uint `json:",omitempty"`
	// ReplacePeerASEnabled shall indicate whether peer Autonomous System (AS)
	// numbers should be replaced. If 'true', private ASNs are removed and replaced
	// with the peer AS. If 'false', they remain unchanged.
	//
	// Version added: v1.1.0
	ReplacePeerASEnabled bool
	// TCPMaxSegmentSizeBytes shall contain the TCP max segment size in bytes
	// signifying the number of bytes that shall be transported in a single packet.
	//
	// Version added: v1.1.0
	TCPMaxSegmentSizeBytes *int `json:",omitempty"`
	// TreatAsWithdrawEnabled shall indicate Border Gateway Protocol (BGP) withdraw
	// status. If 'true', the UPDATE message containing the path attribute shall be
	// treated as though all contained routes had been withdrawn. If 'false', they
	// should remain.
	//
	// Version added: v1.1.0
	TreatAsWithdrawEnabled bool
}

// BGPRoute shall contain properties that are applicable to configuring Border
// Gateway Protocol (BGP) route related properties.
type BGPRoute struct {
	// AdvertiseInactiveRoutesEnabled shall indicate whether inactive routes should
	// be advertised. If 'true', advertise the best Border Gateway Protocol (BGP)
	// route that is inactive because of Interior Gateway Protocol (IGP)
	// preference. If 'false', do not use as part of BGP best path selection.
	//
	// Version added: v1.1.0
	AdvertiseInactiveRoutesEnabled bool
	// DistanceExternal shall modify the administrative distance for routes learned
	// via External BGP (eBGP).
	//
	// Version added: v1.1.0
	DistanceExternal *int `json:",omitempty"`
	// DistanceInternal shall modify the administrative distance for routes learned
	// via Internal BGP (iBGP).
	//
	// Version added: v1.1.0
	DistanceInternal *int `json:",omitempty"`
	// DistanceLocal shall modify the administrative distance for routes configured
	// on a local router.
	//
	// Version added: v1.1.0
	DistanceLocal *int `json:",omitempty"`
	// ExternalCompareRouterIDEnabled shall indicate whether external router
	// identifiers should be compared. If 'true', prefer the route that comes from
	// the Border Gateway Protocol (BGP) router with the lowest router identifier.
	// If 'false', do not use as part of BGP best path selection.
	//
	// Version added: v1.1.0
	ExternalCompareRouterIDEnabled bool `json:"ExternalCompareRouterIdEnabled"`
	// FlapDampingEnabled shall indicate whether route flap dampening should be
	// enabled.
	//
	// Version added: v1.1.0
	FlapDampingEnabled bool
	// SendDefaultRouteEnabled shall indicate whether the default route should be
	// advertised. If 'true', the default route is advertised to all Border Gateway
	// Protocol (BGP) neighbors unless specifically denied. If 'false', the default
	// route is not advertised.
	//
	// Version added: v1.1.0
	SendDefaultRouteEnabled bool
}

// CommonBGPProperties shall contain properties shared across both External and
// Internal Border Gateway Protocol (BGP) related properties.
type CommonBGPProperties struct {
	// ASNumberRange shall contain the range of Autonomous System (AS) numbers
	// assigned to each Border Gateway Protocol (BGP) peer within the fabric.
	//
	// Version added: v1.1.0
	ASNumberRange ASNumberRange
	// BGPNeighbor shall contain all Border Gateway Protocol (BGP) neighbor related
	// properties.
	//
	// Version added: v1.1.0
	BGPNeighbor BGPNeighbor
	// BGPRoute shall contain Border Gateway Protocol (BGP) route-related
	// properties.
	//
	// Version added: v1.1.0
	BGPRoute BGPRoute
	// GracefulRestart shall contain all graceful restart related properties.
	//
	// Version added: v1.1.0
	GracefulRestart GracefulRestart
	// MultiplePaths shall contain all multiple path related properties.
	//
	// Version added: v1.1.0
	MultiplePaths MultiplePaths
	// SendCommunityEnabled shall indicate whether community attributes are sent to
	// BGP neighbors.
	//
	// Version added: v1.1.0
	SendCommunityEnabled bool
}

// DHCP shall contain settings for assigning DHCP-related properties to the
// Ethernet fabric.
type DHCP struct {
	// DHCPInterfaceMTUBytes shall contain the Maximum Transmission Unit (MTU) to
	// use on this interface in bytes.
	//
	// Version added: v1.1.0
	DHCPInterfaceMTUBytes *uint `json:",omitempty"`
	// DHCPRelayEnabled shall indicate whether Dynamic Host Configuration Protocol
	// (DHCP) Relay is enabled.
	//
	// Version added: v1.1.0
	DHCPRelayEnabled bool
	// DHCPServer shall contain an array of addresses assigned to the Dynamic Host
	// Configuration Protocol (DHCP) server for this Ethernet fabric.
	//
	// Version added: v1.1.0
	DHCPServer []string
}

// EBGP shall contain the External BGP (eBGP) related properties for an Ethernet
// fabric.
type EBGP struct {
	// ASNumberRange shall contain the range of Autonomous System (AS) numbers
	// assigned to each Border Gateway Protocol (BGP) peer within the fabric.
	//
	// Version added: v1.1.0
	ASNumberRange ASNumberRange
	// AllowDuplicateASEnabled shall indicate whether duplicate Autonomous System
	// (AS) numbers are allowed. If 'true', routes with the same AS number as the
	// receiving router should be allowed. If 'false', routes should be dropped if
	// the router receives its own AS number in a Border Gateway Protocol (BGP)
	// update.
	//
	// Version added: v1.1.0
	AllowDuplicateASEnabled bool
	// AllowOverrideASEnabled shall indicate whether Autonomous System (AS) numbers
	// should be overridden. If 'true', AS number should be overridden with the AS
	// number of the sending peer. If 'false', AS number override is disabled.
	//
	// Version added: v1.1.0
	AllowOverrideASEnabled bool
	// AlwaysCompareMEDEnabled shall indicate whether neighbor Multi Exit
	// Discriminator (MED) attributes should be compared.
	//
	// Version added: v1.1.0
	AlwaysCompareMEDEnabled bool
	// BGPLocalPreference shall contain the local preference value. Highest local
	// preference value is preferred for Border Gateway Protocol (BGP) best path
	// selection.
	//
	// Version added: v1.1.0
	BGPLocalPreference *int `json:",omitempty"`
	// BGPNeighbor shall contain all Border Gateway Protocol (BGP) neighbor related
	// properties.
	//
	// Version added: v1.1.0
	BGPNeighbor BGPNeighbor
	// BGPRoute shall contain Border Gateway Protocol (BGP) route-related
	// properties.
	//
	// Version added: v1.1.0
	BGPRoute BGPRoute
	// BGPWeight shall contain the Border Gateway Protocol (BGP) weight attribute
	// value for external peers. A higher BGP weight value is preferred for BGP
	// best path selection.
	//
	// Version added: v1.1.0
	BGPWeight *int `json:",omitempty"`
	// GracefulRestart shall contain all graceful restart related properties.
	//
	// Version added: v1.1.0
	GracefulRestart GracefulRestart
	// MED shall contain the Border Gateway Protocol (BGP) Multi Exit Discriminator
	// (MED) value. A lower MED value is preferred for BGP best path selection.
	//
	// Version added: v1.1.0
	MED *int `json:",omitempty"`
	// MultihopEnabled shall indicate whether External BGP (eBGP) multihop is
	// enabled.
	//
	// Version added: v1.1.0
	MultihopEnabled bool
	// MultihopTTL shall contain the External BGP (eBGP) multihop Time to Live
	// (TTL) value.
	//
	// Version added: v1.1.0
	MultihopTTL *int `json:",omitempty"`
	// MultiplePaths shall contain all multiple path related properties.
	//
	// Version added: v1.1.0
	MultiplePaths MultiplePaths
	// SendCommunityEnabled shall indicate whether community attributes are sent to
	// BGP neighbors.
	//
	// Version added: v1.1.0
	SendCommunityEnabled bool
}

// ESINumberRange shall contain Ethernet Segment Identifier (ESI) number ranges
// for allocation in supporting functions such as multihoming.
type ESINumberRange struct {
	// Lower shall contain the lower Ethernet Segment Identifier (ESI) number to be
	// used as part of a range of ESI numbers.
	//
	// Version added: v1.1.0
	Lower int
	// Upper shall contain the upper Ethernet Segment Identifier (ESI) number to be
	// used as part of a range of ESI numbers.
	//
	// Version added: v1.1.0
	Upper int
}

// EVINumberRange shall contain the Ethernet Virtual Private Network (EVPN)
// Instance (EVI) number range for EVPN-based fabrics.
type EVINumberRange struct {
	// Lower shall contain the lower Ethernet Virtual Private Network (EVPN)
	// Instance (EVI) number to be used as part of a range of EVI numbers.
	//
	// Version added: v1.1.0
	Lower int
	// Upper shall contain the upper Ethernet Virtual Private Network (EVPN)
	// Instance (EVI) number to be used as part of a range of EVI numbers.
	//
	// Version added: v1.1.0
	Upper int
}

// APEthernet shall contain the Ethernet-related properties for an address pool.
type APEthernet struct {
	// BFDSingleHopOnly shall contain the Bidirectional Forwarding Detection (BFD)
	// related properties for this Ethernet fabric.
	//
	// Version added: v1.1.0
	BFDSingleHopOnly BFDSingleHopOnly
	// BGPEvpn shall contain the BGP Ethernet Virtual Private Network (EVPN)
	// related properties for this Ethernet fabric.
	//
	// Version added: v1.1.0
	BGPEvpn BGPEvpn
	// EBGP shall contain the External BGP (eBGP) related properties for this
	// Ethernet fabric.
	//
	// Version added: v1.1.0
	EBGP EBGP
	// IPv4 shall contain IPv4 and Virtual LAN (VLAN) addressing-related properties
	// for this Ethernet fabric.
	//
	// Version added: v1.1.0
	IPv4 IPv4
	// MultiProtocolEBGP shall contain the Multi Protocol eBGP (MP eBGP) related
	// properties for this Ethernet fabric.
	//
	// Version added: v1.1.0
	MultiProtocolEBGP EBGP
	// MultiProtocolIBGP shall contain the Multi Protocol iBGP (MP iBGP) related
	// properties for this Ethernet fabric.
	//
	// Version added: v1.1.0
	MultiProtocolIBGP CommonBGPProperties
}

// GatewayIPAddressRange shall contain the IPv4 address range for gateway nodes
// for Ethernet Virtual Private Network (EVPN) based fabrics.
type GatewayIPAddressRange struct {
	// Lower shall contain the lower IP address to be used as part of a range of
	// addresses for gateway nodes in Ethernet Virtual Private Network (EVPN) based
	// fabrics.
	//
	// Version added: v1.2.0
	Lower string
	// Upper shall contain the upper IP address to be used as part of a range of
	// addresses for gateway nodes in Ethernet Virtual Private Network (EVPN) based
	// fabrics.
	//
	// Version added: v1.2.0
	Upper string
}

// APGenZ shall contain Gen-Z related properties for an address pool.
type APGenZ struct {
	// AccessKey shall contain the Gen-Z Core Specification-defined Access Key
	// required for this address pool.
	AccessKey string
	// MaxCID shall contain the maximum value for the Gen-Z Core
	// Specification-defined Component Identifier (CID).
	MaxCID *int `json:",omitempty"`
	// MaxSID shall contain the maximum value for the Gen-Z Core
	// Specification-defined Subnet Identifier (SID).
	MaxSID *int `json:",omitempty"`
	// MinCID shall contain the minimum value for the Gen-Z Core
	// Specification-defined Component Identifier (CID).
	MinCID *int `json:",omitempty"`
	// MinSID shall contain the minimum value for the Gen-Z Core
	// Specification-defined Subnet Identifier (SID).
	MinSID *int `json:",omitempty"`
}

// GracefulRestart shall contain properties that are applicable to configuring
// Border Gateway Protocol (BGP) graceful restart related properties.
type GracefulRestart struct {
	// GracefulRestartEnabled shall indicate whether to enable Border Gateway
	// Protocol (BGP) graceful restart features.
	//
	// Version added: v1.1.0
	GracefulRestartEnabled bool
	// HelperModeEnabled shall indicate what to do with stale routes. If 'true',
	// the router continues to be forward packets to stale routes. If 'false', it
	// does not forward packets to stale routes.
	//
	// Version added: v1.1.0
	HelperModeEnabled bool
	// StaleRoutesTimeSeconds shall contain the time in seconds to hold stale
	// routes for a restarting peer.
	//
	// Version added: v1.1.0
	StaleRoutesTimeSeconds *int `json:",omitempty"`
	// TimeSeconds shall contain the time in seconds to wait for a graceful restart
	// capable neighbor to re-establish Border Gateway Protocol (BGP) peering.
	//
	// Version added: v1.1.0
	TimeSeconds *int `json:",omitempty"`
}

// IPv4 shall contain IPv4 and Virtual LAN (VLAN) addressing-related properties
// for an Ethernet fabric.
type IPv4 struct {
	// AnycastGatewayIPAddress shall contain the anycast gateway IPv4 address for a
	// host subnet.
	//
	// Version added: v1.1.0
	AnycastGatewayIPAddress string
	// AnycastGatewayMACAddress shall contain the anycast gateway MAC address for a
	// host subnet.
	//
	// Version added: v1.1.0
	AnycastGatewayMACAddress string
	// DHCP shall contain the primary and secondary Dynamic Host Configuration
	// Protocol (DHCP) server addressing for this Ethernet fabric.
	//
	// Version added: v1.1.0
	DHCP DHCP
	// DNSDomainName shall contain the Domain Name Service (DNS) domain name for
	// this Ethernet fabric.
	//
	// Version added: v1.1.0
	DNSDomainName string
	// DNSServer shall contain an array of the Domain Name Service (DNS) servers
	// for this Ethernet fabric.
	//
	// Version added: v1.1.0
	DNSServer []string
	// DistributeIntoUnderlayEnabled shall indicate whether host subnets are
	// distributed into the fabric underlay.
	//
	// Version added: v1.1.0
	DistributeIntoUnderlayEnabled bool
	// EBGPAddressRange shall contain the range of IPv4 addresses assigned to
	// External BGP (eBGP) neighbors belonging to different ASes (Autonomous
	// Systems).
	//
	// Version added: v1.1.0
	EBGPAddressRange IPv4AddressRange
	// FabricLinkAddressRange shall contain link-related IPv4 addressing for this
	// Ethernet fabric typically applied to connections between spine and leaf
	// Ethernet switches.
	//
	// Version added: v1.1.0
	FabricLinkAddressRange IPv4AddressRange
	// GatewayIPAddress shall contain the gateway IPv4 address for a host subnet.
	//
	// Version added: v1.1.0
	GatewayIPAddress string
	// HostAddressRange shall contain the IP subnet range for host addressing for
	// physical device endpoints that connect to this Ethernet fabric. An endpoint
	// shall be allocated an IP address from this host address range. The Ethernet
	// fabric should provide IP unicast or multicast connectivity for host device
	// endpoints belonging to this host address range.
	//
	// Version added: v1.1.0
	HostAddressRange IPv4AddressRange
	// IBGPAddressRange shall contain the range of IPv4 addresses assigned to
	// Internal BGP (iBGP) neighbors belonging to the same AS (Autonomous System).
	//
	// Version added: v1.1.0
	IBGPAddressRange IPv4AddressRange
	// LoopbackAddressRange shall contain the range of loopback-related IPv4
	// addresses assigned to this Ethernet fabric's Ethernet switches. A loopback
	// interface provides a stable interface to which an IP address is then
	// assigned. This address can be configured as the source address when the
	// networking device needs to send data for control-plane protocols such as BGP
	// and OSPF.
	//
	// Version added: v1.1.0
	LoopbackAddressRange IPv4AddressRange
	// ManagementAddressRange shall contain the range of management IPv4 addresses
	// assigned to manage this Ethernet fabric's Ethernet switches.
	//
	// Version added: v1.1.0
	ManagementAddressRange IPv4AddressRange
	// NTPOffsetHoursMinutes shall contain the Network Time Protocol (NTP) offset.
	// The NTP offset property is used to calculate the time from UTC (Universal
	// Time Coordinated) time in hours and minutes.
	//
	// Version added: v1.1.0
	NTPOffsetHoursMinutes *int `json:",omitempty"`
	// NTPServer shall contain an array of the Network Time Protocol (NTP) servers
	// for this Ethernet fabric.
	//
	// Version added: v1.1.0
	NTPServer []string
	// NTPTimezone shall contain the Network Time Protocol (NTP) time zone name
	// assigned to this Ethernet fabric.
	//
	// Version added: v1.1.0
	NTPTimezone string
	// NativeVLAN shall contain the Virtual LAN (VLAN) ID value for untagged
	// traffic.
	//
	// Version added: v1.1.0
	NativeVLAN *uint `json:",omitempty"`
	// SystemMACRange shall contain the Media Access Control (MAC) address range
	// for systems in Ethernet Virtual Private Network (EVPN) based fabrics.
	//
	// Version added: v1.2.0
	SystemMACRange SystemMACRange
	// VLANIdentifierAddressRange shall contain Virtual LAN (VLAN) tags for the
	// entire fabric as well as to end hosts.
	//
	// Version added: v1.1.0
	VLANIdentifierAddressRange VLANIdentifierAddressRange
}

// IPv4AddressRange shall contain an IPv4-related address range for an Ethernet
// fabric.
type IPv4AddressRange struct {
	// Lower shall contain the lower IPv4 network address to be used as part of a
	// subnet.
	//
	// Version added: v1.1.0
	Lower string
	// Upper shall contain the upper IPv4 network address to be used as part of a
	// host subnet.
	//
	// Version added: v1.1.0
	Upper string
}

// MaxPrefix shall contain properties that are applicable to configuring Border
// Gateway Protocol (BGP) max prefix related properties.
type MaxPrefix struct {
	// MaxPrefixNumber shall contain the maximum number of prefixes allowed from
	// the neighbor.
	//
	// Version added: v1.1.0
	MaxPrefixNumber *int `json:",omitempty"`
	// RestartTimerSeconds This property determines how long peer routers will wait
	// to delete stale routes before a Border Gateway Protocol (BGP) open message
	// is received. This timer should be less than the BGP HoldTimeSeconds
	// property.
	//
	// Version added: v1.1.0
	RestartTimerSeconds *int `json:",omitempty"`
	// ShutdownThresholdPercentage shall contain the percentage of the maximum
	// prefix received value, '1' to '100', at which the router starts to generate
	// a warning message.
	//
	// Version added: v1.1.0
	ShutdownThresholdPercentage *float64 `json:",omitempty"`
	// ThresholdWarningOnlyEnabled shall indicate what action to take if the Border
	// Gateway Protocol (BGP) route threshold is reached. If 'true', when the
	// Maximum-Prefix limit is exceeded, a log message is generated. If 'false',
	// when the Maximum-Prefix limit is exceeded, the peer session is terminated.
	//
	// Version added: v1.1.0
	ThresholdWarningOnlyEnabled bool
}

// MultiplePaths shall contain properties that are applicable to configuring
// Border Gateway Protocol (BGP) multiple path related properties.
type MultiplePaths struct {
	// MaximumPaths shall contain the maximum number of paths for multiple path
	// operation.
	//
	// Version added: v1.1.0
	MaximumPaths *int `json:",omitempty"`
	// UseMultiplePathsEnabled shall indicate whether multiple paths should be
	// advertised. If 'true', Border Gateway Protocol (BGP) advertises multiple
	// paths for the same prefix for path diversity. If 'false', it advertises
	// based on best path selection.
	//
	// Version added: v1.1.0
	UseMultiplePathsEnabled bool
}

// RouteDistinguisherRange shall contain the Route Distinguisher (RD) Instance
// number range for Ethernet Virtual Private Network (EVPN) based fabrics.
type RouteDistinguisherRange struct {
	// Lower shall contain the lower Route Distinguisher (RD) number to be used as
	// part of a range of Route Distinguisher values.
	//
	// Version added: v1.1.0
	Lower int
	// Upper shall contain the upper Route Distinguisher (RD) number to be used as
	// part of a range of Route Distinguisher values.
	//
	// Version added: v1.1.0
	Upper int
}

// RouteTargetRange shall contain the Route Target (RT) Instance number range
// for Ethernet Virtual Private Network (EVPN) based fabrics.
type RouteTargetRange struct {
	// Lower shall contain the lower Route Target (RT) number to be used as part of
	// a range of Route Target values.
	//
	// Version added: v1.1.0
	Lower *int `json:",omitempty"`
	// Upper shall contain the upper Route Target (RT) number to be used as part of
	// a range of Route Target values.
	//
	// Version added: v1.1.0
	Upper *int `json:",omitempty"`
}

// SystemMACRange shall contain the Media Access Control (MAC) address range for
// Ethernet Virtual Private Network (EVPN) based fabrics.
type SystemMACRange struct {
	// Lower shall contain the lower system Media Access Control (MAC) address to
	// be used as part of a range of system MAC addresses.
	//
	// Version added: v1.2.0
	Lower string
	// Upper shall contain the upper system Media Access Control (MAC) address to
	// be used as part of a range of system MAC addresses.
	//
	// Version added: v1.2.0
	Upper string
}

// VLANIdentifierAddressRange shall contain settings for assigning Virtual LAN
// (VLAN) tags for the entire fabric as well as for end hosts.
type VLANIdentifierAddressRange struct {
	// Lower shall contain the Virtual LAN (VLAN) tag lower value.
	//
	// Version added: v1.1.0
	Lower *uint `json:",omitempty"`
	// Upper shall contain the Virtual LAN (VLAN) tag upper value.
	//
	// Version added: v1.1.0
	Upper *uint `json:",omitempty"`
}
