//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var messageRegistryFileBody = `{
		"@odata.context": "/redfish/v1/$metadata#MessageRegistryFile.MessageRegistryFile",
		"@odata.id": "/redfish/v1/Registries/MyRegistry",
		"@odata.type": "#MessageRegistryFile.v1_0_4.MessageRegistryFile",
		"Id": "MyRegistry",
		"Description": "Registry Definition File for MyRegistry",
		"Languages": [
			"en"
		],
		"Location": [
			{
				"Language": "en",
				"Uri": "/redfish/v1/RegistryStore/registries/en/MyRegistry.json"
			}
		],
		"Name": "MyRegistry Message Registry File",
		"Registry": "MyRegistry.2.2.0"
	}`

// TestMessageRegistryFile tests the parsing of MessageRegistryFile objects.
func TestMessageRegistryFile(t *testing.T) {
	var result MessageRegistryFile
	err := json.NewDecoder(strings.NewReader(messageRegistryFileBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.Registry != "MyRegistry.2.2.0" {
		t.Errorf("Received invalid Registry: %s", result.Registry)
	}

	if result.ODataContext != "/redfish/v1/$metadata#MessageRegistryFile.MessageRegistryFile" {
		t.Errorf("Received invalid ODataContext: %s", result.ODataContext)
	}

	if result.ODataID != "/redfish/v1/Registries/MyRegistry" {
		t.Errorf("Received invalid ODataID: %s", result.ODataID)
	}

	if result.ODataType != "#MessageRegistryFile.v1_0_4.MessageRegistryFile" {
		t.Errorf("Received invalid ODataType: %s", result.ODataType)
	}
	if result.ID != "MyRegistry" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Description != "Registry Definition File for MyRegistry" {
		t.Errorf("Received invalid Description: %s", result.Description)
	}

	if result.Languages[0] != "en" {
		t.Errorf("Received invalid Languages: %s", result.Languages[0])
	}

	if result.Name != "MyRegistry Message Registry File" {
		t.Errorf("Received invalid Name: %s", result.Name)
	}

	if result.Location[0].Language != "en" {
		t.Errorf("Received invalid Location[0].Language: %s", result.Location[0].Language)
	}

	if result.Location[0].URI != "/redfish/v1/RegistryStore/registries/en/MyRegistry.json" {
		t.Errorf("Received invalid Location[0].Uri: %s", result.Location[0].URI)
	}
}
