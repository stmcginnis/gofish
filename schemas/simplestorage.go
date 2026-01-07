//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.3 - #SimpleStorage.v1_3_2.SimpleStorage

package schemas

import (
	"encoding/json"
)

// SimpleStorage This resource contains a storage controller and its
// directly-attached devices.
type SimpleStorage struct {
	Entity
	// Devices shall contain a list of storage devices related to this resource.
	Devices []Device
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
	// UefiDevicePath shall contain the UEFI device path that identifies and
	// locates the specific storage controller.
	UefiDevicePath string
	// chassis is the URI for Chassis.
	chassis string
	// storage is the URI for Storage.
	storage string
}

// UnmarshalJSON unmarshals a SimpleStorage object from the raw JSON.
func (s *SimpleStorage) UnmarshalJSON(b []byte) error {
	type temp SimpleStorage
	type sLinks struct {
		Chassis Link `json:"Chassis"`
		Storage Link `json:"Storage"`
	}
	var tmp struct {
		temp
		Links sLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SimpleStorage(tmp.temp)

	// Extract the links to other entities for later
	s.chassis = tmp.Links.Chassis.String()
	s.storage = tmp.Links.Storage.String()

	return nil
}

// GetSimpleStorage will get a SimpleStorage instance from the service.
func GetSimpleStorage(c Client, uri string) (*SimpleStorage, error) {
	return GetObject[SimpleStorage](c, uri)
}

// ListReferencedSimpleStorages gets the collection of SimpleStorage from
// a provided reference.
func ListReferencedSimpleStorages(c Client, link string) ([]*SimpleStorage, error) {
	return GetCollectionObjects[SimpleStorage](c, link)
}

// Chassis gets the Chassis linked resource.
func (s *SimpleStorage) Chassis() (*Chassis, error) {
	if s.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](s.client, s.chassis)
}

// Storage gets the Storage linked resource.
func (s *SimpleStorage) Storage() (*Storage, error) {
	if s.storage == "" {
		return nil, nil
	}
	return GetObject[Storage](s.client, s.storage)
}

// Device shall describe a storage device visible to simple storage.
type Device struct {
	// CapacityBytes shall represent the size, in bytes, of the storage device.
	//
	// Version added: v1.1.0
	CapacityBytes *uint `json:",omitempty"`
	// Manufacturer shall indicate the name of the manufacturer of this storage
	// device.
	Manufacturer string
	// Model shall indicate the model information as provided by the manufacturer
	// of this storage device.
	Model string
	// Name is the name of the resource or array element.
	Name string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
}
