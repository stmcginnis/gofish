//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/EndpointGroup.v1_3_4.json
// 2020.3 - #EndpointGroup.v1_3_4.EndpointGroup

package schemas

import (
	"encoding/json"
)

type GroupType string

const (
	// ClientGroupType shall indicate that the endpoint group contains client
	// (initiator) endpoints. If the associated endpoints contain the 'EntityRole'
	// property, the 'EntityRole' property shall contain the value 'Initiator' or
	// 'Both'.
	ClientGroupType GroupType = "Client"
	// ServerGroupType shall indicate that the endpoint group contains server
	// (target) endpoints. If the associated endpoints contain the 'EntityRole'
	// property, the 'EntityRole' property shall contain the value 'Target' or
	// 'Both'.
	ServerGroupType GroupType = "Server"
	// InitiatorGroupType shall indicate that the endpoint group contains initiator
	// endpoints. If the associated endpoints contain the 'EntityRole' property,
	// the 'EntityRole' property shall contain the value 'Initiator' or 'Both'.
	InitiatorGroupType GroupType = "Initiator"
	// TargetGroupType shall indicate that the endpoint group contains target
	// endpoints. If the associated endpoints contain the 'EntityRole' property,
	// the 'EntityRole' property shall contain the value 'Target' or 'Both'.
	TargetGroupType GroupType = "Target"
)

// EndpointGroup shall represent a group of endpoints that are managed as a unit
// for a Redfish implementation.
type EndpointGroup struct {
	Entity
	// AccessState shall contain the access state for all associated resources in
	// this endpoint group.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'AccessState' property in
	// the connection resource.
	AccessState AccessState
	// EndpointsCount
	EndpointsCount int `json:"Endpoints@odata.count"`
	// GroupType shall contain the endpoint group type. If this endpoint group
	// represents a SCSI target group, the value of this property shall contain
	// 'Server' or 'Target'.
	GroupType GroupType
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Preferred shall indicate if access to the resources through the endpoint
	// group is preferred over access through other endpoints. The default value
	// for this property is 'false'.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'AccessState' property in
	// the connection resource.
	Preferred bool
	// TargetEndpointGroupIdentifier shall contain a SCSI-defined identifier for
	// this group that corresponds to the TARGET PORT GROUP field in the REPORT
	// TARGET PORT GROUPS response and the TARGET PORT GROUP field in an INQUIRY
	// VPD page 85 response, type 5h identifier. See the INCITS SAM-5
	// specification. This property may not be present if the endpoint group does
	// not represent a SCSI target group.
	TargetEndpointGroupIdentifier *int `json:",omitempty"`
	// connections are the URIs for Connections.
	connections []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (e *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	type eLinks struct {
		Connections Links `json:"Connections"`
		Endpoints   Links `json:"Endpoints"`
	}
	var tmp struct {
		temp
		Links eLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = EndpointGroup(tmp.temp)

	// Extract the links to other entities for later
	e.connections = tmp.Links.Connections.ToStrings()
	e.endpoints = tmp.Links.Endpoints.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	e.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (e *EndpointGroup) Update() error {
	readWriteFields := []string{
		"AccessState",
		"Endpoints",
		"GroupType",
		"Preferred",
		"TargetEndpointGroupIdentifier",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c Client, uri string) (*EndpointGroup, error) {
	return GetObject[EndpointGroup](c, uri)
}

// ListReferencedEndpointGroups gets the collection of EndpointGroup from
// a provided reference.
func ListReferencedEndpointGroups(c Client, link string) ([]*EndpointGroup, error) {
	return GetCollectionObjects[EndpointGroup](c, link)
}

// Connections gets the Connections linked resources.
func (e *EndpointGroup) Connections() ([]*Connection, error) {
	return GetObjects[Connection](e.client, e.connections)
}

// Endpoints gets the Endpoints linked resources.
func (e *EndpointGroup) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](e.client, e.endpoints)
}
