//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var processorMetricsBody = `{
	"@odata.type": "#ProcessorMetrics.v1_6_3.ProcessorMetrics",
	"Id": "Metrics",
	"Name": "Processor Metrics",
	"BandwidthPercent": 62,
	"OperatingSpeedMHz": 2400,
	"ThrottlingCelsius": 65,
	"FrequencyRatio": 0.00432,
	"Cache": [
	  {
		"Level": "3",
		"CacheMiss": 0.12,
		"HitRatio": 0.719,
		"CacheMissesPerInstruction": 0.00088,
		"OccupancyBytes": 3030144,
		"OccupancyPercent": 90.1
	  }
	],
	"LocalMemoryBandwidthBytes": 18253611008,
	"RemoteMemoryBandwidthBytes": 81788928,
	"KernelPercent": 2.3,
	"UserPercent": 34.7,
	"CoreMetrics": [
	  {
		"CoreId": "core0",
		"InstructionsPerCycle": 1.16,
		"UnhaltedCycles": 6254383746,
		"MemoryStallCount": 58372,
		"IOStallCount": 2634872,
		"CoreCache": [
		  {
			"Level": "2",
			"CacheMiss": 0.472,
			"HitRatio": 0.57,
			"CacheMissesPerInstruction": 0.00346,
			"OccupancyBytes": 198231,
			"OccupancyPercent": 77.4
		  }
		],
		"CStateResidency": [
		  {
			"Level": "C0",
			"Residency": 1.13
		  },
		  {
			"Level": "C1",
			"Residency": 26
		  },
		  {
			"Level": "C3",
			"Residency": 0.00878
		  },
		  {
			"Level": "C6",
			"Residency": 0.361
		  },
		  {
			"Level": "C7",
			"Residency": 72.5
		  }
		]
	  }
	],
	"@odata.id": "/redfish/v1/Systems/1/Processors/FPGA1/ProcessorMetrics"
  }`

// TestProcessorMetrics tests the parsing of ProcessorMetrics objects.
func TestProcessorMetrics(t *testing.T) {
	var result ProcessorMetrics
	err := json.NewDecoder(strings.NewReader(processorMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Processor Metrics", result.Name)
	assertEquals(t, "core0", result.CoreMetrics[0].CoreID)

	if result.BandwidthPercent != 62 {
		t.Errorf("Unexpected BandwidthPercent: %.2f", result.BandwidthPercent)
	}

	if result.FrequencyRatio != 0.00432 {
		t.Errorf("Unexpected FrequencyRatio: %.5f", result.FrequencyRatio)
	}

	if result.Cache[0].Level != "3" {
		t.Errorf("Unexpected Cache[0].Level: %s", result.Cache[0].Level)
	}

	if result.Cache[0].CacheMissesPerInstruction != 0.00088 {
		t.Errorf("Unexpected Cache[0].CacheMissesPerInstruction: %.5f", result.Cache[0].CacheMissesPerInstruction)
	}

	if result.CoreMetrics[0].CoreCache[0].CacheMiss != 0.472 {
		t.Errorf("Unexpected CoreMetrics[0].CoreCache[0].CacheMiss: %.5f", result.CoreMetrics[0].CoreCache[0].CacheMiss)
	}
}
