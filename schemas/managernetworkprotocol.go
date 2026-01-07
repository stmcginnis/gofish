//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #ManagerNetworkProtocol.v1_12_0.ManagerNetworkProtocol

package schemas

import (
	"encoding/json"
)

type NotifyIPv6Scope string

const (
	// LinkNotifyIPv6Scope SSDP NOTIFY messages are sent to addresses in the IPv6
	// local link scope.
	LinkNotifyIPv6Scope NotifyIPv6Scope = "Link"
	// SiteNotifyIPv6Scope SSDP NOTIFY messages are sent to addresses in the IPv6
	// local site scope.
	SiteNotifyIPv6Scope NotifyIPv6Scope = "Site"
	// OrganizationNotifyIPv6Scope SSDP NOTIFY messages are sent to addresses in
	// the IPv6 local organization scope.
	OrganizationNotifyIPv6Scope NotifyIPv6Scope = "Organization"
)

type SNMPCommunityAccessMode string

const (
	// FullSNMPCommunityAccessMode shall indicate the RFC1157-defined READ-WRITE
	// access mode.
	FullSNMPCommunityAccessMode SNMPCommunityAccessMode = "Full"
	// LimitedSNMPCommunityAccessMode shall indicate the RFC1157-defined READ-ONLY
	// access mode.
	LimitedSNMPCommunityAccessMode SNMPCommunityAccessMode = "Limited"
)

// ManagerNetworkProtocol shall represent the network service settings for the
// manager.
type ManagerNetworkProtocol struct {
	Entity
	// DHCP shall contain the DHCPv4 protocol settings for the manager.
	//
	// Version added: v1.1.0
	DHCP ProtocolSetting
	// DHCPv6 shall contain the DHCPv6 protocol settings for the manager.
	//
	// Version added: v1.3.0
	DHCPv6 ProtocolSetting
	// FQDN shall contain the fully qualified domain name for the manager.
	FQDN string
	// FTP shall contain the File Transfer Protocol (FTP) settings for the manager.
	// The default 'Port' property value should be '21' for compatibility with
	// established client implementations.
	//
	// Version added: v1.10.0
	FTP ProtocolSetting
	// FTPS shall contain the File Transfer Protocol over SSL (FTPS) settings for
	// the manager. The default value should be '21' for compatibility with
	// established client implementations.
	//
	// Version added: v1.10.0
	FTPS ProtocolSetting
	// HTTP shall contain the HTTP protocol settings for the manager. The default
	// 'Port' property value should be '80' for compatibility with established
	// client implementations.
	HTTP ProtocolSetting
	// HTTPS shall contain the HTTPS/SSL protocol settings for this manager. The
	// default 'Port' property value should be '443' for compatibility with
	// established client implementations.
	HTTPS HTTPSProtocolSettings
	// HostName shall contain the host name without any domain information.
	HostName string
	// IPMI shall contain the IPMI over LAN protocol settings for the manager. The
	// default 'Port' property value should be '623' for compatibility with
	// established client implementations.
	IPMI ProtocolSetting
	// KVMIP shall contain the KVM-IP (Keyboard, Video, Mouse over IP) protocol
	// settings for the manager. If multiple systems are supported by this manager,
	// these properties, if present, apply to all instances of KVM-IP controlled by
	// this manager.
	KVMIP ProtocolSetting
	// Modbus shall contain the Modbus TCP protocol settings for the manager.
	//
	// Version added: v1.12.0
	Modbus ModbusProtocol
	// NTP shall contain the NTP protocol settings for the manager.
	//
	// Version added: v1.2.0
	NTP NTPProtocol
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Proxy shall contain the HTTP/HTTPS proxy configuration for this manager.
	//
	// Version added: v1.8.0
	Proxy ProxyConfiguration
	// RDP shall contain the Remote Desktop Protocol settings for the manager.
	//
	// Version added: v1.3.0
	RDP ProtocolSetting
	// RFB shall contain the Remote Frame Buffer protocol settings for the manager.
	//
	// Version added: v1.3.0
	RFB ProtocolSetting
	// SFTP shall contain the Secure Shell File Transfer Protocol (SFTP) protocol
	// settings for the manager. The default value should be '22' for compatibility
	// with established client implementations.
	//
	// Version added: v1.10.0
	SFTP ProtocolSetting
	// SNMP shall contain the SNMP protocol settings for this manager. The default
	// 'Port' property value should be '161' for compatibility with established
	// client implementations.
	SNMP SNMPProtocol
	// SSDP shall contain the SSDP protocol settings for this manager. Simple
	// Service Discovery Protocol (SSDP) is for network discovery of devices
	// supporting the Redfish Service. The default 'Port' property value should be
	// '1900' for compatibility with established client implementations.
	SSDP SSDProtocol
	// SSH shall contain the Secure Shell (SSH) protocol settings for the manager.
	// The default value should be '22' for compatibility with established client
	// implementations.
	SSH ProtocolSetting
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Telnet shall contain the Telnet protocol settings for this manager. The
	// default 'Port' property value should be '23' for compatibility with
	// established client implementations.
	Telnet ProtocolSetting
	// VirtualMedia shall contain the virtual media protocol settings for this
	// manager. The 'Port' property shall contain the TCP port assigned for virtual
	// media usage. If multiple systems are supported by this manager, these
	// properties, if present, apply to all instances of virtual media controlled
	// by this manager.
	VirtualMedia ProtocolSetting
	// mDNS shall contain the Multicast Domain Name Service (mDNS) settings for the
	// manager. The default 'Port' property value should be '5353' for
	// compatibility with established client implementations.
	//
	// Version added: v1.11.0
	MDNS ProtocolSetting
}

