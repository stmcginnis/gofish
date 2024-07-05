//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ClientType string

const (
	// MonitorClientType The registered client only performs read operations on this service.
	MonitorClientType ClientType = "Monitor"
	// ConfigureClientType The registered client performs update, create, and delete operations on the resources listed
	// in the ManagedResources property as well as read operations on the service.
	ConfigureClientType ClientType = "Configure"
)

// ManagedResource shall contain information about a resource managed by a client. The managed resource may specify
// subordinate resources.
type ManagedResource struct {
	// IncludesSubordinates shall indicate whether the subordinate resources of the managed resource referenced by the
	// ManagedResourceURI property are also managed by the registered client. If not specified, the value is assumed to
	// be 'false' unless ManagedResourceURI references a resource collection.
	IncludesSubordinates bool
	// ManagedResourceURI shall contain the URI of the Redfish resource or Redfish resource collection managed by the
	// registered client. When the URI references a resource collection, all members of the resource collection may be
	// monitored or configured by the client, and the IncludesSubordinates property shall contain 'true'.
	ManagedResourceURI string
	// PreferExclusive shall indicate whether the registered client expects to have exclusive access to the managed
	// resource referenced by the ManagedResourceURI property, and also its subordinate resources if
	// IncludesSubordinates contains 'true'. If not specified, the value is assumed to be 'false'.
	PreferExclusive bool
}

// RegisteredClient shall represent a registered client for a Redfish implementation. It is not expected that
// transient tools, such as a short-lived CLI tool, register. Clients and management tools that live for long
// periods of time can create RegisteredClient resources so that other clients are aware the service might be
// configured or monitored by the client.
type RegisteredClient struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClientType shall contain the type of registered client.
	ClientType ClientType
	// ClientURI shall contain the URI of the registered client.
	ClientURI string
	// Context shall contain data provided by the owning client used to identify the service, provide context about its
	// state, or other information. The value of this property shall not contain unencrypted sensitive data such as
	// user credentials. Services shall support values of at least 256 bytes in length.
	Context string
	// CreatedDate shall contain the date and time when the client entry was created.
	CreatedDate string
	// Description provides a description of this resource.
	Description string
	// ExpirationDate shall contain the date and time when the client entry expires. Registered clients that are
	// actively managing or monitoring should periodically update this value. The value should not be more than 7 days
	// after the date when it was last set. If the current date is beyond this date, the service may delete this client
	// entry.
	ExpirationDate string
	// ManagedResources shall contain an array of resources that the registered client monitors or configures. Other
	// clients can use this property to understand which resources are monitored or configured by the registered
	// client.
	ManagedResources []ManagedResource
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SubContext shall contain additional data provided by the owning client used to identify the service, provide
	// context about its state, or other information. The value of this property shall not contain unencrypted
	// sensitive data such as user credentials. Services shall support values of at least 256 bytes in length.
	SubContext string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a RegisteredClient object from the raw JSON.
func (registeredclient *RegisteredClient) UnmarshalJSON(b []byte) error {
	type temp RegisteredClient
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*registeredclient = RegisteredClient(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	registeredclient.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (registeredclient *RegisteredClient) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(RegisteredClient)
	original.UnmarshalJSON(registeredclient.rawData)

	readWriteFields := []string{
		"ClientType",
		"ClientURI",
		"Context",
		"ExpirationDate",
		"SubContext",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(registeredclient).Elem()

	return registeredclient.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRegisteredClient will get a RegisteredClient instance from the service.
func GetRegisteredClient(c common.Client, uri string) (*RegisteredClient, error) {
	return common.GetObject[RegisteredClient](c, uri)
}

// ListReferencedRegisteredClients gets the collection of RegisteredClient from
// a provided reference.
func ListReferencedRegisteredClients(c common.Client, link string) ([]*RegisteredClient, error) {
	return common.GetCollectionObjects[RegisteredClient](c, link)
}
