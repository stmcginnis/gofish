//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var ethernetInterfaceBody = `{
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
	}`

// TestEthernetInterface tests the parsing of EthernetInterface objects.
func TestEthernetInterface(t *testing.T) {
	var result EthernetInterface
	err := json.NewDecoder(strings.NewReader(ethernetInterfaceBody)).Decode(&result)

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

// TestEthernetInterfaceUpdate tests the Update call.
func TestEthernetInterfaceUpdate(t *testing.T) {
	var result EthernetInterface
	err := json.NewDecoder(strings.NewReader(ethernetInterfaceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.AutoNeg = false
	result.FQDN = "test.local"
	result.FullDuplex = true
	result.HostName = "test"
	result.InterfaceEnabled = false
	result.MACAddress = "de:ad:de:ad:de:ad"
	result.MTUSize = 9216
	result.SpeedMbps = 1000
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AutoNeg:false") {
		t.Errorf("Unexpected AutoNeg update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "FQDN:test.local") {
		t.Errorf("Unexpected FQDN update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "FullDuplex") {
		t.Errorf("Unexpected FullDuplex in update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "HostName:test") {
		t.Errorf("Unexpected HostName update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "InterfaceEnabled:false") {
		t.Errorf("Unexpected InterfaceEnabled update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "MACAddress:de:ad:de:ad:de:ad") {
		t.Errorf("Unexpected MACAddress update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "MTUSize:9216") {
		t.Errorf("Unexpected MTUSize update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "SpeedMbps:1000") {
		t.Errorf("Unexpected SpeedMbps update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "FullDuplex") {
		t.Errorf("Unexpected update for FullDuplex in payload: %s", calls[0].Payload)
	}
}

var ethernetInterfaceIPv6Body = `{
		"@odata.context": "/redfish/v1/$metadata#EthernetInterface.EthernetInterface",
		"@odata.id": "/redfish/v1/Systems/System-1/EthernetInterfaces/NIC-0",
		"@odata.type": "#EthernetInterface.v1_3_0.EthernetInterface",
		"IPv6Addresses": [
			{
				"Address": "FE80::B67A:F1FF:FECF:6462",
				"AddressOrigin": "SLAAC",
				"AddressState": "Preferred",
				"PrefixLength": 64
			},
			{
				"Address": "FDE1:53BA:E9A0:DE41::1649",
				"AddressOrigin": "DHCP",
				"AddressState": "Preferred",
				"PrefixLength": 128
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
	}`

// TestEthernetInterface tests the parsing of EthernetInterface objects.
func TestEthernetInterfaceIPv6(t *testing.T) {
	var result EthernetInterface
	err := json.NewDecoder(strings.NewReader(ethernetInterfaceIPv6Body)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NIC-0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Ethernet Interface" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.IPv6Addresses) != 2 {
		t.Errorf("Should be 2 IPv6 addresses, got: %d", len(result.IPv4Addresses))
	}

	if result.IPv6Addresses[0].PrefixLength != 64 {
		t.Errorf("The 1st IPv6 address's prefix length should be 64, got: %d", result.IPv6Addresses[0].PrefixLength)
	}

	if result.IPv6Addresses[1].PrefixLength != 128 {
		t.Errorf("The 3nd IPv6 address's prefix length should be 128, got: %d", result.IPv6Addresses[1].PrefixLength)
	}
}
