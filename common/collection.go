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
	Items           json.RawMessage
	MembersNextLink string `json:"Members@odata.nextLink,omitempty"`
}

// UnmarshalJSON unmarshals a collection from the raw JSON.
func (c *Collection) UnmarshalJSON(b []byte) error {
	type temp Collection
	var t struct {
		temp
		ODataCount   int `json:"@odata.count"`
		Count        int `json:"Members@odata.count"`
		MembersLinks Links
		MembersRaw   json.RawMessage `json:"Members"`
		Links        LinksCollection `json:"Links"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*c = Collection(t.temp)

	// Redfish objects store collection items under Links
	c.ItemLinks = t.Links.ToStrings()
	c.Items = t.MembersRaw

	if t.MembersRaw != nil {
		err = json.Unmarshal(t.MembersRaw, &t.MembersLinks)
		if err != nil {
			return err
		}
	}

	// Swordfish has them at the root
	if len(c.ItemLinks) == 0 &&
		(t.Count > 0 || t.ODataCount > 0 || len(t.MembersLinks) > 0) {
		c.ItemLinks = t.MembersLinks.ToStrings()
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
func CollectList[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryOption), c Client, link string, queryOpts ...QueryOption) error {
	collection, err := GetCollection(c, BuildQueryForCollection(link, queryOpts...))
	if err != nil {
		return err
	}

	var items []PT
	if collection.Items != nil {
		err = json.Unmarshal(collection.Items, &items)
		if err != nil {
			return err
		}
	}

	CollectCollection(get, items, queryOpts...)
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
func CollectCollection[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryOption), items []PT, queryOpts ...QueryOption) {
	// Only allow three concurrent requests to avoid overwhelming the service
	limiter := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for _, item := range items {

		wg.Add(1)
		limiter <- struct{}{}

		go func(item PT) {
			defer wg.Done()
			get(item, queryOpts...)
			<-limiter
		}(item)
	}

	wg.Wait()
}

func GetCollectionObjects[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryOption) ([]*T, error) {
	var result []*T
	if uri == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *T
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := NewCollectionError()
	get := func(item PT, queryOpts ...QueryOption) {
		if item.GetID() != "" {
			item.SetClient(c)
			ch <- GetResult{Item: item, Link: item.GetODataID(), Error: nil}
			return
		}
		entity, err := GetObject[T, PT](c, BuildQuery(item.GetODataID(), queryOpts...))
		ch <- GetResult{Item: entity, Link: item.GetODataID(), Error: err}
	}

	go func() {
		err := CollectList(get, c, uri, queryOpts...)
		if err != nil {
			collectionError.Failures[uri] = err
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
