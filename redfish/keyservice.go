//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// KeyService shall represent the key service properties for a Redfish implementation.
type KeyService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
func (keyservice *KeyService) NVMeoFKeyPolicies(queryOpts ...common.QueryGroupOption) ([]*KeyPolicy, error) {
	return keyservice.NVMeoFKeyPoliciesWithContext(common.ContextOf(keyservice.GetClient()), queryOpts...)
}

// NVMeoFKeyPoliciesWithContext gets the NVMe-oF key policies maintained by this service.
func (keyservice *KeyService) NVMeoFKeyPoliciesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*KeyPolicy, error) {
	return ListReferencedKeyPolicysWithContext(ctx, keyservice.GetClient(), keyservice.nvmeoFKeyPolicies, queryOpts...)
}

// NVMeofSecrets gets the NVMe-oF keys maintained by this service.
func (keyservice *KeyService) NVMeoFSecrets(queryOpts ...common.QueryGroupOption) ([]*Key, error) {
	return keyservice.NVMeoFSecretsWithContext(common.ContextOf(keyservice.GetClient()), queryOpts...)
}

// NVMeofSecrets gets the NVMe-oF keys maintained by this service.
func (keyservice *KeyService) NVMeoFSecretsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Key, error) {
	return ListReferencedKeysWithContext(ctx, keyservice.GetClient(), keyservice.nvmeoFSecrets, queryOpts...)
}

// GetKeyService will get a KeyService instance from the service.
func GetKeyService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*KeyService, error) {
	return GetKeyServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetKeyServiceWithContext will get a KeyService instance from the service.
func GetKeyServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*KeyService, error) {
	return common.GetObjectWithContext[KeyService](ctx, c, uri, queryOpts...)
}

// ListReferencedKeyServices gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServices(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*KeyService, error) {
	return ListReferencedKeyServicesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedKeyServicesWithContext gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServicesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*KeyService, error) {
	return common.GetCollectionObjectsWithContext[KeyService](ctx, c, link, queryOpts...)
}
