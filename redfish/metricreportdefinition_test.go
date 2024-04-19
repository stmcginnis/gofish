//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var metricReportDefinitonBody = `{
	"@odata.type": "#MetricReportDefinition.v1_4_2.MetricReportDefinition",
	"@odata.context": "/redfish/v1/$metadata#MetricReportDefinition.MetricReportDefinition",
	"@odata.id": "/redfish/v1/TelemetryService/MetricReportDefinitions/CPUSensor",
	"Id": "CPUSensor",
	"Name": "CPU Sensor Metric Report",
	"Description": "CPU Sensor",
	"AppendLimit": 2400,
	"MetricReportDefinitionEnabled": false,
	"MetricReportDefinitionType": "Periodic",
	"MetricReportHeartbeatInterval": "PT0H0M0S",
	"SuppressRepeatedMetricValue": false,
	"ReportTimespan": "PT0H0M0S",
	"ReportUpdates": "Overwrite",
	"Wildcards": [],
	"MetricReportDefinitionType@Redfish.AllowableValues": [
	  "Periodic",
	  "OnChange",
	  "OnRequest"
	],
	"ReportUpdates@Redfish.AllowableValues": [
	  "AppendStopsWhenFull",
	  "AppendWrapsWhenFull",
	  "NewReport",
	  "Overwrite"
	],
	"ReportActions": [
	  "RedfishEvent",
	  "LogToMetricReportsCollection"
	],
	"ReportActions@Redfish.AllowableValues": [
	  "LogToMetricReportsCollection",
	  "RedfishEvent"
	],
	"Status": {
	  "State": "Disabled"
	},
	"Schedule": {
	  "RecurrenceInterval": "PT0H1M0S"
	},
	"Metrics": [
	  {
		"MetricId": "TemperatureReading",
		"MetricProperties": [],
		"CollectionFunction": null,
		"CollectionDuration": null,
		"CollectionTimeScope": "Point",
		"Oem": {
		  "Dell": {
			"@odata.type": "#DellMetric.v1_1_0.DellMetric",
			"CustomLabel": null,
			"FQDD": "iDRAC.Embedded.1#CPU%Temp",
			"Source": null
		  }
		}
	  }
	],
	"Links": {
	  "Triggers": [
		{
		  "@odata.id": "/redfish/v1/TelemetryService/Triggers/CPUCriticalTrigger"
		},
		{
		  "@odata.id": "/redfish/v1/TelemetryService/Triggers/CPUWarnTrigger"
		},
		{
		  "@odata.id": "/redfish/v1/TelemetryService/Triggers/TMPCpuCriticalTrigger"
		},
		{
		  "@odata.id": "/redfish/v1/TelemetryService/Triggers/TMPCpuWarnTrigger"
		}
	  ]
	},
	"Oem": {
	  "Dell": {
		"@odata.type": "#DellMetricReportDefinition.v1_1_0.DellMetricReportDefinition",
		"Digest": "5230bb90347b1362c6a2b6c69a26ada70369b5c8c5e4379f4d932bb639b6a630",
		"iDRACFirmwareVersion": "7.00.60.00"
	  }
	}
  }`

// TestMetricReportDefinition tests the parsing of MetricReportDefinition objects.
func TestMetricReportDefinition(t *testing.T) {
	var result MetricReportDefinition
	err := json.NewDecoder(strings.NewReader(metricReportDefinitonBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "CPUSensor", result.ID)
	assertEquals(t, "CPU Sensor Metric Report", result.Name)
	assertEquals(t, "CPU Sensor", result.Description)
	assertEquals(t, "2400", fmt.Sprintf("%d", result.AppendLimit))
	assertEquals(t, "false", fmt.Sprintf("%t", result.MetricReportDefinitionEnabled))
	assertEquals(t, "PT0H0M0S", result.ReportTimespan)
	assertEquals(t, "2", fmt.Sprintf("%d", len(result.ReportActions)))
	assertEquals(t, "1", fmt.Sprintf("%d", len(result.Metrics)))
	assertEquals(t, "TemperatureReading", result.Metrics[0].MetricID)
	assertEquals(t, "4", fmt.Sprintf("%d", len(result.triggers)))
	assertEquals(t, "/redfish/v1/TelemetryService/Triggers/CPUCriticalTrigger", result.triggers[0])
}
