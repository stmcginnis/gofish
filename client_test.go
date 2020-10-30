//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
	expectErrorStatus400 = `{"error": ` + errMsg + "}"
	expectErrorStatus404 = `404: {"error": ` + errMsg + "}"
)

// TestError400 tests the parsing of error reply.
func TestError400(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(expectErrorStatus400))
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
	errBody, err := json.MarshalIndent(errStruct, "  ", "    ")
	if err != nil {
		t.Errorf("Marshall error %v got: %s", errStruct, err)
	}
	if errMsg != string(errBody) {
		t.Errorf("Expect:\n%s\nGot:\n%s", errMsg, string(errBody))
	}
}

// TestErrorNon400 tests the parsing of error reply for non 400 reply.
func TestErrorNon400(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(expectErrorStatus400))
	}))
	defer ts.Close()

	_, err := Connect(ClientConfig{Endpoint: ts.URL, HTTPClient: ts.Client()})
	if err == nil {
		t.Error("Update call should fail")
	}
	_, ok := err.(*common.Error)
	if ok {
		t.Errorf("404 should not return known error type: %v", err)
	}
	if expectErrorStatus404 != err.Error() {
		t.Errorf("Expect:\n%s\nGot:\n%s", expectErrorStatus404, err.Error())
	}
}
