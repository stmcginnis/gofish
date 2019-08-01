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

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// Bios is used to represent BIOS attributes.
type Bios struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AttributeRegistry is the Resource ID of the Attribute Registry that has
	// the system-specific information about a BIOS resource.
	AttributeRegistry string
	// Attributes are additional properties in this object, and can be looked up
	// in the Attribute Registry by their AttributeName.
	// Attributes string
	// Description provides a description of this resource.
	Description string
}

// GetBios will get a Bios instance from the service.
func GetBios(c common.Client, uri string) (*Bios, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bios Bios
	err = json.NewDecoder(resp.Body).Decode(&bios)
	if err != nil {
		return nil, err
	}

	bios.SetClient(c)
	return &bios, nil
}

// ListReferencedBioss gets the collection of Bios from a provided reference.
func ListReferencedBioss(c common.Client, link string) ([]*Bios, error) {
	var result []*Bios
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, biosLink := range links.ItemLinks {
		bios, err := GetBios(c, biosLink)
		if err != nil {
			return result, err
		}
		result = append(result, bios)
	}

	return result, nil
}
