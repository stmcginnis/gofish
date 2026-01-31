//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var mediaControllerBody = `{
	"@odata.type": "#MediaController.v1_3_1.MediaController",
	"Id": "MediaController1",
	"Name": "Media Controller 1",
	"MediaControllerType": "Memory",
	"Manufacturer": "Contoso",
	"Model": "Contoso MediaController",
	"SerialNumber": "2M220100SL",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"UUID": "41784113-ed6b-2284-1414-916520dc1dd1",
	"Ports": {
	  "@odata.id": "/redfish/v1/Chassis/GenZ/MediaControllers/1/Ports"
	},
	"Actions": {
	  "#MediaController.Reset": {
		"target": "/redfish/v1/Chassis/GenZ/MediaControllers/1/Actions/MediaController.Reset",
		"ResetType@Redfish.AllowableValues": [
		  "ForceRestart"
		]
	  }
	},
	"Links": {
	  "Endpoints": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/GenZ/Endpoints/1"
		}
	  ],
	  "MemoryDomains": [
		{
		  "@odata.id": "/redfish/v1/Chassis/GenZ/MemoryDomains/1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Chassis/GenZ/MediaControllers/1"
  }`

// TestMediaController tests the parsing of MediaController objects.
func TestMediaController(t *testing.T) {
	var result MediaController
	err := json.NewDecoder(strings.NewReader(mediaControllerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "MediaController1", result.ID)
	assertEquals(t, "Media Controller 1", result.Name)
	assertEquals(t, "41784113-ed6b-2284-1414-916520dc1dd1", result.UUID)
	assertEquals(t, "/redfish/v1/Chassis/GenZ/MediaControllers/1/Ports", result.ports)
	assertEquals(t, "/redfish/v1/Fabrics/GenZ/Endpoints/1", result.endpoints[0])
	assertEquals(t, "/redfish/v1/Chassis/GenZ/MemoryDomains/1", result.memoryDomains[0])
	assertEquals(t, "/redfish/v1/Chassis/GenZ/MediaControllers/1/Actions/MediaController.Reset", result.resetTarget)
}

// TestMediaControllerReset tests the MediaController Reset call.
func TestMediaControllerReset(t *testing.T) {
	var result MediaController
	err := json.NewDecoder(strings.NewReader(applicationBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.Reset(OnResetType)
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
