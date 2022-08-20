//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var messageRegistryBody = `{
		"@odata.type": "#MessageRegistry.v1_2_0.MessageRegistry",
		"Description": "This registry is an example.",
		"Id": "MyRegistry.json",
		"Language": "en",
		"Messages": {
			"FirstMessage": {
				"Description": "Example of message with one arg.",
				"Message": "This message has only one arg: %1",
				"NumberOfArgs": 1,
				"ParamTypes": [
					"string"
				],
				"Resolution": "The resolution for the first message.",
				"Severity": "OK"
			},
			"SecondMessage": {
				"Description": "Example of message without args.",
				"Message": "This message has no args.",
				"NumberOfArgs": 0,
				"ParamTypes": [],
				"Resolution": "The resolution for the second message.",
				"Severity": "Critical"
			},
			"ThirdMessage": {
				"Description": "Example of message with two args.",
				"Message": "This message has two args: %1 and %2",
				"NumberOfArgs": 2,
				"ParamTypes": [
					"string",
					"string"
				],
				"Resolution": "The resolution for the third message.",
				"Severity": "Warning"
			},
			"MessageWithOem": {
				"Description": "Example of message with Oem.",
				"Message": "This message has Oem info.",
				"NumberOfArgs": 0,
				"Oem": {
					"VendorName": {
						"OemInfo1": "The Oem info 1.",
						"OemInfoN": "The Oem info N."
					}
				},
				"ParamTypes": [],
				"Resolution": "The resolution for the message with Oem.",
				"Severity": "Critical"
			}
		},
		"Name": "MyRegistry Registry",
		"OwningEntity": "The vendor name",
		"RegistryPrefix": "MyRegistry",
		"RegistryVersion": "2.2.0"
	}`

// TestMessageRegistry tests the parsing of MessageRegistry objects.
func TestMessageRegistry(t *testing.T) { //nolint:funlen,gocyclo
	var result MessageRegistry
	err := json.NewDecoder(strings.NewReader(messageRegistryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "MyRegistry.json" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Description != "This registry is an example." {
		t.Errorf("Received invalid Description: %s", result.Description)
	}

	if result.Language != "en" {
		t.Errorf("Received invalid Language: %s", result.Language)
	}

	if result.Name != "MyRegistry Registry" {
		t.Errorf("Received invalid Name: %s", result.Name)
	}

	if result.ODataType != "#MessageRegistry.v1_2_0.MessageRegistry" {
		t.Errorf("Received invalid ODataType: %s", result.ODataType)
	}

	if result.OwningEntity != "The vendor name" {
		t.Errorf("Received invalid OwningEntity: %s", result.OwningEntity)
	}

	if result.RegistryPrefix != "MyRegistry" {
		t.Errorf("Received invalid RegistryPrefix: %s", result.RegistryPrefix)
	}

	if result.RegistryVersion != "2.2.0" {
		t.Errorf("Received invalid RegistryVersion: %s", result.RegistryVersion)
	}

	// test the messages

	if len(result.Messages) != 4 {
		t.Errorf("Received invalid number of Messages: %d", len(result.Messages))
	}

	// FirstMessage
	messageKey := "FirstMessage"
	if m, ok := result.Messages[messageKey]; ok {
		if m.Description != "Example of message with one arg." {
			t.Errorf("Received invalid Description: %s for the messageKey: %s", m.Description, messageKey)
		}
		if m.Message != "This message has only one arg: %1" {
			t.Errorf("Received invalid Message: %s for the messageKey: %s", m.Message, messageKey)
		}
		if m.NumberOfArgs != 1 {
			t.Errorf("Received invalid NumberOfArgs: %d for the messageKey: %s", m.NumberOfArgs, messageKey)
		}
		if m.ParamTypes[0] != "string" {
			t.Errorf("Received invalid ParamTypes: %s for the messageKey: %s", m.ParamTypes[0], messageKey)
		}
		if m.Resolution != "The resolution for the first message." {
			t.Errorf("Received invalid Resolution: %s for the messageKey: %s", m.Resolution, messageKey)
		}
		if m.Severity != "OK" {
			t.Errorf("Received invalid Severity: %s for the messageKey: %s", m.Severity, messageKey)
		}
	} else {
		t.Errorf("MessageKey %s not found.", messageKey)
	}

	// SecondMessage
	messageKey = "SecondMessage"
	if m, ok := result.Messages[messageKey]; ok {
		if m.Description != "Example of message without args." {
			t.Errorf("Received invalid Description: %s for the messageKey: %s", m.Description, messageKey)
		}
		if m.Message != "This message has no args." {
			t.Errorf("Received invalid Message: %s for the messageKey: %s", m.Message, messageKey)
		}
		if m.NumberOfArgs != 0 {
			t.Errorf("Received invalid NumberOfArgs: %d for the messageKey: %s", m.NumberOfArgs, messageKey)
		}
		if len(m.ParamTypes) > 0 {
			t.Errorf("Received invalid ParamTypes: %v for the messageKey: %s", m.ParamTypes, messageKey)
		}
		if m.Resolution != "The resolution for the second message." {
			t.Errorf("Received invalid Resolution: %s for the messageKey: %s", m.Resolution, messageKey)
		}
		if m.Severity != "Critical" {
			t.Errorf("Received invalid Severity: %s for the messageKey: %s", m.Severity, messageKey)
		}
	} else {
		t.Errorf("MessageKey %s not found.", messageKey)
	}

	// ThirdMessage
	messageKey = "ThirdMessage"
	if m, ok := result.Messages[messageKey]; ok {
		if m.Description != "Example of message with two args." {
			t.Errorf("Received invalid Description: %s for the messageKey: %s", m.Description, messageKey)
		}
		if m.Message != "This message has two args: %1 and %2" {
			t.Errorf("Received invalid Message: %s for the messageKey: %s", m.Message, messageKey)
		}
		if m.NumberOfArgs != 2 {
			t.Errorf("Received invalid NumberOfArgs: %d for the messageKey: %s", m.NumberOfArgs, messageKey)
		}
		if m.ParamTypes[0] != "string" {
			t.Errorf("Received invalid ParamTypes[0]: %s for the messageKey: %s", m.ParamTypes[0], messageKey)
		}
		if m.ParamTypes[1] != "string" {
			t.Errorf("Received invalid ParamTypes[1]: %s for the messageKey: %s", m.ParamTypes[1], messageKey)
		}
		if m.Resolution != "The resolution for the third message." {
			t.Errorf("Received invalid Resolution: %s for the messageKey: %s", m.Resolution, messageKey)
		}
		if m.Severity != "Warning" {
			t.Errorf("Received invalid Severity: %s for the messageKey: %s", m.Severity, messageKey)
		}
	} else {
		t.Errorf("MessageKey %s not found.", messageKey)
	}

	// MessageWithOem
	messageKey = "MessageWithOem"
	if m, ok := result.Messages[messageKey]; ok {
		if m.Description != "Example of message with Oem." {
			t.Errorf("Received invalid Description: %s for the messageKey: %s", m.Description, messageKey)
		}
		if m.Message != "This message has Oem info." {
			t.Errorf("Received invalid Message: %s for the messageKey: %s", m.Message, messageKey)
		}
		if m.NumberOfArgs != 0 {
			t.Errorf("Received invalid NumberOfArgs: %d for the messageKey: %s", m.NumberOfArgs, messageKey)
		}
		if len(m.ParamTypes) > 0 {
			t.Errorf("Received invalid ParamTypes: %v for the messageKey: %s", m.ParamTypes, messageKey)
		}
		if m.Resolution != "The resolution for the message with Oem." {
			t.Errorf("Received invalid Resolution: %s for the messageKey: %s", m.Resolution, messageKey)
		}
		if m.Severity != "Critical" {
			t.Errorf("Received invalid Severity: %s for the messageKey: %s", m.Severity, messageKey)
		}

		// test oem
		switch oem := m.Oem.(type) {
		case map[string]interface{}:
			for vendor, values := range oem {
				if vendor != "VendorName" {
					t.Errorf("Received invalid Oem vendor: %s for the messageKey: %s", vendor, messageKey)
				}
				switch val := values.(type) {
				case map[string]interface{}:
					for k, v := range val {
						if k != "OemInfo1" && k != "OemInfoN" {
							t.Errorf("Received invalid Oem key %s for vendor: %s for the messageKey: %s", k, vendor, messageKey)
						}
						if k == "OemInfo1" && v != "The Oem info 1." {
							t.Errorf("Received invalid OemInfo1: %s for the messageKey: %s", v, messageKey)
						}
						if v == "OemInfoN" && v != "The Oem info N." {
							t.Errorf("Received invalid OemInfoN: %s for the messageKey: %s", v, messageKey)
						}
					}
				default:
					t.Error("Unexpected value format")
				}
			}
		default:
			t.Errorf("Received invalid Oem for the messageKey: %s", messageKey)
		}
	} else {
		t.Errorf("MessageKey %s not found.", messageKey)
	}
}
