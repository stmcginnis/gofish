//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.4 - #Connection.v1_4_0.Connection

package schemas

import (
	"encoding/json"
)

type AccessCapability string

const (
	// ReadAccessCapability Endpoints are allowed to perform reads from the
	// specified resource.
	ReadAccessCapability AccessCapability = "Read"
	// WriteAccessCapability Endpoints are allowed to perform writes to the
	// specified resource.
	WriteAccessCapability AccessCapability = "Write"
	// ReadWriteAccessCapability Endpoints are allowed to perform reads from and
	// writes to the specified resource.
	ReadWriteAccessCapability AccessCapability = "ReadWrite" // from storagegroup
)

// AccessState is This type shall describe the access to the associated resource
// in this connection.
type AccessState string

const (
	// OptimizedAccessState shall indicate the resource is in an active and
	// optimized state.
	OptimizedAccessState AccessState = "Optimized"
	// NonOptimizedAccessState shall indicate the resource is in an active and
	// non-optimized state.
	NonOptimizedAccessState AccessState = "NonOptimized"
	// StandbyAccessState shall indicate the resource is in a standby state.
	StandbyAccessState AccessState = "Standby"
	// UnavailableAccessState shall indicate the resource is in an unavailable
	// state.
	UnavailableAccessState AccessState = "Unavailable"
	// TransitioningAccessState shall indicate the resource is transitioning to a
	// new state.
	TransitioningAccessState AccessState = "Transitioning"
)

type ConnectionType string

const (
	// StorageConnectionType is a connection to storage-related resources, such as
	// volumes.
	StorageConnectionType ConnectionType = "Storage"
	// MemoryConnectionType is a connection to memory-related resources.
	MemoryConnectionType ConnectionType = "Memory"
)

// Connection shall represent information about a connection in the Redfish
// Specification.
type Connection struct {
	Entity
	// ConnectionKeys shall contain the permission keys required to access the
	// specified resources for this connection. Some fabrics require permission
	// checks on transactions from authorized initiators.
	//
	// Version added: v1.1.0
	ConnectionKeys ConnectionKey
	// ConnectionType shall contain the type of resources this connection
	// specifies.
	ConnectionType ConnectionType
	// MemoryChunkInfo shall contain the set of memory chunks and access
	// capabilities specified for this connection.
	//
	// Version added: v1.1.0
	MemoryChunkInfo []MemoryChunkInfo
	// MemoryRegionInfo shall contain the set of memory regions and access
	// capabilities specified for this connection.
	//
	// Version added: v1.3.0
	MemoryRegionInfo []MemoryRegionInfo
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// VolumeInfo shall contain the set of volumes and access capabilities
	// specified for this connection.
	VolumeInfo []VolumeInfo
	// addVolumeInfoTarget is the URL to send AddVolumeInfo requests.
	addVolumeInfoTarget string
	// removeVolumeInfoTarget is the URL to send RemoveVolumeInfo requests.
	removeVolumeInfoTarget string
	// initiatorEndpointGroups are the URIs for InitiatorEndpointGroups.
	initiatorEndpointGroups []string
	// initiatorEndpoints are the URIs for InitiatorEndpoints.
	initiatorEndpoints []string
	// targetEndpointGroups are the URIs for TargetEndpointGroups.
	targetEndpointGroups []string
	// targetEndpoints are the URIs for TargetEndpoints.
	targetEndpoints []string
}

