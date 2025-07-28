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

// powerBody contains a complete Power resource example for testing.
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
			"Redundancy": [
				{
				"@odata.id": "/redfish/v1/Chassis/1/Power#/Redundancy/0"
				}
			],
			"Redundancy@odata.count": 1,
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
		"Redundancy": [
		{
			"@odata.id": "/redfish/v1/Chassis/1/Power#/Redundancy/0",
			"@odata.type": "#Redundancy.v1_2_0.Redundancy",
			"MemberId": "0",
			"Name": "PowerSupply Redundancy Group 1",
			"Mode": "Failover",
			"MaxNumSupported": 4,
			"MinNumNeeded": 1,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"RedundancySet": [
				{
				"@odata.id": "/redfish/v1/Chassis/1/Power#/PowerSupplies/0"
				},
				{
				"@odata.id": "/redfish/v1/Chassis/1/Power#/PowerSupplies/1"
				}
			]
			}
		],
		"Redundancy@odata.count": 1,
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

// invalidPowerBody contains an invalid Power resource example for testing error cases.
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

// TestPower verifies the parsing of Power objects from JSON.
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

	if *result.PowerControl[0].PowerLimit.CorrectionInMs != 10000 {
		t.Errorf("Invalid CorrectionInMs: %d", result.PowerControl[0].PowerLimit.CorrectionInMs)
	}

	if result.PowerControl[0].PowerLimit.LimitException != HardPowerOffPowerLimitException {
		t.Errorf("Invalid LimitException: %s", result.PowerControl[0].PowerLimit.LimitException)
	}

	if result.PowerSupplies[0].IndicatorLED != common.LitIndicatorLED {
		t.Errorf("Invalid PowerSupply IndicatorLED: %s",
			result.PowerSupplies[0].IndicatorLED)
	}

	if *result.Voltages[0].MaxReadingRange != 10 {
		t.Errorf("Invalid MaxReadingRange: %f", *result.Voltages[0].MaxReadingRange)
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

// TestVoltageUnmarshalJSON_MemberIDAsString verifies parsing Voltage with string MemberId.
func TestVoltageUnmarshalJSON_MemberIDAsString(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/1/Power#/Voltages/1",
		"MemberId": "1",
		"Name": "CPU Voltage",
		"ReadingVolts": 3.3,
		"RelatedItem@odata.count": 2,
		"RelatedItem": [
			{"@odata.id": "/redfish/v1/Chassis/1"},
			{"@odata.id": "/redfish/v1/Systems/1"}
		]
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.MemberID != "1" {
		t.Errorf("Expected MemberID '1', got '%s'", voltage.MemberID)
	}
	if voltage.Name != "CPU Voltage" {
		t.Errorf("Expected Name 'CPU Voltage', got '%s'", voltage.Name)
	}
	if *voltage.ReadingVolts != 3.3 {
		t.Errorf("Expected ReadingVolts 3.3, got %f", *voltage.ReadingVolts)
	}
	if len(voltage.RelatedItem) != 2 {
		t.Errorf("Expected 2 RelatedItems, got %d", len(voltage.RelatedItem))
	}
}

// TestVoltageUnmarshalJSON_MemberIDAsNumber verifies parsing Voltage with numeric MemberId.
func TestVoltageUnmarshalJSON_MemberIDAsNumber(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/1/Power#/Voltages/2",
		"MemberId": 2,
		"Name": "Memory Voltage",
		"ReadingVolts": 1.2,
		"RelatedItem@odata.count": 1,
		"RelatedItem": [
			{"@odata.id": "/redfish/v1/Chassis/1"}
		]
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.MemberID != "2" {
		t.Errorf("Expected MemberID '2', got '%s'", voltage.MemberID)
	}
	if voltage.Name != "Memory Voltage" {
		t.Errorf("Expected Name 'Memory Voltage', got '%s'", voltage.Name)
	}
	if *voltage.ReadingVolts != 1.2 {
		t.Errorf("Expected ReadingVolts 1.2, got %f", *voltage.ReadingVolts)
	}
	if len(voltage.RelatedItem) != 1 {
		t.Errorf("Expected 1 RelatedItem, got %d", len(voltage.RelatedItem))
	}
}

// TestVoltageUnmarshalJSON_InvalidJSON verifies error handling for invalid JSON.
func TestVoltageUnmarshalJSON_InvalidJSON(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/1/Power#/Voltages/3",
		"MemberId": "3",
		"ReadingVolts: 5.0,  // Missing quote
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err == nil {
		t.Error("Expected error for invalid JSON, but got none")
	}
}

// TestVoltageUnmarshalJSON_OptionalFields verifies parsing with optional fields.
func TestVoltageUnmarshalJSON_OptionalFields(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/1/Power#/Voltages/4",
		"MemberId": "4",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.MemberID != "4" {
		t.Errorf("Expected MemberID '4', got '%s'", voltage.MemberID)
	}
	if voltage.Status.State != "Enabled" {
		t.Errorf("Expected Status.State 'Enabled', got '%s'", voltage.Status.State)
	}
	if voltage.Status.Health != common.OKHealth {
		t.Errorf("Expected Status.Health 'OK', got '%s'", voltage.Status.Health)
	}
	if voltage.ReadingVolts != nil {
		t.Error("Expected ReadingVolts to be nil")
	}
}

// TestVoltageUnmarshalJSON_Thresholds verifies parsing of threshold values.
func TestVoltageUnmarshalJSON_Thresholds(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/1/Power#/Voltages/5",
		"MemberId": "5",
		"LowerThresholdCritical": 2.8,
		"UpperThresholdCritical": 3.6,
		"ReadingVolts": 3.3
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if *voltage.LowerThresholdCritical != 2.8 {
		t.Errorf("Expected LowerThresholdCritical 2.8, got %f", *voltage.LowerThresholdCritical)
	}
	if *voltage.UpperThresholdCritical != 3.6 {
		t.Errorf("Expected UpperThresholdCritical 3.6, got %f", *voltage.UpperThresholdCritical)
	}
	if voltage.LowerThresholdFatal != nil {
		t.Errorf("Expected UpperThresholdCritical nil, got %f", *voltage.UpperThresholdCritical)
	}
	if voltage.UpperThresholdFatal != nil {
		t.Errorf("Expected UpperThresholdFatal nil, got %f", *voltage.UpperThresholdCritical)
	}
}

// TestVoltageUnmarshalJSON_CompleteData verifies parsing of complete Voltage data.
func TestVoltageUnmarshalJSON_CompleteData(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/0",
		"LowerThresholdCritical": 1.1,
		"LowerThresholdNonCritical": 1.15,
		"MaxReadingRange": 255,
		"MemberId": "VR_PVDQ_EFGH1_Output_Voltage",
		"MinReadingRange": 0,
		"Name": "VR PVDQ EFGH1 Output Voltage",
		"ReadingVolts": 1.24,
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"UpperThresholdCritical": 1.35,
		"UpperThresholdNonCritical": 1.3
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.ODataID != "/redfish/v1/Chassis/Serve/Power#/Voltages/0" {
		t.Errorf("Expected ODataID '/redfish/v1/Chassis/Serve/Power#/Voltages/0', got '%s'", voltage.ODataID)
	}
	if *voltage.LowerThresholdCritical != 1.1 {
		t.Errorf("Expected LowerThresholdCritical 1.1, got %f", *voltage.LowerThresholdCritical)
	}
	if *voltage.LowerThresholdNonCritical != 1.15 {
		t.Errorf("Expected LowerThresholdNonCritical 1.15, got %f", *voltage.LowerThresholdNonCritical)
	}
	if *voltage.MaxReadingRange != 255 {
		t.Errorf("Expected MaxReadingRange 255, got %f", *voltage.MaxReadingRange)
	}
	if voltage.MemberID != "VR_PVDQ_EFGH1_Output_Voltage" {
		t.Errorf("Expected MemberID 'VR_PVDQ_EFGH1_Output_Voltage', got '%s'", voltage.MemberID)
	}
	if *voltage.MinReadingRange != 0 {
		t.Errorf("Expected MinReadingRange 0, got %f", *voltage.MinReadingRange)
	}
	if voltage.Name != "VR PVDQ EFGH1 Output Voltage" {
		t.Errorf("Expected Name 'VR PVDQ EFGH1 Output Voltage', got '%s'", voltage.Name)
	}
	if *voltage.ReadingVolts != 1.24 {
		t.Errorf("Expected ReadingVolts 1.24, got %f", *voltage.ReadingVolts)
	}
	if voltage.Status.Health != common.OKHealth {
		t.Errorf("Expected Status.Health 'OK', got '%s'", voltage.Status.Health)
	}
	if voltage.Status.State != "Enabled" {
		t.Errorf("Expected Status.State 'Enabled', got '%s'", voltage.Status.State)
	}
	if *voltage.UpperThresholdCritical != 1.35 {
		t.Errorf("Expected UpperThresholdCritical 1.35, got %f", *voltage.UpperThresholdCritical)
	}
	if *voltage.UpperThresholdNonCritical != 1.3 {
		t.Errorf("Expected UpperThresholdNonCritical 1.3, got %f", *voltage.UpperThresholdNonCritical)
	}
}

// TestVoltageUnmarshalJSON_NullThresholds verifies parsing with null threshold values.
func TestVoltageUnmarshalJSON_NullThresholds(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/31",
		"LowerThresholdCritical": null,
		"LowerThresholdNonCritical": 198,
		"MaxReadingRange": 300,
		"MemberId": "PSU0_Input_Voltage",
		"MinReadingRange": 0,
		"Name": "PSU0 Input Voltage",
		"ReadingVolts": 230.75,
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"UpperThresholdCritical": 260,
		"UpperThresholdNonCritical": 253
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.LowerThresholdCritical != nil {
		t.Error("Expected LowerThresholdCritical to be nil")
	}
	if *voltage.LowerThresholdNonCritical != 198 {
		t.Errorf("Expected LowerThresholdNonCritical 198, got %f", *voltage.LowerThresholdNonCritical)
	}
	if *voltage.ReadingVolts != 230.75 {
		t.Errorf("Expected ReadingVolts 230.75, got %f", *voltage.ReadingVolts)
	}
}

// TestVoltageUnmarshalJSON_MinimalData verifies parsing with minimal required fields.
func TestVoltageUnmarshalJSON_MinimalData(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/4",
		"MemberId": "VR_PVCSA1_Output_Voltage",
		"ReadingVolts": 0.851
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if voltage.ODataID != "/redfish/v1/Chassis/Serve/Power#/Voltages/4" {
		t.Errorf("Unexpected ODataID: %s", voltage.ODataID)
	}
	if voltage.MemberID != "VR_PVCSA1_Output_Voltage" {
		t.Errorf("Unexpected MemberID: %s", voltage.MemberID)
	}
	if *voltage.ReadingVolts != 0.851 {
		t.Errorf("Unexpected ReadingVolts: %f", *voltage.ReadingVolts)
	}
	if voltage.LowerThresholdCritical != nil {
		t.Error("Expected LowerThresholdCritical to be nil")
	}
	if voltage.Name != "" {
		t.Error("Expected Name to be empty")
	}
}

// TestVoltageUnmarshalJSON_InvalidData verifies error handling for invalid data.
func TestVoltageUnmarshalJSON_InvalidData(t *testing.T) {
	jsonData := `{
		"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/5",
		"MemberId": "VR_P5V_AUX_Output_Voltage",
		"ReadingVolts: 5.243260534504318
	}`

	var voltage Voltage
	err := json.Unmarshal([]byte(jsonData), &voltage)
	if err == nil {
		t.Error("Expected error for invalid JSON, but got none")
	}
}

// TestVoltageUnmarshalJSON_AllVoltages verifies parsing of multiple Voltage objects.
func TestVoltageUnmarshalJSON_AllVoltages(t *testing.T) {
	fullJSON := `{
		"@odata.context": "/redfish/v1/$metadata#Power.Power",
		"@odata.id": "/redfish/v1/Chassis/Serve/Power/",
		"@odata.type": "#Power.v1_7_1.Power",
		"Description": "The Power schema describes power metrics and represents the properties for power consumption and power limiting.",
		"Id": "Power",
		"Name": "Power",
		"Voltages": [
			{
				"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/0",
				"LowerThresholdCritical": 1.1,
				"LowerThresholdNonCritical": 1.15,
				"MaxReadingRange": 255,
				"MemberId": "VR_PVDQ_EFGH1_Output_Voltage",
				"MinReadingRange": 0,
				"Name": "VR PVDQ EFGH1 Output Voltage",
				"ReadingVolts": 1.24,
				"Status": {
					"Health": "OK",
					"State": "Enabled"
				},
				"UpperThresholdCritical": 1.35,
				"UpperThresholdNonCritical": 1.3
			},
			{
				"@odata.id": "/redfish/v1/Chassis/Serve/Power#/Voltages/30",
				"LowerThresholdCritical": null,
				"LowerThresholdNonCritical": 198,
				"MaxReadingRange": 300,
				"MemberId": "PSU1_Input_Voltage",
				"MinReadingRange": 0,
				"Name": "PSU1 Input Voltage",
				"ReadingVolts": 229.25,
				"Status": {
					"Health": "OK",
					"State": "Enabled"
				},
				"UpperThresholdCritical": 260,
				"UpperThresholdNonCritical": 253
			}
		]
	}`

	type PowerResponse struct {
		Voltages []Voltage `json:"Voltages"`
	}

	var power PowerResponse
	err := json.Unmarshal([]byte(fullJSON), &power)
	if err != nil {
		t.Fatalf("Failed to unmarshal full JSON: %v", err)
	}

	if len(power.Voltages) != 2 {
		t.Fatalf("Expected 2 voltages, got %d", len(power.Voltages))
	}

	v0 := power.Voltages[0]
	if v0.MemberID != "VR_PVDQ_EFGH1_Output_Voltage" {
		t.Errorf("First voltage: expected MemberID 'VR_PVDQ_EFGH1_Output_Voltage', got '%s'", v0.MemberID)
	}

	v30 := power.Voltages[1]
	if v30.MemberID != "PSU1_Input_Voltage" {
		t.Errorf("Second voltage: expected MemberID 'PSU1_Input_Voltage', got '%s'", v30.MemberID)
	}
	if v30.LowerThresholdCritical != nil {
		t.Error("Second voltage: expected LowerThresholdCritical to be nil")
	}
}
