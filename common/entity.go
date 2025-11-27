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
	// ODataID is the location of the resource
	ODataID string `json:"@odata.id"`
	// ID uniquely identifies the resource
	ID string `json:"Id"`
	// Name is the name of the resource or array element
	Name string `json:"Name"`

	// client is the REST client interface to the system
	client Client

	// ODataEtag contains either the entity's declared ETag or the ETag header
	// when fetching the object to control updates and prevent conflicts between
	// concurrent modifications
	ODataEtag string `json:"@odata.etag,omitempty"`

	ExtendedInfo []MessageExtendedInfo `json:"@Message.ExtendedInfo,omitempty"`

	// stripEtagQuotes removes surrounding quotes of etag used in If-Match header
	// Only use for vendor implementations where If-Match only matches unquoted etag
	stripEtagQuotes bool

	// disableEtagMatch skips the If-Match header from PATCH and POST requests
	// Workaround for vendor implementations where If-Match header doesn't work properly
	disableEtagMatch bool
}

func (e *Entity) GetID() string {
	return e.ID
}

func (e *Entity) GetODataID() string {
	return e.ODataID
}

func (e *Entity) GetExtendedInfo() []MessageExtendedInfo {
	return e.ExtendedInfo
}

// SetClient sets the API client connection for accessing this entity.
func (e *Entity) SetClient(c Client) {
	e.client = c
}

// SetETag sets the etag value of this API object.
func (e *Entity) SetETag(tag string) {
	e.ODataEtag = tag
}

func (e *Entity) GetETag() string {
	return e.ODataEtag
}

// GetClient returns the API client connection for this entity.
func (e *Entity) GetClient() Client {
	return e.client
}

// StripEtagQuotes enables/disables stripping etag quotes.
func (e *Entity) StripEtagQuotes(b bool) {
	e.stripEtagQuotes = b
}

// DisableEtagMatch enables/disables the etag match header.
func (e *Entity) DisableEtagMatch(b bool) {
	e.disableEtagMatch = b
}

// IsEtagMatchDisabled indicates if etag matching is disabled for this entity.
func (e *Entity) IsEtagMatchDisabled() bool {
	return e.disableEtagMatch
}

// Update commits changes to an entity after validating allowed updates.
func (e *Entity) Update(originalEntity, updatedEntity reflect.Value, allowedUpdates []string) error {
	payload := getPatchPayloadFromUpdate(originalEntity, updatedEntity)

	// Validate that all fields being updated are allowed
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

	// Send updates if there are any allowed changes
	if len(payload) > 0 {
		return e.Patch(e.ODataID, payload)
	}

	return nil
}

// Get performs a GET request against the Redfish service and saves the etag.
func (e *Entity) Get(c Client, uri string, payload any) error {
	resp, err := c.Get(uri)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		return err
	}

	// if the entity already has an etag, don't override it, but otherwise pull from the HTTP header
	if etag := resp.Header.Get("Etag"); etag != "" && e.ODataEtag == "" {
		e.ODataEtag = etag
	}
	e.SetClient(c)

	return nil
}

// Patch performs a PATCH request against the Redfish service with etag headers.
func (e *Entity) Patch(uri string, payload any) error {
	resp, err := e.client.PatchWithHeaders(uri, payload, e.Headers())
	if err != nil {
		return err
	}
	return CleanupHTTPResponse(resp)
}

// Post performs a POST request against the Redfish service with etag headers.
func (e *Entity) Post(uri string, payload any) error {
	resp, err := e.PostWithResponse(uri, payload)
	if err != nil {
		return err
	}
	return CleanupHTTPResponse(resp)
}

// PostWithResponse performs a POST request and returns the full response.
// Callers must close the response body when done.
func (e *Entity) PostWithResponse(uri string, payload any) (*http.Response, error) {
	return e.client.PostWithHeaders(uri, payload, e.Headers())
}

// Headers generates the appropriate Headers including etag if configured.
func (e *Entity) Headers() map[string]string {
	header := make(map[string]string)
	if e.ODataEtag != "" && !e.disableEtagMatch {
		if e.stripEtagQuotes {
			e.ODataEtag = strings.Trim(e.ODataEtag, `"`)
		}
		header["If-Match"] = e.ODataEtag
	}
	return header
}

// Filter represents query filter options for API requests.
type Filter string

// FilterOption defines functions that can modify Filter settings.
type FilterOption func(*Filter)

// WithSkip adds a $skip parameter to the filter.
func WithSkip(skipNum int) FilterOption {
	return func(e *Filter) {
		*e = Filter(fmt.Sprintf("%s$skip=%d", *e, skipNum))
	}
}

