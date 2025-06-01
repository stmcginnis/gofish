//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const (
	errMsg = `{
      "code": "Base.1.0.GeneralError",
      "message": "A general error has occurred. See ExtendedInfo for more information.",
      "@Message.ExtendedInfo": [
          {
              "MessageId": "Base.1.0.PropertyValueNotInList",
              "Message": "The value Red for the property IndicatorLED is not in the list of acceptable values",
              "MessageArgs": [
                  "RED",
                  "IndicatorLED"
              ],
              "Severity": "Warning",
              "Resolution": "Remove the property from the request body and resubmit the request if the operation failed"
          },
          {
              "MessageId": "Base.1.0.PropertyNotWriteable",
              "Message": "The property SKU is a read only property and cannot be assigned a value",
              "MessageArgs": [
                  "SKU"
              ],
              "Severity": "Warning",
              "Resolution": "Remove the property from the request body and resubmit the request if the operation failed"
          }
      ]
  }`
	expectErrorStatus         = `{"error": ` + errMsg + "}"
	nonErrorStructErrorStatus = "Internal Server Error"
)

func testError(code int, t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(code)
		w.Write([]byte(expectErrorStatus)) //nolint
	}))
	defer ts.Close()

	_, err := Connect(ClientConfig{Endpoint: ts.URL, HTTPClient: ts.Client()})
	if err == nil {
		t.Error("Update call should fail")
	}
	errStruct, ok := err.(*common.Error)
	if !ok {
		t.Errorf("%d should return known error type: %v", code, err)
	}
	if errStruct.HTTPReturnedStatusCode != code {
		t.Errorf("The error code is different from %d", code)
	}
	errBody, err := json.MarshalIndent(errStruct, "  ", "    ")
	if err != nil {
		t.Errorf("Marshall error %v got: %s", errStruct, err)
	}
	if errMsg != string(errBody) {
		t.Errorf("Expect:\n%s\nGot:\n%s", errMsg, string(errBody))
	}
}

// TestError400 tests the parsing of error reply.
func TestError400(t *testing.T) {
	testError(400, t)
}

// TestError404 tests the parsing of error reply.
func TestError404(t *testing.T) {
	testError(404, t)
}

// TestErrorOther tests failures that do not return an Error struct
func TestErrorOther(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(nonErrorStructErrorStatus)) //nolint
	}))
	defer ts.Close()

	_, err := Connect(ClientConfig{Endpoint: ts.URL, HTTPClient: ts.Client()})
	if err == nil {
		t.Error("connect should fail")
	}
	errStruct, ok := err.(*common.Error)
	if !ok {
		t.Errorf("call should return known error type: %v", err)
	}
	if errStruct.HTTPReturnedStatusCode != 500 {
		t.Errorf("The error code is different from 500")
	}
	if err.Error() != "500: Internal Server Error" {
		t.Errorf("Expected '500: %s', got '%s'", nonErrorStructErrorStatus, err.Error())
	}
}

// TestConnectContextTimeout
func TestConnectContextTimeout(t *testing.T) {
	// ctx will timeout very quickly
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Microsecond)
	defer cancel()

	_, err := ConnectContext(
		ctx,
		ClientConfig{
			Endpoint: "https://testContextTimeout.com",
		})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("Context should timeout")
	}
}

func TestServiceGetter(t *testing.T) {
	type serviceGetter interface {
		GetService() *Service
	}

	var sg serviceGetter = &APIClient{}
	if sg.GetService() != nil {
		t.Errorf("Empty client should return a nil service")
	}
}

// TestConnectContextCancel
func TestConnectContextCancel(t *testing.T) {
	// ctx will be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ConnectContext(
		ctx,
		ClientConfig{
			Endpoint: "https://testContextCancel.com",
		})

	if !errors.Is(err, context.Canceled) {
		t.Error("Context should be cancelled")
	}
}

// TestConnectDefaultContextTimeout
func TestConnectDefaultContextTimeout(t *testing.T) {
	// ctx will timeout very quickly
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Microsecond)
	defer cancel()

	_, err := ConnectDefaultContext(
		ctx,
		"https://testContextTimeout.com",
	)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("Context should timeout")
	}
}

// TestConnectDefaultContextCancel
func TestConnectDefaultContextCancel(t *testing.T) {
	// ctx will be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ConnectDefaultContext(
		ctx,
		"https://testContextCancel.com",
	)

	if !errors.Is(err, context.Canceled) {
		t.Error("Context should be cancelled")
	}
}

func TestClientRunRawRequestNoURL(t *testing.T) {
	client := APIClient{sem: make(chan bool, 1)}

	_, err := client.runRawRequest("", "", nil, "") //nolint:bodyclose
	if err == nil {
		t.Error("Request without relative path should have failed")
	}

	if err.Error() != "unable to execute request, no target provided" {
		t.Errorf("Unexpected error response: %s", err.Error())
	}
}

func TestAuthTokenAccessors(t *testing.T) {
	c := APIClient{}

	if c.GetAuthToken() != nil {
		t.Error("Empty client should return a nil auth token")
	}

	token := &redfish.AuthToken{}
	c.SetAuthToken(token)

	if c.GetAuthToken() != token {
		t.Error("Client should return current auth token")
	}

	c.SetAuthToken(nil)

	if c.GetAuthToken() != nil {
		t.Error("Should return nil")
	}
}

