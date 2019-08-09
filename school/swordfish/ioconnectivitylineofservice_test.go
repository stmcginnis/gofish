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

	"github.com/stmcginnis/gofish/school/common"
)

var ioConnectivityLineOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
		"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
		"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
		"Id": "IOConnectivityLineOfService-1",
		"Name": "IOConnectivityLineOfServiceOne",
		"Description": "IOConnectivityLineOfService One",
		"AccessProtocols": [
			"FC",
			"FCP",
			"FCoE",
			"iSCSI"
		],
		"MaxBytesPerSecond": 5000000000,
		"MaxIOPS": 1000000000
	}`)

// TestIOConnectivityLineOfService tests the parsing of IOConnectivityLineOfService objects.
func TestIOConnectivityLineOfService(t *testing.T) {
	var result IOConnectivityLineOfService
	err := json.NewDecoder(ioConnectivityLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOConnectivityLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOConnectivityLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessProtocols[3] != common.ISCSIProtocol {
		t.Errorf("Invalid access protocol: %s", result.AccessProtocols[3])
	}

	if result.MaxBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxBytesPerSecond: %d", result.MaxBytesPerSecond)
	}

	if result.MaxIOPS != 1000000000 {
		t.Errorf("Invalid MaxIOPS: %d", result.MaxIOPS)
	}
}
