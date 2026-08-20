//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/stmcginnis/gofish/schemas"
)

// IDRACResetType defines the type of reset to perform.
type IDRACResetType string

const (
	// GracefuliDRACReset performs a graceful reset of the iDRAC.
	GracefuliDRACReset IDRACResetType = "Graceful"
	// ForceiDRACReset performs a forced reset of the iDRAC.
	ForceiDRACReset IDRACResetType = "Force"
)

// iDRACResetRequest represents the body for the iDRAC reset action.
type iDRACResetRequest struct {
	Force IDRACResetType `json:"Force"`
}

// IDRACCardService represents Dell's DelliDRACCardService OEM resource.
type IDRACCardService struct {
	schemas.Entity

	iDRACResetTarget string
}

// UnmarshalJSON unmarshals an IDRACCardService and its action targets.
func (s *IDRACCardService) UnmarshalJSON(b []byte) error {
	type temp IDRACCardService
	var payload struct {
		temp
		Actions struct {
			IDRACReset schemas.ActionTarget `json:"#DelliDRACCardService.iDRACReset"`
		}
	}

	if err := json.Unmarshal(b, &payload); err != nil {
		return err
	}

	*s = IDRACCardService(payload.temp)
	s.iDRACResetTarget = payload.Actions.IDRACReset.Target
	return nil
}

// Reset performs a graceful or forced iDRAC reset. A non-nil task monitor is
// returned when the service accepts the reset asynchronously.
func (s *IDRACCardService) Reset(resetType IDRACResetType) (*schemas.TaskMonitorInfo, error) {
	if resetType != GracefuliDRACReset && resetType != ForceiDRACReset {
		return nil, fmt.Errorf("invalid iDRAC reset type: %s", resetType)
	}
	if s.iDRACResetTarget == "" {
		return nil, errors.New("iDRAC reset is not supported by this service")
	}

	request := iDRACResetRequest{Force: resetType}
	resp, taskInfo, err := schemas.PostWithTask(
		s.GetClient(), s.iDRACResetTarget, request, s.Headers(), false,
	)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to reset iDRAC: %w", err)
	}

	return taskInfo, nil
}

// GetIDRACCardService gets an IDRACCardService from the service.
func GetIDRACCardService(c schemas.Client, uri string) (*IDRACCardService, error) {
	return schemas.GetObject[IDRACCardService](c, uri)
}

// ResetIDRAC performs a graceful or forced reset using the action advertised
// by Dell's iDRAC card service. A non-nil task monitor indicates an asynchronous
// response.
func (m *Manager) ResetIDRAC(resetType IDRACResetType) (*schemas.TaskMonitorInfo, error) {
	service, err := m.IDRACCardService()
	if err != nil {
		return nil, err
	}

	return service.Reset(resetType)
}

// ResetiDRAC performs a reset of the iDRAC card and retains the original API.
// Use ResetIDRAC when the asynchronous task monitor is needed.
func (m *Manager) ResetiDRAC(resetType IDRACResetType) error {
	_, err := m.ResetIDRAC(resetType)
	return err
}
