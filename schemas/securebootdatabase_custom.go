//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

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
