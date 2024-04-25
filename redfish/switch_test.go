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

var switchBody = `{
	"@odata.type": "#Switch.v1_9_2.Switch",
	"Id": "Switch1",
	"Name": "SAS Switch",
	"SwitchType": "SAS",
	"Manufacturer": "Contoso",
	"Model": "SAS1000",
	"SKU": "67B",
	"SerialNumber": "2M220100SL",
	"PartNumber": "76-88883",
	"Ports": {
	  "@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch1/Ports"
	},
	"Redundancy": [
	  {
		"@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch1#/Redundancy/0",
		"MemberId": "Redundancy",
		"Mode": "Sharing",
		"MaxNumSupported": 2,
		"MinNumNeeded": 1,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"RedundancySet": [
		  {
			"@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch1"
		  },
		  {
			"@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch2"
		  }
		]
	  }
	],
	"Links": {
	  "Chassis": {
		"@odata.id": "/redfish/v1/Chassis/Switch1"
	  },
	  "ManagedBy": [
		{
		  "@odata.id": "/redfish/v1/Managers/Switch1"
		},
		{
		  "@odata.id": "/redfish/v1/Managers/Switch2"
		}
	  ]
	},
	"Actions": {
	  "#Switch.Reset": {
		"target": "/redfish/v1/Fabrics/SAS/Switches/Switch1/Actions/Switch.Reset",
		"ResetType@Redfish.AllowableValues": [
		  "ForceRestart",
		  "GracefulRestart"
		]
	  }
	},
	"@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch1"
  }`

// TestSwitch tests the parsing of Switch objects.
func TestSwitch(t *testing.T) {
	var result Switch
	err := json.NewDecoder(strings.NewReader(switchBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Switch1", result.ID)
	assertEquals(t, "SAS Switch", result.Name)
	assertEquals(t, "SAS", string(result.SwitchType))
	assertEquals(t, "/redfish/v1/Fabrics/SAS/Switches/Switch1/Ports", result.ports)
	assertEquals(t, "/redfish/v1/Fabrics/SAS/Switches/Switch2", result.Redundancy[0].redundancySet[1])
	assertEquals(t, "/redfish/v1/Chassis/Switch1", result.chassis)
	assertEquals(t, "/redfish/v1/Managers/Switch2", result.managedBy[1])
}

// TestSwitchReset tests the Switch Reset call.
func TestSwitchReset(t *testing.T) {
	var result Switch
	err := json.NewDecoder(strings.NewReader(switchBody)).Decode(&result)

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

	if !strings.Contains(calls[0].Payload, "ResetType:On") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}
}
