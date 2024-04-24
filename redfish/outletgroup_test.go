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

var outletGroupBody = `{
	"@odata.type": "#OutletGroup.v1_1_2.OutletGroup",
	"Id": "Rack5Storage",
	"Name": "Outlet Group Rack5Storage",
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"CreatedBy": "Bob",
	"PowerOnDelaySeconds": 4,
	"PowerOffDelaySeconds": 0,
	"PowerState": "On",
	"PowerEnabled": true,
	"PowerWatts": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/GroupPowerA",
	  "Reading": 412.36
	},
	"EnergykWh": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/GroupEnergyA",
	  "Reading": 26880
	},
	"Links": {
	  "Outlets": [
		{
		  "@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1"
		},
		{
		  "@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A2"
		},
		{
		  "@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A3"
		}
	  ]
	},
	"Actions": {
	  "#OutletGroup.PowerControl": {
		"target": "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups/Rack5Storage/OutletGroup.PowerControl"
	  },
	  "#OutletGroup.ResetMetrics": {
		"target": "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups/Rack5Storage/OutletGroup.ResetMetrics"
	  }
	},
	"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups/Rack5Storage"
  }`

// TestOutletGroup tests the parsing of OutletGroup objects.
func TestOutletGroup(t *testing.T) {
	var result OutletGroup
	err := json.NewDecoder(strings.NewReader(outletGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Rack5Storage", result.ID)
	assertEquals(t, "Outlet Group Rack5Storage", result.Name)
	assertEquals(t, "Bob", result.CreatedBy)
	assertEquals(t, "On", string(result.PowerState))
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/GroupPowerA", result.PowerWatts.DataSourceURI)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/GroupEnergyA", result.EnergykWh.DataSourceURI)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1", result.outlets[0])

	if !result.PowerEnabled {
		t.Error("Expected power to be enabled")
	}

	if result.PowerOnDelaySeconds != 4 {
		t.Errorf("Unexpected PowerOnDelaySeconds value: %.2f", result.PowerOnDelaySeconds)
	}

	if result.PowerOffDelaySeconds != 0 {
		t.Errorf("Unexpected PowerOffDelaySeconds value: %.2f", result.PowerOffDelaySeconds)
	}

	if result.PowerWatts.Reading != 412.36 {
		t.Errorf("Unexpected PowerWatts.Reading value: %.2f", result.PowerWatts.Reading)
	}
}

// TestOutletGroupResetMetrics tests the OutletGroup ResetMetrics call.
func TestOutletGroupResetMetrics(t *testing.T) {
	var result OutletGroup
	err := json.NewDecoder(strings.NewReader(outletGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.ResetMetrics()
	if err != nil {
		t.Errorf("Error making ResetMetrics call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if calls[0].Payload != "" {
		t.Errorf("Expected payload: %s", calls[0].Payload)
	}

	if calls[0].URL != "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups/Rack5Storage/OutletGroup.ResetMetrics" {
		t.Errorf("Expected target URL: %s", calls[0].URL)
	}
}

// TestOutletGroupPowerControl tests the OutletGroup PowerControl call.
func TestOutletGroupPowerControl(t *testing.T) {
	var result OutletGroup
	err := json.NewDecoder(strings.NewReader(outletGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.PowerControl(OffPowerState)
	if err != nil {
		t.Errorf("Error making PowerControl call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "PowerState:Off") {
		t.Errorf("Expected payload: %s", calls[0].Payload)
	}

	if calls[0].URL != "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups/Rack5Storage/OutletGroup.PowerControl" {
		t.Errorf("Expected target URL: %s", calls[0].URL)
	}
}
