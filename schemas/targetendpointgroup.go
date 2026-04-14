//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
)

// TargetEndpointGroup shall be managed as a unit.
type TargetEndpointGroup struct {
	Entity
	// AccessState shall share this access state.
	AccessState AccessState
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall be of the format for the reserved word *Oem*.
	OEM json.RawMessage `json:"Oem"`
	// Preferred shall indicate that access to the associated resource through the
	// endpoints in this target endpoint group is preferred over access through
	// other endpoints.
	Preferred bool
	// TargetEndpointGroupIdentifier shall contain a SCSI defined identifier for
	// this group, which corresponds to the TARGET PORT GROUP field in the REPORT
	// TARGET PORT GROUPS response and the TARGET PORT GROUP field in an INQUIRY
	// VPD page 85 response, type 5h identifier. See the INCITS SAM-5
	// specification.
	TargetEndpointGroupIdentifier *float64 `json:",omitempty"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a TargetEndpointGroup object from the raw JSON.
func (t *TargetEndpointGroup) UnmarshalJSON(b []byte) error {
	type temp TargetEndpointGroup
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TargetEndpointGroup(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *TargetEndpointGroup) Update() error {
	readWriteFields := []string{
		"AccessState",
		"Preferred",
		"TargetEndpointGroupIdentifier",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetTargetEndpointGroup will get a TargetEndpointGroup instance from the service.
func GetTargetEndpointGroup(c Client, uri string) (*TargetEndpointGroup, error) {
	return GetObject[TargetEndpointGroup](c, uri)
}

// ListReferencedTargetEndpointGroups gets the collection of TargetEndpointGroup from
// a provided reference.
func ListReferencedTargetEndpointGroups(c Client, link string) ([]*TargetEndpointGroup, error) {
	return GetCollectionObjects[TargetEndpointGroup](c, link)
}
