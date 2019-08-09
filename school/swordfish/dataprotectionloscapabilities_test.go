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

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var dataProtectionLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataProtectionLoSCapabilities.DataProtectionLoSCapabilities",
		"@odata.type": "#DataProtectionLoSCapabilities.v1_1_2.DataProtectionLoSCapabilities",
		"@odata.id": "/redfish/v1/DataProtectionLoSCapabilities",
		"Id": "DataProtectionLoSCapabilities-1",
		"Name": "DataProtectionLoSCapabilitiesOne",
		"Description": "DataProtectionLoSCapabilities One",
		"Links": {
			"SupportedReplicaOptions": [{
				"@odata.id": "/redfish/v1/ClassesOfService/1"
			}],
			"SupportedReplicaOptions@odata.count": 1
		},
		"SupportedLinesOfService": [{
			"@odata.id": "/redfish/v1/LinesOfService/1"
		}],
		"SupportedLinesOfService@odata.count": 1,
		"SupportedMinLifetimes": [
			"P0Y6M0DT0H0M0S"
		],
		"SupportedRecoveryGeographicObjectives": [
			"Datacenter",
			"Region"
		],
		"SupportedRecoveryPointObjectiveTimes": [
			"P0Y0M0DT0H30M0S"
		],
		"SupportedRecoveryTimeObjectives": [
			"OnlinePassive"
		],
		"SupportedReplicaTypes": [
			"Clone"
		],
		"SupportsIsolated": true
	}`)

// TestDataProtectionLoSCapabilities tests the parsing of DataProtectionLoSCapabilities objects.
func TestDataProtectionLoSCapabilities(t *testing.T) {
	var result DataProtectionLoSCapabilities
	err := json.NewDecoder(dataProtectionLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataProtectionLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataProtectionLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.SupportsIsolated {
		t.Error("SupportsIsolated should be true")
	}

	if result.SupportedRecoveryTimeObjectives[0] != OnlinePassiveRecoveryAccessScope {
		t.Errorf("Invalid SupportedRecoveryTimeObjective: %s",
			result.SupportedRecoveryTimeObjectives[0])
	}

	if result.SupportedReplicaTypes[0] != CloneReplicaType {
		t.Errorf("Invalid supported replica type: %s", result.SupportedReplicaTypes[0])
	}
}
