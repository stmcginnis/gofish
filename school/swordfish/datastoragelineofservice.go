// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// DataStorageLineOfService is used to describe a service option covering
// storage provisioning and availability.
type DataStorageLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessCapabilities is Each entry specifies a required storage access
	// capability.
	AccessCapabilities []StorageAccessCapability
	// Description provides a description of this resource.
	Description string
	// IsSpaceEfficient is A value of true shall indicate that the storage is
	// compressed or deduplicated. The default value for this property is
	// false.
	IsSpaceEfficient bool
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM string `json:"Oem"`
	// ProvisioningPolicy is The enumeration literal shall define the
	// provisioning policy for storage.
	ProvisioningPolicy ProvisioningPolicy
	// RecoverableCapacitySourceCount is The value is minimum required number
	// of available capacity source resources that shall be available in the
	// event that an equivalent capacity source resource fails.  It is
	// assumed that drives and memory components can be replaced, repaired or
	// otherwise added to increase an associated resource's
	// RecoverableCapacitySourceCount.
	RecoverableCapacitySourceCount int
	// RecoveryTimeObjectives is The enumeration literal specifies the time
	// after a disaster that the client shall regain conformant service level
	// access to the primary store, typical values are 'immediate' or
	// 'offline'. The expectation is that the services required to implement
	// this capability are part of the advertising system.
	RecoveryTimeObjectives RecoveryAccessScope
}

// UnmarshalJSON unmarshals a DataStorageLineOfService object from the raw JSON.
func (datastoragelineofservice *DataStorageLineOfService) UnmarshalJSON(b []byte) error {
	type temp DataStorageLineOfService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*datastoragelineofservice = DataStorageLineOfService(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetDataStorageLineOfService will get a DataStorageLineOfService instance from the service.
func GetDataStorageLineOfService(c common.Client, uri string) (*DataStorageLineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var datastoragelineofservice DataStorageLineOfService
	err = json.NewDecoder(resp.Body).Decode(&datastoragelineofservice)
	if err != nil {
		return nil, err
	}

	datastoragelineofservice.SetClient(c)
	return &datastoragelineofservice, nil
}

// ListReferencedDataStorageLineOfServices gets the collection of DataStorageLineOfService from
// a provided reference.
func ListReferencedDataStorageLineOfServices(c common.Client, link string) ([]*DataStorageLineOfService, error) {
	var result []*DataStorageLineOfService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, datastoragelineofserviceLink := range links.ItemLinks {
		datastoragelineofservice, err := GetDataStorageLineOfService(c, datastoragelineofserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, datastoragelineofservice)
	}

	return result, nil
}
