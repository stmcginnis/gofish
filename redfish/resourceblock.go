//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type CompositionState string

const (
	// ComposingCompositionState Intermediate state indicating composition is in progress.
	ComposingCompositionState CompositionState = "Composing"
	// ComposedAndAvailableCompositionState The resource block is currently participating in one or more compositions,
	// and is available to use in more compositions.
	ComposedAndAvailableCompositionState CompositionState = "ComposedAndAvailable"
	// ComposedCompositionState Final successful state of a resource block that has participated in composition.
	ComposedCompositionState CompositionState = "Composed"
	// UnusedCompositionState The resource block is free and can participate in composition.
	UnusedCompositionState CompositionState = "Unused"
	// FailedCompositionState The final composition resulted in failure and manual intervention might be required to
	// fix it.
	FailedCompositionState CompositionState = "Failed"
	// UnavailableCompositionState The resource block has been made unavailable by the service, such as due to
	// maintenance being performed on the resource block.
	UnavailableCompositionState CompositionState = "Unavailable"
)

type PoolType string

const (
	// FreePoolType This resource block is in the free pool and is not contributing to any composed resources.
	FreePoolType PoolType = "Free"
	// ActivePoolType This resource block is in the active pool and is contributing to at least one composed resource
	// as a result of a composition request.
	ActivePoolType PoolType = "Active"
	// UnassignedPoolType This resource block is not assigned to any pools.
	UnassignedPoolType PoolType = "Unassigned"
)

type ResourceBlockType string

const (
	// ComputeResourceBlockType This resource block contains resources of type 'Processor' and 'Memory' in a manner
	// that creates a compute complex.
	ComputeResourceBlockType ResourceBlockType = "Compute"
	// ProcessorResourceBlockType This resource block contains resources of type 'Processor'.
	ProcessorResourceBlockType ResourceBlockType = "Processor"
	// MemoryResourceBlockType This resource block contains resources of type 'Memory'.
	MemoryResourceBlockType ResourceBlockType = "Memory"
	// NetworkResourceBlockType This resource block contains network resources, such as resources of type
	// 'EthernetInterface' and 'NetworkInterface'.
	NetworkResourceBlockType ResourceBlockType = "Network"
	// StorageResourceBlockType This resource block contains storage resources, such as resources of type 'Storage' and
	// 'SimpleStorage'.
	StorageResourceBlockType ResourceBlockType = "Storage"
	// ComputerSystemResourceBlockType This resource block contains resources of type 'ComputerSystem'.
	ComputerSystemResourceBlockType ResourceBlockType = "ComputerSystem"
	// ExpansionResourceBlockType This resource block is capable of changing over time based on its configuration.
	// Different types of devices within this resource block can be added and removed over time.
	ExpansionResourceBlockType ResourceBlockType = "Expansion"
	// IndependentResourceResourceBlockType This resource block is capable of being consumed as a standalone component.
	// This resource block can represent things such as a software platform on one or more computer systems or an
	// appliance that provides composable resources and other services and can be managed independently of the Redfish
	// service.
	IndependentResourceResourceBlockType ResourceBlockType = "IndependentResource"
)

