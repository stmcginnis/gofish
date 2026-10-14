//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetSupportedResetTypes returns any reset types that the Manager declares as supported
// via either ActionInfo or AllowableValues.
func (m *Manager) GetSupportedResetTypes() ([]ResetType, error) {
	if len(m.SupportedResetTypes) > 0 {
		return m.SupportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if m.resetActionInfo != "" {
		resetActionInfo, err := m.ResetActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			m.SupportedResetTypes = append(m.SupportedResetTypes, ResetType(val))
		}
	}

	return m.SupportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the Manager reset action if supported
func (m *Manager) ResetActionInfo() (*ActionInfo, error) {
	if m.resetActionInfo == "" {
		return nil, errors.New("Manager Reset ActionInfo not supported")
	}

	return GetObject[ActionInfo](m.GetClient(), m.resetActionInfo)
}

// This action shall reset the manager. If this manager provides the Redfish
// service, the service shall send the action response before resetting to
// prevent client timeouts.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Manager) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	if len(m.SupportedResetTypes) == 0 {
		if m.resetActionInfo != "" {
			// reset without confirming the type is supported by the manager.
			// done to minimize overhead though technically not as correct as first checking the supported reset types
			payload := struct {
				ResetType ResetType
			}{ResetType: resetType}
			resp, taskInfo, err := PostWithTask(m.client,
				m.resetTarget, payload, m.Headers(), false)
			defer DeferredCleanupHTTPResponse(resp)
			return taskInfo, err
		}
		// reset directly without reset type. HPE server has the behavior
		resp, taskInfo, err := PostWithTask(m.client,
			m.resetTarget, struct{}{}, m.Headers(), false)
		defer DeferredCleanupHTTPResponse(resp)
		return taskInfo, err
	}
	// Make sure the requested reset type is supported by the manager.
	valid := false
	for _, allowed := range m.SupportedResetTypes {
		if resetType == allowed {
			valid = true
			break
		}
	}

	if !valid {
		return nil, fmt.Errorf("reset type '%s' is not supported by this manager",
			resetType)
	}

	payload := struct {
		ResetType ResetType
	}{ResetType: resetType}
	resp, taskInfo, err := PostWithTask(m.client,
		m.resetTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// GetSupportedResetToDefaultsTypes returns any reset to defaults
// types that the Manager declares as supported via either ActionInfo or AllowableValues.
func (m *Manager) GetSupportedResetToDefaultsTypes() ([]ResetToDefaultsType, error) {
	if len(m.SupportedResetToDefaultsTypes) > 0 {
		return m.SupportedResetToDefaultsTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if m.resetToDefaultsActionInfoTarget != "" {
		resetActionInfo, err := m.ResetToDefaultsActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			m.SupportedResetToDefaultsTypes = append(m.SupportedResetToDefaultsTypes, ResetToDefaultsType(val))
		}
	}

	return m.SupportedResetToDefaultsTypes, nil
}

// ResetToDefaultsActionInfo returns the ActionInfo for the Manager ResetToDefaults action if supported
func (m *Manager) ResetToDefaultsActionInfo() (*ActionInfo, error) {
	if m.resetToDefaultsActionInfoTarget == "" {
		return nil, errors.New("Manager ResetToDefaults ActionInfo not supported")
	}

	return GetObject[ActionInfo](m.GetClient(), m.resetToDefaultsActionInfoTarget)
}

// postUnmarshal records the allowed reset types, their ActionInfo URIs, and the
// OEM section of Links.
func (m *Manager) postUnmarshal(b []byte) error {
	var t struct {
		Actions struct {
			Reset struct {
				ActionTarget
				AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			} `json:"#Manager.Reset"`
			ResetToDefaults struct {
				ActionTarget
				AllowedResetTypes []ResetToDefaultsType `json:"ResetType@Redfish.AllowableValues"`
			} `json:"#Manager.ResetToDefaults"`
		}
		Links struct {
			OEM json.RawMessage `json:"Oem"`
		}
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	m.SupportedResetTypes = t.Actions.Reset.AllowedResetTypes
	m.resetActionInfo = t.Actions.Reset.ActionInfoTarget
	m.SupportedResetToDefaultsTypes = t.Actions.ResetToDefaults.AllowedResetTypes
	m.resetToDefaultsActionInfoTarget = t.Actions.ResetToDefaults.ActionInfoTarget
	m.OEMLinks = t.Links.OEM

	return nil
}
