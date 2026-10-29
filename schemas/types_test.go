//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func TestIs412PreconditionFailed(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "412 Redfish error",
			err:      ConstructError(http.StatusPreconditionFailed, []byte(`{"error":{"code":"Base.1.18.PreconditionFailed","message":"ETag mismatch"}}`)),
			expected: true,
		},
		{
			name:     "500 Redfish error",
			err:      ConstructError(http.StatusInternalServerError, []byte(`{"error":{"code":"Base.1.18.GeneralError","message":"Internal error"}}`)),
			expected: false,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "non-Redfish error",
			err:      fmt.Errorf("some generic error"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Is412PreconditionFailed(tt.err)
			if result != tt.expected {
				t.Errorf("Is412PreconditionFailed() = %v, want %v", result, tt.expected)
			}
		})
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

// TestLinkMarshalJSON tests that a Link is sent as a resource reference object and decoded back.
func TestLinkMarshalJSON(t *testing.T) {
	link := Link("/redfish/v1/Managers/1")

	b, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("Error marshaling link: %s", err)
	}
	if string(b) != `{"@odata.id":"/redfish/v1/Managers/1"}` {
		t.Errorf("Unexpected link JSON: %s", b)
	}

	var decoded Link
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Error unmarshaling link: %s", err)
	}
	if decoded != link {
		t.Errorf("Link round trip = %q, want %q", decoded, link)
	}

	params := struct {
		Required Link   `json:"Required"`
		Optional Link   `json:"Optional,omitempty"`
		Many     []Link `json:"Many,omitempty"`
	}{
		Required: "/redfish/v1/Systems/1",
		Many:     []Link{"/redfish/v1/Chassis/1", "/redfish/v1/Chassis/2"},
	}
	b, err = json.Marshal(params)
	if err != nil {
		t.Fatalf("Error marshaling parameters: %s", err)
	}
	want := `{"Required":{"@odata.id":"/redfish/v1/Systems/1"},"Many":[{"@odata.id":"/redfish/v1/Chassis/1"},{"@odata.id":"/redfish/v1/Chassis/2"}]}`
	if string(b) != want {
		t.Errorf("Unexpected parameters JSON:\n got %s\nwant %s", b, want)
	}
}
