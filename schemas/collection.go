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
	// Concurrency is bounded at the HTTP layer by the client's
	// MaxConcurrentRequests setting, so no additional cap is applied here.
	var wg sync.WaitGroup

	for _, itemLink := range entities {
		wg.Add(1)

		go func(itemLink PT, _ ...QueryGroupOption) {
			defer wg.Done()
			get(itemLink)
		}(itemLink, queryOpts...)
	}

	wg.Wait()
}

func GetCollectionObjects[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryGroupOption) ([]*T, error) {
	if uri == "" {
		return nil, nil
	}

	collectionError := NewCollectionError()

	// Gather the member entities in the order the service returned them, then
	// resolve each concurrently into a slice indexed by that position so the
	// result preserves the service's ordering regardless of fetch completion
	// order.
	members, err := collectMembers[T, PT](c, uri, queryOpts...)
	if err != nil {
		collectionError.Failures[uri] = err
	}

	type memberResult struct {
		item *T
		link string
		err  error
	}

	// One goroutine per member, gated at the HTTP layer by the client
	// semaphore (MaxConcurrentRequests).
	results := make([]memberResult, len(members))
	var wg sync.WaitGroup
	for i, entity := range members {
		wg.Add(1)
		go func(i int, entity PT) {
			defer wg.Done()
			item, link, err := resolveMember[T, PT](c, entity, queryOpts...)
			results[i] = memberResult{item: item, link: link, err: err}
		}(i, entity)
	}
	wg.Wait()

	var result []*T
	for _, r := range results {
		switch {
		case r.err != nil:
			collectionError.Failures[r.link] = r.err
		case r.item != nil:
			result = append(result, r.item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// collectMembers gathers all member entities of a collection in the order the
// service returned them, following pagination and falling back from $expand to
// a plain query if the expanded request fails.
func collectMembers[T any, PT interface {
	*T
	SchemaObject
}](c Client, link string, queryOpts ...QueryGroupOption) ([]PT, error) {
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
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	members := collection.Members
	if collection.MembersNextLink != "" {
		next, err := collectMembers[T, PT](c, collection.MembersNextLink)
		members = append(members, next...)
		if err != nil {
			return members, err
		}
	}
	return members, nil
}

// resolveMember turns a collection member entity into a fully populated object.
// A member returned inline (already carrying an ID) is used as-is; otherwise it
// is fetched from its @odata.id. It returns the resolved object, the member
// link (for error reporting), and any error.
func resolveMember[T any, PT interface {
	*T
	SchemaObject
}](c Client, entity PT, opts ...QueryGroupOption) (item *T, link string, err error) {
	if entity == nil {
		return nil, "", nil
	}

	if entity.GetID() != "" {
		// if the entity has any ExtendedInfo, we assume it's an error
		extendedInfo := entity.GetExtendedInfo()
		if len(extendedInfo) > 0 {
			errE := &Error{}
			for i := range extendedInfo {
				errE.ExtendedInfos = append(errE.ExtendedInfos, ErrExtendedInfo(extendedInfo[i]))
			}
			err = errE
		}

		entity.SetClient(c)
		return entity, entity.GetODataID(), err
	}

	if entity.GetODataID() != "" {
		link = entity.GetODataID()
		item, err = GetObject[T, PT](c, link, opts...)
		return item, link, err
	}

	return nil, "", nil
}
