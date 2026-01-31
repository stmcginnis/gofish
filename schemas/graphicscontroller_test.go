//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var graphicsControllerBody = `{
	"@odata.type": "#GraphicsController.v1_0_1.GraphicsController",
	"Id": "GPU1",
	"Name": "Contoso Graphics Controller 1",
	"AssetTag": "",
	"Manufacturer": "Contoso",
	"Model": "GPU1",
	"SKU": "80937",
	"SerialNumber": "2M220100SL",
	"PartNumber": "G37891",
	"SparePartNumber": "G37890",
	"BiosVersion": "90.02.17.00.7D",
	"DriverVersion": "27.21.14.6079 (Contoso 460.79) DCH / Win 10 64",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Location": {
	  "PartLocation": {
		"ServiceLabel": "Slot 1",
		"LocationOrdinalValue": 1,
		"LocationType": "Slot",
		"Orientation": "LeftToRight",
		"Reference": "Rear"
	  }
	},
	"Ports": {
	  "@odata.id": "/redfish/v1/Systems/1/GraphicsControllers/GPU1/Ports"
	},
	"Links": {
	  "Processors": [
		{
		  "@odata.id": "/redfish/v1/Systems/1/Processors/GPU"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Systems/1/GraphicsControllers/GPU1"
  }`

// TestGraphicsController tests the parsing of GraphicsController objects.
func TestGraphicsController(t *testing.T) {
	var result GraphicsController
	err := json.NewDecoder(strings.NewReader(graphicsControllerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "GPU1", result.ID)
	assertEquals(t, "Contoso Graphics Controller 1", result.Name)
	assertEquals(t, "90.02.17.00.7D", result.BiosVersion)
	assertEquals(t, "27.21.14.6079 (Contoso 460.79) DCH / Win 10 64", result.DriverVersion)
	assertEquals(t, "LeftToRight", string(result.Location.PartLocation.Orientation))
	assertEquals(t, "/redfish/v1/Systems/1/GraphicsControllers/GPU1/Ports", result.ports)
	assertEquals(t, "/redfish/v1/Systems/1/Processors/GPU", result.processors[0])
}
