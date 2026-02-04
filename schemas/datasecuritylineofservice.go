//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/DataSecurityLineOfService.v1_1_1.json
// 1.1.0 - #DataSecurityLineOfService.v1_1_1.DataSecurityLineOfService

package schemas

import (
	"encoding/json"
)

// DataSecurityLineOfService shall be used to describe data security service
// level requirements.
type DataSecurityLineOfService struct {
	Entity
	// AntivirusEngineProvider shall specify an AntiVirus provider.
	AntivirusEngineProvider string
	// AntivirusScanPolicies shall specify the policy for triggering an AntiVirus
	// scan.
	AntivirusScanPolicies []AntiVirusScanTrigger
	// ChannelEncryptionStrength shall specify a key size in a symmetric encryption
	// algorithm for transport channel encryption.
	ChannelEncryptionStrength KeySize
	// DataSanitizationPolicy shall specify the data sanitization policy.
	DataSanitizationPolicy DataSanitizationPolicy
	// HostAuthenticationType shall specify the authentication type for hosts
	// (servers) or initiator endpoints.
	HostAuthenticationType AuthenticationType
	// MediaEncryptionStrength shall specify a key size in a symmetric encryption
	// algorithm for media encryption.
	MediaEncryptionStrength KeySize
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SecureChannelProtocol shall specify the protocol that provide encrypted
	// communication.
	SecureChannelProtocol SecureChannelProtocol
	// UserAuthenticationType shall specify the authentication type for users (or
	// programs).
	UserAuthenticationType AuthenticationType
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a DataSecurityLineOfService object from the raw JSON.
func (d *DataSecurityLineOfService) UnmarshalJSON(b []byte) error {
	type temp DataSecurityLineOfService
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataSecurityLineOfService(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataSecurityLineOfService) Update() error {
	readWriteFields := []string{
		"AntivirusEngineProvider",
		"AntivirusScanPolicies",
		"ChannelEncryptionStrength",
		"DataSanitizationPolicy",
		"HostAuthenticationType",
		"MediaEncryptionStrength",
		"SecureChannelProtocol",
		"UserAuthenticationType",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDataSecurityLineOfService will get a DataSecurityLineOfService instance from the service.
func GetDataSecurityLineOfService(c Client, uri string) (*DataSecurityLineOfService, error) {
	return GetObject[DataSecurityLineOfService](c, uri)
}

// ListReferencedDataSecurityLineOfServices gets the collection of DataSecurityLineOfService from
// a provided reference.
func ListReferencedDataSecurityLineOfServices(c Client, link string) ([]*DataSecurityLineOfService, error) {
	return GetCollectionObjects[DataSecurityLineOfService](c, link)
}
