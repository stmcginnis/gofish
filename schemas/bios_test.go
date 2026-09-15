//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

var biosBody = `{
		"@Redfish.Settings": {
			"@odata.context": "/redfish/v1/$metadata#Settings.Settings",
			"@odata.type": "#Settings.v1_2_1.Settings",
			"SettingsObject": {
			    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Bios/Settings"
			},
			"SupportedApplyTimes": [
			    "OnReset",
			    "AtMaintenanceWindowStart",
			    "InMaintenanceWindowOnReset"
			]
		},
		"@odata.type": "#Bios.v1_0_6.Bios",
		"@odata.context": "/redfish/v1/$metadata#Bios.Bios",
		"@odata.id": "/redfish/v1/Systems/437XR1138R2/BIOS",
		"Id": "BIOS",
		"Name": "BIOS Configuration Current Attributes",
		"Description": "BIOS Attributes",
		"AttributeRegistry": "BiosAttributeRegistryP89.v1_0_0",
		"Attributes": {
			"AdminPhone": "",
			"BootMode": "Uefi",
			"EmbeddedSata": "Raid",
			"NicBoot1": "NetworkBoot",
			"NicBoot2": "Disabled",
			"PowerProfile": "MaxPerf",
			"ProcCoreDisable": 3,
			"ProcHyperthreading": "Enabled",
			"ProcTurboMode": "Enabled",
			"UsbControl": "UsbEnabled",
			"BoolTest1": "True",
			"BoolTest2": 1,
			"BoolTest3": "NotBool"
		},
		"Actions": {
			"#Bios.ResetBios": {
				"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ResetBios"
			},
			"#Bios.ChangePassword": {
				"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ChangePassword"
			}
		},
		"Links": {
			"ActiveSoftwareImage": {
				"@odata.id": "/redfish/v1/Systems/437XR1138R2/BIOS/FirmwareInventory"
			}
		}
	}`

var biosNoAttributesBody = `{
			"@odata.type": "#Bios.v1_0_6.Bios",
			"@odata.context": "/redfish/v1/$metadata#Bios.Bios",
			"@odata.id": "/redfish/v1/Systems/437XR1138R2/BIOS",
			"Id": "BIOS",
			"Name": "BIOS Configuration Current Attributes",
			"Description": "BIOS Attributes",
			"AttributeRegistry": "BiosAttributeRegistryP89.v1_0_0",
			"Attributes": {
				"AdminPhone": "",
				"BootMode": "Uefi",
				"EmbeddedSata": "Raid",
				"NicBoot1": "NetworkBoot",
				"NicBoot2": "Disabled",
				"PowerProfile": "MaxPerf",
				"ProcCoreDisable": 3,
				"ProcHyperthreading": "Enabled",
				"ProcTurboMode": "Enabled",
				"UsbControl": "UsbEnabled",
				"BoolTest1": "True",
				"BoolTest2": 1,
				"BoolTest3": "NotBool"
			},
			"Actions": {
				"#Bios.ResetBios": {
					"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ResetBios"
				},
				"#Bios.ChangePassword": {
					"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ChangePassword"
				}
			}
		}`

// TestBios tests the parsing of Bios objects.
func TestBios(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "BIOS" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "BIOS Configuration Current Attributes" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AttributeRegistry != "BiosAttributeRegistryP89.v1_0_0" {
		t.Errorf("Received incorrect attribute registry: %s", result.AttributeRegistry)
	}

	if result.resetBiosTarget != "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ResetBios" {
		t.Errorf("Invalid ResetBios link: %s", result.resetBiosTarget)
	}

	if result.changePasswordTarget != "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ChangePassword" {
		t.Errorf("Invalid ChangePassword target: %s", result.changePasswordTarget)
	}

	if result.Attributes.String("AdminPhone") != "" {
		t.Errorf("Invalid 'AdminPhone' attribute: %s", result.Attributes["AdminPhone"])
	}

	if result.Attributes.String("PowerProfile") != "MaxPerf" {
		t.Errorf("Invalid 'PowerProfile' attribute: %s", result.Attributes["PowerProfile"])
	}

	if result.Attributes.Int("ProcCoreDisable") != 3 {
		t.Errorf("Invalid 'ProcCoreDisable' attribute: %v", result.Attributes["ProcCoreDisable"])
	}

	if !result.Attributes.Bool("BoolTest1") {
		t.Errorf("Expected True boolean value for 'BoolTest1': %v", result.Attributes["BoolTest1"])
	}

	if !result.Attributes.Bool("BoolTest2") {
		t.Errorf("Expected True boolean value for 'BoolTest2': %v", result.Attributes["BoolTest1"])
	}

	if result.Attributes.Bool("BoolTest3") {
		t.Errorf("Expected False boolean value for 'BoolTest3': %v", result.Attributes["BoolTest1"])
	}

	if len(result.settingsApplyTimes) != 3 {
		t.Errorf("Invalid settings support apply times: %s", result.settingsApplyTimes)
	}

	if result.activeSoftwareImage != "/redfish/v1/Systems/437XR1138R2/BIOS/FirmwareInventory" {
		t.Errorf("Invalid value of activeSoftwareImage: '%s'", result.activeSoftwareImage)
	}
}

