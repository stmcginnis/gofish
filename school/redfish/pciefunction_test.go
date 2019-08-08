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
)

var pcieFunctionBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#PCIeFunction.PCIeFunction",
		"@odata.type": "#PCIeFunction.v1_0_0.PCIeFunction",
		"@odata.id": "/redfish/v1/PCIeFunction",
		"Id": "PCIeFunction-1",
		"Name": "PCIeFunctionOne",
		"Description": "PCIeFunction One",
		"ClassCode": "01",
		"DeviceClass": "MassStorageController",
		"DeviceId": "01",
		"FunctionId": 10,
		"FunctionType": "Virtual",
		"Links": {
			"Drives": [{
				"@odata.id": "/redfish/v1/Drives/1"
			}],
			"Drives@odata.count": 1,
			"EthernetInterfaces": [{
				"@odata.id": "/redfish/v1/EthernetInterfaces/1"
			}],
			"EthernetInterfaces@odata.count": 1,
			"NetworkDeviceFunctions": [{
				"@odata.id": "/redfish/v1/NetworkDeviceFunction/1"
			}],
			"NetworkDeviceFunctions@odata.count": 1,
			"PCIeDevice": {
				"@odata.id": "/redfish/v1/PCIeDevices/1"
			},
			"StorageControllers": [{
				"@odata.id": "/redfish/v1/StorageControllers/1"
			}],
			"StorageControllers@odata.count": 1
		},
		"RevisionId": "1.0",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"SubsystemId": "1F",
		"SubsystemVendorId": "0a",
		"VendorId": "4f"
	}`)

// TestPCIeFunction tests the parsing of PCIeFunction objects.
func TestPCIeFunction(t *testing.T) {
	var result PCIeFunction
	err := json.NewDecoder(pcieFunctionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "PCIeFunction-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "PCIeFunctionOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.DeviceClass != MassStorageControllerDeviceClass {
		t.Errorf("Invalid device class: %s", result.DeviceClass)
	}

	if result.FunctionID != 10 {
		t.Errorf("Invalid function ID: %d", result.FunctionID)
	}

	if result.FunctionType != VirtualFunctionType {
		t.Errorf("Invalid function type: %s", result.FunctionType)
	}
}
