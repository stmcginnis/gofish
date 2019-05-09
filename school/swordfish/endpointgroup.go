// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
	"github.com/stmcginnis/gofish/school/redfish"
)

// DefaultEndpointGroupPath is the default URI for the EndpointGroup
// object.
const DefaultEndpointGroupPath = "/redfish/v1/EndpointGroup"

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
	// TransitioningAccessState shall be transitioning to a new AccesState.
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
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessState is used for associated resources through all
	// aggregated endpoints shall share this access state.
	AccessState AccessState
	// Description provides a description of this resource.
	Description string
	// endpoints shall reference an Endpoint resource.
	endpoints string
	// GroupType contains only endpoints of a given type
	// Client/Initiator or Server/Target.  If this endpoint group represents
	// a SCSI target group, the value of GroupType shall be Server.
	GroupType GroupType
	// ID uniquely identifies the resource.
	ID string `json:"Id"`
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// Name is the name of the resource or array element.
	Name string
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
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (endpointgroup *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	var t struct {
		temp
		Endpoints common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpointgroup = EndpointGroup(t.temp)

	// Extract the links to other entities for later
	endpointgroup.endpoints = string(t.Endpoints)

	return nil
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c common.Client, uri string) (*EndpointGroup, error) {
	resp, err := c.Get(uri)
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

// ListEndpointGroups gets all EndpointGroup in the system.
func ListEndpointGroups(c common.Client) ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(c, DefaultEndpointGroupPath)
}

// Endpoints gets the group's endpoints.
func (endpointgroup *EndpointGroup) Endpoints() ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(endpointgroup.Client, endpointgroup.endpoints)
}
