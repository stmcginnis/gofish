//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ResourceBlock.v1_4_3.json
// 2020.4 - #ResourceBlock.v1_4_3.ResourceBlock

package schemas

import (
	"encoding/json"
)

type CompositionState string

const (
	// ComposingCompositionState Intermediate state indicating composition is in
	// progress.
	ComposingCompositionState CompositionState = "Composing"
	// ComposedAndAvailableCompositionState The resource block is currently
	// participating in one or more compositions, and is available to use in more
	// compositions.
	ComposedAndAvailableCompositionState CompositionState = "ComposedAndAvailable"
	// ComposedCompositionState Final successful state of a resource block that has
	// participated in composition.
	ComposedCompositionState CompositionState = "Composed"
	// UnusedCompositionState The resource block is free and can participate in
	// composition.
	UnusedCompositionState CompositionState = "Unused"
	// FailedCompositionState The final composition resulted in failure and manual
	// intervention might be required to fix it.
	FailedCompositionState CompositionState = "Failed"
	// UnavailableCompositionState The resource block has been made unavailable by
	// the service, such as due to maintenance being performed on the resource
	// block.
	UnavailableCompositionState CompositionState = "Unavailable"
)

type PoolType string

const (
	// FreePoolType This resource block is in the free pool and is not contributing
	// to any composed resources.
	FreePoolType PoolType = "Free"
	// ActivePoolType This resource block is in the active pool and is contributing
	// to at least one composed resource as a result of a composition request.
	ActivePoolType PoolType = "Active"
	// UnassignedPoolType This resource block is not assigned to any pools.
	UnassignedPoolType PoolType = "Unassigned"
)

type ResourceBlockType string

const (
	// ComputeResourceBlockType This resource block contains resources of type
	// 'Processor' and 'Memory' in a manner that creates a compute complex.
	ComputeResourceBlockType ResourceBlockType = "Compute"
	// ProcessorResourceBlockType This resource block contains resources of type
	// 'Processor'.
	ProcessorResourceBlockType ResourceBlockType = "Processor"
	// MemoryResourceBlockType This resource block contains resources of type
	// 'Memory'.
	MemoryResourceBlockType ResourceBlockType = "Memory"
	// NetworkResourceBlockType This resource block contains network resources,
	// such as resources of type 'EthernetInterface' and 'NetworkInterface'.
	NetworkResourceBlockType ResourceBlockType = "Network"
	// StorageResourceBlockType This resource block contains storage resources,
	// such as resources of type 'Storage' and 'SimpleStorage'.
	StorageResourceBlockType ResourceBlockType = "Storage"
	// ComputerSystemResourceBlockType This resource block contains resources of
	// type 'ComputerSystem'.
	ComputerSystemResourceBlockType ResourceBlockType = "ComputerSystem"
	// ExpansionResourceBlockType This resource block is capable of changing over
	// time based on its configuration. Different types of devices within this
	// resource block can be added and removed over time.
	ExpansionResourceBlockType ResourceBlockType = "Expansion"
	// IndependentResourceResourceBlockType This resource block is capable of being
	// consumed as a standalone component. This resource block can represent things
	// such as a software platform on one or more computer systems or an appliance
	// that provides composable resources and other services and can be managed
	// independently of the Redfish service.
	IndependentResourceResourceBlockType ResourceBlockType = "IndependentResource"
)

