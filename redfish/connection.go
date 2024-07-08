//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type AccessCapability string

const (
	// ReadAccessCapability Endpoints are allowed to perform reads from the specified resource.
	ReadAccessCapability AccessCapability = "Read"
	// WriteAccessCapability Endpoints are allowed to perform writes to the specified resource.
	WriteAccessCapability AccessCapability = "Write"
)

// AccessState describes the access to the associated resource in this connection.
type AccessState string

const (
	// OptimizedAccessState shall indicate the resource is in an active and optimized state.
	OptimizedAccessState AccessState = "Optimized"
	// NonOptimizedAccessState shall indicate the resource is in an active and non-optimized state.
	NonOptimizedAccessState AccessState = "NonOptimized"
	// StandbyAccessState shall indicate the resource is in a standby state.
	StandbyAccessState AccessState = "Standby"
	// UnavailableAccessState shall indicate the resource is in an unavailable state.
	UnavailableAccessState AccessState = "Unavailable"
	// TransitioningAccessState shall indicate the resource is transitioning to a new state.
	TransitioningAccessState AccessState = "Transitioning"
)

type ConnectionType string

const (
	// StorageConnectionType is a connection to storage-related resources, such as volumes.
	StorageConnectionType ConnectionType = "Storage"
	// MemoryConnectionType is a connection to memory-related resources.
	MemoryConnectionType ConnectionType = "Memory"
)

// CHAPConnectionKey shall contain the CHAP-specific permission key information for a connection.
type CHAPConnectionKey struct {
	// CHAPPassword shall contain the password for CHAP authentication. The value shall be 'null' in responses.
	CHAPPassword string
	// CHAPUsername shall contain the username for CHAP authentication.
	CHAPUsername string
	// InitiatorCHAPPassword shall contain the initiator shared secret for mutual (2-way) CHAP authentication. The
	// value shall be 'null' in responses.
	InitiatorCHAPPassword string
	// InitiatorCHAPUsername shall contain the initiator username for mutual (2-way) CHAP authentication. For example,
	// this would be the initiator iQN in iSCSI environments.
	InitiatorCHAPUsername string
	// TargetCHAPPassword shall contain the target shared secret for mutual (2-way) CHAP authentication. The value
	// shall be 'null' in responses.
	TargetCHAPPassword string
}

// Connection shall represent information about a connection in the Redfish Specification.
type Connection struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConnectionKeys shall contain the permission keys required to access the specified resources for this connection.
	// Some fabrics require permission checks on transactions from authorized initiators.
	ConnectionKeys ConnectionKey
	// ConnectionType shall contain the type of resources this connection specifies.
	ConnectionType ConnectionType
	// Description provides a description of this resource.
	Description string
	// MemoryChunkInfo shall contain the set of memory chunks and access capabilities specified for this connection.
	MemoryChunkInfo []MemoryChunkInfo
	// MemoryRegionInfo shall contain the set of memory regions and access capabilities specified for this connection.
	MemoryRegionInfo []MemoryRegionInfo
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VolumeInfo shall contain the set of volumes and access capabilities specified for this connection.
	VolumeInfo []VolumeInfo

	initiatorEndpointGroups []string
	// InitiatorEndpointGroupsCount is the number of initiator endpoint groups associated with this connection.
	InitiatorEndpointGroupsCount int
	initiatorEndpoints           []string
	// InititiatorEndpointsCount is the number of initiator endpoints associated with this connection.
	InitiatorEndpointsCount int
	targetEndpointGroups    []string
	// TargetEndpointGroupsCount is the number of target endpoint groups associated with this connection.
	TargetEndpointGroupsCount int
	targetEndpoints           []string
	// TargetEndpointsCount is the number of target endpoints associated with this connection.
	TargetEndpointsCount int
}

