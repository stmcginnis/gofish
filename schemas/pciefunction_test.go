//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"net/http"
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

	if *result.FunctionID != 10 {
		t.Errorf("Invalid function ID: %d", result.FunctionID)
	}

	if result.FunctionType != VirtualFunctionType {
		t.Errorf("Invalid function type: %s", result.FunctionType)
	}
}

// TestPCIeFunctionStorageControllers tests that the StorageControllers link
// resolves to standalone StorageController resources.
func TestPCIeFunctionStorageControllers(t *testing.T) {
	var result PCIeFunction
	err := json.Unmarshal([]byte(`{
		"Id": "PCIeFunction-1",
		"Links": {
			"StorageControllers": [{"@odata.id": "/redfish/v1/Systems/1/Storage/1/Controllers/1"}]
		}
	}`), &result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				getCall(`{
					"@odata.id": "/redfish/v1/Systems/1/Storage/1/Controllers/1",
					"Id": "1",
					"Model": "PERC H755",
					"Links": {
						"AttachedVolumes": [{"@odata.id": "/redfish/v1/Systems/1/Storage/1/Volumes/1"}]
					}
				}`),
			},
		},
	}
	result.SetClient(testClient)

	controllers, err := result.StorageControllers()
	if err != nil {
		t.Fatalf("Error getting storage controllers: %s", err)
	}

	if len(controllers) != 1 {
		t.Fatalf("Expected 1 storage controller, got %d", len(controllers))
	}

	if controllers[0].ID != "1" || controllers[0].Model != "PERC H755" {
		t.Errorf("Received invalid storage controller: %s %s", controllers[0].ID, controllers[0].Model)
	}

	if len(controllers[0].attachedVolumes) != 1 {
		t.Errorf("Expected 1 attached volume link, got %d", len(controllers[0].attachedVolumes))
	}

	if calls := testClient.CapturedCalls(); len(calls) != 1 ||
		calls[0].URL != "/redfish/v1/Systems/1/Storage/1/Controllers/1" {
		t.Errorf("Unexpected calls: %+v", calls)
	}
}
