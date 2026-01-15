//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2017.3 - #BootOption.v1_0_6.BootOption

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// BootOption shall represent a single boot option within a system.
type BootOption struct {
	common.Entity
	// Alias shall contain the string alias of this boot source that describes the
	// type of boot.
	Alias BootSource
	// BootOptionEnabled shall indicate whether the boot option is enabled. If
	// 'true', it is enabled. If 'false', the boot option that the boot order array
	// on the computer system contains shall be skipped. In the UEFI context, this
	// property shall influence the load option active flag for the boot option.
	BootOptionEnabled bool
	// BootOptionReference shall correspond to the boot option or device. For UEFI
	// systems, this string shall match the UEFI boot option variable name, such as
	// 'Boot####'. The BootOrder array of a computer system resource contains this
	// value.
	BootOptionReference string
	// DisplayName shall contain a user-readable boot option name, as it should
	// appear in the boot order list in the user interface.
	DisplayName string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedItem shall contain an array of links to resources or objects that are
	// associated with this boot option.
	relatedItem []string
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// UefiDevicePath shall contain the UEFI Specification-defined UEFI device path
	// that identifies and locates the device for this boot option.
	UefiDevicePath string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a BootOption object from the raw JSON.
func (bo *BootOption) UnmarshalJSON(b []byte) error {
	type temp BootOption
	var tmp struct {
		temp
		RelatedItem common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*bo = BootOption(tmp.temp)

	// Extract the links to other entities for later
	bo.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	bo.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (bo *BootOption) Update() error {
	readWriteFields := []string{
		"BootOptionEnabled",
		"RelatedItem@odata.count",
	}

	return bo.UpdateFromRawData(bo, bo.rawData, readWriteFields)
}

// GetBootOption will get a BootOption instance from the service.
func GetBootOption(c common.Client, uri string) (*BootOption, error) {
	return common.GetObject[BootOption](c, uri)
}

// ListReferencedBootOptions gets the collection of BootOption from
// a provided reference.
func ListReferencedBootOptions(c common.Client, link string) ([]*BootOption, error) {
	return common.GetCollectionObjects[BootOption](c, link)
}
