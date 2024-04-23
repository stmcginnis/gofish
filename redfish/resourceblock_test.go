//
// SPDX-License-Identifier: BSD-3-Clause
//

//nolint:dupl
package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var resourceBlockComputeBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ResourceBlock.ResourceBlock",
		"@odata.etag": "\"1712866586\"",
		"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/ComputeBlock",
		"@odata.type": "#ResourceBlock.v1_3_2.ResourceBlock",
		"CompositionStatus": {
		  "CompositionState": "Unused",
		  "MaxCompositions": 1,
		  "Reserved": false
		},
		"Description": "ComputeBlock",
		"Id": "ComputeBlock",
		"Links": {
		  "Chassis": [
			{
			  "@odata.id": "/redfish/v1/Chassis/BMC_0"
			}
		  ],
		  "Zones": [
			{
			  "@odata.id": "/redfish/v1/CompositionService/ResourceZones/1"
			}
		  ]
		},
		"Memory": [
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/Memory/DevType2_DIMM0"
		  }
		],
		"Memory@odata.count": 1,
		"Name": "ComputeBlock",
		"Processors": [
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/Processors/DevType1_CPU0"
		  }
		],
		"Processors@odata.count": 1,
		"ResourceBlockType": [
		  "Compute"
		],
		"Status": {
		  "Health": "OK",
		  "HealthRollup": "OK",
		  "State": "Enabled"
		}
	  }`)

var resourceBlockDriveBody = strings.NewReader(
	`{
			"@odata.context": "/redfish/v1/$metadata#ResourceBlock.ResourceBlock",
			"@odata.etag": "\"1712866586\"",
			"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/DrivesBlock",
			"@odata.type": "#ResourceBlock.v1_3_2.ResourceBlock",
			"CompositionStatus": {
			  "CompositionState": "Unused",
			  "MaxCompositions": 1,
			  "Reserved": false
			},
			"Description": "DrivesBlock",
			"Drives": [
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/NVMe_Device0_NSID1"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device1_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device2_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device3_Port2"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device4_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device5_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device6_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device7_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device8_Port4"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0/Drives/USB_Device9_Port4"
			  }
			],
			"Drives@odata.count": 10,
			"Id": "DrivesBlock",
			"Links": {
			  "Chassis": [
				{
				  "@odata.id": "/redfish/v1/Chassis/BMC_0"
				}
			  ],
			  "Zones": [
				{
				  "@odata.id": "/redfish/v1/CompositionService/ResourceZones/1"
				}
			  ]
			},
			"Name": "DrivesBlock",
			"ResourceBlockType": [
			  "Storage"
			],
			"SimpleStorage": [
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/SimpleStorage/0"
			  },
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/SimpleStorage/1"
			  }
			],
			"SimpleStorage@odata.count": 2,
			"Status": {
			  "Health": "OK",
			  "HealthRollup": "OK",
			  "State": "Enabled"
			},
			"Storage": [
			  {
				"@odata.id": "/redfish/v1/Systems/System_0/Storage/StorageUnit_0"
			  }
			],
			"Storage@odata.count": 1
		  }`)

var resourceBlockNetworkBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ResourceBlock.ResourceBlock",
		"@odata.etag": "\"1712866586\"",
		"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks/NetworkBlock",
		"@odata.type": "#ResourceBlock.v1_3_2.ResourceBlock",
		"CompositionStatus": {
		  "CompositionState": "Unused",
		  "MaxCompositions": 1,
		  "Reserved": false
		},
		"Description": "NetworkBlock",
		"EthernetInterfaces": [
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/EthernetInterfaces/EthernetInterface0"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/EthernetInterfaces/EthernetInterface1"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/EthernetInterfaces/EthernetInterface2"
		  },
		  {
			"@odata.id": "/redfish/v1/Systems/System_0/EthernetInterfaces/EthernetInterface3"
		  }
		],
		"EthernetInterfaces@odata.count": 4,
		"Id": "NetworkBlock",
		"Links": {
		  "Chassis": [
			{
			  "@odata.id": "/redfish/v1/Chassis/BMC_0"
			}
		  ],
		  "Zones": [
			{
			  "@odata.id": "/redfish/v1/CompositionService/ResourceZones/1"
			}
		  ]
		},
		"Name": "NetworkBlock",
		"ResourceBlockType": [
		  "Network"
		],
		"Status": {
		  "Health": "OK",
		  "HealthRollup": "OK",
		  "State": "Enabled"
		}
	  }`)

