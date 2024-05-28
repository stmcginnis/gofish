//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Device shall describe a storage device visible to SimpleStorage.
type Device struct {
	// CapacityBytes shall represent the size (in bytes) of the Storage Device.
	CapacityBytes int64
	// Manufacturer shall indicate the name of the manufacturer of this storage device.
	Manufacturer string
	// Model shall indicate the model information as provided by the manufacturer
	// of this storage device.
	Model string
	// Name shall be a user-friendly name for the device.
	Name string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// SimpleStorage is used to represent a storage controller and its
// directly-attached devices.
type SimpleStorage struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Devices shall contain a list of storage devices
	// associated with this resource.
	Devices []Device
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// UefiDevicePath is used to identify and locate the specific storage
	// controller.
	UefiDevicePath string

	// chassis shall be a reference to a resource of type Chassis that
	// represent the physical container associated with this Simple Storage.
	chassis string
	// Storage shall contain a link to a Resource of type Storage that represents the same storage subsystem as this
	// Resource.
	storage string
}

// UnmarshalJSON unmarshals a SimpleStorage object from the raw JSON.
func (simplestorage *SimpleStorage) UnmarshalJSON(b []byte) error {
	type temp SimpleStorage
	var t struct {
		temp
		Links struct {
			Chassis common.Link
			Storage common.Link
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*simplestorage = SimpleStorage(t.temp)
	simplestorage.chassis = t.Links.Chassis.String()
	simplestorage.storage = t.Links.Storage.String()

	return nil
}

// GetSimpleStorage will get a SimpleStorage instance from the service.
func GetSimpleStorage(c common.Client, uri string) (*SimpleStorage, error) {
	var simpleStorage SimpleStorage
	return &simpleStorage, simpleStorage.Get(c, uri, &simpleStorage)
}

// ListReferencedSimpleStorages gets the collection of SimpleStorage from
// a provided reference.
func ListReferencedSimpleStorages(c common.Client, link string) ([]*SimpleStorage, error) { //nolint:dupl
	if link == "" {
		return nil, nil
	}

	type GetResult struct {
		Item  *SimpleStorage
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		simplestorage, err := GetSimpleStorage(c, link)
		ch <- GetResult{Item: simplestorage, Link: link, Error: err}
	}

	var links []string
	var err error
	go func() {
		links, err = common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	// Save unordered results into link-to-SimpleStorage helper map.
	unorderedResults := map[string]*SimpleStorage{}
	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			unorderedResults[r.Link] = r.Item
		}
	}

	if !collectionError.Empty() {
		return nil, collectionError
	}
	// Build the final ordered slice based on the original order from the links list.
	results := make([]*SimpleStorage, len(links))
	for i, link := range links {
		results[i] = unorderedResults[link]
	}

	return results, nil
}

// Chassis gets the chassis containing this storage service.
func (simplestorage *SimpleStorage) Chassis() (*Chassis, error) {
	if simplestorage.chassis == "" {
		return nil, nil
	}

	return GetChassis(simplestorage.GetClient(), simplestorage.chassis)
}

// Storage gets the chassis containing this storage service.
func (simplestorage *SimpleStorage) Storage() (*Storage, error) {
	if simplestorage.storage == "" {
		return nil, nil
	}

	return GetStorage(simplestorage.GetClient(), simplestorage.storage)
}
