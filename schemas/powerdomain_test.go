//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var powerDomainBody = `{
	"@odata.type": "#PowerDomain.v1_2_1.PowerDomain",
	"Id": "Row1",
	"Name": "Row #1 Domain",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Links": {
	  "ManagedBy": [
		{
		  "@odata.id": "/redfish/v1/Managers/BMC"
		}
	  ],
	  "RackPDUs": [
		{
		  "@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Facilities/Room237/PowerDomains/Row1"
  }`

// TestPowerDomain tests the parsing of PowerDomain objects.
func TestPowerDomain(t *testing.T) {
	var result PowerDomain
	err := json.NewDecoder(strings.NewReader(powerDomainBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Row1", result.ID)
	assertEquals(t, "Row #1 Domain", result.Name)
	assertEquals(t, "/redfish/v1/Managers/BMC", result.managedBy[0])
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1", result.rackPDUs[0])
}
