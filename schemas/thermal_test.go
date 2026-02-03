//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var thermalBody = `{
	"@odata.type": "#Thermal.v1_6_0.Thermal",
	"@odata.id": "/redfish/v1/Chassis/1/Thermal",
	"Id": "Thermal-1",
	"Name": "Thermal",
	"Temperatures": [
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/0",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "0",
		"Name": "CPU Temp",
		"SensorNumber": 1,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 63,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 100,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 102,
		"PhysicalContext": "CPU",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/1",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "1",
		"Name": "Inlet Temp",
		"SensorNumber": 9,
		"Status": {
		  "State": "Absent"
		},
		"ReadingCelsius": 0,
		"UpperThresholdCritical": 0,
		"UpperThresholdFatal": 0,
		"LowerThresholdCritical": 0,
		"LowerThresholdFatal": 0,
		"MinReadingRangeTemp": 0,
		"MaxReadingRangeTemp": 0,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/2",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "2",
		"Name": "System Temp",
		"SensorNumber": 11,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 30,
		"UpperThresholdCritical": 85,
		"UpperThresholdFatal": 90,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 92,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/3",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "3",
		"Name": "Peripheral Temp",
		"SensorNumber": 12,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 44,
		"UpperThresholdCritical": 85,
		"UpperThresholdFatal": 90,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 92,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/4",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "4",
		"Name": "CPU_VRM Temp",
		"SensorNumber": 16,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 41,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "CPU",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/5",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "5",
		"Name": "SOC_VRM Temp",
		"SensorNumber": 18,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 57,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/6",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "6",
		"Name": "VRMABCD Temp",
		"SensorNumber": 20,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 52,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/7",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "7",
		"Name": "VRMEFGH Temp",
		"SensorNumber": 22,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 51,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "SystemBoard",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/8",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "8",
		"Name": "DIMMABCD Temp",
		"SensorNumber": 176,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 39,
		"UpperThresholdCritical": 85,
		"UpperThresholdFatal": 90,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 92,
		"PhysicalContext": "Memory",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/9",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "9",
		"Name": "DIMMEFGH Temp",
		"SensorNumber": 177,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 38,
		"UpperThresholdCritical": 85,
		"UpperThresholdFatal": 90,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 92,
		"PhysicalContext": "Memory",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/10",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "10",
		"Name": "M2_SSD1 Temp",
		"SensorNumber": 140,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 37,
		"UpperThresholdCritical": 70,
		"UpperThresholdFatal": 75,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 77,
		"PhysicalContext": "StorageDevice",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/11",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "11",
		"Name": "M2_SSD2 Temp",
		"SensorNumber": 141,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 46,
		"UpperThresholdCritical": 70,
		"UpperThresholdFatal": 75,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 77,
		"PhysicalContext": "StorageDevice",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/12",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "12",
		"Name": "AIOM_NIC1 Temp",
		"SensorNumber": 160,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 47,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "ASIC",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1/NetworkAdapters/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/13",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "13",
		"Name": "AOC_NIC1 Temp",
		"SensorNumber": 161,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 71,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "ASIC",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1/NetworkAdapters/2"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Temperatures/14",
		"@odata.type": "#Thermal.v1_6_0.Temperature",
		"MemberId": "14",
		"Name": "AOC_NIC2 Temp",
		"SensorNumber": 162,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingCelsius": 61,
		"UpperThresholdCritical": 100,
		"UpperThresholdFatal": 105,
		"LowerThresholdCritical": 5,
		"LowerThresholdFatal": 5,
		"MinReadingRangeTemp": 3,
		"MaxReadingRangeTemp": 107,
		"PhysicalContext": "ASIC",
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1/NetworkAdapters/3"
		  }
		]
	  }
	],
	"Fans": [
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/0",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "0",
		"Name": "FAN1",
		"SensorNumber": 65,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10780,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/1",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "1",
		"Name": "FAN2",
		"SensorNumber": 66,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10920,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/2",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "2",
		"Name": "FAN3",
		"SensorNumber": 67,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10920,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/3",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "3",
		"Name": "FAN4",
		"SensorNumber": 68,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10920,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/4",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "4",
		"Name": "FAN5",
		"SensorNumber": 69,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10780,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  },
	  {
		"@odata.id": "/redfish/v1/Chassis/1/Thermal#/Fans/5",
		"@odata.type": "#Thermal.v1_6_0.Fan",
		"MemberId": "5",
		"Name": "FAN6",
		"SensorNumber": 70,
		"PhysicalContext": "Fan",
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		},
		"ReadingUnits": "RPM",
		"Reading": 10780,
		"UpperThresholdCritical": 35560,
		"UpperThresholdFatal": 35700,
		"LowerThresholdCritical": 420,
		"LowerThresholdFatal": 280,
		"MinReadingRange": 180,
		"MaxReadingRange": 35800,
		"RelatedItem": [
		  {
			"@odata.id": "/redfish/v1/Systems/1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1"
		  }
		]
	  }
	]
  }
  `

// TestThermal tests the parsing of Thermal objects.
func TestThermal(t *testing.T) {
	var result Thermal
	err := json.NewDecoder(strings.NewReader(thermalBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Thermal-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Thermal" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Fans[0].Name != "FAN1" {
		t.Errorf("Invalid fan name: %s", result.Fans[0].Name)
	}
}

func TestThermalFanParsing(t *testing.T) {
	var fan ThermalFan
	err := json.NewDecoder(strings.NewReader(`{
		"@odata.id": "/redfish/v1/Chassis/1/Thermal/Fans/1",
		"Name": "FAN1",
		"Reading": 10780,
		"ReadingUnits": "RPM",
		"UpperThresholdCritical": 35560,
		"LowerThresholdCritical": 420,
		"PhysicalContext": "Fan",
		"HotPluggable": true,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)).Decode(&fan)

	if err != nil {
		t.Fatalf("Error decoding fan JSON: %v", err)
	}

	if fan.Name != "FAN1" {
		t.Errorf("Expected fan name 'FAN1', got '%s'", fan.Name)
	}

	if fan.Reading == nil || *fan.Reading != 10780 {
		t.Errorf("Expected fan reading 10780, got %v", fan.Reading)
	}

	if fan.ReadingUnits != "RPM" {
		t.Errorf("Expected fan units RPM, got %v", fan.ReadingUnits)
	}

	if fan.PhysicalContext != FanPhysicalContext {
		t.Errorf("Expected physical context 'Fan', got %v", fan.PhysicalContext)
	}

	if !fan.HotPluggable {
		t.Error("Expected fan to be hot-pluggable")
	}

	if fan.Status.Health != OKHealth {
		t.Error("Expected fan health to be OK")
	}
}

