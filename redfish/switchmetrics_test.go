//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var switchMetricsBody = `{
	"@odata.type": "#SwitchMetrics.v1_0_1.SwitchMetrics",
	"Id": "SwitchMetrics",
	"Name": "PCIe Switch Metrics",
	"PCIeErrors": {
	  "CorrectableErrorCount": 1,
	  "NonFatalErrorCount": 0,
	  "FatalErrorCount": 0,
	  "L0ToRecoveryCount": 0,
	  "ReplayCount": 0,
	  "ReplayRolloverCount": 0,
	  "NAKSentCount": 0,
	  "NAKReceivedCount": 0
	},
	"InternalMemoryMetrics": {
	  "CurrentPeriod": {
		"CorrectableECCErrorCount": 1,
		"UncorrectableECCErrorCount": 0
	  },
	  "LifeTime": {
		"CorrectableECCErrorCount": 0,
		"UncorrectableECCErrorCount": 1
	  }
	},
	"@odata.id": "/redfish/v1/Fabrics/PCIe/Switches/1/SwitchMetrics"
  }`

// TestSwitchMetrics tests the parsing of SwitchMetrics objects.
func TestSwitchMetrics(t *testing.T) {
	var result SwitchMetrics
	err := json.NewDecoder(strings.NewReader(switchMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "SwitchMetrics", result.ID)
	assertEquals(t, "PCIe Switch Metrics", result.Name)

	if result.PCIeErrors.CorrectableErrorCount != 1 {
		t.Errorf("Unexpected PCIeErrors.CorrectableErrorCount: %d", result.PCIeErrors.CorrectableErrorCount)
	}

	if result.InternalMemoryMetrics.CurrentPeriod.CorrectableECCErrorCount != 1 {
		t.Errorf("Unexpected InternalMemoryMetrics.CurrentPeriod.CorrectableECCErrorCount: %d", result.InternalMemoryMetrics.CurrentPeriod.CorrectableECCErrorCount)
	}

	if result.InternalMemoryMetrics.LifeTime.UncorrectableECCErrorCount != 1 {
		t.Errorf("Unexpected InternalMemoryMetrics.LifeTime.UncorrectableECCErrorCount: %d", result.InternalMemoryMetrics.LifeTime.UncorrectableECCErrorCount)
	}
}
