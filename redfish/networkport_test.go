//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkPortBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#NetworkPort.NetworkPort",
		"@odata.type": "#NetworkPort.v1_2_2.NetworkPort",
		"@odata.id": "/redfish/v1/NetworkPort",
		"Id": "NetworkPort-1",
		"Name": "NetworkPortOne",
		"Description": "NetworkPort One",
		"ActiveLinkTechnology": "Ethernet",
		"AssociatedNetworkAddress": [
			"98:E7:43:00:01:0A"
		],
		"CurrentLinkSpeedMbps": 1000,
		"EEEEnabled": true,
		"FlowControlConfiguration": "TX_RX",
		"FlowControlStatus": "TX_RX",
		"LinkStatus": "Up",
		"MaxFrameSize": 900,
		"NetDevFuncMaxBWAlloc": [{
				"MaxBWAllocPercent": 100,
				"NetworkDeviceFunction": {
					"@odata.id": "/redfish/v1/Function/1"
				}
			},
			{
				"MaxBWAllocPercent": 100,
				"NetworkDeviceFunction": {
					"@odata.id": "/redfish/v1/Function/2"
				}
			}
		],
		"NetDevFuncMinBWAlloc": [{
				"MinBWAllocPercent": 25,
				"NetworkDeviceFunction": {
					"@odata.id": "/redfish/v1/Function/1"
				}
			},
			{
				"MinBWAllocPercent": 10,
				"NetworkDeviceFunction": {
					"@odata.id": "/redfish/v1/Function/2"
				}
			}
		],
		"NumberDiscoveredRemotePorts": 42,
		"PhysicalPortNumber": "10",
		"PortMaximumMTU": 100,
		"SignalDetected": true,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"SupportedEthernetCapabilities": [
			"WakeOnLAN",
			"EEE"
		],
		"SupportedLinkCapabilities": [{
			"AutoSpeedNegotiation": true,
			"CapableLinkSpeedMbps": [1000, 100],
			"LinkNetworkTechnology": "Ethernet",
			"LinkSpeedMbps": 1000
		}],
		"VendorId": "Vendor-ID",
		"WakeOnLANEnabled": false
	}`)

// TestNetworkPort tests the parsing of NetworkPort objects.
func TestNetworkPort(t *testing.T) {
	var result NetworkPort
	err := json.NewDecoder(networkPortBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NetworkPort-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "NetworkPortOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ActiveLinkTechnology != EthernetLinkNetworkTechnology {
		t.Errorf("Invalid active link technology: %s", result.ActiveLinkTechnology)
	}

	if result.CurrentLinkSpeedMbps != 1000 {
		t.Errorf("Invalid current link speed: %d", result.CurrentLinkSpeedMbps)
	}

	if result.FlowControlConfiguration != TXRXFlowControl {
		t.Errorf("Invalid flow control config: %s", result.FlowControlConfiguration)
	}

	if result.LinkStatus != UpPortLinkStatus {
		t.Errorf("Invalid link status: %s", result.LinkStatus)
	}

	if !result.SignalDetected {
		t.Error("Signal detected should be true")
	}
}