// UnmarshalJSON unmarshals a Connection object from the raw JSON.
func (connection *Connection) UnmarshalJSON(b []byte) error {
	type temp Connection
	type Links struct {
		// InitiatorEndpointGroups shall contain an array of links to resources of type EndpointGroup that are the
		// initiator endpoint groups associated with this connection. If the referenced endpoint groups contain the
		// GroupType property, the GroupType property shall contain the value 'Initiator' or 'Client'. This property shall
		// not be present if InitiatorEndpoints is present.
		InitiatorEndpointGroups      common.Links
		InitiatorEndpointGroupsCount int `json:"InitiatorEndpointGroups@odata.count"`
		// InitiatorEndpoints shall contain an array of links to resources of type Endpoint that are the initiator
		// endpoints associated with this connection. If the referenced endpoints contain the EntityRole property, the
		// EntityRole property shall contain the value 'Initiator' or 'Both'. This property shall not be present if
		// InitiatorEndpointGroups is present.
		InitiatorEndpoints      common.Links
		InitiatorEndpointsCount int `json:"InitiatorEndpoints@odata.count"`
		// TargetEndpointGroups shall contain an array of links to resources of type EndpointGroup that are the target
		// endpoint groups associated with this connection. If the referenced endpoint groups contain the GroupType
		// property, the GroupType property shall contain the value 'Target' or 'Server'. This property shall not be
		// present if TargetEndpoints is present.
		TargetEndpointGroups      common.Links
		TargetEndpointGroupsCount int `json:"TargetEndpointGroups@odata.count"`
		// TargetEndpoints shall contain an array of links to resources of type Endpoint that are the target endpoints
		// associated with this connection. If the referenced endpoints contain the EntityRole property, the EntityRole
		// property shall contain the value 'Target' or 'Both'. This property shall not be present if TargetEndpointGroups
		// is present.
		TargetEndpoints      common.Links
		TargetEndpointsCount int `json:"TargetEndpoints@odata.count"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connection = Connection(t.temp)

	// Extract the links to other entities for later
	connection.initiatorEndpointGroups = t.Links.InitiatorEndpointGroups.ToStrings()
	connection.InitiatorEndpointGroupsCount = t.Links.InitiatorEndpointGroupsCount
	connection.initiatorEndpoints = t.Links.InitiatorEndpoints.ToStrings()
	connection.InitiatorEndpointsCount = t.Links.InitiatorEndpointsCount
	connection.targetEndpointGroups = t.Links.TargetEndpointGroups.ToStrings()
	connection.TargetEndpointGroupsCount = t.Links.TargetEndpointGroupsCount
	connection.targetEndpoints = t.Links.TargetEndpoints.ToStrings()
	connection.TargetEndpointsCount = t.Links.TargetEndpointsCount

	return nil
}

// GetConnection will get a Connection instance from the service.
func GetConnection(c common.Client, uri string) (*Connection, error) {
	return common.GetObject[Connection](c, uri)
}

// ListReferencedConnections gets the collection of Connection from
// a provided reference.
func ListReferencedConnections(c common.Client, link string) ([]*Connection, error) {
	return common.GetCollectionObjects[Connection](c, link)
}

// InitiatorEndpointGroups get the initiator endpoint groups associated with this connection.
func (connection *Connection) InitiatorEndpointGroups() ([]*EndpointGroup, error) {
	return common.GetObjects[EndpointGroup](connection.GetClient(), connection.initiatorEndpointGroups)
}

// InitiatorEndpoints get the initiator endpoint associated with this connection.
func (connection *Connection) InitiatorEndpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](connection.GetClient(), connection.initiatorEndpoints)
}

// TargetEndpointGroups get the target endpoint groups associated with this connection.
func (connection *Connection) TargetEndpointGroups() ([]*EndpointGroup, error) {
	return common.GetObjects[EndpointGroup](connection.GetClient(), connection.targetEndpointGroups)
}

// TargetEndpoints get the target endpoint associated with this connection.
func (connection *Connection) TargetEndpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](connection.GetClient(), connection.targetEndpoints)
}

// ConnectionKey shall contain the permission key information required to access the target resources for a
// connection.
type ConnectionKey struct {
	// CHAP shall contain the CHAP-specific permission key information for this connection. This property shall not be
	// present if DHCHAP is present.
	CHAP CHAPConnectionKey
	// DHCHAP shall contain the DHCHAP-specific permission key information for this connection. This property shall not
	// be present if CHAP is present.
	DHCHAP DHCHAPKey
	// GenZ shall contain the Gen-Z-specific permission key information for this connection.
	GenZ GenZConnectionKey
}

// DHCHAPKey shall contain the DHCHAP-specific permission key information for this connection.
type DHCHAPKey struct {
	// LocalDHCHAPAuthSecret shall contain the local DHCHAP authentication secret. The value shall be 'null' in
	// responses.
	LocalDHCHAPAuthSecret string
	// PeerDHCHAPAuthSecret shall contain the peer DHCHAP authentication secret. The value shall be 'null' in
	// responses.
	PeerDHCHAPAuthSecret string
}

// GenZConnectionKey shall contain the Gen-Z-specific permission key information for a connection.
type GenZConnectionKey struct {
	// AccessKey shall contain the Gen-Z Core Specification-defined Access Key for this connection.
	AccessKey string
	// RKeyDomainCheckingEnabled shall indicate whether Region Key domain checking is enabled for this connection.
	RKeyDomainCheckingEnabled string
	// RKeyReadOnlyKey shall contain the Gen-Z Core Specification-defined read-only Region Key for this connection.
	RKeyReadOnlyKey string
	// RKeyReadWriteKey shall contain the Gen-Z Core Specification-defined read-write Region Key for this connection.
	RKeyReadWriteKey string
}

// MemoryChunkInfo shall contain the combination of permissions and memory chunk information.
type MemoryChunkInfo struct {
	// AccessCapabilities shall specify a current memory access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in this connection.
	AccessState AccessState
	// MemoryChunk shall contain a link to a resource of type MemoryChunk. The endpoints referenced by the
	// InitiatorEndpoints or InitiatorEndpointGroups properties shall be given access to this memory chunk as described
	// by this object. If TargetEndpoints or TargetEndpointGroups is present, the referenced initiator endpoints shall
	// be required to access the referenced memory chunk through one of the referenced target endpoints.
	MemoryChunk MemoryChunks
}

// MemoryRegionInfo shall contain the combination of permissions and memory region information.
type MemoryRegionInfo struct {
	// AccessCapabilities shall specify a current memory access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in this connection.
	AccessState AccessState
	// MemoryRegion shall contain a link to a resource of type MemoryRegion. The endpoints referenced by the
	// InitiatorEndpoints or InitiatorEndpointGroups properties shall be given access to this memory region as
	// described by this object. If TargetEndpoints or TargetEndpointGroups is present, the referenced initiator
	// endpoints shall be required to access the referenced memory region through one of the referenced target
	// endpoints.
	MemoryRegion MemoryRegion
}

// VolumeInfo shall contain the combination of permissions and volume information.
type VolumeInfo struct {
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in this connection.
	AccessState AccessState
	// LUN shall contain the initiator-visible logical unit number (LUN) assigned to this volume for initiators
	// referenced by the InitiatorEndpoints or InitiatorEndpointGroups properties.
	LUN int
	// Volume shall contain a link to a resource of type Volume. The endpoints referenced by the InitiatorEndpoints or
	// InitiatorEndpointGroups properties shall be given access to this volume as described by this object. If
	// TargetEndpoints or TargetEndpointGroups is present, the referenced initiator endpoints shall be required to
	// access the referenced volume through one of the referenced target endpoints.
	Volume common.Link
}
