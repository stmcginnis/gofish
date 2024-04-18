//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var tsBody = strings.NewReader(
	`{
		"@odata.type": "#TelemetryService.v1_2_0.TelemetryService",
		"@odata.id": "/redfish/v1/TelemetryService",
		"Id": "TelemetryService",
		"Name": "Telemetry Service",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"SupportedCollectionFunctions": [
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

	if result.ID != "TelemetryService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Telemetry Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.SupportedCollectionFunctions) != 3 {
		t.Errorf("Expected 3 supported collection functions, got: %d", len(result.SupportedCollectionFunctions))
	}

	if result.SupportedCollectionFunctions[0] != "Average" {
		t.Errorf("Unexpected collection function content: %v", result.SupportedCollectionFunctions)
	}

	if TaskState(result.metricDefinitions) != "/redfish/v1/TelemetryService/MetricDefinitions" {
		t.Errorf("Invalid metric definition target: %s", result.metricDefinitions)
	}

	if TaskState(result.metricReportDefinitions) != "/redfish/v1/TelemetryService/MetricReportDefinitions" {
		t.Errorf("Invalid metric report definition target: %s", result.metricReportDefinitions)
	}

	if TaskState(result.metricReports) != "/redfish/v1/TelemetryService/MetricReports" {
		t.Errorf("Invalid metric definition target: %s", result.metricReports)
	}
}

// TestTelemetryServiceSMC tests the workaround for a Supermicro typo bug in the  TelemetryService.
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