// WithTop adds a $top parameter to the filter.
func WithTop(topNum int) FilterOption {
	return func(e *Filter) {
		*e = Filter(fmt.Sprintf("%s$top=%d", *e, topNum))
	}
}

// SetFilter configures the filter with the provided options.
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

// ClearFilter resets the filter to empty.
func (e *Filter) ClearFilter() {
	*e = ""
}

// UpdateFromRawData provides a generic update implementation for resources
// that store their original JSON data in a RawData field.
func (e *Entity) UpdateFromRawData(resource any, rawData []byte, allowedUpdates []string) error {
	if e == nil {
		return fmt.Errorf("entity is nil")
	}
	if resource == nil {
		return fmt.Errorf("resource is nil")
	}
	if len(rawData) == 0 {
		return fmt.Errorf("rawData is empty")
	}

	resourceType := reflect.TypeOf(resource)
	if resourceType == nil {
		return fmt.Errorf("resource type is nil")
	}

	if resourceType.Kind() == reflect.Ptr {
		resourceType = resourceType.Elem()
	}

	original := reflect.New(resourceType).Interface()
	if original == nil {
		return fmt.Errorf("failed to create original instance")
	}

	// Handle custom unmarshalers if implemented
	if unmarshaler, ok := original.(json.Unmarshaler); ok {
		if err := unmarshaler.UnmarshalJSON(rawData); err != nil {
			return fmt.Errorf("custom unmarshal failed: %w", err)
		}
	} else if err := json.Unmarshal(rawData, original); err != nil {
		return fmt.Errorf("standard unmarshal failed: %w", err)
	}

	originalValue := reflect.ValueOf(original)
	if !originalValue.IsValid() {
		return fmt.Errorf("invalid original value")
	}

	currentValue := reflect.ValueOf(resource)
	if !currentValue.IsValid() {
		return fmt.Errorf("invalid current value")
	}

	if originalValue.Kind() == reflect.Ptr {
		originalValue = originalValue.Elem()
	}
	if currentValue.Kind() == reflect.Ptr {
		currentValue = currentValue.Elem()
	}

	if !originalValue.IsValid() || !currentValue.IsValid() {
		return fmt.Errorf("invalid dereferenced values")
	}

	// Recover from any panics during update
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic in Update: %v\n", r)
		}
	}()

	return e.Update(originalValue, currentValue, allowedUpdates)
}

// getPatchPayloadFromUpdate compares original and updated structures and returns
// a map containing only the changed fields between them. Handles nested structs,
// slices, maps, and respects JSON tags while ignoring private fields.
// getPatchPayloadFromUpdate compares original and updated structures and returns
// a map containing only the changed fields between them.
func getPatchPayloadFromUpdate(original, updated reflect.Value) map[string]any {
	payload := make(map[string]any)

	if !isValidForComparison(original, updated) {
		return payload
	}

	original, updated = derefPointers(original, updated)

	if !isValidStructPair(original, updated) {
		return payload
	}

	return compareStructFields(original, updated)
}

// isValidForComparison checks if the input values are valid for comparison.
func isValidForComparison(original, updated reflect.Value) bool {
	return original.IsValid() && updated.IsValid()
}

// derefPointers dereferences pointer values if needed and returns the dereferenced values.
func derefPointers(original, updated reflect.Value) (derefOriginal, derefUpdated reflect.Value) {
	derefOriginal = original
	derefUpdated = updated

	if original.Kind() == reflect.Ptr {
		derefOriginal = original.Elem()
	}
	if updated.Kind() == reflect.Ptr {
		derefUpdated = updated.Elem()
	}
	return
}

// isValidStructPair checks if both values are structs of the same type.
func isValidStructPair(original, updated reflect.Value) bool {
	return original.Kind() == reflect.Struct &&
		updated.Kind() == reflect.Struct &&
		original.Type() == updated.Type()
}

// compareStructFields compares all fields of two structs and returns differences.
func compareStructFields(original, updated reflect.Value) map[string]any {
	payload := make(map[string]any)

	for i := 0; i < original.NumField(); i++ {
		field := original.Type().Field(i)
		if shouldSkipField(&field) {
			continue
		}

		fieldName := getFieldName(&field)
		originalField := original.Field(i)
		updatedField := updated.Field(i)

		if field.Anonymous {
			addEmbeddedFields(payload, originalField, updatedField)
			continue
		}

		compareField(payload, fieldName, originalField, updatedField)
	}

	return payload
}

