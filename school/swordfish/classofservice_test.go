//
// SPDX-License-Identifier: BSD-3-Clause
//

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
