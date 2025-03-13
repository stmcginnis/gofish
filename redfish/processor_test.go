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
		"@odata.context":"/redfish/v1/$metadata#Processor.Processor",
		"@odata.id":"/redfish/v1/Systems/System.Embedded.1/Processors/CPU.Socket.2",
		"@odata.type":"#Processor.v1_3_1.Processor",
		"Assembly":{
		   "@odata.id":"/redfish/v1/Chassis/System.Embedded.1/Assembly"
		},
		"CacheMemory": {
			"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Processors/CPU.Socket.2/CacheMemory"
		},
		"Description":"Represents the properties of a Processor attached to this System",
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
		},
		"Id":"CPU.Socket.2",
		"InstructionSet":"x86-64",
		"Links":{
		   "Chassis":{
			  "@odata.id":"/redfish/v1/Chassis/System.Embedded.1"
		   }
		},
		"Manufacturer":"Intel",
		"MaxSpeedMHz":4000,
		"Metrics": {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1/ProcessorMetrics"
		},
		"SubProcessors": {
			"@odata.id": "/redfish/v1/Systems/Self/Processors/CPU_1/SubProcessors"
		},
		"Model":"Intel(R) Xeon(R) Gold 6136 CPU @ 3.00GHz",
		"Name":"CPU 2",
		"Oem":{
		   "Dell":{
			  "DellProcessor":{
				 "@odata.context":"/redfish/v1/$metadata#DellProcessor.DellProcessor",
				 "@odata.id":"/redfish/v1/Dell/Systems/System.Embedded.1/Processors/DellProcessor/CPU.Socket.2",
				 "@odata.type":"#DellProcessor.v1_1_0.DellProcessor",
				 "CPUFamily":"Intel(R)Xeon(TM)",
				 "CPUStatus":"CPUEnabled",
				 "Cache1Associativity":"8-WaySet-Associative",
				 "Cache1ErrorMethodology":"Parity",
				 "Cache1InstalledSizeKB":768,
				 "Cache1Level":"L1",
				 "Cache1Location":"Internal",
				 "Cache1PrimaryStatus":"OK",
				 "Cache1SRAMType":"Unknown",
				 "Cache1SizeKB":768,
				 "Cache1Type":"Unified",
				 "Cache1WritePolicy":"WriteBack",
				 "Cache2Associativity":"16-WaySet-Associative",
				 "Cache2ErrorMethodology":"Single-bitECC",
				 "Cache2InstalledSizeKB":12288,
				 "Cache2Level":"L2",
				 "Cache2Location":"Internal",
				 "Cache2PrimaryStatus":"OK",
				 "Cache2SRAMType":"Unknown",
				 "Cache2SizeKB":12288,
				 "Cache2Type":"Unified",
				 "Cache2WritePolicy":"WriteBack",
				 "Cache3Associativity":"FullyAssociative",
				 "Cache3ErrorMethodology":"Single-bitECC",
				 "Cache3InstalledSizeKB":25344,
				 "Cache3Level":"L3",
				 "Cache3Location":"Internal",
				 "Cache3PrimaryStatus":"OK",
				 "Cache3SRAMType":"Unknown",
				 "Cache3SizeKB":25344,
				 "Cache3Type":"Unified",
				 "Cache3WritePolicy":"WriteBack",
				 "CurrentClockSpeedMhz":3000,
				 "ExternalBusClockSpeedMhz":10400,
				 "HyperThreadingCapable":"Yes",
				 "HyperThreadingEnabled":"Yes",
				 "LastSystemInventoryTime":"2020-01-07T00:53:28+00:00",
				 "LastUpdateTime":"2019-08-27T14:52:23+00:00",
				 "TurboModeCapable":"Yes",
				 "TurboModeEnabled":"Yes",
				 "VirtualizationTechnologyCapable":"Yes",
				 "VirtualizationTechnologyEnabled":"Yes",
				 "Volts":"1.8"
			  }
		   }
		},
		"ProcessorArchitecture":"x86",
		"ProcessorId":{
		   "EffectiveFamily":"179",
		   "EffectiveModel":"85",
		   "IdentificationRegisters":"0x00050654",
		   "MicrocodeInfo":null,
		   "Step":"4",
		   "VendorId":"GenuineIntel"
		},
		"ProcessorType":"CPU",
		"Socket":"CPU.Socket.2",
		"Status":{
		   "Health":"OK",
		   "State":"Enabled"
		},
		"TotalCores":12,
		"TotalThreads":24
	 }`)

var processorBodyWithMaxSpeedMHzString = strings.NewReader(
	`{
		"@odata.context":"/redfish/v1/$metadata#Processor.Processor",
		"@odata.id":"/redfish/v1/Systems/System.Embedded.1/Processors/CPU.Socket.2",
		"@odata.type":"#Processor.v1_3_1.Processor",
		"Assembly":{
		   "@odata.id":"/redfish/v1/Chassis/System.Embedded.1/Assembly"
		},
		"Description":"Represents the properties of a Processor attached to this System",
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
		},
		"Id":"CPU.Socket.2",
		"InstructionSet":"x86-64",
		"Links":{
		   "Chassis":{
			  "@odata.id":"/redfish/v1/Chassis/System.Embedded.1"
		   }
		},
		"Manufacturer":"Intel",
		"MaxSpeedMHz":"4000",
		"Metrics": {
			"@odata.id": "/redfish/v1/Systems/1/Processors/1/ProcessorMetrics"
		},
		"Model":"Intel(R) Xeon(R) Gold 6136 CPU @ 3.00GHz",
		"Name":"CPU 2",
		"Oem":{
		   "Dell":{
			  "DellProcessor":{
				 "@odata.context":"/redfish/v1/$metadata#DellProcessor.DellProcessor",
				 "@odata.id":"/redfish/v1/Dell/Systems/System.Embedded.1/Processors/DellProcessor/CPU.Socket.2",
				 "@odata.type":"#DellProcessor.v1_1_0.DellProcessor",
				 "CPUFamily":"Intel(R)Xeon(TM)",
				 "CPUStatus":"CPUEnabled",
				 "Cache1Associativity":"8-WaySet-Associative",
				 "Cache1ErrorMethodology":"Parity",
				 "Cache1InstalledSizeKB":768,
				 "Cache1Level":"L1",
				 "Cache1Location":"Internal",
				 "Cache1PrimaryStatus":"OK",
				 "Cache1SRAMType":"Unknown",
				 "Cache1SizeKB":768,
				 "Cache1Type":"Unified",
				 "Cache1WritePolicy":"WriteBack",
				 "Cache2Associativity":"16-WaySet-Associative",
				 "Cache2ErrorMethodology":"Single-bitECC",
				 "Cache2InstalledSizeKB":12288,
				 "Cache2Level":"L2",
				 "Cache2Location":"Internal",
				 "Cache2PrimaryStatus":"OK",
				 "Cache2SRAMType":"Unknown",
				 "Cache2SizeKB":12288,
				 "Cache2Type":"Unified",
				 "Cache2WritePolicy":"WriteBack",
				 "Cache3Associativity":"FullyAssociative",
				 "Cache3ErrorMethodology":"Single-bitECC",
				 "Cache3InstalledSizeKB":25344,
				 "Cache3Level":"L3",
				 "Cache3Location":"Internal",
				 "Cache3PrimaryStatus":"OK",
				 "Cache3SRAMType":"Unknown",
				 "Cache3SizeKB":25344,
				 "Cache3Type":"Unified",
				 "Cache3WritePolicy":"WriteBack",
				 "CurrentClockSpeedMhz":3000,
				 "ExternalBusClockSpeedMhz":10400,
				 "HyperThreadingCapable":"Yes",
				 "HyperThreadingEnabled":"Yes",
				 "LastSystemInventoryTime":"2020-01-07T00:53:28+00:00",
				 "LastUpdateTime":"2019-08-27T14:52:23+00:00",
				 "TurboModeCapable":"Yes",
				 "TurboModeEnabled":"Yes",
				 "VirtualizationTechnologyCapable":"Yes",
				 "VirtualizationTechnologyEnabled":"Yes",
				 "Volts":"1.8"
			  }
		   }
		},
		"ProcessorArchitecture":"x86",
		"ProcessorId":{
		   "EffectiveFamily":"179",
		   "EffectiveModel":"85",
		   "IdentificationRegisters":"0x00050654",
		   "MicrocodeInfo":null,
		   "Step":"4",
		   "VendorId":"GenuineIntel"
		},
		"ProcessorType":"CPU",
		"Socket":"CPU.Socket.2",
		"Status":{
		   "Health":"OK",
		   "State":"Enabled"
		},
		"TotalCores":12,
		"TotalThreads":24
	 }`)

var invalidProcessorBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Processor.Processor(*)",
		"@odata.etag": "W/\"1604509181\"",
		"@odata.id": "/redfish/v1/Systems/Self/Processors/1",
		"@odata.type": "#Processor.v1_0_3.Processor",
		"Id": "1",
		"InstructionSet": "X86-64",
		"Manufacturer": "Intel(R) Corporation",
		"MaxSpeedMHz": "",
		"Model": "Intel Xeon",
		"Name": "Processor1",
		"Oem": {
		  "Intel_RackScale": {
			"@odata.type": "#Intel.Oem.Processor",
			"Brand": "E5",
			"Capabilities": [
			  "fpu",
			  "vme",
			  "de",
			  "pse",
			  "tsc",
			  "msr",
			  "pae",
			  "mce",
			  "cx8",
			  "apic",
			  "sep",
			  "mtrr",
			  "pge",
			  "mca",
			  "cmov",
			  "pat",
			  "pse-36",
			  "clfsh",
			  "ds",
			  "acpi",
			  "mmx",
			  "fxsr",
			  "sse",
			  "sse2",
			  "ss",
			  "htt",
			  "tm",
			  "pbe"
			]
		  },
		  "Quanta_RackScale": {
			"Version": "Intel(R) Xeon(R) Gold 6242 CPU @ 2.80GHz"
		  }
		},
		"ProcessorArchitecture": "x86",
		"ProcessorId": {
		  "EffectiveFamily": "0x6",
		  "EffectiveModel": "0x55",
		  "IdentificationRegisters": "0xbfebfbff00050657",
		  "MicrocodeInfo": "0x50024",
		  "Step": "0x7",
		  "VendorId": "GenuineIntel"
		},
		"ProcessorType": "CPU",
		"Socket": "CPU_0",
		"Status": {
		  "Health": "OK",
		  "State": "Enabled"
		},
		"TotalCores": 16,
		"TotalThreads": 32
	  }`)

