//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/StorageGroup.v1_6_0.json
// 1.2.6 - #StorageGroup.v1_6_0.StorageGroup

package schemas

import (
	"encoding/json"
)

type AuthenticationMethod string

const (
	// NoneAuthenticationMethod No authentication is used.
	NoneAuthenticationMethod AuthenticationMethod = "None"
	// CHAPAuthenticationMethod shall be used when type CHAP is selected.
	CHAPAuthenticationMethod AuthenticationMethod = "CHAP"
	// MutualCHAPAuthenticationMethod shall be used when type MutualCHAP is
	// selected.
	MutualCHAPAuthenticationMethod AuthenticationMethod = "MutualCHAP"
	// DHCHAPAuthenticationMethod shall be used instead of CHAPInfo, and the
	// LocalDHCHAPAuthSecret and PeerDHCHAPAuthSecret properties shall be used.
	DHCHAPAuthenticationMethod AuthenticationMethod = "DHCHAP"
)

// StorageGroup is a storage group collects a set of related storage entities
// (volumes, file systems...) The collection should be useful for managing the
// storage of a set of related client applications.
type StorageGroup struct {
	Entity
	// AccessState shall describe the access characteristics of this storage group.
	// All associated logical units through all aggregated ports shall share this
	// access state.
	AccessState AccessState
	// AuthenticationMethod The value of this property must be what kind of
	// authentication that the endpoints in this StorageGroup understands.
	//
	// Version added: v1.2.0
	AuthenticationMethod AuthenticationMethod
	// ChapInfo The value of this property must reflect the authentication used by
	// this specific endpoint. If this endpoint represents an initiator, and
	// AuthenticationMethod is CHAP or MutualCHAP, the Credentials fields
	// CHAPUsername and CHAPSecret must be used. If this endpoint represents a
	// target endpoint and AuthenticationMethod is MutualCHAP, then
	// MutualCHAPUsername and MutualCHAPSecret must be used.
	//
	// Version added: v1.2.0
	ChapInfo []CHAPInformation
	// ClientEndpointGroups shall not allow access to the storage via any
	// client-side endpoint.
	clientEndpointGroups []string
	// ClientEndpointGroupsCount
	ClientEndpointGroupsCount int `json:"ClientEndpointGroups@odata.count"`
	// DHChapInfo The value of this property must reflect the authentication used
	// by this specific endpoint when the authentication type is specificed as
	// DHCHAP. If this endpoint represents an initiator, and AuthenticationMethod
	// is DHCHAP, the Credentials fields LocalDHCHAPAuthSecret and
	// PeerDHCHAPAuthSecret must be used.
	//
	// Version added: v1.3.0
	DHChapInfo []DHCHAPInformation
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// MappedVolumes is an array of mapped volumes managed by this storage group.
	//
	// Version added: v1.1.0
	MappedVolumes []MappedVolume
	// MembersAreConsistent shall be set to true if all members are in a consistent
	// state. The default value for this property is false.
	//
	// Deprecated
	// Deprecated in favor of using the ConsistencyGroup for Consistency set
	// management.
	MembersAreConsistent bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReplicaInfo shall describe the replication relationship between this storage
	// group and a corresponding source storage group.
	//
	// Deprecated
	// Deprecated in favor of using the ConsistencyGroup for Consistency set
	// management.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this
	// replica.
	//
	// Version added: v1.1.1
	//
	// Deprecated
	// Deprecated in favor of using the ConsistencyGroup for Consistency set
	// management.
	ReplicaTargets []Entity
	// ReplicaTargetsCount
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// ServerEndpointGroups shall not allow access to the storage via any
	// server-side endpoint.
	serverEndpointGroups []string
	// ServerEndpointGroupsCount
	ServerEndpointGroupsCount int `json:"ServerEndpointGroups@odata.count"`
	// Status shall contain the status of the StorageGroup.
	Status Status
	// Volumes is an array of references to volumes managed by this storage group.
	//
	// Deprecated
	// These references are replaced by the MappedVolumes array in StorageGroup.
	volumes []string
	// VolumesAreExposed shall be set to true if storage volumes are exposed to the
	// paths defined by the client and server endpoints. The default value for this
	// property is false.
	VolumesAreExposed bool
	// VolumesCount
	VolumesCount int `json:"Volumes@odata.count"`
	// exposeVolumesTarget is the URL to send ExposeVolumes requests.
	exposeVolumesTarget string
	// hideVolumesTarget is the URL to send HideVolumes requests.
	hideVolumesTarget string
	// childStorageGroups are the URIs for ChildStorageGroups.
	childStorageGroups []string
	// classOfService is the URI for ClassOfService.
	classOfService string
	// parentStorageGroups are the URIs for ParentStorageGroups.
	parentStorageGroups []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a StorageGroup object from the raw JSON.
func (s *StorageGroup) UnmarshalJSON(b []byte) error {
	type temp StorageGroup
	type sActions struct {
		ExposeVolumes ActionTarget `json:"#StorageGroup.ExposeVolumes"`
		HideVolumes   ActionTarget `json:"#StorageGroup.HideVolumes"`
	}
	type sLinks struct {
		ChildStorageGroups  Links `json:"ChildStorageGroups"`
		ClassOfService      Link  `json:"ClassOfService"`
		ParentStorageGroups Links `json:"ParentStorageGroups"`
	}
	var tmp struct {
		temp
		Actions              sActions
		Links                sLinks
		ClientEndpointGroups Links `json:"ClientEndpointGroups"`
		ServerEndpointGroups Links `json:"ServerEndpointGroups"`
		Volumes              Links `json:"Volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageGroup(tmp.temp)

	// Extract the links to other entities for later
	s.exposeVolumesTarget = tmp.Actions.ExposeVolumes.Target
	s.hideVolumesTarget = tmp.Actions.HideVolumes.Target
	s.childStorageGroups = tmp.Links.ChildStorageGroups.ToStrings()
	s.classOfService = tmp.Links.ClassOfService.String()
	s.parentStorageGroups = tmp.Links.ParentStorageGroups.ToStrings()
	s.clientEndpointGroups = tmp.ClientEndpointGroups.ToStrings()
	s.serverEndpointGroups = tmp.ServerEndpointGroups.ToStrings()
	s.volumes = tmp.Volumes.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageGroup) Update() error {
	readWriteFields := []string{
		"AccessState",
		"AuthenticationMethod",
		"ClientEndpointGroups",
		"MembersAreConsistent",
		"ServerEndpointGroups",
		"Volumes",
		"VolumesAreExposed",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetStorageGroup will get a StorageGroup instance from the service.
func GetStorageGroup(c Client, uri string) (*StorageGroup, error) {
	return GetObject[StorageGroup](c, uri)
}

// ListReferencedStorageGroups gets the collection of StorageGroup from
// a provided reference.
func ListReferencedStorageGroups(c Client, link string) ([]*StorageGroup, error) {
	return GetCollectionObjects[StorageGroup](c, link)
}

// Exposes the storage of this group via the target endpoints named in the
// ServerEndpointGroups to the initiator endpoints named in the
// ClientEndpointGroups. The property VolumesAreExposed shall be set to true
// when this action is completed.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageGroup) ExposeVolumes() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.exposeVolumesTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Hide the storage of this group from the initiator endpoints named in the
// ClientEndpointGroups. The property VolumesAreExposed shall be set to false
// when this action is completed.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageGroup) HideVolumes() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.hideVolumesTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ChildStorageGroups gets the ChildStorageGroups linked resources.
func (s *StorageGroup) ChildStorageGroups() ([]*StorageGroup, error) {
	return GetObjects[StorageGroup](s.client, s.childStorageGroups)
}

// ClassOfService gets the ClassOfService linked resource.
func (s *StorageGroup) ClassOfService() (*ClassOfService, error) {
	if s.classOfService == "" {
		return nil, nil
	}
	return GetObject[ClassOfService](s.client, s.classOfService)
}

// ParentStorageGroups gets the ParentStorageGroups linked resources.
func (s *StorageGroup) ParentStorageGroups() ([]*StorageGroup, error) {
	return GetObjects[StorageGroup](s.client, s.parentStorageGroups)
}

// ClientEndpointGroups gets the ClientEndpointGroups linked resources.
func (s *StorageGroup) ClientEndpointGroups() ([]*EndpointGroup, error) {
	return GetObjects[EndpointGroup](s.client, s.clientEndpointGroups)
}

// ServerEndpointGroups gets the ServerEndpointGroups linked resources.
func (s *StorageGroup) ServerEndpointGroups() ([]*EndpointGroup, error) {
	return GetObjects[EndpointGroup](s.client, s.serverEndpointGroups)
}

// Volumes gets the Volumes linked resources.
func (s *StorageGroup) Volumes() ([]*Volume, error) {
	return GetObjects[Volume](s.client, s.volumes)
}

// CHAPInformation User name and password values for target and initiators
// Endpoints when CHAP authentication is used.
type CHAPInformation struct {
	// CHAPPassword shall be the password when CHAP authentication is specified.
	//
	// Version added: v1.3.0
	CHAPPassword string
	// CHAPUser shall be the username when CHAP authentication is specified.
	//
	// Version added: v1.3.0
	CHAPUser string
	// InitiatorCHAPPassword shall be the shared secret for Mutual (2-way)CHAP
	// authentication.
	//
	// Version added: v1.2.0
	InitiatorCHAPPassword string
	// InitiatorCHAPUser If present, this property is the initiator CHAP username
	// for Mutual (2-way) authentication. For example, with an iSCSI scenario, use
	// the initiator iQN.
	//
	// Version added: v1.2.0
	InitiatorCHAPUser string
	// TargetCHAPPassword shall be the CHAP Secret for 2-way CHAP authentication.
	//
	// Version added: v1.3.0
	TargetCHAPPassword string
	// TargetCHAPUser shall be the Target CHAP Username for Mutual (2-way) CHAP
	// authentication. For example, with an iSCSI scenario, use the target iQN.
	//
	// Version added: v1.2.0
	TargetCHAPUser string
	// TargetPassword shall be the CHAP Secret for 2-way CHAP authentication.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.3.0
	// This property is deprecated in favor of TargetCHAPPassword.
	TargetPassword string
}

// DHCHAPInformation User name and password values for target and initiator
// endpoints when CHAP authentication is used.
type DHCHAPInformation struct {
	// LocalDHCHAPAuthSecret shall be the local DHCHAP auth secret for DHCHAP
	// authentication.
	//
	// Version added: v1.3.0
	LocalDHCHAPAuthSecret string
	// PeerDHCHAPAuthSecret shall be the peer DHCHAP auth secret for DHCHAP
	// authentication.
	//
	// Version added: v1.3.0
	PeerDHCHAPAuthSecret string
}

// MappedVolume Relate a SCSI Logical Unit Number to a Volume.
type MappedVolume struct {
	// AccessCapability shall specify the storage access capability for this mapped
	// volume.
	//
	// Version added: v1.4.0
	AccessCapability AccessCapability
	// LogicalUnitNumber If present, the value is a SCSI Logical Unit Number for
	// the Volume.
	LogicalUnitNumber string
	// Volume shall reference a mapped Volume.
	volume string
}

// UnmarshalJSON unmarshals a MappedVolume object from the raw JSON.
func (m *MappedVolume) UnmarshalJSON(b []byte) error {
	type temp MappedVolume
	var tmp struct {
		temp
		LogicalUnitNumber any
		Volume            Link `json:"Volume"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MappedVolume(tmp.temp)

	// Extract the links to other entities for later
	m.volume = tmp.Volume.String()
	m.LogicalUnitNumber = parseString(tmp.LogicalUnitNumber)

	return nil
}

// Volume gets the Volume linked resource.
func (m *MappedVolume) Volume(client Client) (*Volume, error) {
	if m.volume == "" {
		return nil, nil
	}
	return GetObject[Volume](client, m.volume)
}
