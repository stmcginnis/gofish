//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Memory.Memory",
		"@odata.id": "/redfish/v1/Systems/System-1/Memory/NVRAM4",
		"@odata.type": "#Memory.v1_2_0.Memory",
		"Name": "Memory",
		"Id": "NVRAM4",
		"Links": {
			"Chassis": {
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			}
		},
		"AllocationAlignmentMiB": 1024,
		"AllocationIncrementMiB": 1024,
		"AllowedSpeedsMHz": [3200],
		"Assembly": {
			"@odata.id": "/redfish/v1/Assembly/1"
		},
		"BaseModuleType": "SO_DIMM",
		"BusWidthBits": 1024,
		"CacheSizeMiB": 256,
		"CapacityMiB": 2097152,
		"ConfigurationLocked": false,
		"DataWidthBits": 256,
		"DeviceLocator": "Inside",
		"ErrorCorrection": "SingleBitECC",
		"FirmwareApiVersion": "1.0",
		"FirmwareRevision": "3",
		"IsRankSpareEnabled": false,
		"IsSpareDeviceEnabled": false,
		"LogicalSizeMiB": 2097152,
		"Manufacturer": "Generic",
		"MemoryDeviceType": "DDR4",
		"MemoryLocation": {
			"Channel": 1,
			"MemoryController": 2,
			"Slot": 3,
			"Socket": 4
		},
		"MemoryMedia": ["DRAM"],
		"MemoryType": "NVDIMM_N",
		"OperatingMemoryModes": [
			"PMEM"
		],
		"PowerManagementPolicy": {
			"AveragePowerBudgetMilliWatts": 42,
			"MaxTDPMilliWatts": 12,
			"PeakPowerBudgetMilliWatts": 84,
			"PolicyEnabled": false
		},
		"RankCount": 12,
		"SerialNumber": "TJ27JXQY",
		"SecurityCapabilities": {
			"ConfigurationLockCapable": false,
			"DataLockCapable": false,
			"MaxPassphraseCount": 2,
			"PassphraseCapable": true,
			"PassphraseLockLimit": 2
		},
		"SecurityState": "Unlocked",
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"VendorID": "Generic"
	}`)

// TestMemory tests the parsing of Memory objects.
func TestMemory(t *testing.T) {
	var result Memory
	err := json.NewDecoder(memoryBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NVRAM4" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Memory" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.chassis != "/redfish/v1/Chassis/Chassis-1" {
		t.Errorf("Invalid chassis link: %s", result.chassis)
	}

	if result.MemoryType != NVDIMMNMemoryType {
		t.Errorf("Invalid memory type: %s", result.MemoryType)
	}

	if result.MemoryDeviceType != DDR4MemoryDeviceType {
		t.Errorf("Invalid memory device type: %s", result.MemoryDeviceType)
	}

	if result.OperatingMemoryModes[0] != PMEMOperatingMemoryModes {
		t.Errorf("Invalid operating memory mode results: %s", result.OperatingMemoryModes)
	}

	if result.PowerManagementPolicy.AveragePowerBudgetMilliWatts != 42 {
		t.Errorf("Invalid power management policy average power budget: %d",
			result.PowerManagementPolicy.AveragePowerBudgetMilliWatts)
	}

	if result.SecurityCapabilities.DataLockCapable {
		t.Error("Security capabilities data lock capable should be false")
	}
}
