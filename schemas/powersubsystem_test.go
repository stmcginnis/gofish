//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

var powerSubsystemBody = `{
	"@odata.type": "#PowerSubsystem.v1_1_1.PowerSubsystem",
	"Id": "PowerSubsystem",
	"Name": "Power Subsystem for Chassis",
	"CapacityWatts": 2000,
	"Allocation": {
	  "RequestedWatts": 1500,
	  "AllocatedWatts": 1200
	},
	"PowerSupplyRedundancy": [
	  {
		"RedundancyType": "Failover",
		"MaxSupportedInGroup": 2,
		"MinNeededInGroup": 1,
		"RedundancyGroup": [
		  {
			"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay2"
		  }
		],
		"Status": {
		  "State": "UnavailableOffline",
		  "Health": "OK"
		}
	  }
	],
	"PowerSupplies": {
	  "@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies"
	},
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem"
  }`

var powerSubsystemNVBody = `{
	"@odata.context": "/redfish/v1/$metadata#PowerSubsystem.PowerSubsystem",
	"@odata.etag": "\"1715094363\"",
	"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem",
	"@odata.type": "#PowerSubsystem.v1_1_0.PowerSubsystem",
	"Allocation": [],
	"Description": "PowerSubsytem for this Chassis",
	"Id": "PowerSubsytem",
	"Name": "PowerSubsytem",
	"PowerSupplies": {
	  "@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies"
	},
	"PowerSupplyRedundancy": [
	  {
		"MaxSupportedInGroup": 6,
		"MinNeededInGroup": 1,
		"RedundancyGroup": [
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU0"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU1"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU2"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU3"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU4"
		  },
		  {
			"@odata.id": "/redfish/v1/Chassis/DGX/PowerSubsystem/PowerSupplies/PSU5"
		  }
		],
		"RedundancyGroup@odata.count": 6,
		"RedundancyType": "NPlusM",
		"Status": {
		  "Health": "OK",
		  "State": "Enabled"
		}
	  }
	],
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	}
  }`

// TestPowerSubsystem tests the parsing of PowerSubsystem objects.
func TestPowerSubsystem(t *testing.T) {
	var result PowerSubsystem
	err := json.NewDecoder(strings.NewReader(powerSubsystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PowerSubsystem", result.ID)
	assertEquals(t, "Power Subsystem for Chassis", result.Name)
	assertEquals(t, "2000", fmt.Sprintf("%.0f", *result.CapacityWatts))
	assertEquals(t, "1500", fmt.Sprintf("%.0f", *result.Allocation.RequestedWatts))
	assertEquals(t, "1200", fmt.Sprintf("%.0f", *result.Allocation.AllocatedWatts))
	assertEquals(t, "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies", result.powerSupplies)
}

// TestPowerSubsystemNVWorkaround tests the workaround for a non-spec implementation of PowerSubsystem objects.
func TestPowerSubsystemNVWorkaround(t *testing.T) {
	var result PowerSubsystem
	err := json.NewDecoder(strings.NewReader(powerSubsystemNVBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
}

// Some NVIDIA services send non-numeric Allocation watt values; those decode as
// 0 rather than failing the whole resource.
func TestPowerSubsystemNonNumericAllocation(t *testing.T) {
	var result PowerSubsystem
	body := `{
		"@odata.id": "/redfish/v1/Chassis/1/PowerSubsystem",
		"Id": "PowerSubsystem",
		"Allocation": {"AllocatedWatts": "N/A", "RequestedWatts": 1200}
	}`
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "0", fmt.Sprintf("%.0f", *result.Allocation.AllocatedWatts))
	assertEquals(t, "1200", fmt.Sprintf("%.0f", *result.Allocation.RequestedWatts))
}

// TestPowerSubsystemPowerSupplies tests that the PowerSupplies link resolves
// to PowerSupplyUnit resources (PowerSupplyCollection), not to the legacy
// Power.PowerSupply type, which has no Metrics link.
func TestPowerSubsystemPowerSupplies(t *testing.T) {
	const collection = `{
		"Members": [{"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1"}]
	}`
	const supply = `{
		"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1",
		"Id": "Bay1",
		"PowerSupplyType": "AC",
		"PowerCapacityWatts": 400,
		"Metrics": {"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Metrics"}
	}`
	const metricsBody = `{
		"InputPowerWatts": {"Reading": 937.4},
		"OutputPowerWatts": {"Reading": 905.3},
		"TemperatureCelsius": {"Reading": 43.9}
	}`
	var result PowerSubsystem
	err := json.NewDecoder(strings.NewReader(powerSubsystemBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				getCall(collection),
				getCall(supply),
				getCall(metricsBody),
			},
		},
	}
	result.SetClient(testClient)

	supplies, err := result.PowerSupplies()
	if err != nil {
		t.Fatalf("Error getting power supplies: %s", err)
	}

	if len(supplies) != 1 {
		t.Fatalf("Expected 1 power supply, got %d", len(supplies))
	}

	psu := supplies[0]
	if psu.ID != "Bay1" {
		t.Errorf("Received invalid power supply ID: %s", psu.ID)
	}

	if psu.PowerSupplyType != ACPowerSupplyUnitType {
		t.Errorf("Received invalid power supply type: %s", psu.PowerSupplyType)
	}

	if psu.PowerCapacityWatts == nil || *psu.PowerCapacityWatts != 400 {
		t.Errorf("Received invalid power capacity: %v", psu.PowerCapacityWatts)
	}

	metrics, err := psu.Metrics()
	if err != nil {
		t.Fatalf("Error getting power supply metrics: %s", err)
	}

	if metrics.InputPowerWatts.Reading == nil || *metrics.InputPowerWatts.Reading != 937.4 {
		t.Errorf("Received invalid input power reading: %v", metrics.InputPowerWatts.Reading)
	}

	if metrics.OutputPowerWatts.Reading == nil || *metrics.OutputPowerWatts.Reading != 905.3 {
		t.Errorf("Received invalid output power reading: %v", metrics.OutputPowerWatts.Reading)
	}

	if metrics.TemperatureCelsius.Reading == nil || *metrics.TemperatureCelsius.Reading != 43.9 {
		t.Errorf("Received invalid temperature reading: %v", metrics.TemperatureCelsius.Reading)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 3 ||
		calls[0].URL != "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies" ||
		calls[1].URL != "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1" ||
		calls[2].URL != "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Metrics" {
		t.Errorf("Unexpected calls: %+v", calls)
	}
}

// TestPowerSubsystemNoPowerSupplies tests that a subsystem without a
// PowerSupplies link returns no supplies and makes no request.
func TestPowerSubsystemNoPowerSupplies(t *testing.T) {
	var result PowerSubsystem
	err := json.NewDecoder(strings.NewReader(`{"Id": "PowerSubsystem", "Name": "Power Subsystem"}`)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	supplies, err := result.PowerSupplies()
	if err != nil {
		t.Fatalf("Error getting power supplies: %s", err)
	}

	if supplies != nil {
		t.Errorf("Expected no power supplies, got %d", len(supplies))
	}

	if calls := testClient.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no requests, got %+v", calls)
	}
}
