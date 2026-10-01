//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"context"
	"crypto/tls"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stmcginnis/gofish/schemas"
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
	errStruct, ok := err.(*schemas.Error)
	if !ok {
		t.Errorf("%d should return known error type: %v", code, err)
	}

	schemas.AssertEqual(t, code, errStruct.HTTPReturnedStatusCode)
	schemas.AssertEqual(t, "A general error has occurred. See ExtendedInfo for more information.", errStruct.Message)
	schemas.AssertEqual(t, 2, len(errStruct.ExtendedInfos))
	schemas.AssertEqual(t, "Base.1.0.GeneralError", errStruct.Code)
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
	errStruct, ok := err.(*schemas.Error)
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

// TestGetWithHeadersNotModified verifies that a 304 Not Modified response to a
// conditional GET is reported as schemas.ErrNotModified with the response left
// intact so the caller can read the Etag for cache bookkeeping.
func TestGetWithHeadersNotModified(t *testing.T) {
	const etag = `"abc123"`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-None-Match") != etag {
			t.Errorf("expected If-None-Match %q, got %q", etag, r.Header.Get("If-None-Match"))
		}
		w.Header().Set("Etag", etag)
		w.WriteHeader(http.StatusNotModified)
	}))
	defer ts.Close()

	client := &APIClient{
		ctx:        context.Background(),
		endpoint:   ts.URL,
		HTTPClient: ts.Client(),
		sem:        make(chan bool, 1),
	}

	resp, err := client.GetWithHeaders("/redfish/v1/Chassis/1/Thermal",
		map[string]string{"If-None-Match": etag})
	if resp != nil {
		defer schemas.DeferredCleanupHTTPResponse(resp)
	}

	if !errors.Is(err, schemas.ErrNotModified) {
		t.Fatalf("expected schemas.ErrNotModified, got %v", err)
	}
	if resp == nil {
		t.Fatal("expected the response to be returned intact on 304, got nil")
	}
	if got := resp.Header.Get("Etag"); got != etag {
		t.Errorf("expected Etag %q on 304 response, got %q", etag, got)
	}
	if resp.StatusCode != http.StatusNotModified {
		t.Errorf("expected status %d, got %d", http.StatusNotModified, resp.StatusCode)
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

// TestConnectLeavesDefaultTransportTLSConfigAlone verifies that building the
// default HTTP client does not mutate the tls.Config held by
// http.DefaultTransport. Programs that install their own config there (a custom
// root CA pool, for example) would otherwise have InsecureSkipVerify turned on
// under them, and concurrent Connect calls would race on the shared config.
func TestConnectLeavesDefaultTransportTLSConfigAlone(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("{}")) //nolint
	}))
	defer ts.Close()

	defaultTransport := http.DefaultTransport.(*http.Transport)
	original := defaultTransport.TLSClientConfig
	defer func() { defaultTransport.TLSClientConfig = original }()

	shared := &tls.Config{MinVersion: tls.VersionTLS12}
	defaultTransport.TLSClientConfig = shared

	// No HTTPClient, so gofish builds its own transport seeded from the default.
	_, err := Connect(ClientConfig{Endpoint: ts.URL, Insecure: true})
	if err != nil {
		t.Fatalf("connect failed: %v", err)
	}

	if shared.InsecureSkipVerify {
		t.Error("InsecureSkipVerify was set on http.DefaultTransport's tls.Config")
	}
	if len(shared.NextProtos) != 0 {
		t.Errorf("NextProtos on http.DefaultTransport's tls.Config was modified: %v", shared.NextProtos)
	}
}

// autoExpandServiceRoot is a minimal ServiceRoot advertising all three $expand
// dialects, $levels support, and a MaxLevels cap - enough to exercise every
// branch of AutoExpand/AutoExpandLevels.
const autoExpandServiceRoot = `{
	"ProtocolFeaturesSupported": {
		"ExpandQuery": {
			"ExpandAll": true,
			"Levels": true,
			"Links": true,
			"MaxLevels": 2,
			"NoLinks": true
		}
	}
}`

