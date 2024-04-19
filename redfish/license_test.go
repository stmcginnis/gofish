//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var licenseBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#License.License",
		"@odata.id": "/redfish/v1/LicenseService/Licenses/FD00000032890062",
		"@odata.type": "#License.v1_1_1.License",
		"AuthorizationScope": "Service",
		"Description": "iDRAC9 16G Datacenter License",
		"DownloadURI": "/redfish/v1/LicenseService/Licenses/FD00000032890062/DownloadURI",
		"EntitlementId": "FD00000032890062",
		"ExpirationDate": null,
		"Id": "FD00000032890062",
		"InstallDate": null,
		"LicenseInfoURI": "",
		"LicenseOrigin": "Installed",
		"LicenseType": "Production",
		"Links": {},
		"Name": "FD00000032890062",
		"Removable": true,
		"Status": {
		  "Health": "OK",
		  "State": "Enabled"
		}
	  }`)

// TestLicense tests the parsing of License objects.
func TestLicense(t *testing.T) {
	var result License
	err := json.NewDecoder(licenseBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "FD00000032890062", result.ID)
	assertEquals(t, "FD00000032890062", result.Name)
	assertEquals(t, "Installed", string(result.LicenseOrigin))
	assertEquals(t, "true", fmt.Sprintf("%t", result.Removable))
}
