//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/schemas"
)

const (
	idracManagerURI        = "/redfish/v1/Managers/iDRAC.Embedded.1"
	monitorServiceRootBody = "{\"@odata.id\":\"/redfish/v1/\",\"Id\":\"RootService\",\"Name\":\"Root Service\"}"
)

func TestIDRACMonitorCheckHealth(t *testing.T) {
	var requestedURI string
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, r *http.Request) {
		requestedURI = r.URL.Path
		writeHealthOK(t, w)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		Timeout: time.Second,
	})

	if err := monitor.CheckHealth(context.Background()); err != nil {
		t.Fatalf("health check failed: %v", err)
	}
	if requestedURI != idracManagerURI {
		t.Errorf("health check requested %q, expected %q", requestedURI, idracManagerURI)
	}
}

func TestIDRACMonitorCheckHealthTimeout(t *testing.T) {
	client := newMonitorAPIClient(t, func(_ http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		Timeout: 10 * time.Millisecond,
	})

	err := monitor.CheckHealth(context.Background())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestIDRACMonitorCheckHealthValidation(t *testing.T) {
	managerWithoutClient := newMonitorManager(nil)
	managerWithUnsupportedClient := newMonitorManager(&schemas.TestClient{})

	tests := []struct {
		name    string
		monitor *IDRACMonitor
		ctx     context.Context
		want    string
	}{
		{
			name:    "nil context",
			monitor: NewIDRACMonitor(nil, managerWithoutClient, nil),
			want:    "context is nil",
		},
		{
			name:    "missing manager",
			monitor: NewIDRACMonitor(nil, nil, nil),
			ctx:     context.Background(),
			want:    "manager is not configured",
		},
		{
			name: "missing resource URI",
			monitor: NewIDRACMonitor(nil, &Manager{
				Manager: schemas.Manager{},
			}, nil),
			ctx:  context.Background(),
			want: "manager has no @odata.id",
		},
		{
			name:    "missing client",
			monitor: NewIDRACMonitor(nil, managerWithoutClient, nil),
			ctx:     context.Background(),
			want:    "resource has no client",
		},
		{
			name:    "client is not APIClient",
			monitor: NewIDRACMonitor(nil, managerWithUnsupportedClient, nil),
			ctx:     context.Background(),
			want:    "resource client is not an APIClient",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.monitor.CheckHealth(test.ctx)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected error containing %q, got %v", test.want, err)
			}
		})
	}
}

func TestIDRACMonitorExecuteWithRetry(t *testing.T) {
	var healthChecks atomic.Int32
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, _ *http.Request) {
		healthChecks.Add(1)
		writeHealthOK(t, w)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		MaxRetries: 2,
	})

	operations := 0
	err := monitor.ExecuteWithRetry(context.Background(), func() error {
		operations++
		if operations < 3 {
			return context.DeadlineExceeded
		}
		return nil
	})
	if err != nil {
		t.Fatalf("operation failed: %v", err)
	}
	if operations != 3 {
		t.Errorf("operation called %d times, expected 3", operations)
	}
	if healthChecks.Load() != 3 {
		t.Errorf("health checked %d times, expected 3", healthChecks.Load())
	}
}

func TestIDRACMonitorExecuteWithRetryStopsOnPermanentError(t *testing.T) {
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, _ *http.Request) {
		writeHealthOK(t, w)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		MaxRetries: 3,
	})

	operations := 0
	expectedErr := errors.New("invalid request")
	err := monitor.ExecuteWithRetry(context.Background(), func() error {
		operations++
		return expectedErr
	})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected permanent operation error, got %v", err)
	}
	if operations != 1 {
		t.Errorf("operation called %d times, expected 1", operations)
	}
}

