//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

// BiosAttributes handles the Bios attribute values that may be any of several
// types and adds some basic helper methods to make accessing values easier.
type BiosAttributes map[string]interface{}

// String gets the string representation of the attribute value.
func (ba BiosAttributes) String(name string) string {
	if val, ok := ba[name]; ok {
		return fmt.Sprintf("%v", val)
	}

	return ""
}

// Float64 gets the value as a float64 or 0 if that is not possible.
func (ba BiosAttributes) Float64(name string) float64 {
	if val, ok := ba[name]; ok {
		return val.(float64)
	}

	return 0
}

// Int gets the value as an integer or 0 if that is not possible.
func (ba BiosAttributes) Int(name string) int {
	// Integer values may be interpeted as float64, so get it as that first,
	// then coerce down to int.
	floatVal := int(ba.Float64(name))
	return (floatVal)
}

// Bool gets the value as a boolean or returns false.
func (ba BiosAttributes) Bool(name string) bool {
	maybeBool := ba.String(name)
	maybeBool = strings.ToLower(maybeBool)
	return (maybeBool == "true" ||
		maybeBool == "1" ||
		maybeBool == "enabled")
}

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
	Attributes BiosAttributes
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an Bios object from the raw JSON.
func (bios *Bios) UnmarshalJSON(b []byte) error {
	type temp Bios
	type Actions struct {
		ChangePassword struct {
			Target string
		} `json:"#Bios.ChangePassword"`
		ResetBios struct {
			Target string
		} `json:"#Bios.ResetBios"`
	}
	var t struct {
		temp
		Actions  Actions
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

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	bios.settingsTarget = string(t.Settings.SettingsObject)
	if bios.settingsTarget == "" {
		bios.settingsTarget = bios.ODataID
	}

	// This is a read/write object, so we need to save the raw object data for later
	bios.rawData = b

	return nil
}

// GetBios will get a Bios instance from the service.
func GetBios(c common.Client, uri string) (*Bios, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bios Bios
	err = json.NewDecoder(resp.Body).Decode(&bios)
	if err != nil {
		return nil, err
	}

	bios.SetClient(c)
	return &bios, nil
}

// ListReferencedBioss gets the collection of Bios from a provided reference.
func ListReferencedBioss(c common.Client, link string) ([]*Bios, error) {
	var result []*Bios
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, biosLink := range links.ItemLinks {
		bios, err := GetBios(c, biosLink)
		if err != nil {
			return result, err
		}
		result = append(result, bios)
	}

	return result, nil
}

// ChangePassword shall change the selected BIOS password.
func (bios *Bios) ChangePassword(passwordName, oldPassword, newPassword string) error {
	if passwordName == "" {
		return fmt.Errorf("password name must be supplied")
	}
	if oldPassword == "" {
		return fmt.Errorf("existing password must be supplied")
	}
	if newPassword == "" {
		return fmt.Errorf("new password must be supplied")
	}

	type temp struct {
		PasswordName string
		OldPassword  string
		NewPassword  string
	}
	t := temp{
		PasswordName: passwordName,
		OldPassword:  oldPassword,
		NewPassword:  newPassword,
	}

	_, err := bios.Client.Post(bios.changePasswordTarget, t)
	return err
}

// ResetBios shall perform a reset of the BIOS attributes to their default values.
// A system reset may be required for the default values to be applied. This
// action may impact other resources.
func (bios *Bios) ResetBios() error {
	_, err := bios.Client.Post(bios.resetBiosTarget, nil)
	return err
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

// UpdateBiosAttributes is used to update attribute values.
func (bios *Bios) UpdateBiosAttributes(attrs BiosAttributes) error {
	payload := make(map[string]interface{})

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(bios.rawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]interface{}{"Attributes": payload}
		_, err := bios.Client.Patch(bios.settingsTarget, data)
		if err != nil {
			return err
		}
	}

	return nil
}
