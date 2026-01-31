//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.2 - #KeyService.v1_0_1.KeyService

package schemas

import (
	"encoding/json"
)

// KeyService shall represent the key service properties for a Redfish
// implementation.
type KeyService struct {
	Entity
	// NVMeoFKeyPolicies shall contain a link to a resource collection of type
	// 'KeyPolicyCollection' that contains the NVMe-oF key policies maintained by
	// this service. The 'KeyPolicyType' property for all members of this
	// collection shall contain the value 'NVMeoF'.
	nVMeoFKeyPolicies string
	// NVMeoFSecrets shall contain a link to a resource collection of type
	// 'KeyCollection' that contains the NVMe-oF keys maintained by this service.
	// The 'KeyType' property for all members of this collection shall contain the
	// value 'NVMeoF'.
	nVMeoFSecrets string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a KeyService object from the raw JSON.
func (k *KeyService) UnmarshalJSON(b []byte) error {
	type temp KeyService
	var tmp struct {
		temp
		NVMeoFKeyPolicies Link `json:"NVMeoFKeyPolicies"`
		NVMeoFSecrets     Link `json:"NVMeoFSecrets"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*k = KeyService(tmp.temp)

	// Extract the links to other entities for later
	k.nVMeoFKeyPolicies = tmp.NVMeoFKeyPolicies.String()
	k.nVMeoFSecrets = tmp.NVMeoFSecrets.String()

	return nil
}

// GetKeyService will get a KeyService instance from the service.
func GetKeyService(c Client, uri string) (*KeyService, error) {
	return GetObject[KeyService](c, uri)
}

// ListReferencedKeyServices gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServices(c Client, link string) ([]*KeyService, error) {
	return GetCollectionObjects[KeyService](c, link)
}

// NVMeoFKeyPolicies gets the NVMeoFKeyPolicies collection.
func (k *KeyService) NVMeoFKeyPolicies() ([]*KeyPolicy, error) {
	if k.nVMeoFKeyPolicies == "" {
		return nil, nil
	}
	return GetCollectionObjects[KeyPolicy](k.client, k.nVMeoFKeyPolicies)
}

// NVMeoFSecrets gets the NVMeoFSecrets collection.
func (k *KeyService) NVMeoFSecrets() ([]*Key, error) {
	if k.nVMeoFSecrets == "" {
		return nil, nil
	}
	return GetCollectionObjects[Key](k.client, k.nVMeoFSecrets)
}
