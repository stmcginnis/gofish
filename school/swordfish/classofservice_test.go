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

var classOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ClassOfService.ClassOfService",
		"@odata.type": "#ClassOfService.v1_0_0.ClassOfService",
		"@odata.id": "/redfish/v1/ClassOfService",
		"Id": "ClassOfService-1",
		"Name": "ClassOfServiceOne",
		"Description": "ClassOfService One",
		"ClassOfServiceVersion": "1.0.0",
		"DataProtectionLinesOfService": [
			{
				"@odata.id": "/redfish/v1/DataProtectionLineOfService/1"
			}
		],
		"DataProtectionLinesOfService@odata.count": 1,
		"DataSecurityLinesOfService": [
			{
				"@odata.id": "/redfish/v1/DataSecurityLineOfService/1"
			}
		],
		"DataSecurityLinesOfService@odata.count": 1,
		"DataStorageLinesOfService": [
			{
				"@odata.id": "/redfish/v1/DataStorageLineOfService/1"
			}
		],
		"DataStorageLinesOfService@odata.count": 1,
		"IOConnectivityLinesOfService": [
			{
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService/1"
			}
		],
		"IOConnectivityLinesOfService@odata.count": 1,
		"IOPerformanceLinesOfService": [
			{
				"@odata.id": "/redfish/v1/IOPerformanceLineOfService/1"
			}
		],
		"IOPerformanceLinesOfService@odata.count": 1
	}`)

// TestClassOfService tests the parsing of ClassOfService objects.
func TestClassOfService(t *testing.T) {
	var result ClassOfService
	err := json.NewDecoder(classOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "ClassOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "ClassOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.dataProtectionLinesOfService[0] != "/redfish/v1/DataProtectionLineOfService/1" {
		t.Errorf("Invalid DataProtectionLineOfService link: %s", result.dataProtectionLinesOfService[0])
	}
}
