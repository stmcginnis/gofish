//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageControllerMetricsBody = `{
	"@odata.type": "#StorageControllerMetrics.v1_0_2.StorageControllerMetrics",
	"Id": "Metrics",
	"Name": "Storage Controller Metrics for NVMe IO Controller",
	"NVMeSMART": {
	  "CriticalWarnings": {
		"PMRUnreliable": false,
		"PowerBackupFailed": false,
		"MediaInReadOnly": false,
		"OverallSubsystemDegraded": true,
		"SpareCapacityWornOut": false
	  },
	  "CompositeTemperatureCelsius": 34,
	  "AvailableSparePercent": 50,
	  "AvailableSpareThresholdPercent": 30,
	  "PercentageUsed": 50,
	  "EGCriticalWarningSummary": {
		"NamespacesInReadOnlyMode": false,
		"ReliabilityDegraded": false,
		"SpareCapacityUnderThreshold": false
	  },
	  "DataUnitsRead": 0,
	  "DataUnitsWritten": 0,
	  "HostReadCommands": 0,
	  "HostWriteCommands": 0,
	  "ControllerBusyTimeMinutes": 20,
	  "PowerCycles": 49,
	  "PowerOnHours": 3,
	  "UnsafeShutdowns": 4,
	  "MediaAndDataIntegrityErrors": 0,
	  "NumberOfErrorInformationLogEntries": 100,
	  "WarningCompositeTempTimeMinutes": 0,
	  "CriticalCompositeTempTimeMinutes": 0,
	  "TemperatureSensorsCelsius": [
		34,
		34,
		34,
		26,
		31,
		35,
		33,
		32
	  ],
	  "ThermalMgmtTemp1TransitionCount": 10,
	  "ThermalMgmtTemp2TransitionCount": 2,
	  "ThermalMgmtTemp1TotalTimeSeconds": 20,
	  "ThermalMgmtTemp2TotalTimeSeconds": 42
	},
	"@odata.id": "/redfish/v1/Systems/Sys-1/Storage/SimplestNVMeSSD/Controllers/NVMeIOController/Metrics"
  }`

// TestStorageControllerMetrics tests the parsing of StorageControllerMetrics objects.
func TestStorageControllerMetrics(t *testing.T) {
	var result StorageControllerMetrics
	err := json.NewDecoder(strings.NewReader(storageControllerMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Storage Controller Metrics for NVMe IO Controller", result.Name)

	if !result.NVMeSMART.CriticalWarnings.OverallSubsystemDegraded {
		t.Error("Expected NVMeSMART.CriticalWarnings.OverallSubsystemDegraded to be true")
	}

	if result.NVMeSMART.PercentageUsed != 50 {
		t.Errorf("Unexpected NVMeSMART.PercentageUsed value: %.2f", result.NVMeSMART.PercentageUsed)
	}

	if result.NVMeSMART.PowerCycles != 49 {
		t.Errorf("Unexpected NVMeSMART.PowerCycles value: %d", result.NVMeSMART.PowerCycles)
	}
}
