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

var redundancyBody = `{
		"@odata.id": "/redfish/v1/Redundancy",
		"Id": "Redundancy-1",
		"Name": "RedundancyOne",
		"MaxNumSupported": 2,
		"MemberId": "Redundancy1",
		"MinNumNeeded": 2,
		"Mode": "Sparing",
		"RedundancyEnabled": true,
		"RedundancySet": [
			"One",
			"Two"
		],
		"RedundancySet@odata.count": 2,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`

// TestRedundancy tests the parsing of Redundancy objects.
func TestRedundancy(t *testing.T) {
	var result Redundancy
	err := json.NewDecoder(strings.NewReader(redundancyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Redundancy-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "RedundancyOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaxNumSupported != 2 {
		t.Errorf("Invalid MaxNumSupported: %d", result.MaxNumSupported)
	}

	if result.Mode != SparingRedundancyMode {
		t.Errorf("Invalid mode: %s", result.Mode)
	}
}

// TestRedundancyUpdate tests the Update call.
func TestRedundancyUpdate(t *testing.T) {
	var result Redundancy
	err := json.NewDecoder(strings.NewReader(redundancyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.Mode = NotRedundantRedundancyMode
	result.RedundancyEnabled = true
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Mode:NotRedundant") {
		t.Errorf("Unexpected Mode update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "RedundancyEnabled") {
		t.Errorf("Unexpected update for RedundancyEnabled in payload: %s", calls[0].Payload)
	}
}
