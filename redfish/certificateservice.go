//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #CertificateService.v1_2_0.CertificateService

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CertificateService shall represent the certificate service properties for a
// Redfish implementation.
type CertificateService struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// generateCSRTarget is the URL to send GenerateCSR requests.
	generateCSRTarget string
	// replaceCertificateTarget is the URL to send ReplaceCertificate requests.
	replaceCertificateTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a CertificateService object from the raw JSON.
func (c *CertificateService) UnmarshalJSON(b []byte) error {
	type temp CertificateService
	type cActions struct {
		GenerateCSR        common.ActionTarget `json:"#CertificateService.GenerateCSR"`
		ReplaceCertificate common.ActionTarget `json:"#CertificateService.ReplaceCertificate"`
	}
	var tmp struct {
		temp
		Actions                  cActions
		CertificateEnrollments   common.Link `json:"certificateEnrollments"`
		CertificateLocations     common.Link `json:"certificateLocations"`
		EnrollmentCACertificates common.Link `json:"enrollmentCACertificates"`
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

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CertificateService) Update() error {
	readWriteFields := []string{
		"AutomaticCertificateEnrollment",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
}

// GetCertificateService will get a CertificateService instance from the service.
func GetCertificateService(c common.Client, uri string) (*CertificateService, error) {
	return common.GetObject[CertificateService](c, uri)
}

// ListReferencedCertificateServices gets the collection of CertificateService from
// a provided reference.
func ListReferencedCertificateServices(c common.Client, link string) ([]*CertificateService, error) {
	return common.GetCollectionObjects[CertificateService](c, link)
}

// GenerateCSR shall make a certificate signing request. The response shall
// contain a signing request that a certificate authority (CA) will sign. The
// service should retain the private key that was generated during this request
// for installation of the certificate. The private key should not be part of
// the response.
// alternativeNames - This parameter shall contain an array of additional host
// names of the component to secure, as defined by the RFC5280 'subjectAltName'
// attribute.
// certificateCollection - This parameter shall contain a link to a resource
// collection of type 'CertificateCollection' where the certificate is
// installed after the certificate authority (CA) signs the certificate.
// challengePassword - This property shall contain the challenge password to
// apply to the certificate for revocation requests as defined by the RFC2985
// 'challengePassword' attribute.
// city - This parameter shall contain the city or locality of the organization
// making the request, as defined by the RFC5280 'localityName' attribute.
// commonName - This parameter shall contain the fully qualified domain name of
// the component to secure, as defined by the RFC5280 'commonName' attribute.
// contactPerson - This property shall contain the name of the user making the
// request, as defined by the RFC5280 'name' attribute.
// country - This parameter shall contain the two-letter ISO code for the
// country of the organization making the request, as defined by the RFC5280
// 'countryName' attribute.
// email - This parameter shall contain the email address of the contact within
// the organization making the request, as defined by the RFC2985
// 'emailAddress' attribute.
// givenName - This parameter shall contain the given name of the user making
// the request, as defined by the RFC5280 'givenName' attribute.
// initials - This parameter shall contain the initials of the user making the
// request, as defined by the RFC5280 'initials' attribute.
// keyBitLength - This parameter shall contain the length of the key, in bits,
// if needed based on the 'KeyPairAlgorithm' parameter value.
// keyCurveID - This parameter shall contain the curve ID to use with the key,
// if needed based on the 'KeyPairAlgorithm' parameter value. The allowable
// values for this parameter shall be the strings in the 'Name' field of the
// 'TCG_ECC_CURVE Constants' table, formerly the 'TPM_ECC_CURVE Constants'
// table, within the 'Trusted Computing Group Algorithm Registry'.
// keyPairAlgorithm - This parameter shall contain the type of key-pair for use
// with signing algorithms. The allowable values for this parameter shall be
// the strings in the 'Algorithm Name' field of the 'TCG_ALG_ID Constants'
// table, formerly the 'TPM_ALG_ID Constants' table, within the 'Trusted
// Computing Group Algorithm Registry'.
// keyUsage - This parameter shall contain the usage of the key contained in
// the certificate. If the client does not provide this value, the service can
// determine the appropriate key usage settings in the certificate signing
// request.
// organization - This parameter shall contain the name of the organization
// making the request, as defined by the RFC5280 'organizationName' attribute.
// organizationalUnit - This parameter shall contain the name of the unit or
// division of the organization making the request, as defined by the RFC5280
// 'organizationalUnitName' attribute.
// state - This parameter shall contain the state, province, or region of the
// organization making the request, as defined by the RFC5280
// 'stateOrProvinceName' attribute.
// surname - This parameter shall contain the surname of the user making the
// request, as defined by the RFC5280 'surname' attribute.
// unstructuredName - This property shall contain the unstructured name of the
// subject, as defined by the RFC2985 'unstructuredName' attribute.
func (c *CertificateService) GenerateCSR(alternativeNames string, certificateCollection string, challengePassword string, city string, commonName string, contactPerson string, country string, email string, givenName string, initials string, keyBitLength int, keyCurveID string, keyPairAlgorithm string, keyUsage KeyUsage, organization string, organizationalUnit string, state string, surname string, unstructuredName string) (*GenerateCSRResponse, error) {
	payload := make(map[string]any)
	payload["AlternativeNames"] = alternativeNames
	payload["CertificateCollection"] = certificateCollection
	payload["ChallengePassword"] = challengePassword
	payload["City"] = city
	payload["CommonName"] = commonName
	payload["ContactPerson"] = contactPerson
	payload["Country"] = country
	payload["Email"] = email
	payload["GivenName"] = givenName
	payload["Initials"] = initials
	payload["KeyBitLength"] = keyBitLength
	payload["KeyCurveId"] = keyCurveID
	payload["KeyPairAlgorithm"] = keyPairAlgorithm
	payload["KeyUsage"] = keyUsage
	payload["Organization"] = organization
	payload["OrganizationalUnit"] = organizationalUnit
	payload["State"] = state
	payload["Surname"] = surname
	payload["UnstructuredName"] = unstructuredName

	resp, err := c.PostWithResponse(c.generateCSRTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, common.CleanupHTTPResponse(resp)
	}

	var result GenerateCSRResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ReplaceCertificate shall replace a certificate. The 'Location' header in the
// response shall contain the URI of the new certificate resource.
// certificateString - This parameter shall contain the string of the
// certificate, and the format shall follow the requirements specified by the
// 'CertificateType' parameter value. If the certificate contains any private
// keys, they shall be removed from the string in responses. If the service
// does not know the private key for the certificate and it is needed to use
// the certificate, the client shall provide the private key as part of the
// string in the 'POST' request.
// certificateType - This parameter shall contain the format type for the
// certificate.
// certificateURI - This parameter shall contain a link to a resource of type
// 'Certificate' that is being replaced.
// password - This parameter shall contain the password for the certificate
// contained in the 'CertificateString' parameter. This parameter shall be
// required if the 'CertificateType' parameter contains 'PKCS12' and the
// client-provided certificate is password protected.
func (c *CertificateService) ReplaceCertificate(certificateString string, certificateType CertificateType, certificateURI string, password string) error {
	payload := make(map[string]any)
	payload["CertificateString"] = certificateString
	payload["CertificateType"] = certificateType
	payload["CertificateUri"] = certificateURI
	payload["Password"] = password
	return c.Post(c.replaceCertificateTarget, payload)
}

// CertificateEnrollments gets the CertificateEnrollments linked resource.
func (c *CertificateService) CertificateEnrollments(client common.Client) (*CertificateEnrollment, error) {
	if c.certificateEnrollments == "" {
		return nil, nil
	}
	return common.GetObject[CertificateEnrollment](client, c.certificateEnrollments)
}

// CertificateLocations gets the CertificateLocations linked resource.
func (c *CertificateService) CertificateLocations(client common.Client) (*CertificateLocations, error) {
	if c.certificateLocations == "" {
		return nil, nil
	}
	return common.GetObject[CertificateLocations](client, c.certificateLocations)
}

// EnrollmentCACertificates gets the EnrollmentCACertificates linked resource.
func (c *CertificateService) EnrollmentCACertificates(client common.Client) (*Certificate, error) {
	if c.enrollmentCACertificates == "" {
		return nil, nil
	}
	return common.GetObject[Certificate](client, c.enrollmentCACertificates)
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
		CertificateCollection common.Link `json:"certificateCollection"`
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
func (g *GenerateCSRResponse) CertificateCollection(client common.Client) ([]*Certificate, error) {
	if g.certificateCollection == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, g.certificateCollection)
}