// TestComputeResourceBlock tests the parsing of Compute ResourceBlock objects.
func TestComputeResourceBlock(t *testing.T) {
	var result ResourceBlock
	err := json.NewDecoder(resourceBlockComputeBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ComputeBlock", result.ID)
	assertEquals(t, "ComputeBlock", result.Description)
	assertEquals(t, "ComputeBlock", result.Name)

	if len(result.ResourceBlockType) != 1 {
		t.Errorf("Expected 1 ResourceBlockType, got: %#v", result.ResourceBlockType)
	}

	assertEquals(t, "Compute", string(result.ResourceBlockType[0]))

	if len(result.processors) != 1 {
		t.Errorf("Expected 1 processor, got: %#v", result.processors)
	}

	if len(result.memory) != 1 {
		t.Errorf("Expected 1 memory, got: %#v", result.memory)
	}

	if len(result.drives) != 0 {
		t.Errorf("Expected 0 drives, got: %#v", result.drives)
	}

	if len(result.storage) != 0 {
		t.Errorf("Expected 0 storage links, got: %#v", result.storage)
	}

	if len(result.simpleStorage) != 0 {
		t.Errorf("Expected 0 simple storage links, got: %#v", result.simpleStorage)
	}

	if len(result.ethernetInterfaces) != 0 {
		t.Errorf("Expected 0 ethernet interfaces, got: %#v", result.ethernetInterfaces)
	}

	assertEquals(t, "Unused", string(result.CompositionStatus.CompositionState))
}

// TestDrivesResourceBlock tests the parsing of Storage ResourceBlock objects.
func TestDrivesResourceBlock(t *testing.T) {
	var result ResourceBlock
	err := json.NewDecoder(resourceBlockDriveBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "DrivesBlock", result.ID)
	assertEquals(t, "DrivesBlock", result.Description)
	assertEquals(t, "DrivesBlock", result.Name)

	if len(result.ResourceBlockType) != 1 {
		t.Errorf("Expected 1 ResourceBlockType, got: %#v", result.ResourceBlockType)
	}

	assertEquals(t, "Storage", string(result.ResourceBlockType[0]))

	if len(result.processors) != 0 {
		t.Errorf("Expected 0 processor, got: %#v", result.processors)
	}

	if len(result.memory) != 0 {
		t.Errorf("Expected 0 memory, got: %#v", result.memory)
	}

	if len(result.drives) != 10 {
		t.Errorf("Expected 10 drives, got: %#v", result.drives)
	}

	if len(result.storage) != 1 {
		t.Errorf("Expected 1 storage links, got: %#v", result.storage)
	}

	if len(result.simpleStorage) != 2 {
		t.Errorf("Expected 2 simple storage links, got: %#v", result.simpleStorage)
	}

	if len(result.ethernetInterfaces) != 0 {
		t.Errorf("Expected 0 ethernet interfaces, got: %#v", result.ethernetInterfaces)
	}
}

// TestNetworkResourceBlock tests the parsing of Network ResourceBlock objects.
func TestNetworkResourceBlock(t *testing.T) {
	var result ResourceBlock
	err := json.NewDecoder(resourceBlockNetworkBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "NetworkBlock", result.ID)
	assertEquals(t, "NetworkBlock", result.Description)
	assertEquals(t, "NetworkBlock", result.Name)

	if len(result.ResourceBlockType) != 1 {
		t.Errorf("Expected 1 ResourceBlockType, got: %#v", result.ResourceBlockType)
	}

	assertEquals(t, "Network", string(result.ResourceBlockType[0]))

	if len(result.processors) != 0 {
		t.Errorf("Expected 0 processor, got: %#v", result.processors)
	}

	if len(result.memory) != 0 {
		t.Errorf("Expected 0 memory, got: %#v", result.memory)
	}

	if len(result.drives) != 0 {
		t.Errorf("Expected 0 drives, got: %#v", result.drives)
	}

	if len(result.storage) != 0 {
		t.Errorf("Expected 0 storage links, got: %#v", result.storage)
	}

	if len(result.simpleStorage) != 0 {
		t.Errorf("Expected 0 simple storage links, got: %#v", result.simpleStorage)
	}

	if len(result.ethernetInterfaces) != 4 {
		t.Errorf("Expected 4 ethernet interfaces, got: %#v", result.ethernetInterfaces)
	}
}
