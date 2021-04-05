//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// DefaultServiceRoot is the default path to the Redfish service endpoint.
const DefaultServiceRoot = "/redfish/v1/"

// Client is a connection to a Redfish service.
type Client interface {
	Get(url string) (*http.Response, error)
	GetWithHeaders(url string, customHeaders map[string]string) (*http.Response, error)
	Post(url string, payload interface{}) (*http.Response, error)
	PostWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error)
	PostMultipart(url string, payload map[string]io.Reader) (*http.Response, error)
	PostMultipartWithHeaders(url string, payload map[string]io.Reader, customHeaders map[string]string) (*http.Response, error)
	Patch(url string, payload interface{}) (*http.Response, error)
	PatchWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error)
	Put(url string, payload interface{}) (*http.Response, error)
	PutWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error)
	Delete(url string) (*http.Response, error)
	DeleteWithHeaders(url string, customHeaders map[string]string) (*http.Response, error)
}

// Entity provides the common basis for all Redfish and Swordfish objects.
type Entity struct {
	// ODataID is the location of the resource.
	ODataID string `json:"@odata.id"`
	// ID uniquely identifies the resource.
	ID string `json:"Id"`
	// Name is the name of the resource or array element.
	Name string `json:"Name"`
	// Client is the REST client interface to the system.
	Client Client
}

// SetClient sets the API client connection to use for accessing this
// entity.
func (e *Entity) SetClient(c Client) {
	e.Client = c
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
		originalValue := originalEntity.Field(i).Interface()
		currentValue := currentEntity.Field(i).Interface()
		if originalValue == nil && currentValue == nil {
			continue
		} else if originalValue == nil {
			payload[fieldName] = currentValue
		} else if reflect.TypeOf(originalValue).Kind() != reflect.Map {
			if originalValue != currentValue {
				// TODO: Handle JSON name being different than field name
				payload[fieldName] = currentValue
			}
		} else if !reflect.DeepEqual(originalValue, currentValue) {
			// TODO: Handle JSON name being different than field name
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
		_, err := e.Client.Patch(e.ODataID, payload) // nolint:bodyclose
		if err != nil {
			return err
		}
	}

	return nil
}

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
	// Unique Identifier as defined in the International Telecom Union's OSI networking
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
	// DisabledState State indicates this function or resource has been disabled.
	DisabledState State = "Disabled"
	// StandbyOfflineState indicates this function or resource is enabled,
	// but awaiting an external action to activate it.
	StandbyOfflineState State = "StandbyOffline"
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

// LocationType shall name the type of location in use.
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
	// OpenURackUnits shall be specified in terms of the Open Compute Open
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
	Latitude float32
	// Longitude shall be the longitude of the resource specified in degrees
	// using a decimal format and not minutes or seconds.
	Longitude float32
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
	// Division is used to name a city division, borough, city district,
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
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
}

// ResourceCollection is
type ResourceCollection struct {
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
	// PercentageComplete shall be an integer of the percentage of the
	// operation that has been completed.
	PercentageComplete int
}

// ApplyTime is when to apply a change.
type ApplyTime string

const (

	// ImmediateApplyTime shall be used to indicate the values within the
	// Settings resource are applied immediately.
	ImmediateApplyTime ApplyTime = "Immediate"
	// OnResetApplyTime shall be used to indicate the values within the
	// Settings resource are applied when the system or service is reset.
	OnResetApplyTime ApplyTime = "OnReset"
	// AtMaintenanceWindowStartApplyTime shall be used to indicate the values
	// within the Settings resource are applied during the maintenance window
	// specified by the MaintenanceWindowStartTime and
	// MaintenanceWindowDurationInSeconds properties. A service may perform
	// resets during this maintenance window.
	AtMaintenanceWindowStartApplyTime ApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetApplyTime shall be used to indicate the
	// values within the Settings resource are applied during the maintenance
	// window specified by the MaintenanceWindowStartTime and
	// MaintenanceWindowDurationInSeconds properties, and if a reset occurs
	// within the maintenance window.
	InMaintenanceWindowOnResetApplyTime ApplyTime = "InMaintenanceWindowOnReset"
)

// OperationApplyTime is when to perform the application.
type OperationApplyTime string

const (

	// ImmediateOperationApplyTime shall be used to indicate the requested
	// Create, Delete, or Action operation is applied immediately.
	ImmediateOperationApplyTime OperationApplyTime = "Immediate"
	// OnResetOperationApplyTime shall be used to indicate the requested
	// Create, Delete, or Action operation is applied when the system or
	// service is reset.
	OnResetOperationApplyTime OperationApplyTime = "OnReset"
	// AtMaintenanceWindowStartOperationApplyTime shall be used to indicate
	// the requested Create, Delete, or Action operation is applied during
	// the maintenance window specified by the MaintenanceWindowStartTime and
	// MaintenanceWindowDurationInSeconds properties. A service may perform
	// resets during this maintenance window.
	AtMaintenanceWindowStartOperationApplyTime OperationApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetOperationApplyTime shall be used to indicate
	// the requested Create, Delete, or Action operation is applied during
	// the maintenance window specified by the MaintenanceWindowStartTime and
	// MaintenanceWindowDurationInSeconds properties, and if a reset occurs
	// within the maintenance window.
	InMaintenanceWindowOnResetOperationApplyTime OperationApplyTime = "InMaintenanceWindowOnReset"
)

