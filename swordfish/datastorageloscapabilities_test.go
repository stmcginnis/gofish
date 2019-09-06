//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var dataStorageLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataStorageLoSCapabilities.DataStorageLoSCapabilities",
		"@odata.type": "#DataStorageLoSCapabilities.v1_2_0.DataStorageLoSCapabilities",
		"@odata.id": "/redfish/v1/DataStorageLoSCapabilities",
		"Id": "DataStorageLoSCapabilities-1",
		"Name": "DataStorageLoSCapabilitiesOne",
		"Description": "DataStorageLoSCapabilities One",
		"MaximumRecoverableCapacitySourceCount": 5,
		"SupportedAccessCapabilities": [
			"Read",
			"Write",
			"Append",
			"Streaming"
		],
		"SupportedLinesOfService": [{
				"@odata.id": "/redfish/v1/DataStorageLineOfService/1"
			},
			{
				"@odata.id": "/redfish/v1/DataStorageLineOfService/2"
			}
		],
		"SupportedLinesOfService@odata.count": 2,
		"SupportedProvisioningPolicies": [
			"Fixed",
			"Thin"
		],
		"SupportedRecoveryTimeObjectives": [
			"OnlineActive",
			"OnlinePassive",
			"Nearline",
			"Offline"
		],
		"SupportsSpaceEfficiency": true
	}`)

// TestDataStorageLoSCapabilities tests the parsing of DataStorageLoSCapabilities objects.
func TestDataStorageLoSCapabilities(t *testing.T) {
	var result DataStorageLoSCapabilities
	err := json.NewDecoder(dataStorageLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataStorageLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataStorageLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaximumRecoverableCapacitySourceCount != 5 {
		t.Errorf("Invalid MaximumRecoverableCapacitySource: %d",
			result.MaximumRecoverableCapacitySourceCount)
	}

	if result.SupportedAccessCapabilities[1] != WriteStorageAccessCapability {
		t.Errorf("Invalid SupportedAccessCapability: %s", result.SupportedAccessCapabilities[1])
	}

	if result.SupportedProvisioningPolicies[0] != FixedProvisioningPolicy {
		t.Errorf("Invalid SupportedProvisioningPolicy: %s", result.SupportedAccessCapabilities[0])
	}

	if !result.SupportsSpaceEfficiency {
		t.Error("SupportsSpaceEfficiency should be true")
	}
}
