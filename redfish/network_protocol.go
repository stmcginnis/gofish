//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type Protocol struct {
	// The protocol port
	Port int64
	// An indication of whether the protocol is enabled
	ProtocolEnabled bool
}

type HTTPS struct {
	Protocol

	// certificates shall be a link to HTTP certificates resource
	certificates string
}

func (https *HTTPS) UnmarshalJSON(b []byte) error {
	type temp HTTPS

	var t struct {
		temp
		Certificates common.Link
	}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	*https = HTTPS(t.temp)
	https.certificates = t.Certificates.String()
	return nil
}

type NTP struct {
	Protocol

	// Indicates to which user-supplied NTP servers this manager is
	// subscribed
	NTPServers []string
	// The NTP servers supplied by other network protocols to this
	// manager
	NetworkSuppliedServers []string
}

type SNMPAuthenticationProtocol string

const (
	// Authentication is determined by account settings
	SNMPAuthAccount SNMPAuthenticationProtocol = "Account"
	// SNMP community string authentication
	SNMPAuthCommunityString SNMPAuthenticationProtocol = "CommunityString"
	// HMAC-MD5-96 authentication
	SNMPAuthHMAC_MD5 SNMPAuthenticationProtocol = "HMAC_MD5" //nolint
	// HMAC-SHA-96 authentication
	SNMPAuthHMAC_SHA96 SNMPAuthenticationProtocol = "HMAC_SHA96" //nolint
	// HMAC-128-SHA-224 authentication
	SNMPAuthHMAC128_SHA224 SNMPAuthenticationProtocol = "HMAC128_SHA224" //nolint
	// HMAC-192-SHA-256 authentication
	SNMPAuthHMAC192_SHA256 SNMPAuthenticationProtocol = "HMAC192_SHA256" //nolint
	// HMAC-256-SHA-384 authentication
	SNMPAuthHMAC256_SHA384 SNMPAuthenticationProtocol = "HMAC256_SHA384" //nolint
	// HMAC-384-SHA-512 authentication
	SNMPAuthHMAC384_SHA512 SNMPAuthenticationProtocol = "HMAC384_SHA512" //nolint
)

type SNMPCommunityAccessMode string

const (
	// READ-WRITE access mode
	SNMPCommunityAccessFull SNMPCommunityAccessMode = "Full"
	// READ-ONLY access mode
	SNMPCommunityAccessLimited SNMPCommunityAccessMode = "Limited"
)

type SNMPCommunity struct {
	// The access level of the SNMP community
	AccessMode SNMPCommunityAccessMode
	// The SNMP community string
	CommunityString string
	// The name of the SNMP community
	Name string
}

type SNMPEncryptionProtocol string

const (
	// No encryption
	NoneEncryption SNMPEncryptionProtocol = "None"
	// Encryption is determined by account settings
	AccountEncryption SNMPEncryptionProtocol = "Account"
	// CBC-DES encryption
	CBC_DES_Encryption SNMPEncryptionProtocol = "CBC_DES" //nolint
	// CFB128-AES-128 encryption
	CFB128_AES128_Encryption SNMPEncryptionProtocol = "CFB128_AES128" //nolint
)

type EngineID struct {
	// The architecture identifier
	ArchitectureID string `json:"ArchitectureId"`
	// The enterprise specific method
	EnterpriseSpecificMethod string
	// The private enterprise ID
	PrivateEnterpriseID string `json:"PrivateEnterpriseId"`
}

type SNMP struct {
	Protocol

	// The authentication protocol used for SNMP access to this manager
	AuthenticationProtocol SNMPAuthenticationProtocol
	// The access level of the SNMP community
	CommunityAccessMode SNMPCommunityAccessMode
	// The SNMP community strings
	CommunityStrings []SNMPCommunity
	// Indicates if access via SNMPv1 is enabled
	EnableSNMPv1 bool
	// Indicates if access via SNMPv2c is enabled
	EnableSNMPv2c bool
	// Indicates if access via SNMPv3 is enabled
	EnableSNMPv3 bool
	// The encryption protocol used for SNMPv3 access to this manager
	EncryptionProtocol SNMPEncryptionProtocol
	// The engine ID
	EngineID EngineID `json:"EngineId"`
	// Indicates if the community strings should be hidden
	HideCommunityStrings bool
}

