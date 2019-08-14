//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var dataStorageLineOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataStorageLineOfService.DataStorageLineOfService",
		"@odata.type": "#DataStorageLineOfService.v1_2_0.DataStorageLineOfService",
		"@odata.id": "/redfish/v1/DataStorageLineOfService",
		"Id": "DataStorageLineOfService-1",
		"Name": "DataStorageLineOfServiceOne",
		"Description": "DataStorageLineOfService One",
		"AccessCapabilities": [
			"Read",
			"Write",
			"Append",
			"Streaming"
		],
		"IsSpaceEfficient": true,
		"ProvisioningPolicy": "Thin",
		"RecoverableCapacitySourceCount": 1,
		"RecoveryTimeObjectives": "Nearline"
	}`)

// TestDataStorageLineOfService tests the parsing of DataStorageLineOfService objects.
func TestDataStorageLineOfService(t *testing.T) {
	var result DataStorageLineOfService
	err := json.NewDecoder(dataStorageLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataStorageLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataStorageLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessCapabilities[2] != AppendStorageAccessCapability {
		t.Errorf("Invalid AccessCapability: %s", result.AccessCapabilities[2])
	}

	if !result.IsSpaceEfficient {
		t.Error("IsSpaceEfficient should be true")
	}

	if result.ProvisioningPolicy != ThinProvisioningPolicy {
		t.Errorf("Invalid provisioning policy: %s", result.ProvisioningPolicy)
	}

	if result.RecoverableCapacitySourceCount != 1 {
		t.Errorf("RecoverableCapacitySource should be 1, was %d", result.RecoverableCapacitySourceCount)
	}

	if result.RecoveryTimeObjectives != NearlineRecoveryAccessScope {
		t.Errorf("RecoveryTimeObjective was: %s", result.RecoveryTimeObjectives)
	}
}
