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

var endpointBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Endpoint.Endpoint",
		"@odata.type": "#Endpoint.v1_0_0.Endpoint",
		"@odata.id": "/redfish/v1/Endpoint",
		"Id": "Endpoint-1",
		"Name": "EndpointOne",
		"Description": "Endpoint One",
		"ConnectedEntities": [{
				"EntityLink": {
					"@odata.id": "/redfish/v1/EntityLink"
				},
				"EntityPciId": {
					"ClassCode": "01:00",
					"DeviceID": "0497",
					"FunctionNumber": "00",
					"SubsystemID": "103",
					"SubsystemVendorID": "7a03",
					"VendorID": "Joe's Storage"
				},
				"EntityRole": "Both",
				"EntityType": "StorageInitiator",
				"Identifiers": [{
						"DurableName": "iqn.1994-05.com.redhat:b87750935dd5",
						"DurableNameFormat": "iQN"
					}
				]
			}
		],
		"EndpointProtocol": "NVMe",
		"HostReservationMemoryBytes": 8589934592,
		"IPTransportDetails": [{
				"IPv4Address": "127.0.0.1",
				"IPv6Address": "2001:db8:85a3:0:0:8a2e:370:7334",
				"Port": 8080,
				"TransportProtocol": "iSCSI"
			}
		],
		"Identifiers": [{
				"DurableName": "iqn.1994-05.com.redhat:b87750935dd5",
				"DurableNameFormat": "iQN"
			}
		],
		"Links": {
			"MutuallyExclusiveEndpoints": [
				{
					"@odata.id": "/redfish/v1/Endpoints/Endpoint-1"
				}
			],
			"MutuallyExclusiveEndpointsCount": 1,
			"NetworkDeviceFunction": [{
				"@odata.id": "/redfish/v1/NetworkDeviceFunction/1"
				}
			],
			"NetworkDeviceFunctionCount": 1,
			"Ports": [{
				"@odata.id": "/redfish/v1/Ports/1"
				}
			],
			"PortCount": 1
		},
		"PciID": {
			"ClassCode": "01:00",
			"DeviceID": "0497",
			"FunctionNumber": "00",
			"SubsystemID": "103",
			"SubsystemVendorID": "7a03",
			"VendorID": "Joe's Storage"
		},
		"Redundancy": [{
				"MaxNumSupported": 1,
				"MemberId": "1",
				"MinNumNeeded": 1,
				"Mode": "NotRedundant",
				"Name": "Kevin",
				"RedundancyEnabled": false,
				"RedundancySet": [],
				"RedundancySet@odata.count": 0,
				"Status": {
					"State": "Enabled",
					"Health": "OK"
				}
			}
		],
		"Redundancy@odata.count": 1,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestEndpoint tests the parsing of Endpoint objects.
func TestEndpoint(t *testing.T) {
	var result Endpoint
	err := json.NewDecoder(endpointBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Endpoint-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "EndpointOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.ConnectedEntities) != 1 {
		t.Errorf("Expected on connected entity, got: %d", len(result.ConnectedEntities))
	}

	if result.EndpointProtocol != common.NVMeProtocol {
		t.Errorf("Received endpoint protocol: %s", result.EndpointProtocol)
	}

	if result.HostReservationMemoryBytes != 8589934592 {
		t.Errorf("Received host reservation memory bytes: %d", result.HostReservationMemoryBytes)
	}

	if len(result.IPTransportDetails) != 1 {
		t.Errorf("Received %d IP transport details", len(result.IPTransportDetails))
	}

	if result.IPTransportDetails[0].IPv4Address != "127.0.0.1" {
		t.Errorf("Received IP transport IPv4: %s", result.IPTransportDetails[0].IPv4Address)
	}

	if len(result.Identifiers) != 1 {
		t.Errorf("Received %d identifiers", len(result.Identifiers))
	}

	if result.Identifiers[0].DurableNameFormat != common.IQNDurableNameFormat {
		t.Errorf("Received durable name format: %s", result.Identifiers[0].DurableNameFormat)
	}
}
