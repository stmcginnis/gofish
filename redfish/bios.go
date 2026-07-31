//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreweave/gofish/common"
)

// Bios is used to represent BIOS attributes.
type Bios struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AttributeRegistry is the Resource ID of the Attribute Registry that has
	// the system-specific information about a BIOS resource.
	AttributeRegistry string
	// This property shall contain the list of BIOS attributes and their values
	// as determined by the manufacturer or provider. This object shall
	// describe BIOS attribute settings as additional properties. If the object
	// specifies a BIOS Attribute Registry, attributes shall be looked up in
	// that Attribute Registry by their attribute name. Attributes in this
	// Attribute Registry with the AttributeType of Enumeration shall use valid
	// ValueName values in this object, as listed in that Attribute Registry.
	Attributes SettingsAttributes
	// Attributes are additional properties in this object, and can be looked up
	// in the Attribute Registry by their AttributeName.
	// Attributes string
	// Description provides a description of this resource.
	Description string
	// changePasswordTarget is the URL to send ChangePassword requests.
	changePasswordTarget string
	// resetBiosTarget is the URL to send ResetBios requests.
	resetBiosTarget string
	// settingsTarget is the URL to send settings update requests to.
	settingsTarget string
	// settingsApplyTimes is a set of allowed settings update apply times. If none
	// are specified, then the system does not provide that information.
	settingsApplyTimes []common.ApplyTime
	// activeSoftwareImage is the @odata.id of SoftwareInventory responsible
	// for the active BIOS firmware image (see Links.ActiveSoftwareImage.@odata.id).
	activeSoftwareImage string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an Bios object from the raw JSON.
func (bios *Bios) UnmarshalJSON(b []byte) error {
	type temp Bios
	type Actions struct {
		ChangePassword common.ActionTarget `json:"#Bios.ChangePassword"`
		ResetBios      common.ActionTarget `json:"#Bios.ResetBios"`
	}
	type Links struct {
		ActiveSoftwareImage struct {
			ODataID string `json:"@odata.id"`
		}
	}
	var t struct {
		temp
		Actions  Actions
		Links    Links
		Settings common.Settings `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*bios = Bios(t.temp)

	// Extract the links to other entities for later
	bios.changePasswordTarget = t.Actions.ChangePassword.Target
	bios.resetBiosTarget = t.Actions.ResetBios.Target
	bios.settingsApplyTimes = t.Settings.SupportedApplyTimes
	bios.activeSoftwareImage = t.Links.ActiveSoftwareImage.ODataID

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	bios.settingsTarget = t.Settings.SettingsObject.String()
	if bios.settingsTarget == "" {
		bios.settingsTarget = bios.ODataID
	}

	// This is a read/write object, so we need to save the raw object data for later
	bios.rawData = b

	return nil
}

// GetBios will get a Bios instance from the service.
func GetBios(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Bios, error) {
	return GetBiosWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetBiosWithContext will get a Bios instance from the service.
func GetBiosWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Bios, error) {
	return common.GetObjectWithContext[Bios](ctx, c, uri, queryOpts...)
}

// ListReferencedBioss gets the collection of Bios from a provided reference.
func ListReferencedBioss(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Bios, error) {
	return ListReferencedBiossWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedBiossWithContext gets the collection of Bios from a provided reference.
func ListReferencedBiossWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Bios, error) {
	return common.GetCollectionObjectsWithContext[Bios](ctx, c, link, queryOpts...)
}

// ChangePasswordTarget returns the action target URL for #Bios.ChangePassword,
// or an empty string if the BIOS resource does not advertise that action.
func (bios *Bios) ChangePasswordTarget() string {
	return bios.changePasswordTarget
}

// ChangePassword shall change the selected BIOS password.
func (bios *Bios) ChangePassword(passwordName, oldPassword, newPassword string) error {
	return bios.ChangePasswordWithContext(common.ContextOf(bios.GetClient()), passwordName, oldPassword, newPassword)
}

// ChangePasswordWithContext shall change the selected BIOS password.
func (bios *Bios) ChangePasswordWithContext(ctx context.Context, passwordName, oldPassword, newPassword string) error {
	if passwordName == "" {
		return fmt.Errorf("password name must be supplied")
	}

	t := struct {
		PasswordName string
		OldPassword  string
		NewPassword  string
	}{
		PasswordName: passwordName,
		OldPassword:  oldPassword,
		NewPassword:  newPassword,
	}
	return bios.PostWithContext(ctx, bios.changePasswordTarget, t)
}

// ResetBios shall perform a reset of the BIOS attributes to their default values.
// A system reset may be required for the default values to be applied. This
// action may impact other resources.
func (bios *Bios) ResetBios() error {
	return bios.ResetBiosWithContext(common.ContextOf(bios.GetClient()))
}

// ResetBiosWithContext shall perform a reset of the BIOS attributes to their default values.
// A system reset may be required for the default values to be applied. This
// action may impact other resources.
func (bios *Bios) ResetBiosWithContext(ctx context.Context) error {
	payload := make(map[string]interface{})
	return bios.PostWithContext(ctx, bios.resetBiosTarget, payload)
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
	return bios.UpdateBiosAttributesApplyAtWithContext(common.ContextOf(bios.GetClient()), attrs, applyTime)
}

// UpdateBiosAttributesApplyAtWithContext is used to update attribute values and set apply time together
func (bios *Bios) UpdateBiosAttributesApplyAtWithContext(ctx context.Context, attrs SettingsAttributes, applyTime common.ApplyTime) error {
	payload := make(map[string]interface{})

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

	resp, err := bios.GetClient().GetWithContext(ctx, bios.settingsTarget)
	defer common.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]interface{}{"Attributes": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = bios.GetClient().PatchWithHeadersWithContext(ctx, bios.settingsTarget, data, header)
		defer common.DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBiosAttributes is used to update attribute values.
func (bios *Bios) UpdateBiosAttributes(attrs SettingsAttributes) error {
	return bios.UpdateBiosAttributesWithContext(common.ContextOf(bios.GetClient()), attrs)
}

// UpdateBiosAttributesWithContext is used to update attribute values.
func (bios *Bios) UpdateBiosAttributesWithContext(ctx context.Context, attrs SettingsAttributes) error {
	return bios.UpdateBiosAttributesApplyAtWithContext(ctx, attrs, "")
}

// GetActiveSoftwareImage gets the SoftwareInventory which represents the
// active BIOS firmware image.
func (bios *Bios) GetActiveSoftwareImage(queryOpts ...common.QueryGroupOption) (*SoftwareInventory, error) {
	return bios.GetActiveSoftwareImageWithContext(common.ContextOf(bios.GetClient()), queryOpts...)
}

// GetActiveSoftwareImageWithContext gets the SoftwareInventory which represents the
// active BIOS firmware image.
func (bios *Bios) GetActiveSoftwareImageWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*SoftwareInventory, error) {
	if bios.activeSoftwareImage == "" {
		return nil, nil
	}

	return GetSoftwareInventoryWithContext(ctx, bios.GetClient(), bios.activeSoftwareImage, queryOpts...)
}
