//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.1c - #DataSecurityLoSCapabilities.v1_2_0.DataSecurityLoSCapabilities

package schemas

import (
	"encoding/json"
)

// AntiVirusScanTrigger is The enumberation literals shall specify types of
// antivirus scan triggers.
type AntiVirusScanTrigger string

const (
	// NoneAntiVirusScanTrigger This enumeration literal specifies No trigger.
	NoneAntiVirusScanTrigger AntiVirusScanTrigger = "None"
	// OnFirstReadAntiVirusScanTrigger This enumeration literal specifies to
	// trigger on first read.
	OnFirstReadAntiVirusScanTrigger AntiVirusScanTrigger = "OnFirstRead"
	// OnPatternUpdateAntiVirusScanTrigger This enumeration literal specifies to
	// trigger on antivirus pattern file update.
	OnPatternUpdateAntiVirusScanTrigger AntiVirusScanTrigger = "OnPatternUpdate"
	// OnUpdateAntiVirusScanTrigger This enumeration literal specifies to trigger
	// on object update.
	OnUpdateAntiVirusScanTrigger AntiVirusScanTrigger = "OnUpdate"
	// OnRenameAntiVirusScanTrigger This enumeration literal specifies to trigger
	// on object rename.
	OnRenameAntiVirusScanTrigger AntiVirusScanTrigger = "OnRename"
)

// SFAuthenticationType is The enumeration literals shall specify authentication
// algorithms.
type SFAuthenticationType string

const (
	// NoneSFAuthenticationType This enumeration literal specifies No authentication.
	NoneSFAuthenticationType SFAuthenticationType = "None"
	// PKISFAuthenticationType This enumeration literal specifies a Public Key
	// Infrastructure. Customers with the highest assurance requirements roll PKI
	// out to hosts and users (it is more common for hosts than users. User
	// PKI-based authentication has significant operational complications and
	// administrative overheads, e.g., smart cards may be involved.
	PKISFAuthenticationType SFAuthenticationType = "PKI"
	// TicketSFAuthenticationType This enumeration literal specifies Ticket-based
	// (e.g., Kerberos): This is the most common class of authentication
	// infrastructure used in enterprises. Kerberos is the best known example, and
	// Windows usage of that via Active Directory is so widely deployed as to be a
	// de facto standard. In other areas (e.g., academia) there are comparable
	// ticket-based systems.
	TicketSFAuthenticationType SFAuthenticationType = "Ticket"
	// PasswordSFAuthenticationType This enumeration literal specifies
	// Password/shared-secret: Absent an distributed authentication infrastructure,
	// this is what is typically done.
	PasswordSFAuthenticationType SFAuthenticationType = "Password"
)

// DataSanitizationPolicy is The enumberation literals shall specify types of
// data sanitization policies.
type DataSanitizationPolicy string

const (
	// NoneDataSanitizationPolicy This enumeration literal specifies no
	// sanitization.
	NoneDataSanitizationPolicy DataSanitizationPolicy = "None"
	// ClearDataSanitizationPolicy This enumeration literal specifies to sanitize
	// data in all user-addressable storage locations for protection against simple
	// non-invasive data recovery techniques.
	ClearDataSanitizationPolicy DataSanitizationPolicy = "Clear"
	// CryptographicEraseDataSanitizationPolicy This enumeration literal specifies
	// to leverages the encryption of target data by enabling sanitization of the
	// target data's encryption key. This leaves only the ciphertext remaining on
	// the media, effectively sanitizing the data by preventing read-access. For
	// more information, see NIST800-88 and ISO/IEC 27040.
	CryptographicEraseDataSanitizationPolicy DataSanitizationPolicy = "CryptographicErase"
)

// KeySize is The enumeration literals shall specify Key sizes in a symmetric
// encryption algorithm, (see NIST SP 800-57 part 1
// (http:/csrc.nist.gov/publications/nistpubs/800-57/sp800-57_part1_rev3_general.pdf).
type KeySize string

const (
	// Bits0KeySize This enumeration literal specifies that there is no key.
	Bits0KeySize KeySize = "Bits_0"
	// Bits112KeySize This enumeration literal specifies a 3DES 112 bit key.
	Bits112KeySize KeySize = "Bits_112"
	// Bits128KeySize This enumeration literal specifies an AES 128 bit key.
	Bits128KeySize KeySize = "Bits_128"
	// Bits192KeySize This enumeration literal specifies an AES 192 bit key.
	Bits192KeySize KeySize = "Bits_192"
	// Bits256KeySize This enumeration literal specifies an AES 256 bit key.
	Bits256KeySize KeySize = "Bits_256"
)

