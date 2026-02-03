//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var portBody = `{
    "@odata.type": "#Port.v1_10_0.Port",
    "Id": "1",
    "Name": "SAS Port 1",
    "Description": "SAS Port 1",
    "Status": {
        "State": "Enabled",
        "Health": "OK"
    },
    "LinkStatus": "LinkUp",
    "PortId": "1",
    "PortProtocol": "SAS",
    "PortType": "BidirectionalPort",
    "CurrentSpeedGbps": 48,
    "Width": 4,
    "MaxSpeedGbps": 48,
    "Links": {
        "AssociatedEndpoints": [
            {
                "@odata.id": "/redfish/v1/Fabrics/SAS/Endpoints/Initiator1"
            }
        ]
    },
    "@odata.id": "/redfish/v1/Fabrics/SAS/Switches/Switch1/Ports/1"
}`

// TestPort tests the parsing of Port objects.
func TestPort(t *testing.T) {
	var result Port
	err := json.NewDecoder(strings.NewReader(portBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "SAS Port 1", result.Name)
	assertEquals(t, "1", result.PortID)
	assertEquals(t, "BidirectionalPort", fmt.Sprint(result.PortType))
	assertEquals(t, "SAS", fmt.Sprint(result.PortProtocol))
	assertEquals(t, "48", fmt.Sprint(*result.CurrentSpeedGbps))
	assertEquals(t, "LinkUp", fmt.Sprint(result.LinkStatus))
}
