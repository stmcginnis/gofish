//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageControllerBody = `{
	"@odata.type": "#StorageController.v1_7_2.StorageController",
	"Id": "1",
	"Name": "NVMe IO Controller",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"SupportedControllerProtocols": [
	  "NVMeOverFabrics"
	],
	"NVMeControllerProperties": {
	  "NVMeVersion": "1.4",
	  "ControllerType": "IO",
	  "NVMeControllerAttributes": {
		"ReportsUUIDList": false,
		"SupportsSQAssociations": false,
		"ReportsNamespaceGranularity": false,
		"SupportsTrafficBasedKeepAlive": false,
		"SupportsPredictableLatencyMode": false,
		"SupportsEnduranceGroups": false,
		"SupportsReadRecoveryLevels": false,
		"SupportsNVMSets": true,
		"SupportsExceedingPowerOfNonOperationalState": false,
		"Supports128BitHostId": false
	  },
	  "NVMeSMARTCriticalWarnings": {
		"PMRUnreliable": false,
		"PowerBackupFailed": false,
		"MediaInReadOnly": false,
		"OverallSubsystemDegraded": false,
		"SpareCapacityWornOut": false
	  }
	},
	"Links": {
	  "Endpoints": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator1"
		},
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Endpoints/Target1"
		}
	  ],
	  "AttachedVolumes": [
		{
		  "@odata.id": "/redfish/v1/Storage/NVMeoF/Volumes/1"
		},
		{
		  "@odata.id": "/redfish/v1/Storage/NVMeoF/Volumes/3"
		},
		{
		  "@odata.id": "/redfish/v1/Storage/NVMeoF/Volumes/4"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Storage/NVMeoF/Controllers/1"
  }`

// TestStorageController tests the parsing of StorageController objects.
func TestStorageController(t *testing.T) {
	var result StorageController
	err := json.NewDecoder(strings.NewReader(storageControllerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "NVMe IO Controller", result.Name)
	assertEquals(t, "NVMeOverFabrics", string(result.SupportedControllerProtocols[0]))
	assertEquals(t, "IO", string(result.NVMeControllerProperties.ControllerType))
	assertEquals(t, "/redfish/v1/Fabrics/NVMeoF/Endpoints/Target1", result.endpoints[1])
	assertEquals(t, "/redfish/v1/Storage/NVMeoF/Volumes/1", result.attachedVolumes[0])

	if !result.NVMeControllerProperties.NVMeControllerAttributes.SupportsNVMSets {
		t.Error("Expected NVMeControllerProperties.NVMeControllerAttributes.SupportsNVMSets to be true")
	}
}