func TestIDRACMonitorExecuteWithRetryHonorsCancellation(t *testing.T) {
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, _ *http.Request) {
		writeHealthOK(t, w)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		MaxRetries:    2,
		RetryInterval: time.Hour,
	})
	ctx, cancel := context.WithCancel(context.Background())

	operations := 0
	err := monitor.ExecuteWithRetry(ctx, func() error {
		operations++
		cancel()
		return context.DeadlineExceeded
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
	if operations != 1 {
		t.Errorf("operation called %d times, expected 1", operations)
	}
}

func TestIDRACMonitorUnhealthyRetries(t *testing.T) {
	var healthChecks atomic.Int32
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, _ *http.Request) {
		healthChecks.Add(1)
		http.Error(w, "temporarily unavailable", http.StatusServiceUnavailable)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		MaxRetries: 2,
	})

	operations := 0
	err := monitor.ExecuteWithRetry(context.Background(), func() error {
		operations++
		return nil
	})
	if err == nil || !strings.Contains(err.Error(), "iDRAC unresponsive after 3 attempts") {
		t.Fatalf("expected unresponsive error, got %v", err)
	}
	if healthChecks.Load() != 3 {
		t.Errorf("health checked %d times, expected 3", healthChecks.Load())
	}
	if operations != 0 {
		t.Errorf("operation called %d times, expected 0", operations)
	}
}

func TestIDRACMonitorUnhealthyStopsOnPermanentError(t *testing.T) {
	var healthChecks atomic.Int32
	client := newMonitorAPIClient(t, func(w http.ResponseWriter, _ *http.Request) {
		healthChecks.Add(1)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	})
	monitor := NewIDRACMonitor(nil, newMonitorManager(client), &IDRACMonitorConfig{
		MaxRetries: 2,
	})

	err := monitor.ExecuteWithRetry(context.Background(), func() error {
		t.Fatal("operation should not run after a failed health check")
		return nil
	})
	var redfishErr *schemas.Error
	if !errors.As(err, &redfishErr) || redfishErr.HTTPReturnedStatusCode != http.StatusUnauthorized {
		t.Fatalf("expected unauthorized Redfish error, got %v", err)
	}
	if healthChecks.Load() != 1 {
		t.Errorf("health checked %d times, expected 1", healthChecks.Load())
	}
}

func TestIsRetryableError(t *testing.T) {
	timeoutError := &net.DNSError{IsTimeout: true}
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "nil", err: nil},
		{name: "canceled context", err: context.Canceled},
		{name: "deadline exceeded", err: context.DeadlineExceeded, want: true},
		{name: "network timeout", err: timeoutError, want: true},
		{name: "connection refused", err: fmt.Errorf("dial failed: %w", syscall.ECONNREFUSED), want: true},
		{name: "HTTP 429", err: schemas.ConstructError(http.StatusTooManyRequests, nil), want: true},
		{name: "HTTP 503", err: schemas.ConstructError(http.StatusServiceUnavailable, nil), want: true},
		{name: "HTTP 400", err: schemas.ConstructError(http.StatusBadRequest, nil)},
		{name: "permanent error", err: errors.New("invalid request")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isRetryableError(test.err); got != test.want {
				t.Errorf("isRetryableError(%v) = %t, expected %t", test.err, got, test.want)
			}
		})
	}
}

func newMonitorAPIClient(t *testing.T, healthHandler http.HandlerFunc) *gofish.APIClient {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case schemas.DefaultServiceRoot:
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(monitorServiceRootBody)); err != nil {
				t.Errorf("failed to write service root response: %v", err)
			}
		case idracManagerURI:
			healthHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	client, err := gofish.Connect(gofish.ClientConfig{
		Endpoint:   server.URL,
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("failed to create API client: %v", err)
	}
	return client
}

func newMonitorManager(client schemas.Client) *Manager {
	manager := schemas.Manager{}
	manager.ODataID = idracManagerURI
	manager.ID = "iDRAC.Embedded.1"
	manager.SetClient(client)
	return &Manager{Manager: manager}
}

func writeHealthOK(t *testing.T, w http.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte("{}")); err != nil {
		t.Errorf("failed to write health response: %v", err)
	}
}
