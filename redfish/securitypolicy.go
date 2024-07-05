//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SPDMAlgorithmSet shall contain SPDM algorithm settings.
type SPDMAlgorithmSet struct {
	// AEAD shall contain an array of AEAD algorithms. The allowable values for this property shall be the AEAD
	// algorithm names found in the 'AlgSupported' field of the 'AEAD structure' table in DSP0274, 'ALL', and 'NONE'.
	// An array containing one element with the value of 'ALL' or an empty array shall indicate all AEAD algorithms. An
	// array containing one element with a value of 'NONE' shall indicate no AEAD algorithms.
	AEAD []string
	// BaseAsym shall contain an array of asymmetric signature algorithms. The allowable values for this property shall
	// be the asymmetric key signature algorithm names found in the 'BaseAsymAlgo' field of the 'NEGOTIATE_ALGORITHMS'
	// request message in DSP0274, 'ALL', and 'NONE'. An array containing one element with the value of 'ALL' or an
	// empty array shall indicate all asymmetric signature algorithms. An array containing one element with a value of
	// 'NONE' shall indicate no asymmetric signature algorithms.
	BaseAsym []string
	// BaseHash shall contain an array of hash algorithms. The allowable values for this property shall be the hash
	// algorithm names found in the 'BaseHashAlgo' field of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274,
	// 'ALL', and 'NONE'. An array containing one element with the value of 'ALL' or an empty array shall indicate all
	// hash algorithms. An array containing one element with a value of 'NONE' shall indicate no hash algorithms.
	BaseHash []string
}

// SPDMParameterSet shall contain SPDM policy settings.
type SPDMParameterSet struct {
	// Algorithms shall contain the SPDM algorithms.
	Algorithms SPDMAlgorithmSet
	// Versions shall contain an array of SPDM versions. An array containing one element with the value of 'ALL' or an
	// empty array shall indicate all versions. An array containing one element with a value of 'NONE' shall indicate
	// no versions.
	Versions []string
}

// SPDMPolicy shall contain SPDM policy settings.
type SPDMPolicy struct {
	// AllowExtendedAlgorithms shall indicate whether SPDM extended algorithms as defined in DSP0274 are allowed.
	AllowExtendedAlgorithms bool
	// Allowed shall contain the SPDM policy settings that are allowed, such as the allowable SPDM versions and
	// algorithms.
	Allowed SPDMParameterSet
	// Denied shall contain the SPDM policy settings that are prohibited, such as the prohibited SPDM versions and
	// algorithms.
	Denied SPDMParameterSet
	// Enabled shall indicate whether SPDM communication with devices as defined in DSP0274 is enabled.
	Enabled bool
	// RevokedCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the set of revoked SPDM device certificates. Certificates in this collection may contain leaf certificates,
	// partial certificate chains, or complete certificate chains, where a partial certificate chain is a chain
	// containing only CA certificates. If 'VerifyCertificate' contains the value 'true' and if an SPDM endpoint
	// verifies successfully against a partial chain or exactly matches a leaf certificate, that SPDM endpoint shall
	// fail authentication.
	revokedCertificates string
	// SecureSessionEnabled shall indicate whether SPDM secure sessions with devices as defined in DSP0274 is enabled.
	SecureSessionEnabled bool
	// TrustedCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the set of trusted SPDM device certificates. Certificates in this collection may contain leaf certificates,
	// partial certificate chains, or complete certificate chains, where a partial certificate chain is a chain
	// containing only CA certificates. If 'VerifyCertificate' contains the value 'true' and if an SPDM endpoint
	// verifies successfully against a partial chain or exactly matches a leaf certificate, that SPDM endpoint shall be
	// considered verified and other authentications checks are performed.
	trustedCertificates string
	// VerifyCertificate shall indicate whether the manager will verify the certificate of the SPDM endpoint. If
	// 'true', the manager shall verify the device certificate with the certificates found in the collections
	// referenced by the 'RevokedCertificates' and 'TrustedCertificates' properties. If 'false', the manager shall not
	// perform verification of the endpoint certificate.
	VerifyCertificate bool
}

