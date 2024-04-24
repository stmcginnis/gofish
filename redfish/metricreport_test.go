//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var metricReportBody = `{
	"@odata.type": "#MetricReport.v1_5_0.MetricReport",
	"Id": "AvgPlatformPowerUsage",
	"Name": "Average Platform Power Usage metric report",
	"MetricReportDefinition": {
	  "@odata.id": "/redfish/v1/TelemetryService/MetricReportDefinitions/AvgPlatformPowerUsage"
	},
	"MetricValues": [
	  {
		"MetricId": "AverageConsumedWatts",
		"MetricValue": "100",
		"Timestamp": "2016-11-08T12:25:00-05:00",
		"MetricProperty": "/redfish/v1/Chassis/Tray_1/Power#/0/PowerConsumedWatts"
	  },
	  {
		"MetricId": "AverageConsumedWatts",
		"MetricValue": "94",
		"Timestamp": "2016-11-08T13:25:00-05:00",
		"MetricProperty": "/redfish/v1/Chassis/Tray_1/Power#/0/PowerConsumedWatts"
	  },
	  {
		"MetricId": "AverageConsumedWatts",
		"MetricValue": "100",
		"Timestamp": "2016-11-08T14:25:00-05:00",
		"MetricProperty": "/redfish/v1/Chassis/Tray_1/Power#/0/PowerConsumedWatts"
	  }
	],
	"@odata.id": "/redfish/v1/TelemetryService/MetricReports/AvgPlatformPowerUsage"
  }`

// TestMetricReport tests the parsing of MetricReport objects.
func TestMetricReport(t *testing.T) {
	var result MetricReport
	err := json.NewDecoder(strings.NewReader(metricReportBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "AvgPlatformPowerUsage", result.ID)
	assertEquals(t, "Average Platform Power Usage metric report", result.Name)
	assertEquals(t, "/redfish/v1/TelemetryService/MetricReportDefinitions/AvgPlatformPowerUsage", result.metricReportDefinition)
	assertEquals(t, "AverageConsumedWatts", result.MetricValues[0].MetricID)
	assertEquals(t, "94", result.MetricValues[1].MetricValue)
	assertEquals(t, "/redfish/v1/Chassis/Tray_1/Power#/0/PowerConsumedWatts", result.MetricValues[2].MetricProperty)
}
