//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var licenseServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#LicenseService.LicenseService",
		"@odata.id": "/redfish/v1/LicenseService",
		"@odata.type": "#LicenseService.v1_1_0.LicenseService",
		"Actions": {
		  "#LicenseService.Install": {
			"TransferProtocol@Redfish.AllowableValues": [
			  "HTTP",
			  "HTTPS",
			  "NFS",
			  "CIFS"
			],
			"target": "/redfish/v1/LicenseService/Actions/LicenseService.Install"
		  }
		},
		"Description": "This resource represent a license service and the properties that affect the service itself.",
		"Id": "LicenseService",
		"Licenses": {
		  "@odata.id": "/redfish/v1/LicenseService/Licenses"
		},
		"Name": "LicenseService",
		"ServiceEnabled": true
	  }`)

// TestLicenseService tests the parsing of LicenseService objects.
func TestLicenseService(t *testing.T) {
	var result LicenseService
	err := json.NewDecoder(licenseServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "LicenseService", result.ID)
	assertEquals(t, "/redfish/v1/LicenseService/Licenses", result.licenses)
	assertEquals(t, "/redfish/v1/LicenseService/Actions/LicenseService.Install", result.installTarget)
}