func TestTemperatureParsing(t *testing.T) {
	var temp Temperature
	err := json.NewDecoder(strings.NewReader(`{
		"@odata.id": "/redfish/v1/Chassis/1/Thermal/Temperatures/0",
		"Name": "CPU Temp",
		"ReadingCelsius": 63,
		"UpperThresholdCritical": 100,
		"LowerThresholdCritical": 5,
		"PhysicalContext": "CPU",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"RelatedItem": [
			{"@odata.id": "/redfish/v1/Systems/1/Processors/1"}
		]
	}`)).Decode(&temp)

	if err != nil {
		t.Fatalf("Error decoding temperature JSON: %v", err)
	}

	if temp.Name != "CPU Temp" {
		t.Errorf("Expected temp name 'CPU Temp', got '%s'", temp.Name)
	}

	if temp.ReadingCelsius == nil || *temp.ReadingCelsius != 63 {
		t.Errorf("Expected temp reading 63, got %v", temp.ReadingCelsius)
	}

	if temp.PhysicalContext != "CPU" {
		t.Errorf("Expected physical context 'CPU', got %v", temp.PhysicalContext)
	}

	if len(temp.relatedItem) != 1 || temp.relatedItem[0] != "/redfish/v1/Systems/1/Processors/1" {
		t.Errorf("Expected related item to processor, got %v", temp.relatedItem)
	}
}

func TestThermalFanUpdate(t *testing.T) {
	var fan ThermalFan
	err := json.NewDecoder(strings.NewReader(`{
		"@odata.id": "/redfish/v1/Chassis/1/Thermal/Fans/1",
		"IndicatorLED": "Off",
		"Reading": 10780
	}`)).Decode(&fan)

	if err != nil {
		t.Fatalf("Error decoding fan JSON: %v", err)
	}

	testClient := &TestClient{}
	fan.SetClient(testClient)

	// Change LED state
	fan.IndicatorLED = BlinkingIndicatorLED

	// Test update
	err = fan.Update()
	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected 1 API call, found %d", len(calls))
	}

	if !strings.Contains(calls[0].Payload, "IndicatorLED:Blinking") {
		t.Errorf("Unexpected IndicatorLED update payload: %s", calls[0].Payload)
	}
}

