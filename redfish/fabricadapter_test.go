//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var fabricAdapterBody = `{
	"@odata.type": "#FabricAdapter.v1_5_2.FabricAdapter",
	"Id": "Bridge",
	"Name": "Gen-Z Bridge",
	"Manufacturer": "Contoso",
	"Model": "Gen-Z Bridge Model X",
	"PartNumber": "975999-001",
	"SparePartNumber": "152111-A01",
	"SKU": "Contoso 2-port Gen-Z Bridge",
	"SerialNumber": "2M220100SL",
	"ASICRevisionIdentifier": "A0",
	"ASICPartNumber": "53312",
	"ASICManufacturer": "Contoso",
	"FirmwareVersion": "7.4.10",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Ports": {
	  "@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/Ports"
	},
	"PCIeInterface": {
	  "MaxPCIeType": "Gen4",
	  "MaxLanes": 64,
	  "PCIeType": "Gen4",
	  "LanesInUse": 64
	},
	"UUID": "45724775-ed3b-2214-1313-9865200c1cc1",
	"Links": {
	  "Endpoints": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/GenZ/Endpoints/3"
		}
	  ]
	},
	"GenZ": {
	  "SSDT": {
		"@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/SSDT"
	  },
	  "MSDT": {
		"@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/MSDT"
	  },
	  "RequestorVCAT": {
		"@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/REQ-VCAT"
	  },
	  "ResponderVCAT": {
		"@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/RSP-VCAT"
	  },
	  "RITable": [
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E",
		"0x12",
		"0x3E"
	  ],
	  "PIDT": [
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568",
		"0x12234568"
	  ]
	},
	"@odata.id": "/redfish/v1/Systems/GenZ-example/FabricAdapters/1"
  }`

// TestFabricAdapter tests the parsing of FabricAdapter objects.
func TestFabricAdapter(t *testing.T) {
	var result FabricAdapter
	err := json.NewDecoder(strings.NewReader(fabricAdapterBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Bridge", result.ID)
	assertEquals(t, "Gen-Z Bridge", result.Name)
	assertEquals(t, "975999-001", result.PartNumber)
	assertEquals(t, "Contoso 2-port Gen-Z Bridge", result.SKU)
	assertEquals(t, "A0", result.ASICRevisionIdentifier)
	assertEquals(t, "53312", result.ASICPartNumber)
	assertEquals(t, "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/Ports", result.ports)
	assertEquals(t, "Gen4", string(result.PCIeInterface.PCIeType))
	assertEquals(t, "Gen4", string(result.PCIeInterface.MaxPCIeType))
	assertEquals(t, "/redfish/v1/Fabrics/GenZ/Endpoints/3", result.endpoints[0])
	assertEquals(t, "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/SSDT", result.GenZ.ssdt)
	assertEquals(t, "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/MSDT", result.GenZ.msdt)
	assertEquals(t, "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/REQ-VCAT", result.GenZ.requestorVCAT)
	assertEquals(t, "/redfish/v1/Systems/GenZ-example/FabricAdapters/1/RSP-VCAT", result.GenZ.responderVCAT)

	if len(result.GenZ.RITable) != 16 {
		t.Errorf("Unexpected RITable elements: %#v", result.GenZ.RITable)
	}

	if len(result.GenZ.PIDT) != 32 {
		t.Errorf("Unexpected PIDT elements: %#v", result.GenZ.PIDT)
	}

	if result.PCIeInterface.LanesInUse != 64 {
		t.Errorf("Unexpected PCIeInterface lanes in use: %d", result.PCIeInterface.LanesInUse)
	}
}
