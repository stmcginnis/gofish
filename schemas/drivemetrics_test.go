//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var driveMetricsBody = `{
	"@odata.type": "#DriveMetrics.v1_2_0.DriveMetrics",
	"Id": "Metrics",
	"Name": "Drive Metrics",
	"CorrectableIOReadErrorCount": 184,
	"UncorrectableIOReadErrorCount": 0,
	"CorrectableIOWriteErrorCount": 18,
	"UncorrectableIOWriteErrorCount": 0,
	"BadBlockCount": 123098,
	"PowerOnHours": 3,
	"NVMeSMART": {
	  "CriticalWarnings": {
		"PMRUnreliable": false,
		"PowerBackupFailed": false,
		"MediaInReadOnly": false,
		"OverallSubsystemDegraded": false,
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
		34,
		34,
		35,
		33,
		32
	  ],
	  "ThermalMgmtTemp1TransitionCount": 10,
	  "ThermalMgmtTemp2TransitionCount": 2,
	  "ThermalMgmtTemp1TotalTimeSeconds": 20,
	  "ThermalMgmtTemp2TotalTimeSeconds": 42
	},
	"@odata.id": "/redfish/v1/Chassis/StorageEnclosure1/Drives/0THGR0KP/Metrics"
  }`

// TestDriveMetrics tests the parsing of DriveMetrics objects.
func TestDriveMetrics(t *testing.T) {
	var result DriveMetrics
	err := json.NewDecoder(strings.NewReader(driveMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Drive Metrics", result.Name)

	if *result.CorrectableIOReadErrorCount != 184 {
		t.Errorf("Unexpected CorrectableIOReadErrorCount, %d", result.CorrectableIOReadErrorCount)
	}

	if *result.UncorrectableIOReadErrorCount != 0 {
		t.Errorf("Unexpected UncorrectableIOReadErrorCount, %d", result.UncorrectableIOReadErrorCount)
	}

	if *result.CorrectableIOWriteErrorCount != 18 {
		t.Errorf("Unexpected CorrectableIOWriteErrorCount, %d", result.CorrectableIOWriteErrorCount)
	}

	if *result.UncorrectableIOWriteErrorCount != 0 {
		t.Errorf("Unexpected UncorrectableIOWriteErrorCount, %d", result.UncorrectableIOWriteErrorCount)
	}

	if *result.BadBlockCount != 123098 {
		t.Errorf("Unexpected BadBlockCount, %d", result.BadBlockCount)
	}

	if *result.PowerOnHours != 3 {
		t.Errorf("Unexpected PowerOnHours, %.2f", *result.PowerOnHours)
	}

	if result.NVMeSMART.CriticalWarnings.PMRUnreliable {
		t.Errorf("Unexpected NVMeSMART.CriticalWarnings.PMRUnreliable, %t", result.NVMeSMART.CriticalWarnings.PMRUnreliable)
	}

	if result.NVMeSMART.EGCriticalWarningSummary.ReliabilityDegraded {
		t.Errorf("Unexpected NVMeSMART.EGCriticalWarningSummary.ReliabilityDegraded, %t", result.NVMeSMART.EGCriticalWarningSummary.ReliabilityDegraded)
	}

	if *result.NVMeSMART.CompositeTemperatureCelsius != 34 {
		t.Errorf("Unexpected NVMeSMART.CompositeTemperatureCelsius, %.2f", *result.NVMeSMART.CompositeTemperatureCelsius)
	}
}
