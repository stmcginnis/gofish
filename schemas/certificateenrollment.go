//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/CertificateEnrollment.v1_0_1.json
// 2025.3 - #CertificateEnrollment.v1_0_1.CertificateEnrollment

package schemas

import (
	"encoding/json"
)

type ACMEChallengeType string

const (
	// HTTP01ACMEChallengeType shall indicate the RFC8555-defined http-01 challenge
	// type for domain validation.
	HTTP01ACMEChallengeType ACMEChallengeType = "Http01"
	// DNS01ACMEChallengeType shall indicate the RFC8555-defined dns-01 challenge
	// type for domain validation.
	DNS01ACMEChallengeType ACMEChallengeType = "Dns01"
)

type EnrollmentProtocolType string

const (
	// ACMEEnrollmentProtocolType shall indicate the enrollment uses the Automatic
	// Certificate Management Environment (ACME) protocol as defined by RFC8555.
	ACMEEnrollmentProtocolType EnrollmentProtocolType = "ACME"
	// SCEPEnrollmentProtocolType shall indicate the enrollment uses the Simple
	// Certificate Enrollment Protocol (SCEP) protocol as defined by RFC8894.
	SCEPEnrollmentProtocolType EnrollmentProtocolType = "SCEP"
	// OEMEnrollmentProtocolType shall indicate the OEM enrollment protocol type.
	OEMEnrollmentProtocolType EnrollmentProtocolType = "OEM"
)

type LastOperationType string

const (
	// RenewLastOperationType Certificate renewal operation.
	RenewLastOperationType LastOperationType = "Renew"
	// UpdateAcmeEmailLastOperationType Update ACME email operation. Applicable
	// only when the enrollment protocol is ACME.
	UpdateAcmeEmailLastOperationType LastOperationType = "UpdateAcmeEmail"
)

type OperationStatus string

const (
	// SuccessOperationStatus The operation completed successfully.
	SuccessOperationStatus OperationStatus = "Success"
	// FailedOperationStatus The operation failed.
	FailedOperationStatus OperationStatus = "Failed"
	// InProgressOperationStatus The operation is in progress.
	InProgressOperationStatus OperationStatus = "InProgress"
	// UnknownOperationStatus The operation status is unknown.
	UnknownOperationStatus OperationStatus = "Unknown"
)

