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
)

// AuthenticationMethod is method used to authenticate.
type AuthenticationMethod string

const (
	// NoneAuthenticationMethod No authentication is used.
	NoneAuthenticationMethod AuthenticationMethod = "None"
	// CHAPAuthenticationMethod iSCSI Challenge Handshake Authentication
	// Protocol (CHAP) authentication is used.
	CHAPAuthenticationMethod AuthenticationMethod = "CHAP"
	// MutualCHAPAuthenticationMethod iSCSI Mutual Challenge Handshake
	// Authentication Protocol (CHAP) authentication is used.
	MutualCHAPAuthenticationMethod AuthenticationMethod = "MutualCHAP"
	// DHCHAPAuthenticationMethod Diffie-Hellman Challenge Handshake
	// Authentication Protocol (DHCHAP) is an authentication protocol used in
	// Fibre Channel. DHCHAP implies that only properties 'TargetCHAPUser'
	// and 'TargetPassword' need to be present.
	DHCHAPAuthenticationMethod AuthenticationMethod = "DHCHAP"
)

// CHAPInformation is used for CHAP auth.
type CHAPInformation struct {
	// InitiatorCHAPPassword shall be the
	// shared secret for CHAP authentication.
	InitiatorCHAPPassword string
	// InitiatorCHAPUser is If present, this property is the initiator CHAP
	// username for authentication. For example, with an iSCSI scenario, use
	// the initiator iQN.
	InitiatorCHAPUser string
	// TargetCHAPUser shall be the CHAP
	// Username for 2-way CHAP authentication. For example, with an iSCSI
	// scenario, use the target iQN. In a FC with DHCHAP, this value will be
	// a FC WWN.
	TargetCHAPUser string
	// TargetPassword shall be the CHAP Secret
	// for 2-way CHAP authentication.
	TargetPassword string
}

