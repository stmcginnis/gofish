//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var bootOptionBody = `{
	"@odata.context": "/redfish/v1/$metadata#BootOption.BootOption",
	"@odata.etag": "W/\"A3A6BF43\"",
	"@odata.id": "/redfish/v1/Systems/1/BootOptions/1/",
	"@odata.type": "#BootOption.v1_0_1.BootOption",
	"Id": "1",
	"Alias": "None",
	"BootOptionReference": "Boot0015",
	"DisplayName": "URL File : http://assets.example.com/ipxe.efi (IPv4)",
	"Name": "Boot Option",
	"UefiDevicePath": "IPv4(0.0.0.0)/Uri(http://assets.example.com/ipxe.efi)"
}`

func TestBootOption(t *testing.T) {
	var result BootOption
	err := json.NewDecoder(strings.NewReader(bootOptionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Boot Option" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Alias != "None" {
		t.Errorf("Received invalid alias: %s", result.Alias)
	}

	if result.BootOptionReference != "Boot0015" {
		t.Errorf("Received invalid bootoptionreference: %s", result.BootOptionReference)
	}

	if result.DisplayName != "URL File : http://assets.example.com/ipxe.efi (IPv4)" {
		t.Errorf("Received invalid displayname: %s", result.DisplayName)
	}

	if result.UefiDevicePath != "IPv4(0.0.0.0)/Uri(http://assets.example.com/ipxe.efi)" {
		t.Errorf("Received invalid uefidevicepath: %s", result.UefiDevicePath)
	}
}
