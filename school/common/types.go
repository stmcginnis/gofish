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
	Delete(url string) error
}

// Entity provides the common basis for all Redfish and Swordfish objects.
type Entity struct {
	// ID uniquely identifies the resource.
	ID string `json:"Id"`
	// Name is the name of the resource or array element.
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

// LocationType shall name the type of locatio in use.
type LocationType string

const (
	// SlotLocationType shall be used to indicate the type of PartLocation is
	// of type slot.
	SlotLocationType LocationType = "Slot"
	// BayLocationType shall be used to indicate the type of PartLocation is
	// of type bay.
	BayLocationType LocationType = "Bay"
	// ConnectorLocationType shall be used to indicate the type of
	// PartLocation is of type connector.
	ConnectorLocationType LocationType = "Connector"
	// SocketLocationType shall be used to indicate the type of PartLocation
	// is of type socket.
	SocketLocationType LocationType = "Socket"
)

// Orientation shall name the orientation for the location type ordering in
// determining the LocationOrdinalValue.
type Orientation string

const (
	// FrontToBackOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is front to back.
	FrontToBackOrientation Orientation = "FrontToBack"
	// BackToFrontOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is back to front.
	BackToFrontOrientation Orientation = "BackToFront"
	// TopToBottomOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is top to bottom.
	TopToBottomOrientation Orientation = "TopToBottom"
	// BottomToTopOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is bottom to top.
	BottomToTopOrientation Orientation = "BottomToTop"
	// LeftToRightOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is left to right.
	LeftToRightOrientation Orientation = "LeftToRight"
	// RightToLeftOrientation shall be used to specify the ordering for
	// LocationOrdinalValue is right to left.
	RightToLeftOrientation Orientation = "RightToLeft"
)

// RackUnits shall name the type of rack units in use.
type RackUnits string

const (
	// OpenURackUnits shall be specifie3d in terms of the Open Compute Open
	// Rack specification.
	OpenURackUnits RackUnits = "OpenU"
	// EIA310RackUnits shall be specified as defined by the EIA-310
	// standard.
	EIA310RackUnits RackUnits = "EIA_310"
)

// Reference shall name the reference for the part location.
type Reference string

const (
	// TopReference shall be used to specify the part location is in the top
	// of the unit.
	TopReference Reference = "Top"
	// BottomReference shall be used to specify the part location is in the
	// bottom of the unit.
	BottomReference Reference = "Bottom"
	// FrontReference shall be used to specify the part location is in the
	// front of the unit.
	FrontReference Reference = "Front"
	// RearReference shall be used to specify the part location is in the
	// rear of the unit.
	RearReference Reference = "Rear"
	// LeftReference shall be used to specify the part location is in the
	// left of the unit.
	LeftReference Reference = "Left"
	// RightReference shall be used to specify the part location is in the
	// right of the unit.
	RightReference Reference = "Right"
	// MiddleReference shall be used to specify the part location is in the
	// middle of the unit.
	MiddleReference Reference = "Middle"
)

// ContactInfo is used to obtain more information from an individual or
// organization responsible for this resource.
type ContactInfo struct {
	// ContactName shall contain the name of a person or
	// organization to contact for information about this resource.
	ContactName string
	// EmailAddress shall contain the email address for a person
	// or organization to contact for information about this resource.
	EmailAddress string
	// PhoneNumber shall contain the phone number for a person
	// or organization to contact for information about this resource.
	PhoneNumber string
}

// Location shall describe the location of a resource.
type Location struct {
	// AltitudeMeters is the altitude of the resource in meters.
	AltitudeMeters int
	// Contacts is used to obtain more information from an individual or
	// organization responsible for this resource.
	Contacts []ContactInfo
	// Info shall represent the location of the resource.
	Info string
	// InfoFormat shall represent the format of the Info property.
	InfoFormat string
	// Latitude shall be the latitude of the resource specified
	// in degrees using a decimal format and not minutes or seconds.
	Latitude int
	// Longitude shall be the longitude of the resource specified in degrees
	// using a decimal format and not minutes or seconds.
	Longitude int
	// PartLocation is used to indicate the location within the Placement.
	PartLocation PartLocation
	// Placement shall be a place within the addressed location.
	Placement Placement
	// PostalAddress shall be a postal address of the resource.
	PostalAddress PostalAddress
}

// PartLocation is used to indicate the location within the Placement.
type PartLocation struct {
	// LocationOrdinalValue shall be the number that represents the location of
	// the part based on the LocationType. LocationOrdinalValue shall be
	// measured based on the Orientation value starting with 0.
	LocationOrdinalValue int
	// LocationType shall be a LocationType enumeration literal indicating the
	// type of rack units in use.
	LocationType LocationType
	// Orientation is used by the LocationOrdinalValue property.
	Orientation Orientation
	// Reference shall be a Reference enumeration literal indicating the general
	// location within the unit of the part.
	Reference Reference
	// ServiceLabel shall be the label assigned for service at the part location.
	ServiceLabel string
}

// Placement shall describe a location within a resource. Examples include a
// shelf in a rack.
type Placement struct {
	// AdditionalInfo is used to describe a location that cannot be conveyed
	// with other properties defined for the Placement object.
	AdditionalInfo string
	// Rack shall be the name of the rack within a row.
	Rack string
	// RackOffset is the vertical location of the item in the rack. Rack offset
	// units shall be measured from bottom to top starting with 0.
	RackOffset int
	// RackOffsetUnits shall be a RackUnit enumeration literal indicating the
	// type of rack units in use.
	RackOffsetUnits RackUnits
	// Row shall be the name of the row.
	Row string
}

// PostalAddress shall describe a postal address for a resource. For more
// information see RFC5139. Depending on use, the instance may represent a past,
// current, or future location.
type PostalAddress struct {
	// AdditionalCode shall conform the requirements of the ADDCODE field as
	// defined in RFC5139.
	AdditionalCode string
	// AdditionalInfo is used to provide additional information.
	AdditionalInfo string
	// Building is used to locate the resource.
	Building string
	// City is used to name a city, township, or shi (JP).
	City string
	// Community shall conform to the requirements of the PCN field
	// as defined in RFC5139. The value shall be a postal community name.
	Community string
	// Country shall conform the requirements of the Country
	// field as defined in RFC5139.
	Country int
	// District is used to name a county, parish, gun (JP), or district
	// (IN).
	District string
	// Division is used to name a city division, borough, dity district,
	// ward, chou (JP).
	Division string
	// Floor is used to provide a floor designation.
	Floor string
	// GPSCoords shall conform the requirements of the ADDCODE
	// field as defined in RFC5139. The value shall be the GPS coordinates of
	// the location. If furnished, this shall be expressed in the format
	// '[-][nn]n.nnnnnn, [-][nn]n.nnnnn', i.e. two numbers, either positive
	// or negative, with six decimal places of precision, comma-separated.
	GPSCoords string
	// HouseNumber shall conform the requirements of the HNO
	// field as defined in RFC5139. It is the numeric portion of the house
	// number.
	HouseNumber int
	// HouseNumberSuffix is used to provide a suffix to a house number, (F,
	// B, 1/2).
	HouseNumberSuffix string
	// Landmark is used to identify a landmark or vanity address.
	Landmark string
	// LeadingStreetDirection is used to name a leading street direction, (N,
	// W, SE).
	LeadingStreetDirection string
	// Location is used to provide additional information.
	Location string
	// Name is the name of this resource.
	Name string
	// Neighborhood is used to name a neighborhood or block.
	Neighborhood string
	// POBox shall conform the requirements of the POBOX field
	// as defined in RFC5139. The value shall be a Post office box (P.O.
	// box).
	POBox string
	// PlaceType shall conform the requirements of the PLC field
	// as defined in RFC5139. Examples include: office, residence,...).
	PlaceType string
	// PostalCode shall conform the requirements of the PC field
	// as defined in RFC5139. The value shall be a Postal code (or zip code).
	PostalCode string
	// Road shall conform the requirements of the RD field as
	// defined in RFC5139. The value designates a primary road or street.
	Road string
	// RoadBranch shall conform the requirements of the RDBR
	// field as defined in RFC5139. The value shall be a Post office box
	// (P.O. box)road branch.
	RoadBranch string
	// RoadPostModifier shall conform the requirements of the
	// POM field as defined in RFC5139. (Extended).
	RoadPostModifier string
	// RoadPreModifier shall conform the requirements of the PRM
	// field as defined in RFC5139. (Old, New).
	RoadPreModifier string
	// RoadSection shall conform the requirements of the RDSEC
	// field as defined in RFC5139. The value shall be a road section.
	RoadSection string
	// RoadSubBranch shall conform the requirements of the
	// RDSUBBR field as defined in RFC5139.
	RoadSubBranch string
	// Room is used to locate the resource within the unit.
	Room string
	// Seat shall conform the requirements of the SEAT field as
	// defined in RFC5139. The value shall be a name or number of a Seat
	// (desk, cubicle, workstation).
	Seat string
	// Street is used to name a street.
	Street string
	// StreetSuffix is used to name a street suffix.
	StreetSuffix string
	// Territory is used to name a territory, state, region, province, or
	// prefecture within a country.
	Territory string
	// TrailingStreetSuffix is used to name a trailing street suffix.
	TrailingStreetSuffix string
	// Unit is used to locate the resource.
	Unit string
}

// ReferenceableMember array members can be referenced using the value
// returned in the @odata.id property which may or may not be a
// dereferenceable URL. The @odata.id of this entity shall be the
// location of this element within an Item.
type ReferenceableMember struct {
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// MemberId shall uniquely identify the member within the collection.
	MemberID string
}

// Resource is the base type for resources and referenceable members.
type Resource struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
}

// ResourceCollection is
type ResourceCollection struct {
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Name is the name of this resource
	Name string
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM string `json:"Oem"`
}

// Operations shall describe a currently running operation on the resource.
type Operations struct {
	// AssociatedTask shall be a reference to a resource of type Task that
	// represents the task associated with the operation.
	AssociatedTask string
	// OperationName shall be a string of the name of the operation.
	OperationName string
	// PercentageComplete shall be an interger of the percentage of the
	// operation that has been completed.
	PercentageComplete int
}
