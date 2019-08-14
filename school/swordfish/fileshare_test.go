//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var fileShareBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#FileShare.FileShare",
		"@odata.type": "#FileShare.v1_1_2.FileShare",
		"@odata.id": "/redfish/v1/FileShare",
		"Id": "FileShare-1",
		"Name": "FileShareOne",
		"Description": "FileShare One",
		"CASupported": true,
		"DefaultAccessCapabilities": [
			"Read",
			"Write",
			"Append",
			"Streaming"
		],
		"EthernetInterfaces": {
			"@odata.id": "/redfish/v1/EthernetInterfaces"
		},
		"ExecuteSupport": true,
		"FileSharePath": "/exports/data",
		"FileShareQuotaType": "Hard",
		"FileShareRemainingQuotaBytes": 536870912000,
		"FileShareTotalQuotaBytes": 2147483648000,
		"FileSharingProtocols": [
			"NFSv3",
			"NFSv4_0",
			"NFSv4_1",
			"SMBv3_0"
		],
		"Links": {
			"ClassOfService": {
				"@odata.id": "/redfish/v1/ClassOfService/1"
			},
			"FileSystem": {
				"@odata.id": "/redfish/v1/FileSystem/1"
			}
		},
		"LowSpaceWarningThresholdPercents": [
			15,
			10,
			5
		],
		"RemainingCapacityPercent": 43,
		"RootAccess": false,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"WritePolicy": "Asynchronous"
	}`)

// TestFileShare tests the parsing of FileShare objects.
func TestFileShare(t *testing.T) {
	var result FileShare
	err := json.NewDecoder(fileShareBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "FileShare-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "FileShareOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.CASupported {
		t.Error("CASupported should be true")
	}

	if result.DefaultAccessCapabilities[3] != StreamingStorageAccessCapability {
		t.Errorf("Invalid access capability: %s", result.DefaultAccessCapabilities[3])
	}

	if result.ethernetInterfaces != "/redfish/v1/EthernetInterfaces" {
		t.Errorf("Invalid ethernet interfaces link: %s", result.ethernetInterfaces)
	}

	if !result.ExecuteSupport {
		t.Error("ExecuteSupport should be true")
	}

	if result.FileSharePath != "/exports/data" {
		t.Errorf("Invalid FileSharePath: %s", result.FileSharePath)
	}

	if result.FileShareQuotaType != HardQuotaType {
		t.Errorf("Invalid FileShareQuotaType: %s", result.FileShareQuotaType)
	}

	if result.FileShareRemainingQuotaBytes != 536870912000 {
		t.Errorf("Invalid FileShareRemainingQuotaBytes: %d", result.FileShareRemainingQuotaBytes)
	}

	if result.FileShareTotalQuotaBytes != 2147483648000 {
		t.Errorf("Invalid FileShareTotalQuotaBytes: %d", result.FileShareTotalQuotaBytes)
	}

	if result.FileSharingProtocols[3] != SMBv30FileProtocol {
		t.Errorf("Invalid FileSharingProtocol: %s", result.FileSharingProtocols[3])
	}

	if result.fileSystem != "/redfish/v1/FileSystem/1" {
		t.Errorf("Invalid FileSystem link: %s", result.fileSystem)
	}

	if len(result.LowSpaceWarningThresholdPercents) != 3 {
		t.Errorf("Wrong number of warning threshold percents: %d", len(result.LowSpaceWarningThresholdPercents))
	}

	if result.WritePolicy != AsynchronousReplicaUpdateMode {
		t.Errorf("Invalid WritePolicy: %s", result.WritePolicy)
	}
}
