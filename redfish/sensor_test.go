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

var sensorBody = strings.NewReader(
	`{
		"@odata.type": "#Sensor.v1_7_0.Sensor",
		"Id": "CabinetTemp",
		"Name": "Rack Temperature",
		"ReadingType": "Temperature",
		"ReadingTime": "2019-12-25T04:14:33+06:00",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Reading": 31.6,
		"ReadingUnits": "C",
		"ReadingRangeMin": -1.7976931348623157e+308,
		"ReadingRangeMax": 1.7976931348623157e+308,
		"Accuracy": 0.25,
		"Precision": 1,
		"SensingInterval": "PT3S",
		"PhysicalContext": "Chassis",
		"Thresholds": {
			"UpperCritical": {
				"Reading": 40,
				"Activation": "Increasing"
			},
			"UpperCaution": {
				"Reading": 35,
				"Activation": "Decreasing"
			},
			"LowerCaution": {
				"Reading": 10,
				"Activation": "Increasing"
			}
		},
		"@odata.id": "/redfish/v1/Chassis/1/Sensors/CabinetTemp"
	}`)

// TestSensor tests the parsing of Sensor objects.
func TestSensor(t *testing.T) {
	var result Sensor

	if err := json.NewDecoder(sensorBody).Decode(&result); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "CabinetTemp", result.ID)
	assertEquals(t, "Rack Temperature", result.Name)
	assertEquals(t, "Temperature", string(result.ReadingType))
	assertEquals(t, "2019-12-25T04:14:33+06:00", result.ReadingTime)
	assertEquals(t, "31.6", fmt.Sprint(*result.Reading))
	assertEquals(t, "C", result.ReadingUnits)
	assertEquals(t, "Decreasing", string(result.Thresholds.UpperCaution.Activation))
}

func TestSensorUnmarshalJSON(t *testing.T) {
	testData := `{
		"@odata.context": "/redfish/v1/$metadata#Sensor.Sensor",
		"@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Sensors/SSB_Temp/",
		"@odata.type": "#Sensor.v1_5_0.Sensor",
		"Description": "The Sensor schema describes a sensor and its properties.",
		"Id": "SSB_Temp",
		"Name": "SSB Temp",
		"Reading": 23,
		"ReadingRangeMax": 255,
		"ReadingRangeMin": 0,
		"ReadingType": "Temperature",
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"Thresholds": {
			"LowerCaution": {
				"Reading": 5
			},
			"LowerCritical": {
				"Reading": 0
			},
			"UpperCaution": {
				"Reading": 98
			},
			"UpperCritical": {
				"Reading": 103
			}
		}
	}`

	var sensor Sensor
	err := json.Unmarshal([]byte(testData), &sensor)
	if err != nil {
		t.Fatalf("Failed to unmarshal sensor: %v", err)
	}

	t.Run("Check basic fields", func(t *testing.T) {
		if sensor.ODataContext != "/redfish/v1/$metadata#Sensor.Sensor" {
			t.Errorf("Unexpected ODataContext: %s", sensor.ODataContext)
		}
		if sensor.ODataID != "/redfish/v1/Chassis/Server_R120_G2_Server/Sensors/SSB_Temp/" {
			t.Errorf("Unexpected ODataID: %s", sensor.ODataID)
		}
		if sensor.ODataType != "#Sensor.v1_5_0.Sensor" {
			t.Errorf("Unexpected ODataType: %s", sensor.ODataType)
		}
		if sensor.ID != "SSB_Temp" {
			t.Errorf("Unexpected ID: %s", sensor.ID)
		}
		if sensor.Name != "SSB Temp" {
			t.Errorf("Unexpected Name: %s", sensor.Name)
		}
	})

	t.Run("Check reading values", func(t *testing.T) {
		if *sensor.Reading != 23 {
			t.Errorf("Unexpected Reading value: %v", *sensor.Reading)
		}
		if *sensor.ReadingRangeMax != 255 {
			t.Errorf("Unexpected ReadingRangeMax: %v", *sensor.ReadingRangeMax)
		}
		if *sensor.ReadingRangeMin != 0 {
			t.Errorf("Unexpected ReadingRangeMin: %v", *sensor.ReadingRangeMin)
		}
		if sensor.ReadingType != ReadingTypeTemperature {
			t.Errorf("Unexpected ReadingType: %s", sensor.ReadingType)
		}
	})

	t.Run("Check status", func(t *testing.T) {
		if sensor.Status.Health != "OK" {
			t.Errorf("Unexpected Health status: %s", sensor.Status.Health)
		}
		if sensor.Status.State != "Enabled" {
			t.Errorf("Unexpected State: %s", sensor.Status.State)
		}
	})

	t.Run("Check thresholds", func(t *testing.T) {
		if sensor.Thresholds.LowerCaution == nil || *sensor.Thresholds.LowerCaution.Reading != 5 {
			t.Error("LowerCaution threshold not parsed correctly")
		}
		if sensor.Thresholds.LowerCritical == nil || *sensor.Thresholds.LowerCritical.Reading != 0 {
			t.Error("LowerCritical threshold not parsed correctly")
		}
		if sensor.Thresholds.UpperCaution == nil || *sensor.Thresholds.UpperCaution.Reading != 98 {
			t.Error("UpperCaution threshold not parsed correctly")
		}
		if sensor.Thresholds.UpperCritical == nil || *sensor.Thresholds.UpperCritical.Reading != 103 {
			t.Error("UpperCritical threshold not parsed correctly")
		}
	})
}

