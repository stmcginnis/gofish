//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var nodeManagerBody = `{
  "@odata.type": "#SmcNodeManager.v1_0_1.SmcNodeManager",
  "@odata.id": "/redfish/v1/Systems/1/Oem/Supermicro/NodeManager",
  "Id": "Node Manager",
  "Name": "Node Manager",
  "Capabilities": [
    {
      "DomainID": "Entire platform",
      "PolicyType": "Power Control Policy",
      "MaxConcurrentSettings": 16,
      "MaxValueAfterReset": 32767,
      "MinValueAfterReset": 1,
      "MaxCorrectionTime": 600000,
      "MinCorrectionTime": 3000,
      "MaxReportingPeriod": 3600,
      "MinReportingPeriod": 1,
      "DomainLimitingScope": 128
    }
  ],
  "Statistics": [
    {
      "Mode": "Global power statistics",
      "DomainID": "Entire platform",
      "Timestamp": "2024-10-14T16:58:30:+00:00",
      "CurrentValue": 8934,
      "MaximumValue": 9587,
      "MinimumValue": 88,
      "AverageValue": 4692,
      "ReportingPeriod": 94032
    }
  ],
  "IntelPsysEnabled": false,
  "IntelPsysSupported": false,
  "Version": {
    "IntelNMVersion": "Supported Intel NM 6.0",
    "IPMIVersion": "Intel NM IPMI version 6.0",
    "PatchVersion": 0,
    "MajorRevision": 6,
    "MinorRevision": 20
  },
  "Selftest": {
    "MajorCode": 85,
    "MinorCode": 0,
    "ImageFlags": "Operational"
  },
  "Policy": [
    {
      "DomainID": "Entire platform",
      "PolicyID": 1,
      "PolicyType": 0,
      "PolicyExceptionActions": 0,
      "PowerLimit": 0,
      "CorrectionTimeLimit": 0,
      "PolicyTriggerLimit": 0,
      "StatReportingPeriod": 0
    }
  ],
  "Actions": {
    "#SmcNodeManager.ClearAllPolicies": {
      "target": "/redfish/v1/Systems/1/Oem/Supermicro/NodeManager/Actions/SmcNodeManager.ClearAllPolicies"
    }
  },
  "@odata.etag": "\"ec6bda0f0947b85ca44e4a068acd2e66\""
}`

// TestNodeManager tests the parsing of NodeManager objects.
func TestNodeManager(t *testing.T) {
	var result NodeManager
	err := json.NewDecoder(strings.NewReader(nodeManagerBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Node Manager" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.clearAllPoliciesTarget != "/redfish/v1/Systems/1/Oem/Supermicro/NodeManager/Actions/SmcNodeManager.ClearAllPolicies" {
		t.Errorf("Invalid clear all policies link: %s", result.clearAllPoliciesTarget)
	}

	if len(result.Statistics) != 1 {
		t.Errorf("Expected 1 statistic, got %d", len(result.Statistics))
	}

	if result.Statistics[0].Mode != "Global power statistics" {
		t.Errorf("Expected 'Global power statistics', got %s", result.Statistics[0].Mode)
	}
}
