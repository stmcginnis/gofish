//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var ioConnectivityLoSCapabilitiesBody = `{
		"@odata.context": "/redfish/v1/$metadata#IOConnectivityLoSCapabilities.IOConnectivityLoSCapabilities",
		"@odata.type": "#IOConnectivityLoSCapabilities.v1_1_1.IOConnectivityLoSCapabilities",
		"@odata.id": "/redfish/v1/IOConnectivityLoSCapabilities",
		"Id": "IOConnectivityLoSCapabilities-1",
		"Name": "IOConnectivityLoSCapabilitiesOne",
		"Description": "IOConnectivityLoSCapabilities One",
		"MaxSupportedBytesPerSecond": 5000000000,
		"MaxSupportedIOPS": 1000000000,
		"SupportedAccessProtocols": [
			"FC",
			"FCP",
			"FCoE",
			"iSCSI"
		],
		"SupportedLinesOfService": [{
				"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
				"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
				"Id": "IOConnectivityLineOfService-1",
				"Name": "IOConnectivityLineOfServiceOne",
				"Description": "IOConnectivityLineOfService One",
				"AccessProtocols": [
					"FC",
					"FCP",
					"FCoE",
					"iSCSI"
				],
				"MaxBytesPerSecond": 5000000000,
				"MaxIOPS": 1000000000
			},
			{
				"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
				"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
				"Id": "IOConnectivityLineOfService-2",
				"Name": "IOConnectivityLineOfServiceTwo",
				"Description": "IOConnectivityLineOfService Two",
				"AccessProtocols": [
					"FC",
					"FCP",
					"FCoE"
				],
				"MaxBytesPerSecond": 5000000000,
				"MaxIOPS": 1000000000
			}
		]
	}`

// TestIOConnectivityLoSCapabilities tests the parsing of IOConnectivityLoSCapabilities objects.
func TestIOConnectivityLoSCapabilities(t *testing.T) {
	var result IOConnectivityLoSCapabilities
	err := json.NewDecoder(strings.NewReader(ioConnectivityLoSCapabilitiesBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOConnectivityLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOConnectivityLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if *result.MaxSupportedBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.MaxSupportedBytesPerSecond)
	}

	if *result.MaxSupportedIOPS != 1000000000 {
		t.Errorf("MaxSupportedIOPS: %d", result.MaxSupportedIOPS)
	}

	if result.SupportedAccessProtocols[1] != FCPProtocol {
		t.Errorf("Invalid AccessProtocol: %s", result.SupportedAccessProtocols[1])
	}

	if *result.SupportedLinesOfService[0].MaxBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.SupportedLinesOfService[0].MaxBytesPerSecond)
	}
}

// TestIOConnectivityLoSCapabilitiesUpdate tests the Update call.
func TestIOConnectivityLoSCapabilitiesUpdate(t *testing.T) {
	var result IOConnectivityLoSCapabilities
	err := json.NewDecoder(strings.NewReader(ioConnectivityLoSCapabilitiesBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	*result.MaxSupportedBytesPerSecond = 500
	*result.MaxSupportedIOPS = 10000
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "MaxSupportedBytesPerSecond:500") {
		t.Errorf("Unexpected MaxSupportedBytesPerSecond update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "MaxSupportedIOPS:10000") {
		t.Errorf("Unexpected MaxSupportedIOPS update payload: %s", calls[0].Payload)
	}
}
