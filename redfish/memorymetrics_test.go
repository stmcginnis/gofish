//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryMetricsBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#MemoryMetrics.MemoryMetrics",
		"@odata.type": "#MemoryMetrics.v1_0_0.MemoryMetrics",
		"@odata.id": "/redfish/v1/MemoryMetrics",
		"Id": "MemoryMetrics-1",
		"Name": "MemoryMetricsOne",
		"Description": "MemoryMetrics One",
		"BlockSizeBytes": 512,
		"CurrentPeriod": {
			"BlocksRead": 123456,
			"BlockWritten": 54321
		},
		"HealthData": {
			"AlarmTrips": {
				"AddressParityError": false,
				"CorrectableECCError": false,
				"SpareBlock": false,
				"Temperature": true,
				"UncorrectableECCError": false
			},
			"DataLossDetected": false,
			"LastShutdownSuccess": true,
			"PerformanceDegraded": false,
			"PredictedMediaLifeLeftPercent": 85,
			"RemainingSpareblockPercentage": 95
		},
		"LifeTime": {
			"BlocksRead": 1234567890,
			"BlocksWritten": 9876543210
		}
	}`)

// TestMemoryMetrics tests the parsing of MemoryMetrics objects.
func TestMemoryMetrics(t *testing.T) {
	var result MemoryMetrics
	err := json.NewDecoder(memoryMetricsBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "MemoryMetrics-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "MemoryMetricsOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.CurrentPeriod.BlocksRead != 123456 {
		t.Errorf("Invalid current period blocks read: %d", result.CurrentPeriod.BlocksRead)
	}

	if result.HealthData.AlarmTrips.AddressParityError {
		t.Error("Address parity error should be false")
	}

	if result.HealthData.DataLossDetected {
		t.Error("Data loss detected should be false")
	}

	if result.LifeTime.BlocksWritten != 9876543210 {
		t.Errorf("Invalid lifetime blocks written: %d", result.LifeTime.BlocksWritten)
	}
}
