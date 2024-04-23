//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var aggregationSourceBody = `{
		"@odata.type": "#AggregationSource.v1_4_0.AggregationSource",
		"Id": "AggregationSource1",
		"Name": "AggregationSource One",
		"HostName": "https://Someserver.Contoso.com/redfish/v1",
		"UserName": "root",
		"Password": null,
		"Links": {
			"ConnectionMethod": {
				"@odata.id": "/redfish/v1/AggregationService/ConnectionMethods/ConnectionMethod1"
			},
			"ResourcesAccessed": [
				{
				"@odata.id": "/redfish/v1/Managers/1"
				}
			]
		},
		"@odata.id": "/redfish/v1/AggregationService/AggregationSources/AggregationSource1"
	}`

// TestAggregationSource tests the parsing of AggregationSource objects.
func TestAggregationSource(t *testing.T) {
	var result AggregationSource
	err := json.NewDecoder(strings.NewReader(aggregationSourceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "AggregationSource1", result.ID)
	assertEquals(t, "AggregationSource One", result.Name)
	assertEquals(t, "https://Someserver.Contoso.com/redfish/v1", result.HostName)
	assertEquals(t, "root", result.UserName)
	assertEquals(t, "", result.Password)

	if result.connectionMethod != "/redfish/v1/AggregationService/ConnectionMethods/ConnectionMethod1" {
		t.Errorf("Unexpected connection method: %s", result.connectionMethod)
	}

	if len(result.resourcesAccessed) != 1 {
		t.Errorf("Expected 1 ResourcesAccessed, got: %#v", result.resourcesAccessed)
	}
}
