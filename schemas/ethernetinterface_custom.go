//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// Update commits updates to this object's properties to the running system.
func (e *EthernetInterface) Update() error {
	readWriteFields := []string{
		"AutoNeg",
		"DHCPv4",
		"DHCPv6",
		"FQDN",
		"FullDuplex",
		"HostName",
		"IPv4Addresses",
		"IPv4StaticAddresses",
		"IPv6AddressPolicyTable",
		"IPv6Addresses",
		"IPv6Enabled",
		"IPv6StaticAddresses",
		"IPv6StaticDefaultGateways",
		"InterfaceEnabled",
		"MACAddress",
		"MTUSize",
		"SpeedMbps",
		"StatelessAddressAutoConfig",
		"StaticNameServers",
		"Status",
		"TeamMode",
		"VLAN",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}