// TestBiosAttributes tests the parsing of Bios objects @Redfish.Attributes.
func TestBiosAttributes(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.settingsTarget != "/redfish/v1/Systems/System.Embedded.1/Bios/Settings" {
		t.Errorf("Invalid settings update target: %s", result.settingsTarget)
	}
}

// TestBiosNoAttributes tests the parsing of Bios objects without @Redfish.Attributes.
func TestBiosNoAttributes(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosNoAttributesBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.settingsTarget != "/redfish/v1/Systems/437XR1138R2/BIOS" {
		t.Errorf("Invalid settings update target: %s", result.settingsTarget)
	}
}

// TestUpdateBiosAttributes tests the UpdateBiosAttributes call.
func TestUpdateBiosAttributes(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	update := SettingsAttributes{"AssetTag": "test"}
	err = result.UpdateBiosAttributes(update)

	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 2 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[1].Payload, "AssetTag") {
		t.Errorf("Unexpected update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[1].Payload, "@Redfish.SettingsApplyTime") {
		t.Error("Expected 'SettingsApplyTime' to not be present")
	}
}

// TestUpdateBiosAttributesApplyAt tests the TestUpdateBiosAttributesApplyAt call.
func TestUpdateBiosAttributesApplyAt(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	update := SettingsAttributes{"AssetTag": "test"}
	err = result.UpdateBiosAttributesApplyAt(update, AtMaintenanceWindowStartSettingsApplyTime)

	if err != nil {
		t.Errorf("Error making UpdateBiosAttributesApplyAt call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 2 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[1].Payload, "AssetTag") {
		t.Errorf("Unexpected update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[1].Payload, "@Redfish.SettingsApplyTime") {
		t.Error("Expected 'SettingsApplyTime' to be present")
	}
}

// TestUpdateBiosAttributesApplyAtWithTask tests that UpdateBiosAttributesApplyAtWithTask
// returns a TaskMonitorInfo when the service processes the update asynchronously.
func TestUpdateBiosAttributesApplyAtWithTask(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPatch: {
				&http.Response{
					StatusCode: http.StatusAccepted,
					Header:     http.Header{"Location": []string{"/redfish/v1/TaskService/Tasks/1"}},
					Body:       io.NopCloser(bytes.NewBufferString("")),
				},
			},
		},
	}
	result.SetClient(testClient)

	update := SettingsAttributes{"AssetTag": "test"}
	taskInfo, err := result.UpdateBiosAttributesApplyAtWithTask(update, AtMaintenanceWindowStartSettingsApplyTime)
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributesApplyAtWithTask call: %s", err)
	}

	if taskInfo == nil {
		t.Fatal("Expected a TaskMonitorInfo to be returned, got nil")
	}
	if taskInfo.TaskMonitor != "/redfish/v1/TaskService/Tasks/1" {
		t.Errorf("Unexpected task monitor URI: %s", taskInfo.TaskMonitor)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Errorf("Expected two calls to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[1].Payload, "AssetTag") {
		t.Errorf("Unexpected update payload: %s", calls[1].Payload)
	}

	if !strings.Contains(calls[1].Payload, "@Redfish.SettingsApplyTime") {
		t.Error("Expected 'SettingsApplyTime' to be present")
	}
}

// TestUpdateBiosAttributesApplyAtWithTaskSync tests that UpdateBiosAttributesApplyAtWithTask
// returns a nil TaskMonitorInfo when the service applies the update synchronously.
func TestUpdateBiosAttributesApplyAtWithTaskSync(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	update := SettingsAttributes{"AssetTag": "test"}
	taskInfo, err := result.UpdateBiosAttributesApplyAtWithTask(update, AtMaintenanceWindowStartSettingsApplyTime)
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributesApplyAtWithTask call: %s", err)
	}

	if taskInfo != nil {
		t.Errorf("Expected no TaskMonitorInfo for a synchronous update, got: %v", taskInfo)
	}
}

// biosSettingsBody is a settings object reporting a pending BootMode that differs from the
// applied value in biosBody, plus a slice-valued attribute.
var biosSettingsBody = `{
		"@odata.type": "#Bios.v1_0_6.Bios",
		"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Bios/Settings",
		"Id": "Settings",
		"Name": "BIOS Configuration Pending Settings",
		"Attributes": {
			"BootMode": "Legacy",
			"BootOrder": ["Hdd", "Pxe"]
		}
	}`

// settingsClient returns a TestClient whose first GET answers with the given settings object
// body, as the settings target read does before an attribute update.
func settingsClient(body string) *TestClient {
	return &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(body)),
				},
			},
		},
	}
}

