//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// ActionTarget is contains the target endpoint for object Actions.
type ActionTarget struct {
	Target string
}

// Action contains the target and ActionInfo endpoints for object Actions.
type Action struct {
	Target string
	Info   string `json:"@Redfish.ActionInfo"`
}
