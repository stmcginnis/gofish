//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var fixedBootOrderBody = `{
  "@odata.type": "#SmcFixedBootOrder.v1_0_0.SmcFixedBootOrder",
  "@odata.id": "/redfish/v1/Systems/1/Oem/Supermicro/FixedBootOrder",
  "Id": "FixedBootOrder",
  "Name": "Fixed Boot Order",
  "BootModeSelected": "UEFI",
  "FixedBootOrder": [
    "UEFI CD/DVD",
    "UEFI USB CD/DVD",
    "UEFI Network:(B83/D0/F0) UEFI PXE IPv4 Intel(R) Ethernet Controller X550 - 905A0839F618",
    "UEFI Hard Disk:ubuntu (HDD,Port:900)",
    "UEFI USB Hard Disk",
    "UEFI USB Key",
    "UEFI USB Floppy",
    "UEFI USB Lan",
    "UEFI AP:UEFI: Built-in EFI Shell"
  ],
  "FixedBootOrderDisabledItem": [
    "Disabled"
  ],
  "UEFINetwork": [
    "(B83/D0/F0) UEFI PXE IPv4 Intel(R) Ethernet Controller X550 - 905A0839F618",
    "(B83/D0/F1) UEFI PXE IPv4 Intel(R) Ethernet Controller X550 - 905A0839F619",
    "(B210/D0/F0) UEFI PXE IPv4 Nvidia Network Adapter - 5C:25:73:60:C5:D8 - 5C257360C5D8",
    "(B210/D0/F1) UEFI PXE IPv4 Nvidia Network Adapter - 5C:25:73:60:C5:D9 - 5C257360C5D9"
  ],
  "UEFINetworkDisabledItem": [
    "Disabled"
  ],
  "UEFIHardDisk": [
    "ubuntu (HDD,Port:900)"
  ],
  "UEFIHardDiskDisabledItem": [
    "Disabled"
  ],
  "UEFIAP": [
    "UEFI: Built-in EFI Shell"
  ],
  "UEFIAPDisabledItem": [
    "Disabled"
  ],
  "@odata.etag": "\"120671877241e67076141a0d63fc7c7b\""
}`

// TestFixedBootOrder tests the parsing of FixedBootOrder objects.
func TestFixedBootOrder(t *testing.T) {
	var result FixedBootOrder
	err := json.NewDecoder(strings.NewReader(fixedBootOrderBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "FixedBootOrder" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.BootModeSelected != "UEFI" {
		t.Errorf("Invalid BootModeSelected: %s", result.BootModeSelected)
	}

	if len(result.FixedBootOrder) != 9 {
		t.Errorf("Expected 9 fixed boot order entries, got %d", len(result.FixedBootOrder))
	}
}
