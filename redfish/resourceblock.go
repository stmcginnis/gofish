//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #ResourceBlock.v1_4_3.ResourceBlock

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	common.Entity
	// Client shall contain the client to which this resource block is assigned.
	//
	// Version added: v1.4.0
	Client string
	// CompositionStatus shall contain composition status information about this
	// resource block.
	CompositionStatus CompositionStatus
	// ComputerSystems shall contain an array of links to resources of type
	// 'ComputerSystem' that this resource block contains.
	computerSystems []string
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// Drives shall contain an array of links to resources of type 'Drive' that
	// this resource block contains.
	//
	// Version added: v1.3.0
	drives []string
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// EthernetInterfaces shall contain an array of links to resources of type
	// 'EthernetInterface' that this resource block contains.
	ethernetInterfaces []string
	// EthernetInterfaces@odata.count
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Memory shall contain an array of links to resources of type 'Memory' that
	// this resource block contains.
	memory []string
	// Memory@odata.count
	MemoryCount int `json:"Memory@odata.count"`
	// NetworkInterfaces shall contain an array of links to resources of type
	// 'NetworkInterface' that this resource block contains.
	networkInterfaces []string
	// NetworkInterfaces@odata.count
	NetworkInterfacesCount int `json:"NetworkInterfaces@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
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
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
	// ResourceBlockType shall contain an array of enumerated values that describe
	// the type of resources available.
	ResourceBlockType []ResourceBlockType
	// SimpleStorage shall contain an array of links to resources of type
	// 'SimpleStorage' that this resource block contains.
	simpleStorage []string
	// SimpleStorage@odata.count
	SimpleStorageCount int `json:"SimpleStorage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Storage shall contain an array of links to resources of type 'Storage' that
	// this resource block contains.
	storage []string
	// Storage@odata.count
	StorageCount int `json:"Storage@odata.count"`
	// chassis are the URIs for Chassis.
	chassis      []string
	ChassisCount int
	// consumingResourceBlocks are the URIs for ConsumingResourceBlocks.
	consumingResourceBlocks      []string
	ConsumingResourceBlocksCount int
	// supplyingResourceBlocks are the URIs for SupplyingResourceBlocks.
	supplyingResourceBlocks      []string
	SupplyingResourceBlocksCount int
	// zones are the URIs for Zones.
	zones      []string
	ZonesCount int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ResourceBlock object from the raw JSON.
