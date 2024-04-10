//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var body = `{
	"@odata.type": "#CertificateLocations.v1_0_1.CertificateLocations",
	"@odata.id": "/redfish/v1/CertificateService/CertificateLocations",
	"Id": "CertificateLocations",
	"Name": "Certificate Locations",
	"Links": {
	  "Certificates": [
		{
		  "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates/1"
		}
	  ]
	}
  }`

// TestCertificateLocations tests the parsing of CertificateLocations objects.
func TestCertificateLocations(t *testing.T) {
	var result CertificateLocations
	err := json.NewDecoder(strings.NewReader(body)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CertificateLocations" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Certificate Locations" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.certificates) != 1 && result.certificates[0] != "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates/1" {
		t.Errorf("Received invalid certificate link: %#v", result.certificates)
	}
}
