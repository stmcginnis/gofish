//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var powerSupplyUnitBody = strings.NewReader(
	`{
		"@odata.type": "#PowerSupply.v1_5_1.PowerSupply",
		"Id": "Bay1",
		"Name": "Power Supply Bay 1",
		"Status": {
			"State": "Enabled",
			"Health": "Warning"
		},
		"LineInputStatus": "Normal",
		"Model": "RKS-440DC",
		"Manufacturer": "Contoso Power",
		"FirmwareVersion": "1.00",
		"SerialNumber": "3488247",
		"PartNumber": "23456-133",
		"SparePartNumber": "93284-133",
		"LocationIndicatorActive": false,
		"HotPluggable": false,
		"PowerCapacityWatts": 400,
		"PhaseWiringType": "OnePhase3Wire",
		"PlugType": "IEC_60320_C14",
		"InputRanges": [
			{
				"NominalVoltageType": "AC200To240V",
				"CapacityWatts": 400
			},
			{
				"NominalVoltageType": "AC120V",
				"CapacityWatts": 350
			},
			{
				"NominalVoltageType": "DC380V",
				"CapacityWatts": 400
			}
		],
		"EfficiencyRatings": [
			{
				"LoadPercent": 25,
				"EfficiencyPercent": 75
			},
			{
				"LoadPercent": 50,
				"EfficiencyPercent": 85
			},
			{
				"LoadPercent": 90,
				"EfficiencyPercent": 80
			}
		],
		"OutputRails": [
			{
				"NominalVoltage": 3.3,
				"PhysicalContext": "SystemBoard"
			},
			{
				"NominalVoltage": 5,
				"PhysicalContext": "SystemBoard"
			},
			{
				"NominalVoltage": 12,
				"PhysicalContext": "StorageDevice"
			}
		],
		"Location": {
			"PartLocation": {
				"ServiceLabel": "PSU 1",
				"LocationType": "Bay",
				"LocationOrdinalValue": 0
			}
		},
		"Links": {
			"Outlet": {
				"@odata.id": "https://redfishpdu.contoso.com/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A4"
			}
		},
		"Assembly": {
			"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Assembly"
		},
		"Metrics": {
			"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Metrics"
		},
		"Actions": {
			"#PowerSupply.Reset": {
				"target": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/PowerSupply.Reset"
			}
		},
		"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1"
	}`)

// TestPowerSupplyUnit tests the parsing of PowerSupplyUnit objects.
func TestPowerSupplyUnit(t *testing.T) {
	var result PowerSupplyUnit
	err := json.NewDecoder(powerSupplyUnitBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Bay1", result.ID)
	assertEquals(t, "Power Supply Bay 1", result.Name)
	assertEquals(t, "RKS-440DC", result.Model)
	assertEquals(t, "Contoso Power", result.Manufacturer)
	assertEquals(t, "3488247", result.SerialNumber)
	assertEquals(t, "23456-133", result.PartNumber)
	assertEquals(t, "93284-133", result.SparePartNumber)
	assertEquals(t, "OnePhase3Wire", fmt.Sprint(result.PhaseWiringType))
	assertEquals(t, "IEC_60320_C14", fmt.Sprint(result.PlugType))
}
