//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/school/common"
)

var compositionServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#CompositionService.CompositionService",
		"@odata.type": "#CompositionService.v1_0_0.CompositionService",
		"@odata.id": "/redfish/v1/CompositionService",
		"Id": "CompositionService-1",
		"Name": "Composition Service",
		"Description": "Composition Service",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"AllowOverprovisioning": true,
		"AllowZoneAffinity": false,
		"ServiceEnabled": true,
		"ResourceBlocks": {
			"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks"
		},
		"ResourceZones": {
			"@odata.id": "/redfish/v1/CompositionService/ResourceZones"
		}
	}`)

// TestCompositionService tests the parsing of CompositionService objects.
func TestCompositionService(t *testing.T) {
	var result CompositionService
	err := json.NewDecoder(compositionServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CompositionService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Composition Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.AllowOverprovisioning {
		t.Error("Expected AllowOverprovisioning to be true")
	}

	if result.AllowZoneAffinity {
		t.Error("Expected AllowZoneAffinity to be false")
	}

	if !result.ServiceEnabled {
		t.Error("Expected ServiceEnabled to be true")
	}

	if result.Status.Health != common.OKHealth {
		t.Errorf("Received invalid health status: %s", result.Status.Health)
	}

	if result.resourceBlocks != "/redfish/v1/CompositionService/ResourceBlocks" {
		t.Errorf("Received invalid resource blocks reference: %s", result.resourceBlocks)
	}

	if result.resourceZones != "/redfish/v1/CompositionService/ResourceZones" {
		t.Errorf("Received invalid resource zones reference: %s", result.resourceZones)
	}
}
