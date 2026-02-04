//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Certificate.v1_11_0.json
// 2025.3 - #Certificate.v1_11_0.Certificate

package schemas

import (
	"encoding/json"
)

type CertificateType string

const (
	// PEMCertificateType shall indicate the format of the certificate shall
	// contain a Privacy Enhanced Mail (PEM)-encoded string, containing
	// RFC5280-defined structures, representing a single certificate.
	PEMCertificateType CertificateType = "PEM"
	// PEMchainCertificateType shall indicate the format of the certificate shall
	// contain a Privacy Enhanced Mail (PEM)-encoded string, containing
	// RFC5280-defined structures, representing a certificate chain. When this
	// value is specified, the properties of the resource, except for
	// 'CertificateString', shall contain the information from the leaf
	// certificate. When this value is specified, 'CertificateString' shall contain
	// the entire PEM chain.
	PEMchainCertificateType CertificateType = "PEMchain"
	// PKCS7CertificateType shall contain a Privacy Enhanced Mail (PEM)-encoded
	// string, containing RFC5280-defined and RFC2315-defined structures. The
	// service may discard additional certificates or other data in the structure.
	PKCS7CertificateType CertificateType = "PKCS7"
	// PKCS12CertificateType shall indicate the format of the certificate shall
	// contain a Base64-encoded string, with padding characters, containing
	// RFC7292-defined structures, representing a PKCS12 archive of one or more
	// certificates, keys, or other cryptographic data. The service may discard
	// additional certificates or other data in the structure.
	PKCS12CertificateType CertificateType = "PKCS12"
)

type CertificateUsageType string

const (
	// UserCertificateUsageType This certificate is a user certificate like those
	// associated with a manager account.
	UserCertificateUsageType CertificateUsageType = "User"
	// WebCertificateUsageType This certificate is a web or HTTPS certificate like
	// those used for event destinations.
	WebCertificateUsageType CertificateUsageType = "Web"
	// SSHCertificateUsageType This certificate is used for SSH.
	SSHCertificateUsageType CertificateUsageType = "SSH"
	// DeviceCertificateUsageType This certificate is a device type certificate
	// like those associated with SPDM and other standards.
	DeviceCertificateUsageType CertificateUsageType = "Device"
	// PlatformCertificateUsageType This certificate is a platform type certificate
	// like those associated with SPDM and other standards.
	PlatformCertificateUsageType CertificateUsageType = "Platform"
	// BIOSCertificateUsageType This certificate is a BIOS certificate like those
	// associated with UEFI.
	BIOSCertificateUsageType CertificateUsageType = "BIOS"
	// IDevIDCertificateUsageType This certificate is an IDevID certificate like
	// those associated with TCG TPMs.
	IDevIDCertificateUsageType CertificateUsageType = "IDevID"
	// LDevIDCertificateUsageType This certificate is an LDevID certificate like
	// those associated with TCG TPMs.
	LDevIDCertificateUsageType CertificateUsageType = "LDevID"
	// IAKCertificateUsageType This certificate is an IAK certificate like those
	// associated with TCG TPMs.
	IAKCertificateUsageType CertificateUsageType = "IAK"
	// LAKCertificateUsageType This certificate is an LAK certificate like those
	// associated with TCG TPMs.
	LAKCertificateUsageType CertificateUsageType = "LAK"
	// EKCertificateUsageType This certificate is an EK certificate like those
	// associated with TCG TPMs.
	EKCertificateUsageType CertificateUsageType = "EK"
)

// KeyUsage is This type shall describe the usages of a key within a
// certificate, as specified by the 'Key Usage' and 'Extended Key Usage'
// definitions in RFC5280.
type KeyUsage string

