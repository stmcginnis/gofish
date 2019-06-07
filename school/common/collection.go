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

package common

import (
	"encoding/json"
)

// Collection represents a collection of entity references.
type Collection struct {
	Name      string `json:"Name"`
	ItemLinks []string
}

// UnmarshalJSON unmarshals a collection from the raw JSON.
func (c *Collection) UnmarshalJSON(b []byte) error {
	type temp Collection
	var t struct {
		temp
		LinksCollection
		Links LinksCollection `json:"Links"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*c = Collection(t.temp)

	// Redfish objects store collection items under Links
	c.ItemLinks = t.Links.ToStrings()

	// Swordfish has them at the root
	if len(c.ItemLinks) == 0 && t.Count > 0 {
		c.ItemLinks = t.Members.ToStrings()
	}

	return nil
}

// GetCollection retrieves a collection from the service.
func GetCollection(c Client, uri string) (*Collection, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Collection
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