// GetManagerNetworkProtocol will get a ManagerNetworkProtocol instance from the service.
func GetManagerNetworkProtocol(c Client, uri string) (*ManagerNetworkProtocol, error) {
	return GetObject[ManagerNetworkProtocol](c, uri)
}

// ListReferencedManagerNetworkProtocols gets the collection of ManagerNetworkProtocol from
// a provided reference.
func ListReferencedManagerNetworkProtocols(c Client, link string) ([]*ManagerNetworkProtocol, error) {
	return GetCollectionObjects[ManagerNetworkProtocol](c, link)
}

// EngineID shall contain the RFC3411-defined engine ID.
type EngineID struct {
	// ArchitectureID shall contain the architecture identifier as described in
	// item 3 of the snmpEngineID syntax of RFC3411. The full RFC3411-defined
	// snmpEngineID is formed from the concatenation of the value in the
	// 'PrivateEnterpriseId' property and the value in this property. If the most
	// significant bit in 'PrivateEnterpriseId' is set to zero, this property shall
	// not be present.
	//
	// Version added: v1.6.0
	ArchitectureID string `json:"ArchitectureId"`
	// EnterpriseSpecificMethod shall contain the enterprise-specific method as
	// described in item 2 of the snmpEngineID syntax of RFC3411. The full
	// RFC3411-defined snmpEngineID is formed from the concatenation of the value
	// in the 'PrivateEnterpriseId' property and the value in this property. If the
	// most significant bit in 'PrivateEnterpriseId' is set to one, this property
	// shall not be present.
	//
	// Version added: v1.5.0
	EnterpriseSpecificMethod string
	// PrivateEnterpriseID shall contain an RFC3411-defined private enterprise ID.
	//
	// Version added: v1.5.0
	PrivateEnterpriseID string `json:"PrivateEnterpriseId"`
}

// HTTPSProtocolSettings shall describe information about a protocol setting for a
// manager.
type HTTPSProtocolSettings struct {
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection'.
	//
	// Version added: v1.4.0
	certificates string
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
}

// UnmarshalJSON unmarshals a HTTPSProtocolSettings object from the raw JSON.
func (h *HTTPSProtocolSettings) UnmarshalJSON(b []byte) error {
	type temp HTTPSProtocolSettings
	var tmp struct {
		temp
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = HTTPSProtocolSettings(tmp.temp)

	// Extract the links to other entities for later
	h.certificates = tmp.Certificates.String()

	return nil
}

// Certificates gets the Certificates collection.
func (h *HTTPSProtocolSettings) Certificates(client Client) ([]*Certificate, error) {
	if h.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, h.certificates)
}

