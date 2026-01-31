//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.1.0 - #DataStorageLineOfService.v1_3_1.DataStorageLineOfService

package schemas

import (
	"encoding/json"
)

// DataStorageLineOfService This structure may be used to describe a service
// option covering storage provisioning and availability.
type DataStorageLineOfService struct {
	Entity
	// AccessCapabilities Each entry specifies a required storage access
	// capability.
	//
	// Version added: v1.1.0
	AccessCapabilities []StorageAccessCapability
	// IsSpaceEfficient shall indicate that the storage is compressed or
	// deduplicated. The default value for this property is false.
	IsSpaceEfficient bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProvisioningPolicy shall define the provisioning policy for storage.
	ProvisioningPolicy ProvisioningPolicy
	// RecoverableCapacitySourceCount shall be available in the event that an
	// equivalent capacity source resource fails. It is assumed that drives and
	// memory components can be replaced, repaired or otherwise added to increase
	// an associated resource's RecoverableCapacitySourceCount.
	//
	// Version added: v1.2.0
	RecoverableCapacitySourceCount *int `json:",omitempty"`
	// RecoveryTimeObjectives shall regain conformant service level access to the
	// primary store, typical values are 'immediate' or 'offline'. The expectation
	// is that the services required to implement this capability are part of the
	// advertising system.
	RecoveryTimeObjectives RecoveryAccessScope
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a DataStorageLineOfService object from the raw JSON.
func (d *DataStorageLineOfService) UnmarshalJSON(b []byte) error {
	type temp DataStorageLineOfService
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataStorageLineOfService(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataStorageLineOfService) Update() error {
	readWriteFields := []string{
		"AccessCapabilities",
		"IsSpaceEfficient",
		"ProvisioningPolicy",
		"RecoverableCapacitySourceCount",
		"RecoveryTimeObjectives",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDataStorageLineOfService will get a DataStorageLineOfService instance from the service.
func GetDataStorageLineOfService(c Client, uri string) (*DataStorageLineOfService, error) {
	return GetObject[DataStorageLineOfService](c, uri)
}

// ListReferencedDataStorageLineOfServices gets the collection of DataStorageLineOfService from
// a provided reference.
func ListReferencedDataStorageLineOfServices(c Client, link string) ([]*DataStorageLineOfService, error) {
	return GetCollectionObjects[DataStorageLineOfService](c, link)
}
