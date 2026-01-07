//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var portMetricsBody = `{
	"@odata.type": "#PortMetrics.v1_5_1.PortMetrics",
	"Id": "Metrics",
	"Name": "Gen-Z Port 1 Metrics",
	"GenZ": {
	  "PacketCRCErrors": 24,
	  "EndToEndCRCErrors": 3,
	  "RXStompedECRC": 1,
	  "TXStompedECRC": 2,
	  "NonCRCTransientErrors": 2,
	  "LLRRecovery": 1,
	  "MarkedECN": 1,
	  "PacketDeadlineDiscards": 1,
	  "AccessKeyViolations": 1,
	  "LinkNTE": 1,
	  "ReceivedECN": 1
	},
	"@odata.id": "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/Metrics"
  }`

// TestPortMetrics tests the parsing of PortMetrics objects.
func TestPortMetrics(t *testing.T) {
	var result PortMetrics
	err := json.NewDecoder(strings.NewReader(portMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Gen-Z Port 1 Metrics", result.Name)

	if *result.GenZ.PacketCRCErrors != 24 {
		t.Errorf("Unexpected GenZ.PacketCRCErrors value: %d", result.GenZ.PacketCRCErrors)
	}

	if *result.GenZ.EndToEndCRCErrors != 3 {
		t.Errorf("Unexpected GenZ.EndToEndCRCErrors value: %d", result.GenZ.EndToEndCRCErrors)
	}

	if *result.GenZ.RXStompedECRC != 1 {
		t.Errorf("Unexpected GenZ.RXStompedECRC value: %d", result.GenZ.RXStompedECRC)
	}
}
