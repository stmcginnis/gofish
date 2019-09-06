//
// SPDX-License-Identifier: BSD-3-Clause
//

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