// UnmarshalJSON unmarshals a SPDMPolicy object from the raw JSON.
func (spdmpolicy *SPDMPolicy) UnmarshalJSON(b []byte) error {
	type temp SPDMPolicy
	var t struct {
		temp
		RevokedCertificates common.Link
		TrustedCertificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmpolicy = SPDMPolicy(t.temp)

	// Extract the links to other entities for later
	spdmpolicy.revokedCertificates = t.RevokedCertificates.String()
	spdmpolicy.trustedCertificates = t.TrustedCertificates.String()

	return nil
}

// RevokedCertificates gets the set of revoked SPDM device certificates.
func (spdmpolicy *SPDMPolicy) RevokedCertificates(c common.Client) ([]*Certificate, error) {
	return ListReferencedCertificates(c, spdmpolicy.revokedCertificates)
}

// TrustedCertificates gets the set of trusted SPDM device certificates.
func (spdmpolicy *SPDMPolicy) TrustedCertificates(c common.Client) ([]*Certificate, error) {
	return ListReferencedCertificates(c, spdmpolicy.trustedCertificates)
}

// SecurityPolicy shall represent configurable security-related policies managed by a manager. All security
// parameters in other resources that are controlled by the manager shall follow the related settings in this
// security policy. For example, an outbound TLS connection established per an EventDestination resource will
// follow the values of the properties in the TLS property.
type SecurityPolicy struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OverrideParentManager shall indicate whether this security policy overrides the security policy of the managers
	// referenced by the ManagedBy property within the Links property of the Manager resource for this security policy.
	// If this property is absent, the value shall be assumed to be 'false'.
	OverrideParentManager bool
	// SPDM shall contain the policy requirements for SPDM communication and usage.
	SPDM SPDMPolicy
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TLS shall contain the policy requirements for TLS communication and usage.
	TLS TLSCommunication
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SecurityPolicy object from the raw JSON.
func (securitypolicy *SecurityPolicy) UnmarshalJSON(b []byte) error {
	type temp SecurityPolicy
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*securitypolicy = SecurityPolicy(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	securitypolicy.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (securitypolicy *SecurityPolicy) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SecurityPolicy)
	original.UnmarshalJSON(securitypolicy.rawData)

	readWriteFields := []string{
		"OverrideParentManager",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(securitypolicy).Elem()

	return securitypolicy.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSecurityPolicy will get a SecurityPolicy instance from the service.
func GetSecurityPolicy(c common.Client, uri string) (*SecurityPolicy, error) {
	return common.GetObject[SecurityPolicy](c, uri)
}

// ListReferencedSecurityPolicys gets the collection of SecurityPolicy from
// a provided reference.
func ListReferencedSecurityPolicys(c common.Client, link string) ([]*SecurityPolicy, error) {
	return common.GetCollectionObjects[SecurityPolicy](c, link)
}

// TLSAlgorithmSet shall contain TLS algorithm settings.
type TLSAlgorithmSet struct {
	// CipherSuites shall contain an array of TLS cipher suites. The allowable values for this property shall be the
	// TLS cipher suites listed in 'CipherSuites' defined in, but not limited to, RFC4346, RFC5246, or RFC8446, 'ALL',
	// and 'NONE'. An array containing one element with the value of 'ALL' or an empty array shall indicate all TLS
	// cipher suites. An array containing one element with a value of 'NONE' shall indicate no TLS cipher suites.
	CipherSuites []string
	// SignatureAlgorithms shall contain an array of TLS signature algorithms. The allowable values for this property
	// shall be the TLS signature algorithms listed in 'SignatureScheme' or the concatenation of 'SignatureAlgorithm',
	// '_', and 'HashAlgorithm' defined in, but not limited to, RFC4346, RFC5246, or RFC8446, 'ALL', and 'NONE'. An
	// array containing one element with the value of 'ALL' or an empty array shall indicate all TLS signature
	// algorithms. An array containing one element with a value of 'NONE' shall indicate no TLS signature algorithms.
	SignatureAlgorithms []string
}

// TLSCommunication shall contain the policy requirements for TLS communication and usage for a TLS client and
// server.
type TLSCommunication struct {
	// Client shall contain the policy requirements and usage for TLS connections where the manager acts as a TLS
	// client.
	Client TLSPolicy
	// Server shall contain the policy requirements and usage for TLS connections where the manager acts as a TLS
	// server.
	Server TLSPolicy
}

// TLSParameterSet shall contain TLS policy settings.
type TLSParameterSet struct {
	// Algorithms shall contain the TLS algorithms.
	Algorithms TLSAlgorithmSet
	// Versions shall contain an array of TLS versions. An array containing one element with the value of 'ALL' or an
	// empty array shall indicate all versions. An array containing one element with a value of 'NONE' shall indicate
	// no versions.
	Versions []string
}

// TLSPolicy shall contain TLS policy settings.
type TLSPolicy struct {
	// Allowed shall contain the TLS policy settings that are allowed, such as the allowable TLS versions and
	// algorithms. If a value is missing for the same property in the 'Allowed' and 'Denied' object, the missing value
	// shall behave as if the value is present in the same property under the 'Denied' object. If a value conflicts for
	// the same property between the 'Allowed' and 'Denied' object, the value of the same property in the 'Denied'
	// object shall take precedence. A Redfish service can resolve or prevent conflicts at time of request as well.
	Allowed TLSParameterSet
	// Denied shall contain the TLS policy settings that are prohibited, such as the prohibited TLS versions and
	// algorithms. If a value is missing for the same property in the 'Allowed' and 'Denied' object, the missing value
	// shall behave as if the value is present in the same property under the 'Denied' object. If a value conflicts for
	// the same property between the 'Allowed' and 'Denied' object, the value of the same property in the 'Denied'
	// object shall take precedence. A Redfish service can resolve or prevent conflicts at time of request as well.
	Denied TLSParameterSet
	// RevokedCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the set of revoked TLS certificates. Certificates in this collection may contain leaf certificates, partial
	// certificate chains, or complete certificate chains, where a partial certificate chain is a chain containing only
	// CA certificates. If 'VerifyCertificate' contains the value 'true' and if a TLS endpoint verifies successfully
	// against a partial chain or exactly matches a leaf certificate, that TLS endpoint shall fail authentication.
	revokedCertificates string
	// TrustedCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the set of trusted TLS certificates. Certificates in this collection may contain leaf certificates, partial
	// certificate chains, or complete certificate chains, where a partial certificate chain is a chain containing only
	// CA certificates. If 'VerifyCertificate' contains the value 'true' and if a TLS endpoint verifies successfully
	// against a partial chain or exactly matches a leaf certificate, that TLS endpoint shall be considered verified
	// and other authentications checks are performed.
	trustedCertificates string
	// VerifyCertificate shall indicate whether the manager will verify the certificate of the remote endpoint in a TLS
	// connection. If 'true', the manager shall verify the remote endpoint certificate with the certificates found in
	// the collections referenced by the 'RevokedCertificates' and 'TrustedCertificates' properties. If 'false' or not
	// present, the manager shall not perform verification of the endpoint certificate.
	VerifyCertificate bool
}

// UnmarshalJSON unmarshals a TLSPolicy object from the raw JSON.
func (tlspolicy *TLSPolicy) UnmarshalJSON(b []byte) error {
	type temp TLSPolicy
	var t struct {
		temp
		RevokedCertificates common.Link
		TrustedCertificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tlspolicy = TLSPolicy(t.temp)

	// Extract the links to other entities for later
	tlspolicy.revokedCertificates = t.RevokedCertificates.String()
	tlspolicy.trustedCertificates = t.TrustedCertificates.String()

	return nil
}

// RevokedCertificates gets the set of revoked TLS certificates.
func (tlspolicy *TLSPolicy) RevokedCertificates(c common.Client) ([]*Certificate, error) {
	return ListReferencedCertificates(c, tlspolicy.revokedCertificates)
}

// TrustedCertificates gets the set of trusted TLS certificates.
func (tlspolicy *TLSPolicy) TrustedCertificates(c common.Client) ([]*Certificate, error) {
	return ListReferencedCertificates(c, tlspolicy.trustedCertificates)
}
