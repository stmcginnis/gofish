// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/school/common"
)

var chassisBody = strings.NewReader(
	`{
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
		}
	}`)

// TestChassis tests the parsing of Chassis objects.
func TestChassis(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(chassisBody).Decode(&result)

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
		t.Errorf("Recieved invalid asset tag: %s", result.AssetTag)
	}

	if result.ChassisType != RackMountChassisType {
		t.Errorf("Received invalid chassis type: %s", result.ChassisType)
	}

	if result.Status.Health != common.OKHealth {
		t.Errorf("Received invalid health status: %s", result.Status.Health)
	}

	if result.thermal != "/redfish/v1/Chassis/Chassis-1/Thermal" {
		t.Errorf("Recieved invalid thermal reference: %s", result.thermal)
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
}
