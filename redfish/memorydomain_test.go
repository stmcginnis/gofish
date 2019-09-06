//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryDomainBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#MemoryDomain.MemoryDomain",
		"@odata.type": "#MemoryDomain.v1_0_0.MemoryDomain",
		"@odata.id": "/redfish/v1/MemoryDomain",
		"Id": "MemoryDomain-1",
		"Name": "MemoryDomainOne",
		"Description": "MemoryDomain One",
		"AllowBlockProvisioning": false,
		"AllowsMemoryChunkCreation": false,
		"AllowsMirroring": true,
		"AllowsSparing": true,
		"InterleavableMemorySets": [{
			"MemorySet": [{
					"@odata.id": "/redfish/v1/System/System-1/Memory/NVRAM1"
				},
				{
					"@odata.id": "/redfish/v1/System/System-1/Memory/NVRAM2"
				}
			],
			"MemorySet@odata.count": 2
		}],
		"MemoryChunks": {
			"@odata.id": "/redfish/v1/System/System-1/Memory/NVRAM1/Chunks"
		}
	}`)

// TestMemoryDomain tests the parsing of MemoryDomain objects.
func TestMemoryDomain(t *testing.T) {
	var result MemoryDomain
	err := json.NewDecoder(memoryDomainBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "MemoryDomain-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "MemoryDomainOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AllowsBlockProvisioning {
		t.Error("Allow block provisioning should be false")
	}

	if result.AllowsMemoryChunkCreation {
		t.Error("Allow memory chunk creation should be false")
	}

	if !result.AllowsMirroring {
		t.Error("Allow mirroring should be true")
	}

	if len(result.InterleavableMemorySets) != 1 {
		t.Errorf("Should have one interleavable memory set, got: %d",
			len(result.InterleavableMemorySets))
	}

	if result.memoryChunks != "/redfish/v1/System/System-1/Memory/NVRAM1/Chunks" {
		t.Errorf("Invalid memory chunk link: %s", result.memoryChunks)
	}
}
