//
// SPDX-License-Identifier: BSD-3-Clause
//

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
		"EndpointGroups": {
			"@odata.id": "/redfish/v1/EndpointGroups"
		},
		"Endpoints": {
			"@odata.id": "/redfish/v1/Endpoints"
		},
		"FileSystems": {
			"@odata.id": "/redfish/v1/FileSystems"
		},
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
		"Volumes": {
			"@odata.id": "/redfish/v1/Volumes/1"
		}
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

	if result.endpointGroups != "/redfish/v1/EndpointGroups" {
		t.Errorf("Invalid EndpointGroups link: %s", result.endpointGroups)
	}

	if result.endpoints != "/redfish/v1/Endpoints" {
		t.Errorf("Invalid Endpoints link: %s", result.endpoints)
	}

	if result.IOStatistics.NonIORequests != 1000 {
		t.Errorf("Invalid IOStats NonIORequests: %d", result.IOStatistics.NonIORequests)
	}

	if result.volumes != "/redfish/v1/Volumes/1" {
		t.Errorf("Invalid volumes collection link: %s", result.volumes)
	}
}
