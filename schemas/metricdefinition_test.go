//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var metricDefinitonBody = `{
	"@odata.type": "#MetricDefinition.v1_3_1.MetricDefinition",
	"@odata.context": "/redfish/v1/$metadata#MetricDefinition.MetricDefinition",
	"@odata.id": "/redfish/v1/TelemetryService/MetricDefinitions/VoltageStatus",
	"Id": "VoltageStatus",
	"Name": "Voltage status Metric Definition",
	"Description": "Status of the Voltage of Small Form Factor pluggable(SFP) Transceiver",
	"MetricType": "Discrete",
	"MetricDataType": "Enumeration",
	"Accuracy": 0,
	"SensingInterval": "PT60.0S",
	"DiscreteValues": [
	  "Unknown",
	  "OK",
	  "Warning",
	  "Critical"
	]
  }`

// TestMetricDefinition tests the parsing of MetricDefinition objects.
func TestMetricDefinition(t *testing.T) {
	var result MetricDefinition
	err := json.NewDecoder(strings.NewReader(metricDefinitonBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "VoltageStatus", result.ID)
	assertEquals(t, "Voltage status Metric Definition", result.Name)
	assertEquals(t, "Status of the Voltage of Small Form Factor pluggable(SFP) Transceiver", result.Description)
	assertEquals(t, "0", fmt.Sprintf("%.0f", *result.Accuracy))
	assertEquals(t, "PT60.0S", result.SensingInterval)
	assertEquals(t, "4", fmt.Sprintf("%d", len(result.DiscreteValues)))
	assertEquals(t, "Discrete", string(result.MetricType))
	assertEquals(t, "Enumeration", string(result.MetricDataType))
}