// CompositionStatus shall contain properties that describe the high level composition status of the resource
// block.
type CompositionStatus struct {
	// CompositionState shall contain an enumerated value that describes the composition state of the resource block.
	CompositionState CompositionState
	// MaxCompositions shall contain a number indicating the maximum number of compositions in which this resource
	// block can participate simultaneously. Services can have additional constraints that prevent this value from
	// being achieved, such as due to system topology and current composed resource utilization. If SharingCapable is
	// 'false', this value shall be set to '1'. The service shall support this property if SharingCapable supported.
	MaxCompositions int
	// NumberOfCompositions shall contain the number of compositions in which this resource block is currently
	// participating.
	NumberOfCompositions int
	// Reserved shall indicate whether any client has reserved the resource block. A client sets this property after
	// the resource block is identified as composed. It shall provide a way for multiple clients to negotiate the
	// ownership of the resource block.
	Reserved bool
	// SharingCapable shall indicate whether this resource block can participate in multiple compositions
	// simultaneously. If this property is not provided, it shall be assumed that this resource block is not capable of
	// being shared.
	SharingCapable bool
	// SharingEnabled shall indicate whether this resource block can participate in multiple compositions
	// simultaneously. The service shall reject modifications of this property with HTTP 400 Bad Request if this
	// resource block is already being used as part of a composed resource. If 'false', the service shall not use the
	// 'ComposedAndAvailable' state for this resource block.
	SharingEnabled bool
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type rbLinks struct {
	// Chassis shall contain an array of links to resources of type Chassis that represent the physical containers
	// associated with this resource block.
	Chassis common.Links
	// Chassis@odata.count
	ChassisCount int `json:"Chassis@odata.count"`
	// ComputerSystems shall contain an array of links to resources of type ComputerSystem that represent the computer
	// systems composed from this resource block.
	ComputerSystems common.Links
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// ConsumingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that depend on this resource block as a component.
	ConsumingResourceBlocks common.Links
	// ConsumingResourceBlocks@odata.count
	ConsumingResourceBlocksCount int `json:"ConsumingResourceBlocks@odata.count"`
	// SupplyingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that this resource block depends on as components.
	SupplyingResourceBlocks common.Links
	// SupplyingResourceBlocks@odata.count
	SupplyingResourceBlocksCount int `json:"SupplyingResourceBlocks@odata.count"`
	// Zones shall contain an array of links to resources of type Zone that represent the binding constraints
	// associated with this resource block.
	Zones common.Links
	// Zones@odata.count
	ZonesCount int `json:"Zones@odata.count"`
}

// ResourceBlock shall represent a resource block for a Redfish implementation.
type ResourceBlock struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Client shall contain the client to which this resource block is assigned.
	Client string
	// CompositionStatus shall contain composition status information about this resource block.
	CompositionStatus CompositionStatus
	// ComputerSystems shall contain an array of links to resources of type ComputerSystem that this resource block
	// contains.
	computerSystems []string
	// ComputerSystemsCount is the number of ComputerSystems that this resource block contains.
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// Description provides a description of this resource.
	Description string
	// Drives shall contain an array of links to resources of type Drive that this resource block contains.
	drives []string
	// DrivesCount is the number of Drives that this resource block contains.
	DrivesCount int `json:"Drives@odata.count"`
	// EthernetInterfaces shall contain an array of links to resources of type EthernetInterface that this resource
	// block contains.
	ethernetInterfaces []string
	// EthernetInterfacesCount is the number of EthernetInterfaces this resource block contains.
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Memory shall contain an array of links to resources of type Memory that this resource block contains.
	memory []string
	// MemoryCount is the number of Memory objects this resource block contains.
	MemoryCount int `json:"Memory@odata.count"`
	// NetworkInterfaces shall contain an array of links to resources of type NetworkInterface that this resource block
	// contains.
	networkInterfaces []string
	// NetworkInterfacesCount is the number of NetworkInterfaces this resource block contains.
	NetworkInterfacesCount int `json:"NetworkInterfaces@odata.count"`
	// Pool shall contain the pool to which this resource block belongs. If this resource block is not assigned to a
	// client, this property shall contain the value 'Unassigned'. If this resource block is assigned to a client, this
	// property shall not contain the value 'Unassigned'.
	Pool PoolType
	// Processors shall contain an array of links to resources of type Processor that this resource block contains.
	processors []string
	// ProcessorsCount is the number of processors this resource block contains.
	ProcessorsCount int `json:"Processors@odata.count"`
	// ResourceBlockType shall contain an array of enumerated values that describe the type of resources available.
	ResourceBlockType []ResourceBlockType
	// SimpleStorage shall contain an array of links to resources of type SimpleStorage that this resource block
	// contains.
	simpleStorage []string
	// SimpleStorageCount is the number of SimpleStorage instances this resource block contains.
	SimpleStorageCount int `json:"SimpleStorage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Storage shall contain an array of links to resources of type Storage that this resource block contains.
	storage []string
	// StorageCount is the number of Storage instances this resource block contains.
	StorageCount int `json:"Storage@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// Chassis is the set of links to associated chassis objects.
	chassis []string
	// ChassisCount is the number of associated Chassis objects.
	ChassisCount int
	// ConsumingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that depend on this resource block as a component.
	consumingResourceBlocks []string
	// ConsumingResourceBlocksCount is the number of ResourceBlock objects that depend on this resource block.
	ConsumingResourceBlocksCount int
	// SupplyingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that this resource block depends on as components.
	supplyingResourceBlocks []string
	// SupplyingResourceBlocksCount is the number of ResourceBlocks that this resource block depends on.
	SupplyingResourceBlocksCount int
	// Zones shall contain an array of links to resources of type Zone that represent the binding constraints
	// associated with this resource block.
	zones []string
	// ZonesCount is the number of Zone objects associated with this resource block.
	ZonesCount int
}

