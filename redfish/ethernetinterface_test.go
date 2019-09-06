//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var ethernetInterfaceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#EthernetInterface.EthernetInterface",
		"@odata.id": "/redfish/v1/Systems/System-1/EthernetInterfaces/NIC-0",
		"@odata.type": "#EthernetInterface.v1_3_0.EthernetInterface",
		"AutoNeg": true,
		"Description": "Ethernet Interface",
		"FQDN": "default.local",
		"FullDuplex": true,
		"HostName": "default",
		"IPv4Addresses": [
			{
				"Address": "172.16.3.39",
				"AddressOrigin": "IPv4LinkLocal",
				"Gateway": "0.0.0.0",
				"SubnetMask": "255.255.0.0"
			}
		],
		"Id": "NIC-0",
		"InterfaceEnabled": true,
		"LinkStatus": "LinkUp",
		"Links": {
			"Chassis": {
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			}
		},
		"MACAddress": "f6:a9:26:e3:e6:32",
		"MTUSize": 1500,
		"Name": "Ethernet Interface",
		"NameServers": [
			"8.8.8.8"
		],
		"PermanentMACAddress": "f6:a9:26:e3:e6:32",
		"SpeedMbps": 10000,
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"VLAN": {
			"VLANId": 0
		}
	}`)

// TestEthernetInterface tests the parsing of EthernetInterface objects.
func TestEthernetInterface(t *testing.T) {
	var result EthernetInterface
	err := json.NewDecoder(ethernetInterfaceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NIC-0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Ethernet Interface" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.AutoNeg {
		t.Error("Auto negotiate should be True")
	}

	if !result.FullDuplex {
		t.Error("Full duplex should be True")
	}

	if len(result.IPv4Addresses) != 1 {
		t.Errorf("Expected number of IPv4Addresses, got: %d", len(result.IPv4Addresses))
	}

	if result.IPv4Addresses[0].AddressOrigin != IPv4LinkLocalIPv4AddressOrigin {
		t.Errorf("Should have received IPv4LinkLocal address origin, got: %s",
			result.IPv4Addresses[0].AddressOrigin)
	}

	if len(result.IPv6Addresses) != 0 {
		t.Errorf("Should be 0 IPv6 addresses, got: %d", len(result.IPv4Addresses))
	}

	if result.LinkStatus != LinkUpLinkStatus {
		t.Errorf("Should have received LinkUp status, got %s", result.LinkStatus)
	}

	if result.SpeedMbps != 10000 {
		t.Errorf("Expected 10000 speed, got %d", result.SpeedMbps)
	}
}
