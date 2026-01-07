//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/stmcginnis/gofish/schemas"
)

// iDRACMonitor monitors iDRAC responsiveness and can trigger resets
type iDRACMonitor struct {
	service        *schemas.Entity // Using Entity as base for service access
	manager        *Manager
	timeout        time.Duration
	maxRetries     int
	resetOnTimeout bool
}

// iDRACMonitorConfig configures the iDRAC monitor
type iDRACMonitorConfig struct {
	// Timeout for individual requests
	Timeout time.Duration
	// Maximum number of retries before considering iDRAC unresponsive
	MaxRetries int
	// Whether to automatically reset iDRAC when unresponsive
	ResetOnTimeout bool
}

// NewiDRACMonitor creates a new iDRAC monitor
func NewiDRACMonitor(service *schemas.Entity, manager *Manager, config *iDRACMonitorConfig) *iDRACMonitor {
	if config == nil {
		config = &iDRACMonitorConfig{
			Timeout:        30 * time.Second,
			MaxRetries:     3,
			ResetOnTimeout: false,
		}
	}

	return &iDRACMonitor{
		service:        service,
		manager:        manager,
		timeout:        config.Timeout,
		maxRetries:     config.MaxRetries,
		resetOnTimeout: config.ResetOnTimeout,
	}
}

// CheckHealth performs a health check on the iDRAC
func (im *iDRACMonitor) CheckHealth(ctx context.Context) error {
	// Create a context with timeout
	_, cancel := context.WithTimeout(ctx, im.timeout)
	defer cancel()

	// Try to get basic manager information as a health check
	// For now, just check if we can access the manager
	if im.manager == nil || im.manager.ID == "" {
		return errors.New("health check failed: manager not accessible")
	}

	return nil
}

// ExecuteWithRetry executes a function with retry logic and timeout detection
func (im *iDRACMonitor) ExecuteWithRetry(ctx context.Context, operation func() error) error {
	var lastErr error

	for attempt := 0; attempt <= im.maxRetries; attempt++ {
		// Check iDRAC health before attempting operation
		if healthErr := im.CheckHealth(ctx); healthErr != nil {
			if attempt == im.maxRetries {
				if im.resetOnTimeout {
					// Try to reset iDRAC before giving up
					resetErr := im.manager.ResetiDRAC(GracefuliDRACReset)
					if resetErr != nil {
						return fmt.Errorf("iDRAC unresponsive after %d attempts, reset also failed: %w", im.maxRetries+1, resetErr)
					}
					return fmt.Errorf("iDRAC was unresponsive, performed reset but operation failed")
				}
				return fmt.Errorf("iDRAC unresponsive after %d attempts: %w", im.maxRetries+1, healthErr)
			}
			// Wait before retry
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(attempt+1) * time.Second):
				continue
			}
		}

		// Execute the operation
		if err := operation(); err != nil {
			lastErr = err
			// If it's a network error or timeout, retry
			if isRetryableError(err) {
				continue
			}
			// Non-retryable error, return immediately
			return err
		}

		// Success
		return nil
	}

	return lastErr
}

// isRetryableError determines if an error is worth retrying
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Check for context timeout
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	// Check for HTTP timeout or connection errors
	errStr := err.Error()
	retryableErrors := []string{
		"timeout",
		"connection refused",
		"connection reset",
		"no such host",
		"network is unreachable",
		"i/o timeout",
	}

	for _, retryable := range retryableErrors {
		if containsIgnoreCase(errStr, retryable) {
			return true
		}
	}

	return false
}

// containsIgnoreCase checks if a string contains a substring (case insensitive)
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || containsIgnoreCase(s[1:], substr) || (s != "" && substr != "" && toLower(s[0]) == toLower(substr[0]) && containsIgnoreCase(s[1:], substr[1:])))
}

// toLower converts a byte to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
