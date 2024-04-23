//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var applicationBody = `{
		"@odata.type": "#Application.v1_0_0.Application",
		"Id": "Logger",
		"Name": "Logging Agent",
		"Version": "1.5.1",
		"Vendor": "Contoso",
		"StartTime": "2021-10-29T10:42:38+06:00",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"DestinationURIs": [
			"https://listeners.contoso.org:8000/handler"
		],
		"MetricsURIs": [
			"https://192.168.0.12:7000"
		],
		"Actions": {
			"#Application.Reset": {
				"target": "/redfish/v1/Systems/VM1/OperatingSystem/Applications/Logger/Actions/Application.Reset",
				"ResetType@Redfish.AllowableValues": [
					"On",
					"ForceOff",
					"GracefulShutdown",
					"GracefulRestart",
					"ForceRestart",
					"ForceOn"
				]
			}
		},
		"@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Applications/Logger"
	}`

// TestApplication tests the parsing of Application objects.
func TestApplication(t *testing.T) {
	var result Application
	err := json.NewDecoder(strings.NewReader(applicationBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Logger", result.ID)
	assertEquals(t, "Logging Agent", result.Name)
	assertEquals(t, "1.5.1", result.Version)
	assertEquals(t, "Contoso", result.Vendor)

	if len(result.DestinationURIs) != 1 {
		t.Errorf("Expected 1 destination URI, got %#v", result.DestinationURIs)
	}

	if len(result.MetricsURIs) != 1 {
		t.Errorf("Expected 1 metrics URI, got %#v", result.MetricsURIs)
	}
}

// TestApplicationReset tests the Application Reset call.
func TestApplicationReset(t *testing.T) {
	var result Application
	err := json.NewDecoder(strings.NewReader(applicationBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.Reset(OnResetType)
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "On") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}
}
