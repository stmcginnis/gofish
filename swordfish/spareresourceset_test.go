//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var spareResourceSetBody = `{
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
				"Country": "UK",
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
	}`

// TestSpareResourceSet tests the parsing of SpareResourceSet objects.
func TestSpareResourceSet(t *testing.T) {
	var result SpareResourceSet
	err := json.NewDecoder(strings.NewReader(spareResourceSetBody)).Decode(&result)

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

// TestSpareResourceSetUpdate tests the Update call.
func TestSpareResourceSetUpdate(t *testing.T) {
	var result SpareResourceSet
	err := json.NewDecoder(strings.NewReader(spareResourceSetBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.OnLine = true
	result.ResourceType = "Hat"
	result.TimeToProvision = "P0DT06H30M5S"
	result.TimeToReplenish = "P5DT0H12M0S"
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if strings.Contains(calls[0].Payload, "OnLine") {
		t.Errorf("Unexpected OnLine update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "ResourceType:Hat") {
		t.Errorf("Unexpected ResourceType update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "TimeToProvision:P0DT06H30M5S") {
		t.Errorf("Unexpected TimeToProvision update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "TimeToReplenish:P5DT0H12M0S") {
		t.Errorf("Unexpected TimeToReplenish update payload: %s", calls[0].Payload)
	}
}
