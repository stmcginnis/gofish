//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SoftwareInventoryCollection is used to represent a collection of software or firmware items on the server
type SoftwareInventoryCollection struct {
	common.Entity

	Name        string
	Description string
	members     []string
	rawData     []byte
}

// UnmarshalJSON unmarshals a SoftwareInventoryCollection object from the raw JSON
func (softwareInventory *SoftwareInventoryCollection) UnmarshalJSON(b []byte) error {
	type temp SoftwareInventoryCollection
	var t struct {
		temp
		members common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*softwareInventory = SoftwareInventoryCollection(t.temp)
	softwareInventory.rawData = b
	softwareInventory.members = t.members.ToStrings()
	return nil
}

// Members retrieves the SoftwareInventory objects that are in this collection
func (softwareInventoryCollection *SoftwareInventoryCollection) Members() ([]*SoftwareInventory, error) {
	var result []*SoftwareInventory
	for _, softwareInventoryLink := range softwareInventoryCollection.members {
		softwareInventory, err := GetSoftwareInventory(softwareInventoryCollection.Client, softwareInventoryLink)
		if err != nil {
			return result, nil
		}
		result = append(result, softwareInventory)
	}
	return result, nil
}

// GetSoftwareInventoryCollection will get a SoftwareInventoryCollection instance from the service
func GetSoftwareInventoryCollection(c common.Client, uri string) (*SoftwareInventoryCollection, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var softwareInventoryCollection SoftwareInventoryCollection
	err = json.NewDecoder(resp.Body).Decode(&softwareInventoryCollection)
	if err != nil {
		return nil, err
	}
	softwareInventoryCollection.SetClient(c)
	return &softwareInventoryCollection, nil
}
