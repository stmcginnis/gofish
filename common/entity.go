//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
	"reflect"
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
}

// SetClient sets the API client connection to use for accessing this
// entity.
func (e *Entity) SetClient(c Client) {
	e.client = c
}

// GetClient get the API client connection to use for accessing this
// entity.
func (e *Entity) GetClient() Client {
	return e.client
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
	if e.etag != "" {
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
	if e.etag != "" {
		header["If-Match"] = e.etag
	}

	resp, err := e.client.PostWithHeaders(uri, payload, header)
	if err == nil {
		return resp.Body.Close()
	}
	return err
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
