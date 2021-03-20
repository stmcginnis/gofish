//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// AntiVirusScanTrigger shall specify types of antivirus scan triggers.
type AntiVirusScanTrigger string

const (
	// NoneAntiVirusScanTrigger specifies there is no trigger.
	NoneAntiVirusScanTrigger AntiVirusScanTrigger = "None"
	// OnFirstReadAntiVirusScanTrigger specifies to trigger on first read.
	OnFirstReadAntiVirusScanTrigger AntiVirusScanTrigger = "OnFirstRead"
	// OnPatternUpdateAntiVirusScanTrigger specifies to trigger on antivirus
	// pattern file update.
	OnPatternUpdateAntiVirusScanTrigger AntiVirusScanTrigger = "OnPatternUpdate"
	// OnUpdateAntiVirusScanTrigger specifies to trigger on object update.
	OnUpdateAntiVirusScanTrigger AntiVirusScanTrigger = "OnUpdate"
	// OnRenameAntiVirusScanTrigger specifies to trigger on object rename.
	OnRenameAntiVirusScanTrigger AntiVirusScanTrigger = "OnRename"
)

// AuthenticationType shall specify authentication algorithms.
type AuthenticationType string

const (
	// NoneAuthenticationType specifies No authentication.
	NoneAuthenticationType AuthenticationType = "None"
	// PKIAuthenticationType specifies a Public Key Infrastructure. Customers
	// with the highest assurance requirements roll PKI out to hosts and users
	// (it is more common for hosts than users. User PKI-based authentication
	// has significant operational complications and administrative overheads,
	// e.g., smart cards may be involved.
	PKIAuthenticationType AuthenticationType = "PKI"
	// TicketAuthenticationType specifies Ticket-
	// based (e.g., Kerberos): This is the most common class of
	// authentication infrastructure used in enterprises. Kerberos is the
	// best known example, and Windows usage of that via Active Directory is
	// so widely deployed as to be a de facto standard. In other areas (e.g.,
	// academia) there are comparable ticket-based systems.
	TicketAuthenticationType AuthenticationType = "Ticket"
	// PasswordAuthenticationType specifies
	// Password/shared-secret: Absent an distributed authentication
	// infrastructure, this is what is typically done.
	PasswordAuthenticationType AuthenticationType = "Password"
)

// DataSanitizationPolicy shall specify types of data sanitization policies.
type DataSanitizationPolicy string

const (
	// NoneDataSanitizationPolicy specifies no sanitization.
	NoneDataSanitizationPolicy DataSanitizationPolicy = "None"
	// ClearDataSanitizationPolicy specifies to
	// sanitize data in all user-addressable storage locations for protection
	// against simple non-invasive data recovery techniques.
	ClearDataSanitizationPolicy DataSanitizationPolicy = "Clear"
	// CryptographicEraseDataSanitizationPolicy This enumeration literal
	// specifies to leverages the encryption of target data by enabling
	// sanitization of the target data's encryption key. This leaves only the
	// ciphertext remaining on the media, effectively sanitizing the data by
	// preventing read-access. For more information, see NIST800-88 and
	// ISO/IEC 27040.
	CryptographicEraseDataSanitizationPolicy DataSanitizationPolicy = "CryptographicErase"
)

// KeySize shall specify Key sizes in a
// symmetric encryption algorithm, (see NIST SP 800-57 part 1 (
// http:/csrc.nist.gov/publications/nistpubs/800-57/sp800-57_part1_rev3_general.pdf)
type KeySize string

const (
	// Bits0KeySize specifies that there is no key.
	Bits0KeySize KeySize = "Bits_0"
	// Bits112KeySize specifies a 3DES 112 bit key.
	Bits112KeySize KeySize = "Bits_112"
	// Bits128KeySize specifies an AES 128 bit key.
	Bits128KeySize KeySize = "Bits_128"
	// Bits192KeySize specifies an AES 192 bit key.
	Bits192KeySize KeySize = "Bits_192"
	// Bits256KeySize specifies an AES 256 bit key.
	Bits256KeySize KeySize = "Bits_256"
)

// SecureChannelProtocol shall specify types
// of Secure channel protocols.
type SecureChannelProtocol string

const (
	// NoneSecureChannelProtocol specifies no encryption.
	NoneSecureChannelProtocol SecureChannelProtocol = "None"
	// TLSSecureChannelProtocol specifies Transport
	// Layer Security (TLS), as defined by IETF RFC 5246.
	TLSSecureChannelProtocol SecureChannelProtocol = "TLS"
	// IPsecSecureChannelProtocol specifies Internet
	// Protocol Security (IPsec), as defined by IETF RFC 2401.
	IPsecSecureChannelProtocol SecureChannelProtocol = "IPsec"
	// RPCSECGSSSecureChannelProtocol specifies RPC
	// access to the Generic Security Services Application Programming
	// Interface (GSS-API), as defined by IETF RPC 2203.
	RPCSECGSSSecureChannelProtocol SecureChannelProtocol = "RPCSEC_GSS"
)

