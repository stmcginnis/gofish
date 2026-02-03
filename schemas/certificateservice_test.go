//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var serviceBody = `{
	"@odata.type": "#CertificateService.v1_0_1.CertificateService",
	"@odata.id": "/redfish/v1/CertificateService",
	"Id": "CertificateService",
	"Name": "Certificate Service",
	"CertificateLocations": {
	  "@odata.id": "/redfish/v1/CertificateService/CertificateLocations"
	},
	"Actions": {
	  "Oem": {},
	  "#CertificateService.GenerateCSR": {
		"target": "/redfish/v1/CertificateService/Actions/CertificateService.GenerateCSR",
		"@Redfish.ActionInfo": "/redfish/v1/CertificateService/GenerateCSRActionInfo"
	  },
	  "#CertificateService.ReplaceCertificate": {
		"target": "/redfish/v1/CertificateService/Actions/CertificateService.ReplaceCertificate",
		"@Redfish.ActionInfo": "/redfish/v1/CertificateService/ReplaceCertificateActionInfo"
	  }
	}
  }`

// TestCertificateService tests the parsing of CertificateService objects.
func TestCertificateService(t *testing.T) {
	var result CertificateService
	err := json.NewDecoder(strings.NewReader(serviceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CertificateService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Certificate Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.certificateLocations != "/redfish/v1/CertificateService/CertificateLocations" {
		t.Errorf("Received invalid certificate locations: %s",
			result.certificateLocations)
	}
}
