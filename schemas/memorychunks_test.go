//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryChunksBody = `{
	"@odata.type": "#MemoryChunks.v1_6_1.MemoryChunks",
	"Name": "Memory Chunk - Whole System",
	"Id": "1",
	"MemoryChunkSizeMiB": 32768,
	"AddressRangeType": "Volatile",
	"IsMirrorEnabled": false,
	"IsSpare": false,
	"InterleaveSets": [
	  {
		"Memory": {
		  "@odata.id": "/redfish/v1/Systems/2/Memory/1"
		}
	  },
	  {
		"Memory": {
		  "@odata.id": "/redfish/v1/Systems/2/Memory/2"
		}
	  },
	  {
		"Memory": {
		  "@odata.id": "/redfish/v1/Systems/2/Memory/3"
		}
	  },
	  {
		"Memory": {
		  "@odata.id": "/redfish/v1/Systems/2/Memory/4"
		}
	  }
	],
	"@Redfish.Settings": {
	  "@odata.type": "#Settings.v1_3_5.Settings",
	  "SettingsObject": {
		"@odata.id": "/redfish/v1/Systems/2/MemoryDomains/1/MemoryChunks/1/SD"
	  },
	  "Time": "2012-03-07T14:44.30-05:00",
	  "ETag": "someetag",
	  "Messages": [
		{
		  "MessageId": "Base.1.0.Success"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Systems/2/MemoryDomains/1/MemoryChunks/1"
  }`

// TestMemoryChunks tests the parsing of MemoryChunks objects.
func TestMemoryChunks(t *testing.T) {
	var result MemoryChunks
	err := json.NewDecoder(strings.NewReader(memoryChunksBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Memory Chunk - Whole System", result.Name)
	assertEquals(t, "Volatile", string(result.AddressRangeType))
	assertEquals(t, "/redfish/v1/Systems/2/Memory/4", result.InterleaveSets[3].memory)

	if result.IsSpare {
		t.Error("Expected not to be a spare")
	}

	if result.IsMirrorEnabled {
		t.Error("Expected mirror not to be enabled")
	}

	if *result.MemoryChunkSizeMiB != 32768 {
		t.Errorf("Unexpected memory chunk size: %d", result.MemoryChunkSizeMiB)
	}
}
