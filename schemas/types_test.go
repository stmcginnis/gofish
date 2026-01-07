//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var partLocationBodyString = strings.NewReader(`
	{
		"ServiceLabel": "Battery 1",
		"LocationType": "Bay",
		"LocationOrdinalValue": "5"
	}`)

var partLocationBodyInt = strings.NewReader(`
	{
		"ServiceLabel": "Battery 1",
		"LocationType": "Bay",
		"LocationOrdinalValue": 5
	}`)

func TestPartLocationString(t *testing.T) {
	var result PartLocation
	err := json.NewDecoder(partLocationBodyString).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ServiceLabel != "Battery 1" {
		t.Errorf("Received invalid service label: %s", result.ServiceLabel)
	}

	if result.LocationType != "Bay" {
		t.Errorf("Received invalid location type: %s", result.LocationType)
	}

	if *result.LocationOrdinalValue != 5 {
		t.Errorf("Received invalid location ordinal value: %d", result.LocationOrdinalValue)
	}
}

func TestPartLocationInt(t *testing.T) {
	var result PartLocation
	err := json.NewDecoder(partLocationBodyInt).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ServiceLabel != "Battery 1" {
		t.Errorf("Received invalid service label: %s", result.ServiceLabel)
	}

	if result.LocationType != "Bay" {
		t.Errorf("Received invalid location type: %s", result.LocationType)
	}

	if *result.LocationOrdinalValue != 5 {
		t.Errorf("Received invalid location ordinal value: %d", result.LocationOrdinalValue)
	}
}
