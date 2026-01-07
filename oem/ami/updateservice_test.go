//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var updateServiceBody = `{
  "@odata.context": "/redfish/v1/$metadata#UpdateService.UpdateService",
  "@odata.etag": "\"1729105654\"",
  "@odata.id": "/redfish/v1/UpdateService",
  "@odata.type": "#UpdateService.v1_6_0.UpdateService",
  "Actions": {
    "#UpdateService.SimpleUpdate": {
      "@Redfish.ActionInfo": "/redfish/v1/UpdateService/SimpleUpdateActionInfo",
      "target": "/redfish/v1/UpdateService/Actions/SimpleUpdate"
    },
    "Oem": {
      "#UpdateService.UploadCABundle": {
        "@Redfish.ActionInfo": "/redfish/v1/UpdateService/UploadCABundleActionInfo",
        "target": "/redfish/v1/UpdateService/Actions/Oem/UpdateService.UploadCABundle"
      }
    }
  },
  "Description": "Redfish Update Service",
  "FirmwareInventory": {
    "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory"
  },
  "Id": "UpdateService",
  "MaxImageSizeBytes": 441393152,
  "MultipartHttpPushUri": "/redfish/v1/UpdateService/upload",
  "Name": "Update Service",
  "Oem": {
    "AMIUpdateService": {
      "@odata.type": "#AMIUpdateService.v1_0_0.AMIUpdateService",
      "FlashPercentage": null,
      "PreserveConfiguration": true,
      "UpdateStatus": null,
      "UpdateTarget": null
    },
    "BIOS": {
      "BIOSPreserveNVRAM": true
    },
    "BMC": {
      "@odata.type": "#AMIUpdateService.v1_0_0.BMC",
      "DualImageConfigurations": {
        "ActiveImage": "1",
        "BootImage": "1",
        "FirmwareImage1Name": "Image1",
        "FirmwareImage1Version": "3.35.00",
        "FirmwareImage2Name": "Image2",
        "FirmwareImage2Version": "3.35.00"
      }
    }
  },
  "ServiceEnabled": true,
  "Status": {
    "Health": "OK",
    "State": "Enabled"
  }
}`

// TestUpdateService tests the parsing of the UpdateService oem fields.
func TestUpdateService(t *testing.T) {
	us := &schemas.UpdateService{}
	if err := json.Unmarshal([]byte(updateServiceBody), us); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	updateService, err := FromUpdateService(us)
	if err != nil {
		t.Fatalf("error getting oem object: %v", err)
	}

	if updateService.ID != "UpdateService" {
		t.Errorf("unexpected ID: %s", updateService.ID)
	}

	if updateService.uploadCABundleTarget != "/redfish/v1/UpdateService/Actions/Oem/UpdateService.UploadCABundle" {
		t.Errorf("unexpected uploadCABundle target: %s", updateService.uploadCABundleTarget)
	}

	if !updateService.AMIUpdateService.PreserveConfiguration {
		t.Errorf("unexpected preserve configuration: %t", updateService.AMIUpdateService.PreserveConfiguration)
	}

	if !updateService.BIOS.BIOSPreserveNVRAM {
		t.Errorf("unexpected preserve nvram: %t", updateService.BIOS.BIOSPreserveNVRAM)
	}

	if updateService.BMC.DualImageConfigurations.ActiveImage != "1" {
		t.Errorf("unexpected active image: %s", updateService.BMC.DualImageConfigurations.ActiveImage)
	}
}
