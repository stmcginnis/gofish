//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkDeviceFunctionMetricsBody = `{
	"@odata.type": "#NetworkDeviceFunctionMetrics.v1_1_2.NetworkDeviceFunctionMetrics",
	"Id": "NetworkDeviceFunctionMetrics",
	"Name": "Network Device Function Metrics",
	"TXAvgQueueDepthPercent": 13.7,
	"RXAvgQueueDepthPercent": 21.2,
	"RXFrames": 27193387,
	"RXBytes": 7754199970,
	"RXUnicastFrames": 26193387,
	"RXMulticastFrames": 1000000,
	"TXFrames": 18205770,
	"TXBytes": 9436506547,
	"TXUnicastFrames": 17205770,
	"TXMulticastFrames": 1000000,
	"TXQueuesEmpty": true,
	"RXQueuesEmpty": true,
	"TXQueuesFull": 0,
	"RXQueuesFull": 0,
	"Ethernet": {
	  "NumOffloadedIPv4Conns": 0,
	  "NumOffloadedIPv6Conns": 0
	},
	"@odata.id": "/redfish/v1/Chassis/1U/NetworkAdapters/Slot1/NetworkDeviceFunctions/SC2KP1F0/Metrics"
  }`

// TestNetworkDeviceFunctionMetrics tests the parsing of NetworkDeviceFunctionMetrics objects.
func TestNetworkDeviceFunctionMetrics(t *testing.T) {
	var result NetworkDeviceFunctionMetrics
	err := json.NewDecoder(strings.NewReader(networkDeviceFunctionMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "NetworkDeviceFunctionMetrics", result.ID)
	assertEquals(t, "Network Device Function Metrics", result.Name)

	if *result.TXAvgQueueDepthPercent != 13.7 {
		t.Errorf("Unexpected TXAvgQueueDepthPercent value: %.2f", *result.TXAvgQueueDepthPercent)
	}

	if *result.RXAvgQueueDepthPercent != 21.2 {
		t.Errorf("Unexpected RXAvgQueueDepthPercent value: %.2f", *result.RXAvgQueueDepthPercent)
	}

	if *result.RXFrames != 27193387 {
		t.Errorf("Unexpected RXFrames value: %d", result.RXFrames)
	}

	if *result.TXUnicastFrames != 17205770 {
		t.Errorf("Unexpected TXUnicastFrames value: %d", result.TXUnicastFrames)
	}

	if *result.RXMulticastFrames != 1000000 {
		t.Errorf("Unexpected RXMulticastFrames value: %d", result.RXMulticastFrames)
	}

	if *result.Ethernet.NumOffloadedIPv4Conns != 0 {
		t.Errorf("Unexpected NumOffloadedIPv4Conns value: %d", result.Ethernet.NumOffloadedIPv4Conns)
	}

	if *result.Ethernet.NumOffloadedIPv6Conns != 0 {
		t.Errorf("Unexpected NumOffloadedIPv4Conns value: %d", result.Ethernet.NumOffloadedIPv6Conns)
	}

	if !result.TXQueuesEmpty {
		t.Error("Expected TXQueues to be empty")
	}

	if !result.RXQueuesEmpty {
		t.Error("Expected RXQueues to be empty")
	}
}
