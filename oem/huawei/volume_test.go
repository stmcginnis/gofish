//
// SPDX-License-Identifier: BSD-3-Clause
//

package huawei

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

// data from /redfish/v1/Systems/1/Storages/RAIDStorage0/Volumes/LogicalDrive0

var huaweiVolumeBody = `
{
  "@odata.context": "/redfish/v1/$metadata#Systems/Members/1/Storages/Members/RAIDStorage0/Volumes/Members/$entity",
  "@odata.id": "/redfish/v1/Systems/1/Storages/RAIDStorage0/Volumes/LogicalDrive0",
  "@odata.type": "#Volume.v1_0_1.Volume",
  "Id": "LogicalDrive0",
  "Name": "LogicalDrive0",
  "CapacityBytes": 1198999470080,
  "VolumeType": "Mirrored",
  "OptimumIOSizeBytes": 262144,
  "Status": {
    "State": "Enabled",
    "Health": "OK",
    "Oem": {
      "Huawei": {
        "Severity": "Informational"
      }
    }
  },
  "Operations": [],
  "Oem": {
    "Huawei": {
      "State": "Optimal",
      "VolumeName": "N/A",
      "RaidControllerID": 0,
      "VolumeRaidLevel": "RAID1",
      "DefaultReadPolicy": "ReadAhead",
      "DefaultWritePolicy": "WriteBackWithBBU",
      "DefaultCachePolicy": "DirectIO",
      "ConsistencyCheck": false,
      "SpanNumber": 1,
      "NumDrivePerSpan": 2,
      "Spans": [
        {
          "SpanName": "Span0",
          "Drives": [
            {
              "@odata.id": "/redfish/v1/Chassis/1/Drives/HDDPlaneDisk0"
            },
            {
              "@odata.id": "/redfish/v1/Chassis/1/Drives/HDDPlaneDisk1"
            }
          ]
        }
      ],
      "CurrentReadPolicy": "ReadAhead",
      "CurrentWritePolicy": "WriteBackWithBBU",
      "CurrentCachePolicy": "DirectIO",
      "AccessPolicy": "ReadWrite",
      "BootEnable": true,
      "BGIEnable": true,
      "SSDCachecadeVolume": false,
      "SSDCachingEnable": false,
      "AssociatedCacheCadeVolume": [],
      "DriveCachePolicy": "Unchanged",
      "OSDriveName": null,
      "InitializationMode": "UnInit"
    }
  },
  "Links": {
    "Drives@odata.count": 2,
    "Drives": [
      {
        "@odata.id": "/redfish/v1/Chassis/1/Drives/HDDPlaneDisk0"
      },
      {
        "@odata.id": "/redfish/v1/Chassis/1/Drives/HDDPlaneDisk1"
      }
    ]
  },
  "Actions": {
    "#Volume.Initialize": {
      "target": "/redfish/v1/Systems/1/Storages/RAIDStorage0/Volumes/LogicalDrive0/Actions/Volume.Initialize",
      "@Redfish.ActionInfo": "/redfish/v1/Systems/1/Storages/RAIDStorage0/Volumes/LogicalDrive0/InitializeActionInfo"
    }
  }
}
`

// TestHpeThermalOem tests the parsing of Thermal objects and support oem field.
func TestHuaweiVolumeOem(t *testing.T) {
	var volume redfish.Volume
	err := json.NewDecoder(strings.NewReader(huaweiVolumeBody)).Decode(&volume)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	huaweiVolume, err := FromVolume(&volume)
	if err != nil {
		t.Errorf("Convert volume failed: %s", err)
		return
	}

	if huaweiVolume.Oem.Huawei.SSDCachecadeVolume != false {
		t.Errorf("Received invalid huaweiVolume configuration: %v", huaweiVolume.Oem.Huawei.SSDCachecadeVolume)
	}
	if huaweiVolume.Oem.Huawei.State != "Optimal" {
		t.Errorf("Received invalid huaweiVolume configuration: %s", huaweiVolume.Oem.Huawei.State)
	}
}
