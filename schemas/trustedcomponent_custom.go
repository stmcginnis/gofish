//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a TrustedComponent object from the raw JSON.
func (t *TrustedComponent) UnmarshalJSON(b []byte) error {
	type temp TrustedComponent
	type tActions struct {
		TPMClear       ActionTarget `json:"#TrustedComponent.TPMClear"`
		TPMGetEventLog ActionTarget `json:"#TrustedComponent.TPMGetEventLog"`
	}
	type tLinks struct {
		ActiveSoftwareImage Link  `json:"ActiveSoftwareImage"`
		ComponentIntegrity  Links `json:"ComponentIntegrity"`
		ComponentsProtected Links `json:"ComponentsProtected"`
		IntegratedInto      Link  `json:"IntegratedInto"`
		Owner               Link  `json:"Owner"`
		SoftwareImages      Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions      tActions
		Links        tLinks
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TrustedComponent(tmp.temp)

	// Extract the links to other entities for later
	t.tpmClearTarget = tmp.Actions.TPMClear.Target
	t.tpmGetEventLogTarget = tmp.Actions.TPMGetEventLog.Target
	t.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	t.componentIntegrity = tmp.Links.ComponentIntegrity.ToStrings()
	t.componentsProtected = tmp.Links.ComponentsProtected.ToStrings()
	t.integratedInto = tmp.Links.IntegratedInto.String()
	t.owner = tmp.Links.Owner.String()
	t.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	t.certificates = tmp.Certificates.String()

	return nil
}

// This action shall reset all user-accessible TPM settings to factory
// defaults.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TrustedComponent) TPMClear() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.tpmClearTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall return the event log for TPM 2.0 devices.
func (t *TrustedComponent) TPMGetEventLog() (*TPMGetEventLogResponse, error) {
	payload := make(map[string]any)

	resp, err := t.PostWithResponse(t.tpmGetEventLogTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result TPMGetEventLogResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
