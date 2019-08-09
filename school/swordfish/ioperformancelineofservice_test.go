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

var ioPerformanceLineOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#IOPerformanceLineOfService.IOPerformanceLineOfService",
		"@odata.type": "#IOPerformanceLineOfService.v1_0_2.IOPerformanceLineOfService",
		"@odata.id": "/redfish/v1/IOPerformanceLineOfService",
		"Id": "IOPerformanceLineOfService-1",
		"Name": "IOPerformanceLineOfServiceOne",
		"Description": "IOPerformanceLineOfService One",
		"AverageIOOperationLatencyMicroseconds": 500,
		"IOOperationsPerSecondIsLimited": true,
		"IOWorkload": {
			"Components": [{
				"AverageIOBytes": 102400,
				"Duration": "P3Y6M4DT12H30M5S",
				"IOAccessPattern": "RandomReadNew",
				"PercentOfData": 99,
				"PercentOfIOPS": 99,
				"Schedule": {
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
				}
			}]
		},
		"MaxIOOperationsPerSecondPerTerabyte": 1000,
		"SamplePeriod": "ISO8601Duration"
	}`)

// TestIOPerformanceLineOfService tests the parsing of IOPerformanceLineOfService objects.
func TestIOPerformanceLineOfService(t *testing.T) {
	var result IOPerformanceLineOfService
	err := json.NewDecoder(ioPerformanceLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOPerformanceLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOPerformanceLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.IOOperationsPerSecondIsLimited {
		t.Error("IOOperationsPerSecondIsLimited should be true")
	}

	if result.MaxIOOperationsPerSecondPerTerabyte != 1000 {
		t.Errorf("Invalid MaxIOOperationsPerSecondPerTerabyte: %d",
			result.MaxIOOperationsPerSecondPerTerabyte)
	}
}
