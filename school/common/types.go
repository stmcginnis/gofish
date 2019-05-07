// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"encoding/json"
	"net/http"
)

// DefaultServiceRoot is the default path to the Redfish service endpoint.
const DefaultServiceRoot = "/redfish/v1/"

// Client is a connection to a Redfish service.
type Client interface {
	Get(url string) (*http.Response, error)
	// Post()
	// Patch()
	// Put()
	// Delete()
}

// Entity provides the common basis for all Redfish and Swordfish objects.
type Entity struct {
	ID     string `json:"Id"`
	Name   string `json:"Name"`
	Client Client
}

// SetClient sets the API client connection to use for accessing this
// entity.
func (e *Entity) SetClient(c Client) {
	e.Client = c
}

// Link is an OData link reference
type Link string

// UnmarshalJSON unmarshals a Link
func (l *Link) UnmarshalJSON(b []byte) error {
	var t struct {
		ODataID string `json:"@odata.id"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		*l = ""
	}

	*l = Link(t.ODataID)
	return nil
}

// Links are a collection of Link references
type Links []Link

// ToStrings converts a Link collection to a list of strings
func (l Links) ToStrings() []string {
	var result []string
	for _, link := range l {
		result = append(result, string(link))
	}
	return result
}

// LinksCollection contains links to other entities
type LinksCollection struct {
	Count   int   `json:"Members@odata.count"`
	Members Links `json:"Members"`
}

// ToStrings will extract the URI for all linked entities.
func (l LinksCollection) ToStrings() []string {
	return l.Members.ToStrings()
}

// Health indicates the health of a resource.
type Health string

const (
	// OKHealth indicates the health is normal.
	OKHealth Health = "OK"
	// WarningHealth indicates a condition exists that requires attention.
	WarningHealth Health = "Warning"
	// CriticalHealth indicates a critical condition exists that requires
	// immediate attention.
	CriticalHealth Health = "Critical"
)

// DurableNameFormat indicates the type of durable name.
type DurableNameFormat string

const (
	// NAADurableNameFormat shall be a hexadecimal representation of the Name Address
	// Authority structure as defined in the T11 Fibre Channel - Framing and Signaling
	// - 3 (FC-FS-3) specification.
	NAADurableNameFormat DurableNameFormat = "NAA"
	// IQNDurableNameFormat shall be in the iSCSI Qualified Name format as defined
	// in RFC 3720 and RFC 3721.
	IQNDurableNameFormat DurableNameFormat = "iQN"
	// FCWWNDurableNameFormat shall be a hexadecimal representation of the World
	// Wide Name format as defined in the T11 Fibre Channel Physical and Signaling
	// Interface Specification.
	FCWWNDurableNameFormat DurableNameFormat = "FC_WWN"
	// UUIDDurableNameFormat shall be the hexadecimal representation of the Universal
	// Unique Identifier as defined in the Internation Telecom Union's OSI networking
	// and system aspects - Naming, Addressing and Registration Specification.
	UUIDDurableNameFormat DurableNameFormat = "UUID"
	// EUIDurableNameFormat shall be the hexadecimal representation of the IEEE-defined
	// 64-bit Extended Unique Identifier as defined in the IEEE's Guidelines for
	// 64-bit Global Identifier (EUI-64) Specification.
	EUIDurableNameFormat DurableNameFormat = "EUI"
	// NQNDurableNameFormat shall be in the NVMe Qualified Name format as defined
	// in the NVN Express over Fabric Specification.
	NQNDurableNameFormat DurableNameFormat = "NQN"
	// NSIDDurableNameFormat shall be in the NVM Namespace Identifier format as
	// defined in the NVN Express Specification.
	NSIDDurableNameFormat DurableNameFormat = "NSID"
)

// Identifier shall contain any additional identifiers of a resource.
type Identifier struct {
	// DurableName indicates the world wide, persistent name of the resource.
	DurableName string
	// DurableNameFormat
	DurableNameFormat DurableNameFormat
}

// IndicatorLED represents LED indicator states
type IndicatorLED string

const (
	// UnknownIndicatorLED indicates the state of the Indicator LED cannot be
	// determined.
	UnknownIndicatorLED IndicatorLED = "Unknown"
	// LitIndicatorLED indicates the Indicator LED is lit.
	LitIndicatorLED IndicatorLED = "Lit"
	// BlinkingIndicatorLED indicates the Indicator LED is blinking.
	BlinkingIndicatorLED IndicatorLED = "Blinking"
	// OffIndicatorLED indicates the Indicator LED is off.
	OffIndicatorLED IndicatorLED = "Off"
)

// State indicates the known state of the resource, such as if it is enabled.
type State string

const (
	// EnabledState indicates this function or resource has been enabled.
	EnabledState State = "Enabled"
	// DisabledState Stateindicates this function or resource has been disabled.
	DisabledState State = "Disabled"
	// StandbyOfflineState indicates this function or resource is enabled,
	// but awaiting an external action to activate it.
	StandbyOfflineState State = "StandbyOffinline"
	// StandbySpareState indicates this function or resource is part of a
	// redundancy set and is awaiting a failover or other external action to
	// activate it.
	StandbySpareState State = "StandbySpare"
	// InTestState indicates this function or resource is undergoing testing.
	InTestState State = "InTest"
	// StartingState indicates this function or resource is starting.
	StartingState State = "Starting"
	// AbsentState indicates this function or resource is not present or not
	// detected.
	AbsentState State = "Absent"
	// UnavailableOfflineState indicates this function or resource is present
	// but cannot be used.
	UnavailableOfflineState State = "UnavailableOffline"
	// DeferringState indicates the element will not process any commands but
	// will queue new requests.
	DeferringState State = "Deferring"
	// QuiescedState indicates the element is enabled but only processes a
	// restricted set of commands.
	QuiescedState State = "Quiesced"
	// UpdatingState indicates the element is updating and may be unavailable
	// or degraded.
	UpdatingState State = "Updating"
)

// Status describes the status and health of a resource and its children.
type Status struct {
	Health Health `json:"Health"`
	State  State  `json:"State"`
}