func TestTemperatureUpdate(t *testing.T) {
	var temp Temperature
	err := json.NewDecoder(strings.NewReader(`{
		"@odata.id": "/redfish/v1/Chassis/1/Thermal/Temperatures/0",
		"LowerThresholdUser": 30.0,
		"UpperThresholdUser": 80.0
	}`)).Decode(&temp)

	if err != nil {
		t.Fatalf("Error decoding temperature JSON: %v", err)
	}

	testClient := &TestClient{}
	temp.SetClient(testClient)

	// Change thresholds
	temp.LowerThresholdUser = toRef(float32(35.0))
	temp.UpperThresholdUser = toRef(float32(75.0))

	// Test update
	err = temp.Update()
	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected 1 API call, found %d", len(calls))
	}

	if !strings.Contains(calls[0].Payload, "LowerThresholdUser:35") {
		t.Errorf("Unexpected LowerThresholdUser update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "UpperThresholdUser:75") {
		t.Errorf("Unexpected LowerThresholdUser update payload: %s", calls[0].Payload)
	}
}

func TestServerThermalParsing(t *testing.T) {
	var thermal Thermal
	err := json.NewDecoder(strings.NewReader(serverThermalBody)).Decode(&thermal)
	if err != nil {
		t.Fatalf("Error decoding thermal JSON: %v", err)
	}

	if thermal.ID != "Thermal" {
		t.Errorf("Expected Thermal ID 'Thermal', got '%s'", thermal.ID)
	}
	if thermal.Name != "Thermal" {
		t.Errorf("Expected Thermal name 'Thermal', got '%s'", thermal.Name)
	}

	if len(thermal.Fans) != 7 {
		t.Errorf("Expected 7 fans, got %d", len(thermal.Fans))
	} else {
		fan := thermal.Fans[0]
		if fan.Name != "System Fan 3" {
			t.Errorf("Expected fan name 'System Fan 3', got '%s'", fan.Name)
		}
		if *fan.Reading != 45 {
			t.Errorf("Expected fan reading 45, got %d", *fan.Reading)
		}
		if fan.ReadingUnits != "Percent" {
			t.Errorf("Expected fan units Percent, got %v", fan.ReadingUnits)
		}
		if fan.Status.Health != OKHealth {
			t.Error("Expected fan health to be OK")
		}
	}

	if len(thermal.Temperatures) != 51 {
		t.Errorf("Expected 51 temperature sensors, got %d", len(thermal.Temperatures))
	} else {
		temp := thermal.Temperatures[3] // CPU1 Die
		if temp.Name != "CPU1 Die" {
			t.Errorf("Expected temp name 'CPU1 Die', got '%s'", temp.Name)
		}
		if *temp.ReadingCelsius != 48.157 {
			t.Errorf("Expected temp reading ~48.157, got %v", *temp.ReadingCelsius)
		}
		if temp.Status.Health != OKHealth {
			t.Error("Expected temp health to be OK")
		}

		nullTemp := thermal.Temperatures[8] // RSR Temp 1 ext
		if nullTemp.ReadingCelsius != nil {
			t.Error("Expected null temp reading to be nil")
		}
		if nullTemp.Status.State != "Absent" {
			t.Errorf("Expected absent state, got %s", nullTemp.Status.State)
		}
	}
}

func TestServerFanOemParsing(t *testing.T) {
	var thermal Thermal
	err := json.NewDecoder(strings.NewReader(serverThermalBody)).Decode(&thermal)
	if err != nil {
		t.Fatalf("Error decoding thermal JSON: %v", err)
	}

	fan := thermal.Fans[0]
	if len(fan.OEM) == 0 {
		t.Fatal("Expected OEM data in fan")
	}

	var oem struct {
		Temp struct {
			Connector string `json:"Connector"`
			Readings  []struct {
				Reading float64 `json:"Reading"`
			} `json:"Readings"`
		} `json:"Temp"`
	}
	err = json.Unmarshal(fan.OEM, &oem)
	if err != nil {
		t.Fatalf("Error decoding fan OEM: %v", err)
	}

	if oem.Temp.Connector != "Fan3" {
		t.Errorf("Expected connector 'Fan3', got '%s'", oem.Temp.Connector)
	}
	if len(oem.Temp.Readings) != 2 {
		t.Errorf("Expected 2 readings, got %d", len(oem.Temp.Readings))
	}
	if oem.Temp.Readings[0].Reading != 12032 {
		t.Errorf("Expected first reading 12032, got %v", oem.Temp.Readings[0].Reading)
	}
}

