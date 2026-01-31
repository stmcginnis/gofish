//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var managerDiagnosticDataBody = `{
	"@odata.type": "#ManagerDiagnosticData.v1_2_2.ManagerDiagnosticData",
	"Id": "ManagerDiagnosticData",
	"Name": "Manager Diagnostic Data",
	"I2CBuses": [
	  {
		"I2CBusName": "i2c-0",
		"TotalTransactionCount": 10000,
		"BusErrorCount": 12,
		"NACKCount": 34
	  },
	  {
		"I2CBusName": "i2c-1",
		"TotalTransactionCount": 20000,
		"BusErrorCount": 56,
		"NACKCount": 78
	  }
	],
	"MemoryStatistics": {
	  "TotalBytes": 1013052000,
	  "UsedBytes": 45084000,
	  "FreeBytes": 894820000,
	  "SharedBytes": 19864000,
	  "BuffersAndCacheBytes": 73148000,
	  "AvailableBytes": 928248000
	},
	"ProcessorStatistics": {
	  "KernelPercent": 12.34,
	  "UserPercent": 23.45
	},
	"TopProcesses": [
	  {
		"CommandLine": "dbus-broker",
		"UserTimeSeconds": 14400,
		"KernelTimeSeconds": 10800,
		"ResidentSetSizeBytes": 2300000
	  },
	  {
		"CommandLine": "swampd",
		"UserTimeSeconds": 13200,
		"KernelTimeSeconds": 8441,
		"ResidentSetSizeBytes": 8883000
	  },
	  {
		"CommandLine": "ipmid",
		"UserTimeSeconds": 13100,
		"KernelTimeSeconds": 6650,
		"ResidentSetSizeBytes": 23400000
	  },
	  {
		"CommandLine": "phosphor-hwmon-readd -i iface1",
		"UserTimeSeconds": 5100,
		"KernelTimeSeconds": 3200,
		"ResidentSetSizeBytes": 564000
	  }
	],
	"BootTimeStatistics": {
	  "FirmwareTimeSeconds": 42.3,
	  "LoaderTimeSeconds": 12.3,
	  "KernelTimeSeconds": 33.1,
	  "InitrdTimeSeconds": 3.2,
	  "UserSpaceTimeSeconds": 81.1
	},
	"MemoryECCStatistics": {
	  "CorrectableECCErrorCount": 1,
	  "UncorrectableECCErrorCount": 2
	},
	"@odata.id": "/redfish/v1/Managers/BMC/ManagerDiagnosticData"
  }`

// TestManagerDiagnosticData tests the parsing of ManagerDiagnosticData objects.
func TestManagerDiagnosticData(t *testing.T) {
	var result ManagerDiagnosticData
	err := json.NewDecoder(strings.NewReader(managerDiagnosticDataBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ManagerDiagnosticData", result.ID)
	assertEquals(t, "Manager Diagnostic Data", result.Name)
	assertEquals(t, "i2c-0", result.I2CBuses[0].I2CBusName)

	if *result.MemoryStatistics.TotalBytes != 1013052000 {
		t.Errorf("Unexpected memory stats total bytes: %d", result.MemoryStatistics.TotalBytes)
	}

	if *result.ProcessorStatistics.KernelPercent != 12.34 {
		t.Errorf("Unexpected processor stats kernel percent: %.2f", *result.ProcessorStatistics.KernelPercent)
	}

	if *result.BootTimeStatistics.LoaderTimeSeconds != 12.3 {
		t.Errorf("Unexpected memory stats total bytes: %.2f", *result.BootTimeStatistics.LoaderTimeSeconds)
	}

	if *result.MemoryECCStatistics.CorrectableECCErrorCount != 1 {
		t.Errorf("Unexpected memory ECC stats correctable error count: %d", result.MemoryECCStatistics.CorrectableECCErrorCount)
	}
}