// ModbusProtocol shall describe information about a protocol setting for a
// manager.
type ModbusProtocol struct {
	// AllowedClients shall contain all the clients allowed to access this Modbus
	// TCP server interface. If 'RestrictAccessToAllowedClients' contains 'true',
	// the service shall enforce this restriction. If
	// 'RestrictAccessToAllowedClients' contains 'false', the service shall not
	// enforce this restriction. This property may contain FQDN, IPv4 addresses,
	// IPv6 addresses, or any combination of those formats. Services may reject
	// unsupported formats.
	//
	// Version added: v1.12.0
	AllowedClients []string
	// MaximumConnectedClients shall contain the maximum number of simultaneously
	// connected clients to this Modbus TCP server.
	//
	// Version added: v1.12.0
	MaximumConnectedClients *uint `json:",omitempty"`
	// NumberOfConnectedClients shall contain the current number of connected
	// clients to this Modbus TCP server.
	//
	// Version added: v1.12.0
	NumberOfConnectedClients *uint `json:",omitempty"`
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
	// ReadOnly shall indicate if the Modbus TCP server protocol interface is read
	// only. If 'true', the Modbus TCP server service on this manager shall reject
	// or ignore Modbus requests that modify data.
	//
	// Version added: v1.12.0
	ReadOnly bool
	// RestrictAccessToAllowedClients shall indicate if access to the Modbus TCP
	// server protocol interface is restricted to allowed clients. If 'true', the
	// Modbus TCP server service on this manager shall restrict access to the list
	// of clients defined by the 'AllowedClients' property.
	//
	// Version added: v1.12.0
	RestrictAccessToAllowedClients bool
	// ServerID shall contain the Modbus Messaging On TCP/IP Implementation
	// Guide-defined 'Unit Identifier' that identifies this Modbus TCP protocol
	// server.
	//
	// Version added: v1.12.0
	ServerID *uint `json:"ServerId,omitempty"`
}

// NTPProtocol shall describe information about a protocol setting for a
// manager.
type NTPProtocol struct {
	// NTPServers shall contain all the user-supplied NTP servers for which this
	// manager is using to obtain time. 'NetworkSuppliedServers' is used for NTP
	// servers supplied by other network protocols such as DHCP.
	//
	// Version added: v1.2.0
	NTPServers []string
	// NetworkSuppliedServers shall contain the NTP servers supplied by other
	// network protocols to this manager. DHCP is an example of a protocol that can
	// supply NTP servers to this manager.
	//
	// Version added: v1.9.0
	NetworkSuppliedServers []string
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
}

// ProtocolSetting shall describe information about a protocol setting for a manager.
type ProtocolSetting struct {
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
}

// ProxyConfiguration shall contain the HTTP/HTTPS proxy configuration for a
// manager.
type ProxyConfiguration struct {
	// Enabled shall indicate if the proxy server is used for communications.
	//
	// Version added: v1.8.0
	Enabled bool
	// ExcludeAddresses shall contain a list of hostnames or IP addresses that do
	// not require a connection through the proxy server to access.
	//
	// Version added: v1.8.0
	ExcludeAddresses []string
	// Password shall contain the password for this proxy. The value shall be
	// 'null' in responses.
	//
	// Version added: v1.8.0
	Password string
	// PasswordSet shall contain 'true' if a valid value was provided for the
	// 'Password' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.8.0
	PasswordSet bool
	// ProxyAutoConfigURI shall contain the URI at which to access a proxy
	// auto-configuration (PAC) file containing one or more JavaScript functions
	// for configuring proxy usage for this manager.
	//
	// Version added: v1.8.0
	ProxyAutoConfigURI string
	// ProxyServerURI shall contain the URI of the proxy server. The value shall
	// contain the scheme for accessing the server, and shall include the port if
	// the value is not the default port for the specified scheme.
	//
	// Version added: v1.8.0
	ProxyServerURI string
	// Username shall contain the username for this proxy.
	//
	// Version added: v1.8.0
	Username string
}

