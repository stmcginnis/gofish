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

var networkInterfaceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#NetworkInterface.NetworkInterface",
		"@odata.type": "#NetworkInterface.v1_1_2.NetworkInterface",
		"@odata.id": "/redfish/v1/NetworkInterface",
		"Id": "NetworkInterface-1",
		"Name": "NetworkInterfaceOne",
		"Description": "NetworkInterface One",
		"Links": {
			"NetworkAdapter": {
					"@odata.id": "/redfish/v1/NetworkAdapters/1"
			}
		},
		"NetworkDeviceFunctions": {
			"@odata.id": "/redfish/v1/NetworkFunctions"
		},
		"NetworkPorts": {
			"@odata.id": "/redfish/v1/NetworkPorts"
		},
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestNetworkInterface tests the parsing of NetworkInterface objects.
func TestNetworkInterface(t *testing.T) {
	var result NetworkInterface
	err := json.NewDecoder(networkInterfaceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NetworkInterface-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "NetworkInterfaceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}
}
