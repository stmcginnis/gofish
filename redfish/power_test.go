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

var invalidPowerBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Chassis/Members(*)/Self/Power/$entity",
		"@odata.etag": "W/\"1604509181\"",
		"@odata.id": "/redfish/v1/Chassis/Self/Power",
		"@odata.type": "#Power.v1_2_1.Power",
		"Id": "Power",
		"Name": "Power",
		"PowerControl": [
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/PowerControl/0",
			"MemberId": 0,
			"Name": "Chassis Power Control",
			"PowerLimit": {
			  "CorrectionInMs": 1000,
			  "LimitException": "NoAction",
			  "LimitInWatts": 500
			},
			"PowerMetrics": {
			  "AverageConsumedWatts": 148,
			  "IntervalInMin": 0.083333333333333,
			  "MaxConsumedWatts": 301,
			  "MinConsumedWatts": 0
			},
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			}
		  }
		],
		"PowerSupplies": [
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/PowerSupplies/0",
			"@odata.type": "#Power.v1_2_1.PowerSupply",
			"FirmwareVersion": "00.04.04",
			"InputRanges": [
			  {
				"MaximumVoltage": 264,
				"MinimumVoltage": 90,
				"OutputWattage": 128
			  }
			],
			"LastPowerOutputWatts": 103,
			"LineInputVoltage": 241,
			"Manufacturer": "Liteon Power",
			"MemberId": "1",
			"Model": "PS-2122-7Q",
			"Name": "PSU1",
			"PowerCapacityWatts": 1200,
			"SerialNumber": "6D7QX0101J224CV",
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			}
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/PowerSupplies/1",
			"@odata.type": "#Power.v1_2_1.PowerSupply",
			"FirmwareVersion": "00.04.04",
			"InputRanges": [
			  {
				"MaximumVoltage": 264,
				"MinimumVoltage": 90,
				"OutputWattage": 150
			  }
			],
			"LastPowerOutputWatts": 123,
			"LineInputVoltage": 241,
			"Manufacturer": "Liteon Power",
			"MemberId": "0",
			"Model": "PS-2122-7Q",
			"Name": "PSU0",
			"PowerCapacityWatts": 1200,
			"SerialNumber": "6D7QX0101J2247A",
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			}
		  }
		],
		"Voltages": [
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/0",
			"LowerThresholdCritical": 1.431,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 218,
			"MinReadingRange": 0,
			"Name": "Volt_PVCCIN_CPU1",
			"ReadingVolts": 1.692,
			"SensorNumber": 218,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 2.205,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/1",
			"LowerThresholdCritical": 1.078,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 223,
			"MinReadingRange": 0,
			"Name": "Volt_CPU1_DEF",
			"ReadingVolts": 1.218,
			"SensorNumber": 223,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.323,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/2",
			"LowerThresholdCritical": 2.975,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 209,
			"MinReadingRange": 0,
			"Name": "Volt_P3V3",
			"ReadingVolts": 3.264,
			"SensorNumber": 209,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 3.621,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/3",
			"LowerThresholdCritical": 1.078,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 222,
			"MinReadingRange": 0,
			"Name": "Volt_CPU1_ABC",
			"ReadingVolts": 1.218,
			"SensorNumber": 222,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.323,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/4",
			"LowerThresholdCritical": 10.773,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 208,
			"MinReadingRange": 0,
			"Name": "Volt_P12V",
			"ReadingVolts": 12.033,
			"SensorNumber": 208,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 13.23,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/5",
			"LowerThresholdCritical": 0.882,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 217,
			"MinReadingRange": 0,
			"Name": "Volt_PVCCIO_CPU0",
			"ReadingVolts": 0.973,
			"SensorNumber": 217,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.057,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/6",
			"LowerThresholdCritical": 1.078,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 220,
			"MinReadingRange": 0,
			"Name": "Volt_CPU0_ABC",
			"ReadingVolts": 1.218,
			"SensorNumber": 220,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.323,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/7",
			"LowerThresholdCritical": 0.763,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 213,
			"MinReadingRange": 0,
			"Name": "Volt_PVNN_PCH",
			"ReadingVolts": 0.987,
			"SensorNumber": 213,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.106,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/8",
			"LowerThresholdCritical": 0.882,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 219,
			"MinReadingRange": 0,
			"Name": "Volt_PVCCIO_CPU1",
			"ReadingVolts": 0.987,
			"SensorNumber": 219,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.057,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/9",
			"LowerThresholdCritical": 11.214,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 212,
			"MinReadingRange": 0,
			"Name": "Volt_P12V_AUX",
			"ReadingVolts": 12.033,
			"SensorNumber": 212,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 13.041,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/10",
			"LowerThresholdCritical": 4.498,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 210,
			"MinReadingRange": 0,
			"Name": "Volt_P5V",
			"ReadingVolts": 5.018,
			"SensorNumber": 210,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 5.538,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/11",
			"LowerThresholdCritical": 1.078,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 221,
			"MinReadingRange": 0,
			"Name": "Volt_CPU0_DEF",
			"ReadingVolts": 1.218,
			"SensorNumber": 221,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.323,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/12",
			"LowerThresholdCritical": 1.62,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 214,
			"MinReadingRange": 0,
			"Name": "Volt_P1V8_PCH",
			"ReadingVolts": 1.71,
			"SensorNumber": 214,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.989,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/13",
			"LowerThresholdCritical": 0.945,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 211,
			"MinReadingRange": 0,
			"Name": "Volt_P1V05_PCH",
			"ReadingVolts": 1.036,
			"SensorNumber": 211,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 1.155,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/14",
			"LowerThresholdCritical": 1.431,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 216,
			"MinReadingRange": 0,
			"Name": "Volt_PVCCIN_CPU0",
			"ReadingVolts": 1.692,
			"SensorNumber": 216,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 2.205,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/Voltages/15",
			"LowerThresholdCritical": 2.52,
			"LowerThresholdFatal": null,
			"LowerThresholdNonCritical": null,
			"MemberId": 215,
			"MinReadingRange": 0,
			"Name": "Volt_P3V_BAT",
			"ReadingVolts": 3.003,
			"SensorNumber": 215,
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			},
			"UpperThresholdCritical": 3.591,
			"UpperThresholdFatal": null,
			"UpperThresholdNonCritical": null
		  }
		]
	  }
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

// TestNonconformingPower tests the parsing of nonconforming Power objects.
// Some Dell implementations return MemberID as an integer when they should be
// strings.
func TestNonconformingPower(t *testing.T) {
	var result Power
	err := json.NewDecoder(invalidPowerBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.PowerControl[0].MemberID != "0" {
		t.Errorf("Expected first PowerController MemberID to be '0': %s", result.PowerControl[0].MemberID)
	}

	voltage := result.Voltages[0]
	if voltage.MemberID != "218" {
		t.Errorf("Expected first Voltage MemberID to be '218': %s", voltage.MemberID)
	}
}
