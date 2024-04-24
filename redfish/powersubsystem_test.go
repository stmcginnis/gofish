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

// TestPowerSubsystem tests the parsing of PowerSubsystem objects.
func TestPowerSubsystem(t *testing.T) {
	var result PowerSubsystem
	err := json.NewDecoder(strings.NewReader(powerSubsystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PowerSubsystem", result.ID)
	assertEquals(t, "Power Subsystem for Chassis", result.Name)
	assertEquals(t, "2000", fmt.Sprintf("%.0f", result.CapacityWatts))
	assertEquals(t, "1500", fmt.Sprintf("%.0f", result.Allocation.RequestedWatts))
	assertEquals(t, "1200", fmt.Sprintf("%.0f", result.Allocation.AllocatedWatts))
	assertEquals(t, "Failover", string(result.PowerSupplyRedundancy[0].RedundancyType))
	assertEquals(t, "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay2", result.PowerSupplyRedundancy[0].redundancyGroup[1])
	assertEquals(t, "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies", result.powerSupplies)
}
