//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/DataStorageLoSCapabilities.v1_2_2.json
// 1.0.7a - #DataStorageLoSCapabilities.v1_2_2.DataStorageLoSCapabilities

package schemas

import (
	"encoding/json"
)

// ProvisioningPolicy is The enumeration literals may be used to specify space
// provisioning policy.
type ProvisioningPolicy string

const (
	// FixedProvisioningPolicy shall be fully allocated.
	FixedProvisioningPolicy ProvisioningPolicy = "Fixed"
	// ThinProvisioningPolicy This enumeration literal specifies storage may be
	// over allocated.
	ThinProvisioningPolicy ProvisioningPolicy = "Thin"
)

// DataStorageLoSCapabilities Each instance of DataStorageLoSCapabilities
// describes capabilities of the system to support various data storage service
// options.
type DataStorageLoSCapabilities struct {
	Entity
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// MaximumRecoverableCapacitySourceCount The maximum number of capacity source
	// resources that can be supported for the purpose of recovery when in the
	// event that an equivalent capacity source resource fails.
	//
	// Version added: v1.2.0
	MaximumRecoverableCapacitySourceCount *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupportedAccessCapabilities Each entry specifies a storage access
	// capability.
	SupportedAccessCapabilities []StorageAccessCapability
	// SupportedLinesOfService shall contain known and supported
	// DataStorageLinesOfService.
	supportedLinesOfService []string
	// SupportedLinesOfServiceCount
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedProvisioningPolicies This collection specifies supported storage
	// allocation policies.
	SupportedProvisioningPolicies []ProvisioningPolicy
	// SupportedRecoveryTimeObjectives This collection specifies supported
	// expectations for time to access the primary store after recovery.
	SupportedRecoveryTimeObjectives []RecoveryAccessScope
	// SupportsSpaceEfficiency The value specifies whether storage compression or
	// deduplication is supported. The default value for this property is false.
	SupportsSpaceEfficiency bool
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a DataStorageLoSCapabilities object from the raw JSON.
func (d *DataStorageLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp DataStorageLoSCapabilities
	var tmp struct {
		temp
		SupportedLinesOfService Links `json:"SupportedLinesOfService"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataStorageLoSCapabilities(tmp.temp)

	// Extract the links to other entities for later
	d.supportedLinesOfService = tmp.SupportedLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataStorageLoSCapabilities) Update() error {
	readWriteFields := []string{
		"MaximumRecoverableCapacitySourceCount",
		"SupportedAccessCapabilities",
		"SupportedLinesOfService",
		"SupportedProvisioningPolicies",
		"SupportedRecoveryTimeObjectives",
		"SupportsSpaceEfficiency",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDataStorageLoSCapabilities will get a DataStorageLoSCapabilities instance from the service.
func GetDataStorageLoSCapabilities(c Client, uri string) (*DataStorageLoSCapabilities, error) {
	return GetObject[DataStorageLoSCapabilities](c, uri)
}

// ListReferencedDataStorageLoSCapabilitiess gets the collection of DataStorageLoSCapabilities from
// a provided reference.
func ListReferencedDataStorageLoSCapabilitiess(c Client, link string) ([]*DataStorageLoSCapabilities, error) {
	return GetCollectionObjects[DataStorageLoSCapabilities](c, link)
}

// SupportedLinesOfService gets the SupportedLinesOfService linked resources.
func (d *DataStorageLoSCapabilities) SupportedLinesOfService() ([]*DataStorageLineOfService, error) {
	return GetObjects[DataStorageLineOfService](d.client, d.supportedLinesOfService)
}
