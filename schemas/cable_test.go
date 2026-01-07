//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var cableBody = `{
	"@odata.type": "#Cable.v1_2_2.Cable",
	"Id": "hdmi_dp",
	"Name": "HDMI to DP Cable",
	"UserDescription": "HDMI to DisplayPort Cable",
	"UpstreamName": "HDMI0",
	"DownstreamName": "Video Out",
	"CableType": "HDMI",
	"LengthMeters": 0.1,
	"CableClass": "Video",
	"UpstreamConnectorTypes": [
	  "HDMI"
	],
	"DownstreamConnectorTypes": [
	  "DisplayPort"
	],
	"Links": {
	  "UpstreamChassis": [
		{
		  "@odata.id": "/redfish/v1/Chassis/bmc"
		}
	  ]
	},
	"PartNumber": "934AMS02X",
	"Manufacturer": "Cable Co.",
	"SerialNumber": "2345791",
	"Vendor": "Cablestore",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"CableStatus": "Normal",
	"@odata.id": "/redfish/v1/Cables/hdmi_dp"
  }`

// TestCable tests the parsing of Cable objects.
func TestCable(t *testing.T) {
	var result Cable
	err := json.NewDecoder(strings.NewReader(cableBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "hdmi_dp", result.ID)
	assertEquals(t, "HDMI to DP Cable", result.Name)
	assertEquals(t, "HDMI to DisplayPort Cable", result.UserDescription)
	assertEquals(t, "HDMI0", result.UpstreamName)
	assertEquals(t, "Video Out", result.DownstreamName)
	assertEquals(t, "HDMI", result.CableType)
	assertEquals(t, "Normal", string(result.CableStatus))
	assertEquals(t, "Video", string(result.CableClass))
	assertEquals(t, "/redfish/v1/Chassis/bmc", result.upstreamChassis[0])
	assertEquals(t, "HDMI", string(result.UpstreamConnectorTypes[0]))
	assertEquals(t, "DisplayPort", string(result.DownstreamConnectorTypes[0]))

	if *result.LengthMeters != 0.1 {
		t.Errorf("Unexpected cable meter length: %.2f", *result.LengthMeters)
	}
}