// SecureChannelProtocol is The enumeration literals shall specify types of
// Secure channel protocols.
type SecureChannelProtocol string

const (
	// NoneSecureChannelProtocol This enumeration literal specifies no encryption.
	NoneSecureChannelProtocol SecureChannelProtocol = "None"
	// TLSSecureChannelProtocol This enumeration literal specifies Transport Layer
	// Security (TLS), as defined by IETF RFC 5246.
	TLSSecureChannelProtocol SecureChannelProtocol = "TLS"
	// IPsecSecureChannelProtocol This enumeration literal specifies Internet
	// Protocol Security (IPsec), as defined by IETF RFC 2401.
	IPsecSecureChannelProtocol SecureChannelProtocol = "IPsec"
	// RPCSECGSSSecureChannelProtocol This enumeration literal specifies RPC access
	// to the Generic Security Services Application Programming Interface
	// (GSS-API), as defined by IETF RPC 2203.
	RPCSECGSSSecureChannelProtocol SecureChannelProtocol = "RPCSEC_GSS"
)

// DataSecurityLoSCapabilities This resource may be used to describe data
// security capabilities.
type DataSecurityLoSCapabilities struct {
	Entity
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupportedAntivirusEngineProviders shall specify supported AntiVirus
	// providers.
	SupportedAntivirusEngineProviders []string
	// SupportedAntivirusScanPolicies shall specify supported policies that trigger
	// an AntiVirus scan.
	SupportedAntivirusScanPolicies []AntiVirusScanTrigger
	// SupportedChannelEncryptionStrengths shall specify supported key sizes in a
	// symmetric encryption algorithm (AES) for transport channel encryption.
	SupportedChannelEncryptionStrengths []KeySize
	// SupportedDataSanitizationPolicies shall specify supported data sanitization
	// policies.
	SupportedDataSanitizationPolicies []DataSanitizationPolicy
	// SupportedHostAuthenticationTypes shall specify supported authentication
	// types for hosts (servers) or initiator endpoints.
	SupportedHostAuthenticationTypes []SFAuthenticationType
	// SupportedLinesOfService shall contain supported DataSecurity service
	// options.
	supportedLinesOfService []string
	// SupportedLinesOfServiceCount
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedMediaEncryptionStrengths shall specify supported key sizes in a
	// symmetric encryption algorithm (AES) for media encryption.
	SupportedMediaEncryptionStrengths []KeySize
	// SupportedSecureChannelProtocols shall specify supported protocols that
	// provide encrypted communication.
	SupportedSecureChannelProtocols []SecureChannelProtocol
	// SupportedUserAuthenticationTypes shall specify supported authentication
	// types for users (or programs).
	SupportedUserAuthenticationTypes []SFAuthenticationType
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a DataSecurityLoSCapabilities object from the raw JSON.
func (d *DataSecurityLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp DataSecurityLoSCapabilities
	var tmp struct {
		temp
		SupportedLinesOfService Links `json:"SupportedLinesOfService"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataSecurityLoSCapabilities(tmp.temp)

	// Extract the links to other entities for later
	d.supportedLinesOfService = tmp.SupportedLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataSecurityLoSCapabilities) Update() error {
	readWriteFields := []string{
		"SupportedAntivirusEngineProviders",
		"SupportedAntivirusScanPolicies",
		"SupportedChannelEncryptionStrengths",
		"SupportedDataSanitizationPolicies",
		"SupportedHostAuthenticationTypes",
		"SupportedLinesOfService",
		"SupportedMediaEncryptionStrengths",
		"SupportedSecureChannelProtocols",
		"SupportedUserAuthenticationTypes",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDataSecurityLoSCapabilities will get a DataSecurityLoSCapabilities instance from the service.
func GetDataSecurityLoSCapabilities(c Client, uri string) (*DataSecurityLoSCapabilities, error) {
	return GetObject[DataSecurityLoSCapabilities](c, uri)
}

// ListReferencedDataSecurityLoSCapabilitiess gets the collection of DataSecurityLoSCapabilities from
// a provided reference.
func ListReferencedDataSecurityLoSCapabilitiess(c Client, link string) ([]*DataSecurityLoSCapabilities, error) {
	return GetCollectionObjects[DataSecurityLoSCapabilities](c, link)
}

// SupportedLinesOfService gets the SupportedLinesOfService linked resources.
func (d *DataSecurityLoSCapabilities) SupportedLinesOfService() ([]*DataSecurityLineOfService, error) {
	return GetObjects[DataSecurityLineOfService](d.client, d.supportedLinesOfService)
}