func TestSensorThresholds(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		check   func(*Sensor) bool
	}{
		{
			name: "All thresholds",
			json: `{"Thresholds": {
				"LowerCaution": {"Reading": 5},
				"LowerCritical": {"Reading": 0},
				"UpperCaution": {"Reading": 98},
				"UpperCritical": {"Reading": 103}
			}}`,
			wantErr: false,
			check: func(s *Sensor) bool {
				return s.Thresholds.LowerCaution != nil &&
					s.Thresholds.LowerCritical != nil &&
					s.Thresholds.UpperCaution != nil &&
					s.Thresholds.UpperCritical != nil
			},
		},
		{
			name:    "No thresholds",
			json:    `{}`,
			wantErr: false,
			check: func(s *Sensor) bool {
				return s.Thresholds.LowerCaution == nil &&
					s.Thresholds.LowerCritical == nil &&
					s.Thresholds.UpperCaution == nil &&
					s.Thresholds.UpperCritical == nil
			},
		},
		{
			name: "Partial thresholds",
			json: `{"Thresholds": {
				"UpperCaution": {"Reading": 98},
				"UpperCritical": {"Reading": 103}
			}}`,
			wantErr: false,
			check: func(s *Sensor) bool {
				return s.Thresholds.LowerCaution == nil &&
					s.Thresholds.LowerCritical == nil &&
					s.Thresholds.UpperCaution != nil &&
					s.Thresholds.UpperCritical != nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Sensor
			err := json.Unmarshal([]byte(tt.json), &s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sensor.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.check(&s) {
				t.Error("Thresholds check failed")
			}
		})
	}
}

func TestSensorReadingTypes(t *testing.T) {
	tests := []struct {
		json      string
		wantType  ReadingType
		wantValue float64
	}{
		{
			json:      `{"ReadingType": "Temperature", "Reading": 23.5}`,
			wantType:  ReadingTypeTemperature,
			wantValue: 23.5,
		},
		{
			json:      `{"ReadingType": "Voltage", "Reading": 12.2}`,
			wantType:  ReadingTypeVoltage,
			wantValue: 12.2,
		},
		{
			json:      `{"ReadingType": "Power", "Reading": 150.75}`,
			wantType:  ReadingTypePower,
			wantValue: 150.75,
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.wantType), func(t *testing.T) {
			var s Sensor
			if err := json.Unmarshal([]byte(tt.json), &s); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if s.ReadingType != tt.wantType {
				t.Errorf("Got ReadingType %s, want %s", s.ReadingType, tt.wantType)
			}

			if *s.Reading != tt.wantValue {
				t.Errorf("Got Reading %v, want %v", *s.Reading, tt.wantValue)
			}
		})
	}
}

func TestSensorStatus(t *testing.T) {
	testData := `{
		"Status": {
			"Health": "Warning",
			"HealthRollup": "OK",
			"State": "StandbyOffline"
		}
	}`

	var s Sensor
	if err := json.Unmarshal([]byte(testData), &s); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if s.Status.Health != "Warning" {
		t.Errorf("Unexpected Health: %s", s.Status.Health)
	}
	if s.Status.HealthRollup != "OK" {
		t.Errorf("Unexpected HealthRollup: %s", s.Status.HealthRollup)
	}
	if s.Status.State != "StandbyOffline" {
		t.Errorf("Unexpected State: %s", s.Status.State)
	}
}

