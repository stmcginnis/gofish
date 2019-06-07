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

	"github.com/stmcginnis/gofish/school/redfish"

	"github.com/stmcginnis/gofish/school/common"
)

// StoragePool is a container of data storage capable of providing
// capacity conforming to one of its supported classes of service. The
// storage pool does not support IO to its data storage.
type StoragePool struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
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
	// DefaultClassOfService is used.
	defaultClassOfService string
	// Description provides a description of this resource.
	Description string
	// IOStatistics is the value shall represent IO statistics for this
	// StoragePool.
	ioStatistics string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// LowSpaceWarningThresholdPercents is each time the following value is
	// less than one of the values in the array the
	// LOW_SPACE_THRESHOLD_WARNING event shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []string
	// MaxBlockSizeBytes if present, the value is the maximum block size
	// of an allocated resource. If the block size is unknown or if a block
	// concept is not valid (for example, with Memory), this property shall
	// be NULL.
	MaxBlockSizeBytes int
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
	//this StoragePool.
	dedicatedSpareDrives []string
	// DedicatedSpareDrivesCount is the number of drives.
	DedicatedSpareDrivesCount int
	// SpareResourceSets shall contain resources that may be utilized to replace
	// the capacity provided by a failed resource having a compatible type.
	spareResourceSets []string
	// SpareResourceSetsCount is the number of spare resource sets.
	SpareResourceSetsCount int
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
		IOStatistics          common.Link
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
	storagepool.ioStatistics = string(t.IOStatistics)
	storagepool.allocatedPools = string(t.AllocatedPools)
	storagepool.allocatedVolumes = string(t.AllocatedVolumes)
	storagepool.capacitySources = t.CapacitySource.ToStrings()
	storagepool.classesOfService = string(t.ClassesOfService)
	storagepool.defaultClassOfService = string(t.DefaultClassOfService)

	return nil
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

// IOStatistics gets IO statistics for this storage pool.
func (storagepool *StoragePool) IOStatistics() (*IOStatistics, error) {
	if storagepool.ioStatistics == "" {
		return nil, nil
	}
	return GetIOStatistics(storagepool.Client, storagepool.ioStatistics)
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
		cap, err := GetCapacitySource(storagepool.Client, capLink)
		if err != nil {
			return result, nil
		}
		result = append(result, cap)
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