// DataSecurityLoSCapabilities is used to describe data security capabilities.
type DataSecurityLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// SupportedAntivirusEngineProviders shall specify supported AntiVirus providers.
	SupportedAntivirusEngineProviders []string
	// SupportedAntivirusScanPolicies shall specify supported policies that
	// trigger an AntiVirus scan.
	SupportedAntivirusScanPolicies []AntiVirusScanTrigger
	// SupportedChannelEncryptionStrengths shall specify supported key sizes in
	// a symmetric encryption algorithm (AES) for transport channel encryption.
	SupportedChannelEncryptionStrengths []KeySize
	// SupportedDataSanitizationPolicies shall specify supported data
	// sanitization policies.
	SupportedDataSanitizationPolicies []DataSanitizationPolicy
	// SupportedHostAuthenticationTypes shall specify supported authentication
	// types for hosts (servers) or initiator endpoints.
	SupportedHostAuthenticationTypes []AuthenticationType
	// SupportedLinesOfService shall contain supported DataSecurity service options.
	SupportedLinesOfService []DataSecurityLineOfService
	// SupportedLinesOfServiceCount is the number of supported lines of service.
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedMediaEncryptionStrengths shall specify supported key sizes in a
	// symmetric encryption algorithm (AES) for media encryption.
	SupportedMediaEncryptionStrengths []KeySize
	// SupportedSecureChannelProtocols shall specify supported protocols that
	// provide encrypted communication.
	SupportedSecureChannelProtocols []SecureChannelProtocol
	// SupportedUserAuthenticationTypes shall specify supported authentication
	// types for users (or programs).
	SupportedUserAuthenticationTypes []AuthenticationType
	// rawData holds the original serialized JSON so we can compare updates.
	// rawData []byte
}

// // UnmarshalJSON unmarshals a DataSecurityLoSCapabilities object from the raw JSON.
// func (datasecurityloscapabilities *DataSecurityLoSCapabilities) UnmarshalJSON(b []byte) error {
// 	type temp DataSecurityLoSCapabilities
// 	var t struct {
// 		temp
// 	}

// 	err := json.Unmarshal(b, &t)
// 	if err != nil {
// 		return err
// 	}

// 	*datasecurityloscapabilities = DataSecurityLoSCapabilities(t.temp)

// 	// Extract the links to other entities for later

// 	// This is a read/write object, so we need to save the raw object data for later
// 	datasecurityloscapabilities.rawData = b

// 	return nil
// }

// // Update commits updates to this object's properties to the running system.
// func (datasecurityloscapabilities *DataSecurityLoSCapabilities) Update() error {

// 	// Get a representation of the object's original state so we can find what
// 	// to update.
// 	original := new(DataSecurityLoSCapabilities)
// 	original.UnmarshalJSON(datasecurityloscapabilities.rawData)

// 	readWriteFields := []string{
// 		"SupportedAntivirusEngineProviders",
// 		"SupportedAntivirusScanPolicies",
// 		"SupportedChannelEncryptionStrengths",
// 		"SupportedDataSanitizationPolicies",
// 		"SupportedHostAuthenticationTypes",
// 		"SupportedLinesOfService",
// 		"SupportedMediaEncryptionStrengths",
// 		"SupportedSecureChannelProtocols",
// 		"SupportedUserAuthenticationTypes",
// 	}

// 	originalElement := reflect.ValueOf(original).Elem()
// 	currentElement := reflect.ValueOf(datasecurityloscapabilities).Elem()

// 	return datasecurityloscapabilities.Entity.Update(originalElement, currentElement, readWriteFields)
// }

// GetDataSecurityLoSCapabilities will get a DataSecurityLoSCapabilities instance from the service.
func GetDataSecurityLoSCapabilities(c common.Client, uri string) (*DataSecurityLoSCapabilities, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var datasecurityloscapabilities DataSecurityLoSCapabilities
	err = json.NewDecoder(resp.Body).Decode(&datasecurityloscapabilities)
	if err != nil {
		return nil, err
	}

	datasecurityloscapabilities.SetClient(c)
	return &datasecurityloscapabilities, nil
}

// ListReferencedDataSecurityLoSCapabilities gets the collection of DataSecurityLoSCapabilities from
// a provided reference.
func ListReferencedDataSecurityLoSCapabilities(c common.Client, link string) ([]*DataSecurityLoSCapabilities, error) {
	var result []*DataSecurityLoSCapabilities
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, datasecurityloscapabilitiesLink := range links.ItemLinks {
		datasecurityloscapabilities, err := GetDataSecurityLoSCapabilities(c, datasecurityloscapabilitiesLink)
		if err != nil {
			return result, err
		}
		result = append(result, datasecurityloscapabilities)
	}

	return result, nil
}
