//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var usbControllerBody = `{
	"@odata.type": "#USBController.v1_0_0.USBController",
	"Id": "USB1",
	"Name": "Contoso USB Controller 1",
	"Manufacturer": "Contoso",
	"Model": "USBv3",
	"SKU": "80937",
	"SerialNumber": "2M220100SL",
	"PartNumber": "G37891",
	"SparePartNumber": "G37890",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Ports": {
	  "@odata.id": "/redfish/v1/Systems/1/USBControllers/USB1/Ports"
	},
	"Links": {
	  "Processors": [
		{
		  "@odata.id": "/redfish/v1/Systems/1/Processors/1"
		},
		{
		  "@odata.id": "/redfish/v1/Systems/1/Processors/2"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Systems/1/USBControllers/USB1"
  }`

// TestUSBController tests the parsing of USBController objects.
func TestUSBController(t *testing.T) {
	var result USBController
	err := json.NewDecoder(strings.NewReader(usbControllerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "USB1", result.ID)
	assertEquals(t, "/redfish/v1/Systems/1/USBControllers/USB1/Ports", result.ports)
	assertEquals(t, "/redfish/v1/Systems/1/Processors/2", result.processors[1])
}
