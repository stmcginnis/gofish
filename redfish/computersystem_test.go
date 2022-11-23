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

var computerSystemBody = `{
		"@odata.context": "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		"@odata.id": "/redfish/v1/Systems/System-1",
		"@odata.type": "#ComputerSystem.v1_3_0.ComputerSystem",
		"Id": "System-1",
		"Name": "My Computer System",
		"SystemType": "Physical",
		"AssetTag": "free form asset tag",
		"Manufacturer": "Manufacturer Name",
		"Model": "Model Name",
		"SKU": "",
		"SerialNumber": "2M220100SL",
		"PartNumber": "",
		"Description": "Description of server",
		"UUID": "00000000-0000-0000-0000-000000000000",
		"HostName": "web-srv344",
		"Status": {
			"State": "Enabled",
			"Health": "OK",
			"HealthRollup": "OK"
		},
		"IndicatorLED": "Off",
		"PowerState": "On",
		"Boot": {
			"AutomaticRetryAttempts": 3,
			"AutomaticRetryConfig": "Disabled",
			"AutomaticRetryConfig@Redfish.AllowableValues": [
				"Disabled",
				"RetryAttempts"
			],
			"BootSourceOverrideEnabled": "Once",
			"BootSourceOverrideMode": "UEFI",
			"BootSourceOverrideTarget": "Pxe",
			"BootSourceOverrideTarget@Redfish.AllowableValues": [
				"None",
				"Pxe",
				"Floppy",
				"Cd",
				"Usb",
				"Hdd",
				"BiosSetup",
				"Utilities",
				"Diags",
				"UefiTarget",
				"SDCard",
				"UefiHttp"
			],
			"UefiTargetBootSourceOverride": "uefi device path"
		},
		"BiosVersion": "P79 v1.00 (09/20/2013)",
		"ProcessorSummary": {
			"Status": {
				"Health": "OK",
				"State": "Enabled"
			},
			"Count": 2,
			"Model": "Multi-Core Intel(R) Xeon(R) processor 7500 Series"
		},
		"MemorySummary": {
			"Status": {
				"Health": "OK",
				"State": "Enabled"
			},
			"TotalSystemMemoryGiB": 65536,
			"TotalSystemPersistentMemoryGiB": 262144
		},
		"TrustedModules": [
			{
				"Status": {
					"State": "Enabled",
					"Health": "OK"
				},
				"ModuleType": "TPM2_0",
				"FirmwareVersion": "3.1",
				"FirmwareVersion2": "1",
				"InterfaceTypeSelection": "None"
			}
		],
		"Processors": {
			"@odata.id": "/redfish/v1/Systems/System-1/Processors"
		},
		"Memory": {
			"@odata.id": "/redfish/v1/Systems/System-1/Memory"
		},
		"EthernetInterfaces": {
			"@odata.id": "/redfish/v1/Systems/System-1/EthernetInterfaces"
		},
		"SimpleStorage": {
			"@odata.id": "/redfish/v1/Systems/System-1/SimpleStorage"
		},
		"Links": {
			"Chassis": [
				{
					"@odata.id": "/redfish/v1/Chassis/Chassis-1"
				}
			],
			"ManagedBy": [
				{
					"@odata.id": "/redfish/v1/Managers/BMC-1"
				}
			],
			"Oem": {}
		},
		"Actions": {
			"#ComputerSystem.Reset": {
				"target": "/redfish/v1/Systems/System-1/Actions/ComputerSystem.Reset",
				"@Redfish.ActionInfo": "/redfish/v1/Systems/System-1/ResetActionInfo",
				"ResetType@Redfish.AllowableValues": [
					"On",
					"ForceOff",
					"ForceRestart",
					"Nmi",
					"ForceOn",
					"PushPowerButton"
				]
			},
			"#ComputerSystem.SetDefaultBootOrder": {
				"target": "/redfish/v1/Systems/System-1/Actions/ComputerSystem.SetDefaultBootOrder"
			}
		}
	}`

