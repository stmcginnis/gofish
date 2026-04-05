//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var attributeRegistryBody = strings.NewReader(
	`{
		"@odata.type": "#AttributeRegistry.v1_3_6.AttributeRegistry",
		"Description": "This registry defines a representation of BIOS Attribute instances",
		"Id": "BiosAttributeRegistryG9000.v1_0_0",
		"Language": "en",
		"Name": "G9000 BIOS Attribute Registry",
		"OwningEntity": "Contoso",
		"RegistryVersion": "1.0.0",
		"SupportedSystems": [
			{
				"ProductName": "Contoso Server GLH9000",
				"SystemId": "G9000",
				"FirmwareVersion": "v1.00 (06/02/2014)"
			}
		],
		"RegistryEntries": {
			"Attributes": [
				{
					"CurrentValue": null,
					"DisplayName": "Embedded NIC 1 Boot",
					"DisplayOrder": 5,
					"HelpText": "Select this option to enable network boot (PXE, iSCSI, or FCoE) for the selected NIC. You may need to configure the NIC firmware for the boot option to be active.",
					"MenuPath": "./SystemOptions/NetworkBootOptions",
					"AttributeName": "NicBoot1",
					"ReadOnly": false,
					"Hidden": false,
					"Type": "Enumeration",
					"Value": [
						{
							"ValueDisplayName": "Network Boot",
							"ValueName": "NetworkBoot"
						},
						{
							"ValueDisplayName": "Disabled",
							"ValueName": "Disabled"
						}
					],
					"WarningText": "Important: When enabling network boot support for an embedded NIC, the NIC boot option does not appear in the UEFI Boot Order or Legacy IPL lists until the next system reboot."
				},
				{
					"CurrentValue": null,
					"DisplayName": "Embedded SATA Configuration",
					"DisplayOrder": 74,
					"HelpText": "Important: Select this option to configure the embedded chipset SATA controller.",
					"MenuPath": "./SystemOptions/SataOptions",
					"AttributeName": "EmbeddedSata",
					"ReadOnly": false,
					"Hidden": false,
					"Type": "Enumeration",
					"Value": [
						{
							"ValueDisplayName": "Enable SATA AHCI Support",
							"ValueName": "Ahci"
						},
						{
							"ValueDisplayName": "Enable Software RAID Support",
							"ValueName": "Raid"
						}
					],
					"WarningText": "Important: Software RAID is not supported when the Boot Mode is configured in Legacy BIOS Mode."
				}
			],
			"Dependencies": [
				{
					"Dependency": {
						"MapFrom": [
							{
								"MapFromAttribute": "BootMode",
								"MapFromCondition": "EQU",
								"MapFromProperty": "CurrentValue",
								"MapFromValue": "LegacyBios"
							}
						],
						"MapToAttribute": "EmbeddedSata",
						"MapToProperty": "ReadOnly",
						"MapToValue": true
					},
					"DependencyFor": "EmbeddedSata",
					"Type": "Map"
				}
			],
			"Menus": [
				{
					"DisplayName": "BIOS Configuration",
					"DisplayOrder": 1,
					"MenuPath": "./",
					"MenuName": "BiosMainMenu",
					"Hidden": false,
					"ReadOnly": false
				},
				{
					"DisplayName": "System Options",
					"DisplayOrder": 2,
					"MenuPath": "./SystemOptions",
					"MenuName": "SystemOptions",
					"Hidden": false,
					"ReadOnly": false
				}
			]
		}
	}`)

// TestAttributeRegistryMultiTypeCurrentValue verifies that CurrentValue and
// DefaultValue unmarshal correctly when the BMC returns string, boolean, or
// numeric values. Prior to the fix for #514, these fields were *float64 and
// string values silently failed to parse.
func TestAttributeRegistryMultiTypeCurrentValue(t *testing.T) {
	body := `{
		"@odata.type": "#AttributeRegistry.v1_4_0.AttributeRegistry",
		"Id": "BiosAttributeRegistry",
		"Name": "BIOS Attribute Registry",
		"Language": "en",
		"OwningEntity": "Dell",
		"RegistryVersion": "1.0.0",
		"RegistryEntries": {
			"Attributes": [
				{
					"AttributeName": "BootMode",
					"CurrentValue": "UEFI",
					"DefaultValue": "UEFI",
					"DisplayName": "Boot Mode",
					"Type": "Enumeration"
				},
				{
					"AttributeName": "NumLock",
					"CurrentValue": true,
					"DefaultValue": false,
					"DisplayName": "NumLock State",
					"Type": "Boolean"
				},
				{
					"AttributeName": "AcPwrRcvryDelay",
					"CurrentValue": 30,
					"DefaultValue": 0,
					"DisplayName": "AC Power Recovery Delay",
					"Type": "Integer"
				},
				{
					"AttributeName": "Unset",
					"CurrentValue": null,
					"DefaultValue": null,
					"DisplayName": "Unset Attribute",
					"Type": "String"
				}
			]
		}
	}`

	var result AttributeRegistry
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&result); err != nil {
		t.Fatalf("Failed to decode AttributeRegistry: %v", err)
	}

	attrs := result.RegistryEntries.Attributes
	if len(attrs) != 4 {
		t.Fatalf("Expected 4 attributes, got %d", len(attrs))
	}

	if v, ok := attrs[0].CurrentValue.(string); !ok || v != "UEFI" {
		t.Errorf("BootMode.CurrentValue = %v (%T), want \"UEFI\"", attrs[0].CurrentValue, attrs[0].CurrentValue)
	}
	if v, ok := attrs[1].CurrentValue.(bool); !ok || v != true {
		t.Errorf("NumLock.CurrentValue = %v (%T), want true", attrs[1].CurrentValue, attrs[1].CurrentValue)
	}
	if v, ok := attrs[2].CurrentValue.(float64); !ok || v != 30 {
		t.Errorf("AcPwrRcvryDelay.CurrentValue = %v (%T), want 30", attrs[2].CurrentValue, attrs[2].CurrentValue)
	}
	if attrs[3].CurrentValue != nil {
		t.Errorf("Unset.CurrentValue = %v, want nil", attrs[3].CurrentValue)
	}
}