// TestProcessor tests the parsing of Processor objects.
func TestProcessor(t *testing.T) {
	var result Processor
	err := json.NewDecoder(processorBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CPU.Socket.2" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "CPU 2" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.assembly != "/redfish/v1/Chassis/System.Embedded.1/Assembly" {
		t.Errorf("Invalid assembly link: %s", result.assembly)
	}

	if result.cacheMemory != "/redfish/v1/Systems/System.Embedded.1/Processors/CPU.Socket.2/CacheMemory" {
		t.Errorf("Invalid cache memory link: %s", result.cacheMemory)
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

// TestMaxSpeedMHzString tests the parsing of Processor objects.
func TestMaxSpeedMHzString(t *testing.T) {
	var result Processor
	err := json.NewDecoder(processorBodyWithMaxSpeedMHzString).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CPU.Socket.2" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "CPU 2" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.assembly != "/redfish/v1/Chassis/System.Embedded.1/Assembly" {
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

	if result.MaxSpeedMHz != 4000 {
		t.Errorf("Expected MaxSpeedMhz to be 4000 but got %f", result.MaxSpeedMHz)
	}
}

// TestNonconformingProcessor tests the parsing of Processor objects from certain
// Dell implementations that do not fully conform to the spec.
func TestNonconformingProcessor(t *testing.T) {
	var result Processor
	err := json.NewDecoder(invalidProcessorBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.MaxSpeedMHz != 0 {
		t.Errorf("Expected MaxSpeedMhz to be 0 but got %f", result.MaxSpeedMHz)
	}
}
