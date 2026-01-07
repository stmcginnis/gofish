//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryRegionBody = `{
	"@odata.type": "#MemoryRegion.v1_0_1.MemoryRegion",
	"Id": "1",
	"Name": "Dynamic Memory Region 1",
	"Description": "CXL Dynamic Memory Region 1 of LD 1 in Device 1",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK",
	  "HealthRollup": "OK"
	},
	"RegionType": "Dynamic",
	"RegionNumber": 0,
	"RegionBaseOffsetMiB": 0,
	"RegionSizeMiB": 65536,
	"ShareableRegion": false,
	"SanitizeOnRelease": true,
	"BlockSizeMiB": 128,
	"ExtentsCount": 1,
	"MemoryExtents": [
	  {
		"ExtentOffsetMiB": 0,
		"ExtentSizeMiB": 4096,
		"Tag": "User Defined Tag",
		"SequenceNumber": 0
	  }
	],
	"MemoryChunks": [
	  {
		"ChunkOffsetMiB": 0,
		"ChunkLink": {
		  "@odata.id": "/redfish/v1/Chassis/1/MemoryDomains/1/MemoryChunks/1"
		}
	  }
	],
	"@odata.id": "/redfish/v1/Chassis/1/PCIeDevices/1/CXLLogicalDevices/1/MemoryRegions/1"
  }`

// TestMemoryRegion tests the parsing of MemoryRegion objects.
func TestMemoryRegion(t *testing.T) {
	var result MemoryRegion
	err := json.NewDecoder(strings.NewReader(memoryRegionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Dynamic Memory Region 1", result.Name)
	assertEquals(t, "Dynamic", string(result.RegionType))
	assertEquals(t, "User Defined Tag", result.MemoryExtents[0].Tag)
	assertEquals(t, "/redfish/v1/Chassis/1/MemoryDomains/1/MemoryChunks/1", result.MemoryChunks[0].chunkLink)

	if result.RegionNumber != 0 {
		t.Errorf("Unexpected region number: %d", result.RegionNumber)
	}

	if result.RegionBaseOffsetMiB != 0 {
		t.Errorf("Unexpected RegionBaseOffsetMiB: %d", result.RegionBaseOffsetMiB)
	}

	if result.RegionSizeMiB != 65536 {
		t.Errorf("Unexpected RegionSizeMiB: %d", result.RegionSizeMiB)
	}

	if !result.SanitizeOnRelease {
		t.Error("Expected SanitizeOnRelease to be enabled")
	}
}