// MaintenanceWindow shall indicate if a given resource
// has a maintenance window assignment for applying settings or
// operations. Other resources may reference this object in order to
// convey a common control surface for the configuration of the
// maintenance window.
type MaintenanceWindow struct {
	// MaintenanceWindowDurationInSeconds shall
	// indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property.
	MaintenanceWindowDurationInSeconds int
	// MaintenanceWindowStartTime shall
	// indicate the date and time as to when the service is allowed to start
	// applying the requested settings or operation as part of a maintenance
	// window.
	MaintenanceWindowStartTime string
}

// OperationApplyTimeSupport shall specify the support a
// service has for a client to request a specific apply time of a Create,
// Delete, or Action operation of a given resource.
type OperationApplyTimeSupport struct {
	// MaintenanceWindowDurationInSeconds shall
	// be the same as the MaintenanceWindowDurationInSeconds property found
	// in the MaintenanceWindow structure on the MaintenanceWindowResource.
	// This property shall be required if the SupportedValues property
	// contains AtMaintenanceWindowStart or InMaintenanceWindowOnReset.
	MaintenanceWindowDurationInSeconds int
	// MaintenanceWindowResource shall be a
	// reference to a resource that contains the @Redfish.MaintenanceWindow
	// property which governs this resource. This property shall be required
	// if the SupportedValues property contains AtMaintenanceWindowStart or
	// InMaintenanceWindowOnReset.
	MaintenanceWindowResource string
	// MaintenanceWindowStartTime shall be the
	// same as the MaintenanceWindowStartTime property found in the
	// MaintenanceWindow structure on the MaintenanceWindowResource. This
	// property shall be required if the SupportedValues property contains
	// AtMaintenanceWindowStart or InMaintenanceWindowOnReset.
	MaintenanceWindowStartTime string
	// SupportedValues shall indicate the types
	// of apply times the client is allowed request when performing a Create,
	// Delete, or Action operation.
	SupportedValues []OperationApplyTime
}

// PreferredApplyTime shall be specified by client in a request to indicate its
// preference on when to apply the values in this Settings resource.
type PreferredApplyTime struct {
	// ApplyTime shall indicate the preference
	// on to when to apply the values in this Settings resource.
	ApplyTime ApplyTime
	// MaintenanceWindowDurationInSeconds shall
	// indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property.
	// This property shall be required if the ApplyTime property is specified
	// as AtMaintenanceWindowStart or InMaintenanceWindowOnReset.
	MaintenanceWindowDurationInSeconds int
	// MaintenanceWindowStartTime shall
	// indicate the date and time as to when the service is allowed to start
	// applying the future configuration as part of a maintenance window.
	// This property shall be required if the ApplyTime property is specified
	// as AtMaintenanceWindowStart or InMaintenanceWindowOnReset.
	MaintenanceWindowStartTime string
}

// Settings shall describe any attributes of a resource.
type Settings struct {
	// ETag shall be the ETag of the resource to which the settings were
	// applied, after the application. This is here so that the client can check
	// it against the ETag of the current resource to see if any other changes
	// have also happened to the resource.
	ETag string
	// MaintenanceWindowResource shall be a
	// reference to a resource that contains the @Redfish.MaintenanceWindow
	// property which governs this resource. This property should be
	// supported if the SupportedApplyTimes property contains
	// AtMaintenanceWindowStart or InMaintenanceWindowOnReset.
	MaintenanceWindowResource string
	// Messages shall be an array of messages
	// associated with the task.
	Messages []Message
	// SettingsObject shall be the URI of the resource to which a client must do
	// a PUT or PATCH in order to modify this resource.
	SettingsObject Link
	// SupportedApplyTimes is A service shall advertise its applytime
	// capabilities using this property as to when a Setting resource can be
	// applied.
	SupportedApplyTimes []ApplyTime
	// Time shall indicate the time that the settings object was applied to the
	// resource.
	Time string
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

func (e *Error) Error() string {
	if e.HTTPReturnedStatusCode != 0 {
		return fmt.Sprintf("%d: %s", e.HTTPReturnedStatusCode, e.rawData)
	}
	return string(e.rawData)
}

// ErrExtendedInfo is for redfish ExtendedInfo error response
// TODO: support RelatedProperties
type ErrExtendedInfo struct {
	// Indicating a specific error or message (not to be confused with the HTTP status code).
	// This code can be used to access a detailed message from a message registry.
	MessageID string `json:"MessageId"`
	// A human readable error message indicating the semantics associated with the error.
	// This shall be the complete message, and not rely on substitution variables.
	Message string
	// An optional array of strings representing the substitution parameter values for the message.
	// This shall be included in the response if a MessageId is specified for a parameterized message.
	MessageArgs []string
	// An optional string representing the severity of the error.
	Severity string
	// An optional string describing recommended action(s) to take to resolve the error.
	Resolution string
}