func TestPowerSensorUnmarshalJSON(t *testing.T) {
	testData := `{
		"@odata.context": "/redfish/v1/$metadata#Sensor.Sensor",
		"@odata.id": "/redfish/v1/Chassis/Server_R120_G2_Server/Sensors/PWR_AUX_PWR1/",
		"@odata.type": "#Sensor.v1_5_0.Sensor",
		"Description": "The Sensor schema describes a sensor and its properties.",
		"Id": "PWR_AUX_PWR1",
		"Name": "PWR AUX PWR1",
		"Reading": 0,
		"ReadingRangeMax": 400,
		"ReadingRangeMin": 0,
		"ReadingType": "Power",
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"Thresholds": {
			"UpperCaution": {
				"Reading": 300
			},
			"UpperCritical": {
				"Reading": 300
			}
		}
	}`

	var sensor Sensor
	err := json.Unmarshal([]byte(testData), &sensor)
	if err != nil {
		t.Fatalf("Failed to unmarshal power sensor: %v", err)
	}

	t.Run("Check power sensor basic fields", func(t *testing.T) {
		if sensor.ODataID != "/redfish/v1/Chassis/Server_R120_G2_Server/Sensors/PWR_AUX_PWR1/" {
			t.Errorf("Unexpected ODataID: %s", sensor.ODataID)
		}
		if sensor.ID != "PWR_AUX_PWR1" {
			t.Errorf("Unexpected ID: %s", sensor.ID)
		}
		if sensor.Name != "PWR AUX PWR1" {
			t.Errorf("Unexpected Name: %s", sensor.Name)
		}
	})

	t.Run("Check power sensor specific values", func(t *testing.T) {
		if *sensor.Reading != 0 {
			t.Errorf("Expected zero reading, got: %v", *sensor.Reading)
		}
		if *sensor.ReadingRangeMax != 400 {
			t.Errorf("Unexpected max range: %v", *sensor.ReadingRangeMax)
		}
		if *sensor.ReadingRangeMin != 0 {
			t.Errorf("Unexpected min range: %v", *sensor.ReadingRangeMin)
		}
		if sensor.ReadingType != ReadingTypePower {
			t.Errorf("Expected Power reading type, got: %s", sensor.ReadingType)
		}
	})

	t.Run("Check power sensor thresholds", func(t *testing.T) {
		if sensor.Thresholds.UpperCaution == nil || *sensor.Thresholds.UpperCaution.Reading != 300 {
			t.Error("UpperCaution threshold not parsed correctly")
		}
		if sensor.Thresholds.UpperCritical == nil || *sensor.Thresholds.UpperCritical.Reading != 300 {
			t.Error("UpperCritical threshold not parsed correctly")
		}
		// Verify lower thresholds are nil for power sensor
		if sensor.Thresholds.LowerCaution != nil {
			t.Error("LowerCaution threshold should be nil for power sensor")
		}
		if sensor.Thresholds.LowerCritical != nil {
			t.Error("LowerCritical threshold should be nil for power sensor")
		}
	})

	t.Run("Check power sensor status", func(t *testing.T) {
		if sensor.Status.Health != "OK" {
			t.Errorf("Expected health OK, got: %s", sensor.Status.Health)
		}
		if sensor.Status.State != "Enabled" {
			t.Errorf("Expected state Enabled, got: %s", sensor.Status.State)
		}
	})

	t.Run("Check power sensor additional fields", func(t *testing.T) {
		if sensor.ApparentVA != nil {
			t.Error("ApparentVA should be nil when not in payload")
		}
		if sensor.PowerFactor != nil {
			t.Error("PowerFactor should be nil when not in payload")
		}
		if sensor.ReactiveVAR != nil {
			t.Error("ReactiveVAR should be nil when not in payload")
		}
	})
}

func TestPowerSensorWithElectricalProperties(t *testing.T) {
	testData := `{
		"@odata.type": "#Sensor.v1_5_0.Sensor",
		"ReadingType": "Power",
		"Reading": 125.5,
		"ApparentVA": 150.2,
		"PowerFactor": 0.85,
		"ReactiveVAR": 75.3,
		"PhaseAngleDegrees": 25.1
	}`

	var sensor Sensor
	err := json.Unmarshal([]byte(testData), &sensor)
	if err != nil {
		t.Fatalf("Failed to unmarshal power sensor with electrical properties: %v", err)
	}

	t.Run("Check electrical properties", func(t *testing.T) {
		if *sensor.ApparentVA != 150.2 {
			t.Errorf("Unexpected ApparentVA: %v", *sensor.ApparentVA)
		}
		if *sensor.PowerFactor != 0.85 {
			t.Errorf("Unexpected PowerFactor: %v", *sensor.PowerFactor)
		}
		if *sensor.ReactiveVAR != 75.3 {
			t.Errorf("Unexpected ReactiveVAR: %v", *sensor.ReactiveVAR)
		}
		if *sensor.PhaseAngleDegrees != 25.1 {
			t.Errorf("Unexpected PhaseAngleDegrees: %v", *sensor.PhaseAngleDegrees)
		}
	})
}

func TestPowerSensorEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		check   func(*testing.T, *Sensor)
	}{
		{
			name: "Null reading",
			json: `{"ReadingType": "Power", "Reading": null}`,
			check: func(t *testing.T, s *Sensor) {
				if s.Reading != nil {
					t.Error("Reading should be nil for null value")
				}
			},
		},
		{
			name: "Missing thresholds",
			json: `{"ReadingType": "Power"}`,
			check: func(t *testing.T, s *Sensor) {
				if s.Thresholds.UpperCaution != nil {
					t.Error("UpperCaution should be nil when missing")
				}
			},
		},
		{
			name:    "Invalid power value",
			json:    `{"ReadingType": "Power", "Reading": -1}`,
			wantErr: false, // Should accept negative values for power
			check: func(t *testing.T, s *Sensor) {
				if *s.Reading != -1 {
					t.Errorf("Should accept negative power values, got: %v", *s.Reading)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Sensor
			err := json.Unmarshal([]byte(tt.json), &s)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.check != nil {
				tt.check(t, &s)
			}
		})
	}
}
