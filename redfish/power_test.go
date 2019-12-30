//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var powerBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Power.Power",
		"@odata.type": "#Power.v1_5_3.Power",
		"@odata.id": "/redfish/v1/Power",
		"Id": "Power-1",
		"Name": "PowerOne",
		"Description": "Power One",
		"PowerControl": [{
			"@odata.id": "/redfish/v1/PowerControl",
			"MemberId": "PC1",
			"Name": "Fred",
			"PhysicalContext": "Upper",
			"PowerAllocatedWatts": 100.0,
			"PowerAvailableWatts": 100.0,
			"PowerCapacityWatts": 100.0,
			"PowerConsumedWatts": 100.0,
			"PowerLimit": {
				"CorrectionInMs": 10000,
				"LimitException": "HardPowerOff",
				"LimitInWatts": 1000.0
			},
			"PowerMetrics": {
				"AverageConsumedWatts": 1000.0,
				"IntervalInMin": 5,
				"MaxConsumedWatts": 1000.0,
				"MinConsumedWatts": 1000.0
			},
			"PowerRequestedWatts": 1000.0,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			}
		}],
		"PowerControl@odata.count": 1,
		"PowerSupplies": [{
			"@odata.id": "/redfish/v1/PowerSupply",
			"MemberId": "PS1",
			"Assembly": {
				"@odata.id": "/redfish/v1/Assembly/1"
			},
			"EfficiencyPercentage": 73,
			"FirmwareVersion": "1.0",
			"HotPluggable": true,
			"IndicatorLED": "Lit",
			"InputRanges": [{
					"InputType": "AC",
					"MaximumFrequencyHz": 99.0,
					"MaximumVoltage": 9.0,
					"MinimumFrequencyHz": 10.0,
					"MinimumVoltage": 1.0,
					"OutputWattage": 100.0
				},
				{
					"InputType": "DC",
					"MaximumFrequencyHz": 88.0,
					"MaximumVoltage": 9.0,
					"MinimumFrequencyHz": 8.0,
					"MinimumVoltage": 1.0,
					"OutputWattage": 22.0
				}
			],
			"LastPowerOutputWatts": 100.0,
			"LineInputVoltage": 9.0,
			"LineInputVoltageType": "ACandDCWideRange",
			"Location": {},
			"Manufacturer": "Acme Inc",
			"Model": "Power2000",
			"Name": "Power 2000",
			"PartNumber": "P2000",
			"PowerCapacityWatts": 100,
			"PowerInputWatts": 100,
			"PowerOutputWatts": 100,
			"PowerSupplyType": "ACorDC",
			"Redundancy": [],
			"Redundancy@odata.count": 0,
			"RelatedItems": [],
			"RelatedItems@odata.count": 0,
			"SerialNumber": "1234",
			"SparePartNumber": "P2000",
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			}
		}],
		"PowerSupplies@odata.count": 1,
		"Redundancy": [],
		"Redundancy@odata.count": 0,
		"Voltages": [{
			"@odata.id": "/redfish/v1/Voltage/1",
			"LowerThresholdCritical": 1.0,
			"LowerThresholdFatal": 0.0,
			"LowerThresholdNonCritical": 5.0,
			"MaxReadingRange": 10.0,
			"MemberId": "Voltage1",
			"MinReadingRange": 1.0,
			"Name": "Voltage-1",
			"PhysicalContext": "Upper",
			"ReadingVolts": 12.0,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"SensorNumber": 1,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"UpperThresholdCritical": 10000.0,
			"UpperThresholdFatal": 10001.0,
			"UpperThresholdNonCritical": 1000.0
		}],
		"Voltages@odata.count": 1
	}`)

// TestPower tests the parsing of Power objects.
func TestPower(t *testing.T) {
	var result Power
	err := json.NewDecoder(powerBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Power-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "PowerOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.PowerControl[0].PhysicalContext != common.UpperPhysicalContext {
		t.Errorf("Invalid physical context: %s", result.PowerControl[0].PhysicalContext)
	}

	if result.PowerControl[0].PowerLimit.CorrectionInMs != 10000 {
		t.Errorf("Invalid CorrectionInMs: %d", result.PowerControl[0].PowerLimit.CorrectionInMs)
	}

	if result.PowerControl[0].PowerLimit.LimitException != HardPowerOffPowerLimitException {
		t.Errorf("Invalid LimitException: %s", result.PowerControl[0].PowerLimit.LimitException)
	}

	if result.PowerSupplies[0].IndicatorLED != common.LitIndicatorLED {
		t.Errorf("Invalid PowerSupply IndicatorLED: %s",
			result.PowerSupplies[0].IndicatorLED)
	}

	if result.Voltages[0].MaxReadingRange != 10 {
		t.Errorf("Invalid MaxReadingRange: %f", result.Voltages[0].MaxReadingRange)
	}
}
