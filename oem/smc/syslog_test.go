//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var syslogBody = `{
  "@odata.type": "#Syslog.v1_0_1.Syslog",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog",
  "Id": "Syslog",
  "Name": "Syslog",
  "EnableSyslog": true,
  "SyslogServer": "localhost",
  "SyslogPortNumber": 514,
  "@odata.etag": "\"b27af6393687bb1810b00fe52874e053\""
}`

var updatedBody = `{
  "@odata.type": "#Syslog.v1_0_2.Syslog",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog",
  "Id": "Syslog",
  "Name": "Syslog",
  "EnableSyslog": true,
  "SyslogServer": [
    {
      "ServerIP": "syslog.xxxx.yyy",
      "Port": 514
    },
    {
      "ServerIP": null,
      "Port": null
    },
    {
      "ServerIP": null,
      "Port": null
    }
  ],
  "@odata.etag": "\"deb0534b5f0661ee3e0ea15dc3d6023b\""
}`

// TestSyslog tests the parsing of Syslog objects.
func TestSyslog(t *testing.T) {
	var result Syslog
	err := json.NewDecoder(strings.NewReader(syslogBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Syslog" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}

	if result.Server != "localhost" {
		t.Errorf("Invalid server: %s", result.Server)
	}

	if result.Port != 514 {
		t.Errorf("Invalid port: %d", result.Port)
	}
}

// TestSyslogUpdate tests updates against firmware reporting a single server.
func TestSyslogUpdate(t *testing.T) {
	var result Syslog
	if err := json.NewDecoder(strings.NewReader(syslogBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &schemas.TestClient{}
	result.SetClient(testClient)
	result.Server = "syslog.example.com"
	result.Port = 1514

	if err := result.Update(); err != nil {
		t.Fatalf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call, got %d", len(calls))
	}

	if !strings.Contains(calls[0].Payload, "SyslogServer:syslog.example.com") {
		t.Errorf("Unexpected server update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "SyslogPortNumber:1514") {
		t.Errorf("Unexpected port update payload: %s", calls[0].Payload)
	}
}

// TestNewSyslogUpdate tests updates against firmware reporting a collection of
// servers.
func TestNewSyslogUpdate(t *testing.T) {
	var result Syslog
	if err := json.NewDecoder(strings.NewReader(updatedBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &schemas.TestClient{}
	result.SetClient(testClient)
	result.Servers[1].ServerIP = "syslog2.example.com"

	if err := result.Update(); err != nil {
		t.Fatalf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call, got %d", len(calls))
	}

	if !strings.Contains(calls[0].Payload, "SyslogServer:[") {
		t.Errorf("Expected a SyslogServer collection payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "syslog2.example.com") {
		t.Errorf("Unexpected server update payload: %s", calls[0].Payload)
	}
}

// TestNewSyslogUpdateFoldsSingularFields tests that changes to the backwards
// compatible Server and Port fields are applied to the first collection entry.
func TestNewSyslogUpdateFoldsSingularFields(t *testing.T) {
	var result Syslog
	if err := json.NewDecoder(strings.NewReader(updatedBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &schemas.TestClient{}
	result.SetClient(testClient)
	result.Server = "syslog.example.com"
	result.Port = 1514

	if err := result.Update(); err != nil {
		t.Fatalf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call, got %d", len(calls))
	}

	if !strings.Contains(calls[0].Payload, "syslog.example.com") {
		t.Errorf("Server change not folded into the collection: %s", calls[0].Payload)
	}
}

// TestSyslogUpdateNoChanges tests that no request is sent when nothing changed.
func TestSyslogUpdateNoChanges(t *testing.T) {
	for name, body := range map[string]string{"singular": syslogBody, "collection": updatedBody} {
		t.Run(name, func(t *testing.T) {
			var result Syslog
			if err := json.NewDecoder(strings.NewReader(body)).Decode(&result); err != nil {
				t.Fatalf("Error decoding JSON: %s", err)
			}

			testClient := &schemas.TestClient{}
			result.SetClient(testClient)

			if err := result.Update(); err != nil {
				t.Fatalf("Error making Update call: %s", err)
			}

			if calls := testClient.CapturedCalls(); len(calls) != 0 {
				t.Errorf("Expected no calls, got %d", len(calls))
			}
		})
	}
}

// TestNewSyslog tests the parsing of updated Syslog objects that contain an
// array of syslog server entries.
func TestNewSyslog(t *testing.T) {
	var result Syslog
	err := json.NewDecoder(strings.NewReader(updatedBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Syslog" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}

	if len(result.Servers) != 3 {
		t.Errorf("Expected 3 syslog server entries, got %d", len(result.Servers))
	}

	if result.Servers[0].ServerIP != "syslog.xxxx.yyy" {
		t.Errorf("Invalid server: %s", result.Server)
	}

	if *result.Servers[0].Port != 514 {
		t.Errorf("Invalid port: %d", result.Port)
	}

	// Check backwards compatibility
	if result.Server != "syslog.xxxx.yyy" {
		t.Errorf("Invalid server: %s", result.Server)
	}

	if result.Port != 514 {
		t.Errorf("Invalid port: %d", result.Port)
	}
}