// shouldSkipField determines if a field should be skipped during comparison.
func shouldSkipField(field *reflect.StructField) bool {
	// Skip private fields (those with PkgPath set)
	return field.PkgPath != "" && !field.Anonymous
}

// getFieldName extracts the JSON field name from the struct tag.
func getFieldName(field *reflect.StructField) string {
	fieldName := field.Name
	jsonTag := field.Tag.Get("json")

	if jsonTag == "-" {
		return ""
	}

	if jsonTag != "" {
		if jsonTag == ",omitempty" {
			return fieldName
		}
		if name := strings.Split(jsonTag, ",")[0]; name != "" {
			return name
		}
	}

	return fieldName
}

// addEmbeddedFields handles comparison of embedded struct fields.
func addEmbeddedFields(payload map[string]any, original, updated reflect.Value) {
	embeddedPayload := getPatchPayloadFromUpdate(original, updated)
	for k, v := range embeddedPayload {
		payload[k] = v
	}
}

// compareField compares individual fields and adds differences to the payload.
func compareField(payload map[string]any, fieldName string, original, updated reflect.Value) {
	if fieldName == "" {
		return
	}

	switch original.Kind() {
	case reflect.Struct:
		handleStructField(payload, fieldName, original, updated)
	case reflect.Ptr:
		handlePointerField(payload, fieldName, original, updated)
	case reflect.Slice, reflect.Map:
		handleCompositeField(payload, fieldName, original, updated)
	default:
		handleSimpleField(payload, fieldName, original, updated)
	}
}

// handleStructField handles comparison of struct fields.
func handleStructField(payload map[string]any, fieldName string, original, updated reflect.Value) {
	nestedPayload := getPatchPayloadFromUpdate(original, updated)
	if len(nestedPayload) > 0 {
		payload[fieldName] = nestedPayload
	}
}

// handlePointerField handles comparison of pointer fields.
func handlePointerField(payload map[string]any, fieldName string, original, updated reflect.Value) {
	if original.IsNil() && updated.IsNil() {
		return
	}
	if original.IsNil() || updated.IsNil() {
		if !reflect.DeepEqual(original.Interface(), updated.Interface()) {
			payload[fieldName] = updated.Interface()
		}
		return
	}
	compareField(payload, fieldName, original.Elem(), updated.Elem())
}

// handleCompositeField handles comparison of slice and map fields.
func handleCompositeField(payload map[string]any, fieldName string, original, updated reflect.Value) {
	if !reflect.DeepEqual(original.Interface(), updated.Interface()) {
		payload[fieldName] = updated.Interface()
	}
}

// handleSimpleField handles comparison of basic type fields.
func handleSimpleField(payload map[string]any, fieldName string, original, updated reflect.Value) {
	if !reflect.DeepEqual(original.Interface(), updated.Interface()) {
		payload[fieldName] = updated.Interface()
	}
}

// SchemaObject defines the minimum interface required for API objects.
type SchemaObject interface {
	SetClient(Client)
	GetETag() string
	SetETag(string)
	GetID() string
	GetODataID() string
	GetExtendedInfo() []MessageExtendedInfo
}

type GenericSchemaObjectPointer[T any] interface {
	*T
	SchemaObject
}

// GetObject retrieves a single API object from the service.
func GetObject[T any, PT GenericSchemaObjectPointer[T]](c Client, uri string, opts ...QueryGroupOption) (*T, error) {
	uri = BuildQuery(c, uri, false, opts...)

	resp, err := c.Get(uri)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	return DecodeGenericEntity[T, PT](c, resp)
}

// DecodeGenericEntity attempts to decode an HTTP response into an Entity struct
func DecodeGenericEntity[T any, PT GenericSchemaObjectPointer[T]](c Client, resp *http.Response) (*T, error) {
	entity := PT(new(T))
	if err := json.NewDecoder(resp.Body).Decode(entity); err != nil {
		return nil, err
	}

	if etag := resp.Header.Get("Etag"); etag != "" && entity.GetETag() == "" {
		entity.SetETag(etag)
	}
	entity.SetClient(c)
	return entity, nil
}

// GetObjects retrieves multiple API objects concurrently from the service.
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

	// Worker function to get a single object
	get := func(link string) {
		entity, err := GetObject[T, PT](c, link)
		ch <- GetResult{Item: entity, Link: link, Error: err}
	}

	// Start workers for each URI
	go func() {
		CollectCollection(get, uris)
		close(ch)
	}()

	// Process results
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
