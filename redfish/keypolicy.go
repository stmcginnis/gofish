//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type KeyPolicyType string

const (
	// NVMeoFKeyPolicyType shall indicate the key policy is for an NVMe-oF key.
	NVMeoFKeyPolicyType KeyPolicyType = "NVMeoF"
)

// NVMeoFCipherSuiteType is This enumeration shall list the NVMe cipher suites that a key is allowed to use.
type NVMeoFCipherSuiteType string

const (
	// TLSAES128GCMSHA256NVMeoFCipherSuiteType shall indicate TLS_AES_128_GCM_SHA256 as defined by the 'Mandatory and
	// Recommended Cipher Suites' clause in the NVMe TCP Transport Specification.
	TLSAES128GCMSHA256NVMeoFCipherSuiteType NVMeoFCipherSuiteType = "TLS_AES_128_GCM_SHA256"
	// TLSAES256GCMSHA384NVMeoFCipherSuiteType shall indicate TLS_AES_256_GCM_SHA384 as defined by the 'Mandatory and
	// Recommended Cipher Suites' clause in the NVMe TCP Transport Specification.
	TLSAES256GCMSHA384NVMeoFCipherSuiteType NVMeoFCipherSuiteType = "TLS_AES_256_GCM_SHA384"
)

// NVMeoFDHGroupType is This enumeration shall list the Diffie-Hellman (DH) groups that a key is allowed to use.
type NVMeoFDHGroupType string

const (
	// FFDHE2048NVMeoFDHGroupType shall indicate the 2048-bit Diffie-Hellman (DH) group as defined by the 'DH-HMAC-CHAP
	// Diffie-Hellman group identifiers' figure in the NVMe Base Specification.
	FFDHE2048NVMeoFDHGroupType NVMeoFDHGroupType = "FFDHE2048"
	// FFDHE3072NVMeoFDHGroupType shall indicate the 3072-bit Diffie-Hellman (DH) group as defined by the 'DH-HMAC-CHAP
	// Diffie-Hellman group identifiers' figure in the NVMe Base Specification.
	FFDHE3072NVMeoFDHGroupType NVMeoFDHGroupType = "FFDHE3072"
	// FFDHE4096NVMeoFDHGroupType shall indicate the 4096-bit Diffie-Hellman (DH) group as defined by the 'DH-HMAC-CHAP
	// Diffie-Hellman group identifiers' figure in the NVMe Base Specification.
	FFDHE4096NVMeoFDHGroupType NVMeoFDHGroupType = "FFDHE4096"
	// FFDHE6144NVMeoFDHGroupType shall indicate the 2048-bit Diffie-Hellman (DH) group as defined by the 'DH-HMAC-CHAP
	// Diffie-Hellman group identifiers' figure in the NVMe Base Specification.
	FFDHE6144NVMeoFDHGroupType NVMeoFDHGroupType = "FFDHE6144"
	// FFDHE8192NVMeoFDHGroupType shall indicate the 8192-bit Diffie-Hellman (DH) group as defined by the 'DH-HMAC-CHAP
	// Diffie-Hellman group identifiers' figure in the NVMe Base Specification.
	FFDHE8192NVMeoFDHGroupType NVMeoFDHGroupType = "FFDHE8192"
)

// NVMeoFSecureHashType is This enumeration shall list the NVMe secure hash algorithms that a key is allowed to
// use.
type NVMeoFSecureHashType string

const (
	// SHA256NVMeoFSecureHashType shall indicate the SHA-256 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA256NVMeoFSecureHashType NVMeoFSecureHashType = "SHA256"
	// SHA384NVMeoFSecureHashType shall indicate the SHA-384 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA384NVMeoFSecureHashType NVMeoFSecureHashType = "SHA384"
	// SHA512NVMeoFSecureHashType shall indicate the SHA-512 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA512NVMeoFSecureHashType NVMeoFSecureHashType = "SHA512"
)

// NVMeoFSecurityProtocolType is a list of the NVMe security protocols that a key is allowed to
// use.
type NVMeoFSecurityProtocolType string

const (
	// DHHCNVMeoFSecurityProtocolType shall indicate the Diffie-Hellman Hashed Message Authentication Code Challenge
	// Handshake Authentication Protocol (DH-HMAC-CHAP) as defined by the NVMe Base Specification.
	DHHCNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "DHHC"
	// TLSPSKNVMeoFSecurityProtocolType shall indicate Transport Layer Security Pre-Shared Key (TLS PSK) as defined by
	// the NVMe TCP Transport Specification.
	TLSPSKNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "TLS_PSK"
	// OEMNVMeoFSecurityProtocolType shall indicate an OEM-defined security protocol. The OEMSecurityProtocolAllowList
	// property shall contain the specific OEM protocol.
	OEMNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "OEM"
)