func (r *ResourceBlock) UnmarshalJSON(b []byte) error {
	type temp ResourceBlock
	type rLinks struct {
		Chassis                      common.Links `json:"Chassis"`
		ChassisCount                 int          `json:"Chassis@odata.count"`
		ComputerSystems              common.Links `json:"ComputerSystems"`
		ComputerSystemsCount         int          `json:"ComputerSystems@odata.count"`
		ConsumingResourceBlocks      common.Links `json:"ConsumingResourceBlocks"`
		ConsumingResourceBlocksCount int          `json:"ConsumingResourceBlocks@odata.count"`
		SupplyingResourceBlocks      common.Links `json:"SupplyingResourceBlocks"`
		SupplyingResourceBlocksCount int          `json:"SupplyingResourceBlocks@odata.count"`
		Zones                        common.Links `json:"Zones"`
		ZonesCount                   int          `json:"Zones@odata.count"`
	}
	var tmp struct {
		temp
		Links              rLinks
		ComputerSystems    common.Links
		Drives             common.Links
		EthernetInterfaces common.Links
		Memory             common.Links
		NetworkInterfaces  common.Links
		Processors         common.Links
		SimpleStorage      common.Links
		Storage            common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = ResourceBlock(tmp.temp)

	// Extract the links to other entities for later
	r.chassis = tmp.Links.Chassis.ToStrings()
	r.ChassisCount = tmp.Links.ChassisCount
	r.computerSystems = tmp.Links.ComputerSystems.ToStrings()
	r.consumingResourceBlocks = tmp.Links.ConsumingResourceBlocks.ToStrings()
	r.drives = tmp.Drives.ToStrings()
	r.ethernetInterfaces = tmp.EthernetInterfaces.ToStrings()
	r.memory = tmp.Memory.ToStrings()
	r.networkInterfaces = tmp.NetworkInterfaces.ToStrings()
	r.processors = tmp.Processors.ToStrings()
	r.simpleStorage = tmp.SimpleStorage.ToStrings()
	r.storage = tmp.Storage.ToStrings()
	r.supplyingResourceBlocks = tmp.Links.SupplyingResourceBlocks.ToStrings()
	r.zones = tmp.Links.Zones.ToStrings()

	if r.ComputerSystemsCount == 0 {
		// Via Links instead of directly in object
		r.computerSystems = tmp.Links.ComputerSystems.ToStrings()
		r.ComputerSystemsCount = tmp.Links.ComputerSystemsCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	r.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *ResourceBlock) Update() error {
	readWriteFields := []string{
		"Client",
		"CompositionStatus",
		"ComputerSystems@odata.count",
		"Drives@odata.count",
		"EthernetInterfaces@odata.count",
		"Memory@odata.count",
		"NetworkInterfaces@odata.count",
		"Pool",
		"Processors@odata.count",
		"SimpleStorage@odata.count",
		"Status",
		"Storage@odata.count",
	}

	return r.UpdateFromRawData(r, r.rawData, readWriteFields)
}

// GetResourceBlock will get a ResourceBlock instance from the service.
func GetResourceBlock(c common.Client, uri string) (*ResourceBlock, error) {
	return common.GetObject[ResourceBlock](c, uri)
}

// ListReferencedResourceBlocks gets the collection of ResourceBlock from
// a provided reference.
func ListReferencedResourceBlocks(c common.Client, link string) ([]*ResourceBlock, error) {
	return common.GetCollectionObjects[ResourceBlock](c, link)
}

// Chassis gets the Chassis linked resources.
func (r *ResourceBlock) Chassis(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, r.chassis)
}

// ComputerSystems gets the ComputerSystems linked resources.
func (r *ResourceBlock) ComputerSystems(client common.Client) ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](client, r.computerSystems)
}

// ConsumingResourceBlocks gets the ConsumingResourceBlocks linked resources.
func (r *ResourceBlock) ConsumingResourceBlocks(client common.Client) ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](client, r.consumingResourceBlocks)
}

// Drives gets the Drives linked resources.
func (r *ResourceBlock) Drives(client common.Client) ([]*Drive, error) {
	return common.GetObjects[Drive](client, r.drives)
}

// EthernetInterfaces gets the EthernetInterface linked resources.
func (r *ResourceBlock) EthernetInterfaces(client common.Client) ([]*EthernetInterface, error) {
	return common.GetObjects[EthernetInterface](client, r.ethernetInterfaces)
}

// Memory gets the Memory linked resources.
func (r *ResourceBlock) Memory(client common.Client) ([]*Memory, error) {
	return common.GetObjects[Memory](client, r.memory)
}

// NetworkInterfaces gets the NetworkInterface linked resources.
func (r *ResourceBlock) NetworkInterfaces(client common.Client) ([]*NetworkInterface, error) {
	return common.GetObjects[NetworkInterface](client, r.networkInterfaces)
}

// Processors gets the Processor linked resources.
func (r *ResourceBlock) Processors(client common.Client) ([]*Processor, error) {
	return common.GetObjects[Processor](client, r.processors)
}

// SupplyingResourceBlocks gets the SupplyingResourceBlocks linked resources.
func (r *ResourceBlock) SupplyingResourceBlocks(client common.Client) ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](client, r.supplyingResourceBlocks)
}

// Zones gets the Zones linked resources.
func (r *ResourceBlock) Zones(client common.Client) ([]*Zone, error) {
	return common.GetObjects[Zone](client, r.zones)
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