const brokenHuaweiServiceRoot = `{
  "@odata.context": "/redfish/v1/$metadata#ServiceRoot",
  "@odata.id": "/redfish/v1/",
  "@odata.type": "#ServiceRoot.v1_1_0.ServiceRoot",
  "Id": "RootService",
  "Name": "Root Service",
  "RedfishVersion": "1.0.2",
  "UUID": "2D0B7460-48DC-BB02-EB11-592730502FEC",
  "Systems": {
    "@odata.id": "/redfish/v1/Systems"
  },
  "Chassis": {
    "@odata.id": "/redfish/v1/Chassis"
  },
  "Managers": {
    "@odata.id": "/redfish/v1/Managers"
  },
  "Tasks": {
    "@odata.id": "/redfish/v1/TaskService"
  },
  "SessionService": {
    "@odata.id": "/redfish/v1/SessionService"
  },
  "AccountService": {
    "@odata.id": "/redfish/v1/AccountService"
  },
  "EventService": {
    "@odata.id": "/redfish/v1/EventService"
  },
  "UpdateService": {
    "@odata.id": "/redfish/v1/UpdateService"
  },
  "Registries": {
    "@odata.id": "/redfish/v1/Registries"
  },
  "JsonSchemas": {
    "@odata.id": "/redfish/v1/JSONSchemas"
  },
  "Oem": {
    "Huawei": {
      "SmsUpdateService": null,
      "ProductName": "2288H V5",
      "HostName": "2102311XBxxxxxxxxxxx",
      "LanguageSet": "en,zh,ja,fr",
      "Copyright": "Huawei Technologies Co., Ltd. 2004-2020. All rights reserved.",
      "DomainName": null,
      "AccountLockoutDuration": 300
    }
  }
}`

func huaweiPatch(s *Service) error {
	// check if this is a Huawei BMC
	if s.Vendor != "Huawei" && s.Vendor != "" {
		return nil
	}
	var result map[string]any
	err := json.Unmarshal(s.Oem, &result)
	if err != nil {
		return err
	}
	var isHuawei = false
	_, isHuawei = result["Huawei"]
	if !isHuawei {
		return nil
	}

	// This code would be called from an external package
	// => can only use public API!
	if s.sessions == "" {
		s.sessions = "/redfish/v1/SessionService/Sessions"
	}

	return nil
}

func TestSetupClientWithConfig(t *testing.T) {
	const xAuthToken = "f38d4148774822d02e4fe8e5bd85ad41"
	const sessionID = "d8899013a6beb9e6"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redfish/v1/" && r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.WriteHeader(200)
			io.WriteString(w, brokenHuaweiServiceRoot)
		} else if r.URL.Path == "/redfish/v1/SessionService/Sessions" && r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.Header().Set("X-Auth-Token", xAuthToken)
			w.Header().Set("Location", fmt.Sprintf("%s/%s", r.URL.Path, sessionID))
			w.WriteHeader(201)
			io.WriteString(w, fmt.Sprintf(
				`{
  "@odata.context": "/redfish/v1/$metadata#Session.Session",
  "@odata.id": "/redfish/v1/SessionService/Sessions/%s",
  "@odata.type": "#Session.v1_0_2.Session",
  "Id": %s,
  "Name": "User Session",
}`, sessionID, sessionID))
		} else {
			w.WriteHeader(404)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	config := ClientConfig{
		Endpoint:   ts.URL,
		HTTPClient: ts.Client(),
		Username:   "foo",
		Password:   "bar",
		BasicAuth:  false,
	}

	// First test the failing case
	// see https://github.com/stmcginnis/gofish/issues/162, and
	//     https://github.com/stmcginnis/gofish/issues/425
	_, err := ConnectContext(ctx, config)
	if err != nil {
		if err.Error() != "unable to execute request, no target provided" {
			t.Errorf("Unexpected error: %v", err)
		}
	} else {
		t.Errorf("Should fail with 'unable to execute request, no target provided'")
	}

	// Now test the lower level API that allows for a client-side vendor patch
	c, err := SetupClientWithConfig(ctx, &config)
	if err != nil {
		t.Errorf("Should succeed, got: %s", err)
	}
	err = huaweiPatch(c.GetService())
	if err != nil {
		t.Errorf("Should succeed, got: %s", err)
	}
	token, err := c.GetService().CreateSession(config.Username, config.Password)
	if err != nil {
		t.Errorf("Should create session, instead: %v", err)
	}
	c.SetAuthToken(token)
	session, err := c.GetSession()
	if err != nil {
		t.Errorf("Should succeed, got: %s", err)
	}
	if session.Token != xAuthToken {
		t.Errorf("Expected X-Auth-Token to be %s, got %s", xAuthToken, session.Token)
	}
	if !strings.HasSuffix(session.ID, sessionID) {
		t.Errorf("Expected session ID suffix %s, got %s", sessionID, session.ID)
	}
	// c.Logout()
}