const (
	// DigitalSignatureKeyUsage Verifies digital signatures, other than signatures
	// on certificates and CRLs.
	DigitalSignatureKeyUsage KeyUsage = "DigitalSignature"
	// NonRepudiationKeyUsage Verifies digital signatures, other than signatures on
	// certificates and CRLs, and provides a non-repudiation service that protects
	// against the signing entity falsely denying some action.
	NonRepudiationKeyUsage KeyUsage = "NonRepudiation"
	// KeyEnciphermentKeyUsage Enciphers private or secret keys.
	KeyEnciphermentKeyUsage KeyUsage = "KeyEncipherment"
	// DataEnciphermentKeyUsage Directly enciphers raw user data without an
	// intermediate symmetric cipher.
	DataEnciphermentKeyUsage KeyUsage = "DataEncipherment"
	// KeyAgreementKeyUsage Key agreement.
	KeyAgreementKeyUsage KeyUsage = "KeyAgreement"
	// KeyCertSignKeyUsage Verifies signatures on public key certificates.
	KeyCertSignKeyUsage KeyUsage = "KeyCertSign"
	// CRLSigningKeyUsage Verifies signatures on certificate revocation lists
	// (CRLs).
	CRLSigningKeyUsage KeyUsage = "CRLSigning"
	// EncipherOnlyKeyUsage Enciphers data while performing a key agreement.
	EncipherOnlyKeyUsage KeyUsage = "EncipherOnly"
	// DecipherOnlyKeyUsage Deciphers data while performing a key agreement.
	DecipherOnlyKeyUsage KeyUsage = "DecipherOnly"
	// ServerAuthenticationKeyUsage TLS WWW server authentication.
	ServerAuthenticationKeyUsage KeyUsage = "ServerAuthentication"
	// ClientAuthenticationKeyUsage TLS WWW client authentication.
	ClientAuthenticationKeyUsage KeyUsage = "ClientAuthentication"
	// CodeSigningKeyUsage Signs downloadable executable code.
	CodeSigningKeyUsage KeyUsage = "CodeSigning"
	// EmailProtectionKeyUsage Email protection.
	EmailProtectionKeyUsage KeyUsage = "EmailProtection"
	// TimestampingKeyUsage Binds the hash of an object to a time.
	TimestampingKeyUsage KeyUsage = "Timestamping"
	// OCSPSigningKeyUsage Signs OCSP responses.
	OCSPSigningKeyUsage KeyUsage = "OCSPSigning"
)

