//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var radiusBody = `{
  "@odata.type": "#RADIUS.v1_0_2.RADIUS",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/RADIUS",
  "Id": "RADIUS",
  "Name": "RADIUS",
  "RadiusEnabled": true,
  "RadiusServer": "10.10.10.10",
  "RadiusPortNumber": 1812,
  "RadiusSecret": "SECRET",
  "@odata.etag": "\"23907a0ed1b646fed1e96d0fe734ddd9\""
}`

// TestRadius tests the parsing of Radius objects.
func TestRadius(t *testing.T) {
	var result RADIUS
	err := json.NewDecoder(strings.NewReader(radiusBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "RADIUS" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}

	if result.Server != "10.10.10.10" {
		t.Errorf("Invalid server: %s", result.Server)
	}

	if result.Port != 1812 {
		t.Errorf("Invalid port: %d", result.Port)
	}

	if result.Secret != "SECRET" {
		t.Errorf("Invalid secret: %s", result.Secret)
	}
}
