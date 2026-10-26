//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"reflect"
	"strings"
)

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

func (bi *Bios) ReportedAttributeUpdateApplyTimes() []SettingsApplyTime {
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

// postUnmarshal records the settings resource from @Redfish.Settings and keeps
// the raw payload for UpdateBiosAttributes.
func (bi *Bios) postUnmarshal(b []byte) error {
	var t struct {
		Settings Settings `json:"@Redfish.Settings"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	bi.settingsApplyTimes = t.Settings.SupportedApplyTimes
	bi.settingsTarget = t.Settings.SettingsObject
	if bi.settingsTarget == "" {
		bi.settingsTarget = bi.ODataID
	}
	bi.rawData = b

	return nil
}
