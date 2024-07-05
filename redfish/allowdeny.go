//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AllowType is a set of allow or deny types.
type AllowType string

const (
	// AllowAllowType Indicates that traffic that matches the criteria in this resource is permitted.
	AllowAllowType AllowType = "Allow"
	// DenyAllowType Indicates that traffic that matches the criteria in this resource is not permitted.
	DenyAllowType AllowType = "Deny"
)

// DataDirection is the direction of data flow.
type DataDirection string

const (
	// NoneDataDirection Indicates that this limit not enforced.
	NoneDataDirection DataDirection = "None"
	// IngressDataDirection Indicates that this limit is enforced on packets and bytes received by the network device
	// function.
	IngressDataDirection DataDirection = "Ingress"
	// EgressDataDirection Indicates that this limit is enforced on packets and bytes transmitted by the network device
	// function.
	EgressDataDirection DataDirection = "Egress"
)

// AllowDeny shall represent an AllowDeny resource in a Redfish implementation.
type AllowDeny struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowType shall indicate the type of permission.
	AllowType AllowType
	// Description provides a description of this resource.
	Description string
	// DestinationPortLower shall contain the TCP, UDP, or other destination port to which this rule begins
	// application, inclusive.
	DestinationPortLower int
	// DestinationPortUpper shall contain the TCP, UDP, or other destination port to which this rule ends application,
	// inclusive.
	DestinationPortUpper int
	// Direction shall indicate the direction of the data to which this permission applies for this network device
	// function.
	Direction DataDirection
	// IANAProtocolNumber shall contain the IANA protocol number to which this permission applies.
	IANAProtocolNumber int
	// IPAddressLower shall contain the lower IP address to which this permission applies.
	IPAddressLower string
	// IPAddressType shall contain the type of IP address populated in the IPAddressLower and IPAddressUpper
	// properties. Services shall not permit mixing IPv6 and IPv4 addresses on the same resource.
	IPAddressType IPAddressType
	// IPAddressUpper shall contain the upper IP address to which this permission applies.
	IPAddressUpper string
	// SourcePortLower shall contain the TCP, UDP, or other source port to which this rule begins application,
	// inclusive.
	SourcePortLower int
	// SourcePortUpper shall contain the TCP, UDP, or other source port to which this rule ends application, inclusive.
	SourcePortUpper int
	// StatefulSession shall indicate if this permission only applies to stateful connections, which are those using
	// SYN, ACK, and FIN.
	StatefulSession bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AllowDeny object from the raw JSON.
func (allowdeny *AllowDeny) UnmarshalJSON(b []byte) error {
	type temp AllowDeny
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*allowdeny = AllowDeny(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	allowdeny.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (allowdeny *AllowDeny) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AllowDeny)
	_ = original.UnmarshalJSON(allowdeny.rawData)

	readWriteFields := []string{
		"AllowType",
		"DestinationPortLower",
		"DestinationPortUpper",
		"Direction",
		"IANAProtocolNumber",
		"IPAddressLower",
		"IPAddressType",
		"IPAddressUpper",
		"SourcePortLower",
		"SourcePortUpper",
		"StatefulSession",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(allowdeny).Elem()

	return allowdeny.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAllowDeny will get a AllowDeny instance from the service.
func GetAllowDeny(c common.Client, uri string) (*AllowDeny, error) {
	return common.GetObject[AllowDeny](c, uri)
}

// ListReferencedAllowDenys gets the collection of AllowDeny from
// a provided reference.
func ListReferencedAllowDenys(c common.Client, link string) ([]*AllowDeny, error) {
	return common.GetCollectionObjects[AllowDeny](c, link)
}