type Proxy struct {
	// Indicates if the manager uses the proxy server
	Enabled bool
	// Addresses that do not require the proxy server to access
	ExcludeAddresses []string
	// The username for the proxy.
	Username string
	// The password for the proxy.  The value is `null` in responses
	Password string
	// Indicates if the Password property is set
	PasswordSet bool
	// The URI used to access a proxy auto-configuration (PAC) file
	ProxyAutoConfigURI string
	// The URI of the proxy server, including the scheme and any non-default
	// port value
	ProxyServerURI string
}

type NotifyIPv6Scope string

const (
	// SSDP NOTIFY messages are sent to addresses in the IPv6 local link scope
	NotifyIPv6ScopeLink NotifyIPv6Scope = "Link"
	// SSDP NOTIFY messages are sent to addresses in the IPv6 local
	// organization scope
	NotifyIPv6ScopeSite NotifyIPv6Scope = "Site"
	// SSDP NOTIFY messages are sent to addresses in the IPv6 local site scope
	NotifyIPv6ScopeOrganization NotifyIPv6Scope = "Organization"
)

type SSDP struct {
	Protocol

	// The IPv6 scope for multicast NOTIFY messages for SSDP
	NotifyIPv6Scope NotifyIPv6Scope
	// The time interval, in seconds, between transmissions of the
	// multicast NOTIFY ALIVE message from this service for SSDP
	NotifyMulticastIntervalSeconds int64
	// The time-to-live hop count for SSDP multicast NOTIFY messages
	NotifyTTL int64
}

type NetworkProtocolSettings struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// The description of this resource
	Description string
	// The status and health of the Resource and its subordinate or
	// dependent Resources
	Status common.Status
	// The settings for this manager's DHCPv4 protocol support
	DHCP Protocol
	// The settings for this manager's DHCPv6 protocol support
	DHCPv6 Protocol
	// The fully qualified domain name for the manager obtained by
	// DNS including the host name and top-level domain name
	FQDN string
	// The settings for this manager's HTTP protocol support
	HTTP Protocol
	// The settings for this manager's HTTPS protocol support
	HTTPS HTTPS
	// The DNS host name of this manager, without any domain information
	HostName string
	// The settings for this manager's IPMI-over-LAN protocol support
	IPMI Protocol
	// The settings for this manager's KVM-IP protocol support that
	// apply to all system instances controlled by this manager
	KVMIP Protocol
	// The settings for this manager's NTP protocol support
	NTP NTP
	// The HTTP/HTTPS proxy information for this manager
	Proxy Proxy
	// The settings for this manager's Remote Desktop Protocol support
	RDP Protocol
	// The settings for this manager's Remote Frame Buffer protocol
	// support, which can support VNC
	RFB Protocol
	// The settings for this manager's SNMP support
	SNMP SNMP
	// The settings for this manager's SSDP support
	SSDP SSDP
	// The settings for this manager's Secure Shell (SSH) protocol
	// support
	SSH Protocol
	// The settings for this manager's Telnet protocol support
	Telnet Protocol
	// The settings for this manager's virtual media support that
	// apply to all system instances controlled by this manager
	VirtualMedia Protocol
	// The OEM extension property
	Oem json.RawMessage
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

func (networkProtocol *NetworkProtocolSettings) UnmarshalJSON(b []byte) error {
	type temp NetworkProtocolSettings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkProtocol = NetworkProtocolSettings(t.temp)

	networkProtocol.rawData = b
	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkProtocol *NetworkProtocolSettings) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkProtocolSettings)
	err := original.UnmarshalJSON(networkProtocol.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"DHCP",
		"DHCPv6",
		"HTTP",
		"HTTPS",
		"IPMI",
		"KVMIP",
		"NTP",
		"Proxy",
		"RDP",
		"RFB",
		"SNMP",
		"SSDP",
		"SSH",
		"Telnet",
		"VirtualMedia",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkProtocol).Elem()

	return networkProtocol.Entity.Update(originalElement, currentElement, readWriteFields)
}

func GetNetworkProtocol(c common.Client, uri string) (*NetworkProtocolSettings, error) {
	var networkProtocolSettings NetworkProtocolSettings
	return &networkProtocolSettings, networkProtocolSettings.Get(c, uri, &networkProtocolSettings)
}
