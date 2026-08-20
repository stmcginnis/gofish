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
	"syscall"
	"time"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/schemas"
)

const (
	defaultIDRACMonitorTimeout       = 30 * time.Second
	defaultIDRACMonitorMaxRetries    = 3
	defaultIDRACMonitorRetryInterval = time.Second
)

// IDRACMonitor monitors iDRAC responsiveness and can trigger resets.
type IDRACMonitor struct {
	service        *schemas.Entity
	manager        *Manager
	timeout        time.Duration
	maxRetries     int
	retryInterval  time.Duration
	resetOnTimeout bool
}

// IDRACMonitorConfig configures the iDRAC monitor.
type IDRACMonitorConfig struct {
	// Timeout is the maximum duration of an individual health check. A value of
	// zero disables the per-request timeout.
	Timeout time.Duration
	// MaxRetries is the number of retries after the initial attempt.
	MaxRetries int
	// RetryInterval is the base delay between attempts. Each subsequent delay is
	// increased linearly. A value of zero disables the delay.
	RetryInterval time.Duration
	// ResetOnTimeout enables a graceful iDRAC reset after all health checks fail.
	ResetOnTimeout bool
}

// NewIDRACMonitor creates a new iDRAC monitor. If service is nil, the manager
// resource is used for health checks.
func NewIDRACMonitor(service *schemas.Entity, manager *Manager, config *IDRACMonitorConfig) *IDRACMonitor {
	effectiveConfig := IDRACMonitorConfig{
		Timeout:       defaultIDRACMonitorTimeout,
		MaxRetries:    defaultIDRACMonitorMaxRetries,
		RetryInterval: defaultIDRACMonitorRetryInterval,
	}
	if config != nil {
		effectiveConfig = *config
	}

	return &IDRACMonitor{
		service:        service,
		manager:        manager,
		timeout:        effectiveConfig.Timeout,
		maxRetries:     max(effectiveConfig.MaxRetries, 0),
		retryInterval:  max(effectiveConfig.RetryInterval, 0),
		resetOnTimeout: effectiveConfig.ResetOnTimeout,
	}
}

// NewiDRACMonitor creates a new iDRAC monitor.
// Deprecated: Use NewIDRACMonitor.
func NewiDRACMonitor(service *schemas.Entity, manager *Manager, config *IDRACMonitorConfig) *IDRACMonitor {
	return NewIDRACMonitor(service, manager, config)
}

// CheckHealth verifies that the configured iDRAC resource responds to a GET.
func (m *IDRACMonitor) CheckHealth(ctx context.Context) error {
	if ctx == nil {
		return errors.New("health check failed: context is nil")
	}

	resource, err := m.healthResource()
	if err != nil {
		return err
	}
	resourceClient := resource.GetClient()
	if resourceClient == nil {
		return errors.New("health check failed: resource has no client")
	}
	client, ok := resourceClient.(*gofish.APIClient)
	if !ok {
		return errors.New("health check failed: resource client is not an APIClient")
	}

	requestContext := ctx
	if m.timeout > 0 {
		var cancel context.CancelFunc
		requestContext, cancel = context.WithTimeout(ctx, m.timeout)
		defer cancel()
	}

	resp, err := client.WithContext(requestContext).Get(resource.ODataID)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}

	return nil
}

func (m *IDRACMonitor) healthResource() (*schemas.Entity, error) {
	if m.service != nil {
		if m.service.ODataID == "" {
			return nil, errors.New("health check failed: service has no @odata.id")
		}
		return m.service, nil
	}
	if m.manager == nil {
		return nil, errors.New("health check failed: manager is not configured")
	}
	if m.manager.ODataID == "" {
		return nil, errors.New("health check failed: manager has no @odata.id")
	}
	return &m.manager.Entity, nil
}

// ExecuteWithRetry executes an operation after checking iDRAC health. Network
// and service availability errors are retried according to the monitor config.
func (m *IDRACMonitor) ExecuteWithRetry(ctx context.Context, operation func() error) error {
	if operation == nil {
		return errors.New("operation is nil")
	}

	attempts := m.maxRetries + 1
	for attempt := 0; attempt < attempts; attempt++ {
		if err := m.CheckHealth(ctx); err != nil {
			if !isRetryableError(err) {
				return err
			}
			if attempt == attempts-1 {
				return m.handleUnhealthyIDRAC(attempts, err)
			}
			if err := m.waitForRetry(ctx, attempt); err != nil {
				return err
			}
			continue
		}

		err := operation()
		if err == nil {
			return nil
		}
		if !isRetryableError(err) {
			return err
		}
		if attempt == attempts-1 {
			return fmt.Errorf("operation failed after %d attempts: %w", attempts, err)
		}
		if err := m.waitForRetry(ctx, attempt); err != nil {
			return err
		}
	}

	return nil
}

func (m *IDRACMonitor) handleUnhealthyIDRAC(attempts int, healthErr error) error {
	if !m.resetOnTimeout {
		return fmt.Errorf("iDRAC unresponsive after %d attempts: %w", attempts, healthErr)
	}
	if m.manager == nil {
		return fmt.Errorf("iDRAC unresponsive after %d attempts and reset is unavailable: %w", attempts, healthErr)
	}
	if err := m.manager.ResetiDRAC(GracefuliDRACReset); err != nil {
		return fmt.Errorf("iDRAC unresponsive after %d attempts, reset also failed: %w", attempts, err)
	}
	return fmt.Errorf("iDRAC unresponsive after %d attempts; graceful reset performed: %w", attempts, healthErr)
}

func (m *IDRACMonitor) waitForRetry(ctx context.Context, attempt int) error {
	delay := time.Duration(attempt+1) * m.retryInterval
	if delay <= 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return nil
		}
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// isRetryableError determines whether an operation error is transient.
func isRetryableError(err error) bool {
	if err == nil || errors.Is(err, context.Canceled) {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	var redfishErr *schemas.Error
	if errors.As(err, &redfishErr) {
		status := redfishErr.HTTPReturnedStatusCode
		return status == http.StatusRequestTimeout || status == http.StatusTooManyRequests || status >= http.StatusInternalServerError
	}

	var networkErr net.Error
	if errors.As(err, &networkErr) && networkErr.Timeout() {
		return true
	}

	return errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.ENETUNREACH) ||
		errors.Is(err, syscall.EHOSTUNREACH)
}
