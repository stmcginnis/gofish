//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type CertificateType string

const (
	// A Privacy Enhanced Mail (PEM)-encoded single certificate.
	PEMCertificateType CertificateType = "PEM"
	// A Privacy Enhanced Mail (PEM)-encoded certificate chain.
	PEMChainCertificateType CertificateType = "PEMChain"
	// A Privacy Enhanced Mail (PEM)-encoded PKCS7 certificate.
	PKCS7CertificateType CertificateType = "PKCS7"
)

type CertificateUsageType string

const (
	// This certificate is a BIOS certificate like those associated with UEFI.
	BIOSCertificateUsageType CertificateUsageType = "BIOS"
	// This certificate is a device type certificate like those associated with SPDM and other standards.
	DeviceCertificateUsageType CertificateUsageType = "Device"
	// This certificate is a platform type certificate like those associated with SPDM and other standards.
	PlatformCertificateUsageType CertificateUsageType = "Platform"
	// This certificate is used for SSH.
	SSHCertificateUsageType CertificateUsageType = "SSH"
	// This certificate is a user certificate like those associated with a manager account.
	UserCertificateUsageType CertificateUsageType = "User"
	// This certificate is a web or HTTPS certificate like those used for event destinations.
	WebCertificateUsageType CertificateUsageType = "Web"
)

type KeyUsageExtension string

const (
	// TLS WWW client authentication.
	ClientAuthenticationKeyUsageExtension KeyUsageExtension = "ClientAuthentication"
	// Signs downloadable executable code.
	CodeSigningKeyUsageExtension KeyUsageExtension = "CodeSigning"
	// Verifies signatures on certificate revocation lists (CRLs).
	CRLSigningKeyUsageExtension KeyUsageExtension = "CRLSigning"
	// Directly enciphers raw user data without an intermediate symmetric cipher.
	DataEnciphermentKeyUsageExtension KeyUsageExtension = "DataEncipherment"
	// Deciphers data while performing a key agreement.
	DecipherOnlyKeyUsageExtension KeyUsageExtension = "DecipherOnly"
	// Verifies digital signatures, other than signatures on certificates and CRLs.
	DigitalSignatureKeyUsageExtension KeyUsageExtension = "DigitalSignature"
	// Email protection.
	EmailProtectionKeyUsageExtension KeyUsageExtension = "EmailProtection"
	// Enciphers data while performing a key agreement.
	EncipherOnlyKeyUsageExtension KeyUsageExtension = "EncipherOnly"
	// Key agreement.
	KeyAgreementKeyUsageExtension KeyUsageExtension = "KeyAgreement"
	// Verifies signatures on public key certificates.
	KeyCertSignKeyUsageExtension KeyUsageExtension = "KeyCertSign"
	// Enciphers private or secret keys.
	KeyEnciphermentKeyUsageExtension KeyUsageExtension = "KeyEncipherment"
	// Verifies digital signatures, other than signatures on certificates and CRLs,
	// and provides a non-repudiation service that protects against the signing entity falsely denying some action.
	NonRepudiationKeyUsageExtension KeyUsageExtension = "NonRepudiation"
	// Signs OCSP responses.
	OCSPSigningKeyUsageExtension KeyUsageExtension = "OCSPSigning"
	// TLS WWW server authentication.
	ServerAuthenticationKeyUsageExtension KeyUsageExtension = "ServerAuthentication"
	// Binds the hash of an object to a time.
	TimestampingKeyUsageExtension KeyUsageExtension = "Timestamping"
)

type SPDM struct {
	// Slot identifier of the certificate.
	SlotID int64 `json:"SlotId"`
}

type CertificateIdentifier struct {
	// Additional common names of the entity.
	AdditionalCommonNames []string
	// Additional organizational units of the entity.
	AdditionalOrganizationalUnits []string
	// The additional host names of the entity.
	AlternativeNames []string
	// The city or locality of the organization of the entity.
	City string
	// The common name of the entity.
	CommonName string
	// The country of the organization of the entity.
	Country string
	// A human-readable string for this identifier.
	DisplayString string
	// The domain components of the entity.
	DomainComponents []string
	// The email address of the contact within the organization of the entity.
	Email string
	// The name of the organization of the entity.
	Organization string
	// The name of the unit or division of the organization of the entity.
	OrganizationalUnit string
	// The state, province, or region of the organization of the entity.
	State string
}

