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

var pcieSlotsBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#PCIeSlots.PCIeSlots",
		"@odata.etag": "\"1683342314\"",
		"@odata.id": "/redfish/v1/Chassis/Self/PCIeSlots",
		"@odata.type": "#PCIeSlots.v1_1_1.PCIeSlots",
		"Description": "Pcie Slot #11",
		"Id": "11",
		"Name": "PcieSlot_11",
		"Slots": [
		  {
			"HotPluggable": true,
			"Lanes": 32,
			"Links": {
			  "PCIeDevice": [
				{
				  "@odata.id": "/redfish/v1/Chassis/Self/PCIeDevices/00_00_01"
				}
			  ],
			  "PCIeDevice@odata.count": 1
			},
			"Location": {
			  "PartLocation": {
				"LocationOrdinalValue": 11,
				"LocationType": "Slot",
				"ServiceLabel": "Slot F"
			  }
			},
			"PCIeType": "Gen3",
			"SlotType": "FullLength",
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			}
		  }
		]
	  }`)

// TestPCIeSlots tests the parsing of PCIeSlots objects.
func TestPCIeSlots(t *testing.T) {
	var result PCIeSlots
	err := json.NewDecoder(pcieSlotsBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "11", result.ID)
	assertEquals(t, "PcieSlot_11", result.Name)
	assertEquals(t, "true", fmt.Sprint(result.Slots[0].HotPluggable))
	assertEquals(t, "32", fmt.Sprint(result.Slots[0].Lanes))
	assertEquals(t, "Gen3", fmt.Sprint(result.Slots[0].PCIeType))
	assertEquals(t, "FullLength", fmt.Sprint(result.Slots[0].SlotType))
}
