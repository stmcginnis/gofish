//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var spareResourceSetBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#SpareResourceSet.SpareResourceSet",
		"@odata.type": "#SpareResourceSet.v1_0_0.SpareResourceSet",
		"@odata.id": "/redfish/v1/SpareResourceSet",
		"Id": "SpareResourceSet-1",
		"Name": "SpareResourceSetOne",
		"Description": "SpareResourceSet One",
		"Links": {
			"OnHandSpares": [{
				"@odata.id": "/redfish/v1/SpareType/1"
			}],
			"OnHandSpares@odata.count": 1,
			"ReplacementSpareSets": [],
			"ReplacementSpareSets@odata.count": 0
		},
		"OnHandLocation": {
			"AltitudeMeters": 150,
			"Contacts": [{
					"ContactName": "Fred",
					"EmailAddress": "fred@example.com",
					"PhoneNumber": "+44 303 123 7300"
				},
				{
					"ContactName": "George",
					"EmailAddress": "george@example.com",
					"PhoneNumber": "+44 20 7405 7324"
				}
			],
			"Info": "Down the street",
			"InfoFormat": "Landmarks",
			"Latitude": 51.5074,
			"Longitude": 0.1278,
			"PartLocation": {
				"LocationOrdinalValue": 1,
				"LocationType": "Slot",
				"Orientation": "TopToBottom",
				"Reference": "Top",
				"ServiceLabel": "Label1"
			},
			"Placement": {
				"AdditionalInfo": "Turn left at the fountain.",
				"Rack": "Rack-2",
				"RackOffset": 3,
				"RackOffsetUnits": "OpenU",
				"Row": "Row-4"
			},
			"PostalAddress": {
				"Building": "The big one on the right",
				"City": "London",
				"Country": 826,
				"Floor": "Top floor",
				"HouseNumber": 89,
				"Name": "Big Data Center",
				"PlaceType": "office",
				"Street": "Clerkenwell Rd"
			}
		},
		"OnLine": true,
		"ResourceType": "Box",
		"TimeToProvision": "P0DT12H30M5S",
		"TimeToReplenish": "P5DT0H0M0S"
	}`)

// TestSpareResourceSet tests the parsing of SpareResourceSet objects.
func TestSpareResourceSet(t *testing.T) {
	var result SpareResourceSet
	err := json.NewDecoder(spareResourceSetBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SpareResourceSet-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SpareResourceSetOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.OnHandLocation.AltitudeMeters != 150 {
		t.Errorf("OnHandLocation Altitude incorrect: %d", result.OnHandLocation.AltitudeMeters)
	}

	if !result.OnLine {
		t.Error("OnLine should be true")
	}

	if result.ResourceType != "Box" {
		t.Errorf("Invalid resource type: %s", result.ResourceType)
	}
}
