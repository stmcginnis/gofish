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
	expectErrorStatus = `{"error": ` + errMsg + "}"
)

// TestError400 tests the parsing of error reply.
func TestError400(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(expectErrorStatus))
	}))
	defer ts.Close()

	_, err := Connect(ClientConfig{Endpoint: ts.URL, HTTPClient: ts.Client()})
	if err == nil {
		t.Error("Update call should fail")
	}
	errStruct, ok := err.(*common.Error)
	if !ok {
		t.Errorf("400 should return known error type: %v", err)
	}
	if errStruct.HTTPReturnedStatusCode != 400 {
		t.Errorf("The error code is different from 400")
	}
	errBody, err := json.MarshalIndent(errStruct, "  ", "    ")
	if err != nil {
		t.Errorf("Marshall error %v got: %s", errStruct, err)
	}
	if errMsg != string(errBody) {
		t.Errorf("Expect:\n%s\nGot:\n%s", errMsg, string(errBody))
	}
}

// TestError404 tests the parsing of error reply.
func TestError404(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(expectErrorStatus))
	}))
	defer ts.Close()

	_, err := Connect(ClientConfig{Endpoint: ts.URL, HTTPClient: ts.Client()})
	if err == nil {
		t.Error("Update call should fail")
	}
	errStruct, ok := err.(*common.Error)
	if !ok {
		t.Errorf("404 should return known error type: %v", err)
	}
	if errStruct.HTTPReturnedStatusCode != 404 {
		t.Errorf("The error code is different from 404")
	}
	errBody, err := json.MarshalIndent(errStruct, "  ", "    ")
	if err != nil {
		t.Errorf("Marshall error %v got: %s", errStruct, err)
	}
	if errMsg != string(errBody) {
		t.Errorf("Expect:\n%s\nGot:\n%s", errMsg, string(errBody))
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
