//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
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
	storagepool.allocatedPools = t.AllocatedPools.String()
	storagepool.allocatedVolumes = t.AllocatedVolumes.String()
	storagepool.capacitySources = t.CapacitySource.ToStrings()
	storagepool.classesOfService = t.ClassesOfService.String()
	storagepool.defaultClassOfService = t.DefaultClassOfService.String()

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
	var storagePool StoragePool
	return &storagePool, storagePool.Get(c, uri, &storagePool)
}

// ListReferencedStoragePools gets the collection of StoragePool from
// a provided reference.
func ListReferencedStoragePools(c common.Client, link string) ([]*StoragePool, error) { //nolint:dupl
	var result []*StoragePool
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *StoragePool
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		storagepool, err := GetStoragePool(c, link)
		ch <- GetResult{Item: storagepool, Link: link, Error: err}
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

// DedicatedSpareDrives gets the Drive entities which are currently assigned as
// a dedicated spare and are able to support this StoragePool.
func (storagepool *StoragePool) DedicatedSpareDrives() ([]*redfish.Drive, error) {
	var result []*redfish.Drive

	collectionError := common.NewCollectionError()
	for _, driveLink := range storagepool.dedicatedSpareDrives {
		drive, err := redfish.GetDrive(storagepool.GetClient(), driveLink)
		if err != nil {
			collectionError.Failures[driveLink] = err
		} else {
			result = append(result, drive)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storagepool *StoragePool) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet

	collectionError := common.NewCollectionError()
	for _, srsLink := range storagepool.spareResourceSets {
		srs, err := GetSpareResourceSet(storagepool.GetClient(), srsLink)
		if err != nil {
			collectionError.Failures[srsLink] = err
		} else {
			result = append(result, srs)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// AllocatedPools gets the storage pools allocated from this storage pool.
func (storagepool *StoragePool) AllocatedPools() ([]*StoragePool, error) {
	return ListReferencedStoragePools(storagepool.GetClient(), storagepool.allocatedPools)
}

// AllocatedVolumes gets the volumes allocated from this storage pool.
func (storagepool *StoragePool) AllocatedVolumes() ([]*Volume, error) {
	return ListReferencedVolumes(storagepool.GetClient(), storagepool.allocatedVolumes)
}

// CapacitySources gets space allocations to this pool.
func (storagepool *StoragePool) CapacitySources() ([]*CapacitySource, error) {
	var result []*CapacitySource

	collectionError := common.NewCollectionError()
	for _, capLink := range storagepool.capacitySources {
		capacity, err := GetCapacitySource(storagepool.GetClient(), capLink)
		if err != nil {
			collectionError.Failures[capLink] = err
		} else {
			result = append(result, capacity)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ClassesOfService gets references to all classes of service supported by this
// storage pool. Capacity allocated from this storage pool shall conform to one
// of the referenced classes of service.
func (storagepool *StoragePool) ClassesOfService() ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storagepool.GetClient(), storagepool.classesOfService)
}

// DefaultClassOfService gets the default ClassOfService for this pool.
func (storagepool *StoragePool) DefaultClassOfService() (*ClassOfService, error) {
	if storagepool.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storagepool.GetClient(), storagepool.defaultClassOfService)
}
