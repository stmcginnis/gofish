//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var processorBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Processor.Processor",
		"@odata.type": "#Processor.1_0_0.Processor",
		"@odata.id": "/redfish/v1/Systems/System-1/Processors/CPU0",
		"Name": "Processor",
		"Id": "CPU0",
		"AccelerationFunctions": {
			"@odata.id": "/redfish/v1/Systems/System-1/Processors/CPU0/Functions"
		},
		"Assembly": {
			"@odata.id": "/redfish/v1/Assembly/1"
		},
		"Description": "Processor one",
		"FPGA": {
			"ExternalInterfaces": [{
				"Ethernet": {
					"@odata.id": "/redfish/v1/Systems/System-1/Eth/1"
				},
				"InterfaceType": "Ethernet",
				"PCIe": {
					"@odata.id": "/redfish/v1/Systems/System-1/PCIe/1"
				}
			}],
			"FirmwareId": "A1",
			"FirmwareManufacturer": "Acme",
			"FirmwareVersion": "A01",
			"FpgaType": "Discrete",
			"HostInterface": {
				"Ethernet": {
					"@odata.id": "/redfish/v1/Systems/System-1/Eth/1"
				},
				"InterfaceType": "Ethernet",
				"PCIe": {
					"@odata.id": "/redfish/v1/Systems/System-1/PCIe/1"
				}
			},
			"Model": "FPGA-1",
			"PCIeVirtualFunctions": 4,
			"ProgrammableFromHost": true,
			"ReconfigurationSlots": [{
					"AccelerationFunction": {
						"@odata.id": "/redfish/v1/Functions/1"
					},
					"ProgrammableFromHost": true,
					"SlotId": "1",
					"UUID": "e703fdd1-e386-48cb-b503-097a2b35d3e1"
				},
				{
					"AccelerationFunction": {
						"@odata.id": "/redfish/v1/Functions/2"
					},
					"ProgrammableFromHost": true,
					"SlotId": "2",
					"UUID": "e703fdd1-e386-48cb-beef-097a2b35d3e1"
				}
			]
		}
	}`)

// TestProcessor tests the parsing of Processor objects.
func TestProcessor(t *testing.T) {
	var result Processor
	err := json.NewDecoder(processorBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CPU0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Processor" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.assembly != "/redfish/v1/Assembly/1" {
		t.Errorf("Invalid assembly link: %s", result.assembly)
	}

	if result.FPGA.FpgaType != DiscreteFpgaType {
		t.Errorf("Invalid FPGA type: %s", result.FPGA.FpgaType)
	}

	if result.FPGA.PCIeVirtualFunctions != 4 {
		t.Errorf("Invalid virtual function count: %d", result.FPGA.PCIeVirtualFunctions)
	}

	if len(result.FPGA.ReconfigurationSlots) != 2 {
		t.Errorf("Expected 2 ReconfigurationSlots, got %d", len(result.FPGA.ReconfigurationSlots))
	}
}
