//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var zoneBody = `{
	"@Redfish.CollectionCapabilities": {
	  "@odata.type": "#CollectionCapabilities.v1_2_0.CollectionCapabilities",
	  "Capabilities": [
		{
		  "CapabilitiesObject": {
			"@odata.id": "/redfish/v1/Systems/Capabilities"
		  },
		  "Links": {
			"TargetCollection": {
			  "@odata.id": "/redfish/v1/Systems"
			}
		  },
		  "UseCase": "ComputerSystemComposition"
		}
	  ],
	  "MaxMembers": 1
	},
	"@odata.context": "/redfish/v1/$metadata#Zone.Zone",
	"@odata.etag": "\"1712866586\"",
	"@odata.id": "/redfish/v1/CompositionService/ResourceZones/1",
	"@odata.type": "#Zone.v1_3_1.Zone",
	"Description": "Resource Zone 1",
	"Id": "1",
	"Links": {
	  "ResourceBlocks": [
		{
		  "@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/ComputeBlock"
		},
		{
		  "@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/DrivesBlock"
		},
		{
		  "@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/NetworkBlock"
		}
	  ]
	},
	"Name": "Resource Zone 1",
	"Status": {
	  "Health": "OK",
	  "HealthRollup": "OK",
	  "State": "Enabled"
	}
  }`

// TestZone tests the parsing of Zone objects.
func TestZone(t *testing.T) {
	var result Zone
	err := json.NewDecoder(strings.NewReader(zoneBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Resource Zone 1", result.Name)
	assertEquals(t, "Resource Zone 1", result.Description)

	if len(result.resourceBlocks) != 3 {
		t.Errorf("Expected 3 resource blocks, got %#v", result.resourceBlocks)
	}
}
