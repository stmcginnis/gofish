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

var powerDistributionBody = strings.NewReader(
	`{
		"@odata.type": "#PowerDistribution.v1_3_1.PowerDistribution",
		"Id": "1",
		"EquipmentType": "RackPDU",
		"Name": "RackPDU1",
		"FirmwareVersion": "4.3.0",
		"Version": "1.03b",
		"ProductionDate": "2017-01-11T08:00:00Z",
		"Manufacturer": "Contoso",
		"Model": "ZAP4000",
		"SerialNumber": "29347ZT536",
		"PartNumber": "AA-23",
		"UUID": "32354641-4135-4332-4a35-313735303734",
		"AssetTag": "PDX-92381",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Location": {
			"Placement": {
				"Row": "North 1"
			}
		},
		"Mains": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Mains"
		},
		"Branches": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Branches"
		},
		"Outlets": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets"
		},
		"OutletGroups": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/OutletGroups"
		},
		"Metrics": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Metrics"
		},
		"Sensors": {
			"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors"
		},
		"Links": {
			"Facility": {
				"@odata.id": "/redfish/v1/Facilities/Room237"
			}
		},
		"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1"
	}`)

// TestPowerDistribution tests the parsing of PowerDistribution objects.
func TestPowerDistribution(t *testing.T) {
	var result PowerDistribution
	err := json.NewDecoder(powerDistributionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "RackPDU", fmt.Sprint(result.EquipmentType))
	assertEquals(t, "RackPDU1", result.Name)
	assertEquals(t, "4.3.0", result.FirmwareVersion)
	assertEquals(t, "1.03b", result.Version)
	assertEquals(t, "2017-01-11T08:00:00Z", result.ProductionDate)
	assertEquals(t, "Contoso", result.Manufacturer)
	assertEquals(t, "ZAP4000", result.Model)
	assertEquals(t, "29347ZT536", result.SerialNumber)
	assertEquals(t, "AA-23", result.PartNumber)
	assertEquals(t, "32354641-4135-4332-4a35-313735303734", result.UUID)
	assertEquals(t, "PDX-92381", result.AssetTag)
	assertEquals(t, "North 1", result.Location.Placement.Row)
}
