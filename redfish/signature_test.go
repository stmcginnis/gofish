//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var signatureBody = `{
	"@odata.type": "#Signature.v1_0_2.Signature",
	"Id": "1",
	"Name": "SHA256 Signature",
	"SignatureString": "80B4D96931BF0D02FD91A61E19D14F1DA452E66DB2408CA8604D411F92659F0A",
	"SignatureTypeRegistry": "UEFI",
	"SignatureType": "EFI_CERT_SHA256_GUID",
	"UefiSignatureOwner": "28d5e212-165b-4ca0-909b-c86b9cee0112",
	"@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/db/Signatures/1"
  }`

// TestSignature tests the parsing of Signature objects.
func TestSignature(t *testing.T) {
	var result Signature
	err := json.NewDecoder(strings.NewReader(signatureBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "SHA256 Signature", result.Name)
	assertEquals(t, "80B4D96931BF0D02FD91A61E19D14F1DA452E66DB2408CA8604D411F92659F0A", result.SignatureString)
	assertEquals(t, "UEFI", string(result.SignatureTypeRegistry))
	assertEquals(t, "EFI_CERT_SHA256_GUID", result.SignatureType)
	assertEquals(t, "28d5e212-165b-4ca0-909b-c86b9cee0112", result.UefiSignatureOwner)
}
