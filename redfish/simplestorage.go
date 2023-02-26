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
	// Name shall be a user-friendly name for the device.
	Name string
	// CapacityBytes shall represent the size (in bytes) of the Storage Device.
	CapacityBytes int64
	// Manufacturer shall indicate the name of the manufacturer of this storage device.
	Manufacturer string
	// Model shall indicate the model information as provided by the manufacturer
	// of this storage device.
	Model string
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
	// Actions is The Actions property shall contain the available actions
	// for this resource.
	Actions string
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
}

// UnmarshalJSON unmarshals a SimpleStorage object from the raw JSON.
func (simplestorage *SimpleStorage) UnmarshalJSON(b []byte) error {
	type temp SimpleStorage
	var t struct {
		temp
		Links struct {
			// Chassis shall be a reference to a resource of type Chassis that
			// represent the physical container associated with this Simple Storage.
			Chassis common.Link
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*simplestorage = SimpleStorage(t.temp)
	simplestorage.chassis = t.Links.Chassis.String()

	return nil
}

// GetSimpleStorage will get a SimpleStorage instance from the service.
func GetSimpleStorage(c common.Client, uri string) (*SimpleStorage, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var simplestorage SimpleStorage
	err = json.NewDecoder(resp.Body).Decode(&simplestorage)
	if err != nil {
		return nil, err
	}

	simplestorage.SetClient(c)
	return &simplestorage, nil
}

// ListReferencedSimpleStorages gets the collection of SimpleStorage from
// a provided reference.
func ListReferencedSimpleStorages(c common.Client, link string) ([]*SimpleStorage, error) { //nolint:dupl
	var result []*SimpleStorage
	if link == "" {
		return result, nil
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

// Chassis gets the chassis containing this storage service.
func (simplestorage *SimpleStorage) Chassis() (*Chassis, error) {
	if simplestorage.chassis == "" {
		return nil, nil
	}

	return GetChassis(simplestorage.Client, simplestorage.chassis)
}
