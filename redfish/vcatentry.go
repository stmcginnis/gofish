//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// VCATEntry shall represent an entry of Virtual Channel Action Table in a Redfish implementation.
type VCATEntry struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawEntryHex shall contain the hexadecimal value of the Virtual Channel Action Table entries. The length of the
	// hexadecimal value depends on the number of Virtual Channel Action entries supported by the component.
	RawEntryHex string
	// VCEntries shall contain an array of entries of the Virtual Channel Action Table. The length of the array depends
	// on the number of Virtual Channel Action entries supported by the component.
	VCEntries []VCATableEntry
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VCATEntry object from the raw JSON.
func (vcatentry *VCATEntry) UnmarshalJSON(b []byte) error {
	type temp VCATEntry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*vcatentry = VCATEntry(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	vcatentry.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (vcatentry *VCATEntry) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(VCATEntry)
	original.UnmarshalJSON(vcatentry.rawData)

	readWriteFields := []string{
		"RawEntryHex",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(vcatentry).Elem()

	return vcatentry.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVCATEntry will get a VCATEntry instance from the service.
func GetVCATEntry(c common.Client, uri string) (*VCATEntry, error) {
	return common.GetObject[VCATEntry](c, uri)
}

// ListReferencedVCATEntries gets the collection of VCATEntry from
// a provided reference.
func ListReferencedVCATEntries(c common.Client, link string) ([]*VCATEntry, error) {
	return common.GetCollectionObjects[VCATEntry](c, link)
}

// VCATableEntry shall contain a Virtual Channel entry definition that describes a specific Virtual Channel.
type VCATableEntry struct {
	// Threshold shall contain the Gen-Z Core Specification-defined 'TH' 7-bit threshold.
	Threshold string
	// VCMask shall contain a 32-bit value where the bits correspond to a supported Virtual Channel.
	VCMask string
}
