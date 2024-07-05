//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CertificateService shall represent the certificate service properties for a Redfish implementation.
type CertificateService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CertificateLocations shall contain a link to a resource of type CertificateLocations.
	certificateLocations string
	// Description provides a description of this resource.
	Description              string
	generateCSRTarget        string
	replaceCertificateTarget string
}

// UnmarshalJSON unmarshals a CertificateService object from the raw JSON.
func (certificateservice *CertificateService) UnmarshalJSON(b []byte) error {
	type temp CertificateService
	type linkReference struct {
		AssociatedControls      common.Links
		AssociatedControlsCount int `json:"AssociatedControls@odata.count"`
	}
	type actions struct {
		GenerateCSR        common.ActionTarget `json:"#CertificateService.GenerateCSR"`
		ReplaceCertificate common.ActionTarget `json:"#CertificateService.ReplaceCertificate"`
	}
	var t struct {
		temp
		CertificateLocations common.Link
		Actions              actions
		Links                linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*certificateservice = CertificateService(t.temp)

	// Extract the links to other entities for later
	certificateservice.certificateLocations = t.CertificateLocations.String()
	certificateservice.generateCSRTarget = t.Actions.GenerateCSR.Target
	certificateservice.replaceCertificateTarget = t.Actions.ReplaceCertificate.Target

	return nil
}

// CertificateLocations get the certificate locations.
func (certificateservice *CertificateService) CertificateLocations() (*CertificateLocations, error) {
	if certificateservice.certificateLocations == "" {
		return nil, nil
	}
	return GetCertificateLocations(certificateservice.GetClient(), certificateservice.certificateLocations)
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

// GenerateCSRResponse shall contain the properties found in the response body for the GenerateCSR action.
type GenerateCSRResponse struct {
	// CSRString shall contain the Privacy Enhanced Mail (PEM)-encoded string, which contains RFC2986-specified
	// structures, of the certificate signing request. The private key should not be part of the string.
	CSRString string
	// CertificateCollection shall contain a link to a resource collection of type CertificateCollection where the
	// certificate is installed after the certificate authority (CA) has signed the certificate.
	certificateCollection string
	client                common.Client
}

// setClient provides a reference to the API client for being able to make future calls.
func (generatecsrresponse *GenerateCSRResponse) setClient(c common.Client) {
	generatecsrresponse.client = c
}

// UnmarshalJSON unmarshals a GenerateCSRResponse object from the raw JSON.
func (generatecsrresponse *GenerateCSRResponse) UnmarshalJSON(b []byte) error {
	type temp GenerateCSRResponse
	var t struct {
		temp
		CertificateCollection common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*generatecsrresponse = GenerateCSRResponse(t.temp)

	// Extract the links to other entities for later
	generatecsrresponse.certificateCollection = t.CertificateCollection.String()

	return nil
}

// Certificates gets the collection of where the certificate is installed after
// the certificate authority (CA) has signed the certificate.
func (generatecsrresponse *GenerateCSRResponse) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(generatecsrresponse.client, generatecsrresponse.certificateCollection)
}

type GenerateCSRRequest struct {
	// AlternativeNames is an array of additional host names of the component to secure,
	// as defined by the RFC5280 'subjectAltName' attribute.
	AlternativeNames []string `json:",omitempty"`
	// CertificateCollection is a link to a resource collection of type CertificateCollection where the certificate is
	// installed after the certificate authority (CA) signs the certificate.
	CertificateCollection string
	// ChallengePassword is the challenge password to apply to the certificate for revocation requests as defined by
	// the RFC2985 'challengePassword' attribute.
	ChallengePassword string `json:",omitempty"`
	// City is the city or locality of the organization making the request, as defined by the RFC5280 'localityName' attribute.
	City string
	// CommonName is the fully qualified domain name of the component to secure, as defined by the RFC5280 'commonName' attribute.
	CommonName string
	// ContactPerson is the name of the user making the request, as defined by the RFC5280 'name' attribute.
	ContactPerson string `json:",omitempty"`
	// Country is the two-letter ISO code for the country of the organization making the request, as defined by the RFC5280 'countryName' attribute.
	Country string
	// Email is the email address of the contact within the organization making the request, as defined by the RFC2985 'emailAddress' attribute.
	Email string `json:",omitempty"`
	// GivenName is the given name of the user making the request, as defined by the RFC5280 'givenName' attribute.
	GivenName string `json:",omitempty"`
	// Initials is the initials of the user making the request, as defined by the RFC5280 'initials' attribute.
	Initials string `json:",omitempty"`
	// KeyBitLength is the length of the key, in bits, if needed based on the KeyPairAlgorithm parameter value.
	KeyBitLength int `json:",omitempty"`
	// KeyCurveID is the curve ID to use with the key, if needed based on the KeyPairAlgorithm parameter value.
	// The allowable values for this parameter shall be the strings in the 'Name' field of the 'TPM_ECC_CURVE Constants'
	// table within the 'Trusted Computing Group Algorithm Registry'.
	KeyCurveID string `json:"KeyCurveId,omitempty"`
	// KeyPairAlgorithm is the type of key-pair for use with signing algorithms. The allowable values for this parameter
	// shall be the strings in the 'Algorithm Name' field of the 'TPM_ALG_ID Constants' table within the
	// 'Trusted Computing Group Algorithm Registry'.
	KeyPairAlgorithm string `json:",omitempty"`
	// KeyUsage is the usage of the key contained in the certificate. If the client does not provide this value, the
	// service can determine the appropriate key usage settings in the certificate signing request.
	KeyUsage []KeyUsageExtension `json:",omitempty"`
	// Organization is the name of the organization making the request, as defined by the RFC5280 'organizationName' attribute.
	Organization string
	// OrganizationalUnit is the name of the unit or division of the organization making the request, as defined by the
	// RFC5280 'organizationalUnitName' attribute.
	OrganizationalUnit string
	// State is the state, province, or region of the organization making the request, as defined by the RFC5280
	// 'stateOrProvinceName' attribute.
	State string
	// Surname is the surname of the user making the request, as defined by the RFC5280 'surname' attribute.
	Surname string `json:",omitempty"`
	// UnstructuredName is the unstructured name of the subject, as defined by the RFC2985 'unstructuredName' attribute.
	UnstructuredName string `json:",omitempty"`
}

// GenerateCSR makes a certificate signing request. The response shall contain a signing request that a certificate
// authority (CA) will sign. The service should retain the private key that was generated during this request for
// installation of the certificate.
// WARNING: this has not been fully tested and is subject to change.
func (certificateservice *CertificateService) GenerateCSR(request *GenerateCSRRequest) (*GenerateCSRResponse, error) {
	resp, err := certificateservice.PostWithResponse(certificateservice.generateCSRTarget, request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var csrResponse GenerateCSRResponse
	err = json.NewDecoder(resp.Body).Decode(&csrResponse)
	if err != nil {
		return nil, err
	}

	csrResponse.setClient(certificateservice.GetClient())
	return &csrResponse, nil
}

// ReplaceCertificate replaces a certificate.
//
// `certificateString` is the string of the certificate, and the format shall follow the requirements specified by the
// CertificateType property value. If the service does not know the private key for the certificate and it is needed to
// use the certificate, the client shall provide the private key as part of the string in the request.
//
// `certificateType` specifies the format type of the certificate.
//
// `certificateURI` is a link to a resource of type Certificate that is being replaced.
// WARNING: this has not been fully tested.
func (certificateservice *CertificateService) ReplaceCertificate(certificateString string, certificateType CertificateType, certificateURI string) error {
	// TODO: The new certificate resource is returned in the `Location` header of the response.
	// Need to rework to be able to extract this header to use the URI to get the certificate.
	// If we actually need it, that is.
	payload := struct {
		CertificateString string
		CertificateType   CertificateType
		CertificateURI    string `json:"CertificateUri"`
	}{
		certificateString,
		certificateType,
		certificateURI,
	}

	return certificateservice.Post(certificateservice.replaceCertificateTarget, payload)
}
