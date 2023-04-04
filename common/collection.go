//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Collection represents a collection of entity references.
type Collection struct {
	Name            string `json:"Name"`
	ItemLinks       []string
	MembersNextLink string `json:"Members@odata.nextLink,omitempty"`
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
	if len(c.ItemLinks) == 0 &&
		(t.Count > 0 || t.ODataCount > 0 || len(t.Members) > 0) {
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

// CollectList will retrieve a collection of entities from the Redfish service.
func CollectList(get func(string), c Client, link string) error {
	collection, err := GetCollection(c, link)
	if err != nil {
		return err
	}

	CollectCollection(get, c, collection.ItemLinks)
	if collection.MembersNextLink != "" {
		err := CollectList(get, c, collection.MembersNextLink)
		if err != nil {
			return err
		}
	}
	return nil
}

// CollectCollection will retrieve a collection of entitied from the Redfish service
// when you already have the set of individual links in the collection.
func CollectCollection(get func(string), c Client, links []string) {
	// Only allow three concurrent requests to avoid overwhelming the service
	limiter := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for _, itemLink := range links {
		wg.Add(1)
		limiter <- struct{}{}

		go func(itemLink string) {
			defer wg.Done()
			get(itemLink)
			<-limiter
		}(itemLink)
	}

	wg.Wait()
}
