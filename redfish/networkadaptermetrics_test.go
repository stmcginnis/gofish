//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkAdapterMetricsBody = `{
	"@odata.type": "#NetworkAdapterMetrics.v1_0_1.NetworkAdapterMetrics",
	"Id": "NetworkAdapterMetrics",
	"Name": "Network Adapter Metrics",
	"HostBusRXPercent": 35.53,
	"HostBusTXPercent": 14.17,
	"CPUCorePercent": 8.35,
	"NCSIRXFrames": 0,
	"NCSITXFrames": 0,
	"NCSIRXBytes": 0,
	"NCSITXBytes": 0,
	"RXBytes": 7754199970,
	"RXMulticastFrames": 1941,
	"RXUnicastFrames": 27193387,
	"TXBytes": 9436506547,
	"TXMulticastFrames": 153,
	"TXUnicastFrames": 18205770,
	"@odata.id": "/redfish/v1/Chassis/1U/NetworkAdapters/Slot1/Metrics"
  }`

// TestNetworkAdapterMetrics tests the parsing of NetworkAdapterMetrics objects.
func TestNetworkAdapterMetrics(t *testing.T) {
	var result NetworkAdapterMetrics
	err := json.NewDecoder(strings.NewReader(networkAdapterMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "NetworkAdapterMetrics", result.ID)
	assertEquals(t, "Network Adapter Metrics", result.Name)

	if result.HostBusRXPercent != 35.53 {
		t.Errorf("Unexpected HostBusRXPercent value: %.2f", result.HostBusRXPercent)
	}

	if result.HostBusTXPercent != 14.17 {
		t.Errorf("Unexpected HostBusTXPercent value: %.2f", result.HostBusTXPercent)
	}

	if result.CPUCorePercent != 8.35 {
		t.Errorf("Unexpected CPUCorePercent value: %.2f", result.CPUCorePercent)
	}

	if result.NCSIRXFrames != 0 {
		t.Errorf("Unexpected NCSIRXFrames value: %d", result.NCSIRXFrames)
	}

	if result.RXBytes != 7754199970 {
		t.Errorf("Unexpected RXBytes value: %d", result.RXBytes)
	}

	if result.RXMulticastFrames != 1941 {
		t.Errorf("Unexpected RXMulticastFrames value: %d", result.RXMulticastFrames)
	}

	if result.TXUnicastFrames != 18205770 {
		t.Errorf("Unexpected TXUnicastFrames value: %d", result.TXUnicastFrames)
	}
}
