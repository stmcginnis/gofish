//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.2 - #AllowDeny.v1_0_3.AllowDeny

package schemas

import (
	"encoding/json"
)

type AllowType string

const (
	// AllowAllowType Indicates that traffic that matches the criteria in this
	// resource is permitted.
	AllowAllowType AllowType = "Allow"
	// DenyAllowType Indicates that traffic that matches the criteria in this
	// resource is not permitted.
	DenyAllowType AllowType = "Deny"
)

type DataDirection string

const (
	// IngressDataDirection Indicates that this limit is enforced on packets and
	// bytes received by the network device function.
	IngressDataDirection DataDirection = "Ingress"
	// EgressDataDirection Indicates that this limit is enforced on packets and
	// bytes transmitted by the network device function.
	EgressDataDirection DataDirection = "Egress"
	// NoneDataDirection Indicates that this limit not enforced.
	NoneDataDirection DataDirection = "None" // From networkdevicefunction.go
)

// AllowDeny shall represent an AllowDeny resource in a Redfish implementation.
type AllowDeny struct {
	Entity
	// AllowType shall indicate the type of permission.
	AllowType AllowType
	// DestinationPortLower shall contain the TCP, UDP, or other destination port
	// to which this rule begins application, inclusive.
	DestinationPortLower *int `json:",omitempty"`
	// DestinationPortUpper shall contain the TCP, UDP, or other destination port
	// to which this rule ends application, inclusive.
	DestinationPortUpper *int `json:",omitempty"`
	// Direction shall indicate the direction of the data to which this permission
	// applies for this network device function.
	Direction DataDirection
	// IANAProtocolNumber shall contain the IANA protocol number to which this
	// permission applies.
	IANAProtocolNumber *int `json:",omitempty"`
	// IPAddressLower shall contain the lower IP address to which this permission
	// applies.
	IPAddressLower string
	// IPAddressType shall contain the type of IP address populated in the
	// 'IPAddressLower' and 'IPAddressUpper' properties. Services shall not permit
	// mixing IPv6 and IPv4 addresses on the same resource.
	IPAddressType IPAddressType
	// IPAddressUpper shall contain the upper IP address to which this permission
	// applies.
	IPAddressUpper string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SourcePortLower shall contain the TCP, UDP, or other source port to which
	// this rule begins application, inclusive.
	SourcePortLower *int `json:",omitempty"`
	// SourcePortUpper shall contain the TCP, UDP, or other source port to which
	// this rule ends application, inclusive.
	SourcePortUpper *int `json:",omitempty"`
	// StatefulSession shall indicate if this permission only applies to stateful
	// connections, which are those using SYN, ACK, and FIN.
	StatefulSession bool
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a AllowDeny object from the raw JSON.
func (a *AllowDeny) UnmarshalJSON(b []byte) error {
	type temp AllowDeny
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AllowDeny(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	a.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AllowDeny) Update() error {
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

	return a.UpdateFromRawData(a, a.RawData, readWriteFields)
}

// GetAllowDeny will get a AllowDeny instance from the service.
func GetAllowDeny(c Client, uri string) (*AllowDeny, error) {
	return GetObject[AllowDeny](c, uri)
}

// ListReferencedAllowDenys gets the collection of AllowDeny from
// a provided reference.
func ListReferencedAllowDenys(c Client, link string) ([]*AllowDeny, error) {
	return GetCollectionObjects[AllowDeny](c, link)
}
