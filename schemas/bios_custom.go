//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"reflect"
	"strconv"
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

// biosAttributeEqual reports whether a requested attribute value matches the value the
// service reports. reflect.DeepEqual handles identical types; the remaining case is a
// requested string against a reported number or boolean. The Redfish Bios schema types
// Attributes per the attribute registry (Integer and Boolean entries are JSON numbers
// and booleans), so controller-style callers carrying YAML strings otherwise rewrite an
// unchanged attribute on every reconcile. A string counts as equal only when it parses
// and formats back to itself: "007" and "1e5" never equal 7 and 100000, and integers
// above 2^53 do not round-trip, so they count as changes rather than risk suppressing
// a real one.
func biosAttributeEqual(current, val any) bool {
	if reflect.DeepEqual(current, val) {
		return true
	}

	s, ok := val.(string)
	if !ok {
		return false
	}

	s = strings.TrimSpace(s)

	switch c := current.(type) {
	case float64:
		f, err := strconv.ParseFloat(s, 64)
		return err == nil && f == c && strconv.FormatFloat(f, 'g', -1, 64) == s
	case bool:
		b, err := strconv.ParseBool(strings.ToLower(s))
		return err == nil && b == c
	}

	return false
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

		if strings.HasPrefix(key, "BootTypeOrder") || !biosAttributeEqual(current, val) {
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
	if etag := etagFromHeader(resp.Header); etag != "" {
		header["If-Match"] = etag
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
