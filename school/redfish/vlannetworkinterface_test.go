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

var vlanNetworkInterfaceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#VlanNetworkInterface.VlanNetworkInterface",
		"@odata.type": "#VLanNetworkInterface.v1_1_3.VLanNetworkInterface",
		"@odata.id": "/redfish/v1/VlanNetworkInterface",
		"Id": "VlanNetworkInterface-1",
		"Name": "VlanNetworkInterfaceOne",
		"Description": "VlanNetworkInterface One",
		"VLANEnable": true,
		"VLANId": 200
	}`)

// TestVlanNetworkInterface tests the parsing of VlanNetworkInterface objects.
func TestVlanNetworkInterface(t *testing.T) {
	var result VLanNetworkInterface
	err := json.NewDecoder(vlanNetworkInterfaceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "VlanNetworkInterface-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "VlanNetworkInterfaceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.VLANEnable {
		t.Error("VLAN should be enabled")
	}

	if result.VLANID != 200 {
		t.Errorf("Invalid VLAN ID: %d", result.VLANID)
	}
}
