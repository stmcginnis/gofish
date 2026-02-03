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