// CertificateEnrollment The 'CertificateEnrollment' schema describes an
// automatic certificate enrollment for a specific protocol such as ACME
// (Automatic Certificate Management Environment) or SCEP (Simple Certificate
// Enrollment Protocol).
type CertificateEnrollment struct {
	Entity
	// ACME shall contain configuration specific to the ACME protocol. This
	// property shall only be present when the 'EnrollmentType' property contains
	// 'ACME'.
	ACME ACMEConfiguration
	// CSRParameters shall contain the parameters used for generating the
	// certificate signing request.
	CSRParameters CSRParameters
	// Enabled shall indicate whether this automatic certificate enrollment is
	// enabled. If 'true', the implementation shall automatically enroll and renew
	// certificates according to the configuration. If 'false', the implementation
	// shall not perform automatic certificate enrollment operations. If this
	// property is not specified by the client in the create request, it shall be
	// assumed to be 'false'.
	Enabled bool
	// EnrollmentState shall contain the status information for this enrollment
	// including the last operation performed and its status.
	EnrollmentState EnrollmentState
	// EnrollmentType shall contain the configured automatic certificate enrollment
	// protocol.
	EnrollmentType EnrollmentProtocolType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RenewBeforeExpiryDays shall contain the number of days before certificate
	// expiry to begin automatic renewal of the certificate.
	RenewBeforeExpiryDays *uint `json:",omitempty"`
	// SCEP shall contain configuration specific to the SCEP protocol. This
	// property shall only be present when the 'EnrollmentType' property contains
	// 'SCEP'.
	SCEP SCEPConfiguration
	// ServerURI shall contain the URI of the certificate enrollment server that
	// provides the automatic enrollment service.
	ServerURI string
	// VerifyCertificate shall indicate whether the service will verify the
	// certificate of the server referenced by the 'ServerURI' property with the
	// certificates found in the collection referenced by the 'Certificates'
	// property. If this property is not supported by the service or specified by
	// the client in the create request, it shall be assumed to be 'false'.
	// Regardless of the value of this property, services may perform additional
	// verification based on other factors, such as the configuration of the
	// 'SecurityPolicy' resource.
	VerifyCertificate bool
	// cACertificates are the URIs for CACertificates.
	cACertificates []string
	// enrolledCertificate is the URI for EnrolledCertificate.
	enrolledCertificate string
	// enrolledCertificateLocation is the URI for EnrolledCertificateLocation.
	enrolledCertificateLocation string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a CertificateEnrollment object from the raw JSON.
func (c *CertificateEnrollment) UnmarshalJSON(b []byte) error {
	type temp CertificateEnrollment
	type cLinks struct {
		CACertificates              Links `json:"CACertificates"`
		EnrolledCertificate         Link  `json:"EnrolledCertificate"`
		EnrolledCertificateLocation Link  `json:"EnrolledCertificateLocation"`
	}
	var tmp struct {
		temp
		Links cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CertificateEnrollment(tmp.temp)

	// Extract the links to other entities for later
	c.cACertificates = tmp.Links.CACertificates.ToStrings()
	c.enrolledCertificate = tmp.Links.EnrolledCertificate.String()
	c.enrolledCertificateLocation = tmp.Links.EnrolledCertificateLocation.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CertificateEnrollment) Update() error {
	readWriteFields := []string{
		"Enabled",
		"RenewBeforeExpiryDays",
		"ServerURI",
		"VerifyCertificate",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCertificateEnrollment will get a CertificateEnrollment instance from the service.
func GetCertificateEnrollment(c Client, uri string) (*CertificateEnrollment, error) {
	return GetObject[CertificateEnrollment](c, uri)
}

// ListReferencedCertificateEnrollments gets the collection of CertificateEnrollment from
// a provided reference.
func ListReferencedCertificateEnrollments(c Client, link string) ([]*CertificateEnrollment, error) {
	return GetCollectionObjects[CertificateEnrollment](c, link)
}

// CACertificates gets the CACertificates linked resources.
func (c *CertificateEnrollment) CACertificates() ([]*Certificate, error) {
	return GetObjects[Certificate](c.client, c.cACertificates)
}

// EnrolledCertificate gets the EnrolledCertificate linked resource.
func (c *CertificateEnrollment) EnrolledCertificate() (*Certificate, error) {
	if c.enrolledCertificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](c.client, c.enrolledCertificate)
}

// EnrolledCertificateLocation gets the EnrolledCertificateLocation linked resource.
func (c *CertificateEnrollment) EnrolledCertificateLocation() (*Certificate, error) {
	if c.enrolledCertificateLocation == "" {
		return nil, nil
	}
	return GetObject[Certificate](c.client, c.enrolledCertificateLocation)
}

// ACMEConfiguration shall contain configuration specific to the ACME protocol.
type ACMEConfiguration struct {
	// ChallengeType shall contain the ACME challenge type used for domain
	// validation during automatic certificate enrollment.
	ChallengeType ACMEChallengeType
	// EABKey shall contain a Base64-encoded string, with padding characters, of
	// the external account binding (EAB) key value used for ACME account
	// registration with certificate authorities that require EAB. This property
	// shall be 'null' in responses.
	EABKey string
	// EABKeyID shall contain the external account binding (EAB) key identifier
	// used for ACME account registration with certificate authorities that require
	// EAB. This property shall be 'null' in responses.
	EABKeyID string `json:"EABKeyId"`
	// Email shall contain the email address used for ACME account registration and
	// notifications.
	Email string
}

// CSRParameters shall contain the parameters for generating a certificate
// signing request.
type CSRParameters struct {
	// AlternativeNames shall contain an array of additional host names of the
	// component to secure, as defined by the RFC5280 'subjectAltName' attribute.
	AlternativeNames []string
	// ChallengePassword shall contain the challenge password to apply to the
	// certificate for revocation requests as defined by the RFC2985
	// 'challengePassword' attribute.
	ChallengePassword string
	// City shall contain the city or locality of the organization making the
	// request, as defined by the RFC5280 'localityName' attribute.
	City string
	// CommonName shall contain the of the component to secure, as defined by the
	// RFC5280 'commonName' attribute.
	CommonName string
	// ContactPerson shall contain the name of the user making the request, as
	// defined by the RFC5280 'name' attribute.
	ContactPerson string
	// Country shall contain the two-letter ISO code for the country of the
	// organization making the request, as defined by the RFC5280 'countryName'
	// attribute.
	Country string
	// Email shall contain the email address of the contact within the organization
	// making the request, as defined by the RFC2985 'emailAddress' attribute.
	Email string
	// GivenName shall contain the given name of the user making the request, as
	// defined by the RFC5280 'givenName' attribute.
	GivenName string
	// Initials shall contain the initials of the user making the request, as
	// defined by the RFC5280 'initials' attribute.
	Initials string
	// KeyBitLength shall contain the length of the key, in bits, if needed based
	// on the 'KeyPairAlgorithm' property value.
	KeyBitLength *int `json:",omitempty"`
	// KeyCurveID shall contain the curve ID to use with the key, if needed based
	// on the 'KeyPairAlgorithm' property value. The allowable values for this
	// property shall be the strings in the 'Name' field of the 'TPM_ECC_CURVE
	// Constants' table within the 'Trusted Computing Group Algorithm Registry'.
	KeyCurveID string `json:"KeyCurveId"`
	// KeyPairAlgorithm shall contain the type of key-pair for use with signing
	// algorithms. The allowable values for this property shall be the strings in
	// the 'Algorithm Name' field of the 'TPM_ALG_ID Constants' table within the
	// 'Trusted Computing Group Algorithm Registry'.
	KeyPairAlgorithm string
	// KeyUsage shall contain the usage of the key contained in the certificate. If
	// the client does not provide this value, the service can determine the
	// appropriate key usage settings in the certificate signing request.
	KeyUsage []KeyUsage
	// Organization shall contain the name of the organization making the request,
	// as defined by the RFC5280 'organizationName' attribute.
	Organization string
	// OrganizationalUnit shall contain the name of the unit or division of the
	// organization making the request, as defined by the RFC5280
	// 'organizationalUnitName' attribute.
	OrganizationalUnit string
	// State shall contain the state, province, or region of the organization
	// making the request, as defined by the RFC5280 'stateOrProvinceName'
	// attribute.
	State string
	// Surname shall contain the surname of the user making the request, as defined
	// by the RFC5280 'surname' attribute.
	Surname string
	// UnstructuredName shall contain the unstructured name of the subject, as
	// defined by the RFC2985 'unstructuredName' attribute.
	UnstructuredName string
}

// EnrollmentState shall contain the status information for an enrollment
// including the last operation performed and its status.
type EnrollmentState struct {
	// LastOperation shall contain the last operation performed by the automatic
	// enrollment service.
	LastOperation LastOperationType
	// LastOperationStatus shall describe the status of the last operation
	// performed by automatic enrollment service.
	LastOperationStatus OperationStatus
	// LastOperationTime shall contain the date and time when the last operation
	// was performed by the automatic enrollment service.
	LastOperationTime string
}

// SCEPConfiguration shall contain configuration specific to the SCEP protocol.
type SCEPConfiguration struct {
	// ChallengePassword shall contain the challenge password used for SCEP
	// enrollment. This property shall be 'null' in responses.
	ChallengePassword string
}
