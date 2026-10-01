//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

const metricsURI = "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Metrics"

// batteryWithMetrics builds a Battery payload whose "Metrics" property holds
// the given raw JSON, standing in for whatever a service returns for a
// singleton link.
func batteryWithMetrics(metrics string) string {
	return `{
		"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1",
		"Id": "Module1",
		"Name": "Battery 1",
		"Metrics": ` + metrics + `
	}`
}

// fetchedMetrics is the response a live GET of the metrics URI would return.
func fetchedMetrics() *http.Response {
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
			"@odata.id": "` + metricsURI + `",
			"Id": "Metrics",
			"Name": "Fetched Metrics",
			"DischargeCycles": 8.0
		}`)),
	}
}

// decodeBattery decodes body into a Battery backed by a fresh TestClient that
// answers every GET with fetchedMetrics.
func decodeBattery(t *testing.T, body string) (*Battery, *TestClient) {
	t.Helper()

	var battery Battery
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&battery); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	client := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {fetchedMetrics(), fetchedMetrics()},
		},
	}
	battery.SetClient(client)

	return &battery, client
}

// TestResolvedLinkExpanded verifies that a fully inlined singleton - what a
// service returns for $expand - is handed back without a fetch.
func TestResolvedLinkExpanded(t *testing.T) {
	battery, client := decodeBattery(t, batteryWithMetrics(`{
		"@odata.id": "`+metricsURI+`",
		"@odata.type": "#BatteryMetrics.v1_1_0.BatteryMetrics",
		"Id": "Metrics",
		"Name": "Inlined Metrics",
		"DischargeCycles": 3.0
	}`))

	assertEquals(t, metricsURI, battery.metrics.URI())

	metrics, err := battery.Metrics()
	if err != nil {
		t.Fatalf("Error getting metrics: %s", err)
	}
	if metrics == nil {
		t.Fatal("Expected the inlined metrics to be returned")
	}

	assertEquals(t, "Inlined Metrics", metrics.Name)
	if calls := client.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no calls for an inlined link, captured: %v", calls)
	}
	if metrics.GetClient() != client {
		t.Error("Expected the client to be set on the inlined object")
	}

	// The inlined value is consumed exactly once, so a second call re-reads
	// the resource from the service.
	metrics, err = battery.Metrics()
	if err != nil {
		t.Fatalf("Error getting metrics a second time: %s", err)
	}
	assertEquals(t, "Fetched Metrics", metrics.Name)

	calls := client.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call on the second get, captured: %v", calls)
	}
	assertEquals(t, metricsURI, calls[0].URL)
}

// These Port and PortMetrics fields are trimmed from recorded H100 BMC
// responses. The captures are unexpanded; this test assembles them into a
// level-2 expand response rather than claiming it was captured verbatim.
func TestResolvedLinkExpandedCollectionMember(t *testing.T) {
	const portURI = "/redfish/v1/Fabrics/HGX_NVLinkFabric_0/Switches/NVSwitch_0/Ports/NVLink_0"
	const portMetricsURI = portURI + "/Metrics"
	const collectionURI = "/redfish/v1/Fabrics/HGX_NVLinkFabric_0/Switches/NVSwitch_0/Ports"

	client := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{
					"@odata.id": "` + collectionURI + `",
					"@odata.type": "#PortCollection.PortCollection",
					"Members": [{
						"@odata.id": "` + portURI + `",
						"@odata.type": "#Port.v1_4_0.Port",
						"Id": "NVLink_0",
						"LinkState": "Enabled",
						"Metrics": {
							"@odata.id": "` + portMetricsURI + `",
							"@odata.type": "#PortMetrics.v1_3_0.PortMetrics",
							"Id": "Metrics",
							"Name": "NVLink_0 Port Metrics",
							"RXBytes": 0,
							"TXBytes": 0
						}
					}],
					"Members@odata.count": 1
				}`))},
				&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{
					"@odata.id": "` + portMetricsURI + `",
					"Id": "Metrics",
					"Name": "Refetched Metrics",
					"RXBytes": 1
				}`))},
			},
		},
	}

	ports, err := GetCollectionObjects[Port](client, collectionURI)
	if err != nil {
		t.Fatalf("Error getting ports: %s", err)
	}
	if len(ports) != 1 {
		t.Fatalf("Expected one expanded port, got %d", len(ports))
	}
	assertEquals(t, "NVLink_0", ports[0].ID)

	metrics, err := ports[0].Metrics()
	if err != nil {
		t.Fatalf("Error getting expanded metrics: %s", err)
	}
	if metrics == nil || metrics.RXBytes == nil || metrics.TXBytes == nil {
		t.Fatalf("Expected populated inlined metrics, got: %+v", metrics)
	}
	assertEquals(t, "NVLink_0 Port Metrics", metrics.Name)
	if *metrics.RXBytes != 0 || *metrics.TXBytes != 0 {
		t.Errorf("Expected captured zero byte counters, got RX=%d TX=%d", *metrics.RXBytes, *metrics.TXBytes)
	}
	if metrics.GetClient() != client {
		t.Error("Expected the client to be set on the inlined metrics")
	}
	if calls := client.CapturedCalls(); len(calls) != 1 || calls[0].URL != collectionURI {
		t.Fatalf("Expected only the collection GET, captured: %v", calls)
	}

	metrics, err = ports[0].Metrics()
	if err != nil {
		t.Fatalf("Error refetching metrics: %s", err)
	}
	assertEquals(t, "Refetched Metrics", metrics.Name)
	if calls := client.CapturedCalls(); len(calls) != 2 || calls[1].URL != portMetricsURI {
		t.Fatalf("Expected a live GET for the second metrics call, captured: %v", calls)
	}
}

// TestResolvedLinkNotExpanded covers the payloads that carry a URI but no
// actual content. Each must result in a live fetch: anything else hands the
// caller a hollow object built from a link.
func TestResolvedLinkNotExpanded(t *testing.T) {
	tests := []struct {
		name    string
		metrics string
	}{
		{
			name:    "bare link",
			metrics: `{"@odata.id": "` + metricsURI + `"}`,
		},
		{
			// A link decorated with type/etag annotations has keys beyond
			// "@odata.id" but is still just a link.
			name: "annotated link",
			metrics: `{
				"@odata.id": "` + metricsURI + `",
				"@odata.type": "#BatteryMetrics.v1_1_0.BatteryMetrics",
				"@odata.etag": "W/\"1234567890\""
			}`,
		},
		{
			// Some services describe a link with "href" rather than
			// "@odata.id"; the annotation makes it look inlined.
			name: "annotated href link",
			metrics: `{
				"href": "` + metricsURI + `",
				"@odata.type": "#BatteryMetrics.v1_1_0.BatteryMetrics"
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			battery, client := decodeBattery(t, batteryWithMetrics(test.metrics))

			assertEquals(t, metricsURI, battery.metrics.URI())

			metrics, err := battery.Metrics()
			if err != nil {
				t.Fatalf("Error getting metrics: %s", err)
			}
			assertEquals(t, "Fetched Metrics", metrics.Name)

			calls := client.CapturedCalls()
			if len(calls) != 1 {
				t.Fatalf("Expected one call for a link-only payload, captured: %v", calls)
			}
			assertEquals(t, metricsURI, calls[0].URL)
		})
	}
}

