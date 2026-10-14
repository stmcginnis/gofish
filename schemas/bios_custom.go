//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"reflect"
	"strings"
)

// UnmarshalJSON unmarshals a Bios object from the raw JSON.
func (bi *Bios) UnmarshalJSON(b []byte) error {
	type temp Bios
	type biActions struct {
		ChangePassword ActionTarget `json:"#Bios.ChangePassword"`
		ResetBios      ActionTarget `json:"#Bios.ResetBios"`
	}
	type biLinks struct {
		ActiveSoftwareImage Link  `json:"ActiveSoftwareImage"`
		SoftwareImages      Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions  biActions
		Links    biLinks
		Settings Settings `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*bi = Bios(tmp.temp)

	// Extract the links to other entities for later
	bi.changePasswordTarget = tmp.Actions.ChangePassword.Target
	bi.resetBiosTarget = tmp.Actions.ResetBios.Target
	bi.settingsApplyTimes = tmp.Settings.SupportedApplyTimes
	bi.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	bi.softwareImages = tmp.Links.SoftwareImages.ToStrings()

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	bi.settingsTarget = tmp.Settings.SettingsObject
	if bi.settingsTarget == "" {
		bi.settingsTarget = bi.ODataID
	}

	// This is a read/write object, so we need to save the raw object data for later
	bi.rawData = b

	return nil
}

// AllowedAttributeUpdateApplyTimes returns the set of allowed apply times to request when
// setting the Bios attribute values.
func (bi *Bios) AllowedAttributeUpdateApplyTimes() []SettingsApplyTime {
	if len(bi.settingsApplyTimes) == 0 {
		result := []SettingsApplyTime{
			ImmediateSettingsApplyTime,
			OnResetSettingsApplyTime,
			AtMaintenanceWindowStartSettingsApplyTime,
			InMaintenanceWindowOnResetSettingsApplyTime,
		}
		return result
	}
	return bi.settingsApplyTimes
}

// UpdateBiosAttributesApplyAt is used to update attribute values and set apply time together
func (bi *Bios) UpdateBiosAttributesApplyAt(attrs SettingsAttributes, applyTime SettingsApplyTime) error {
	_, err := bi.UpdateBiosAttributesApplyAtWithTask(attrs, applyTime)
	return err
}

// UpdateBiosAttributesApplyAtWithTask is used to update attribute values and set apply time
// together. It returns a TaskMonitorInfo when the service processes the update asynchronously.
//
// Only attributes whose requested value differs from the current value reported by the settings
// target are sent. For a resource with a separate settings object that target holds the pending
// attribute values, which can differ from the applied values in bi.Attributes, so requesting the
// applied value for an attribute that has something else staged is sent as the change it is.
func (bi *Bios) UpdateBiosAttributesApplyAtWithTask(attrs SettingsAttributes, applyTime SettingsApplyTime) (*TaskMonitorInfo, error) {
	resp, err := bi.client.Get(bi.settingsTarget)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	// The update is written to the settings target, so its attributes are what the requested
	// values are compared against. A target that does not report them leaves this empty and
	// every attribute falls back to the applied values below.
	var settings struct {
		Attributes SettingsAttributes
	}
	_ = json.NewDecoder(resp.Body).Decode(&settings)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err = original.UnmarshalJSON(bi.rawData)
	if err != nil {
		return nil, err
	}

	payload := make(map[string]any)
	for key, val := range attrs {
		current, ok := settings.Attributes[key]
		if !ok {
			// Nothing staged for this attribute, so compare against the applied value.
			current = original.Attributes[key]
		}

		// DeepEqual rather than !=, which panics on the slice-valued attributes that some
		// implementations report.
		if strings.HasPrefix(key, "BootTypeOrder") || !reflect.DeepEqual(current, val) {
			payload[key] = val
		}
	}

	// If there are no allowed updates, there is nothing to send.
	if len(payload) == 0 {
		return nil, nil
	}

	data := map[string]any{"Attributes": payload}
	if applyTime != "" {
		data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
	}

	header := make(map[string]string)
	if resp.Header["Etag"] != nil {
		header["If-Match"] = resp.Header["Etag"][0]
	}

	patchResp, taskInfo, err := PatchWithTask(bi.client, bi.settingsTarget, data, header)
	defer DeferredCleanupHTTPResponse(patchResp)
	return taskInfo, err
}

// UpdateBiosAttributes is used to update attribute values.
func (b *Bios) UpdateBiosAttributes(attrs SettingsAttributes) error {
	return b.UpdateBiosAttributesApplyAt(attrs, "")
}