// TestAttributeRegistry tests the parsing of AttributeRegistry objects.
func TestAttributeRegistry(t *testing.T) {
	var result AttributeRegistry
	err := json.NewDecoder(attributeRegistryBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "BiosAttributeRegistryG9000.v1_0_0", result.ID)
	assertEquals(t, "G9000 BIOS Attribute Registry", result.Name)
	assertEquals(t, "Contoso", result.OwningEntity)
	assertEquals(t, "1.0.0", result.RegistryVersion)
	assertEquals(t, "Embedded NIC 1 Boot", result.RegistryEntries.Attributes[0].DisplayName)
	assertEquals(t, "Enable Software RAID Support", result.RegistryEntries.Attributes[1].Value[1].ValueDisplayName)
	assertEquals(t, "BootMode", result.RegistryEntries.Dependencies[0].Dependency.MapFrom[0].MapFromAttribute)
	assertEquals(t, "EmbeddedSata", result.RegistryEntries.Dependencies[0].DependencyFor)
	assertEquals(t, "System Options", result.RegistryEntries.Menus[1].DisplayName)
}

var attributeRegistryUpperBoundBody = strings.NewReader(
	`{
		"@odata.type": "#AttributeRegistry.v1_3_7.AttributeRegistry",
		"Id": "BiosAttributeRegistryDell",
		"Language": "en",
		"Name": "Dell BIOS Attribute Registry",
		"OwningEntity": "Dell",
		"RegistryVersion": "1.0.0",
		"RegistryEntries": {
			"Attributes": [
				{
					"AttributeName": "SystemModelName",
					"DisplayName": "System Model Name",
					"Type": "String",
					"LowerBound": 0,
					"UpperBound": 18446744073709551615,
					"MaxLength": 40,
					"MinLength": 0,
					"ReadOnly": true,
					"Hidden": false
				},
				{
					"AttributeName": "ProcCores",
					"DisplayName": "Number of Cores per Processor",
					"Type": "Integer",
					"LowerBound": 1,
					"UpperBound": 128,
					"ReadOnly": false,
					"Hidden": false
				}
			]
		}
	}`)

// TestAttributeRegistryUpperBound tests that UpperBound correctly handles
// math.MaxUint64 values (18446744073709551615) as returned by Dell iDRAC
// for string-type BIOS attributes.
func TestAttributeRegistryUpperBound(t *testing.T) {
	var result AttributeRegistry
	err := json.NewDecoder(attributeRegistryUpperBoundBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "BiosAttributeRegistryDell", result.ID)

	if len(result.RegistryEntries.Attributes) != 2 {
		t.Fatalf("Expected 2 attributes, got %d", len(result.RegistryEntries.Attributes))
	}

	strAttr := result.RegistryEntries.Attributes[0]
	assertEquals(t, "SystemModelName", strAttr.AttributeName)
	if strAttr.UpperBound == nil {
		t.Fatal("Expected UpperBound to be set for SystemModelName")
	}
	if *strAttr.UpperBound != 18446744073709551615 {
		t.Errorf("Expected UpperBound 18446744073709551615, got %d", *strAttr.UpperBound)
	}
	if strAttr.LowerBound == nil {
		t.Fatal("Expected LowerBound to be set for SystemModelName")
	}
	if *strAttr.LowerBound != 0 {
		t.Errorf("Expected LowerBound 0, got %d", *strAttr.LowerBound)
	}

	intAttr := result.RegistryEntries.Attributes[1]
	assertEquals(t, "ProcCores", intAttr.AttributeName)
	if intAttr.UpperBound == nil {
		t.Fatal("Expected UpperBound to be set for ProcCores")
	}
	if *intAttr.UpperBound != 128 {
		t.Errorf("Expected UpperBound 128, got %d", *intAttr.UpperBound)
	}
}
