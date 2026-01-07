//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var outboundConnectionBody = `{
	"@odata.type": "#OutboundConnection.v1_0_1.OutboundConnection",
	"Id": "1",
	"Name": "Outbound Connection to contoso app",
	"Status": {
	  "Health": "OK",
	  "HealthRollup": "OK",
	  "State": "Enabled"
	},
	"Authentication": "MTLS",
	"Certificates": {
	  "@odata.id": "/redfish/v1/AccountService/OutboundConnections/1/Certificates"
	},
	"ClientCertificates": {
	  "@odata.id": "/redfish/v1/AccountService/OutboundConnections/1/ClientCertificates"
	},
	"ConnectionEnabled": true,
	"EndpointURI": "wss://ws.contoso.com:443",
	"RetryPolicy": {
	  "ConnectionRetryPolicy": "RetryCount",
	  "RetryIntervalMinutes": 5,
	  "RetryCount": 60
	},
	"Roles": [
	  "Administrator"
	],
	"WebSocketPingIntervalMinutes": 10,
	"@odata.id": "/redfish/v1/AccountService/OutboundConnections/1"
  }`

// TestOutboundConnection tests the parsing of OutboundConnection objects.
func TestOutboundConnection(t *testing.T) {
	var result OutboundConnection
	err := json.NewDecoder(strings.NewReader(outboundConnectionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Outbound Connection to contoso app", result.Name)
	assertEquals(t, "MTLS", string(result.Authentication))
	assertEquals(t, "/redfish/v1/AccountService/OutboundConnections/1/Certificates", result.certificates)
	assertEquals(t, "/redfish/v1/AccountService/OutboundConnections/1/ClientCertificates", result.clientCertificates)
	assertEquals(t, "wss://ws.contoso.com:443", result.EndpointURI)
	assertEquals(t, "RetryCount", string(result.RetryPolicy.ConnectionRetryPolicy))
	assertEquals(t, "Administrator", result.Roles[0])

	if !result.ConnectionEnabled {
		t.Error("Expected connection to be enabled")
	}

	if *result.RetryPolicy.RetryIntervalMinutes != 5 {
		t.Errorf("Unexpected RetryIntervalMinutes value: %d", result.RetryPolicy.RetryIntervalMinutes)
	}

	if *result.WebSocketPingIntervalMinutes != 10 {
		t.Errorf("Unexpected WebSocketPingIntervalMinutes value: %d", result.WebSocketPingIntervalMinutes)
	}
}
