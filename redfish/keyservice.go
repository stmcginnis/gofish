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
	nvmeoFKeyPolicies []string
	// NVMeoFSecrets shall contain a link to a resource collection of type KeyCollection that contains the NVMe-oF keys
	// maintained by this service. The KeyType property for all members of this collection shall contain the value
	// 'NVMeoF'.
	nvmeoFSecrets []string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a KeyService object from the raw JSON.
func (keyservice *KeyService) UnmarshalJSON(b []byte) error {
	type temp KeyService
	var t struct {
		temp
		NVMeoFKeyPolicies common.LinksCollection
		NVMeoFSecrets     common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keyservice = KeyService(t.temp)

	// Extract the links to other entities for later
	keyservice.nvmeoFKeyPolicies = t.NVMeoFKeyPolicies.Members.ToStrings()
	keyservice.nvmeoFSecrets = t.NVMeoFSecrets.Members.ToStrings()

	return nil
}

// NVMeoFKeyPolicies gets the NVMe-oF key policies maintained by this service.
func (keyservice *KeyService) NVMeoFKeyPolicies() ([]*KeyPolicy, error) {
	var result []*KeyPolicy

	collectionError := common.NewCollectionError()
	for _, uri := range keyservice.nvmeoFKeyPolicies {
		unit, err := GetKeyPolicy(keyservice.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, unit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// NVMeofSecrets gets the NVMe-oF keys maintained by this service.
func (keyservice *KeyService) NVMeoFSecrets() ([]*Key, error) {
	var result []*Key

	collectionError := common.NewCollectionError()
	for _, uri := range keyservice.nvmeoFSecrets {
		unit, err := GetKey(keyservice.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, unit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// GetKeyService will get a KeyService instance from the service.
func GetKeyService(c common.Client, uri string) (*KeyService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var keyservice KeyService
	err = json.NewDecoder(resp.Body).Decode(&keyservice)
	if err != nil {
		return nil, err
	}

	keyservice.SetClient(c)
	return &keyservice, nil
}

// ListReferencedKeyServices gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServices(c common.Client, link string) ([]*KeyService, error) {
	var result []*KeyService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *KeyService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		keyservice, err := GetKeyService(c, link)
		ch <- GetResult{Item: keyservice, Link: link, Error: err}
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
