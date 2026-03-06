//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

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
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var result Collection
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetResourceCollection retrieves a ResourceCollection from the service.
func GetResourceCollection[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryGroupOption) (*ResourceCollectionGeneric[PT], error) {
	resp, err := c.Get(BuildQuery(c, uri, true, queryOpts...))
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	result := new(ResourceCollectionGeneric[PT])
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
func CollectList(get func(string), c Client, link string, queryOpts ...QueryGroupOption) error {
	return CollectListGeneric(func(resource *Resource, _ ...QueryGroupOption) {
		get(resource.ODataID)
	}, c, link, queryOpts...)
}

func CollectListGeneric[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryGroupOption), c Client, link string, queryOpts ...QueryGroupOption) error {
	collection, err := GetResourceCollection[T, PT](c, link, queryOpts...)
	if err != nil {
		// allow for auto-fallback from $expand to regular
		// this will only run on the first query, not future pages
		builtOpts := BuildQueryGroup(c, queryOpts...).QueryCollection
		if builtOpts.expand != ExpandNone && builtOpts.expandFallback {
			queryWithoutExpand := queryOpts
			queryWithoutExpand = append(queryWithoutExpand,
				WithCollectionQueryOpts(WithExpand(ExpandNone)))
			collection, err = GetResourceCollection[T, PT](c, link, queryWithoutExpand...)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	CollectResourceCollection(get, collection.Members, queryOpts...)
	if collection.MembersNextLink != "" {
		err := CollectListGeneric(get, c, collection.MembersNextLink)
		if err != nil {
			return err
		}
	}
	return nil
}

// CollectCollection will retrieve a collection of entities from the Redfish service
// when you already have the set of individual links in the collection.
func CollectCollection(get func(string), links []string) {
	linkEntities := []*Resource{}

	for _, link := range links {
		linkEntities = append(linkEntities, &Resource{Entity: Entity{ODataID: link}})
	}

	CollectResourceCollection(func(resource *Resource, _ ...QueryGroupOption) { get(resource.ODataID) }, linkEntities)
}

func CollectResourceCollection[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryGroupOption), entities []PT, queryOpts ...QueryGroupOption) {
	// Only allow three concurrent requests to avoid overwhelming the service
	limiter := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for _, itemLink := range entities {
		wg.Add(1)
		limiter <- struct{}{}

		go func(itemLink PT, _ ...QueryGroupOption) {
			defer wg.Done()
			get(itemLink)
			<-limiter
		}(itemLink, queryOpts...)
	}

	wg.Wait()
}

// extendedInfoToError converts a slice of Message into a *Error.
func extendedInfoToError(ext []Message) *Error {
	errE := &Error{}
	for i := range ext {
		errE.ExtendedInfos = append(errE.ExtendedInfos, ErrExtendedInfo(ext[i]))
	}
	return errE
}

// collectMemberLinks walks all pages of a typed collection starting at uri,
// returning the ordered list of member @odata.id links. queryOpts are applied
// only to the first page; subsequent pages follow MembersNextLink verbatim.
// Inline-expanded members that carry ExtendedInfo are recorded as errors and
// excluded from the returned links. Inline-expanded members with a valid Id
// but no ExtendedInfo are re-fetched individually by GetCollectionObjects to
// guarantee server-provided ordering; this is an intentional trade-off.
func collectMemberLinks[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, collectionError *CollectionError, queryOpts ...QueryGroupOption) []string {
	var links []string
	next := uri
	firstPage := true
	for next != "" {
		var collection *ResourceCollectionGeneric[PT]
		var err error
		if firstPage {
			collection, err = GetResourceCollection[T, PT](c, next, queryOpts...)
			firstPage = false
		} else {
			collection, err = GetResourceCollection[T, PT](c, next)
		}
		if err != nil {
			collectionError.Failures[next] = err
			break
		}
		for _, m := range collection.Members {
			if m == nil {
				continue
			}
			odataID := m.GetODataID()
			if odataID == "" {
				continue
			}
			// Inline-expanded member with ExtendedInfo is an error; skip fetch.
			if m.GetID() != "" {
				if ext := m.GetExtendedInfo(); len(ext) > 0 {
					collectionError.Failures[odataID] = extendedInfoToError(ext)
					continue
				}
			}
			links = append(links, odataID)
		}
		next = collection.MembersNextLink
	}
	return links
}

// GetCollectionObjects retrieves all members of a Redfish collection at uri,
// preserving the server-provided Members order. The returned slice is dense
// (no nil holes); failed fetches are recorded in the returned CollectionError.
func GetCollectionObjects[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryGroupOption) ([]*T, error) {
	if uri == "" {
		return make([]*T, 0), nil
	}

	collectionError := NewCollectionError()
	links := collectMemberLinks[T, PT](c, uri, collectionError, queryOpts...)

	if len(links) == 0 {
		if collectionError.Empty() {
			return make([]*T, 0), nil
		}
		return make([]*T, 0), collectionError
	}

	// Fetch all members concurrently into a preallocated indexed slice.
	// First occurrence of a duplicate link wins.
	result := make([]*T, len(links))
	index := make(map[string]int, len(links))
	for i, l := range links {
		if _, exists := index[l]; !exists {
			index[l] = i
		}
	}

	var mu sync.Mutex
	get := func(link string) {
		entity, err := GetObject[T, PT](c, link, queryOpts...)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			collectionError.Failures[link] = err
			return
		}
		if entity == nil {
			return
		}
		if extInfo := PT(entity).GetExtendedInfo(); len(extInfo) > 0 {
			collectionError.Failures[link] = extendedInfoToError(extInfo)
			return
		}
		if idx, ok := index[PT(entity).GetODataID()]; ok {
			result[idx] = entity
		} else if idx, ok := index[link]; ok {
			result[idx] = entity
		}
	}
	CollectCollection(get, links)

	// Compact to a dense slice preserving relative order.
	dense := make([]*T, 0, len(result))
	for _, it := range result {
		if it != nil {
			dense = append(dense, it)
		}
	}

	if collectionError.Empty() {
		return dense, nil
	}
	return dense, collectionError
}
