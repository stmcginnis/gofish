//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

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