// NVMeoFSecurityTransportType is This enumeration shall list the NVMe security transports that a key is allowed to
// use.
type NVMeoFSecurityTransportType string

const (
	// TLSv2NVMeoFSecurityTransportType shall indicate Transport Layer Security (TLS) v2 as defined by the 'Transport
	// Specific Address Subtype Definition for NVMe/TCP Transport' figure in the NVMe TCP Transport Specification.
	TLSv2NVMeoFSecurityTransportType NVMeoFSecurityTransportType = "TLSv2"
	// TLSv3NVMeoFSecurityTransportType shall indicate Transport Layer Security (TLS) v3 as defined by the 'Transport
	// Specific Address Subtype Definition for NVMe/TCP Transport' figure in the NVMe TCP Transport Specification.
	TLSv3NVMeoFSecurityTransportType NVMeoFSecurityTransportType = "TLSv3"
)

// KeyPolicy shall represent a key policy for a Redfish implementation.
type KeyPolicy struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IsDefault shall indicate if this key policy is the policy applied when no other policies are specified.
	IsDefault bool
	// KeyPolicyType shall contain the type of key policy.
	KeyPolicyType KeyPolicyType
	// NVMeoF shall contain NVMe-oF specific properties for this key policy. This property shall be present if
	// KeyPolicyType contains the value 'NVMeoF'.
	NVMeoF NVMeoF
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a KeyPolicy object from the raw JSON.
func (keypolicy *KeyPolicy) UnmarshalJSON(b []byte) error {
	type temp KeyPolicy
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keypolicy = KeyPolicy(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	keypolicy.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (keypolicy *KeyPolicy) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(KeyPolicy)
	original.UnmarshalJSON(keypolicy.rawData)

	readWriteFields := []string{
		"IsDefault",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(keypolicy).Elem()

	return keypolicy.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetKeyPolicy will get a KeyPolicy instance from the service.
func GetKeyPolicy(c common.Client, uri string) (*KeyPolicy, error) {
	return common.GetObject[KeyPolicy](c, uri)
}

// ListReferencedKeyPolicys gets the collection of KeyPolicy from
// a provided reference.
func ListReferencedKeyPolicys(c common.Client, link string) ([]*KeyPolicy, error) {
	return common.GetCollectionObjects[KeyPolicy](c, link)
}

// NVMeoF shall contain NVMe-oF specific properties for a key policy.
type NVMeoF struct {
	// CipherSuiteAllowList shall contain the cipher suites that this key policy allows. The absence of the property
	// shall indicate any cipher suite is allowed. An empty list shall indicate no cipher suites are allowed.
	CipherSuiteAllowList []NVMeoFCipherSuiteType
	// DHGroupAllowList shall contain the Diffie-Hellman (DH) groups that this key policy allows. The absence of the
	// property shall indicate any DH group is allowed. An empty list shall indicate no DH groups are allowed.
	DHGroupAllowList []NVMeoFDHGroupType
	// OEMSecurityProtocolAllowList shall contain the OEM-defined security protocols that this key policy allows. NVMe-
	// oF channels are restricted to OEM-defined security protocols in this list. An empty list shall indicate no
	// security protocols are allowed. This property shall be present if SecurityProtocolAllowList contains 'OEM'.
	OEMSecurityProtocolAllowList []string
	// SecureHashAllowList shall contain the secure hash algorithms that this key policy allows. The absence of the
	// property shall indicate any secure hash algorithm is allowed. An empty list shall indicate no secure hash
	// algorithms are allowed.
	SecureHashAllowList []NVMeoFSecureHashType
	// SecurityProtocolAllowList shall contain the security protocols that this key policy allows. NVMe-oF channels are
	// restricted to security protocols in this list. The absence of the property shall indicate any security protocol
	// is allowed. An empty list shall indicate no security protocols are allowed.
	SecurityProtocolAllowList []NVMeoFSecurityProtocolType
	// SecurityTransportAllowList shall contain the security transports that this key policy allows. The absence of the
	// property shall indicate any security transport is allowed. An empty list shall indicate no security transports
	// are allowed.
	SecurityTransportAllowList []NVMeoFSecurityTransportType
}
