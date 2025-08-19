//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
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
	  	"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.EjectMedia",
		"@Redfish.ActionInfo": "/redfish/v1/Managers/1/VirtualMedia/1/EjectMediaActionInfo"
	  },
	  "#VirtualMedia.InsertMedia": {
	  	"target": "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.InsertMedia",
		"@Redfish.ActionInfo": "/redfish/v1/Managers/1/VirtualMedia/1/InsertMediaActionInfo"
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

	if result.ejectMedia.Target != "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.EjectMedia" {
		t.Errorf("Received invalid EjectMedia Action target: %s", result.ejectMedia.Target)
	}
	if result.ejectMedia.ActionInfoTarget != "/redfish/v1/Managers/1/VirtualMedia/1/EjectMediaActionInfo" {
		t.Errorf("Received invalid EjectMediaActionInfo target: %s", result.ejectMedia.ActionInfoTarget)
	}

	if result.insertMedia.Target != "/redfish/v1/Managers/1/VirtualMedia/1/Actions/VirtualMedia.InsertMedia" {
		t.Errorf("Received invalid InsertMedia Action target: %s", result.insertMedia.Target)
	}
	if result.insertMedia.ActionInfoTarget != "/redfish/v1/Managers/1/VirtualMedia/1/InsertMediaActionInfo" {
		t.Errorf("Received invalid InsertMediaActionInfo target: %s", result.insertMedia.ActionInfoTarget)
	}

	if *result.Inserted == true {
		t.Error("Expected Inserted to be false")
	}

	if *result.WriteProtected == false {
		t.Error("Expected WriteProtected to be true")
	}

	if *result.Image != "https://example.com/mygoldimage.iso" {
		t.Errorf("Expected Image to be 'https://example.com/mygoldimage.iso', got %s", *result.Image)
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
	name := "Fred"
	fls := false
	testClient := &common.TestClient{}
	result.SetClient(testClient)
	result.UserName = &name
	result.WriteProtected = &fls

	err = result.Update()
	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	// Проверяем, что был сделан хотя бы один вызов
	if len(calls) == 0 {
		t.Errorf("No calls were captured")
	}

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

	if calls[0].Payload != "" {
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

	_, err = result.InsertMedia("https://example.com/image", false, true)
	if err != nil {
		t.Errorf("Error making InsertMedia call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Image:https://example.com/image") {
		t.Errorf("Unexpected InsertMedia Image payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Inserted:true") {
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

	Inserted := true
	MediaType := CDMediaType
	Password := "test1234"
	UserName := "root"
	WriteProtected := true
	virtualMediaConfig := VirtualMediaConfig{
		Image:          "https://example.com/image",
		Inserted:       &Inserted,
		MediaType:      &MediaType,
		Password:       &Password,
		UserName:       &UserName,
		WriteProtected: &WriteProtected,
	}

	_, err = result.InsertMediaConfig(virtualMediaConfig)
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
		t.Errorf("Unexpected InsertMedia MediaType payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "Password:test1234") {
		t.Errorf("Unexpected InsertMedia Password payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "UserName:root") {
		t.Errorf("Unexpected InsertMedia UserName payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "WriteProtected:true") {
		t.Errorf("Unexpected InsertMedia WriteProtected payload: %s", calls[0].Payload)
	}
}

// TestVirtualMediaActionInfo tests the ActionInfo call.
func TestVirtualMediaActionInfo(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(vmBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	want := map[string]interface{}{
		"Id":   "InsertMedia",
		"Name": "InsertMedia",
		"Parameters": []ActionInfoParameter{
			{Name: "Image", DataType: "String", Required: true},
			{Name: "Inserted", DataType: "Boolean"},
			{Name: "WriteProtected", DataType: "Boolean"},
			{Name: "TransferProtocolType", DataType: "String", Required: true, AllowableValues: []string{"NFS", "CIFS"}},
			{Name: "TransferMethod", DataType: "String", AllowableValues: []string{"Stream"}},
			{Name: "UserName", DataType: "String"},
			{Name: "Password", DataType: "String"},
		},
	}
	b, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("Failed to marshal json: %s", err)
	}

	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBuffer(b)),
				},
			},
		},
	}
	result.SetClient(testClient)

	got, err := result.InsertMediaActionInfo()
	if err != nil {
		t.Errorf("Error making InsertMediaActionInfo call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if calls[0].Payload != "" {
		t.Errorf("Unexpected InsertMediaActionInfo payload: %s", calls[0].Payload)
	}

	if got.ID != want["Id"].(string) {
		t.Errorf("Unexpected ID, want: %v, got: %v", want["Id"], got.ID)
	}
	if got.Name != want["Name"].(string) {
		t.Errorf("Unexpected Name, want: %v, got: %v", want["Name"], got.Name)
	}
	if !reflect.DeepEqual(got.Parameters, want["Parameters"].([]ActionInfoParameter)) {
		t.Errorf("Parameters don't match, \nwant: %v\n got: %v", want["Parameters"], got.Parameters)
	}
}
