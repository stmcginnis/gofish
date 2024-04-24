//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var coolingUnitBody = `{
	"@odata.type": "#CoolingUnit.v1_1_1.CoolingUnit",
	"Id": "1",
	"EquipmentType": "CDU",
	"Name": "Rack #4 Cooling Distribution Unit",
	"FirmwareVersion": "3.2.0",
	"Version": "1.03b",
	"ProductionDate": "2020-12-24T08:00:00Z",
	"Manufacturer": "Contoso",
	"Model": "BRRR4000",
	"SerialNumber": "29347ZT536",
	"PartNumber": "ICE-9",
	"AssetTag": "PDX5-92381",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Location": {
	  "Placement": {
		"Row": "North 1"
	  }
	},
	"PrimaryCoolantConnectors": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/PrimaryCoolantConnectors"
	},
	"SecondaryCoolantConnectors": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/SecondaryCoolantConnectors"
	},
	"Pumps": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/Pumps"
	},
	"Filters": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/Filters"
	},
	"EnvironmentMetrics": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/EnvironmentMetrics"
	},
	"LeakDetection": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection"
	},
	"Links": {
	  "Facility": {
		"@odata.id": "/redfish/v1/Facilities/Room237"
	  }
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1"
  }`

// TestCoolingUnit tests the parsing of CoolingUnit objects.
func TestCoolingUnit(t *testing.T) {
	var result CoolingUnit
	err := json.NewDecoder(strings.NewReader(coolingUnitBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Rack #4 Cooling Distribution Unit", result.Name)
	assertEquals(t, "CDU", string(result.EquipmentType))
	assertEquals(t, "3.2.0", result.FirmwareVersion)
	assertEquals(t, "1.03b", result.Version)
	assertEquals(t, "29347ZT536", result.SerialNumber)
	assertEquals(t, "North 1", result.Location.Placement.Row)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/PrimaryCoolantConnectors", result.primaryCoolantConnectors)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/SecondaryCoolantConnectors", result.secondaryCoolantConnectors)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/Pumps", result.pumps)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/Filters", result.filters)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/EnvironmentMetrics", result.environmentMetrics)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection", result.leakDetection)
	assertEquals(t, "/redfish/v1/Facilities/Room237", result.facility)
}
