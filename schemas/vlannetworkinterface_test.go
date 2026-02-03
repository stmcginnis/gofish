//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var vlanNetworkInterfaceBody = `{
		"@odata.context": "/redfish/v1/$metadata#VlanNetworkInterface.VlanNetworkInterface",
		"@odata.type": "#VLanNetworkInterface.v1_1_3.VLanNetworkInterface",
		"@odata.id": "/redfish/v1/VlanNetworkInterface",
		"Id": "VlanNetworkInterface-1",
		"Name": "VlanNetworkInterfaceOne",
		"Description": "VlanNetworkInterface One",
		"VLANEnable": true,
		"VLANId": 200
	}`

// TestVlanNetworkInterface tests the parsing of VlanNetworkInterface objects.
func TestVlanNetworkInterface(t *testing.T) {
	var result VLanNetworkInterface
	err := json.NewDecoder(strings.NewReader(vlanNetworkInterfaceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "VlanNetworkInterface-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "VlanNetworkInterfaceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.VLANEnable {
		t.Error("VLAN should be enabled")
	}

	if result.VLANID != 200 {
		t.Errorf("Invalid VLAN ID: %d", result.VLANID)
	}
}

// TestVlanNetworkInterfaceUpdate tests the Update call.
func TestVlanNetworkInterfaceUpdate(t *testing.T) {
	var result VLanNetworkInterface
	err := json.NewDecoder(strings.NewReader(vlanNetworkInterfaceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	result.VLANEnable = false
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "VLANEnable:false") {
		t.Errorf("Unexpected VLANEnable update payload: %s", calls[0].Payload)
	}
}
