//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"context"
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

// GetCollection retrieves a collection from the service using the client's cached context.
func GetCollection(c Client, uri string) (*Collection, error) {
	return GetCollectionWithContext(ContextOf(c), c, uri)
}

// GetCollectionWithContext retrieves a collection from the service.
func GetCollectionWithContext(ctx context.Context, c Client, uri string) (*Collection, error) {
	resp, err := c.GetWithContext(ctx, uri)
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

// GetResourceCollection retrieves a ResourceCollection from the service using the client's cached context.
func GetResourceCollection[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryGroupOption) (*ResourceCollectionGeneric[PT], error) {
	return GetResourceCollectionWithContext[T, PT](ContextOf(c), c, uri, queryOpts...)
}

// GetResourceCollectionWithContext retrieves a ResourceCollection from the service.
func GetResourceCollectionWithContext[T any, PT interface {
	*T
	SchemaObject
}](ctx context.Context, c Client, uri string, queryOpts ...QueryGroupOption) (*ResourceCollectionGeneric[PT], error) {
	resp, err := c.GetWithContext(ctx, BuildQuery(c, uri, true, queryOpts...))
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

// Unwrap returns the underlying per-member errors so that errors.Is and
// errors.As see through a CollectionError (e.g. errors.Is(err, context.Canceled)
// reports true when member fetches failed due to cancellation).
func (cr *CollectionError) Unwrap() []error {
	errs := make([]error, 0, len(cr.Failures))
	for _, err := range cr.Failures {
		errs = append(errs, err)
	}
	return errs
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

// CollectList will retrieve a collection of entities from the Redfish service using the client's cached context.
func CollectList(get func(string), c Client, link string, queryOpts ...QueryGroupOption) error {
	return CollectListWithContext(ContextOf(c), get, c, link, queryOpts...)
}

// CollectListWithContext will retrieve a collection of entities from the Redfish service.
func CollectListWithContext(ctx context.Context, get func(string), c Client, link string, queryOpts ...QueryGroupOption) error {
	return CollectListGenericWithContext(ctx, func(resource *Resource, _ ...QueryGroupOption) {
		get(resource.ODataID)
	}, c, link, queryOpts...)
}

func CollectListGeneric[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryGroupOption), c Client, link string, queryOpts ...QueryGroupOption) error {
	return CollectListGenericWithContext[T, PT](ContextOf(c), get, c, link, queryOpts...)
}

func CollectListGenericWithContext[T any, PT interface {
	*T
	SchemaObject
}](ctx context.Context, get func(PT, ...QueryGroupOption), c Client, link string, queryOpts ...QueryGroupOption) error {
	collection, err := GetResourceCollectionWithContext[T, PT](ctx, c, link, queryOpts...)
	if err != nil {
		// A dead context fails every retry the same way; surface it as-is.
		if ctx.Err() != nil {
			return err
		}
		// allow for auto-fallback from $expand to regular
		// this will only run on the first query, not future pages
		builtOpts := BuildQueryGroup(c, queryOpts...).QueryCollection
		if builtOpts.expand != ExpandNone && builtOpts.expandFallback {
			queryWithoutExpand := queryOpts
			queryWithoutExpand = append(queryWithoutExpand,
				WithCollectionQueryOpts(WithExpand(ExpandNone)))
			collection, err = GetResourceCollectionWithContext[T, PT](ctx, c, link, queryWithoutExpand...)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	CollectResourceCollectionWithContext(ctx, get, collection.Members, queryOpts...)
	if collection.MembersNextLink != "" {
		if err := ctx.Err(); err != nil {
			return err
		}
		err := CollectListGenericWithContext(ctx, get, c, collection.MembersNextLink, queryOpts...)
		if err != nil {
			return err
		}
	}
	return nil
}

// CollectCollection will retrieve a collection of entities from the Redfish service
// when you already have the set of individual links in the collection.
func CollectCollection(get func(string), links []string) {
	CollectCollectionWithContext(context.Background(), get, links)
}

// CollectCollectionWithContext will retrieve a collection of entities from the
// Redfish service when you already have the set of individual links in the
// collection.
func CollectCollectionWithContext(ctx context.Context, get func(string), links []string) {
	linkEntities := []*Resource{}

	for _, link := range links {
		linkEntities = append(linkEntities, &Resource{Entity: Entity{ODataID: link}})
	}

	CollectResourceCollectionWithContext(ctx, func(resource *Resource, _ ...QueryGroupOption) { get(resource.ODataID) }, linkEntities)
}

func CollectResourceCollection[T any, PT interface {
	*T
	SchemaObject
}](get func(PT, ...QueryGroupOption), entities []PT, queryOpts ...QueryGroupOption) {
	CollectResourceCollectionWithContext(context.Background(), get, entities, queryOpts...)
}

func CollectResourceCollectionWithContext[T any, PT interface {
	*T
	SchemaObject
}](ctx context.Context, get func(PT, ...QueryGroupOption), entities []PT, queryOpts ...QueryGroupOption) {
	// Only allow three concurrent requests to avoid overwhelming the service
	limiter := make(chan struct{}, 3)
	var wg sync.WaitGroup

	canceled := -1
	for i, itemLink := range entities {
		if ctx.Err() != nil {
			canceled = i
			break
		}
		select {
		case <-ctx.Done():
			canceled = i
		case limiter <- struct{}{}:
		}
		if canceled >= 0 {
			break
		}

		wg.Add(1)
		go func(itemLink PT, opts ...QueryGroupOption) {
			defer wg.Done()
			get(itemLink, opts...)
			<-limiter
		}(itemLink, queryOpts...)
	}
	if canceled >= 0 {
		// The context died mid-iteration. Run the remaining members through
		// get synchronously instead of spawning a fetch goroutine per member,
		// so every member is still accounted for by the caller: when get
		// performs the fetch it fails fast with the context error; when the
		// member was already materialized (e.g. via $expand) it is delivered
		// without further I/O.
		for _, itemLink := range entities[canceled:] {
			get(itemLink, queryOpts...)
		}
	}

	wg.Wait()
}

func GetCollectionObjects[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string, queryOpts ...QueryGroupOption) ([]*T, error) {
	return GetCollectionObjectsWithContext[T, PT](ContextOf(c), c, uri, queryOpts...)
}

func GetCollectionObjectsWithContext[T any, PT interface {
	*T
	SchemaObject
}](ctx context.Context, c Client, uri string, queryOpts ...QueryGroupOption) ([]*T, error) {
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
	get := func(entity PT, opts ...QueryGroupOption) {
		if entity != nil && entity.GetID() != "" {
			// if the entity has any ExtendedInfo, we assume it's an error
			var err error
			extendedInfo := entity.GetExtendedInfo()
			if len(extendedInfo) > 0 {
				errE := &Error{}
				for _, info := range extendedInfo {
					errE.ExtendedInfos = append(errE.ExtendedInfos, ErrExtendedInfo(info))
				}
				err = errE
			}

			entity.SetClient(c)

			ch <- GetResult{Item: entity, Link: entity.GetODataID(), Error: err}
		} else if entity != nil && entity.GetODataID() != "" {
			link := entity.GetODataID()
			entity, err := GetObjectWithContext[T, PT](ctx, c, link, opts...)
			ch <- GetResult{Item: entity, Link: link, Error: err}
		}
	}

	go func() {
		// A collection-level failure is delivered through the channel so that
		// Failures is only ever written by the consumer loop below.
		if err := CollectListGenericWithContext(ctx, get, c, uri, queryOpts...); err != nil {
			ch <- GetResult{Link: uri, Error: err}
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
