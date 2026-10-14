//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// If @Redfish.Settings is present, PATCH goes to the System resource first.
// Falls back to the Settings URI with ETag concurrency control only if the
// direct PATCH fails (e.g., Dell iDRAC10 returns read-only).
func (c *ComputerSystem) SetBoot(b *Boot) error {
	t := struct {
		Boot *Boot
	}{Boot: b}

	// Try the main System resource first.
	err := c.Patch(c.ODataID, t)
	if err == nil {
		return nil
	}

	// Fall back to the Settings resource with ETag concurrency control.
	if c.settingsTarget != c.ODataID {
		resp, getErr := c.GetClient().Get(c.settingsTarget)
		defer DeferredCleanupHTTPResponse(resp)
		if getErr != nil {
			return fmt.Errorf("system resource: %w, settings resource: %w", err, getErr)
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = sanitizeETag(resp.Header["Etag"][0])
		}

		resp, patchErr := c.GetClient().PatchWithHeaders(c.settingsTarget, t, header)
		defer DeferredCleanupHTTPResponse(resp)
		if patchErr != nil {
			return fmt.Errorf("system resource: %w, settings resource: %w", err, patchErr)
		}

		return nil
	}

	return err
}

// This action shall reset the system represented by the resource. For systems
// that implement ACPI Power Button functionality, the 'PushPowerButton' value
// shall perform or emulate an ACPI Power Button Push, and the 'ForceOff' value
// shall perform an ACPI Power Button Override, commonly known as a four-second
// hold of the power button.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	// Make sure the requested reset type is supported by the system
	valid := false
	if len(c.supportedResetTypes) > 0 {
		for _, allowed := range c.supportedResetTypes {
			if resetType == allowed {
				valid = true
				break
			}
		}
	} else {
		// No allowed values supplied, assume we are OK
		valid = true
	}

	if !valid {
		return nil, fmt.Errorf("reset type '%s' is not supported by this service",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetTarget, t, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// GetSupportedResetTypes returns any reset types that the ComputerSystem declares as supported
// via either ActionInfo or AllowableValues.
func (c *ComputerSystem) GetSupportedResetTypes() ([]ResetType, error) {
	if len(c.supportedResetTypes) > 0 {
		return c.supportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if c.resetActionInfoTarget != "" {
		resetActionInfo, err := c.ResetActionInfo()
		if err != nil {
			return c.supportedResetTypes, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return c.supportedResetTypes, err
		}

		for _, val := range vals {
			c.supportedResetTypes = append(c.supportedResetTypes, ResetType(val))
		}
	}

	return c.supportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the ComputerSystem reset action if supported
func (c *ComputerSystem) ResetActionInfo() (*ActionInfo, error) {
	if c.resetActionInfoTarget == "" {
		return nil, errors.New("ComputerSystem Reset ActionInfo not supported")
	}

	return GetObject[ActionInfo](c.GetClient(), c.resetActionInfoTarget)
}

// UpdateBootAttributesApplyAt is used to update attribute values and set apply time together
func (c *ComputerSystem) UpdateBootAttributesApplyAt(attrs SettingsAttributes, applyTime SettingsApplyTime) error {
	payload := make(map[string]any)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(c.RawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := c.GetClient().Get(c.settingsTarget)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]any{"Boot": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = c.GetClient().PatchWithHeaders(c.settingsTarget, data, header)
		defer DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBootAttributes is used to update attribute values.
func (c *ComputerSystem) UpdateBootAttributes(attrs SettingsAttributes) error {
	return c.UpdateBootAttributesApplyAt(attrs, "")
}

// BootOptions gets all BootOption items for this system.
func (c *ComputerSystem) BootOptions() ([]*BootOption, error) {
	return GetCollectionObjects[BootOption](
		c.GetClient(),
		c.Boot.bootOptions)
}

// StorageServices gets the StorageServices linked resource.
func (h *HostedServices) StorageServices(client Client) (*StorageService, error) {
	if h.storageServices == "" {
		return nil, nil
	}
	return GetObject[StorageService](client, h.storageServices)
}

// UnmarshalJSON unmarshals an IPMIHostInterface object from the raw JSON.
func (i *IPMIHostInterface) UnmarshalJSON(b []byte) error {
	type temp IPMIHostInterface
	var t temp
	err := json.Unmarshal(b, &t)
	if err == nil {
		*i = IPMIHostInterface(t)
		return nil
	}

	// Some BMCs emit a bare bool rather than the object the schema defines.
	var enabled bool
	if json.Unmarshal(b, &enabled) != nil {
		return err
	}
	i.ServiceEnabled = enabled
	return nil
}

// postUnmarshal records the allowed reset types and the settings resource from
// @Redfish.Settings.
func (c *ComputerSystem) postUnmarshal(b []byte) error {
	var t struct {
		Actions struct {
			Reset struct {
				ActionTarget
				AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			} `json:"#ComputerSystem.Reset"`
		}
		Settings Settings `json:"@Redfish.Settings"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	c.resetActionInfoTarget = t.Actions.Reset.ActionInfoTarget
	c.supportedResetTypes = t.Actions.Reset.AllowedResetTypes
	c.settingsTarget = t.Settings.SettingsObject
	if c.settingsTarget == "" {
		c.settingsTarget = c.ODataID
	}
	c.SettingsApplyTimes = t.Settings.SupportedApplyTimes

	return nil
}

// postUnmarshal records the StorageServices link.
func (h *HostedServices) postUnmarshal(b []byte) error {
	var t struct {
		StorageServices Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	h.storageServices = t.StorageServices.String()

	return nil
}
