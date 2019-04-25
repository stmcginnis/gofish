// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
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

// DefaultCompositionServicePath is the default URI for CompositionService collections.
const DefaultCompositionServicePath = "/redfish/v1/CompositionService"

// CompositionService defines a Composition Service which represents the
// properties for the service and links to the resources available for
// composition.
type CompositionService struct {
	common.Entity
	Status         common.Status
	ServiceEnabled string
	resourceBlocks string
	resourceZones  string
}

// UnmarshalJSON unmarshals CompositionService object from the raw JSON.
func (cs *CompositionService) UnmarshalJSON(b []byte) error {
	type temp CompositionService
	var t struct {
		temp
		ResourceBlocks common.Link
		ResourceZones  common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cs = CompositionService(t.temp)

	// Extract the links to other entities for later
	cs.resourceBlocks = string(t.ResourceBlocks)
	cs.resourceZones = string(t.ResourceZones)

	return nil
}

// GetCompositionService will get a Composition instance from the Redfish service.
func GetCompositionService(c common.Client, uri string) (*CompositionService, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var t CompositionService
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
