// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#StorageService.StorageService",
		"@odata.type": "#StorageService.v1_2_0.StorageService",
		"@odata.id": "/redfish/v1/StorageService",
		"Id": "StorageService-1",
		"Name": "StorageServiceOne",
		"Description": "StorageService One",
		"ClassesOfService": {
			"@odata.id": "/redfish/v1/ClassesOfService"
		},
		"DataProtectionLoSCapabilities": {
			"@odata.id": "/redfish/v1/DataProtectionLoSCapabilities/1"
		},
		"DataSecurityLoSCapabilities": {
			"@odata.id": "/redfish/v1/DataSecurityLoSCapabilities/1"
		},
		"DefaultClassOfService": {
			"@odata.id": "/redfish/v1/ClassesOfService/1"
		},
		"Drives": {
			"@odata.id": "/redfish/v1/Drives"
		},
		"EndpointGroups": [{
				"@odata.id": "/redfish/v1/Endpoints/1"
			},
			{
				"@odata.id": "/redfish/v1/Endpoints/2"
			}
		],
		"Endpoints": [{
				"@odata.id": "/redfish/v1/Endpoints/1"
			},
			{
				"@odata.id": "/redfish/v1/Endpoints/2"
			}
		],
		"FileSystems": [{
				"@odata.id": "/redfish/v1/FileSystem/1"
			},
			{
				"@odata.id": "/redfish/v1/FileSystem/2"
			}
		],
		"IOConnectivityLoSCapabilities": {
			"@odata.id": "/redfish/v1/IOConnectivityLoSCapabilities/1"
		},
		"IOPerformanceLoSCapabilities": {
			"@odata.id": "/redfish/v1/IOPerformanceLoSCapabilities/1"
		},
		"IOStatistics": {
			"NonIORequestTime": "P0Y0M0DT0H0M5S",
			"NonIORequests": 1000,
			"ReadHitIORequests": 500,
			"ReadIOKiBytes": 1024,
			"ReadIORequestTime": "P0Y0M0DT0H0M5S",
			"ReadIORequests": 5000,
			"WriteHitIORequests": 100,
			"WriteIOKiBytes": 1024,
			"WriteIORequestTime": "P0Y0M0DT0H0M5S",
			"WriteIORequests": 5000
		},
		"Links": {
			"HostingSystem": {
				"@odata.id": "/redfish/v1/Hosts/1"
			}
		},
		"SpareResourceSets": [{
			"@odata.id": "/redfish/v1/SpareResourceSets/1"
		}],
		"SpareResourceSets@odata.count": 1,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"StorageGroups": [{
			"@odata.id": "/redfish/v1/StorageGroups/1"
		}],
		"StoragePools": [{
			"@odata.id": "/redfish/v1/StoragePools/1"
		}],
		"StorageSubsystems": {
			"@odata.id": "/redfish/v1/StorageSubsystems/1"
		},
		"Volumes": [{
			"@odata.id": "/redfish/v1/Volumes/1"
		}, {
			"@odata.id": "/redfish/v1/Volumes/2"
		}]
	}`)

// TestStorageService tests the parsing of StorageService objects.
func TestStorageService(t *testing.T) {
	var result StorageService
	err := json.NewDecoder(storageServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "StorageService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "StorageServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.classesOfService != "/redfish/v1/ClassesOfService" {
		t.Errorf("Invalid ClassesOfService link: %s", result.classesOfService)
	}

	if len(result.endpointGroups) != 2 {
		t.Errorf("Expected 2 endpoint groups, got %d", len(result.endpointGroups))
	}

	if len(result.endpoints) != 2 {
		t.Errorf("Expected 2 endpoints, got %d", len(result.endpoints))
	}

	if result.IOStatistics.NonIORequests != 1000 {
		t.Errorf("Invalid IOStats NonIORequests: %d", result.IOStatistics.NonIORequests)
	}
}
