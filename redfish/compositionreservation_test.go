//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var compositionReservationBody = `{
	"@odata.type": "#CompositionReservation.v1_0_1.CompositionReservation",
	"Id": "1",
	"Name": "Composition Reservation 1",
	"ReservationTime": "2019-08-22T10:35:16+06:00",
	"Client": "VCF1001",
	"ReservedResourceBlocks": [
	  {
		"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/BladeServer-2"
	  }
	],
	"Manifest": {
	  "Description": "Description for this Manifest document.",
	  "Timestamp": "2019-08-22T10:35:16+06:00",
	  "Expand": "None",
	  "Stanzas": [
		{
		  "StanzaType": "ComposeSystem",
		  "StanzaId": "Compute1",
		  "Request": {
			"Links": {
			  "ResourceBlocks": [
				{
				  "@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/BladeServer-2"
				}
			  ]
			}
		  }
		}
	  ]
	},
	"@odata.id": "/redfish/v1/CompositionService/CompositionReservations/1"
  }`

// TestCompositionReservation tests the parsing of CompositionReservation objects.
func TestCompositionReservation(t *testing.T) {
	var result CompositionReservation
	err := json.NewDecoder(strings.NewReader(compositionReservationBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Composition Reservation 1", result.Name)
	assertEquals(t, "VCF1001", result.Client)
	assertEquals(t, "/redfish/v1/CompositionService/ResourceBlocks/BladeServer-2", result.reservedResourceBlocks[0])
	assertEquals(t, "Description for this Manifest document.", result.Manifest.Description)
	assertEquals(t, "None", string(result.Manifest.Expand))
	assertEquals(t, "ComposeSystem", string(result.Manifest.Stanzas[0].StanzaType))
}
