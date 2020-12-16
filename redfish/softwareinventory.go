//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SoftwareInventory is used to represent a software or firmware item on the server
type SoftwareInventory struct {
	common.Entity

	Description            string
	Status                 common.Status
	Version                string
	LowestSupportedVersion string
	Manufacturer           string
	ReleaseDate            string
	SoftwareID             string
	UefiDevicePaths        []string
	Updateable             bool
	WriteProtected         bool
	rawData                []byte
}

// UnmarshalJSON unmarshals a SoftwareInventory object from the raw JSON
func (softwareInventory *SoftwareInventory) UnmarshalJSON(b []byte) error {
	type temp SoftwareInventory
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*softwareInventory = SoftwareInventory(t.temp)
	softwareInventory.rawData = b
	return nil
}

// GetSoftwareInventory will get a SoftwareInventory instance from the service
func GetSoftwareInventory(conn common.Client, uri string) (*SoftwareInventory, error) {
	resp, err := conn.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var softwareInventory SoftwareInventory
	err = json.NewDecoder(resp.Body).Decode(&softwareInventory)
	if err != nil {
		return nil, err
	}
	softwareInventory.SetClient(conn)
	return &softwareInventory, nil
}
