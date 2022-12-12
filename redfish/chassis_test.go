//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

const TestAssetTag = "TestAssetTag"
const TestChassisPath = "/redfish/v1/Chassis/Chassis-1"

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
		"Assembly": {
			"@odata.id": "/redfish/v1/Chassis/Chassis-1/Assembly"
		},
		"Drives": {
			"@odata.id": "/redfish/v1/Chassis/Chassis-1/Drives"
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

var supermicroRAIDChassisBody = `{
    "@odata.type": "#Chassis.v1_9_1.Chassis",
    "@odata.id": "/redfish/v1/Chassis/HA-RAID.0.StorageEnclosure.0",
    "Id": "HA-RAID.0.StorageEnclosure.0",
    "Name": "Internal Enclosure 0",
    "ChassisType": "Enclosure",
    "Model": "Internal Enclosure",
    "SerialNumber": "",
    "PartNumber": "",
    "Links": {
        "ManagedBy": [
            {
                "@odata.id": "/redfish/v1/Managers/1"
            }
        ],
        "Storage": [
            {
                "@odata.id": "/redfish/v1/Systems/1/Storage/HA-RAID"
            }
        ],
        "Drives": [
            {
                "@odata.id": "/redfish/v1/Chassis/HA-RAID.0.StorageEnclosure.0/Drives/Disk.Bay.0"
            },
            {
                "@odata.id": "/redfish/v1/Chassis/HA-RAID.0.StorageEnclosure.0/Drives/Disk.Bay.1"
            },
            {
                "@odata.id": "/redfish/v1/Chassis/HA-RAID.0.StorageEnclosure.0/Drives/Disk.Bay.2"
            },
            {
                "@odata.id": "/redfish/v1/Chassis/HA-RAID.0.StorageEnclosure.0/Drives/Disk.Bay.3"
            }
        ]
    },
    "Oem": {}
}`

var driveCollection = `{
    "@odata.id": "/redfish/v1/Chassis/IPAttachedDrive/Drives",
    "@odata.type": "#DriveCollection.DriveCollection",
    "Members": [
        {
            "@odata.id": "/redfish/v1/Chassis/IPAttachedDrive/Drives/IPAttachedDrive"
        }
    ],
    "Members@odata.count": 1,
    "Name": "Drive Collection"
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

	if result.assembly != "/redfish/v1/Chassis/Chassis-1/Assembly" {
		t.Errorf("Received invalid assembly reference: %s", result.assembly)
	}

	if result.drives != "/redfish/v1/Chassis/Chassis-1/Drives" {
		t.Errorf("Received invalid drive reference: %s", result.drives)
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

// TestMinimumChassis tests a failure we had from how SM returns a RAID
// controller chassis.
//
// The required properties according to the spec are:
// "required": [
//
//	"ChassisType",
//	"@odata.id",
//	"@odata.type",
//	"Id",
//	"Name"]
func TestMinimumChassis(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(supermicroRAIDChassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "HA-RAID.0.StorageEnclosure.0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Internal Enclosure 0" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ChassisType != EnclosureChassisType {
		t.Errorf("Received invalid chassis type: %s", result.ChassisType)
	}
}

// TestLinkedDriveChassis tests getting drives from versions supporting the older
// Chassis.Links.Drives location.
func TestLinkedDriveChassis(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(supermicroRAIDChassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if len(result.linkedDrives) != 4 {
		t.Errorf("Expected 3 drive links: %v", result.linkedDrives)
	}

	if result.drives != "" {
		t.Errorf("Expected drives link to be empty, got %q", result.drives)
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

	result.AssetTag = TestAssetTag
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

// getCall returns an http.Response for a GET request.
func getCall(body string) *http.Response {
	return &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header),
	}
}

// TestChassisDrives tests getting the drives for a chassis.
func TestChassisDrives(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(chassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				getCall(driveCollection), //nolint
				getCall(driveBody),       //nolint
			},
		},
	}
	result.SetClient(testClient)

	drives, err := result.Drives()
	if err != nil {
		t.Errorf("Error getting drives: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 2 {
		t.Errorf("Expected two calls to be made, captured: %v", calls)
	}

	if len(drives) != 1 {
		t.Errorf("Expected 1 drive to be returned, got %d", len(drives))
	}
}

// TestChassisLinkedDrives tests getting the drives returned through the links for a chassis.
func TestChassisLinkedDrives(t *testing.T) {
	var result Chassis
	err := json.NewDecoder(strings.NewReader(supermicroRAIDChassisBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				getCall(driveBody), //nolint
				getCall(driveBody), //nolint
				getCall(driveBody), //nolint
				getCall(driveBody), //nolint
			},
		},
	}
	result.SetClient(testClient)

	drives, err := result.Drives()
	if err != nil {
		t.Errorf("Error getting drives: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 4 {
		t.Errorf("Expected four calls to be made, captured: %v", calls)
	}

	if len(drives) != 4 {
		t.Errorf("Expected 4 drives to be returned, got %d", len(drives))
	}
}
