//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stmcginnis/gofish/redfish"
)

func TestIDRACReset(t *testing.T) {
	// This would require a mock server, for now just test the basic functionality exists
	// The actual implementation would need integration testing with a real or mocked iDRAC
	t.Skip("Skipping iDRAC reset test - requires integration testing with real iDRAC or comprehensive mocking")
}

func TestIDRACMonitor_CheckHealth(t *testing.T) {
	// Create a mock manager
	baseManager := redfish.Manager{}
	baseManager.ID = "iDRAC.Embedded.1"
	manager := &Manager{Manager: baseManager}

	// Create monitor
	config := &iDRACMonitorConfig{
		Timeout:        5 * time.Second,
		MaxRetries:     2,
		ResetOnTimeout: false,
	}

	monitor := NewiDRACMonitor(nil, manager, config)

	// Test with valid manager
	ctx := context.Background()
	err := monitor.CheckHealth(ctx)
	if err != nil {
		t.Errorf("Health check failed unexpectedly: %v", err)
	}
}

func TestIDRACMonitor_ExecuteWithRetry(t *testing.T) {
	baseManager := redfish.Manager{}
	baseManager.ID = "iDRAC.Embedded.1"
	manager := &Manager{Manager: baseManager}

	config := &iDRACMonitorConfig{
		Timeout:        1 * time.Second,
		MaxRetries:     1,
		ResetOnTimeout: false,
	}

	monitor := NewiDRACMonitor(nil, manager, config)

	ctx := context.Background()

	// Test successful operation
	callCount := 0
	err := monitor.ExecuteWithRetry(ctx, func() error {
		callCount++
		return nil
	})

	if err != nil {
		t.Errorf("Expected operation to succeed, got error: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected operation to be called once, got %d calls", callCount)
	}
}

func TestIsRetryableError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil error", nil, false},
		{"timeout error", context.DeadlineExceeded, true},
		{"connection refused", errors.New("connection refused"), true},
		{"network unreachable", errors.New("network is unreachable"), true},
		{"i/o timeout", errors.New("i/o timeout"), true},
		{"regular error", errors.New("some other error"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isRetryableError(tt.err)
			if result != tt.expected {
				t.Errorf("isRetryableError(%v) = %v, expected %v", tt.err, result, tt.expected)
			}
		})
	}
}