// StorageGroup is a set of related storage entities (volumes, file systems...)
// The collection should be useful for managing the storage of a set of related
// client applications.
type StorageGroup struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessState shall describe the access
	// characteristics of this storage group. All associated logical units
	// through all aggregated ports shall share this access state.
	AccessState AccessState
	// AuthenticationMethod is The value of this property must be what kind
	// of authentication that the endpoints in this StorageGroup understands.
	AuthenticationMethod AuthenticationMethod
	// ChapInfo is used by this specific endpoint. For example, if this
	// endpoint represents an initiator, and AuthenticationMethod is CHAP or
	// MutualCHAP, the Credentials fields CHAPUsername and CHAPSecret must be
	// used. If this endpoint represents a target endpoint and
	// AuthenticationMethod is MutualCHAP, then MutualCHAPUsername and
	// MutualCHAPSecret must be used.
	ChapInfo []CHAPInformation
	// ClientEndpointGroups is used to make requests to the storage exposed
	// by this StorageGroup. If null, the implementation may allow access to
	// the storage via any client-side endpoint. If empty, the
	// implementation shall not allow access to the storage via any client-
	// side endpoint.
	ClientEndpointGroups []EndpointGroup
	// ClientEndpointGroups@odata.count is
	ClientEndpointGroupsCount int `json:"ClientEndpointGroups@odata.count"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// mappedVolumes is an array of mapped volumes managed by this storage
	// group.
	mappedVolumes []string
	// MembersAreConsistent shall be set to true if all members are in a
	// consistent state. The default value for this property is false.
	MembersAreConsistent bool
	// ReplicaInfo shall describe the replication relationship between this
	// storage group and a corresponding source storage group.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that
	// are sourced by this replica.
	replicaTargets []string
	// ReplicaTargetsCount is number of replica targets.
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// serverEndpointGroups is used to make requests to the storage exposed
	// by this storage group. If null, the implementation may allow access
	// to the storage via any server-side endpoint. If empty, the
	// implementation shall not allow access to the storage via any server-
	// side endpoint.
	serverEndpointGroups []string
	// ServerEndpointGroupsCount is the number of server endpoints.
	ServerEndpointGroupsCount int `json:"ServerEndpointGroups@odata.count"`
	// Status is the status of this group.
	Status common.Status
	// VolumesCount is the number of volumes.
	VolumesCount int `json:"Volumes@odata.count"`
	// VolumesAreExposed shall be set to true if storage volumes are exposed to
	// the paths defined by the client and server endpoints. The default value
	// for this property is false.
	VolumesAreExposed bool
	// ChildStorageGroups is an array of references to StorageGroups are
	// incorporated into this StorageGroup
	childStorageGroups []string
	// ChildStorageGroupsCount is the number of child storage groups.
	ChildStorageGroupsCount int
	// ClassOfService is the ClassOfService that all storage in this
	// StorageGroup conforms to.
	classOfService string
	// ParentStorageGroups is an array of references to StorageGroups that
	// incorporate this StorageGroup
	parentStorageGroups []string
	// ParentStorageGroupsCount is the number of parent storage groups.
	ParentStorageGroupsCount int
}

// UnmarshalJSON unmarshals a StorageGroup object from the raw JSON.
func (storagegroup *StorageGroup) UnmarshalJSON(b []byte) error {
	type temp StorageGroup
	type links struct {
		ChildStorageGroups       common.Links
		ChildStorageGroupsCount  int `json:"ChildStorageGroups@odata.count"`
		ClassOfService           common.Link
		ParentStorageGroups      common.Links
		ParentStorageGroupsCount int `json:"ParentStorageGroups@odata.count"`
	}
	var t struct {
		temp
		Links                links
		MappedVolumes        common.Links
		ServerEndpointGroups common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagegroup = StorageGroup(t.temp)

	// Extract the links to other entities for later
	storagegroup.mappedVolumes = t.MappedVolumes.ToStrings()
	storagegroup.childStorageGroups = t.Links.ChildStorageGroups.ToStrings()
	storagegroup.ChildStorageGroupsCount = t.Links.ChildStorageGroupsCount
	storagegroup.classOfService = string(t.Links.ClassOfService)
	storagegroup.parentStorageGroups = t.Links.ParentStorageGroups.ToStrings()
	storagegroup.ParentStorageGroupsCount = t.Links.ParentStorageGroupsCount

	return nil
}

// GetStorageGroup will get a StorageGroup instance from the service.
func GetStorageGroup(c common.Client, uri string) (*StorageGroup, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storagegroup StorageGroup
	err = json.NewDecoder(resp.Body).Decode(&storagegroup)
	if err != nil {
		return nil, err
	}

	storagegroup.SetClient(c)
	return &storagegroup, nil
}

// ListReferencedStorageGroups gets the collection of StorageGroup from
// a provided reference.
func ListReferencedStorageGroups(c common.Client, link string) ([]*StorageGroup, error) {
	var result []*StorageGroup
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storagegroupLink := range links.ItemLinks {
		storagegroup, err := GetStorageGroup(c, storagegroupLink)
		if err != nil {
			return result, err
		}
		result = append(result, storagegroup)
	}

	return result, nil
}

// MappedVolumes gets mapped volumes managed by this storage group.
func (storagegroup *StorageGroup) MappedVolumes() ([]*Volume, error) {
	var result []*Volume
	for _, volLink := range storagegroup.mappedVolumes {
		vol, err := GetVolume(storagegroup.Client, volLink)
		if err != nil {
			return result, err
		}
		result = append(result, vol)
	}

	return result, nil
}

// ChildStorageGroups gets child groups of this group.
func (storagegroup *StorageGroup) ChildStorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup
	for _, sgLink := range storagegroup.childStorageGroups {
		sg, err := GetStorageGroup(storagegroup.Client, sgLink)
		if err != nil {
			return result, err
		}
		result = append(result, sg)
	}

	return result, nil
}

// ParentStorageGroups gets parent groups of this group.
func (storagegroup *StorageGroup) ParentStorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup
	for _, sgLink := range storagegroup.parentStorageGroups {
		sg, err := GetStorageGroup(storagegroup.Client, sgLink)
		if err != nil {
			return result, err
		}
		result = append(result, sg)
	}

	return result, nil
}

// ClassOfService gets the ClassOfService that all storage in this StorageGroup
// conforms to.
func (storagegroup *StorageGroup) ClassOfService() (*ClassOfService, error) {
	if storagegroup.classOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storagegroup.Client, storagegroup.classOfService)
}
