//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
	// CHAPPassword shall be the password when CHAP authentication is specified.
	CHAPPassword string
	// CHAPUser shall be the username when CHAP authentication is specified.
	CHAPUser string
	// InitiatorCHAPPassword shall be the shared secret for Mutual (2-way) CHAP
	// authentication.
	InitiatorCHAPPassword string
	// InitiatorCHAPUser If present, this property is the initiator CHAP username
	// for Mutual (2-way) authentication. For example, with an iSCSI scenario,
	// use the initiator iQN.
	InitiatorCHAPUser string
	// TargetCHAPPassword shall be the CHAP Secret for 2-way CHAP authentication.
	TargetCHAPPassword string
	// TargetCHAPUser shall be the Target CHAP Username for Mutual (2-way) CHAP
	// authentication. For example, with an iSCSI scenario, use the target iQN.
	TargetCHAPUser string
}

// DHCHAPInformation User name and password values for target and initiator
// endpoints when CHAP authentication is used.
type DHCHAPInformation struct {
	// LocalDHCHAPAuthSecret shall be the local DHCHAP auth secret for DHCHAP authentication.
	LocalDHCHAPAuthSecret string
	// PeerDHCHAPAuthSecret shall be the peer DHCHAP auth secret for DHCHAP authentication.
	PeerDHCHAPAuthSecret string
}

// StorageGroup is a set of related storage entities (volumes, file systems...)
// The collection should be useful for managing the storage of a set of related
// client applications.
type StorageGroup struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
	MappedVolumes []MappedVolume
	// MembersAreConsistent shall be set to true if all members are in a
	// consistent state. The default value for this property is false.
	// Deprecated in favor of using the ConsistencyGroup for Consistency set management.
	MembersAreConsistent bool
	// ReplicaInfo shall describe the replication relationship between this
	// storage group and a corresponding source storage group.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that
	// are sourced by this replica.
	// replicaTargets []string
	// ReplicaTargetsCount is number of replica targets.
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// serverEndpointGroups is used to make requests to the storage exposed
	// by this storage group. If null, the implementation may allow access
	// to the storage via any server-side endpoint. If empty, the
	// implementation shall not allow access to the storage via any server-
	// side endpoint.
	// serverEndpointGroups []string
	// ServerEndpointGroupsCount is the number of server endpoints.
	ServerEndpointGroupsCount int `json:"ServerEndpointGroups@odata.count"`
	// Status is the status of this group.
	Status common.Status
	// VolumesCount is the number of volumes.
	// These references are replaced by the MappedVolumes array in StorageGroup.
	VolumesCount int `json:"Volumes@odata.count"`
	// VolumesAreExposed shall be set to true if storage volumes are exposed to
	// the paths defined by the client and server endpoints. The default value
	// for this property is false.
	VolumesAreExposed bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

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

	// exposeVolumesTarget is the URL to for the ExposeVolumes action.
	exposeVolumesTarget string
	// hideVolumesTarget is the URL to for the HideVolumes action.
	hideVolumesTarget string
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
	type actions struct {
		ExposeVolumes common.ActionTarget `json:"#StorageGroup.ExposeVolumes"`
		HideVolumes   common.ActionTarget `json:"#StorageGroup.HideVolumes"`
	}
	var t struct {
		temp
		Links                links
		ServerEndpointGroups common.Links
		Actions              actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*storagegroup = StorageGroup(t.temp)
	storagegroup.childStorageGroups = t.Links.ChildStorageGroups.ToStrings()
	storagegroup.ChildStorageGroupsCount = t.Links.ChildStorageGroupsCount
	storagegroup.classOfService = t.Links.ClassOfService.String()
	storagegroup.parentStorageGroups = t.Links.ParentStorageGroups.ToStrings()
	storagegroup.ParentStorageGroupsCount = t.Links.ParentStorageGroupsCount

	storagegroup.exposeVolumesTarget = t.Actions.ExposeVolumes.Target
	storagegroup.hideVolumesTarget = t.Actions.HideVolumes.Target

	// This is a read/write object, so we need to save the raw object data for later
	storagegroup.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storagegroup *StorageGroup) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StorageGroup)
	err := original.UnmarshalJSON(storagegroup.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccessState",
		"AuthenticationMethod",
		"ClientEndpointGroups",
		"ServerEndpointGroups",
		"VolumesAreExposed",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagegroup).Elem()

	return storagegroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetStorageGroup will get a StorageGroup instance from the service.
func GetStorageGroup(c common.Client, uri string) (*StorageGroup, error) {
	var storageGroup StorageGroup
	return &storageGroup, storageGroup.Get(c, uri, &storageGroup)
}

// ListReferencedStorageGroups gets the collection of StorageGroup from
// a provided reference.
func ListReferencedStorageGroups(c common.Client, link string) ([]*StorageGroup, error) {
	var result []*StorageGroup
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *StorageGroup
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		storagegroup, err := GetStorageGroup(c, link)
		ch <- GetResult{Item: storagegroup, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ChildStorageGroups gets child groups of this group.
func (storagegroup *StorageGroup) ChildStorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup

	collectionError := common.NewCollectionError()
	for _, sgLink := range storagegroup.childStorageGroups {
		sg, err := GetStorageGroup(storagegroup.GetClient(), sgLink)
		if err != nil {
			collectionError.Failures[sgLink] = err
		} else {
			result = append(result, sg)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ParentStorageGroups gets parent groups of this group.
func (storagegroup *StorageGroup) ParentStorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup

	collectionError := common.NewCollectionError()
	for _, sgLink := range storagegroup.parentStorageGroups {
		sg, err := GetStorageGroup(storagegroup.GetClient(), sgLink)
		if err != nil {
			collectionError.Failures[sgLink] = err
		} else {
			result = append(result, sg)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ClassOfService gets the ClassOfService that all storage in this StorageGroup
// conforms to.
func (storagegroup *StorageGroup) ClassOfService() (*ClassOfService, error) {
	if storagegroup.classOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storagegroup.GetClient(), storagegroup.classOfService)
}

// MappedVolume is an exposed volume mapping.
type MappedVolume struct {
	// LogicalUnitNumber is the value is a SCSI Logical Unit Number for the Volume.
	LogicalUnitNumber int
	// Volume shall reference a mapped Volume.
	Volume common.Link
}

// ExposeVolumes exposes the storage of this group via the target endpoints
// named in the ServerEndpointGroups to the initiator endpoints named in the
// ClientEndpointGroups.  The property VolumesAreExposed shall be set to true
// when this action is completed.
func (storagegroup *StorageGroup) ExposeVolumes() error {
	err := storagegroup.Post(storagegroup.exposeVolumesTarget, nil)
	if err == nil {
		// Only set to exposed if no error. Calling expose when already exposed
		// could fail so we don't want to indicate they are not exposed.
		storagegroup.VolumesAreExposed = true
	}
	return err
}

// HideVolumes hides the storage of this group from the initiator endpoints
// named in the ClientEndpointGroups. The property VolumesAreExposed shall be
// set to false when this action is completed.
func (storagegroup *StorageGroup) HideVolumes() error {
	err := storagegroup.Post(storagegroup.hideVolumesTarget, nil)
	if err == nil {
		storagegroup.VolumesAreExposed = false
	}
	return err
}
