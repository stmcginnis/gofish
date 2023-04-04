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
func (e *Entity) Update(originalEntity, currentEntity reflect.Value, allowedUpdates []string) error {
	payload := make(map[string]interface{})

	for i := 0; i < originalEntity.NumField(); i++ {
		if !originalEntity.Field(i).CanInterface() {
			// Private field or something that we can't access
			continue
		}
		fieldType := originalEntity.Type().Field(i).Type.Kind()
		if fieldType == reflect.Struct || fieldType == reflect.Ptr || fieldType == reflect.Slice {
			// TODO: Handle more complicated data types
			continue
		}
		fieldName := originalEntity.Type().Field(i).Name
		jsonName := originalEntity.Type().Field(i).Tag.Get("json")
		if jsonName != "" {
			fieldName = jsonName
		}
		originalValue := originalEntity.Field(i).Interface()
		currentValue := currentEntity.Field(i).Interface()
		if originalValue == nil && currentValue == nil {
			continue
		} else if originalValue == nil {
			payload[fieldName] = currentValue
		} else if reflect.TypeOf(originalValue).Kind() != reflect.Map {
			if originalValue != currentValue {
				payload[fieldName] = currentValue
			}
		} else if !reflect.DeepEqual(originalValue, currentValue) {
			payload[fieldName] = currentValue
		}
	}

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
