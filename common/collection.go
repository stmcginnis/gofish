//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
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

// CollectionError is used for collecting errors when working with collections
type CollectionError struct {
	Failures map[string]error
}

// NewCollectionError gets you a new *CollectionError
// it's useful for collecting and formatting errors that occur when fetching a collection
func NewCollectionError() *CollectionError {
	return &CollectionError{
		Failures: make(map[string]error),
	}
}

func (cr *CollectionError) Empty() bool {
	return len(cr.Failures) == 0
}

// for associating a linked entity with its error
type entityError struct {
	Link  string `json:"link"`
	Error string `json:"error"`
}

func (cr *CollectionError) Error() string {
	var entityErrors []entityError
	for link, err := range cr.Failures {
		entityErrors = append(entityErrors, entityError{
			Link:  link,
			Error: err.Error(),
		})
	}

	errorsJSON, err := json.Marshal(entityErrors)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("failed to retrieve some items: %s", errorsJSON)
}
