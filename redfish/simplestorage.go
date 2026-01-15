//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.3 - #SimpleStorage.v1_3_2.SimpleStorage

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SimpleStorage This resource contains a storage controller and its
// directly-attached devices.
type SimpleStorage struct {
	common.Entity
	// Devices shall contain a list of storage devices related to this resource.
	Devices []Device
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UefiDevicePath shall contain the UEFI device path that identifies and
	// locates the specific storage controller.
	UefiDevicePath string
	// chassis is the URI for Chassis.
	chassis string
	// storage is the URI for Storage.
	storage string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SimpleStorage object from the raw JSON.
func (s *SimpleStorage) UnmarshalJSON(b []byte) error {
	type temp SimpleStorage
	type sLinks struct {
		Chassis common.Link `json:"Chassis"`
		Storage common.Link `json:"Storage"`
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

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SimpleStorage) Update() error {
	readWriteFields := []string{
		"Devices",
		"Status",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSimpleStorage will get a SimpleStorage instance from the service.
func GetSimpleStorage(c common.Client, uri string) (*SimpleStorage, error) {
	return common.GetObject[SimpleStorage](c, uri)
}

// ListReferencedSimpleStorages gets the collection of SimpleStorage from
// a provided reference.
func ListReferencedSimpleStorages(c common.Client, link string) ([]*SimpleStorage, error) {
	return common.GetCollectionObjects[SimpleStorage](c, link)
}

// Chassis gets the Chassis linked resource.
func (s *SimpleStorage) Chassis(client common.Client) (*Chassis, error) {
	if s.chassis == "" {
		return nil, nil
	}
	return common.GetObject[Chassis](client, s.chassis)
}

// Storage gets the Storage linked resource.
func (s *SimpleStorage) Storage(client common.Client) (*Storage, error) {
	if s.storage == "" {
		return nil, nil
	}
	return common.GetObject[Storage](client, s.storage)
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
	// Oem shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}
