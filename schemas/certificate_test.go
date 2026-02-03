//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var certificateBody = strings.NewReader(
	`{
		"@odata.type": "#Certificate.v1_7_0.Certificate",
		"Id": "1",
		"Name": "HTTPS Certificate",
		"CertificateString": "-----BEGIN CERTIFICATE-----\nMIIFsTCC [*truncated*] GXG5zljlu\n-----END CERTIFICATE-----",
		"CertificateType": "PEM",
		"Issuer": {
			"Country": "US",
			"State": "Oregon",
			"City": "Portland",
			"Organization": "Contoso",
			"OrganizationalUnit": "ABC",
			"CommonName": "manager.contoso.org"
		},
		"Subject": {
			"Country": "US",
			"State": "Oregon",
			"City": "Portland",
			"Organization": "Contoso",
			"OrganizationalUnit": "ABC",
			"CommonName": "manager.contoso.org"
		},
		"ValidNotBefore": "2018-09-07T13:22:05Z",
		"ValidNotAfter": "2019-09-07T13:22:05Z",
		"KeyUsage": [
			"KeyEncipherment",
			"ServerAuthentication"
		],
		"SerialNumber": "5d:7a:d8:df:f6:fc:c1:b3:ca:fe:fb:cc:38:f3:01:64:51:ea:05:cb",
		"Fingerprint": "A6:E9:D2:5C:DC:52:DA:4B:3B:14:97:F3:A4:53:D9:99:A1:0B:56:41",
		"FingerprintHashAlgorithm": "TPM_ALG_SHA1",
		"SignatureAlgorithm": "sha256WithRSAEncryption",
		"@odata.id": "/redfish/v1/Managers/BMC/NetworkProtocol/HTTPS/Certificates/1"
	}`)

// TestCertificate tests the parsing of Certificate objects.
func TestCertificate(t *testing.T) {
	var result Certificate

	if err := json.NewDecoder(certificateBody).Decode(&result); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "HTTPS Certificate", result.Name)
	assertEquals(t, "PEM", string(result.CertificateType))
	assertEquals(t, "Contoso", result.Issuer.Organization)
	assertEquals(t, "2019-09-07T13:22:05Z", result.ValidNotAfter)
	assertEquals(t, "TPM_ALG_SHA1", result.FingerprintHashAlgorithm)
	assertEquals(t, "sha256WithRSAEncryption", result.SignatureAlgorithm)
}
