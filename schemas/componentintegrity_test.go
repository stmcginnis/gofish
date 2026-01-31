//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var componentIntegrityBody = `{
	"@odata.type": "#ComponentIntegrity.v1_2_2.ComponentIntegrity",
	"Id": "TPM-0",
	"Description": "TPM physically attached to a GPU.",
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"ComponentIntegrityType": "TPM",
	"ComponentIntegrityTypeVersion": "1.2.0",
	"ComponentIntegrityEnabled": true,
	"LastUpdated": "2021-11-02T14:09:54-07:00",
	"TargetComponentURI": "/redfish/v1/Systems/437XR1138R2#/TrustedModules/0",
	"Links": {
	  "ComponentsProtected": [
		{
		  "@odata.id": "/redfish/v1/Systems/437XR1138R2/GraphicsControllers/GPU1"
		}
	  ]
	},
	"TPM": {
	  "MeasurementSet": {
		"Measurements": [
		  {
			"PCR": 1,
			"Measurement": "h6spEuxbyOtGhP35UoGhTcVX3iRaZQGDw4Yk5oQcabw=",
			"LastUpdated": "2021-10-31T20:14:27-07:00",
			"MeasurementHashAlgorithm": "TPM_ALG_SHA256"
		  },
		  {
			"PCR": 3,
			"Measurement": "GnbzS4ToNQb+Y7SxXw4AvRDTf4SzO5eeAlAlDca28AA=",
			"LastUpdated": "2021-10-31T20:14:27-07:00",
			"MeasurementHashAlgorithm": "TPM_ALG_SHA256"
		  },
		  {
			"PCR": 1,
			"Measurement": "pLJa5Dyh8CDYFZ1WNOrsiSG1eyCPBlre42CD7CTywg7VkcC4afw4ZG3gQxi2XEFCt5jxz6tN1/cbx/DNx2/tOg==",
			"LastUpdated": "2021-10-31T20:14:27-07:00",
			"MeasurementHashAlgorithm": "TPM_ALG_SHA512"
		  },
		  {
			"PCR": 3,
			"Measurement": "GBgEucATV8omirTmYqY+vvbbisHR1jBKfVAEK1XSifBHnnIYXopsc0NExURDSSyPjO21NrPqnwiq5LhI1p6rzQ==",
			"LastUpdated": "2021-10-31T20:14:27-07:00",
			"MeasurementHashAlgorithm": "TPM_ALG_SHA512"
		  }
		]
	  },
	  "IdentityAuthentication": {
		"VerificationStatus": "Success",
		"ComponentCertificate": {
		  "@odata.id": "/redfish/v1/Systems/437XR1138R2/Certificates/TPMcert"
		}
	  },
	  "ComponentCommunication": {
		"Sessions": [
		  {
			"SessionId": 4556,
			"SessionType": "Plain"
		  }
		]
	  }
	},
	"@odata.id": "/redfish/v1/ComponentIntegrity/TPM-0"
  }`

// TestComponentIntegrity tests the parsing of ComponentIntegrity objects.
func TestComponentIntegrity(t *testing.T) {
	var result ComponentIntegrity
	err := json.NewDecoder(strings.NewReader(componentIntegrityBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "TPM-0", result.ID)
	assertEquals(t, "TPM physically attached to a GPU.", result.Description)
	assertEquals(t, "TPM", string(result.ComponentIntegrityType))
	assertEquals(t, "1.2.0", result.ComponentIntegrityTypeVersion)
	assertEquals(t, "/redfish/v1/Systems/437XR1138R2#/TrustedModules/0", result.TargetComponentURI)
	assertEquals(t, "Success", string(result.TPM.IdentityAuthentication.VerificationStatus))

	if len(result.TPM.MeasurementSet.Measurements) != 4 {
		t.Errorf("Expected 4 measurements, got %#v", result.TPM.MeasurementSet.Measurements)
	}

	if !result.ComponentIntegrityEnabled {
		t.Error("Expected component integrity to be enabled")
	}
}
