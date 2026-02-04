//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/IPAddresses.v1_1_5.json
// 2018.2 - #IPAddresses.v1_1_5

package schemas

import (
	"encoding/json"
)

type AddressState string

const (
	// PreferredAddressState This address is currently within both its
	// RFC4862-defined valid and preferred lifetimes.
	PreferredAddressState AddressState = "Preferred"
	// DeprecatedAddressState This address is currently within its valid lifetime
	// but is now outside its RFC4862-defined preferred lifetime.
	DeprecatedAddressState AddressState = "Deprecated"
	// TentativeAddressState This address is currently undergoing Duplicate Address
	// Detection (DAD) testing, as defined in RFC4862, section 5.4.
	TentativeAddressState AddressState = "Tentative"
	// FailedAddressState This address has failed Duplicate Address Detection (DAD)
	// testing, as defined in RFC4862, section 5.4, and is not currently in use.
	FailedAddressState AddressState = "Failed"
)

type IPv4AddressOrigin string

const (
	// StaticIPv4AddressOrigin is a user-configured static address.
	StaticIPv4AddressOrigin IPv4AddressOrigin = "Static"
	// DHCPIPv4AddressOrigin is a DHCPv4 service-provided address.
	DHCPIPv4AddressOrigin IPv4AddressOrigin = "DHCP"
	// BOOTPIPv4AddressOrigin is a BOOTP service-provided address.
	BOOTPIPv4AddressOrigin IPv4AddressOrigin = "BOOTP"
	// IPv4LinkLocalIPv4AddressOrigin The address is valid for only this network
	// segment, or link.
	IPv4LinkLocalIPv4AddressOrigin IPv4AddressOrigin = "IPv4LinkLocal"
)

type IPv6AddressOrigin string

const (
	// StaticIPv6AddressOrigin is a static user-configured address.
	StaticIPv6AddressOrigin IPv6AddressOrigin = "Static"
	// DHCPv6IPv6AddressOrigin is a DHCPv6 service-provided address.
	DHCPv6IPv6AddressOrigin IPv6AddressOrigin = "DHCPv6"
	// LinkLocalIPv6AddressOrigin The address is valid for only this network
	// segment, or link.
	LinkLocalIPv6AddressOrigin IPv6AddressOrigin = "LinkLocal"
	// SLAACIPv6AddressOrigin is a stateless autoconfiguration (SLAAC)
	// service-provided address.
	SLAACIPv6AddressOrigin IPv6AddressOrigin = "SLAAC"
)

// IPv4Address shall describe an IPv4 address assigned to an interface.
type IPv4Address struct {
	// Address shall contain an IPv4 address assigned to this interface. If DHCPv4
	// is enabled on the interface, this property becomes read-only.
	Address string
	// AddressOrigin shall contain the IP address origin for this network
	// interface.
	AddressOrigin IPv4AddressOrigin
	// Gateway shall contain the IPv4 default gateway address for this interface.
	// If DHCPv4 is enabled on the interface and is configured to set the IPv4
	// default gateway address, this property becomes read-only. If multiple IPv4
	// addresses are present on the same interface, only a single default gateway
	// is allowed. Any additional IPv4 addresses shall not have a default gateway
	// specified.
	Gateway string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SubnetMask shall contain the IPv4 subnet mask for this address. If DHCPv4 is
	// enabled on the interface, this property becomes read-only.
	SubnetMask string
}

// IPv6Address shall describe an IPv6 address assigned to an interface.
type IPv6Address struct {
	// Address This property lists an IPv6 address that is currently assigned on
	// this interface.
	Address string
	// AddressOrigin shall contain the IPv6 address origin for this interface.
	AddressOrigin IPv6AddressOrigin
	// AddressState shall contain the current RFC4862-defined state of this
	// address. Preferred and Deprecated states follow the definitions in RFC4862,
	// section 5.5.4. The Tentative state indicates that the address is undergoing
	// Duplicate Address Detection (DAD), as defined in RFC4862, section 5.4. The
	// Failed state indicates a static address that did not pass DAD. A static
	// address in the Failed state is not in use on the network stack, and
	// corrective action is required to remedy this condition.
	AddressState AddressState
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PrefixLength shall contain the IPv6 address prefix length for this
	// interface.
	PrefixLength uint8
}

// IPv6GatewayStaticAddress shall represent a single IPv6 static address to be
// assigned on a network interface.
type IPv6GatewayStaticAddress struct {
	// Address This property provides access to a static IPv6 address that is
	// currently assigned on a network interface.
	//
	// Version added: v1.1.0
	Address string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.1.0
	OEM json.RawMessage `json:"Oem"`
	// PrefixLength Provides the IPv6 network prefix length, in bits, for this
	// address.
	//
	// Version added: v1.1.0
	PrefixLength uint8
}

// IPv6StaticAddress shall represent a single IPv6 static address to be assigned
// on a network interface.
type IPv6StaticAddress struct {
	// Address This property provides access to a static IPv6 address that is
	// currently assigned on a network interface.
	Address string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PrefixLength shall contain the IPv6 network prefix length, in bits, for this
	// address.
	PrefixLength uint8
}
