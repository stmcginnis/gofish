//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var dataProtectionLineOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataProtectionLineOfService.DataProtectionLineOfService",
		"@odata.type": "#DataProtectionLineOfService.v1_0_0.DataProtectionLineOfService",
		"@odata.id": "/redfish/v1/DataProtectionLineOfService",
		"Id": "DataProtectionLineOfService-1",
		"Name": "DataProtectionLineOfServiceOne",
		"Description": "DataProtectionLineOfService One",
		"IsIsolated": true,
		"MinLifetime": "P3Y6M4DT12H30M5S",
		"RecoveryGeographicObjective": "Row",
		"RecoveryPointObjective": "P0Y0M0DT0H30M0S",
		"RecoveryTimeObjective": "OnlinePassive",
		"ReplicaClassOfService": {
			"@odata.id": "/redfish/v1/ClassOfService"
		},
		"ReplicaType": "Clone",
		"Schedule": {
			"EnabledDaysOfMonth": [1, 5, 10, 15, 20, 25, 30],
			"EnabledDaysOfWeek": [
				"Monday",
				"Tuesday",
				"Wednesday",
				"Thursday",
				"Friday"
			],
			"EnabledIntervals": ["R5/2008-03-01T13:00:00Z/P1Y2M10DT2H30M"],
			"EnabledMonthsOfYear": [
				"January",
				"February",
				"March",
				"April",
				"May",
				"June",
				"July",
				"August",
				"September",
				"October",
				"November",
				"December"
			],
			"InitialStartTime": "2019-08-09T01:29:45+0000",
			"Lifetime": "P3Y6M4DT12H30M5S",
			"MaxOccurrences": 999,
			"RecurrenceInterval": "P0Y0M0DT0H30M0S"
		},
		"DataProtectionLineOfServiceVersion": "1.0.0",
		"DataProtectionLinesOfService": [{
			"@odata.id": "/redfish/v1/DataProtectionLineOfService/1"
		}],
		"DataProtectionLinesOfService@odata.count": 1,
		"DataSecurityLinesOfService": [{
			"@odata.id": "/redfish/v1/DataSecurityLineOfService/1"
		}],
		"DataSecurityLinesOfService@odata.count": 1,
		"DataStorageLinesOfService": [{
			"@odata.id": "/redfish/v1/DataStorageLineOfService/1"
		}],
		"DataStorageLinesOfService@odata.count": 1,
		"IOConnectivityLinesOfService": [{
			"@odata.id": "/redfish/v1/IOConnectivityLineOfService/1"
		}],
		"IOConnectivityLinesOfService@odata.count": 1,
		"IOPerformanceLinesOfService": [{
			"@odata.id": "/redfish/v1/IOPerformanceLineOfService/1"
		}],
		"IOPerformanceLinesOfService@odata.count": 1
	}`)

// TestDataProtectionLineOfService tests the parsing of DataProtectionLineOfService objects.
func TestDataProtectionLineOfService(t *testing.T) {
	var result DataProtectionLineOfService
	err := json.NewDecoder(dataProtectionLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataProtectionLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataProtectionLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.IsIsolated {
		t.Error("IsIsolated should be true")
	}

	if result.RecoveryGeographicObjective != RowFailureDomainScope {
		t.Errorf("Invalid RecoveryGeographicObjective: %s", result.RecoveryGeographicObjective)
	}

	if result.RecoveryTimeObjective != OnlinePassiveRecoveryAccessScope {
		t.Errorf("Invalid RecoveryTimeObjective: %s", result.RecoveryTimeObjective)
	}
}
