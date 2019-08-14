//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var eventDestinationBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#EventDestination.EventDestination",
		"@odata.type": "#EventDestination.v1_0_0.EventDestination",
		"@odata.id": "/redfish/v1/EventDestination",
		"Id": "EventDestination-1",
		"Name": "EventDestinationOne",
		"Description": "EventDestination One",
		"Context": "MyContext",
		"Destination": "http://example.com/events",
		"EventFormatType": "MetricReport",
		"HttpHeaders": [],
		"MessageIds": ["One", "Two"],
		"Protocol": "Redfish",
		"RegistryPrefixes": ["ONE_", "TWO_"],
		"ResourceTypes": ["one", "two"],
		"SubordinateResources": false,
		"SubscriptionType": "SSE"
	}`)

// TestEventDestination tests the parsing of EventDestination objects.
func TestEventDestination(t *testing.T) {
	var result EventDestination
	err := json.NewDecoder(eventDestinationBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "EventDestination-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "EventDestinationOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.HTTPHeaders) != 0 {
		t.Errorf("Expected 0 headers, got: %d", len(result.HTTPHeaders))
	}

	if result.EventFormatType != MetricReportEventFormatType {
		t.Errorf("Invalid event format type: %s", result.EventFormatType)
	}

	if result.Protocol != RedfishEventDestinationProtocol {
		t.Errorf("Invalid event protocol: %s", result.Protocol)
	}

	if result.SubordinateResources {
		t.Error("Subordinate resources should be False")
	}
}
