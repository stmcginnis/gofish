//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// BootOption represents the properties of a bootable device available in the
// system.
type BootOption struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Alias is the alias of this boot source if one exists.
	Alias BootSourceOverrideTarget
	// BootOptionEnabled is an indication of whether the boot option is
	// enabled. If true , it is enabled. If false, the boot option that the
	// boot order array on the computer system contains is skipped. In the
	// UEFI context, this property shall influence the load option active
	// flag for the boot option.
	BootOptionEnabled bool
	// BootOptionReference is the unique identifier seen in Boot.BootOrder.
	BootOptionReference string
	// DisplayName is the user-readable display name of the boot option
	// that appears in the boot order list in the user interface.
	DisplayName string
	// UefiDevicePath is the UEFI device path to access this UEFI boot
	// option.
	UefiDevicePath string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
	// etag contains the etag header when fetching the object. This is used to
	// control updates to make sure the object has not been modified my a different
	// process between fetching and updating that could cause conflicts.
	etag string
}

// UnmarshalJSON unmarshals a BootOption object from the raw JSON.
func (bootoption *BootOption) UnmarshalJSON(b []byte) error {
	type temp BootOption
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*bootoption = BootOption(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	bootoption.rawData = b

	return nil
}

// GetBootOption will get a BootOption instance from the service.
func GetBootOption(c common.Client, uri string) (*BootOption, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bootoption BootOption
	err = json.NewDecoder(resp.Body).Decode(&bootoption)
	if err != nil {
		return nil, err
	}

	if resp.Header["Etag"] != nil {
		bootoption.etag = resp.Header["Etag"][0]
	}

	bootoption.SetClient(c)
	return &bootoption, nil
}