func TestServerTemperatureThresholds(t *testing.T) {
	var thermal Thermal
	err := json.NewDecoder(strings.NewReader(serverThermalBody)).Decode(&thermal)
	if err != nil {
		t.Fatalf("Error decoding thermal JSON: %v", err)
	}
	temp := thermal.Temperatures[42] // Outlet Air Temp
	if *temp.LowerThresholdCritical != 0 {
		t.Errorf("Expected LowerThresholdCritical 0, got %v", *temp.LowerThresholdCritical)
	}
	if *temp.UpperThresholdCritical != 61 {
		t.Errorf("Expected UpperThresholdCritical 61, got %v", *temp.UpperThresholdCritical)
	}
}

var serverThermalBody = `{
    "@odata.context": "/redfish/v1/$metadata#Thermal.Thermal",
    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal/",
    "@odata.type": "#Thermal.v1_7_1.Thermal",
    "Description": "The Thermal schema describes temperature monitoring and thermal management subsystems, such as cooling fans, for a computer system or similar devices contained within a chassis.",
    "Fans": [
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/0",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan3",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 3",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/0/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan3",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/0/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 12032,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/0/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13386,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/1",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan1",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 1",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/1/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan1",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/1/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13064,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/1/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 12234,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/2",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan0",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 0",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/2/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan0",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/2/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 11832,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/2/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13341,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/3",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan5",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 5",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/3/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan5",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/3/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 12049,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/3/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13017,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/4",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan4",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 4",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/4/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan4",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/4/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13291,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/4/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 12210,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/5",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan2",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 2",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/5/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan2",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/5/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13503,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/5/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 12147,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/6",
            "LowerThresholdCritical": 0,
            "Manufacturer": "Temp",
            "MaxReadingRange": 100,
            "MemberId": "Sys_Fan6",
            "MinReadingRange": 0,
            "Model": "Server R120 G2 System Fan",
            "Name": "System Fan 6",
            "Oem": {
                "Temp": {
                    "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/6/Oem/Temp",
                    "@odata.type": "#OemThermal.v1_0_0.Fans",
                    "Connector": "Fan6",
                    "Readings": [
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/6/Oem/Temp/Readings/0",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 1500,
                            "Reading": 13022,
                            "UpperThresholdCritical": null
                        },
                        {
                            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Fans/6/Oem/Temp/Readings/1",
                            "LowerThresholdCritical": 500,
                            "MaxReadingRange": 32000,
                            "MinReadingRange": 0,
                            "Reading": 11984,
                            "UpperThresholdCritical": null
                        }
                    ]
                }
            },
            "PartNumber": "MODFAN781001A",
            "Reading": 45,
            "ReadingUnits": "Percent",
            "SerialNumber": "",
            "Status": {
                "Health": "OK",
                "HealthRollup": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 0
        }
    ],
    "Id": "Thermal",
    "Name": "Thermal",
    "Temperatures": [
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/0",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVDQ_ABCD0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVDQ ABCD0 Temp",
            "ReadingCelsius": 27,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/1",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVCCIO0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVCCIO0 Temp",
            "ReadingCelsius": 36,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/2",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVCCIN0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVCCIN0 Temp",
            "ReadingCelsius": 33,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/3",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Die",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Die",
            "ReadingCelsius": 48.157,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 90,
            "UpperThresholdNonCritical": 82
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/4",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "PSU0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "PSU0 Temp",
            "ReadingCelsius": 28,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 80,
            "UpperThresholdNonCritical": 75
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/5",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_9",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 9",
            "ReadingCelsius": 45,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/6",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_P1V8_0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR P1V8 0 Temp",
            "ReadingCelsius": 33,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/7",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_1",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 1",
            "ReadingCelsius": 41,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/8",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "RSR_Temp_1_ext",
            "MinReadingRangeTemp": -128,
            "Name": "RSR Temp 1 ext",
            "ReadingCelsius": null,
            "Status": {
                "Health": "OK",
                "State": "Absent"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/9",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVCCIO1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVCCIO1 Temp",
            "ReadingCelsius": 42,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/10",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_6",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 6",
            "ReadingCelsius": 38,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/11",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_2",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 2",
            "ReadingCelsius": 46,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/12",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_14",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 14",
            "ReadingCelsius": 40,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/13",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_P1V8_1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR P1V8 1 Temp",
            "ReadingCelsius": 42,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/14",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_15",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 15",
            "ReadingCelsius": 39,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/15",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_4",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 4",
            "ReadingCelsius": 41,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/16",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_11",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 11",
            "ReadingCelsius": 39,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/17",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVDQ_ABCD1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVDQ ABCD1 Temp",
            "ReadingCelsius": 29,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/18",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR 0 Temp",
            "ReadingCelsius": 22.187,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 115,
            "UpperThresholdNonCritical": 110
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/19",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_20",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 20",
            "ReadingCelsius": 43,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/20",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR 1 Temp",
            "ReadingCelsius": 23.125,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 115,
            "UpperThresholdNonCritical": 110
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/21",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_20",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 20",
            "ReadingCelsius": 37,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/22",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVCCIN1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVCCIN1 Temp",
            "ReadingCelsius": 34,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/23",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_17",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 17",
            "ReadingCelsius": 39,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/24",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_3",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 3",
            "ReadingCelsius": 38,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/25",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "PSU1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "PSU1 Temp",
            "ReadingCelsius": 32,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 80,
            "UpperThresholdNonCritical": 75
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/26",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "RSR_Temp_1",
            "MinReadingRangeTemp": -128,
            "Name": "RSR Temp 1",
            "ReadingCelsius": 28.625,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/27",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_10",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 10",
            "ReadingCelsius": 39,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/28",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "DIMM_CPU0_A0",
            "MinReadingRangeTemp": -128,
            "Name": "DIMM CPU0 A0",
            "ReadingCelsius": 24,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 83,
            "UpperThresholdNonCritical": 78
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/29",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_1",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 1",
            "ReadingCelsius": 47,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/30",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_25",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 25",
            "ReadingCelsius": 43,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/31",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 127,
            "MemberId": "PCH_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "PCH Temp",
            "ReadingCelsius": 30.25,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 115,
            "UpperThresholdNonCritical": 110
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/32",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_12",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 12",
            "ReadingCelsius": 43,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/33",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "DIMM_CPU1_A0",
            "MinReadingRangeTemp": -128,
            "Name": "DIMM CPU1 A0",
            "ReadingCelsius": 26,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 83,
            "UpperThresholdNonCritical": 78
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/34",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_P3V3_AUX_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR P3V3 AUX Temp",
            "ReadingCelsius": 36,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/35",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_24",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 24",
            "ReadingCelsius": 44,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/36",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_P5V_AUX_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR P5V AUX Temp",
            "ReadingCelsius": 34,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/37",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 127,
            "MemberId": "System_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "System Temp",
            "ReadingCelsius": 31.5,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 115,
            "UpperThresholdNonCritical": 110
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/38",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVDQ_EFGH1_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVDQ EFGH1 Temp",
            "ReadingCelsius": 28,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/39",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_19",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 19",
            "ReadingCelsius": 36,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/40",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Die",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Die",
            "ReadingCelsius": 41.438,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 90,
            "UpperThresholdNonCritical": 82
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/41",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_22",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 22",
            "ReadingCelsius": 45,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/42",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 127,
            "MemberId": "Outlet_Air_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "Outlet Air Temp",
            "ReadingCelsius": 23.5,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 61,
            "UpperThresholdNonCritical": 56
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/43",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_8",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 8",
            "ReadingCelsius": 45,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/44",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "Inlet_Air_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "Inlet Air Temp",
            "ReadingCelsius": 17,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/45",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_19",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 19",
            "ReadingCelsius": 45,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/46",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU0_Core_25",
            "MinReadingRangeTemp": -128,
            "Name": "CPU0 Core 25",
            "ReadingCelsius": 37,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/47",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_23",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 23",
            "ReadingCelsius": 43,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/48",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "CPU1_Core_5",
            "MinReadingRangeTemp": -128,
            "Name": "CPU1 Core 5",
            "ReadingCelsius": 44,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": null,
            "UpperThresholdNonCritical": null
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/49",
            "LowerThresholdCritical": null,
            "LowerThresholdNonCritical": null,
            "MaxReadingRangeTemp": 127,
            "MemberId": "VR_PVDQ_EFGH0_Temp",
            "MinReadingRangeTemp": -128,
            "Name": "VR PVDQ EFGH0 Temp",
            "ReadingCelsius": 28,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 120,
            "UpperThresholdNonCritical": 115
        },
        {
            "@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Thermal#/Temperatures/50",
            "LowerThresholdCritical": 0,
            "LowerThresholdNonCritical": 5,
            "MaxReadingRangeTemp": 255,
            "MemberId": "SSB_Temp",
            "MinReadingRangeTemp": 0,
            "Name": "SSB Temp",
            "ReadingCelsius": 34,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 103,
            "UpperThresholdNonCritical": 98
        }
    ]
}`
