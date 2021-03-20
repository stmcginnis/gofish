//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/redfish"

	"github.com/stmcginnis/gofish/common"
)

// StoragePool is a container of data storage capable of providing
// capacity conforming to one of its supported classes of service. The
// storage pool does not support IO to its data storage.
type StoragePool struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllocatedPools shall contain a reference
	// to the collection of storage pools allocated from this storage pool.
	allocatedPools string
	// AllocatedVolumes shall contain a
	// reference to the collection of volumes allocated from this storage
	// pool.
	allocatedVolumes string
	// Capacity shall provide an information
	// about the actual utilization of the capacity within this storage pool.
	Capacity Capacity
	// CapacitySources is fully or partially consumed storage from a source
	// resource. Each entry shall provide capacity allocation data from a
	// named source resource.
	capacitySources []string
	// CapacitySourcesCount is the number of capacity sources.
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// ClassesOfService shall contain references to all classes of service
	// supported by this storage pool. Capacity allocated from this storage pool
	// shall conform to one of the referenced classes of service.
	classesOfService string
	// Compressed shall contain a boolean indicator if the StoragePool is
	// currently utilizing compression or not.
	Compressed bool
	// Deduplicted shall contain a boolean indicator if the StoragePool is
	// currently utilizing deduplication or not.
	Deduplicated bool
	// DefaultClassOfService is used.
	defaultClassOfService string
	// Description provides a description of this resource.
	Description string
	// Encrypted shall contain a boolean indicator if the
	// StoragePool is currently utilizing encryption or not.
	Encrypted bool
	// IOStatistics is the value shall represent IO statistics for this
	// StoragePool.
	IOStatistics IOStatistics
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// LowSpaceWarningThresholdPercents is each time the following value is
	// less than one of the values in the array the
	// LOW_SPACE_THRESHOLD_WARNING event shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []int
	// MaxBlockSizeBytes if present, the value is the maximum block size
	// of an allocated resource. If the block size is unknown or if a block
	// concept is not valid (for example, with Memory), this property shall
	// be NULL.
	MaxBlockSizeBytes int64
	// RecoverableCapacitySourceCount is the value of the number of available
	// capacity source resources currently available in the event that an
	// equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent if present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// Status is the storage pool status.
	Status common.Status
	// DedicatedSpareDrives shall be a reference to the resources that this
	// StoragePool is associated with and shall reference resources of type
	// Drive. This property shall only contain references to Drive entities
	// which are currently assigned as a dedicated spare and are able to support
	// this StoragePool.
	dedicatedSpareDrives []string
	// DedicatedSpareDrivesCount is the number of drives.
	DedicatedSpareDrivesCount int
	// SpareResourceSets shall contain resources that may be utilized to replace
	// the capacity provided by a failed resource having a compatible type.
	spareResourceSets []string
	// SpareResourceSetsCount is the number of spare resource sets.
	SpareResourceSetsCount int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StoragePool object from the raw JSON.
func (storagepool *StoragePool) UnmarshalJSON(b []byte) error {
	type temp StoragePool
	type links struct {
		DedicatedSpareDrives      common.Links
		DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
		SpareResourceSets         common.Links
		SpareResourceSetsCount    int `json:"SpareResourceSets@odata.count"`
	}
	var t struct {
		temp
		Links                 links
		AllocatedPools        common.Link
		AllocatedVolumes      common.Link
		CapacitySource        common.Links
		ClassesOfService      common.Link
		DefaultClassOfService common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagepool = StoragePool(t.temp)

	// Extract the links to other entities for later
	storagepool.dedicatedSpareDrives = t.Links.DedicatedSpareDrives.ToStrings()
	storagepool.DedicatedSpareDrivesCount = t.Links.DedicatedSpareDrivesCount
	storagepool.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	storagepool.SpareResourceSetsCount = t.Links.SpareResourceSetsCount
	storagepool.allocatedPools = string(t.AllocatedPools)
	storagepool.allocatedVolumes = string(t.AllocatedVolumes)
	storagepool.capacitySources = t.CapacitySource.ToStrings()
	storagepool.classesOfService = string(t.ClassesOfService)
	storagepool.defaultClassOfService = string(t.DefaultClassOfService)

	// This is a read/write object, so we need to save the raw object data for later
	storagepool.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storagepool *StoragePool) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StoragePool)
	err := original.UnmarshalJSON(storagepool.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"CapacitySources",
		"ClassesOfService",
		"Compressed",
		"Deduplicated",
		"DefaultClassOfService",
		"Encrypted",
		"LowSpaceWarningThresholdPercents",
		"RecoverableCapacitySourceCount",
		"SupportedProvisioningPolicies",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagepool).Elem()

	return storagepool.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetStoragePool will get a StoragePool instance from the service.
func GetStoragePool(c common.Client, uri string) (*StoragePool, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storagepool StoragePool
	err = json.NewDecoder(resp.Body).Decode(&storagepool)
	if err != nil {
		return nil, err
	}

	storagepool.SetClient(c)
	return &storagepool, nil
}

// ListReferencedStoragePools gets the collection of StoragePool from
// a provided reference.
func ListReferencedStoragePools(c common.Client, link string) ([]*StoragePool, error) {
	var result []*StoragePool
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storagepoolLink := range links.ItemLinks {
		storagepool, err := GetStoragePool(c, storagepoolLink)
		if err != nil {
			return result, err
		}
		result = append(result, storagepool)
	}

	return result, nil
}

// DedicatedSpareDrives gets the Drive entities which are currently assigned as
// a dedicated spare and are able to support this StoragePool.
func (storagepool *StoragePool) DedicatedSpareDrives() ([]*redfish.Drive, error) {
	var result []*redfish.Drive
	for _, driveLink := range storagepool.dedicatedSpareDrives {
		drive, err := redfish.GetDrive(storagepool.Client, driveLink)
		if err != nil {
			return result, nil
		}
		result = append(result, drive)
	}
	return result, nil
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storagepool *StoragePool) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet
	for _, srsLink := range storagepool.spareResourceSets {
		srs, err := GetSpareResourceSet(storagepool.Client, srsLink)
		if err != nil {
			return result, nil
		}
		result = append(result, srs)
	}
	return result, nil
}

// AllocatedPools gets the storage pools allocated from this storage pool.
func (storagepool *StoragePool) AllocatedPools() ([]*StoragePool, error) {
	return ListReferencedStoragePools(storagepool.Client, storagepool.allocatedPools)
}

// AllocatedVolumes gets the volumes allocated from this storage pool.
func (storagepool *StoragePool) AllocatedVolumes() ([]*Volume, error) {
	return ListReferencedVolumes(storagepool.Client, storagepool.allocatedVolumes)
}

// CapacitySources gets space allocations to this pool.
func (storagepool *StoragePool) CapacitySources() ([]*CapacitySource, error) {
	var result []*CapacitySource
	for _, capLink := range storagepool.capacitySources {
		capacity, err := GetCapacitySource(storagepool.Client, capLink)
		if err != nil {
			return result, nil
		}
		result = append(result, capacity)
	}
	return result, nil
}

// ClassesOfService gets references to all classes of service supported by this
// storage pool. Capacity allocated from this storage pool shall conform to one
// of the referenced classes of service.
func (storagepool *StoragePool) ClassesOfService() ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storagepool.Client, storagepool.classesOfService)
}

// DefaultClassOfService gets the default ClassOfService for this pool.
func (storagepool *StoragePool) DefaultClassOfService() (*ClassOfService, error) {
	if storagepool.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storagepool.Client, storagepool.defaultClassOfService)
}
