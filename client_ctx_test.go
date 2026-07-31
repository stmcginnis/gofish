//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestClient(ctx context.Context, ts *httptest.Server) *APIClient {
	return &APIClient{
		endpoint:   ts.URL,
		ctx:        ctx,
		HTTPClient: ts.Client(),
		sem:        make(chan bool, 1),
	}
}

func newOKServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
}

// TestPerCallContext verifies that the WithContext request variants use the
// context passed per call rather than the client's cached context, while the
// context-free shims continue to work using the cached context.
func TestPerCallContext(t *testing.T) {
	ts := newOKServer()
	defer ts.Close()

	client := newTestClient(context.Background(), ts)

	// The cached-context shim should succeed.
	resp, err := client.Get("/redfish/v1/")
	if err != nil {
		t.Fatalf("Get with cached context failed: %v", err)
	}
	_ = resp.Body.Close()

	// A per-call cancelled context should abort the request with the context's
	// error, not an arbitrary failure.
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := client.GetWithContext(cancelledCtx, "/redfish/v1/"); !errors.Is(err, context.Canceled) {
		t.Fatalf("GetWithContext with cancelled context: want context.Canceled, got %v", err)
	}

	// The cached context must be unaffected by the per-call cancellation.
	resp, err = client.Get("/redfish/v1/")
	if err != nil {
		t.Fatalf("Get after per-call cancellation failed: %v", err)
	}
	_ = resp.Body.Close()
}

// TestPerCallContextInverse verifies the opposite direction: a dead cached
// context must fail the context-free shims, while a live per-call context
// keeps the WithContext variants working. This catches shims that hardcode
// the cached context.
func TestPerCallContextInverse(t *testing.T) {
	ts := newOKServer()
	defer ts.Close()

	deadCtx, cancel := context.WithCancel(context.Background())
	cancel()
	client := newTestClient(deadCtx, ts)

	if _, err := client.Get("/redfish/v1/"); !errors.Is(err, context.Canceled) {
		t.Fatalf("Get with dead cached context: want context.Canceled, got %v", err)
	}

	resp, err := client.GetWithContext(context.Background(), "/redfish/v1/")
	if err != nil {
		t.Fatalf("GetWithContext with live per-call context failed: %v", err)
	}
	_ = resp.Body.Close()
}

// TestMidFlightCancellation verifies that cancelling the per-call context
// while the request is blocked on the server aborts it with context.Canceled.
func TestMidFlightCancellation(t *testing.T) {
	arrived := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		close(arrived)
		// Block until the client gives up on the request.
		<-r.Context().Done()
	}))
	defer ts.Close()

	client := newTestClient(context.Background(), ts)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-arrived
		cancel()
	}()

	if _, err := client.GetWithContext(ctx, "/redfish/v1/"); !errors.Is(err, context.Canceled) {
		t.Fatalf("mid-flight cancellation: want context.Canceled, got %v", err)
	}
}

// TestPerCallContextOtherVerbs covers the POST and multipart request paths,
// which route through different internals than GET.
func TestPerCallContextOtherVerbs(t *testing.T) {
	ts := newOKServer()
	defer ts.Close()

	client := newTestClient(context.Background(), ts)

	resp, err := client.PostWithContext(context.Background(), "/redfish/v1/x", map[string]string{"a": "b"})
	if err != nil {
		t.Fatalf("PostWithContext failed: %v", err)
	}
	_ = resp.Body.Close()

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	payload := map[string]io.Reader{"file": strings.NewReader("data")}
	if _, err := client.PostMultipartWithContext(cancelledCtx, "/redfish/v1/x", payload); !errors.Is(err, context.Canceled) {
		t.Fatalf("PostMultipartWithContext with cancelled context: want context.Canceled, got %v", err)
	}
}