// TestResolvedLinkUnset verifies that an absent link behaves like GetObject
// against an empty URI: no value, no error, no call.
func TestResolvedLinkUnset(t *testing.T) {
	battery, client := decodeBattery(t, `{
		"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1",
		"Id": "Module1",
		"Name": "Battery 1"
	}`)

	assertEquals(t, "", battery.metrics.URI())

	metrics, err := battery.Metrics()
	if err != nil {
		t.Fatalf("Error getting metrics: %s", err)
	}
	if metrics != nil {
		t.Errorf("Expected no metrics for an unset link, got: %v", metrics)
	}
	if calls := client.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no calls for an unset link, captured: %v", calls)
	}
}

// TestResolvedLinkExpandedError verifies that an inlined payload carrying
// extended error information is surfaced as an error rather than as a
// successfully expanded object.
func TestResolvedLinkExpandedError(t *testing.T) {
	battery, client := decodeBattery(t, batteryWithMetrics(`{
		"@odata.id": "`+metricsURI+`",
		"Id": "Metrics",
		"Name": "Metrics",
		"@Message.ExtendedInfo": [
			{
				"MessageId": "Base.1.0.InsufficientPrivilege",
				"Message": "There are insufficient privileges for the account."
			}
		]
	}`))

	_, err := battery.Metrics()
	if err == nil {
		t.Fatal("Expected an error for an inlined payload with extended info")
	}

	var redfishErr *Error
	if !errors.As(err, &redfishErr) {
		t.Fatalf("Expected a *schemas.Error, got %T: %s", err, err)
	}
	if len(redfishErr.ExtendedInfos) != 1 {
		t.Fatalf("Expected one extended info, got: %v", redfishErr.ExtendedInfos)
	}
	assertEquals(t, "Base.1.0.InsufficientPrivilege", redfishErr.ExtendedInfos[0].MessageID)
	if calls := client.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no calls for an inlined link, captured: %v", calls)
	}
}
