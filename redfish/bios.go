//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #Bios.v1_2_3.Bios

package redfish

import (
	"encoding/json"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

// biosSettings shall describe any settings of a resource, specific to Bios settings.
type biosSettings struct {
	// SettingsObject shall contain the URI of the resource that the client can
	// 'PUT' or 'PATCH' to modify the resource.
	settingsObject string
	// SupportedApplyTimes shall contain the supported apply time values a client
	// is allowed to request when configuring the settings apply time. Services
	// that do not support clients configuring the apply time can support this
	// property with a single array member in order to inform the client when the
	// settings will be applied.
	//
	// Version added: v1.1.0
	SupportedApplyTimes []common.ApplyTime
}

// UnmarshalJSON unmarshals a Settings object from the raw JSON.
func (s *biosSettings) UnmarshalJSON(b []byte) error {
	type temp biosSettings
	var t struct {
		temp
		GetSettingsObject common.Link `json:"settingsObject"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = biosSettings(t.temp)

	// Extract the links to other entities for later
	s.settingsObject = t.GetSettingsObject.String()

	return nil
}

// Bios shall represent BIOS attributes for a Redfish implementation.
type Bios struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
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
	settingsApplyTimes []common.ApplyTime
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
		ChangePassword common.ActionTarget `json:"#Bios.ChangePassword"`
		ResetBios      common.ActionTarget `json:"#Bios.ResetBios"`
	}
	type biLinks struct {
		ActiveSoftwareImage common.Link  `json:"ActiveSoftwareImage"`
		SoftwareImages      common.Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions  biActions
		Links    biLinks
		Settings biosSettings `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*bi = Bios(tmp.temp)

	// Extract the links to other entities for later
	bi.changePasswordTarget = tmp.Actions.ChangePassword.Target
	bi.resetBiosTarget = tmp.Actions.ResetBios.Target
	bi.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	bi.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	bi.settingsApplyTimes = tmp.Settings.SupportedApplyTimes

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	bi.settingsTarget = tmp.Settings.settingsObject
	if bi.settingsTarget == "" {
		bi.settingsTarget = bi.ODataID
	}

	// This is a read/write object, so we need to save the raw object data for later
	bi.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (bi *Bios) Update() error {
	readWriteFields := []string{
		"Attributes",
	}

	return bi.UpdateFromRawData(bi, bi.rawData, readWriteFields)
}

// GetBios will get a Bios instance from the service.
func GetBios(c common.Client, uri string) (*Bios, error) {
	return common.GetObject[Bios](c, uri)
}

// ListReferencedBioss gets the collection of Bios from
// a provided reference.
func ListReferencedBioss(c common.Client, link string) ([]*Bios, error) {
	return common.GetCollectionObjects[Bios](c, link)
}

// AllowedAttributeUpdateApplyTimes returns the set of allowed apply times to request when
// setting the Bios attribute values.
func (bios *Bios) AllowedAttributeUpdateApplyTimes() []common.ApplyTime {
	if len(bios.settingsApplyTimes) == 0 {
		result := []common.ApplyTime{
			common.ImmediateApplyTime,
			common.OnResetApplyTime,
			common.AtMaintenanceWindowStartApplyTime,
			common.InMaintenanceWindowOnResetApplyTime,
		}
		return result
	}
	return bios.settingsApplyTimes
}

// UpdateBiosAttributesApplyAt is used to update attribute values and set apply time together
func (bios *Bios) UpdateBiosAttributesApplyAt(attrs SettingsAttributes, applyTime common.ApplyTime) error {
	payload := make(map[string]any)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(bios.rawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := bios.GetClient().Get(bios.settingsTarget)
	defer common.DeferredCleanupHTTPResponse(resp)
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

		resp, err = bios.GetClient().PatchWithHeaders(bios.settingsTarget, data, header)
		defer common.DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBiosAttributes is used to update attribute values.
func (bios *Bios) UpdateBiosAttributes(attrs SettingsAttributes) error {
	return bios.UpdateBiosAttributesApplyAt(attrs, "")
}

// ChangePassword shall change the selected BIOS password.
// newPassword - This parameter shall contain the new BIOS password.
// oldPassword - This parameter shall contain the existing BIOS password to
// change.
// passwordName - This parameter shall contain the name of the BIOS password to
// change. For instance, 'AdminPassword' or 'UserPassword'.
func (bi *Bios) ChangePassword(newPassword string, oldPassword string, passwordName string) error {
	payload := make(map[string]any)
	payload["NewPassword"] = newPassword
	payload["OldPassword"] = oldPassword
	payload["PasswordName"] = passwordName
	return bi.Post(bi.changePasswordTarget, payload)
}

// ResetBios shall reset the BIOS attributes to their default values. To
// apply the default values, a system reset may be required. This action can
// impact other resources. This action may clear pending values in the settings
// resource.
func (bi *Bios) ResetBios() error {
	payload := make(map[string]any)
	return bi.Post(bi.resetBiosTarget, payload)
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (bi *Bios) ActiveSoftwareImage(client common.Client) (*SoftwareInventory, error) {
	if bi.activeSoftwareImage == "" {
		return nil, nil
	}
	return common.GetObject[SoftwareInventory](client, bi.activeSoftwareImage)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (bi *Bios) SoftwareImages(client common.Client) ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](client, bi.softwareImages)
}
