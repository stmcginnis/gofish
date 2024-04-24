//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var filterBody = `{
	"@odata.type": "#Filter.v1_0_1.Filter",
	"Id": "1",
	"Name": "Cooling Loop Filter",
	"ServicedDate": "2020-12-24T08:00:00Z",
	"ServiceHours": 5791,
	"RatedServiceHours": 10000,
	"Manufacturer": "Contoso",
	"Model": "MyCoffee",
	"PartNumber": "Cone4",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Location": {
	  "Placement": {
		"Row": "North 1"
	  }
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/Filters/1"
  }`

// TestFilter tests the parsing of Filter objects.
func TestFilter(t *testing.T) {
	var result Filter
	err := json.NewDecoder(strings.NewReader(filterBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Cooling Loop Filter", result.Name)
	assertEquals(t, "North 1", result.Location.Placement.Row)

	if result.RatedServiceHours != 10000 {
		t.Errorf("Expected rated service hours: %.2f", result.RatedServiceHours)
	}

	if result.ServiceHours != 5791 {
		t.Errorf("Expected rated service hours: %.2f", result.ServiceHours)
	}
}
