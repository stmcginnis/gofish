//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/SecureBootDatabase.v1_0_3.json
// 2020.1 - #SecureBootDatabase.v1_0_3.SecureBootDatabase

package schemas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type SecureBootDatabaseResetKeysType string

const (
	// ResetAllKeysToDefaultSecureBootDatabaseResetKeysType Reset the contents of this UEFI Secure
	// Boot key database to the default values.
	ResetAllKeysToDefaultSecureBootDatabaseResetKeysType SecureBootDatabaseResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysSecureBootDatabaseResetKeysType Delete the contents of this UEFI Secure Boot key
	// database.
	DeleteAllKeysSecureBootDatabaseResetKeysType SecureBootDatabaseResetKeysType = "DeleteAllKeys"
)

// SecureBootDatabase shall be used to represent a UEFI Secure Boot database for
// a Redfish implementation.
type SecureBootDatabase struct {
	Entity
	// Certificates shall be a link to a resource collection of type
	// 'CertificateCollection'.
	certificates string
	// DatabaseID shall contain the name of the UEFI Secure Boot database. This
	// property shall contain the same value as the 'Id' property. The value shall
	// be one of the UEFI-defined Secure Boot databases: 'PK', 'KEK' 'db', 'dbx',
	// 'dbr', 'dbt', 'PKDefault', 'KEKDefault', 'dbDefault', 'dbxDefault',
	// 'dbrDefault', or 'dbtDefault'.
	DatabaseID string `json:"DatabaseId"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Signatures shall be a link to a resource collection of type
	// 'SignatureCollection'.
	signatures string
	// resetKeysTarget is the URL to send ResetKeys requests.
	resetKeysTarget string
}

// UnmarshalJSON unmarshals a SecureBootDatabase object from the raw JSON.
func (s *SecureBootDatabase) UnmarshalJSON(b []byte) error {
	type temp SecureBootDatabase
	type sActions struct {
		ResetKeys ActionTarget `json:"#SecureBootDatabase.ResetKeys"`
	}
	var tmp struct {
		temp
		Actions      sActions
		Certificates Link `json:"Certificates"`
		Signatures   Link `json:"Signatures"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SecureBootDatabase(tmp.temp)

	// Extract the links to other entities for later
	s.resetKeysTarget = tmp.Actions.ResetKeys.Target
	s.certificates = tmp.Certificates.String()
	s.signatures = tmp.Signatures.String()

	return nil
}

// GetSecureBootDatabase will get a SecureBootDatabase instance from the service.
func GetSecureBootDatabase(c Client, uri string) (*SecureBootDatabase, error) {
	return GetObject[SecureBootDatabase](c, uri)
}

// ListReferencedSecureBootDatabases gets the collection of SecureBootDatabase from
// a provided reference.
func ListReferencedSecureBootDatabases(c Client, link string) ([]*SecureBootDatabase, error) {
	return GetCollectionObjects[SecureBootDatabase](c, link)
}

// This action shall perform a reset of this UEFI Secure Boot key database. The
// 'ResetAllKeysToDefault' value shall reset this UEFI Secure Boot key database
// to the default values. The 'DeleteAllKeys' value shall delete the contents
// of this UEFI Secure Boot key database.
// resetKeysType - This parameter shall specify the type of reset or delete to
// perform on this UEFI Secure Boot database.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *SecureBootDatabase) ResetKeys(resetKeysType SecureBootDatabaseResetKeysType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetKeysType"] = resetKeysType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetKeysTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Certificates gets the Certificates collection.
func (s *SecureBootDatabase) Certificates() ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](s.client, s.certificates)
}

// Signatures gets the Signatures collection.
func (s *SecureBootDatabase) Signatures() ([]*Signature, error) {
	if s.signatures == "" {
		return nil, nil
	}
	return GetCollectionObjects[Signature](s.client, s.signatures)
}

