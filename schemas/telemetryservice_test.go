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

var tsBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#TelemetryService.TelemetryService",
		"@odata.id": "/redfish/v1/TelemetryService",
		"@odata.type": "#TelemetryService.v1_3_1.TelemetryService",
		"Actions": {
		  "#TelemetryService.SubmitTestMetricReport": {
			"target": "/redfish/v1/TelemetryService/Actions/TelemetryService.SubmitTestMetricReport"
		  },
		  "#TelemetryService.ClearMetricReports": {
			"target": "/redfish/v1/TelemetryService/Actions/TelemetryService.ClearMetricReports"
		  },
		  "#TelemetryService.ResetMetricReportDefinitionsToDefaults": {
			"target": "/redfish/v1/TelemetryService/Actions/TelemetryService.ResetMetricReportDefinitionsToDefaults"
		  },
		  "Oem": {
			"#DellTelemetryService.SubmitMetricValue": {
			  "target": "/redfish/v1/TelemetryService/Actions/Oem/DellTelemetryService.SubmitMetricValue"
			}
		  }
		},
		"Id": "TelemetryService",
		"Name": "TelemetryService",
		"Description": "The Manager TelemetryService",
		"MaxReports": 50,
		"MinCollectionInterval": "PT0H0M5S",
		"ServiceEnabled": false,
		"SupportedCollectionFunctions": [
		  "Average",
		  "Maximum",
		  "Minimum",
		  "Summation"
		],
		"MetricDefinitions": {
		  "@odata.id": "/redfish/v1/TelemetryService/MetricDefinitions"
		},
		"MetricReportDefinitions": {
		  "@odata.id": "/redfish/v1/TelemetryService/MetricReportDefinitions"
		},
		"MetricReports": {
		  "@odata.id": "/redfish/v1/TelemetryService/MetricReports"
		},
		"Triggers": {
		  "@odata.id": "/redfish/v1/TelemetryService/Triggers"
		}
	}`)

var smcBody = strings.NewReader(
	`{
			"@odata.type": "#TelemetryService.v1_2_0.TelemetryService",
			"@odata.id": "/redfish/v1/TelemetryService",
			"Id": "TelemetryService",
			"Name": "Telemetry Service",
			"Status": {
			  "State": "Enabled",
			  "Health": "OK"
			},
			"SupportedCollectionFuntions": [
			  "Average",
			  "Minimum",
			  "Maximum"
			],
			"MetricDefinitions": {
			  "@odata.id": "/redfish/v1/TelemetryService/MetricDefinitions"
			},
			"MetricReportDefinitions": {
			  "@odata.id": "/redfish/v1/TelemetryService/MetricReportDefinitions"
			},
			"MetricReports": {
			  "@odata.id": "/redfish/v1/TelemetryService/MetricReports"
			}
	}`)

// TestTelemetryService tests the parsing of TelemetryService objects.
func TestTelemetryService(t *testing.T) {
	var result TelemetryService
	err := json.NewDecoder(tsBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "TelemetryService", result.ID)
	assertEquals(t, "TelemetryService", result.Name)
	assertEquals(t, "The Manager TelemetryService", result.Description)
	assertEquals(t, "50", fmt.Sprintf("%d", *result.MaxReports))
	assertEquals(t, "/redfish/v1/TelemetryService/MetricDefinitions", result.metricDefinitions)
	assertEquals(t, "/redfish/v1/TelemetryService/MetricReportDefinitions", result.metricReportDefinitions)
	assertEquals(t, "/redfish/v1/TelemetryService/MetricReports", result.metricReports)
	assertEquals(t, "/redfish/v1/TelemetryService/Triggers", result.triggers)

	if len(result.SupportedCollectionFunctions) != 4 {
		t.Errorf("Expected 4 supported collection functions, got: %d", len(result.SupportedCollectionFunctions))
	}
	assertEquals(t, "Average", string(result.SupportedCollectionFunctions[0]))
}

// TestTelemetryServiceSMC tests the workaround for a Supermicro typo bug in the TelemetryService.
func TestTelemetryServiceSMC(t *testing.T) {
	var result TelemetryService
	err := json.NewDecoder(smcBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if len(result.SupportedCollectionFunctions) != 3 {
		t.Errorf("Expected 3 supported collection functions, got: %d", len(result.SupportedCollectionFunctions))
	}

	if result.SupportedCollectionFunctions[0] != "Average" {
		t.Errorf("Unexpected collection function content: %v", result.SupportedCollectionFunctions)
	}
}