// UnmarshalJSON unmarshals a ResourceBlock object from the raw JSON.
func (resourceblock *ResourceBlock) UnmarshalJSON(b []byte) error {
	type temp ResourceBlock
	var t struct {
		temp
		ComputerSystems    common.Links
		Drives             common.Links
		EthernetInterfaces common.Links
		Memory             common.Links
		NetworkInterfaces  common.Links
		Processors         common.Links
		SimpleStorage      common.Links
		Storage            common.Links
		Links              rbLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*resourceblock = ResourceBlock(t.temp)

	// Extract the links to other entities for later
	resourceblock.computerSystems = t.ComputerSystems.ToStrings()
	resourceblock.drives = t.Drives.ToStrings()
	resourceblock.ethernetInterfaces = t.EthernetInterfaces.ToStrings()
	resourceblock.memory = t.Memory.ToStrings()
	resourceblock.networkInterfaces = t.NetworkInterfaces.ToStrings()
	resourceblock.processors = t.Processors.ToStrings()
	resourceblock.simpleStorage = t.SimpleStorage.ToStrings()
	resourceblock.storage = t.Storage.ToStrings()
	resourceblock.chassis = t.Links.Chassis.ToStrings()
	resourceblock.ChassisCount = t.Links.ChassisCount
	resourceblock.consumingResourceBlocks = t.Links.ConsumingResourceBlocks.ToStrings()
	resourceblock.ConsumingResourceBlocksCount = t.Links.ConsumingResourceBlocksCount
	resourceblock.supplyingResourceBlocks = t.Links.SupplyingResourceBlocks.ToStrings()
	resourceblock.SupplyingResourceBlocksCount = t.Links.SupplyingResourceBlocksCount
	resourceblock.zones = t.Links.Zones.ToStrings()
	resourceblock.ZonesCount = t.Links.ZonesCount

	if resourceblock.ComputerSystemsCount == 0 {
		// Via Links instead of directly in object
		resourceblock.computerSystems = t.Links.ComputerSystems.ToStrings()
		resourceblock.ComputerSystemsCount = t.Links.ComputerSystemsCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	resourceblock.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (resourceblock *ResourceBlock) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ResourceBlock)
	original.UnmarshalJSON(resourceblock.rawData)

	readWriteFields := []string{
		"Client",
		"Pool",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(resourceblock).Elem()

	return resourceblock.Entity.Update(originalElement, currentElement, readWriteFields)
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

// ResourceBlockLimits shall specify the allowable quantities of types of resource blocks for a given composition
// request.
type ResourceBlockLimits struct {
	// MaxCompute shall contain an integer that specifies the maximum number of resource blocks of type 'Compute'
	// allowed for the composition request.
	MaxCompute int
	// MaxComputerSystem shall contain an integer that specifies the maximum number of resource blocks of type
	// 'ComputerSystem' allowed for the composition request.
	MaxComputerSystem int
	// MaxExpansion shall contain an integer that specifies the maximum number of resource blocks of type 'Expansion'
	// allowed for the composition request.
	MaxExpansion int
	// MaxMemory shall contain an integer that specifies the maximum number of resource blocks of type 'Memory' allowed
	// for the composition request.
	MaxMemory int
	// MaxNetwork shall contain an integer that specifies the maximum number of resource blocks of type 'Network'
	// allowed for the composition request.
	MaxNetwork int
	// MaxProcessor shall contain an integer that specifies the maximum number of resource blocks of type 'Processor'
	// allowed for the composition request.
	MaxProcessor int
	// MaxStorage shall contain an integer that specifies the maximum number of resource blocks of type 'Storage'
	// allowed for the composition request.
	MaxStorage int
	// MinCompute shall contain an integer that specifies the minimum number of resource blocks of type 'Compute'
	// required for the composition request.
	MinCompute int
	// MinComputerSystem shall contain an integer that specifies the minimum number of resource blocks of type
	// 'ComputerSystem' required for the composition request.
	MinComputerSystem int
	// MinExpansion shall contain an integer that specifies the minimum number of resource blocks of type 'Expansion'
	// required for the composition request.
	MinExpansion int
	// MinMemory shall contain an integer that specifies the minimum number of resource blocks of type 'Memory'
	// required for the composition request.
	MinMemory int
	// MinNetwork shall contain an integer that specifies the minimum number of resource blocks of type 'Network'
	// required for the composition request.
	MinNetwork int
	// MinProcessor shall contain an integer that specifies the minimum number of resource blocks of type 'Processor'
	// required for the composition request.
	MinProcessor int
	// MinStorage shall contain an integer that specifies the minimum number of resource blocks of type 'Storage'
	// required for the composition request.
	MinStorage int
}
