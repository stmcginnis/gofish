//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

// AddressState is the state of the address.
type AddressState string

const (
	// PreferredAddressState This address is currently within both it's valid
	// and preferred lifetimes as defined in RFC 4862.
	PreferredAddressState AddressState = "Preferred"
	// DeprecatedAddressState This address is currently within it's valid
	// lifetime, but is now outside of it's preferred lifetime as defined in
	// RFC 4862.
	DeprecatedAddressState AddressState = "Deprecated"
	// TentativeAddressState This address is currently undergoing Duplicate
	// Address Detection testing as defined in RFC 4862 section 5.4.
	TentativeAddressState AddressState = "Tentative"
	// FailedAddressState This address has failed Duplicate Address Detection
	// testing as defined in RFC 4862 section 5.4 and is not currently in
	// use.
	FailedAddressState AddressState = "Failed"
)

// IPv4AddressOrigin is the assignment origin of the address.
type IPv4AddressOrigin string

const (
	// StaticIPv4AddressOrigin A static address as configured by the user.
	StaticIPv4AddressOrigin IPv4AddressOrigin = "Static"
	// DHCPIPv4AddressOrigin Address is provided by a DHCPv4 service.
	DHCPIPv4AddressOrigin IPv4AddressOrigin = "DHCP"
	// BOOTPIPv4AddressOrigin Address is provided by a BOOTP service.
	BOOTPIPv4AddressOrigin IPv4AddressOrigin = "BOOTP"
	// IPv4LinkLocalIPv4AddressOrigin Address is valid only for this network
	// segment (link).
	IPv4LinkLocalIPv4AddressOrigin IPv4AddressOrigin = "IPv4LinkLocal"
)

// IPv6AddressOrigin is the assignment origin of the address.
type IPv6AddressOrigin string

const (
	// StaticIPv6AddressOrigin A static address as configured by the user.
	StaticIPv6AddressOrigin IPv6AddressOrigin = "Static"
	// DHCPv6IPv6AddressOrigin Address is provided by a DHCPv6 service.
	DHCPv6IPv6AddressOrigin IPv6AddressOrigin = "DHCPv6"
	// LinkLocalIPv6AddressOrigin Address is valid only for this network
	// segment (link).
	LinkLocalIPv6AddressOrigin IPv6AddressOrigin = "LinkLocal"
	// SLAACIPv6AddressOrigin Address is provided by a Stateless Address
	// AutoConfiguration (SLAAC) service.
	SLAACIPv6AddressOrigin IPv6AddressOrigin = "SLAAC"
)

// IPv4Address describes an IPv4 address assigned to an interface.
type IPv4Address struct {
	// Address shall be an IPv4 address assigned to this interface. If DHCPv4
	// is enabled on the interface, this property becomes read-only.
	Address string
	// AddressOrigin shall be the IP address origin for this network interface.
	AddressOrigin IPv4AddressOrigin
	// Gateway shall be the IPv4 default gateway address for this interface. If
	// DHCPv4 is enabled on the interface and is configured to set the IPv4
	// default gateway address, this property becomes read-only.
	Gateway string
	// SubnetMask shall be the IPv4 subnet mask for this address. If DHCPv4 is
	// enabled on the interface, this property becomes read-only.
	SubnetMask string
}

// IPv6Address describes an IPv6 address assigned to an interface.
type IPv6Address struct {
	// Address lists an IPv6 address that is currently assigned on this interface.
	Address string
	// AddressOrigin shall be the IPv6 address origin for this interface.
	AddressOrigin IPv6AddressOrigin
	// AddressState Preferred and Deprecated states follow the definitions
	// given RFC4862 Section 5.5.4. An address is in the Tentative state
	// while undergoing Duplicate Address Detection (DAD) per RFC4862 Section
	// 5.4. The Failed state indicates a Static addresses which did not pass
	// DAD. A Static address in the Failed state is not in use on the
	// network stack, and corrective action will be needed to remedy this
	// condition.
	AddressState AddressState
	// PrefixLength shall be the IPv6 address prefix length for this interface.
	PrefixLength int8
}

// IPv6GatewayStaticAddress shall represent a single IPv6 static address to be
// assigned on a network interface.
type IPv6GatewayStaticAddress struct {
	// Address provides access to a static IPv6 address that is currently
	// assigned on a network interface.
	Address string
	// PrefixLength provides the IPv6 network prefix length in bits for this address.
	PrefixLength int8
}

// IPv6StaticAddress shall represent a single IPv6 static address to be assigned
// on a network interface.
type IPv6StaticAddress struct {
	// Address provides access to a static IPv6 address that is currently
	// assigned on a network interface.
	Address string
	// PrefixLength provides the IPv6 network prefix length in bits for this address.
	PrefixLength int8
}
