//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type GroupType string

const (
	// ClientGroupType shall indicate that the endpoint group contains client (initiator) endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Initiator' or
	// 'Both'.
	ClientGroupType GroupType = "Client"
	// ServerGroupType shall indicate that the endpoint group contains server (target) endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Target' or 'Both'.
	ServerGroupType GroupType = "Server"
	// InitiatorGroupType shall indicate that the endpoint group contains initiator endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Initiator' or
	// 'Both'.
	InitiatorGroupType GroupType = "Initiator"
	// TargetGroupType shall indicate that the endpoint group contains target endpoints. If the associated endpoints
	// contain the EntityRole property, the EntityRole property shall contain the value 'Target' or 'Both'.
	TargetGroupType GroupType = "Target"
)

// EndpointGroup shall represent a group of endpoints that are managed as a unit for a Redfish implementation.
type EndpointGroup struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// GroupType shall contain the endpoint group type. If this endpoint group represents a SCSI target group, the
	// value of this property shall contain 'Server' or 'Target'.
	GroupType GroupType
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// TargetEndpointGroupIdentifier shall contain a SCSI-defined identifier for this group that corresponds to the
	// TARGET PORT GROUP field in the REPORT TARGET PORT GROUPS response and the TARGET PORT GROUP field in an INQUIRY
	// VPD page 85 response, type 5h identifier. See the INCITS SAM-5 specification. This property may not be present
	// if the endpoint group does not represent a SCSI target group.
	TargetEndpointGroupIdentifier int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	endpoints []string
	// EndpointsCount is the number of Endpoints in the group.
	EndpointsCount int `json:"Endpoints@odata.count"`
	connections    []string
	// ConnectionsCount is the number of Connections to this group.
	ConnectionsCount int
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (endpointgroup *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	type Links struct {
		// Connections shall contain an array of links to resources of type Connection that represent the connections to
		// which this endpoint group belongs.
		Connections      common.Links
		ConnectionsCount int `json:"Connections@odata.count"`
		// Endpoints shall contain an array of links to resources of type Endpoint that represent the endpoints that are in
		// this endpoint group.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
	}
	var t struct {
		temp
		Endpoints common.Links
		Links     Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpointgroup = EndpointGroup(t.temp)

	// Extract the links to other entities for later
	endpointgroup.connections = t.Links.Connections.ToStrings()
	endpointgroup.ConnectionsCount = t.Links.ConnectionsCount

	// Handle the move of endpoint links in 1.3.0
	if len(endpointgroup.endpoints) == 0 {
		endpointgroup.endpoints = t.Links.Endpoints.ToStrings()
		endpointgroup.EndpointsCount = t.Links.EndpointsCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	endpointgroup.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (endpointgroup *EndpointGroup) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EndpointGroup)
	original.UnmarshalJSON(endpointgroup.rawData)

	readWriteFields := []string{
		"GroupType",
		"TargetEndpointGroupIdentifier",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(endpointgroup).Elem()

	return endpointgroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c common.Client, uri string) (*EndpointGroup, error) {
	return common.GetObject[EndpointGroup](c, uri)
}

// ListReferencedEndpointGroups gets the collection of EndpointGroup from
// a provided reference.
func ListReferencedEndpointGroups(c common.Client, link string) ([]*EndpointGroup, error) {
	return common.GetCollectionObjects[EndpointGroup](c, link)
}

// Endpoints get the endpoints associated with this endpoint group.
func (endpointgroup *EndpointGroup) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](endpointgroup.GetClient(), endpointgroup.endpoints)
}

// Connections get the connections associated with this endpoint group.
func (endpointgroup *EndpointGroup) Connections() ([]*Connection, error) {
	return common.GetObjects[Connection](endpointgroup.GetClient(), endpointgroup.connections)
}
