//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.1 - #Signature.v1_0_3.Signature

package schemas

import (
	"encoding/json"
)

type SignatureTypeRegistry string

const (
	// UEFISignatureTypeRegistry shall indicate that the 'SignatureType' string
	// contains the #define name of the SignatureType member of the
	// EFI_SIGNATURE_LIST, as defined by the UEFI Specification. This value shall
	// also indicate that the format of the 'SignatureString' is a big-endian
	// hex-encoded string of the binary value specified in the UEFI SignatureData
	// array in EFI_SIGNATURE_DATA, as defined by the UEFI Specification.
	UEFISignatureTypeRegistry SignatureTypeRegistry = "UEFI"
)

// Signature This resource contains a signature for a Redfish implementation.
type Signature struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SignatureString shall contain the string of the signature, and the format
	// shall follow the requirements specified by the value of the 'SignatureType'
	// property. If the signature contains any private keys, they shall be removed
	// from the string in responses. If the private key for the signature is not
	// known by the service and is needed to use the signature, the client shall
	// provide the private key as part of the string in the 'POST' request.
	SignatureString string
	// SignatureType shall contain the format type for the signature. The format is
	// qualified by the value of the 'SignatureTypeRegistry' property.
	SignatureType string
	// SignatureTypeRegistry shall contain the type for the signature.
	SignatureTypeRegistry SignatureTypeRegistry
	// UefiSignatureOwner shall contain the GUID of the UEFI signature owner for
	// this signature as defined by the UEFI Specification. This property shall
	// only be present if the 'SignatureTypeRegistry' property is 'UEFI'.
	UefiSignatureOwner string
}

// GetSignature will get a Signature instance from the service.
func GetSignature(c Client, uri string) (*Signature, error) {
	return GetObject[Signature](c, uri)
}

// ListReferencedSignatures gets the collection of Signature from
// a provided reference.
func ListReferencedSignatures(c Client, link string) ([]*Signature, error) {
	return GetCollectionObjects[Signature](c, link)
}
