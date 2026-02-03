//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

var sessionBody = `{
		"@odata.context": "/redfish/v1/$metadata#Session.Session",
		"@odata.type": "#Session.v1_2_0.Session",
		"@odata.id": "/redfish/v1/Session",
		"Id": "Session-1",
		"Name": "SessionOne",
		"Description": "Session One",
		"OemSessionType": "Ticket",
		"SessionType": "OEM",
		"UserName": "mfreeman"
	}`

// TestSession tests the parsing of Session objects.
func TestSession(t *testing.T) {
	var result Session
	err := json.NewDecoder(strings.NewReader(sessionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Session-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SessionOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Password != "" {
		t.Error("Password should be nil")
	}

	if result.SessionType != OEMSessionTypes {
		t.Errorf("Invalid session type: %s", result.SessionType)
	}

	if result.UserName != "mfreeman" {
		t.Errorf("Invalid user name: %s", result.UserName)
	}
}

// TestCreateSession tests the CreateSession call.
func TestCreateSession(t *testing.T) {
	var result Session
	err := json.NewDecoder(strings.NewReader(sessionBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// define the expected values
	expectedSessionURI := "/redfish/v1/SessionService/Sessions/SessionId/"
	expectedXAuthToken := "expectedXAuthToken"

	// create the custom test client
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				// defining the custom return for the first POST operation
				&http.Response{
					Status:        "201 Created",
					StatusCode:    201,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString("")),
					ContentLength: int64(len("")),
					Header: http.Header{
						"Location": []string{
							expectedSessionURI,
						},
						"X-Auth-Token": []string{
							expectedXAuthToken,
						},
					},
				},
			},
		},
	}
	result.SetClient(testClient)

	// create the session
	auth, err := CreateSession(
		testClient,
		"/redfish/v1/SessionService/Sessions/",
		"user",
		"password",
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateSession call: %s", err)
	}

	if auth.Session != expectedSessionURI {
		t.Errorf("Error CreateSession returned: %s expected: %s",
			auth.Session,
			expectedSessionURI)
	}

	if auth.Token != expectedXAuthToken {
		t.Errorf("Error CreateSession returned: %s expected: %s",
			auth.Token,
			expectedXAuthToken)
	}

	// validate the payload
	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "UserName:user") {
		t.Errorf("Unexpected Username CreateSession payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Password:password") {
		t.Errorf("Unexpected Password CreateSession payload: %s", calls[0].Payload)
	}
}

// TestCreateSessionFullURIPath tests the CreateSession call
// when the vendor returns full URI path.
func TestCreateSessionFullURIPath(t *testing.T) {
	var result Session
	err := json.NewDecoder(strings.NewReader(sessionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// define the expected values
	expectedSessionURI := "/redfish/v1/SessionService/Sessions/SessionId/"
	fullURIPath := fmt.Sprintf("https://redfish-server%s", expectedSessionURI)
	expectedXAuthToken := "expectedXAuthToken"

	// create the custom test client
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				// defining the custom return for the first POST operation
				&http.Response{
					Status:        "201 Created",
					StatusCode:    201,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString("")),
					ContentLength: int64(len("")),
					Header: http.Header{
						"Location": []string{
							fullURIPath,
						},
						"X-Auth-Token": []string{
							expectedXAuthToken,
						},
					},
				},
			},
		},
	}
	result.SetClient(testClient)

	// create session
	auth, err := CreateSession(
		testClient,
		"/redfish/v1/SessionService/Sessions/",
		"user",
		"password",
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateSession call: %s", err)
	}

	if auth.Session != expectedSessionURI {
		t.Errorf("Error CreateSession returned: %s expected: %s",
			auth.Session,
			expectedSessionURI)
	}

	if auth.Token != expectedXAuthToken {
		t.Errorf("Error CreateSession returned: %s expected: %s",
			auth.Token,
			expectedXAuthToken)
	}

	// validate the payload
	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "UserName:user") {
		t.Errorf("Unexpected Username CreateSession payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Password:password") {
		t.Errorf("Unexpected Password CreateSession payload: %s", calls[0].Payload)
	}
}
