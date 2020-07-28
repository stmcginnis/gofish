//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"fmt"
	"net/http"
)

// TestAPICall captures the arguments to one of the API calls.
type TestAPICall struct {
	// Action is the REST action (GET, PUT, etc) of the call
	Action string
	// URL is the URL to send to
	URL string
	// Payload is the string representation of the payload
	Payload string
}

// TestClient is a mock client to use for unit testing some of the
// function calls and actions that would normally need to connect
// with a host.
type TestClient struct {
	// calls collects any API calls made through the client
	calls []TestAPICall
	// CustomReturnForOperations can be used to define custom
	// return for operations, valid keys are:
	// http.MethodGet, http.MethodPost, http.MethodPut,
	// http.MethodPatch, http.MethodDelete
	// TODO: This can be improved to map cases where
	// in the test we have more than one operation using
	// the same method, for example, two posts with
	// different returns.
	CustomReturnForOperations map[string]interface{}
}

// CapturedCalls gets all calls that were made through this instance
func (c *TestClient) CapturedCalls() []TestAPICall {
	return c.calls
}

// Reset resets the captured information for this mock client.
func (c *TestClient) Reset() {
	c.calls = []TestAPICall{}
}

// recordCall is a helper to record any API calls made through this client
func (c *TestClient) recordCall(action string, url string, payload interface{}) {
	call := TestAPICall{
		Action:  action,
		URL:     url,
		Payload: fmt.Sprintf("%v", payload),
	}
	c.calls = append(c.calls, call)
}

// Get performs a GET request against the Redfish service.
func (c *TestClient) Get(url string) (*http.Response, error) {
	c.recordCall(http.MethodGet, url, nil)
	customReturnForOperation, ok := c.CustomReturnForOperations[http.MethodGet]
	if !ok {
		return nil, nil
	}
	return customReturnForOperation.(*http.Response), nil
}

// Post performs a Post request against the Redfish service.
func (c *TestClient) Post(url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPost, url, payload)
	customReturnForOperation, ok := c.CustomReturnForOperations[http.MethodPost]
	if !ok {
		return nil, nil
	}
	return customReturnForOperation.(*http.Response), nil
}

// Put performs a Put request against the Redfish service.
func (c *TestClient) Put(url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPut, url, payload)
	customReturnForOperation, ok := c.CustomReturnForOperations[http.MethodPut]
	if !ok {
		return nil, nil
	}
	return customReturnForOperation.(*http.Response), nil
}

// Patch performs a Patch request against the Redfish service.
func (c *TestClient) Patch(url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPatch, url, payload)
	customReturnForOperation, ok := c.CustomReturnForOperations[http.MethodPatch]
	if !ok {
		return nil, nil
	}
	return customReturnForOperation.(*http.Response), nil
}

// Delete performs a Delete request against the Redfish service.
func (c *TestClient) Delete(url string) error {
	c.recordCall(http.MethodDelete, url, nil)
	customReturnForOperation, ok := c.CustomReturnForOperations[http.MethodDelete]
	if !ok {
		return nil
	}
	return customReturnForOperation.(error)
}
