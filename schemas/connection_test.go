//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var connectionBody = `{
	"@odata.type": "#Connection.v1_3_1.Connection",
	"Id": "1",
	"Name": "Connection info for host 1",
	"ConnectionType": "Storage",
	"VolumeInfo": [
	  {
		"AccessCapabilities": [
		  "Read",
		  "Write"
		],
		"Volume": {
		  "@odata.id": "/redfish/v1/Storage/NVMeoF/Volumes/1"
		}
	  },
	  {
		"AccessCapabilities": [
		  "Read",
		  "Write"
		],
		"Volume": {
		  "@odata.id": "/redfish/v1/Storage/NVMeoF/Volumes/3"
		}
	  }
	],
	"Links": {
	  "InitiatorEndpoints": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Fabrics/NVMeoF/Connections/1"
  }`

// TestConnection tests the parsing of Connection objects.
func TestConnection(t *testing.T) {
	var result Connection
	err := json.NewDecoder(strings.NewReader(connectionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Connection info for host 1", result.Name)
	assertEquals(t, "/redfish/v1/Storage/NVMeoF/Volumes/1", result.VolumeInfo[0].volume)
	assertEquals(t, "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator1", result.initiatorEndpoints[0])
}
