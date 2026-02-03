//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"errors"
	"strconv"

	"github.com/stmcginnis/gofish/schemas"
)

// iDRACResetType defines the type of reset to perform
type iDRACResetType string

const (
	// GracefuliDRACReset performs a graceful reset of the iDRAC
	GracefuliDRACReset iDRACResetType = "Graceful"
	// ForceiDRACReset performs a forced reset of the iDRAC
	ForceiDRACReset iDRACResetType = "Force"
)

// iDRACResetRequest represents the body for iDRAC reset action
type iDRACResetRequest struct {
	Force iDRACResetType `json:"Force"`
}

// ResetiDRAC performs a reset of the iDRAC card
//
// resetType specifies whether to perform a graceful or forced reset
func (m *Manager) ResetiDRAC(resetType iDRACResetType) error {
	request := iDRACResetRequest{
		Force: resetType,
	}

	// Use the standard action target for iDRAC reset as documented
	target := "/redfish/v1/Dell/Managers/iDRAC.Embedded.1/DelliDRACCardService/Actions/DelliDRACCardService.iDRACReset"

	resp, err := m.PostWithResponse(target, request)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return errors.New("failed to reset iDRAC: " + err.Error())
	}

	// Check for successful response codes (200, 201, etc.)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("iDRAC reset failed with status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
