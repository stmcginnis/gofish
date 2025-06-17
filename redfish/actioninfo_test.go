//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var actionInfoBody = `{
		"@odata.type": "#ActionInfo.v1_4_2.ActionInfo",
		"Id": "ResetActionInfo",
		"Name": "Reset Action Info",
		"Parameters": [
			{
				"Name": "ResetType",
				"Required": true,
				"DataType": "String",
				"AllowableValues": [
					"On",
					"ForceOff"
				]
			}
		],
		"@odata.id": "/redfish/v1/Systems/1/ResetActionInfo"
	}`

func TestActionInfo(t *testing.T) {
	var result ActionInfo
	err := json.NewDecoder(strings.NewReader(actionInfoBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ODataType != "#ActionInfo.v1_4_2.ActionInfo" {
		t.Errorf("Received invalid ODataType: %s", result.ODataType)
	}

	if result.ID != "ResetActionInfo" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Reset Action Info" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ODataID != "/redfish/v1/Systems/1/ResetActionInfo" {
		t.Errorf("Received invalid ODataID: %s", result.ODataID)
	}

	if len(result.Parameters) != 1 {
		t.Errorf("Received invalid number of parameters: %d", len(result.Parameters))
	}

	firstParam := result.Parameters[0]

	if firstParam.Name != "ResetType" {
		t.Errorf("Received invalid param name: %s", firstParam.Name)
	}

	if firstParam.Required != true {
		t.Errorf("Received invalid param required value: %t", firstParam.Required)
	}

	if firstParam.DataType != StringActionInfoDataTypes {
		t.Errorf("Received invalid param data type: %s", firstParam.DataType)
	}

	if len(firstParam.AllowableValues) != 2 {
		t.Errorf("Received invalid number of allowable values: %d", len(firstParam.AllowableValues))
	}

	if firstParam.AllowableValues[0] != "On" {
		t.Errorf("Received invalid allowable value [0]: %s", firstParam.AllowableValues[0])
	}

	if firstParam.AllowableValues[1] != "ForceOff" {
		t.Errorf("Received invalid allowable value [1]: %s", firstParam.AllowableValues[1])
	}
}