// SNMPCommunity shall contain an SNMP community string used to access an SNMP
// manager.
type SNMPCommunity struct {
	// AccessMode shall contain the access/privilege level of the SNMP community
	// used to access an SNMP manager.
	//
	// Version added: v1.5.0
	AccessMode SNMPCommunityAccessMode
	// CommunityString shall contain the SNMP community string used for accessing
	// an SNMP service on this manager. If 'HideCommunityStrings' is 'true', this
	// value shall be 'null' in responses.
	//
	// Version added: v1.5.0
	CommunityString string
	// IPv4AddressRangeLower shall contain the lowest IPv4 address in the range
	// allowed to access the SNMP service using this community string. If
	// 'RestrictCommunityToIPv4AddressRange' contains 'true', the service shall
	// enforce this range. If 'RestrictCommunityToIPv4AddressRange' contains
	// 'false', the service shall not enforce this range.
	//
	// Version added: v1.10.0
	IPv4AddressRangeLower string
	// IPv4AddressRangeUpper shall contain the upper or highest IPv4 address in the
	// range allowed to access the SNMP service using this community string. If
	// 'RestrictCommunityToIPv4AddressRange' contains 'true', the service shall
	// enforce this range. If 'RestrictCommunityToIPv4AddressRange' contains
	// 'false', the service shall not enforce this range.
	//
	// Version added: v1.10.0
	IPv4AddressRangeUpper string
	// Name is the name of the resource or array element.
	//
	// Version added: v1.5.0
	Name string
	// RestrictCommunityToIPv4AddressRange shall indicate if this community is
	// restricted to accessing the service from a range of IPv4 addresses. If
	// 'true', SNMP access using this community string is restricted to the range
	// of IPv4 addresses defined by the 'IPv4AddressRangeLower' and
	// 'IPv4AddressRangeUpper' properties.
	//
	// Version added: v1.10.0
	RestrictCommunityToIPv4AddressRange bool
}

// SNMPProtocol shall describe information about a protocol setting for a
// manager.
type SNMPProtocol struct {
	// AuthenticationProtocol shall contain the SNMP authentication protocol used
	// to access this manager. When the property contains the value 'Account', the
	// SNMP settings in each manager account are used for authentication.
	//
	// Version added: v1.5.0
	AuthenticationProtocol SNMPAuthenticationProtocols
	// CommunityAccessMode shall contain the access/privilege level of the SNMP
	// community used to access an SNMP manager.
	//
	// Version added: v1.5.0
	//
	// Deprecated: v1.10.0
	// This property has been deprecated in favor of 'AccessMode' inside
	// 'CommunityStrings'.
	CommunityAccessMode SNMPCommunityAccessMode
	// CommunityStrings shall contain an array of the SNMP community strings used
	// to access an SNMP manager.
	//
	// Version added: v1.5.0
	CommunityStrings []SNMPCommunity
	// EnableSNMPv1 shall indicate if access to the SNMP service on this manager
	// using the SNMPv1 protocol is enabled.
	//
	// Version added: v1.5.0
	EnableSNMPv1 bool
	// EnableSNMPv2c shall indicate if access to the SNMP service on this manager
	// using the SNMPv2c protocol is enabled.
	//
	// Version added: v1.5.0
	EnableSNMPv2c bool
	// EnableSNMPv3 shall indicate if access to the SNMP service on this manager
	// using the SNMPv3 protocol is enabled.
	//
	// Version added: v1.5.0
	EnableSNMPv3 bool
	// EncryptionProtocol shall contain the SNMPv3 encryption protocol used to
	// access this manager, unless 'AuthenticationProtocol' contains the value
	// 'Account'.
	//
	// Version added: v1.5.0
	EncryptionProtocol SNMPEncryptionProtocols
	// EngineID shall contain the RFC3411-defined engine ID.
	//
	// Version added: v1.5.0
	EngineID EngineID `json:"EngineId"`
	// HideCommunityStrings shall indicate if the community strings should be
	// hidden in responses.
	//
	// Version added: v1.5.0
	HideCommunityStrings bool
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
	// TrapPort shall contain the port assigned to SNMP traps.
	//
	// Version added: v1.10.0
	TrapPort *uint `json:",omitempty"`
}

// SSDProtocol shall describe information about a protocol setting for a
// manager.
type SSDProtocol struct {
	// NotifyIPv6Scope shall contain the IPv6 scope for multicast NOTIFY messages.
	// The valid enumerations are a subset of the available IPv6 scope types.
	NotifyIPv6Scope NotifyIPv6Scope
	// NotifyMulticastIntervalSeconds shall contain the time interval, in seconds,
	// between transmissions of the multicast NOTIFY ALIVE message. A setting of 0
	// seconds shall disable this functionality. The recommended value is 600
	// seconds. When disabled, other NOTIFY messages are also disabled due to their
	// dependency on previously sent NOTIFY ALIVE messages.
	NotifyMulticastIntervalSeconds *uint `json:",omitempty"`
	// NotifyTTL shall contain the time-to-live hop count used for multicast NOTIFY
	// messages. The recommended value is 2.
	NotifyTTL *uint `json:",omitempty"`
	// Port shall contain the port assigned to the protocol.
	Port *uint `json:",omitempty"`
	// ProtocolEnabled shall indicate whether the protocol is enabled.
	ProtocolEnabled bool
}
