//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

var smcDriveBody = `{
    "@odata.type": "#Drive.v1_6_2.Drive",
    "@odata.id": "/redfish/v1/Chassis/NVMeSSD.0.Group.0.StorageBackplane/Drives/Disk.Bay.22",
    "Name": "Disk.Bay.22",
    "Id": "22",
    "Manufacturer": "INTEL",
    "SerialNumber": "PHLWOOFMEOWIAMCATDOG",
    "Model": "INTEL SSDPE2KX080T8O",
    "StatusIndicator": "OK",
    "FailurePredicted": false,
    "CapacityBytes": 8001563222016,
    "CapableSpeedGbs": 31.5,
    "Oem": {
        "Supermicro": {
            "@odata.type": "#SmcDriveExtensions.v1_0_0.Drive",
            "Temperature": 33,
            "PercentageDriveLifeUsed": 3,
            "DriveFunctional": true
        }
    },
    "IndicatorLED": "Off",
    "Status": {
        "State": "Enabled",
        "Health": "OK"
    },
    "Links": {
        "Volumes": []
    },
    "Actions": {
        "Oem": {
            "#Drive.Indicate": {
                "target": "/redfish/v1/Chassis/NVMeSSD.0.Group.0.StorageBackplane/Drives/Disk.Bay.22/Actions/Oem/Drive.Indicate",
                "@Redfish.ActionInfo": "/redfish/v1/Chassis/NVMeSSD.0.Group.0.StorageBackplane/Drives/Disk.Bay.22/IndicateActionInfo"
            }
        }
    }
}`

// TestSmcDriveOem tests the parsing of the Drive oem field
func TestSmcDriveOem(t *testing.T) {
	drive := &redfish.Drive{}
	if err := json.Unmarshal([]byte(smcDriveBody), drive); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	smcDrive, err := FromDrive(drive)
	if err != nil {
		t.Fatalf("error getting oem info from drive: %v", err)
	}

	if smcDrive.Temperature != 33 {
		t.Errorf("unexpected oem drive temperature: %d", smcDrive.Temperature)
	}

	if smcDrive.indicateTarget != "/redfish/v1/Chassis/NVMeSSD.0.Group.0.StorageBackplane/Drives/Disk.Bay.22/Actions/Oem/Drive.Indicate" {
		t.Errorf("unexpected oem drive indicator target: %s", smcDrive.indicateTarget)
	}
}
