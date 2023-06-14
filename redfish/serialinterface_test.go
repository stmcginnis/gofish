//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var serialInterfaceBody = `{
		"@odata.type": "#SerialInterface.v1_1_8.SerialInterface",
		"Id": "TTY0",
		"Name": "Manager Serial Interface 1",
		"Description": "Management for Serial Interface",
		"InterfaceEnabled": true,
		"SignalType": "Rs232",
		"BitRate": "115200",
		"Parity": "None",
		"DataBits": "8",
		"StopBits": "1",
		"FlowControl": "None",
		"ConnectorType": "RJ45",
		"PinOut": "Cyclades",
		"@odata.id": "/redfish/v1/Managers/BMC/SerialInterfaces/TTY0"
	}`

// TestSerialInterface tests the parsing of SerialInterface objects.
func TestSerialInterface(t *testing.T) {
	var result SerialInterface

	if err := json.NewDecoder(strings.NewReader(serialInterfaceBody)).Decode(&result); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "TTY0", result.ID)
	assertEquals(t, "Manager Serial Interface 1", result.Name)
	assertEquals(t, "Rs232", string(result.SignalType))
	assertEquals(t, "115200", string(result.BitRate))
	assertEquals(t, "RJ45", string(result.ConnectorType))
	assertEquals(t, "Cyclades", string(result.PinOut))
}

// TestSerialInterfaceUpdate tests the Update call.
func TestSerialInterfaceUpdate(t *testing.T) {
	var result SerialInterface

	if err := json.NewDecoder(strings.NewReader(serialInterfaceBody)).Decode(&result); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	newBitRate := BitRate57600
	newFlowControl := HardwareSerialConnectionFlowControl
	newParity := OddParityBit

	result.BitRate = newBitRate
	result.FlowControl = newFlowControl
	result.Parity = newParity

	if err := result.Update(); err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	actual := calls[0].Payload

	propertyName := "BitRate"
	expected := fmt.Sprintf("%s:%v", propertyName, newBitRate)
	assertContains(t, expected, actual)

	propertyName = "FlowControl"
	expected = fmt.Sprintf("%s:%v", propertyName, newFlowControl)
	assertContains(t, expected, actual)

	propertyName = "Parity"
	expected = fmt.Sprintf("%s:%v", propertyName, newParity)
	assertContains(t, expected, actual)
}
