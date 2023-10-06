//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var powerEquipmentBody = strings.NewReader(
	`{
		"@odata.type": "#PowerEquipment.v1_2_0.PowerEquipment",
		"Id": "PowerEquipment",
		"Name": "DCIM Power Equipment",
		"Status": {
			"State": "Enabled",
			"HealthRollup": "OK"
		},
		"FloorPDUs": {
			"@odata.id": "/redfish/v1/PowerEquipment/FloorPDUs"
		},
		"RackPDUs": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs"
		},
		"TransferSwitches": {
			"@odata.id": "/redfish/v1/PowerEquipment/TransferSwitches"
		},
		"@odata.id": "/redfish/v1/PowerEquipment"
	}`)

// TestPowerEquipment tests the parsing of PowerEquipment objects.
func TestPowerEquipment(t *testing.T) {
	var result PowerEquipment
	err := json.NewDecoder(powerEquipmentBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PowerEquipment", result.ID)
	assertEquals(t, "DCIM Power Equipment", result.Name)
	assertEquals(t, "Enabled", fmt.Sprint(result.Status.State))
}
