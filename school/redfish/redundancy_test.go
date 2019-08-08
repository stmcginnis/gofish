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

var redundancyBody = strings.NewReader(
	`{
		"@odata.id": "/redfish/v1/Redundancy",
		"Id": "Redundancy-1",
		"Name": "RedundancyOne",
		"MaxNumSupported": 2,
		"MemberId": "Redundancy1",
		"MinNumNeeded": 2,
		"Mode": "Sparing",
		"RedundancyEnabled": true,
		"RedundancySet": [
			"One",
			"Two"
		],
		"RedundancySet@odata.count": 2,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestRedundancy tests the parsing of Redundancy objects.
func TestRedundancy(t *testing.T) {
	var result Redundancy
	err := json.NewDecoder(redundancyBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Redundancy-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "RedundancyOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaxNumSupported != 2 {
		t.Errorf("Invalid MaxNumSupported: %d", result.MaxNumSupported)
	}

	if result.Mode != SparingRedundancyMode {
		t.Errorf("Invalid mode: %s", result.Mode)
	}
}
