//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var ioConnectivityLineOfServiceBody = strings.NewReader(
	`{
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
	}`)

// TestIOConnectivityLineOfService tests the parsing of IOConnectivityLineOfService objects.
func TestIOConnectivityLineOfService(t *testing.T) {
	var result IOConnectivityLineOfService
	err := json.NewDecoder(ioConnectivityLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOConnectivityLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOConnectivityLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessProtocols[3] != common.ISCSIProtocol {
		t.Errorf("Invalid access protocol: %s", result.AccessProtocols[3])
	}

	if result.MaxBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxBytesPerSecond: %d", result.MaxBytesPerSecond)
	}

	if result.MaxIOPS != 1000000000 {
		t.Errorf("Invalid MaxIOPS: %d", result.MaxIOPS)
	}
}
