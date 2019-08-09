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

var ioConnectivityLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#IOConnectivityLoSCapabilities.IOConnectivityLoSCapabilities",
		"@odata.type": "#IOConnectivityLoSCapabilities.v1_1_1.IOConnectivityLoSCapabilities",
		"@odata.id": "/redfish/v1/IOConnectivityLoSCapabilities",
		"Id": "IOConnectivityLoSCapabilities-1",
		"Name": "IOConnectivityLoSCapabilitiesOne",
		"Description": "IOConnectivityLoSCapabilities One",
		"MaxSupportedBytesPerSecond": 5000000000,
		"MaxSupportedIOPS": 1000000000,
		"SupportedAccessProtocols": [
			"FC",
			"FCP",
			"FCoE",
			"iSCSI"
		],
		"SupportedLinesOfService": [{
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
			},
			{
				"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
				"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
				"Id": "IOConnectivityLineOfService-2",
				"Name": "IOConnectivityLineOfServiceTwo",
				"Description": "IOConnectivityLineOfService Two",
				"AccessProtocols": [
					"FC",
					"FCP",
					"FCoE"
				],
				"MaxBytesPerSecond": 5000000000,
				"MaxIOPS": 1000000000
			}
		]
	}`)

// TestIOConnectivityLoSCapabilities tests the parsing of IOConnectivityLoSCapabilities objects.
func TestIOConnectivityLoSCapabilities(t *testing.T) {
	var result IOConnectivityLoSCapabilities
	err := json.NewDecoder(ioConnectivityLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOConnectivityLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOConnectivityLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaxSupportedBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.MaxSupportedBytesPerSecond)
	}

	if result.MaxSupportedIOPS != 1000000000 {
		t.Errorf("MaxSupportedIOPS: %d", result.MaxSupportedIOPS)
	}

	if result.SupportedAccessProtocols[1] != common.FCPProtocol {
		t.Errorf("Invalid AccessProtocol: %s", result.SupportedAccessProtocols[1])
	}

	if result.SupportedLinesOfService[0].MaxBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.SupportedLinesOfService[0].MaxBytesPerSecond)
	}
}
