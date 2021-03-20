//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// AccessState is used for associated resources through all
// aggregated endpoints shall share this access state.
type AccessState string

const (
	// OptimizedAccessState shall be in an Active/Optimized state.
	OptimizedAccessState AccessState = "Optimized"
	// NonOptimizedAccessState shall be in an Active/NonOptimized state.
	NonOptimizedAccessState AccessState = "NonOptimized"
	// StandbyAccessState shall be in a Standby state.
	StandbyAccessState AccessState = "Standby"
	// UnavailableAccessState shall be in an unavailable state.
	UnavailableAccessState AccessState = "Unavailable"
	// TransitioningAccessState shall be transitioning to a new AccessState.
	TransitioningAccessState AccessState = "Transitioning"
)

// GroupType is the type of endpoint grouping.
type GroupType string

const (
	// ClientGroupType The group contains the client (initiator) endpoints.
	ClientGroupType GroupType = "Client"
	// ServerGroupType The group contains the server (target) endpoints.
	ServerGroupType GroupType = "Server"
)

// EndpointGroup is a group of endpoints that shall be managed as a unit.
type EndpointGroup struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessState is used for associated resources through all
	// aggregated endpoints shall share this access state.
	AccessState AccessState
	// Description provides a description of this resource.
	Description string
	// endpoints shall reference an Endpoint resource.
	endpoints string
	// EndpointsCount is the number of Endpoints
	EndpointsCount int
	// GroupType contains only endpoints of a given type
	// Client/Initiator or Server/Target.  If this endpoint group represents
	// a SCSI target group, the value of GroupType shall be Server.
	GroupType GroupType
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// Preferred with a value of True shall indicate that
	// access to the associated resource through the endpoints in this
	// endpoint group is preferred over access through other endpoints. The
	// default value for this property is false.
	Preferred bool
	// TargetEndpointGroupIdentifier represents a
	// SCSI target group, the value of this property shall contain a SCSI
	// defined identifier for this group, which corresponds to the TARGET
	// PORT GROUP field in the REPORT TARGET PORT GROUPS response and the
	// TARGET PORT GROUP field in an INQUIRY VPD page 85 response, type 5h
	// identifier. See the INCITS SAM-5 specification.
	TargetEndpointGroupIdentifier int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (endpointgroup *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	var t struct {
		temp
		Endpoints      common.Link
		EndpointsCount int `json:"Endpoints@odata.count"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpointgroup = EndpointGroup(t.temp)

	// Extract the links to other entities for later
	endpointgroup.endpoints = string(t.Endpoints)
	endpointgroup.EndpointsCount = t.EndpointsCount

	// This is a read/write object, so we need to save the raw object data for later
	endpointgroup.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (endpointgroup *EndpointGroup) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EndpointGroup)
	err := original.UnmarshalJSON(endpointgroup.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccessState",
		"Endpoints",
		"GroupType",
		"Preferred",
		"TargetEndpointGroupIdentifier",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(endpointgroup).Elem()

	return endpointgroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c common.Client, uri string) (*EndpointGroup, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endpointgroup EndpointGroup
	err = json.NewDecoder(resp.Body).Decode(&endpointgroup)
	if err != nil {
		return nil, err
	}

	endpointgroup.SetClient(c)
	return &endpointgroup, nil
}

// ListReferencedEndpointGroups gets the collection of EndpointGroup from
// a provided reference.
func ListReferencedEndpointGroups(c common.Client, link string) ([]*EndpointGroup, error) {
	var result []*EndpointGroup
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, endpointgroupLink := range links.ItemLinks {
		endpointgroup, err := GetEndpointGroup(c, endpointgroupLink)
		if err != nil {
			return result, err
		}
		result = append(result, endpointgroup)
	}

	return result, nil
}

// Endpoints gets the group's endpoints.
func (endpointgroup *EndpointGroup) Endpoints() ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(endpointgroup.Client, endpointgroup.endpoints)
}