func connectWithAutoExpand(t *testing.T, body string, levels uint) *APIClient {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(body)) //nolint
	}))
	t.Cleanup(ts.Close)

	client, err := Connect(ClientConfig{
		Endpoint:         ts.URL,
		Insecure:         true,
		AutoExpand:       true,
		AutoExpandLevels: levels,
	})
	if err != nil {
		t.Fatalf("connect failed: %v", err)
	}

	return client
}

// TestAutoExpandLevelsAppliesLevelsWhenAdvertised asks for a level-2 expand
// against a service that advertises Levels support up to MaxLevels 2: the
// built collection query should carry both $expand and $levels.
func TestAutoExpandLevelsAppliesLevelsWhenAdvertised(t *testing.T) {
	client := connectWithAutoExpand(t, autoExpandServiceRoot, 2)

	query := schemas.BuildQuery(client, "/redfish/v1/Chassis", true)
	if !strings.Contains(query, "$expand=.") {
		t.Errorf("expected $expand=. in query, got %q", query)
	}
	if !strings.Contains(query, "$levels=2") {
		t.Errorf("expected $levels=2 in query, got %q", query)
	}
}

// TestAutoExpandLevelsClampsToMaxLevels asks for a level deeper than the
// service's advertised MaxLevels: the built query should be clamped down to
// MaxLevels rather than requesting a depth the service didn't advertise.
func TestAutoExpandLevelsClampsToMaxLevels(t *testing.T) {
	client := connectWithAutoExpand(t, autoExpandServiceRoot, 5)

	query := schemas.BuildQuery(client, "/redfish/v1/Chassis", true)
	if !strings.Contains(query, "$levels=2") {
		t.Errorf("expected $levels clamped to MaxLevels=2, got %q", query)
	}
}

// TestAutoExpandLevelsIgnoredWithoutLevelsSupport asks for a level-2 expand
// against a service that supports $expand but not $levels: the built query
// should carry $expand alone, with no $levels term.
func TestAutoExpandLevelsIgnoredWithoutLevelsSupport(t *testing.T) {
	const body = `{
		"ProtocolFeaturesSupported": {
			"ExpandQuery": {"NoLinks": true}
		}
	}`

	client := connectWithAutoExpand(t, body, 2)

	query := schemas.BuildQuery(client, "/redfish/v1/Chassis", true)
	if !strings.Contains(query, "$expand=.") {
		t.Errorf("expected $expand=. in query, got %q", query)
	}
	if strings.Contains(query, "$levels") {
		t.Errorf("expected no $levels in query since the service doesn't support it, got %q", query)
	}
}

// TestAutoExpandLevelsBoundsConversion asks for a nonsensical depth against a
// service that advertises $levels without the MaxLevels the spec requires, so
// nothing clamps it. The depth must still come out positive rather than
// wrapping negative on the way to an int.
func TestAutoExpandLevelsBoundsConversion(t *testing.T) {
	const body = `{
		"ProtocolFeaturesSupported": {
			"ExpandQuery": {"NoLinks": true, "Levels": true}
		}
	}`

	client := connectWithAutoExpand(t, body, math.MaxUint)

	query := schemas.BuildQuery(client, "/redfish/v1/Chassis", true)
	if !strings.Contains(query, "$levels=2147483647") {
		t.Errorf("expected $levels bounded to MaxInt32, got %q", query)
	}
}

// TestConcurrentConnectNoTLSConfigRace exercises the same shared-config path
// from several goroutines so `go test -race` catches a regression.
func TestConcurrentConnectNoTLSConfigRace(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("{}")) //nolint
	}))
	defer ts.Close()

	defaultTransport := http.DefaultTransport.(*http.Transport)
	original := defaultTransport.TLSClientConfig
	defer func() { defaultTransport.TLSClientConfig = original }()

	defaultTransport.TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := Connect(ClientConfig{Endpoint: ts.URL, Insecure: true}); err != nil {
				t.Errorf("connect failed: %v", err)
			}
		}()
	}
	wg.Wait()
}
