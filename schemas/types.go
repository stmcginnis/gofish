//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
)

// DefaultServiceRoot is the default path to the Redfish service endpoint.
const DefaultServiceRoot = "/redfish/v1/"

// Link is an OData link reference
type Link string

// UnmarshalJSON unmarshals a Link
func (l *Link) UnmarshalJSON(b []byte) error {
	var t struct {
		ODataID string `json:"@odata.id"`
		Href    string `json:"href"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		*l = ""
	}

	*l = Link(t.ODataID)
	if *l == "" {
		*l = Link(t.Href)
	}
	return nil
}

func (l Link) String() string {
	return string(l)
}

// Links are a collection of Link references
type Links []Link

// ToStrings converts a Link collection to a list of strings
func (l Links) ToStrings() []string {
	var result []string
	for _, link := range l {
		result = append(result, link.String())
	}
	return result
}

// LinksCollection contains links to other entities
type LinksCollection struct {
	ODataCount int   `json:"@odata.count"`
	Count      int   `json:"Members@odata.count"`
	Members    Links `json:"Members"`
}

// ToStrings will extract the URI for all linked entities.
func (l LinksCollection) ToStrings() []string {
	return l.Members.ToStrings()
}

type Severity string

const (
	// OKSeverity indicates normal operation.
	OKSeverity Severity = "OK"
	// WarningSeverity indicates the condition requires attention.
	WarningSeverity Severity = "Warning"
	// CriticalSeverity indicates the condition requires immediate attention
	CriticalSeverity Severity = "Critical"
)

// Resource is the base type for resources and referenceable members.
type Resource struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM json.RawMessage `json:"Oem"`
}

// ResourceCollection is a group of resources in a list
type ResourceCollection struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM json.RawMessage `json:"Oem"`

	MembersCount    int    `json:"Members@odata.count"`
	MembersNextLink string `json:"Members@odata.nextLink"`
}

type ResourceCollectionGeneric[T SchemaObject] struct {
	ResourceCollection

	Members []T `json:"Members"`
}

// ConstructError tries to create error if body is defined as redfish spec
func ConstructError(statusCode int, b []byte) error {
	var err struct {
		Error *Error
	}
	if e := json.Unmarshal(b, &err); e != nil || err.Error == nil {
		// Construct our own error
		err.Error = new(Error)
		err.Error.Message = string(b)
	}
	err.Error.HTTPReturnedStatusCode = statusCode
	err.Error.rawData = b
	return err.Error
}

// Error is redfish error response object for HTTP status codes different from 200, 201 and 204
type Error struct {
	rawData []byte
	// An integer that represents the status code returned by the API
	HTTPReturnedStatusCode int `json:"-"`
	// A string indicating a specific MessageId from the message registry.
	Code string `json:"code"`
	// A human readable error message corresponding to the message in the message registry.
	Message string `json:"message"`
	// An array of message objects describing one or more error message(s).
	ExtendedInfos []ErrExtendedInfo `json:"@Message.ExtendedInfo"`
}

func (e *Error) UnmarshalJSON(b []byte) error {
	type temp Error
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = Error(tmp.temp)
	e.rawData = b

	return nil
}

func (e *Error) Error() string {
	if e.HTTPReturnedStatusCode != 0 {
		return fmt.Sprintf("%d: %s", e.HTTPReturnedStatusCode, e.rawData)
	}
	return string(e.rawData)
}

// ErrExtendedInfo is for redfish ExtendedInfo error response
type ErrExtendedInfo Message

// ActionTarget is contains the target endpoint for object Actions.
type ActionTarget struct {
	Target string
	// ActionInfoTarget is an optional resource that provides information about parameters
	// that are supported by the associated Target.
	ActionInfoTarget string `json:"@Redfish.ActionInfo"`
}
