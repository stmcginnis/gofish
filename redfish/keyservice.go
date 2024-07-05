//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// KeyService shall represent the key service properties for a Redfish implementation.
type KeyService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// NVMeoFKeyPolicies shall contain a link to a resource collection of type KeyPolicyCollection that contains the
	// NVMe-oF key policies maintained by this service. The KeyPolicyType property for all members of this collection
	// shall contain the value 'NVMeoF'.
	nvmeoFKeyPolicies string
	// NVMeoFSecrets shall contain a link to a resource collection of type KeyCollection that contains the NVMe-oF keys
	// maintained by this service. The KeyType property for all members of this collection shall contain the value
	// 'NVMeoF'.
	nvmeoFSecrets string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a KeyService object from the raw JSON.
func (keyservice *KeyService) UnmarshalJSON(b []byte) error {
	type temp KeyService
	var t struct {
		temp
		NVMeoFKeyPolicies common.Link
		NVMeoFSecrets     common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keyservice = KeyService(t.temp)

	// Extract the links to other entities for later
	keyservice.nvmeoFKeyPolicies = t.NVMeoFKeyPolicies.String()
	keyservice.nvmeoFSecrets = t.NVMeoFSecrets.String()

	return nil
}

// NVMeoFKeyPolicies gets the NVMe-oF key policies maintained by this service.
func (keyservice *KeyService) NVMeoFKeyPolicies() ([]*KeyPolicy, error) {
	return ListReferencedKeyPolicys(keyservice.GetClient(), keyservice.nvmeoFKeyPolicies)
}

// NVMeofSecrets gets the NVMe-oF keys maintained by this service.
func (keyservice *KeyService) NVMeoFSecrets() ([]*Key, error) {
	return ListReferencedKeys(keyservice.GetClient(), keyservice.nvmeoFSecrets)
}

// GetKeyService will get a KeyService instance from the service.
func GetKeyService(c common.Client, uri string) (*KeyService, error) {
	return common.GetObject[KeyService](c, uri)
}

// ListReferencedKeyServices gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServices(c common.Client, link string) ([]*KeyService, error) {
	return common.GetCollectionObjects[KeyService](c, link)
}
