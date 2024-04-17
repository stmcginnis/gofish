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

var vmBody = `{
	"@odata.context": "/redfish/v1/$metadata#VirtualMedia.VirtualMedia",
	"@odata.etag": "W/\"\"",
	"@odata.id": "/redfish/v1/Managers/1/VirtualMedia/1",
	"@odata.type": "#VirtualMedia.v1_2_0.VirtualMedia",
	"Id": "1",
	"Actions": {
	  "#VirtualMedia.EjectMedia": {
		"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.EjectMedia"
	  },
	  "#VirtualMedia.InsertMedia": {
		"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.InsertMedia"
	  }
	},
	"ConnectedVia": "NotConnected",
	"Description": "Virtual Removable Media",
	"Image": "https://example.com/mygoldimage.iso",
	"ImageName": "mygoldimage.iso",
	"Inserted": false,
	"MediaTypes": [
	  "Floppy",
	  "USBStick"
	],
	"Name": "VirtualMedia",
	"Oem": {
	  "Hpe": {
		"@odata.context": "/redfish/v1/$metadata#HpeiLOVirtualMedia.HpeiLOVirtualMedia",
		"@odata.type": "#HpeiLOVirtualMedia.v2_2_0.HpeiLOVirtualMedia",
		"Actions": {
		  "#HpeiLOVirtualMedia.EjectVirtualMedia": {
			"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/Oem/Hpe/HpeiLOVirtualMedia.EjectVirtualMedia"
		  },
		  "#HpeiLOVirtualMedia.InsertVirtualMedia": {
			"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/Oem/Hpe/HpeiLOVirtualMedia.InsertVirtualMedia"
		  }
		}
	  }
	},
	"WriteProtected": true
  }`

// TestVirtualMedia tests the parsing of VirtualMedia objects.
func TestVirtualMedia(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "VirtualMedia" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ejectMediaTarget != "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.EjectMedia" {
		t.Errorf("Received invalid EjectMedia Action target: %s", result.ejectMediaTarget)
	}

	if result.insertMediaTarget != "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.InsertMedia" {
		t.Errorf("Received invalid InsertMedaiaAction target: %s", result.insertMediaTarget)
	}

	if result.Inserted == true {
		t.Error("Expected Inserted to be false")
	}

	if result.WriteProtected == false {
		t.Error("Expected WriteProtected to be true")
	}

	if result.Image != "https://example.com/mygoldimage.iso" {
		t.Errorf("Expected Image to be 'https://example.com/mygoldimage.iso', got %s", result.Image)
	}

	if result.ImageName != "mygoldimage.iso" {
		t.Errorf("Expected ImageName to be 'mygoldimage.iso', got %s", result.ImageName)
	}

	if result.SupportsMediaInsert == false {
		t.Error("Expected SupportsMediaInsert to be true since target is set")
	}
}

// TestVirtualMediaUpdate tests the Update call.
func TestVirtualMediaUpdate(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)
	result.UserName = "Fred"
	result.WriteProtected = false
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "UserName:Fred") {
		t.Errorf("Unexpected UserName update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "WriteProtected:false") {
		t.Errorf("Unexpected WriteProtected update payload: %s", calls[0].Payload)
	}
}

// TestVirtualMediaEject tests the EjectMedia call.
func TestVirtualMediaEject(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.EjectMedia()

	if err != nil {
		t.Errorf("Error making EjectMedia call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if calls[0].Payload != "map[]" {
		t.Errorf("Unexpected EjectMedia payload: %s", calls[0].Payload)
	}
}

// TestVirtualMediaInser tests the InsertMedia call.
func TestVirtualMediaInsert(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.InsertMedia("https://example.com/image", false, true)

	if err != nil {
		t.Errorf("Error making InsertMedia call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Image:https://example.com/image") {
		t.Errorf("Unexpected InsertMedia Image payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Inserted:false") {
		t.Errorf("Unexpected InsertMedia Inserted payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "WriteProtected:true") {
		t.Errorf("Unexpected InsertMedia WriteProtected payload: %s", calls[0].Payload)
	}
}

func TestVirtualMediaInsertConfig(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	virtualMediaConfig := VirtualMediaConfig{
		Image:          "https://example.com/image",
		Inserted:       true,
		MediaType:      "CD",
		Password:       "test1234",
		UserName:       "root",
		WriteProtected: true,
	}

	err = result.InsertMediaConfig(virtualMediaConfig)
	if err != nil {
		t.Errorf("Error making InsertMediaConfig call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Image:https://example.com/image") {
		t.Errorf("Unexpected InsertMedia Image payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "Inserted:true") {
		t.Errorf("Unexpected InsertMedia Inserted payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "MediaType:CD") {
		t.Errorf("Unexpected InsertMedia Inserted payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "Password:test1234") {
		t.Errorf("Unexpected InsertMedia Image payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "UserName:root") {
		t.Errorf("Unexpected InsertMedia Image payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "WriteProtected:true") {
		t.Errorf("Unexpected InsertMedia Inserted payload: %s", calls[0].Payload)
	}
}
