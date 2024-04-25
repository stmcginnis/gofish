//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var serviceConditionsBody = `{
	"@odata.type": "#ServiceConditions.v1_0_0.ServiceConditions",
	"ID": "ServiceConditions",
	"Name": "Redfish Service Conditions",
	"HealthRollup": "Warning",
	"Conditions": [
	  {
		"MessageId": "ThermalEvents.1.0.OverTemperature",
		"Timestamp": "2020-11-08T12:25:00-05:00 ",
		"Message": "Temperature exceeds rated limit in power supply 'A'.",
		"Severity": "Warning",
		"MessageArgs": [
		  "A"
		],
		"OriginOfCondition": {
		  "@odata.id": "/redfish/v1/Chassis/1/Power"
		},
		"LogEntry": {
		  "@odata.id": "/redfish/v1/Managers/1/LogServices/Log1/Entries/1"
		}
	  },
	  {
		"MessageId": "Base.1.9.ConditionInRelatedResource",
		"Message": "One or more conditions exist in a related resource. See theOriginOfCondition property.",
		"Severity": "Warning",
		"OriginOfCondition": {
		  "@odata.id": "/redfish/v1/Systems/cpu-memory-example"
		}
	  }
	],
	"@odata.id": "/redfish/v1/ServiceConditions"
  }`

// TestServiceConditions tests the parsing of ServiceConditions objects.
func TestServiceConditions(t *testing.T) {
	var result ServiceConditions
	err := json.NewDecoder(strings.NewReader(serviceConditionsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ServiceConditions", result.ID)
	assertEquals(t, "Redfish Service Conditions", result.Name)
	assertEquals(t, "Warning", string(result.HealthRollup))
	assertEquals(t, "Temperature exceeds rated limit in power supply 'A'.", result.Conditions[0].Message)
	assertEquals(t, "/redfish/v1/Chassis/1/Power", result.Conditions[0].originOfCondition)
	assertEquals(t, "/redfish/v1/Managers/1/LogServices/Log1/Entries/1", result.Conditions[0].logEntry)
	assertEquals(t, "Base.1.9.ConditionInRelatedResource", result.Conditions[1].MessageID)
	assertEquals(t, "Warning", string(result.Conditions[1].Severity))
}