// AddCertificate enrolls a certificate in this UEFI Secure Boot database by
// posting it to the database's certificate collection.
//
// certificateString shall contain the certificate in the format indicated by
// certificateType, typically a PEM-encoded string for PEMCertificateType.
//
// uefiSignatureOwner is optional. When not empty it shall contain the GUID of
// the UEFI signature owner for this certificate as defined by the UEFI
// Specification.
//
// The created Certificate is returned when the service provides it, either in
// the response body or through the Location header. Services that provide
// neither, or that process the request asynchronously, yield a nil Certificate
// and a nil error; use Certificates to re-read the collection in that case.
func (s *SecureBootDatabase) AddCertificate(
	certificateString string,
	certificateType CertificateType,
	uefiSignatureOwner string,
) (*Certificate, error) {
	if s.certificates == "" {
		return nil, fmt.Errorf("secure boot database %q does not provide a certificate collection", s.ODataID)
	}

	payload := map[string]any{
		"CertificateString": certificateString,
		"CertificateType":   certificateType,
	}
	if uefiSignatureOwner != "" {
		payload["UefiSignatureOwner"] = uefiSignatureOwner
	}

	return postToSecureBootCollection[Certificate](s.client, s.certificates, payload)
}

// AddSignature enrolls a signature in this UEFI Secure Boot database by posting
// it to the database's signature collection. This is the hash-based counterpart
// to AddCertificate, used for databases such as 'dbx'.
//
// signatureString shall contain the signature in the format required by
// signatureTypeRegistry. For UEFISignatureTypeRegistry that is a big-endian
// hex-encoded string of the UEFI SignatureData value.
//
// signatureType shall contain the format type for the signature, qualified by
// signatureTypeRegistry. For UEFISignatureTypeRegistry that is the '#define'
// name of the EFI_SIGNATURE_LIST SignatureType member, for example
// 'EFI_CERT_SHA256_GUID'.
//
// uefiSignatureOwner is optional. When not empty it shall contain the GUID of
// the UEFI signature owner for this signature as defined by the UEFI
// Specification.
//
// The created Signature is returned when the service provides it, either in
// the response body or through the Location header. Services that provide
// neither, or that process the request asynchronously, yield a nil Signature
// and a nil error; use Signatures to re-read the collection in that case.
func (s *SecureBootDatabase) AddSignature(
	signatureString string,
	signatureTypeRegistry SignatureTypeRegistry,
	signatureType string,
	uefiSignatureOwner string,
) (*Signature, error) {
	if s.signatures == "" {
		return nil, fmt.Errorf("secure boot database %q does not provide a signature collection", s.ODataID)
	}

	payload := map[string]any{
		"SignatureString":       signatureString,
		"SignatureTypeRegistry": signatureTypeRegistry,
		"SignatureType":         signatureType,
	}
	if uefiSignatureOwner != "" {
		payload["UefiSignatureOwner"] = uefiSignatureOwner
	}

	return postToSecureBootCollection[Signature](s.client, s.signatures, payload)
}

// postToSecureBootCollection creates a new member in a UEFI Secure Boot key
// database collection and makes a best effort to return it.
//
// Implementations differ in what they return for the create: some send the new
// resource in the response body, some send an empty body with a Location
// header, and some send neither. A nil object with a nil error means the
// service did not identify what it created, not that nothing was created.
func postToSecureBootCollection[T any, PT GenericSchemaObjectPointer[T]](
	c Client,
	uri string,
	payload any,
) (*T, error) {
	resp, taskInfo, err := PostWithTask(c, uri, payload, nil, false)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	// The service accepted the request but is doing the work asynchronously,
	// so there is nothing to return yet.
	if taskInfo != nil {
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// A representation of the created resource is the preferred response, but
	// an empty body is common, in which case json.Unmarshal fails and we fall
	// through to the Location header.
	entity := PT(new(T))
	if json.Unmarshal(body, entity) == nil && entity.GetODataID() != "" {
		if etag := resp.Header.Get("Etag"); etag != "" && entity.GetETag() == "" {
			entity.SetETag(sanitizeETag(etag))
		}
		entity.SetClient(c)
		return (*T)(entity), nil
	}

	location := resp.Header.Get("Location")
	if location == "" {
		return nil, nil
	}

	// The header may hold an absolute URL while the client expects a path.
	if parsed, err := url.ParseRequestURI(location); err == nil {
		location = parsed.RequestURI()
	}

	return GetObject[T, PT](c, location)
}
