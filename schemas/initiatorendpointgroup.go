//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
)

// InitiatorEndpointGroup shall be managed as a unit.
type InitiatorEndpointGroup struct {
	Entity
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall be of the format for the reserved word *Oem*.
	OEM json.RawMessage `json:"Oem"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a InitiatorEndpointGroup object from the raw JSON.
func (i *InitiatorEndpointGroup) UnmarshalJSON(b []byte) error {
	type temp InitiatorEndpointGroup
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = InitiatorEndpointGroup(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *InitiatorEndpointGroup) Update() error {
	readWriteFields := []string{
		"Identifier",
	}

	return i.UpdateFromRawData(i, i.RawData, readWriteFields)
}

// GetInitiatorEndpointGroup will get a InitiatorEndpointGroup instance from the service.
func GetInitiatorEndpointGroup(c Client, uri string) (*InitiatorEndpointGroup, error) {
	return GetObject[InitiatorEndpointGroup](c, uri)
}

// ListReferencedInitiatorEndpointGroups gets the collection of InitiatorEndpointGroup from
// a provided reference.
func ListReferencedInitiatorEndpointGroups(c Client, link string) ([]*InitiatorEndpointGroup, error) {
	return GetCollectionObjects[InitiatorEndpointGroup](c, link)
}
