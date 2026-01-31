//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var tcBody = `{
	"@odata.type": "#TrustedComponent.v1_1_0.TrustedComponent",
	"@odata.context": "/redfish/v1/$metadata#TrustedComponent.TrustedComponent",
	"Id": "TPM",
	"Name": "TrustedComponent for TPM",
	"Description": "TrustedComponent for Trusted Platform Module",
	"TrustedComponentType": "Discrete",
	"Certificates": {
	  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/TrustedComponents/TPM/Certificates"
	},
	"Links": {
	  "ComponentsProtected": [
		{
		  "@odata.id": "/redfish/v1/Systems/System.Embedded.1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Chassis/System.Embedded.1/TrustedComponents/TPM"
  }`

// TestTrustedComponent tests the parsing of TrustedComponent objects.
func TestTrustedComponent(t *testing.T) {
	var result TrustedComponent
	err := json.NewDecoder(strings.NewReader(tcBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "TPM" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "TrustedComponent for TPM" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.certificates != "/redfish/v1/Chassis/System.Embedded.1/TrustedComponents/TPM/Certificates" {
		t.Errorf("Invalid fan name: %s", result.certificates)
	}
}
