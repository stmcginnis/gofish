//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var driveBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Drive.Drive",
		"@odata.type": "#Drive.v1_0_0.Drive",
		"@odata.id": "/redfish/v1/Drive",
		"Id": "Drive-1",
		"Name": "Drive One",
		"Description": "One drive",
		"Assembly": {
			"@odata.id": "/redfish/v1/Assembly/Assembly-1"
		},
		"AssetTag": "Asset 1",
		"BlockSizeBytes": 512,
		"CapableSpeedGbs": 40,
		"CapacityBytes": 1099511627776,
		"EncryptionAbility": "SelfEncryptingDrive",
		"EncryptionStatus": "Unlocked",
		"FailurePredicted": false,
		"HotSpareMode": "Revertible",
		"HotSpareType": "Chassis",
		"Identifiers": [
			{
				"DurableName": "5000D3100101D52E",
				"DurableNameFormat": "FC_WWN"
			}
		],
		"IndicatorLED": "Blinking",
		"Links": {
			"Chassis": {
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			},
			"Endpoints": [],
			"Endpoints@odata.count": 0,
			"PCIeFunctions": [
				{
					"@odata.id": "/redfish/v1/PCIeFunctions/PCIeFunction-1"
				}
			],
			"PCIeFunctions@odata.count": 1,
			"Volumes": [
				{
					"@odata.id": "/redfish/v1/Volumes/Volume-1"
				}
			],
			"Volumes@odata.count": 1
		},
		"Manufacturer": "Joe's Storage",
		"MediaType": "SSD",
		"Model": "Storage One",
		"NegotiatedSpeedGbps": 10,
		"Operations": [],
		"PartNumber": "12345",
		"PhysicalLocation": {},
		"PredictedMediaLifeLeftPercent": 100,
		"Protocol": "FC",
		"Revision": "2.0",
		"RotationSpeedRPM": 5200,
		"SKU": "123456",
		"SerialNumber": "1234567",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"StatusIndicator": "Hotspare"
	}`)

// TestDrive tests the parsing of Drive objects.
func TestDrive(t *testing.T) {
	var result Drive
	err := json.NewDecoder(driveBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Drive-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Drive One" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.assembly != "/redfish/v1/Assembly/Assembly-1" {
		t.Errorf("Incorrect assembly link: %s", result.assembly)
	}

	if result.BlockSizeBytes != 512 {
		t.Errorf("Incorrect block size bytes: %d", result.BlockSizeBytes)
	}

	if result.CapableSpeedGbs != 40 {
		t.Errorf("Incorrect capable speed: %d", result.CapableSpeedGbs)
	}

	if result.CapacityBytes != 1099511627776 {
		t.Errorf("Incorrect capacity: %d", result.CapacityBytes)
	}

	if result.EncryptionAbility != "SelfEncryptingDrive" {
		t.Errorf("Incorrect encryption ability: %s", result.EncryptionAbility)
	}

	if result.EncryptionStatus != "Unlocked" {
		t.Errorf("Incorrect encryption status: %s", result.EncryptionStatus)
	}

	if result.FailurePredicted {
		t.Error("Failure predicted should be false.")
	}

	if result.chassis != "/redfish/v1/Chassis/Chassis-1" {
		t.Errorf("Invalid chassis link: %s", result.chassis)
	}
}
