//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Bios.v1_2_3.json
// 2021.1 - #Bios.v1_2_3.Bios

package schemas

import (
	"encoding/json"
	"strings"
)

// Bios shall represent BIOS attributes for a Redfish implementation.
type Bios struct {
	Entity
	// AttributeRegistry The link to the attribute registry that lists the metadata
	// describing the BIOS attribute settings in this resource.
	AttributeRegistry string
	// Attributes shall contain the list of BIOS attributes specific to the
	// manufacturer or provider. BIOS attribute settings appear as additional
	// properties in this object and can be looked up in the attribute registry by
	// their 'AttributeName'.
	Attributes SettingsAttributes
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ResetBiosToDefaultsPending shall indicate whether there is a pending request
	// to reset the BIOS attributes to default values. A successful completion of
	// the 'ResetBios' action shall set this property to 'true'. Applying the
	// default attribute values to this resource shall set this property to
	// 'false'. Services may reject modification requests to the settings resource
	// if this property contains 'true'.
	//
	// Version added: v1.2.0
	ResetBiosToDefaultsPending bool
	// changePasswordTarget is the URL to send ChangePassword requests.
	changePasswordTarget string
	// resetBiosTarget is the URL to send ResetBios requests.
	resetBiosTarget string
	// settingsTarget is the URL to send settings update requests to.
	settingsTarget string
	// settingsApplyTimes is a set of allowed settings update apply times. If none
	// are specified, then the system does not provide that information.
	settingsApplyTimes []SettingsApplyTime
	// activeSoftwareImage is the URI for ActiveSoftwareImage.
	activeSoftwareImage string
	// softwareImages are the URIs for SoftwareImages.
	softwareImages []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

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

// GetBios will get a Bios instance from the service.
func GetBios(c Client, uri string) (*Bios, error) {
	return GetObject[Bios](c, uri)
}

// ListReferencedBioss gets the collection of Bios from
// a provided reference.
func ListReferencedBioss(c Client, link string) ([]*Bios, error) {
	return GetCollectionObjects[Bios](c, link)
}

// This action shall change the selected BIOS password.
// newPassword - This parameter shall contain the new BIOS password.
// oldPassword - This parameter shall contain the existing BIOS password to
// change.
// passwordName - This parameter shall contain the name of the BIOS password to
// change. For instance, 'AdminPassword' or 'UserPassword'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (bi *Bios) ChangePassword(newPassword string, oldPassword string, passwordName string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["NewPassword"] = newPassword
	payload["OldPassword"] = oldPassword
	payload["PasswordName"] = passwordName
	resp, taskInfo, err := PostWithTask(bi.client,
		bi.changePasswordTarget, payload, bi.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the BIOS attributes to their default values. To
// apply the default values, a system reset may be required. This action can
// impact other resources. This action may clear pending values in the settings
// resource.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (bi *Bios) ResetBios() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(bi.client,
		bi.resetBiosTarget, payload, bi.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (bi *Bios) ActiveSoftwareImage() (*SoftwareInventory, error) {
	if bi.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](bi.client, bi.activeSoftwareImage)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (bi *Bios) SoftwareImages() ([]*SoftwareInventory, error) {
	return GetObjects[SoftwareInventory](bi.client, bi.softwareImages)
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
	payload := make(map[string]any)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(bi.rawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := bi.GetClient().Get(bi.settingsTarget)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]any{"Attributes": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = bi.GetClient().PatchWithHeaders(bi.settingsTarget, data, header)
		defer DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBiosAttributes is used to update attribute values.
func (b *Bios) UpdateBiosAttributes(attrs SettingsAttributes) error {
	return b.UpdateBiosAttributesApplyAt(attrs, "")
}
