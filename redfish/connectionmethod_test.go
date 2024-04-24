//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var connectionMethodBody = `{
	"@odata.type": "#ConnectionMethod.v1_1_0.ConnectionMethod",
	"Id": "ConnectionMethod1",
	"Name": "ConnectionMethod One",
	"ConnectionMethodType": "Redfish",
	"ConnectionMethodVariant": "Contoso",
	"Links": {
		"AggregationSources": [
			{
				"@odata.id": "/redfish/v1/AggregationService/AggregationSources/AggregationSource1"
			}
		]
	},
	"@odata.id": "/redfish/v1/AggregationService/ConnectionMethods/ConnectionMethod1"
	}`

// TestConnectionMethod tests the parsing of ConnectionMethod objects.
func TestConnectionMethod(t *testing.T) {
	var result ConnectionMethod
	err := json.NewDecoder(strings.NewReader(connectionMethodBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ConnectionMethod1", result.ID)
	assertEquals(t, "ConnectionMethod One", result.Name)
	assertEquals(t, "Redfish", string(result.ConnectionMethodType))
	assertEquals(t, "Contoso", result.ConnectionMethodVariant)
	assertEquals(t, "/redfish/v1/AggregationService/AggregationSources/AggregationSource1", result.aggregationSources[0])
}
