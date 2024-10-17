//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"testing"
)

var dumpServiceBody = `{
  "@odata.type": "#DumpService.v1_0_2.DumpService",
  "@odata.id": "/redfish/v1/Oem/Supermicro/DumpService",
  "Id": "DumpService",
  "Name": "Dump Service",
  "Description": "Dump Service",
  "Dumps": {
    "@odata.id": "/redfish/v1/Oem/Supermicro/DumpService/Dumps"
  },
  "Actions": {
    "Oem": {},
    "#SmcDumpService.CreateDump": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Actions/SmcDumpService.CreateDump",
      "@Redfish.ActionInfo": "/redfish/v1/Oem/Supermicro/DumpService/CreateDumpActionInfo"
    },
    "#SmcDumpService.DeleteAll": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Actions/SmcDumpService.DeleteAll"
    },
    "#OemDumpService.Collect": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Actions/OemDumpService.Collect",
      "@Redfish.ActionInfo": "/redfish/v1/Oem/Supermicro/DumpService/CollectActionInfo"
    }
  },
  "@odata.etag": "\"697e2b05b6d5d49940bc0fd68803608b\""
}`

var dumpBody = `{
  "@odata.type": "#Dump.v1_1_0.Dump",
  "@odata.id": "/redfish/v1/Oem/Supermicro/DumpService/Dumps/AttestationDump",
  "Id": "AttestationDump",
  "Name": "AttestationDump",
  "Description": "Supermicro Attestation Dump Service",
  "AttestationFile": [
    "attd_BS=OM243S046922_2024-09-18T18:10:52-07:00.bin",
    "attd_BS=OM243S046922_MB_2024-05-23T13:54:08+08:00.bin"
  ],
  "Actions": {
    "Oem": {},
    "#SmcAttestationDump.Generate": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Dumps/AttestationDump/Actions/SmcAttestationDump.Generate"
    },
    "#SmcAttestationDump.Download": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Dumps/AttestationDump/Actions/SmcAttestationDump.Download"
    },
    "#SmcAttestationDump.Delete": {
      "target": "/redfish/v1/Oem/Supermicro/DumpService/Dumps/AttestationDump/Actions/SmcAttestationDump.Delete"
    }
  },
  "@odata.etag": "\"0e25db1a4c2d3d4a28cda306e1b29abe\""
}`

// TestSmcDumpService tests the parsing of the DumpService oem field
func TestSmcDumpService(t *testing.T) {
	ds := &DumpService{}
	if err := json.Unmarshal([]byte(dumpServiceBody), ds); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	if ds.ID != "DumpService" {
		t.Errorf("unexpected ID: %s", ds.ID)
	}

	if ds.collectTarget != "/redfish/v1/Oem/Supermicro/DumpService/Actions/OemDumpService.Collect" {
		t.Errorf("unexpected install target: %s", ds.collectTarget)
	}

	if ds.createDumpTarget != "/redfish/v1/Oem/Supermicro/DumpService/Actions/SmcDumpService.CreateDump" {
		t.Errorf("unexpected ssl cert link: %s", ds.createDumpTarget)
	}

	if ds.deleteAllTarget != "/redfish/v1/Oem/Supermicro/DumpService/Actions/SmcDumpService.DeleteAll" {
		t.Errorf("unexpected ipmi config link: %s", ds.deleteAllTarget)
	}

	if ds.dumps != "/redfish/v1/Oem/Supermicro/DumpService/Dumps" {
		t.Errorf("unexpected dumps link: %s", ds.dumps)
	}
}

// TestSmcDump tests the parsing of the Dumpobject.
func TestSmcDump(t *testing.T) {
	ds := &Dump{}
	if err := json.Unmarshal([]byte(dumpBody), ds); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	if ds.ID != "AttestationDump" {
		t.Errorf("unexpected ID: %s", ds.ID)
	}

	if len(ds.AttestationFile) != 2 {
		t.Errorf("unexpected number of attestation files: %d", len(ds.AttestationFile))
	}
}