// UnmarshalJSON unmarshals a Connection object from the raw JSON.
func (c *Connection) UnmarshalJSON(b []byte) error {
	type temp Connection
	type cActions struct {
		AddVolumeInfo    ActionTarget `json:"#Connection.AddVolumeInfo"`
		RemoveVolumeInfo ActionTarget `json:"#Connection.RemoveVolumeInfo"`
	}
	type cLinks struct {
		InitiatorEndpointGroups Links `json:"InitiatorEndpointGroups"`
		InitiatorEndpoints      Links `json:"InitiatorEndpoints"`
		TargetEndpointGroups    Links `json:"TargetEndpointGroups"`
		TargetEndpoints         Links `json:"TargetEndpoints"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Connection(tmp.temp)

	// Extract the links to other entities for later
	c.addVolumeInfoTarget = tmp.Actions.AddVolumeInfo.Target
	c.removeVolumeInfoTarget = tmp.Actions.RemoveVolumeInfo.Target
	c.initiatorEndpointGroups = tmp.Links.InitiatorEndpointGroups.ToStrings()
	c.initiatorEndpoints = tmp.Links.InitiatorEndpoints.ToStrings()
	c.targetEndpointGroups = tmp.Links.TargetEndpointGroups.ToStrings()
	c.targetEndpoints = tmp.Links.TargetEndpoints.ToStrings()

	return nil
}

// GetConnection will get a Connection instance from the service.
func GetConnection(c Client, uri string) (*Connection, error) {
	return GetObject[Connection](c, uri)
}

// ListReferencedConnections gets the collection of Connection from
// a provided reference.
func ListReferencedConnections(c Client, link string) ([]*Connection, error) {
	return GetCollectionObjects[Connection](c, link)
}

// This action shall add a volume to the connection. Services shall add the
// volume to the 'VolumeInfo' property.
// accessCapabilities - This parameter shall contain an array of the storage
// access capabilities to assign to the volume. Services shall reject requests
// that do not contain either 'LUN' or 'AccessCapabilities'.
// lUN - This property shall contain the initiator-visible logical unit number
// (LUN) to assign to the volume. Services shall reject requests that do not
// contain either 'LUN' or 'AccessCapabilities'.
// volume - This parameter shall contain a link to a resource of type 'Volume'
// that represents the volume to add.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Connection) AddVolumeInfo(accessCapabilities []AccessCapability, lUN int, volume string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["AccessCapabilities"] = accessCapabilities
	payload["LUN"] = lUN
	payload["Volume"] = volume
	resp, taskInfo, err := PostWithTask(c.client,
		c.addVolumeInfoTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall remove a volume to the connection. Services shall remove
// the volume from the 'VolumeInfo' property.
// lUN - This parameter shall contain the initiator-visible logical unit number
// (LUN) assigned to this volume to remove. If this parameter is not provided,
// the service shall remove all entries associated with volume referenced by
// the 'Volume' parameter.
// volume - This parameter shall contain a link to a resource of type 'Volume'
// that represents the volume to remove.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Connection) RemoveVolumeInfo(lUN int, volume string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["LUN"] = lUN
	payload["Volume"] = volume
	resp, taskInfo, err := PostWithTask(c.client,
		c.removeVolumeInfoTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// InitiatorEndpointGroups gets the InitiatorEndpointGroups linked resources.
func (c *Connection) InitiatorEndpointGroups() ([]*EndpointGroup, error) {
	return GetObjects[EndpointGroup](c.client, c.initiatorEndpointGroups)
}

// InitiatorEndpoints gets the InitiatorEndpoints linked resources.
func (c *Connection) InitiatorEndpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](c.client, c.initiatorEndpoints)
}

// TargetEndpointGroups gets the TargetEndpointGroups linked resources.
func (c *Connection) TargetEndpointGroups() ([]*EndpointGroup, error) {
	return GetObjects[EndpointGroup](c.client, c.targetEndpointGroups)
}

// TargetEndpoints gets the TargetEndpoints linked resources.
func (c *Connection) TargetEndpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](c.client, c.targetEndpoints)
}

// CHAPConnectionKey shall contain the CHAP-specific permission key information
// for a connection.
type CHAPConnectionKey struct {
	// CHAPPassword shall contain the password for CHAP authentication. The value
	// shall be 'null' in responses.
	//
	// Version added: v1.2.0
	CHAPPassword string
	// CHAPUsername shall contain the username for CHAP authentication.
	//
	// Version added: v1.2.0
	CHAPUsername string
	// InitiatorCHAPPassword shall contain the initiator shared secret for mutual
	// (2-way) CHAP authentication. The value shall be 'null' in responses.
	//
	// Version added: v1.2.0
	InitiatorCHAPPassword string
	// InitiatorCHAPUsername shall contain the initiator username for mutual
	// (2-way) CHAP authentication. For example, this would be the initiator iQN in
	// iSCSI environments.
	//
	// Version added: v1.2.0
	InitiatorCHAPUsername string
	// TargetCHAPPassword shall contain the target shared secret for mutual (2-way)
	// CHAP authentication. The value shall be 'null' in responses.
	//
	// Version added: v1.2.0
	TargetCHAPPassword string
}

// ConnectionKey shall contain the permission key information required to access
// the target resources for a connection.
type ConnectionKey struct {
	// CHAP shall contain the CHAP-specific permission key information for this
	// connection. This property shall not be present if 'DHCHAP' is present.
	//
	// Version added: v1.2.0
	CHAP CHAPConnectionKey
	// DHCHAP shall contain the DHCHAP-specific permission key information for this
	// connection. This property shall not be present if 'CHAP' is present.
	//
	// Version added: v1.2.0
	DHCHAP DHCHAPKey
	// GenZ shall contain the Gen-Z-specific permission key information for this
	// connection.
	//
	// Version added: v1.1.0
	GenZ GenZConnectionKey
}

// DHCHAPKey shall contain the DHCHAP-specific permission key information for
// this connection.
type DHCHAPKey struct {
	// LocalDHCHAPAuthSecret shall contain the local DHCHAP authentication secret.
	// The value shall be 'null' in responses.
	//
	// Version added: v1.2.0
	LocalDHCHAPAuthSecret string
	// PeerDHCHAPAuthSecret shall contain the peer DHCHAP authentication secret.
	// The value shall be 'null' in responses.
	//
	// Version added: v1.2.0
	PeerDHCHAPAuthSecret string
}

// GenZConnectionKey shall contain the Gen-Z-specific permission key information
// for a connection.
type GenZConnectionKey struct {
	// AccessKey shall contain the Gen-Z Core Specification-defined Access Key for
	// this connection.
	//
	// Version added: v1.1.0
	AccessKey string
	// RKeyDomainCheckingEnabled shall indicate whether Region Key domain checking
	// is enabled for this connection.
	//
	// Version added: v1.1.0
	RKeyDomainCheckingEnabled bool
	// RKeyReadOnlyKey shall contain the Gen-Z Core Specification-defined read-only
	// Region Key for this connection.
	//
	// Version added: v1.1.0
	RKeyReadOnlyKey string
	// RKeyReadWriteKey shall contain the Gen-Z Core Specification-defined
	// read-write Region Key for this connection.
	//
	// Version added: v1.1.0
	RKeyReadWriteKey string
}

// MemoryChunkInfo shall contain the combination of permissions and memory chunk
// information.
type MemoryChunkInfo struct {
	// AccessCapabilities shall specify a current memory access capability.
	//
	// Version added: v1.1.0
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in
	// this connection.
	//
	// Version added: v1.1.0
	AccessState AccessState
	// MemoryChunk shall contain a link to a resource of type 'MemoryChunk'. The
	// endpoints referenced by the 'InitiatorEndpoints' or
	// 'InitiatorEndpointGroups' properties shall be given access to this memory
	// chunk as described by this object. If 'TargetEndpoints' or
	// 'TargetEndpointGroups' is present, the referenced initiator endpoints shall
	// be required to access the referenced memory chunk through one of the
	// referenced target endpoints.
	//
	// Version added: v1.1.0
	memoryChunk string
}

// UnmarshalJSON unmarshals a MemoryChunkInfo object from the raw JSON.
func (m *MemoryChunkInfo) UnmarshalJSON(b []byte) error {
	type temp MemoryChunkInfo
	var tmp struct {
		temp
		MemoryChunk Link `json:"MemoryChunk"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryChunkInfo(tmp.temp)

	// Extract the links to other entities for later
	m.memoryChunk = tmp.MemoryChunk.String()

	return nil
}

// MemoryChunk gets the MemoryChunk linked resource.
func (m *MemoryChunkInfo) MemoryChunk(client Client) (*MemoryChunks, error) {
	if m.memoryChunk == "" {
		return nil, nil
	}
	return GetObject[MemoryChunks](client, m.memoryChunk)
}

// MemoryRegionInfo shall contain the combination of permissions and memory
// region information.
type MemoryRegionInfo struct {
	// AccessCapabilities shall specify a current memory access capability.
	//
	// Version added: v1.3.0
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in
	// this connection.
	//
	// Version added: v1.3.0
	AccessState AccessState
	// MemoryRegion shall contain a link to a resource of type 'MemoryRegion'. The
	// endpoints referenced by the 'InitiatorEndpoints' or
	// 'InitiatorEndpointGroups' properties shall be given access to this memory
	// region as described by this object. If 'TargetEndpoints' or
	// 'TargetEndpointGroups' is present, the referenced initiator endpoints shall
	// be required to access the referenced memory region through one of the
	// referenced target endpoints. For CXL fabrics, memory regions from
	// 'Connection' resources are not allowed.
	//
	// Version added: v1.3.0
	memoryRegion string
}

// UnmarshalJSON unmarshals a MemoryRegionInfo object from the raw JSON.
func (m *MemoryRegionInfo) UnmarshalJSON(b []byte) error {
	type temp MemoryRegionInfo
	var tmp struct {
		temp
		MemoryRegion Link `json:"MemoryRegion"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryRegionInfo(tmp.temp)

	// Extract the links to other entities for later
	m.memoryRegion = tmp.MemoryRegion.String()

	return nil
}

// MemoryRegion gets the MemoryRegion linked resource.
func (m *MemoryRegionInfo) MemoryRegion(client Client) (*MemoryRegion, error) {
	if m.memoryRegion == "" {
		return nil, nil
	}
	return GetObject[MemoryRegion](client, m.memoryRegion)
}

// VolumeInfo shall contain the combination of permissions and volume
// information.
type VolumeInfo struct {
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in
	// this connection.
	AccessState AccessState
	// LUN shall contain the initiator-visible logical unit number (LUN) assigned
	// to this volume for initiators referenced by the 'InitiatorEndpoints' or
	// 'InitiatorEndpointGroups' properties.
	//
	// Version added: v1.2.0
	LUN *int `json:",omitempty"`
	// Volume shall contain a link to a resource of type 'Volume'. The endpoints
	// referenced by the 'InitiatorEndpoints' or 'InitiatorEndpointGroups'
	// properties shall be given access to this volume as described by this object.
	// If 'TargetEndpoints' or 'TargetEndpointGroups' is present, the referenced
	// initiator endpoints shall be required to access the referenced volume
	// through one of the referenced target endpoints.
	volume string
}

// UnmarshalJSON unmarshals a VolumeInfo object from the raw JSON.
func (v *VolumeInfo) UnmarshalJSON(b []byte) error {
	type temp VolumeInfo
	var tmp struct {
		temp
		Volume Link `json:"Volume"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VolumeInfo(tmp.temp)

	// Extract the links to other entities for later
	v.volume = tmp.Volume.String()

	return nil
}

// Volume gets the Volume linked resource.
func (v *VolumeInfo) Volume(client Client) (*Volume, error) {
	if v.volume == "" {
		return nil, nil
	}
	return GetObject[Volume](client, v.volume)
}