// TestComputerSystem tests the parsing of ComputerSystem objects.
func TestComputerSystem(t *testing.T) { //nolint
	var result ComputerSystem
	err := json.NewDecoder(strings.NewReader(computerSystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "System-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "My Computer System" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.SystemType != PhysicalSystemType {
		t.Errorf("Incorrect system type: %s", result.SystemType)
	}

	if result.AssetTag != "free form asset tag" {
		t.Errorf("Received incorrect asset tag: %s", result.AssetTag)
	}

	if result.Status.Health != common.OKHealth {
		t.Errorf("Received invalid health status: %s", result.Status.Health)
	}

	if result.Status.HealthRollup != common.OKHealth {
		t.Errorf("Received invalid health rollup status: %s", result.Status.HealthRollup)
	}

	if result.IndicatorLED != common.OffIndicatorLED {
		t.Errorf("Received invalid indicator status: %s", result.IndicatorLED)
	}

	if result.PowerState != OnPowerState {
		t.Errorf("Received invalid power status: %s", result.PowerState)
	}
	if result.Boot.AutomaticRetryAttempts != 3 {
		t.Errorf("Received invalid boot automatic retry attempts: %d", result.Boot.AutomaticRetryAttempts)
	}

	if result.Boot.AutomaticRetryConfig != "Disabled" {
		t.Errorf("Received invalid boot automatic retry config: %s", result.Boot.AutomaticRetryConfig)
	}

	if result.Boot.BootSourceOverrideEnabled != "Once" {
		t.Errorf("Received invalid boot source override: %s", result.Boot.BootSourceOverrideEnabled)
	}

	if result.Boot.BootSourceOverrideMode != "UEFI" {
		t.Errorf("Received invalid boot source override mode: %s", result.Boot.BootSourceOverrideMode)
	}

	if result.Boot.BootSourceOverrideTarget != "Pxe" {
		t.Errorf("Received invalid boot source target: %s", result.Boot.BootSourceOverrideTarget)
	}

	if result.Boot.UefiTargetBootSourceOverride != "uefi device path" {
		t.Errorf("Received invalid uefi target boot source: %s", result.Boot.UefiTargetBootSourceOverride)
	}

	if result.ProcessorSummary.Status.State != common.EnabledState {
		t.Errorf("Received invalid processor summary state: %s", result.ProcessorSummary.Status.State)
	}

	if result.ProcessorSummary.Count != 2 {
		t.Errorf("Received invalid processor count: %d", result.ProcessorSummary.Count)
	}

	if result.MemorySummary.Status.State != common.EnabledState {
		t.Errorf("Received invalid memory summary state: %s", result.MemorySummary.Status.State)
	}

	if result.MemorySummary.TotalSystemMemoryGiB != 65536 {
		t.Errorf("Received invalid total system memory: %f", result.MemorySummary.TotalSystemMemoryGiB)
	}

	if result.MemorySummary.TotalSystemPersistentMemoryGiB != 262144 {
		t.Errorf("Received invalid total system persistent memory: %f",
			result.MemorySummary.TotalSystemPersistentMemoryGiB)
	}

	if len(result.TrustedModules) != 1 {
		t.Errorf("Received invalid number of trusted modules: %d", len(result.TrustedModules))
	}

	if result.TrustedModules[0].Status.Health != common.OKHealth {
		t.Errorf("Received invalid trusted module health: %s",
			result.TrustedModules[0].Status.Health)
	}

	if result.TrustedModules[0].InterfaceTypeSelection != NoneInterfaceTypeSelection {
		t.Errorf("Received invalid trusted module interface type selection: %s",
			result.TrustedModules[0].InterfaceTypeSelection)
	}

	if result.processors != "/redfish/v1/Systems/System-1/Processors" {
		t.Errorf("Received invalid processors reference: %s", result.processors)
	}

	if result.memory != "/redfish/v1/Systems/System-1/Memory" {
		t.Errorf("Received invalid memory reference: %s", result.memory)
	}

	if result.ethernetInterfaces != "/redfish/v1/Systems/System-1/EthernetInterfaces" {
		t.Errorf("Received invalid ethernet interface reference: %s", result.ethernetInterfaces)
	}

	if result.simpleStorage != "/redfish/v1/Systems/System-1/SimpleStorage" {
		t.Errorf("Received invalid simple storage reference: %s", result.simpleStorage)
	}

	if len(result.chassis) != 1 {
		t.Errorf("Received invalid number of chassis: %d", len(result.chassis))
	}

	if result.chassis[0] != TestChassisPath {
		t.Errorf("Received invalid chassis reference: %s", result.chassis[0])
	}

	if result.resetTarget != "/redfish/v1/Systems/System-1/Actions/ComputerSystem.Reset" {
		t.Errorf("Invalid reset action target: %s", result.resetTarget)
	}

	if len(result.SupportedResetTypes) != 6 {
		t.Errorf("Invalid allowable reset actions, expected 6, got %d",
			len(result.SupportedResetTypes))
	}
	if len(result.ManagedBy) != 1 {
		t.Errorf("Received invalid number of ManagedBy: %d", len(result.ManagedBy))
	}
	if result.ManagedBy[0] != "/redfish/v1/Managers/BMC-1" {
		t.Errorf("Received invalid Managers reference: %s", result.ManagedBy[0])
	}
}

// TestComputerSystemUpdate tests the Update call.
func TestComputerSystemUpdate(t *testing.T) {
	var result ComputerSystem
	err := json.NewDecoder(strings.NewReader(computerSystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.AssetTag = TestAssetTag
	result.HostName = "TestHostName"
	result.IndicatorLED = common.BlinkingIndicatorLED
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AssetTag:TestAssetTag") {
		t.Errorf("Unexpected AssetTag update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "HostName:TestHostName") {
		t.Errorf("Unexpected HostName update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "IndicatorLED:Blinking") {
		t.Errorf("Unexpected IndicatorLED update payload: %s", calls[0].Payload)
	}
}

var bootOptionBody = `{
	"@odata.context": "/redfish/v1/$metadata#BootOption.BootOption",
	"@odata.etag": "W/\"A3A6BF43\"",
	"@odata.id": "/redfish/v1/Systems/1/BootOptions/1/",
	"@odata.type": "#BootOption.v1_0_1.BootOption",
	"Id": "1",
	"Alias": "None",
	"BootOptionReference": "Boot0015",
	"DisplayName": "URL File : http://assets.example.com/ipxe.efi (IPv4)",
	"Name": "Boot Option",
	"UefiDevicePath": "IPv4(0.0.0.0)/Uri(http://assets.example.com/ipxe.efi)"
}`

func TestBootOption(t *testing.T) {
	var result BootOption
	err := json.NewDecoder(strings.NewReader(bootOptionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Boot Option" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Alias != "None" {
		t.Errorf("Received invalid alias: %s", result.Alias)
	}

	if result.BootOptionReference != "Boot0015" {
		t.Errorf("Received invalid bootoptionreference: %s", result.BootOptionReference)
	}

	if result.DisplayName != "URL File : http://assets.example.com/ipxe.efi (IPv4)" {
		t.Errorf("Received invalid displayname: %s", result.DisplayName)
	}

	if result.UefiDevicePath != "IPv4(0.0.0.0)/Uri(http://assets.example.com/ipxe.efi)" {
		t.Errorf("Received invalid uefidevicepath: %s", result.UefiDevicePath)
	}
}