// TestUpdateBiosAttributesComparesAgainstSettingsTarget tests that an attribute is compared
// against the value staged on the settings target being written to, not the applied value on
// the Bios resource. Requesting the applied value for an attribute with something else staged
// is a change to what is staged, so it has to be sent.
func TestUpdateBiosAttributesComparesAgainstSettingsTarget(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := settingsClient(biosSettingsBody)
	result.SetClient(testClient)

	// BootMode is applied as "Uefi" in biosBody but staged as "Legacy" in biosSettingsBody.
	err = result.UpdateBiosAttributes(SettingsAttributes{"BootMode": "Uefi"})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("Expected two calls to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[1].Payload, "BootMode") {
		t.Errorf("Expected BootMode to be sent to replace the staged value, payload: %s", calls[1].Payload)
	}
}

// TestUpdateBiosAttributesOmitsAttributeMatchingSettingsTarget tests that an attribute whose
// requested value already matches what the settings target reports is omitted, since sending it
// would not change what is staged.
func TestUpdateBiosAttributesOmitsAttributeMatchingSettingsTarget(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := settingsClient(biosSettingsBody)
	result.SetClient(testClient)

	// BootMode is already staged as "Legacy", AssetTag is staged nowhere.
	err = result.UpdateBiosAttributes(SettingsAttributes{"BootMode": "Legacy", "AssetTag": "test"})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("Expected two calls to be made, captured: %v", calls)
	}

	if strings.Contains(calls[1].Payload, "BootMode") {
		t.Errorf("Expected the already staged BootMode to be omitted, payload: %s", calls[1].Payload)
	}
	if !strings.Contains(calls[1].Payload, "AssetTag") {
		t.Errorf("Expected AssetTag to be sent, payload: %s", calls[1].Payload)
	}
}

// TestUpdateBiosAttributesSliceValued tests that a slice-valued attribute can be compared
// without panicking, and is sent when its contents differ from what is staged.
func TestUpdateBiosAttributesSliceValued(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := settingsClient(biosSettingsBody)
	result.SetClient(testClient)

	// BootOrder is staged as ["Hdd", "Pxe"].
	err = result.UpdateBiosAttributes(SettingsAttributes{"BootOrder": []any{"Pxe", "Hdd"}})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("Expected two calls to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[1].Payload, "BootOrder") {
		t.Errorf("Expected the reordered BootOrder to be sent, payload: %s", calls[1].Payload)
	}
}

// TestUpdateBiosAttributesSliceValuedUnchanged tests that a slice-valued attribute matching
// what is staged is omitted rather than compared with !=, which would panic.
func TestUpdateBiosAttributesSliceValuedUnchanged(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := settingsClient(biosSettingsBody)
	result.SetClient(testClient)

	err = result.UpdateBiosAttributes(SettingsAttributes{"BootOrder": []any{"Hdd", "Pxe"}})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected only the settings target read, captured: %v", calls)
	}
}

// TestUpdateBiosAttributesUnreadableSettingsTarget tests that a settings target which does not
// report its attributes falls back to comparing against the applied values.
func TestUpdateBiosAttributesUnreadableSettingsTarget(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := settingsClient("")
	result.SetClient(testClient)

	// BootMode is applied as "Uefi" in biosBody, AssetTag is not set at all.
	err = result.UpdateBiosAttributes(SettingsAttributes{"BootMode": "Uefi", "AssetTag": "test"})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("Expected two calls to be made, captured: %v", calls)
	}

	if strings.Contains(calls[1].Payload, "BootMode") {
		t.Errorf("Expected BootMode matching the applied value to be omitted, payload: %s", calls[1].Payload)
	}
	if !strings.Contains(calls[1].Payload, "AssetTag") {
		t.Errorf("Expected AssetTag to be sent, payload: %s", calls[1].Payload)
	}
}

// biosSliceAttributeBody is a Bios resource whose applied attributes include a slice-valued
// one, and which has no settings object so updates are written to the resource itself.
var biosSliceAttributeBody = `{
		"@odata.type": "#Bios.v1_0_6.Bios",
		"@odata.id": "/redfish/v1/Systems/437XR1138R2/BIOS",
		"Id": "BIOS",
		"Name": "BIOS Configuration Current Attributes",
		"Attributes": {
			"BootMode": "Uefi",
			"BootOrder": ["Hdd", "Pxe"]
		}
	}`

// TestUpdateBiosAttributesSliceValuedAppliedUnchanged tests that a slice-valued attribute is
// compared against an applied value of the same type without panicking, and is omitted when
// its contents match.
func TestUpdateBiosAttributesSliceValuedAppliedUnchanged(t *testing.T) {
	var result Bios
	err := json.NewDecoder(strings.NewReader(biosSliceAttributeBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	// An empty settings target read falls the comparison back to the applied attributes.
	testClient := settingsClient("")
	result.SetClient(testClient)

	err = result.UpdateBiosAttributes(SettingsAttributes{"BootOrder": []any{"Hdd", "Pxe"}})
	if err != nil {
		t.Errorf("Error making UpdateBiosAttributes call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected only the settings target read, captured: %v", calls)
	}
}
