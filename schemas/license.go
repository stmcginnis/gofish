//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/License.v1_1_4.json
// 2022.3 - #License.v1_1_4.License

package schemas

import (
	"encoding/json"
)

// AuthorizationScope is This property shall describe the authorization scope
// for this license.
type AuthorizationScope string

const (
	// DeviceAuthorizationScope shall indicate the license authorizes functionality
	// for one or more specific device instances, listed as values of the
	// 'AuthorizedDevices' property.
	DeviceAuthorizationScope AuthorizationScope = "Device"
	// CapacityAuthorizationScope shall indicate the license authorizes
	// functionality for one or more device instances limited to a maximum number
	// of devices specified by the value of the 'MaxAuthorizedDevices' property. In
	// an aggregator, the aggregating service shall represent the applicable
	// services in the 'TargetServices' property in the 'Links' property.
	CapacityAuthorizationScope AuthorizationScope = "Capacity"
	// ServiceAuthorizationScope shall indicate the license authorizes
	// product-level or service-level functionality for a service. This may include
	// hardware or software features not tied to a specific device or subsystem.
	// 'License' resources using this value shall not include the
	// 'AuthorizedDevices' nor the 'MaxAuthorizedDevices' properties. In an
	// aggregator, the aggregating service shall represent the applicable services
	// in the 'TargetServices' property in the 'Links' property.
	ServiceAuthorizationScope AuthorizationScope = "Service"
)

type LicenseOrigin string

const (
	// BuiltInLicenseOrigin is a license was provided with the product.
	BuiltInLicenseOrigin LicenseOrigin = "BuiltIn"
	// InstalledLicenseOrigin is a license installed by user.
	InstalledLicenseOrigin LicenseOrigin = "Installed"
)

type LicenseType string

const (
	// ProductionLicenseType shall indicate a license purchased or obtained for use
	// in production environments.
	ProductionLicenseType LicenseType = "Production"
	// PrototypeLicenseType shall indicate a license that is designed for
	// development or internal use.
	PrototypeLicenseType LicenseType = "Prototype"
	// TrialLicenseType shall indicate a trial version of a license.
	TrialLicenseType LicenseType = "Trial"
)

// License shall represent a license for a Redfish implementation.
type License struct {
	Entity
	// AuthorizationScope shall contain the authorization scope of the license.
	AuthorizationScope AuthorizationScope
	// Contact shall contain an object containing information about the contact of
	// the license.
	Contact ContactInfo
	// DownloadURI shall contain the URI from which to download the license file,
	// using the Redfish protocol and authentication methods. The service provides
	// this URI for the download of the OEM-specific binary file of license data.
	// An HTTP 'GET' from this URI shall return a response payload of MIME type
	// 'application/octet-stream'.
	DownloadURI string
	// EntitlementID shall contain the entitlement identifier for this license,
	// used to display a license key, partial license key, or other value used to
	// identify or differentiate license instances.
	EntitlementID string `json:"EntitlementId"`
	// ExpirationDate shall contain the date and time when the license expires.
	ExpirationDate string
	// GracePeriodDays shall contain the number of days that the license is still
	// usable after the date and time specified by the 'ExpirationDate' property.
	GracePeriodDays *int `json:",omitempty"`
	// InstallDate shall contain the date and time when the license was installed.
	InstallDate string
	// LicenseInfoURI shall contain the URI at which to provide more information
	// about the license. The information provided at the URI is intended to be
	// general product-related and not tied to a specific user, customer, or
	// license instance.
	LicenseInfoURI string
	// LicenseOrigin shall contain the origin for the license.
	LicenseOrigin LicenseOrigin
	// LicenseString shall contain the Base64-encoded string, with padding
	// characters, of the license. This property shall not appear in response
	// payloads.
	LicenseString string
	// LicenseType shall contain the type for the license.
	LicenseType LicenseType
	// Manufacturer shall represent the name of the manufacturer or producer of
	// this license.
	Manufacturer string
	// MaxAuthorizedDevices shall contain the maximum number of devices that are
	// authorized by the license. This property shall only be present if the
	// 'AuthorizationScope' property contains the value 'Capacity'.
	MaxAuthorizedDevices *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the
	// license.
	PartNumber string
	// RemainingDuration shall contain the remaining usage duration before the
	// license expires. This property shall only be present for licenses that are
	// based on usage time.
	RemainingDuration string
	// RemainingUseCount shall contain the remaining usage count before the license
	// expires. This property shall only be present for licenses that are based on
	// usage count.
	RemainingUseCount *int `json:",omitempty"`
	// Removable shall indicate whether a user can remove the license with an HTTP
	// 'DELETE' operation.
	Removable bool
	// SKU shall contain the SKU number for this license.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the license.
	SerialNumber string
	// Status shall contain the status of license.
	Status Status
	// authorizedDevices are the URIs for AuthorizedDevices.
	authorizedDevices []string
	// targetServices are the URIs for TargetServices.
	targetServices []string
}

// UnmarshalJSON unmarshals a License object from the raw JSON.
func (l *License) UnmarshalJSON(b []byte) error {
	type temp License
	type lLinks struct {
		AuthorizedDevices Links `json:"AuthorizedDevices"`
		TargetServices    Links `json:"TargetServices"`
	}
	var tmp struct {
		temp
		Links lLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = License(tmp.temp)

	// Extract the links to other entities for later
	l.authorizedDevices = tmp.Links.AuthorizedDevices.ToStrings()
	l.targetServices = tmp.Links.TargetServices.ToStrings()

	return nil
}

// GetLicense will get a License instance from the service.
func GetLicense(c Client, uri string) (*License, error) {
	return GetObject[License](c, uri)
}

// ListReferencedLicenses gets the collection of License from
// a provided reference.
func ListReferencedLicenses(c Client, link string) ([]*License, error) {
	return GetCollectionObjects[License](c, link)
}

// AuthorizedDevices gets the AuthorizedDevices linked resources.
func (l *License) AuthorizedDevices() ([]*Entity, error) {
	return GetObjects[Entity](l.client, l.authorizedDevices)
}

// TargetServices gets the TargetServices linked resources.
func (l *License) TargetServices() ([]*Manager, error) {
	return GetObjects[Manager](l.client, l.targetServices)
}