// Certificate shall represent a certificate for a Redfish implementation.
type Certificate struct {
	Entity
	// CertificateString shall contain the certificate, and the format shall follow
	// the requirements specified by the 'CertificateType' property value. If the
	// certificate contains any private keys, they shall be removed from the string
	// in responses. If the service does not know the private key for the
	// certificate and is needed to use the certificate, the client shall provide
	// the private key as part of the string in the 'POST' request. For additional
	// property requirements, see the corresponding definition in the Redfish Data
	// Model Specification.
	CertificateString string
	// CertificateType shall contain the format type for the certificate. For
	// additional property requirements, see the corresponding definition in the
	// Redfish Data Model Specification.
	CertificateType CertificateType
	// CertificateUsageTypes shall contain an array describing the types or
	// purposes for this certificate.
	//
	// Version added: v1.4.0
	CertificateUsageTypes []CertificateUsageType
	// Fingerprint shall be a string containing the ASCII representation of the
	// fingerprint of the certificate. The hash algorithm used to generate this
	// fingerprint shall be specified by the 'FingerprintHashAlgorithm' property.
	//
	// Version added: v1.3.0
	Fingerprint string
	// FingerprintHashAlgorithm shall be a string containing the hash algorithm
	// used for generating the 'Fingerprint' property. The value shall be one of
	// the strings in the 'Algorithm Name' field of the 'TCG_ALG_ID Constants'
	// table, formerly the 'TPM_ALG_ID Constants' table, within the 'Trusted
	// Computing Group Algorithm Registry'.
	//
	// Version added: v1.3.0
	FingerprintHashAlgorithm string
	// Issuer shall contain an object containing information about the issuer of
	// the certificate.
	Issuer CertificateIdentifier
	// KeyUsage shall contain the key usage extension, which defines the purpose of
	// the public keys in this certificate.
	KeyUsage []KeyUsage
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password for the certificate contained in the
	// 'CertificateString' property. This property shall be required in create
	// requests if the 'CertificateType' property contains 'PKCS12' and the
	// client-provided certificate is password protected. This property shall not
	// be present in responses.
	//
	// Version added: v1.10.0
	Password string
	// SPDM shall contain SPDM-related information for the certificate. This
	// property shall only be present for SPDM certificates.
	//
	// Version added: v1.5.0
	SPDM SPDM
	// SerialNumber shall be a string containing the ASCII representation of the
	// serial number of the certificate, as defined by the RFC5280 'serialNumber'
	// field.
	//
	// Version added: v1.3.0
	SerialNumber string
	// SignatureAlgorithm shall be a string containing the algorithm used for
	// generating the signature of the certificate, as defined by the RFC5280
	// 'signatureAlgorithm' field. The value shall be a string representing the
	// ASN.1 OID of the signature algorithm as defined in, but not limited to,
	// RFC3279, RFC4055, or RFC4491.
	//
	// Version added: v1.3.0
	SignatureAlgorithm string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.10.0
	Status Status
	// Subject shall contain an object containing information about the subject of
	// the certificate.
	Subject CertificateIdentifier
	// UefiSignatureOwner shall contain the GUID of the UEFI signature owner for
	// this certificate as defined by the UEFI Specification. This property shall
	// only be present for certificates managed by UEFI.
	//
	// Version added: v1.2.0
	UefiSignatureOwner string
	// ValidNotAfter shall contain the date when the certificate validity period
	// ends.
	ValidNotAfter string
	// ValidNotBefore shall contain the date when the certificate validity period
	// begins.
	ValidNotBefore string
	// forceAutomaticRenewTarget is the URL to send ForceAutomaticRenew requests.
	forceAutomaticRenewTarget string
	// rekeyTarget is the URL to send Rekey requests.
	rekeyTarget string
	// renewTarget is the URL to send Renew requests.
	renewTarget string
	// issuer is the URI for Issuer.
	issuer string
	// subjects are the URIs for Subjects.
	subjects []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Certificate object from the raw JSON.
func (c *Certificate) UnmarshalJSON(b []byte) error {
	type temp Certificate
	type cActions struct {
		ForceAutomaticRenew ActionTarget `json:"#Certificate.ForceAutomaticRenew"`
		Rekey               ActionTarget `json:"#Certificate.Rekey"`
		Renew               ActionTarget `json:"#Certificate.Renew"`
	}
	type cLinks struct {
		Issuer   Link  `json:"Issuer"`
		Subjects Links `json:"Subjects"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Certificate(tmp.temp)

	// Extract the links to other entities for later
	c.forceAutomaticRenewTarget = tmp.Actions.ForceAutomaticRenew.Target
	c.rekeyTarget = tmp.Actions.Rekey.Target
	c.renewTarget = tmp.Actions.Renew.Target
	c.issuer = tmp.Links.Issuer.String()
	c.subjects = tmp.Links.Subjects.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Certificate) Update() error {
	readWriteFields := []string{
		"Password",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCertificate will get a Certificate instance from the service.
func GetCertificate(c Client, uri string) (*Certificate, error) {
	return GetObject[Certificate](c, uri)
}

// ListReferencedCertificates gets the collection of Certificate from
// a provided reference.
func ListReferencedCertificates(c Client, link string) ([]*Certificate, error) {
	return GetCollectionObjects[Certificate](c, link)
}

// This action shall force an automatic renewal of the certificate, if this
// certificate is configured for automatic certificate enrollment with a
// 'CertificateEnrollment' resource. If the certificate is not configured for
// automatic certificate enrollment, the service shall reject the request and
// return the HTTP '400 Bad Request' status code.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Certificate) ForceAutomaticRenew() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(c.client,
		c.forceAutomaticRenewTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// CertificateRekeyParameters holds the parameters for the Rekey action.
type CertificateRekeyParameters struct {
	// ChallengePassword shall contain the challenge password to apply to the
	// certificate for revocation requests as defined by the RFC2985
	// 'challengePassword' attribute.
	ChallengePassword string `json:"ChallengePassword,omitempty"`
	// KeyBitLength shall contain the length of the key, in bits, if needed based
	// on the 'KeyPairAlgorithm' parameter value.
	KeyBitLength int `json:"KeyBitLength,omitempty"`
	// KeyCurveID shall contain the curve ID to use with the key, if needed based
	// on the 'KeyPairAlgorithm' parameter value. The allowable values for this
	// parameter shall be the strings in the 'Name' field of the 'TCG_ECC_CURVE
	// Constants' table, formerly the 'TPM_ECC_CURVE Constants' table, within the
	// 'Trusted Computing Group Algorithm Registry'.
	KeyCurveID string `json:"KeyCurveId,omitempty"`
	// KeyPairAlgorithm shall contain the type of key-pair for use with signing
	// algorithms. The allowable values for this parameter shall be the strings in
	// the 'Algorithm Name' field of the 'TCG_ALG_ID Constants' table, formerly the
	// 'TPM_ALG_ID Constants' table, within the 'Trusted Computing Group Algorithm
	// Registry'.
	KeyPairAlgorithm string `json:"KeyPairAlgorithm,omitempty"`
}

// This action shall use the certificate data to generate a new key-pair for a
// certificate. The response shall contain a signing request that a certificate
// authority (CA) will sign. The service should retain the private key that
// generated this request for installation of the certificate. The private key
// should not be part of the response. The private key should not be part of
// the response.
func (c *Certificate) Rekey(params *CertificateRekeyParameters) (*RekeyResponse, error) {
	resp, err := c.PostWithResponse(c.rekeyTarget, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result RekeyResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall generate a certificate signing request using the existing
// information and key-pair of the certificate. The response shall contain a
// signing request that a certificate authority (CA) will sign. The service
// should retain the private key that this request generates for when the
// certificate is installed. The private key should not be part of the
// response.
// challengePassword - This property shall contain the challenge password to
// apply to the certificate for revocation requests as defined by the RFC2985
// 'challengePassword' attribute.
func (c *Certificate) Renew(challengePassword string) (*RenewResponse, error) {
	payload := make(map[string]any)
	payload["ChallengePassword"] = challengePassword

	resp, err := c.PostWithResponse(c.renewTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result RenewResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Subjects gets the Subjects linked resources.
func (c *Certificate) Subjects() ([]*Certificate, error) {
	return GetObjects[Certificate](c.client, c.subjects)
}

// CertificateIdentifier shall contain the properties that identifies the issuer or subject
// of a certificate.
type CertificateIdentifier struct {
	// AdditionalCommonNames shall contain an array of additional common names for
	// the entity, as defined by the RFC5280 'commonName' attribute, in array order
	// as they appear in the certificate. This property shall not be present if
	// only one common name is found. The first common name shall not appear in
	// this property.
	//
	// Version added: v1.6.0
	AdditionalCommonNames []string
	// AdditionalOrganizationalUnits shall contain an array of additional
	// organizational units for the entity, as defined by the RFC5280
	// 'organizationalUnitName' attribute, in array order as they appear in the
	// certificate. This property shall not be present if only one organizational
	// unit is found. The first organizational unit shall not appear in this
	// property.
	//
	// Version added: v1.6.0
	AdditionalOrganizationalUnits []string
	// AlternativeNames shall contain the additional host names of the entity, as
	// defined by the RFC5280 'subjectAltName' attribute. This property shall not
	// be present in the 'Issuer' property.
	//
	// Version added: v1.7.0
	AlternativeNames []string
	// City shall contain the city or locality of the organization of the entity,
	// as defined by the RFC5280 'localityName' attribute.
	City string
	// CommonName shall contain the common name of the entity, as defined by the
	// RFC5280 'commonName' attribute.
	CommonName string
	// Country shall contain the two-letter ISO code for the country of the
	// organization of the entity, as defined by the RFC5280 'countryName'
	// attribute.
	Country string
	// DisplayString shall contain a display string that represents the entire
	// identifier. The string should be formatted using industry conventions, such
	// as the single-line human-readable string described by RFC2253 and preserving
	// the field order as shown in the certificate.
	//
	// Version added: v1.6.0
	DisplayString string
	// DomainComponents shall contain an array of domain component fields for the
	// entity, as defined by the RFC4519 'domainComponent' attribute, in array
	// order as they appear in the certificate.
	//
	// Version added: v1.6.0
	DomainComponents []string
	// Email shall contain the email address of the contact within the organization
	// of the entity, as defined by the RFC2985 'emailAddress' attribute.
	Email string
	// Organization shall contain the name of the organization of the entity, as
	// defined by the RFC5280 'organizationName' attribute.
	Organization string
	// OrganizationalUnit shall contain the name of the unit or division of the
	// organization of the entity, as defined by the RFC5280
	// 'organizationalUnitName' attribute.
	OrganizationalUnit string
	// State shall contain the state, province, or region of the organization of
	// the entity, as defined by the RFC5280 'stateOrProvinceName' attribute.
	State string
}

// RekeyResponse shall contain the properties found in the response body for the
// 'Rekey' action.
type RekeyResponse struct {
	// CSRString shall contain the certificate signing request as a PEM-encoded
	// string, containing structures specified by RFC2986. The private key should
	// not be part of the string.
	//
	// Version added: v1.1.0
	CSRString string
	// Certificate shall contain a link to a resource of type 'Certificate' that is
	// replaced after the certificate authority (CA) signs the certificate.
	//
	// Version added: v1.1.0
	certificate string
}

// UnmarshalJSON unmarshals a RekeyResponse object from the raw JSON.
func (r *RekeyResponse) UnmarshalJSON(b []byte) error {
	type temp RekeyResponse
	var tmp struct {
		temp
		Certificate Link `json:"Certificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = RekeyResponse(tmp.temp)

	// Extract the links to other entities for later
	r.certificate = tmp.Certificate.String()

	return nil
}

// Certificate gets the Certificate linked resource.
func (r *RekeyResponse) Certificate(client Client) (*Certificate, error) {
	if r.certificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, r.certificate)
}

// RenewResponse shall contain the properties found in the response body for the
// 'Renew' action.
type RenewResponse struct {
	// CSRString shall contain the certificate signing request as a PEM-encoded
	// string, containing structures specified by RFC2986. The private key should
	// not be part of the string.
	//
	// Version added: v1.1.0
	CSRString string
	// Certificate shall contain a link to a resource of type 'Certificate' that is
	// replaced after the certificate authority (CA) signs the certificate.
	//
	// Version added: v1.1.0
	certificate string
}

// UnmarshalJSON unmarshals a RenewResponse object from the raw JSON.
func (r *RenewResponse) UnmarshalJSON(b []byte) error {
	type temp RenewResponse
	var tmp struct {
		temp
		Certificate Link `json:"Certificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = RenewResponse(tmp.temp)

	// Extract the links to other entities for later
	r.certificate = tmp.Certificate.String()

	return nil
}

// Certificate gets the Certificate linked resource.
func (r *RenewResponse) Certificate(client Client) (*Certificate, error) {
	if r.certificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, r.certificate)
}

// SPDM shall contain SPDM-related information for a certificate.
type SPDM struct {
	// SlotID shall contain an integer between 0 and 7, inclusive, that represents
	// the slot identifier for an SPDM-provided certificate.
	//
	// Version added: v1.5.0
	SlotID *int `json:"SlotId,omitempty"`
}
