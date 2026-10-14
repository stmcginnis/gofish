//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
)

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (m *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	type mActions struct {
		ForceFailover       ActionTarget `json:"#Manager.ForceFailover"`
		ModifyRedundancySet ActionTarget `json:"#Manager.ModifyRedundancySet"`
		Reset               struct {
			ActionTarget
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Manager.Reset"`
		ResetToDefaults struct {
			ActionTarget
			AllowedResetTypes []ResetToDefaultsType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Manager.ResetToDefaults"`
		UpdateSecurityMode ActionTarget `json:"#Manager.UpdateSecurityMode"`
	}
	type mLinks struct {
		ActiveSoftwareImage Link            `json:"ActiveSoftwareImage"`
		ManagedBy           Links           `json:"ManagedBy"`
		ManagerForChassis   Links           `json:"ManagerForChassis"`
		ManagerForFabrics   Links           `json:"ManagerForFabrics"`
		ManagerForManagers  Links           `json:"ManagerForManagers"`
		ManagerForServers   Links           `json:"ManagerForServers"`
		ManagerForSwitches  Links           `json:"ManagerForSwitches"`
		ManagerInChassis    Link            `json:"ManagerInChassis"`
		SelectedNetworkPort Link            `json:"SelectedNetworkPort"`
		SoftwareImages      Links           `json:"SoftwareImages"`
		OEM                 json.RawMessage `json:"Oem"`
	}
	var tmp struct {
		temp
		Actions               mActions
		Links                 mLinks
		Certificates          Link                                 `json:"Certificates"`
		DedicatedNetworkPorts Link                                 `json:"DedicatedNetworkPorts"`
		EthernetInterfaces    Link                                 `json:"EthernetInterfaces"`
		HostInterfaces        Link                                 `json:"HostInterfaces"`
		LogServices           Link                                 `json:"LogServices"`
		ManagerDiagnosticData ResolvedLink[ManagerDiagnosticData]  `json:"ManagerDiagnosticData"`
		NetworkProtocol       ResolvedLink[ManagerNetworkProtocol] `json:"NetworkProtocol"`
		Redundancy            ResolvedLink[Redundancy]             `json:"Redundancy"`
		RemoteAccountService  ResolvedLink[AccountService]         `json:"RemoteAccountService"`
		SecurityPolicy        ResolvedLink[SecurityPolicy]         `json:"SecurityPolicy"`
		SerialInterfaces      Link                                 `json:"SerialInterfaces"`
		SharedNetworkPorts    Link                                 `json:"SharedNetworkPorts"`
		USBPorts              Link                                 `json:"USBPorts"`
		VirtualMedia          Link                                 `json:"VirtualMedia"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Manager(tmp.temp)

	// Extract the links to other entities for later
	m.forceFailoverTarget = tmp.Actions.ForceFailover.Target
	m.modifyRedundancySetTarget = tmp.Actions.ModifyRedundancySet.Target
	m.resetTarget = tmp.Actions.Reset.Target
	m.SupportedResetTypes = tmp.Actions.Reset.AllowedResetTypes
	m.resetActionInfo = tmp.Actions.Reset.ActionInfoTarget
	m.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	m.SupportedResetToDefaultsTypes = tmp.Actions.ResetToDefaults.AllowedResetTypes
	m.resetToDefaultsActionInfoTarget = tmp.Actions.ResetToDefaults.ActionInfoTarget
	m.updateSecurityModeTarget = tmp.Actions.UpdateSecurityMode.Target
	m.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	m.managedBy = tmp.Links.ManagedBy.ToStrings()
	m.managerForChassis = tmp.Links.ManagerForChassis.ToStrings()
	m.managerForFabrics = tmp.Links.ManagerForFabrics.ToStrings()
	m.managerForManagers = tmp.Links.ManagerForManagers.ToStrings()
	m.managerForServers = tmp.Links.ManagerForServers.ToStrings()
	m.managerForSwitches = tmp.Links.ManagerForSwitches.ToStrings()
	m.managerInChassis = tmp.Links.ManagerInChassis.String()
	m.selectedNetworkPort = tmp.Links.SelectedNetworkPort.String()
	m.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	m.OEMLinks = tmp.Links.OEM
	m.certificates = tmp.Certificates.String()
	m.dedicatedNetworkPorts = tmp.DedicatedNetworkPorts.String()
	m.ethernetInterfaces = tmp.EthernetInterfaces.String()
	m.hostInterfaces = tmp.HostInterfaces.String()
	m.logServices = tmp.LogServices.String()
	m.managerDiagnosticData = tmp.ManagerDiagnosticData

	m.networkProtocol = tmp.NetworkProtocol

	m.redundancy = tmp.Redundancy

	m.remoteAccountService = tmp.RemoteAccountService

	m.securityPolicy = tmp.SecurityPolicy

	m.serialInterfaces = tmp.SerialInterfaces.String()
	m.sharedNetworkPorts = tmp.SharedNetworkPorts.String()
	m.uSBPorts = tmp.USBPorts.String()
	m.virtualMedia = tmp.VirtualMedia.String()

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

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
