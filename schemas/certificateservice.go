//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #CertificateService.v1_2_1.CertificateService

package schemas

import (
	"encoding/json"
)

// CertificateService shall represent the certificate service properties for a
// Redfish implementation.
type CertificateService struct {
	Entity
	// AutomaticCertificateEnrollment shall contain the configuration and status of
	// automatic certificate enrollment.
	//
	// Version added: v1.2.0
	AutomaticCertificateEnrollment AutomaticCertificateEnrollment
	// CertificateEnrollments shall contain a link to a resource collection of type
	// 'CertificateEnrollmentCollection' that contains the certificate enrollment
	// configurations for this service.
	//
	// Version added: v1.2.0
	certificateEnrollments string
	// CertificateLocations shall contain a link to a resource of type
	// 'CertificateLocations'.
	certificateLocations string
	// EnrollmentCACertificates shall contain a link to a resource collection of
	// type 'CertificateCollection' that contains the server certificates for the
	// automatic certificate enrollment servers.
	//
	// Version added: v1.2.0
	enrollmentCACertificates string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// generateCSRTarget is the URL to send GenerateCSR requests.
	generateCSRTarget string
	// replaceCertificateTarget is the URL to send ReplaceCertificate requests.
	replaceCertificateTarget string
}

// UnmarshalJSON unmarshals a CertificateService object from the raw JSON.
func (c *CertificateService) UnmarshalJSON(b []byte) error {
	type temp CertificateService
	type cActions struct {
		GenerateCSR        ActionTarget `json:"#CertificateService.GenerateCSR"`
		ReplaceCertificate ActionTarget `json:"#CertificateService.ReplaceCertificate"`
	}
	var tmp struct {
		temp
		Actions                  cActions
		CertificateEnrollments   Link `json:"CertificateEnrollments"`
		CertificateLocations     Link `json:"CertificateLocations"`
		EnrollmentCACertificates Link `json:"EnrollmentCACertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CertificateService(tmp.temp)

	// Extract the links to other entities for later
	c.generateCSRTarget = tmp.Actions.GenerateCSR.Target
	c.replaceCertificateTarget = tmp.Actions.ReplaceCertificate.Target
	c.certificateEnrollments = tmp.CertificateEnrollments.String()
	c.certificateLocations = tmp.CertificateLocations.String()
	c.enrollmentCACertificates = tmp.EnrollmentCACertificates.String()

	return nil
}

// GetCertificateService will get a CertificateService instance from the service.
func GetCertificateService(c Client, uri string) (*CertificateService, error) {
	return GetObject[CertificateService](c, uri)
}

// ListReferencedCertificateServices gets the collection of CertificateService from
// a provided reference.
func ListReferencedCertificateServices(c Client, link string) ([]*CertificateService, error) {
	return GetCollectionObjects[CertificateService](c, link)
}

// CertificateServiceGenerateCSRParameters holds the parameters for the GenerateCSR action.
type CertificateServiceGenerateCSRParameters struct {
	// AlternativeNames shall contain an array of additional host names of the
	// component to secure, as defined by the RFC5280 'subjectAltName' attribute.
	AlternativeNames []string `json:"AlternativeNames,omitempty"`
	// CertificateCollection shall contain a link to a resource collection of type
	// 'CertificateCollection' where the certificate is installed after the
	// certificate authority (CA) signs the certificate.
	CertificateCollection string `json:"CertificateCollection,omitempty"`
	// ChallengePassword shall contain the challenge password to apply to the
	// certificate for revocation requests as defined by the RFC2985
	// 'challengePassword' attribute.
	ChallengePassword string `json:"ChallengePassword,omitempty"`
	// City shall contain the city or locality of the organization making the
	// request, as defined by the RFC5280 'localityName' attribute.
	City string `json:"City,omitempty"`
	// CommonName shall contain the of the component to secure, as defined by the
	// RFC5280 'commonName' attribute.
	CommonName string `json:"CommonName,omitempty"`
	// ContactPerson shall contain the name of the user making the request, as
	// defined by the RFC5280 'name' attribute.
	ContactPerson string `json:"ContactPerson,omitempty"`
	// Country shall contain the two-letter ISO code for the country of the
	// organization making the request, as defined by the RFC5280 'countryName'
	// attribute.
	Country string `json:"Country,omitempty"`
	// Email shall contain the email address of the contact within the organization
	// making the request, as defined by the RFC2985 'emailAddress' attribute.
	Email string `json:"Email,omitempty"`
	// GivenName shall contain the given name of the user making the request, as
	// defined by the RFC5280 'givenName' attribute.
	GivenName string `json:"GivenName,omitempty"`
	// Initials shall contain the initials of the user making the request, as
	// defined by the RFC5280 'initials' attribute.
	Initials string `json:"Initials,omitempty"`
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
	// KeyUsage shall contain the usage of the key contained in the certificate. If
	// the client does not provide this value, the service can determine the
	// appropriate key usage settings in the certificate signing request.
	KeyUsage []KeyUsage `json:"KeyUsage,omitempty"`
	// Organization shall contain the name of the organization making the request,
	// as defined by the RFC5280 'organizationName' attribute.
	Organization string `json:"Organization,omitempty"`
	// OrganizationalUnit shall contain the name of the unit or division of the
	// organization making the request, as defined by the RFC5280
	// 'organizationalUnitName' attribute.
	OrganizationalUnit string `json:"OrganizationalUnit,omitempty"`
	// State shall contain the state, province, or region of the organization
	// making the request, as defined by the RFC5280 'stateOrProvinceName'
	// attribute.
	State string `json:"State,omitempty"`
	// Surname shall contain the surname of the user making the request, as defined
	// by the RFC5280 'surname' attribute.
	Surname string `json:"Surname,omitempty"`
	// UnstructuredName shall contain the unstructured name of the subject, as
	// defined by the RFC2985 'unstructuredName' attribute.
	UnstructuredName string `json:"UnstructuredName,omitempty"`
}

// This action shall make a certificate signing request. The response shall
// contain a signing request that a certificate authority (CA) will sign. The
// service should retain the private key that was generated during this request
// for installation of the certificate. The private key should not be part of
// the response.
func (c *CertificateService) GenerateCSR(params *CertificateServiceGenerateCSRParameters) (*GenerateCSRResponse, error) {
	resp, err := c.PostWithResponse(c.generateCSRTarget, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result GenerateCSRResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CertificateServiceReplaceCertificateParameters holds the parameters for the ReplaceCertificate action.
type CertificateServiceReplaceCertificateParameters struct {
	// CertificateString shall contain the string of the certificate, and the
	// format shall follow the requirements specified by the 'CertificateType'
	// parameter value. If the certificate contains any private keys, they shall be
	// removed from the string in responses. If the service does not know the
	// private key for the certificate and it is needed to use the certificate, the
	// client shall provide the private key as part of the string in the 'POST'
	// request.
	CertificateString string `json:"CertificateString,omitempty"`
	// CertificateType shall contain the format type for the certificate.
	CertificateType CertificateType `json:"CertificateType,omitempty"`
	// CertificateURI shall contain a link to a resource of type 'Certificate' that
	// is being replaced.
	CertificateURI string `json:"CertificateUri,omitempty"`
	// Password shall contain the password for the certificate contained in the
	// 'CertificateString' parameter. This parameter shall be required if the
	// 'CertificateType' parameter contains 'PKCS12' and the client-provided
	// certificate is password protected.
	Password string `json:"Password,omitempty"`
}

// This action shall replace a certificate. The 'Location' header in the
// response shall contain the URI of the new certificate resource.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CertificateService) ReplaceCertificate(params *CertificateServiceReplaceCertificateParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(c.client,
		c.replaceCertificateTarget, params, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// CertificateEnrollments gets the CertificateEnrollments linked resource.
func (c *CertificateService) CertificateEnrollments() (*CertificateEnrollment, error) {
	if c.certificateEnrollments == "" {
		return nil, nil
	}
	return GetObject[CertificateEnrollment](c.client, c.certificateEnrollments)
}

// CertificateLocations gets the CertificateLocations linked resource.
func (c *CertificateService) CertificateLocations() (*CertificateLocations, error) {
	if c.certificateLocations == "" {
		return nil, nil
	}
	return GetObject[CertificateLocations](c.client, c.certificateLocations)
}

// EnrollmentCACertificates gets the EnrollmentCACertificates linked resource.
func (c *CertificateService) EnrollmentCACertificates() (*Certificate, error) {
	if c.enrollmentCACertificates == "" {
		return nil, nil
	}
	return GetObject[Certificate](c.client, c.enrollmentCACertificates)
}

// AutomaticCertificateEnrollment shall contain the configuration and status of
// automatic certificate enrollment.
type AutomaticCertificateEnrollment struct {
	// CertificatesSupported shall contain an array of certificate usage types that
	// support automatic enrollments for this service.
	//
	// Version added: v1.2.0
	CertificatesSupported []CertificateUsageType
	// EnrollmentTypes shall contain an array of automatic enrollment protocols
	// supported by this service.
	//
	// Version added: v1.2.0
	EnrollmentTypes []EnrollmentProtocolType
	// ServiceEnabled shall indicate whether automatic certificate enrollment is
	// enabled.
	//
	// Version added: v1.2.0
	ServiceEnabled bool
}

// GenerateCSRResponse shall contain the properties found in the response body
// for the 'GenerateCSR' action.
type GenerateCSRResponse struct {
	// CSRString shall contain the Privacy Enhanced Mail (PEM)-encoded string,
	// which contains RFC2986-specified structures, of the certificate signing
	// request. The private key should not be part of the string.
	CSRString string
	// CertificateCollection shall contain a link to a resource collection of type
	// 'CertificateCollection' where the certificate is installed after the
	// certificate authority (CA) has signed the certificate.
	certificateCollection string
}

// UnmarshalJSON unmarshals a GenerateCSRResponse object from the raw JSON.
func (g *GenerateCSRResponse) UnmarshalJSON(b []byte) error {
	type temp GenerateCSRResponse
	var tmp struct {
		temp
		CertificateCollection Link `json:"CertificateCollection"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = GenerateCSRResponse(tmp.temp)

	// Extract the links to other entities for later
	g.certificateCollection = tmp.CertificateCollection.String()

	return nil
}

// CertificateCollection gets the CertificateCollection collection.
func (g *GenerateCSRResponse) CertificateCollection(client Client) ([]*Certificate, error) {
	if g.certificateCollection == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, g.certificateCollection)
}
