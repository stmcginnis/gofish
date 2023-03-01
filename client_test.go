//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stmcginnis/gofish/common"
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	client := APIClient{mu: &sync.Mutex{}}

	_, err := client.runRawRequest("", "", nil, "") //nolint:bodyclose
	if err == nil {
		t.Error("Request without relative path should have failed")
	}

	if err.Error() != "unable to execute request, no target provided" {
		t.Errorf("Unexpected error response: %s", err.Error())
	}
}