// ResourceBlock shall represent a resource block for a Redfish implementation.
type ResourceBlock struct {
	Entity
	// Client shall contain the client to which this resource block is assigned.
	//
	// Version added: v1.4.0
	Client string
	// CompositionStatus shall contain composition status information about this
	// resource block.
	CompositionStatus CompositionStatus
	// ComputerSystemsCount
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// Drives shall contain an array of links to resources of type 'Drive' that
	// this resource block contains.
	//
	// Version added: v1.3.0
	drives []string
	// DrivesCount
	DrivesCount int `json:"Drives@odata.count"`
	// EthernetInterfaces shall contain an array of links to resources of type
	// 'EthernetInterface' that this resource block contains.
	ethernetInterfaces []string
	// EthernetInterfacesCount
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Memory shall contain an array of links to resources of type 'Memory' that
	// this resource block contains.
	memory []string
	// MemoryCount
	MemoryCount int `json:"Memory@odata.count"`
	// NetworkInterfaces shall contain an array of links to resources of type
	// 'NetworkInterface' that this resource block contains.
	networkInterfaces []string
	// NetworkInterfacesCount
	NetworkInterfacesCount int `json:"NetworkInterfaces@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Pool shall contain the pool to which this resource block belongs. If this
	// resource block is not assigned to a client, this property shall contain the
	// value 'Unassigned'. If this resource block is assigned to a client, this
	// property shall not contain the value 'Unassigned'.
	//
	// Version added: v1.4.0
	Pool PoolType
	// Processors shall contain an array of links to resources of type 'Processor'
	// that this resource block contains.
	processors []string
	// ProcessorsCount
	ProcessorsCount int `json:"Processors@odata.count"`
	// ResourceBlockType shall contain an array of enumerated values that describe
	// the type of resources available.
	ResourceBlockType []ResourceBlockType
	// SimpleStorage shall contain an array of links to resources of type
	// 'SimpleStorage' that this resource block contains.
	simpleStorage []string
	// SimpleStorageCount
	SimpleStorageCount int `json:"SimpleStorage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Storage shall contain an array of links to resources of type 'Storage' that
	// this resource block contains.
	storage []string
	// StorageCount
	StorageCount int `json:"Storage@odata.count"`
	// chassis are the URIs for Chassis.
	chassis []string
	// computerSystems are the URIs for ComputerSystems.
	computerSystems []string
	// consumingResourceBlocks are the URIs for ConsumingResourceBlocks.
	consumingResourceBlocks []string
	// supplyingResourceBlocks are the URIs for SupplyingResourceBlocks.
	supplyingResourceBlocks []string
	// zones are the URIs for Zones.
	zones []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ResourceBlock object from the raw JSON.
func (r *ResourceBlock) UnmarshalJSON(b []byte) error {
	type temp ResourceBlock
	type rLinks struct {
		Chassis                 Links `json:"Chassis"`
		ComputerSystems         Links `json:"ComputerSystems"`
		ConsumingResourceBlocks Links `json:"ConsumingResourceBlocks"`
		SupplyingResourceBlocks Links `json:"SupplyingResourceBlocks"`
		Zones                   Links `json:"Zones"`
	}
	var tmp struct {
		temp
		Links              rLinks
		Drives             Links `json:"Drives"`
		EthernetInterfaces Links `json:"EthernetInterfaces"`
		Memory             Links `json:"Memory"`
		NetworkInterfaces  Links `json:"NetworkInterfaces"`
		Processors         Links `json:"Processors"`
		SimpleStorage      Links `json:"SimpleStorage"`
		Storage            Links `json:"Storage"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = ResourceBlock(tmp.temp)

	// Extract the links to other entities for later
	r.chassis = tmp.Links.Chassis.ToStrings()
	r.computerSystems = tmp.Links.ComputerSystems.ToStrings()
	r.consumingResourceBlocks = tmp.Links.ConsumingResourceBlocks.ToStrings()
	r.supplyingResourceBlocks = tmp.Links.SupplyingResourceBlocks.ToStrings()
	r.zones = tmp.Links.Zones.ToStrings()
	r.drives = tmp.Drives.ToStrings()
	r.ethernetInterfaces = tmp.EthernetInterfaces.ToStrings()
	r.memory = tmp.Memory.ToStrings()
	r.networkInterfaces = tmp.NetworkInterfaces.ToStrings()
	r.processors = tmp.Processors.ToStrings()
	r.simpleStorage = tmp.SimpleStorage.ToStrings()
	r.storage = tmp.Storage.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *ResourceBlock) Update() error {
	readWriteFields := []string{
		"Client",
		"Pool",
	}

	return r.UpdateFromRawData(r, r.RawData, readWriteFields)
}

// GetResourceBlock will get a ResourceBlock instance from the service.
func GetResourceBlock(c Client, uri string) (*ResourceBlock, error) {
	return GetObject[ResourceBlock](c, uri)
}

// ListReferencedResourceBlocks gets the collection of ResourceBlock from
// a provided reference.
func ListReferencedResourceBlocks(c Client, link string) ([]*ResourceBlock, error) {
	return GetCollectionObjects[ResourceBlock](c, link)
}

// Chassis gets the Chassis linked resources.
func (r *ResourceBlock) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](r.client, r.chassis)
}

// ComputerSystems gets the ComputerSystems linked resources.
func (r *ResourceBlock) ComputerSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](r.client, r.computerSystems)
}

// ConsumingResourceBlocks gets the ConsumingResourceBlocks linked resources.
func (r *ResourceBlock) ConsumingResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](r.client, r.consumingResourceBlocks)
}

// SupplyingResourceBlocks gets the SupplyingResourceBlocks linked resources.
func (r *ResourceBlock) SupplyingResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](r.client, r.supplyingResourceBlocks)
}

// Zones gets the Zones linked resources.
func (r *ResourceBlock) Zones() ([]*Zone, error) {
	return GetObjects[Zone](r.client, r.zones)
}

// Drives gets the Drives linked resources.
func (r *ResourceBlock) Drives() ([]*Drive, error) {
	return GetObjects[Drive](r.client, r.drives)
}

// EthernetInterfaces gets the EthernetInterfaces linked resources.
func (r *ResourceBlock) EthernetInterfaces() ([]*EthernetInterface, error) {
	return GetObjects[EthernetInterface](r.client, r.ethernetInterfaces)
}

// Memory gets the Memory linked resources.
func (r *ResourceBlock) Memory() ([]*Memory, error) {
	return GetObjects[Memory](r.client, r.memory)
}

// NetworkInterfaces gets the NetworkInterfaces linked resources.
func (r *ResourceBlock) NetworkInterfaces() ([]*NetworkInterface, error) {
	return GetObjects[NetworkInterface](r.client, r.networkInterfaces)
}

// Processors gets the Processors linked resources.
func (r *ResourceBlock) Processors() ([]*Processor, error) {
	return GetObjects[Processor](r.client, r.processors)
}

// SimpleStorage gets the SimpleStorage linked resources.
func (r *ResourceBlock) SimpleStorage() ([]*SimpleStorage, error) {
	return GetObjects[SimpleStorage](r.client, r.simpleStorage)
}

// Storage gets the Storage linked resources.
func (r *ResourceBlock) Storage() ([]*Storage, error) {
	return GetObjects[Storage](r.client, r.storage)
}

// CompositionStatus shall contain properties that describe the high level
// composition status of the resource block.
type CompositionStatus struct {
	// CompositionState shall contain an enumerated value that describes the
	// composition state of the resource block.
	CompositionState CompositionState
	// MaxCompositions shall contain a number indicating the maximum number of
	// compositions in which this resource block can participate simultaneously.
	// Services can have additional constraints that prevent this value from being
	// achieved, such as due to system topology and current composed resource
	// utilization. If 'SharingCapable' is 'false', this value shall be set to '1'.
	// The service shall support this property if SharingCapable supported.
	//
	// Version added: v1.1.0
	MaxCompositions *uint `json:",omitempty"`
	// NumberOfCompositions shall contain the number of compositions in which this
	// resource block is currently participating.
	//
	// Version added: v1.1.0
	NumberOfCompositions *uint `json:",omitempty"`
	// Reserved shall indicate whether any client has reserved the resource block.
	// A client sets this property after the resource block is identified as
	// composed. It shall provide a way for multiple clients to negotiate the
	// ownership of the resource block.
	Reserved bool
	// SharingCapable shall indicate whether this resource block can participate in
	// multiple compositions simultaneously. If this property is not provided, it
	// shall be assumed that this resource block is not capable of being shared.
	//
	// Version added: v1.1.0
	SharingCapable bool
	// SharingEnabled shall indicate whether this resource block can participate in
	// multiple compositions simultaneously. The service shall reject modifications
	// of this property with the HTTP '400 Bad Request' status code if this
	// resource block is already being used as part of a composed resource. If
	// 'false', the service shall not use the 'ComposedAndAvailable' state for this
	// resource block.
	//
	// Version added: v1.1.0
	SharingEnabled bool
}

// ResourceBlockLimits shall specify the allowable quantities of types of
// resource blocks for a given composition request.
type ResourceBlockLimits struct {
	// MaxCompute shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Compute' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxCompute *uint `json:",omitempty"`
	// MaxComputerSystem shall contain an integer that specifies the maximum number
	// of resource blocks of type 'ComputerSystem' allowed for the composition
	// request.
	//
	// Version added: v1.3.0
	MaxComputerSystem *uint `json:",omitempty"`
	// MaxExpansion shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Expansion' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxExpansion *uint `json:",omitempty"`
	// MaxMemory shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Memory' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxMemory *uint `json:",omitempty"`
	// MaxNetwork shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Network' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxNetwork *uint `json:",omitempty"`
	// MaxProcessor shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Processor' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxProcessor *uint `json:",omitempty"`
	// MaxStorage shall contain an integer that specifies the maximum number of
	// resource blocks of type 'Storage' allowed for the composition request.
	//
	// Version added: v1.3.0
	MaxStorage *uint `json:",omitempty"`
	// MinCompute shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Compute' required for the composition request.
	//
	// Version added: v1.3.0
	MinCompute *uint `json:",omitempty"`
	// MinComputerSystem shall contain an integer that specifies the minimum number
	// of resource blocks of type 'ComputerSystem' required for the composition
	// request.
	//
	// Version added: v1.3.0
	MinComputerSystem *uint `json:",omitempty"`
	// MinExpansion shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Expansion' required for the composition request.
	//
	// Version added: v1.3.0
	MinExpansion *uint `json:",omitempty"`
	// MinMemory shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Memory' required for the composition request.
	//
	// Version added: v1.3.0
	MinMemory *uint `json:",omitempty"`
	// MinNetwork shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Network' required for the composition request.
	//
	// Version added: v1.3.0
	MinNetwork *uint `json:",omitempty"`
	// MinProcessor shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Processor' required for the composition request.
	//
	// Version added: v1.3.0
	MinProcessor *uint `json:",omitempty"`
	// MinStorage shall contain an integer that specifies the minimum number of
	// resource blocks of type 'Storage' required for the composition request.
	//
	// Version added: v1.3.0
	MinStorage *uint `json:",omitempty"`
}