type Certificate struct {
	common.Entity

	Description string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// The string for the certificate.
	CertificateString string
	// The format of the certificate.
	CertificateType CertificateType
	// The types or purposes for this certificate.
	CertificateUsageTypes []CertificateUsageType
	// The fingerprint of the certificate.
	Fingerprint string
	// The hash algorithm for the fingerprint of the certificate.
	FingerprintHashAlgorithm string
	// The issuer of the certificate.
	Issuer CertificateIdentifier
	// The usages of a key contained within a certificate.
	KeyUsage []KeyUsageExtension
	// The serial number of the certificate.
	SerialNumber string
	// The algorithm used for creating the signature of the certificate.
	SignatureAlgorithm string
	// SPDM-related information for the certificate.
	SPDM SPDM
	// The subject of the certificate.
	Subject CertificateIdentifier
	// The UEFI signature owner for this certificate.
	UefiSignatureOwner string
	// The date when the certificate is no longer valid.
	ValidNotAfter string
	// The date when the certificate becomes valid.
	ValidNotBefore string
	Oem            json.RawMessage

	// A link to the certificate of the CA that issued this certificate.
	issuer string
	// An array of links to certificates that were issued by the CA that is represented by this certificate.
	subjects      []string
	SubjectsCount int
	OemLinks      json.RawMessage

	rekeyTarget string
	renewTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (certificate *Certificate) UnmarshalJSON(b []byte) error {
	type temp Certificate
	type linkReference struct {
		Issuer        common.Link
		Subjects      common.Links
		SubjectsCount int `json:"Subjects@odata.count"`
		Oem           json.RawMessage
	}
	type actions struct {
		RekeyCertificate struct {
			Target string
		} `json:"#Certificate.Rekey"`
		RenewCertificate struct {
			Target string
		} `json:"#Certificate.Renew"`
		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		Links   linkReference
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*certificate = Certificate(t.temp)
	certificate.issuer = t.Links.Issuer.String()
	certificate.subjects = t.Links.Subjects.ToStrings()
	certificate.SubjectsCount = t.Links.SubjectsCount
	certificate.OemLinks = t.Links.Oem
	certificate.rekeyTarget = t.Actions.RekeyCertificate.Target
	certificate.renewTarget = t.Actions.RenewCertificate.Target
	certificate.OemActions = t.Actions.Oem

	return nil
}

// GetCertificate will get a Certificate instance from the Redfish service.
func GetCertificate(c common.Client, uri string) (*Certificate, error) {
	var certificate Certificate
	return &certificate, certificate.Get(c, uri, &certificate)
}

// ListReferencedCertificates gets the Certificates collection.
func ListReferencedCertificates(c common.Client, link string) ([]*Certificate, error) { //nolint:dupl
	var result []*Certificate
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Certificate
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		certificate, err := GetCertificate(c, link)
		ch <- GetResult{Item: certificate, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

func (certificate *Certificate) RekeyCertificate(challengePassword, keyCurveID, keyPairAlgorithm string, keyBitLength int) error {
	t := struct {
		ChallengePassword string
		KeyCurveID        string `json:"KeyCurveId"`
		KeyPairAlgorithm  string
		KeyBitLength      int
	}{
		ChallengePassword: challengePassword,
		KeyCurveID:        keyCurveID,
		KeyPairAlgorithm:  keyPairAlgorithm,
		KeyBitLength:      keyBitLength,
	}
	return certificate.Post(certificate.rekeyTarget, t)
}

func (certificate *Certificate) RenewCertificate(challengePassword string) error {
	t := struct {
		ChallengePassword string
	}{ChallengePassword: challengePassword}
	return certificate.Post(certificate.renewTarget, t)
}
