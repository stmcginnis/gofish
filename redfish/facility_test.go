//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var facilityBody = `{
	"@odata.type": "#Facility.v1_4_1.Facility",
	"Id": "Room237",
	"Name": "Room #237, 2nd Floor",
	"FacilityType": "Room",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Location": {
	  "PostalAddress": {
		"Country": "US",
		"Territory": "OR",
		"City": "Portland",
		"Street": "1001 SW 5th Avenue",
		"HouseNumber": 1100,
		"Name": "DMTF, Inc.",
		"PostalCode": "97204",
		"Floor": "2",
		"Room": "237"
	  }
	},
	"PowerDomains": {
	  "@odata.id": "/redfish/v1/Facilities/Room237/PowerDomains"
	},
	"Links": {
	  "ContainedByFacility": {
		"@odata.id": "/redfish/v1/Facilities/Building"
	  },
	  "RackPDUs": [
		{
		  "@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Facilities/Room237"
  }`

// TestFacility tests the parsing of Facility objects.
func TestFacility(t *testing.T) {
	var result Facility
	err := json.NewDecoder(strings.NewReader(facilityBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Room237", result.ID)
	assertEquals(t, "Room #237, 2nd Floor", result.Name)
	assertEquals(t, "Room", string(result.FacilityType))
	assertEquals(t, "Portland", result.Location.PostalAddress.City)
	assertEquals(t, "/redfish/v1/Facilities/Room237/PowerDomains", result.powerDomains)
	assertEquals(t, "/redfish/v1/Facilities/Building", result.containedByFacility)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1", result.rackPDUs[0])
}
