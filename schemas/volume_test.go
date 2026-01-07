//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var volumeBody = `{
	"@odata.type": "#Volume.v1_10_0.Volume",
	"Id": "2",
	"Name": "Virtual Disk 2",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Encrypted": false,
	"RAIDType": "RAID0",
	"CapacityBytes": 107374182400,
	"Identifiers": [
	  {
		"DurableNameFormat": "UUID",
		"DurableName": "0324c96c-8031-4f5e-886c-50cd90aca854"
	  }
	],
	"Links": {
	  "Drives": [
		{
		  "@odata.id": "/redfish/v1/Systems/437XR1138R2/Storage/1/Drives/3D58ECBC375FD9F2"
		}
	  ]
	},
	"Actions": {
	  "#Volume.Initialize": {
		"target": "/redfish/v1/Systems/3/Storage/RAIDIntegrated/Volumes/1/Actions/Volume.Initialize",
		"InitializeType@Redfish.AllowableValues": [
		  "Fast",
		  "Slow"
		]
	  }
	},
	"@odata.id": "/redfish/v1/Systems/437XR1138R2/Storage/1/Volumes/2"
  }`

// TesVolume tests the parsing of Volume objects.
func TestVolume(t *testing.T) {
	var result Volume
	err := json.NewDecoder(strings.NewReader(volumeBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "2", result.ID)
	assertEquals(t, "RAID0", string(result.RAIDType))
	assertEquals(t, "UUID", string(result.Identifiers[0].DurableNameFormat))
	assertEquals(t, "0324c96c-8031-4f5e-886c-50cd90aca854", result.Identifiers[0].DurableName)
	assertEquals(t, "/redfish/v1/Systems/437XR1138R2/Storage/1/Drives/3D58ECBC375FD9F2", result.drives[0])
}

// TestVolumeInitialize tests the Volume Initialize call.
func TestVolumeInitialize(t *testing.T) {
	var result Volume
	err := json.NewDecoder(strings.NewReader(volumeBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.Initialize(BackgroundInitializeMethod, FastInitializeType)
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "InitializeMethod:Background") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "InitializeType:Fast") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}
}
