//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var ioPerformanceLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#IOPerformanceLoSCapabilities.IOPerformanceLoSCapabilities",
		"@odata.type": "#IOPerformanceLoSCapabilities.v1_1_2.IOPerformanceLoSCapabilities",
		"@odata.id": "/redfish/v1/IOPerformanceLoSCapabilities",
		"Id": "IOPerformanceLoSCapabilities-1",
		"Name": "IOPerformanceLoSCapabilitiesOne",
		"Description": "IOPerformanceLoSCapabilities One",
		"IOLimitingIsSupported": true,
		"MaxSamplePeriod": "P3Y6M4DT12H30M5S",
		"MinSamplePeriod": "P0Y0M0DT0H30M5S",
		"MinSupportedIoOperationLatencyMicroseconds": 1000,
		"SupportedIOWorkloads": [{
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
		}],
		"SupportedLinesOfService": [{
				"@odata.context": "/redfish/v1/$metadata#IOPerformanceLineOfService.IOPerformanceLineOfService",
				"@odata.type": "#IOPerformanceLineOfService.v1_1_1.IOPerformanceLineOfService",
				"@odata.id": "/redfish/v1/IOPerformanceLineOfService",
				"Id": "IOPerformanceLineOfService-1",
				"Name": "IOPerformanceLineOfServiceOne",
				"Description": "IOPerformanceLineOfService One",
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
				"@odata.context": "/redfish/v1/$metadata#IOPerformanceLineOfService.IOPerformanceLineOfService",
				"@odata.type": "#IOPerformanceLineOfService.v1_1_1.IOPerformanceLineOfService",
				"@odata.id": "/redfish/v1/IOPerformanceLineOfService",
				"Id": "IOPerformanceLineOfService-2",
				"Name": "IOPerformanceLineOfServiceTwo",
				"Description": "IOPerformanceLineOfService Two",
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

// TestIOPerformanceLoSCapabilities tests the parsing of IOPerformanceLoSCapabilities objects.
func TestIOPerformanceLoSCapabilities(t *testing.T) {
	var result IOPerformanceLoSCapabilities
	err := json.NewDecoder(ioPerformanceLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOPerformanceLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOPerformanceLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.IOLimitingIsSupported {
		t.Error("IOLimitingIsSupported should be true")
	}

	if result.MinSupportedIoOperationLatencyMicroseconds != 1000 {
		t.Errorf("Invalid MinSupportedIoOperationLatencyMicroseconds: %d",
			result.MinSupportedIoOperationLatencyMicroseconds)
	}
}
