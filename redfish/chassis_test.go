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

var chassisBody = `{
		"@odata.context": "/redfish/v1/$metadata#Chassis.Chassis",
		"@odata.id": "/redfish/v1/Chassis/Chassis-1",
		"@odata.type": "#Chassis.v1_0_0.Chassis",
		"Id": "Chassis-1",
		"Name": "Computer System Chassis",
		"ChassisType": "RackMount",
		"Manufacturer": "Redfish Computers",
		"Model": "3500RX",
		"SKU": "8675309",
		"SerialNumber": "437XR1138R2",
		"Version": "1.02",
		"PartNumber": "224071-J23",
		"AssetTag": "Chicago-45Z-2381",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Thermal": {
			"@odata.id": "/redfish/v1/Chassis/Chassis-1/Thermal"
		},
		"Power": {
			"@odata.id": "/redfish/v1/Chassis/Chassis-1/Power"
		},
		"Links": {
			"ComputerSystems": [
				{
					"@odata.id": "/redfish/v1/Systems/System-1"
				}
			],
			"ResourceBlocks": [],
			"ManagedBy": [
				{
					"@odata.id": "/redfish/v1/Managers/BMC-1"
				}
			]
		},
		"Actions": {
			"#Chassis.Reset": {
				"target": "/redfish/v1/Chassis/System.Embedded.1/Actions/Chassis.Reset",
				"ResetType@Redfish.AllowableValues": [
					"On",
					"ForceOff"
				]
			}
		}
	}`

// TestChassis tests the parsing of Chassis objects.
func TestChassis(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(chassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Chassis-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Computer System Chassis" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AssetTag != "Chicago-45Z-2381" {
		t.Errorf("Received invalid asset tag: %s", result.AssetTag)
	}

	if result.ChassisType != RackMountChassisType {
		t.Errorf("Received invalid chassis type: %s", result.ChassisType)
	}

	if result.Status.Health != common.OKHealth {
		t.Errorf("Received invalid health status: %s", result.Status.Health)
	}

	if result.thermal != "/redfish/v1/Chassis/Chassis-1/Thermal" {
		t.Errorf("Received invalid thermal reference: %s", result.thermal)
	}

	if result.power != "/redfish/v1/Chassis/Chassis-1/Power" {
		t.Errorf("Received invalid power reference: %s", result.power)
	}

	if len(result.computerSystems) != 1 {
		t.Errorf("Expected 1 computer system, got %d", len(result.computerSystems))
	}

	if result.computerSystems[0] != "/redfish/v1/Systems/System-1" {
		t.Errorf("Invalid computer system reference: %s", result.computerSystems[0])
	}

	if len(result.resourceBlocks) != 0 {
		t.Errorf("Resource blocks should have been 0, got %d", len(result.resourceBlocks))
	}

	if len(result.managedBy) != 1 {
		t.Errorf("Expected 1 managed by reference, got %d", len(result.managedBy))
	}

	if result.managedBy[0] != "/redfish/v1/Managers/BMC-1" {
		t.Errorf("Invalid managed by reference: %s", result.managedBy[0])
	}

	if result.resetTarget != "/redfish/v1/Chassis/System.Embedded.1/Actions/Chassis.Reset" {
		t.Errorf("Invalid reset action target: %s", result.resetTarget)
	}

	if len(result.SupportedResetTypes) != 2 {
		t.Errorf("Invalid allowable reset actions, expected 2, got %d",
			len(result.SupportedResetTypes))
	}
}

// TestChassisUpdate tests the Update call.
func TestChassisUpdate(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(chassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.AssetTag = "TestAssetTag"
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[0].Payload, result.AssetTag) {
		t.Errorf("Unexpected update payload: %s", calls[0].Payload)
	}
}
