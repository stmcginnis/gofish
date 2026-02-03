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

var triggersBody = strings.NewReader(
	`{
		"@odata.type": "#Triggers.v1_2_0.Triggers",
		"@odata.context": "/redfish/v1/$metadata#Triggers.Triggers",
		"@odata.id": "/redfish/v1/TelemetryService/Triggers/CPUCriticalTrigger",
		"Id": "CPUCriticalTrigger",
		"Name": "Trigger on CPU critical errors",
		"Description": "Trigger when an OEM event is raised",
		"TriggerActions": [
		  "RedfishMetricReport"
		],
		"TriggerActions@Redfish.AllowableValues": [
		  "RedfishMetricReport"
		],
		"EventTriggers": [
		  "iDRAC.1.6.CPU0004",
		  "iDRAC.1.6.CPU0700",
		  "iDRAC.1.6.CPU0702",
		  "iDRAC.1.6.CPU0006",
		  "iDRAC.1.6.CPU0703",
		  "iDRAC.1.6.CPU0701",
		  "iDRAC.1.6.CPU0003"
		],
		"Links": {
		  "MetricReportDefinitions": [
			{
			  "@odata.id": "/redfish/v1/TelemetryService/MetricReportDefinitions/CPUSensor"
			}
		  ]
		}
	  }`)

// TestTriggers tests the parsing of Triggers objects.
func TestTriggers(t *testing.T) {
	var result Triggers
	err := json.NewDecoder(triggersBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "CPUCriticalTrigger", result.ID)
	assertEquals(t, "Trigger on CPU critical errors", result.Name)
	assertEquals(t, "Trigger when an OEM event is raised", result.Description)
	assertEquals(t, "1", fmt.Sprintf("%d", len(result.TriggerActions)))
	assertEquals(t, "7", fmt.Sprintf("%d", len(result.EventTriggers)))
	assertEquals(t, "1", fmt.Sprintf("%d", len(result.metricReportDefinitions)))
	assertEquals(t, "/redfish/v1/TelemetryService/MetricReportDefinitions/CPUSensor", result.metricReportDefinitions[0])
}
