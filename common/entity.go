//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// Entity provides the common basis for all Redfish and Swordfish objects.
type Entity struct {
	// ODataID is the location of the resource.
	ODataID string `json:"@odata.id"`
	// ID uniquely identifies the resource.
	ID string `json:"Id"`
	// Name is the name of the resource or array element.
	Name string `json:"Name"`
	// Client is the REST client interface to the system.
	client Client
	// etag contains the etag header when fetching the object. This is used to
	// control updates to make sure the object has not been modified my a different
	// process between fetching and updating that could cause conflicts.
	etag string
	// Removes surrounding quotes of etag used in If-Match header of PATCH and POST requests.
	// Only use this option to resolve bad vendor implementation where If-Match only matches the unquoted etag string.
	stripEtagQuotes bool
	// DisableEtagMatch when set will skip the If-Match header from PATCH and POST requests.
	//
	// This is a work around for bad vendor implementations where the If-Match header does not work - even with the '*' value
	// and requests are incorrectly denied with an ETag mismatch error.
	disableEtagMatch bool
}

// SetClient sets the API client connection to use for accessing this
// entity.
func (e *Entity) SetClient(c Client) {
	e.client = c
}

// SetETag sets the etag value of this API object.
func (e *Entity) SetETag(tag string) {
	e.etag = tag
}

// GetClient get the API client connection to use for accessing this
// entity.
func (e *Entity) GetClient() Client {
	return e.client
}

// Set stripEtagQuotes to enable/disable strupping etag quotes
func (e *Entity) StripEtagQuotes(b bool) {
	e.stripEtagQuotes = b
}

// Disable Etag Match header from being sent by the client.
func (e *Entity) DisableEtagMatch(b bool) {
	e.disableEtagMatch = b
}

// Update commits changes to an entity.
func (e *Entity) Update(originalEntity, updatedEntity reflect.Value, allowedUpdates []string) error {
	payload := getPatchPayloadFromUpdate(originalEntity, updatedEntity)

	// See if we are attempting to update anything that is not allowed
	for field := range payload {
		found := false
		for _, name := range allowedUpdates {
			if name == field {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("%s field is read only", field)
		}
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		return e.Patch(e.ODataID, payload)
	}

	return nil
}

// Get performs a Get request against the Redfish service and save etag
func (e *Entity) Get(c Client, uri string, payload interface{}) error {
	resp, err := c.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(payload)
	if err != nil {
		return err
	}

	if resp.Header["Etag"] != nil {
		e.etag = resp.Header["Etag"][0]
	}
	e.SetClient(c)
	return nil
}

// Patch performs a Patch request against the Redfish service with etag
func (e *Entity) Patch(uri string, payload interface{}) error {
	header := make(map[string]string)
	if e.etag != "" && !e.disableEtagMatch {
		if e.stripEtagQuotes {
			e.etag = strings.Trim(e.etag, "\"")
		}

		header["If-Match"] = e.etag
	}

	resp, err := e.client.PatchWithHeaders(uri, payload, header)
	if err == nil {
		return resp.Body.Close()
	}
	return err
}

// Post performs a Post request against the Redfish service with etag
func (e *Entity) Post(uri string, payload interface{}) error {
	header := make(map[string]string)
	if e.etag != "" && !e.disableEtagMatch {
		if e.stripEtagQuotes {
			e.etag = strings.Trim(e.etag, "\"")
		}

		header["If-Match"] = e.etag
	}

	resp, err := e.client.PostWithHeaders(uri, payload, header)
	if err == nil {
		return resp.Body.Close()
	}
	return err
}

// PostWithResponse performs a Post request against the Redfish service with etag,
// returning the response from the service.
// Callers should make sure to call `resp.Body.Close()` when done with the response.
func (e *Entity) PostWithResponse(uri string, payload interface{}) (*http.Response, error) {
	header := make(map[string]string)
	if e.etag != "" && !e.disableEtagMatch {
		if e.stripEtagQuotes {
			e.etag = strings.Trim(e.etag, "\"")
		}

		header["If-Match"] = e.etag
	}

	return e.client.PostWithHeaders(uri, payload, header)
}

type Filter string

type FilterOption func(*Filter)

func WithSkip(skipNum int) FilterOption {
	return func(e *Filter) {
		*e = Filter(fmt.Sprintf("%s$skip=%d", *e, skipNum))
	}
}

func WithTop(topNum int) FilterOption {
	return func(e *Filter) {
		*e = Filter(fmt.Sprintf("%s$top=%d", *e, topNum))
	}
}

func (e *Filter) SetFilter(opts ...FilterOption) {
	*e = "?"
	lastIdx := len(opts) - 1
	for idx, opt := range opts {
		opt(e)
		if idx < lastIdx {
			*e += "&"
		}
	}
}

func (e *Filter) ClearFilter() {
	*e = ""
}

func getPatchPayloadFromUpdate(originalEntity, updatedEntity reflect.Value) (payload map[string]interface{}) {
	payload = make(map[string]interface{})

	for i := 0; i < originalEntity.NumField(); i++ {
		if !originalEntity.Field(i).CanInterface() {
			// Private field or something that we can't access
			continue
		}
		field := originalEntity.Type().Field(i)
		if field.Type.Kind() == reflect.Ptr {
			continue
		}
		fieldName := field.Name
		jsonName := field.Tag.Get("json")
		if jsonName == "-" {
			continue
		}
		if jsonName != "" {
			fieldName = jsonName
		}
		originalValue := originalEntity.Field(i).Interface()
		currentValue := updatedEntity.Field(i).Interface()
		if originalValue == nil && currentValue == nil {
			continue
		}

		if originalValue == nil {
			payload[fieldName] = currentValue
			continue
		}

		switch reflect.TypeOf(originalValue).Kind() {
		case reflect.Slice, reflect.Map:
			if !reflect.DeepEqual(originalValue, currentValue) {
				payload[fieldName] = currentValue
			}
		case reflect.Struct:
			structPayload := getPatchPayloadFromUpdate(originalEntity.Field(i), updatedEntity.Field(i))
			if field.Anonymous {
				for k, v := range structPayload {
					payload[k] = v
				}
			} else if len(structPayload) != 0 {
				payload[fieldName] = structPayload
			}
		default:
			if originalValue != currentValue {
				payload[fieldName] = currentValue
			}
		}
	}
	return payload
}

type SchemaObject interface {
	SetClient(Client)
	SetETag(string)
}

// GetObject retrieves an API object from the service.
func GetObject[T any, PT interface {
	*T
	SchemaObject
}](c Client, uri string) (*T, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	entity := PT(new(T))
	err = json.NewDecoder(resp.Body).Decode(&entity)
	if err != nil {
		return nil, err
	}

	if resp.Header["Etag"] != nil {
		entity.SetETag(resp.Header["Etag"][0])
	}
	entity.SetClient(c)
	return entity, nil
}

// GetObject retrieves an API object from the service.
func GetObjects[T any, PT interface {
	*T
	SchemaObject
}](c Client, uris []string) ([]*T, error) {
	var result []*T
	if len(uris) == 0 {
		return result, nil
	}

	type GetResult struct {
		Item  *T
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := NewCollectionError()
	get := func(link string) {
		entity, err := GetObject[T, PT](c, link)
		ch <- GetResult{Item: entity, Link: link, Error: err}
	}

	go func() {
		CollectCollection(get, uris)
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
