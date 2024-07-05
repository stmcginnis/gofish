//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// LicenseService shall represent a license service and the properties that affect the service itself for a Redfish
// implementation.
type LicenseService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// LicenseExpirationWarningDays shall contain the number of days prior to a license expiration that the service
	// shall send the 'DaysBeforeExpiration' message from the License Message Registry at least once. A value of zero
	// shall indicate that no warning messages are sent prior to license expiration.
	LicenseExpirationWarningDays int
	// Licenses shall contain a link to a resource collection of type LicenseCollection. When installing a license with
	// a POST operation to this collection, the service may update an existing License resource instead of creating a
	// new resource. In these cases, the service shall respond with the HTTP '200 OK' status code or HTTP '204 No
	// Content' status code and the 'Location' header in the response shall contain the URI of the updated License
	// resource.
	licenses string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	installTarget string
}

// UnmarshalJSON unmarshals a LicenseService object from the raw JSON.
func (licenseservice *LicenseService) UnmarshalJSON(b []byte) error {
	type temp LicenseService
	var t struct {
		temp
		Licenses common.Link
		Actions  struct {
			Install common.ActionTarget `json:"#LicenseService.Install"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*licenseservice = LicenseService(t.temp)

	// Extract the links to other entities for later
	licenseservice.licenses = t.Licenses.String()
	licenseservice.installTarget = t.Actions.Install.Target

	// This is a read/write object, so we need to save the raw object data for later
	licenseservice.rawData = b

	return nil
}

type InstallLicenseParameters struct {
	// AuthorizedDevices (optional) parameter shall contain an array of links to the devices to be authorized
	// by the license. Clients can provide this parameter when installing a license to apply the license to
	// specific devices. If not provided when installing a license, the service may determine the devices to
	// which the license applies. This parameter shall not be present if the AuthorizationScope property
	// contains the value `Service`.
	AuthorizedDevices []string `json:",omitempty"`
	// LicenseFileURI shall contain an RFC3986-defined URI that links to a file that the license service
	// retrieves to install the license in that file. This URI should contain a scheme that describes the
	// transfer protocol. If the TransferProtocol parameter is absent or not supported, and a transfer protocol
	// is not specified by a scheme contained within this URI, the service shall use HTTP to get the file.
	LicenseFileURI string
	// Password (optional) shall represent the password to access the URI specified by the LicenseFileURI parameter.
	Password string `json:",omitempty"`
	// TargetServices (optional) shall contain an array of links to resources of type Manager that represent the
	// services where the license will be installed, such as remote Redfish services. This parameter shall only
	// be present in aggregators when the AuthorizationScope property contains `Service` or `Capacity`.
	TargetServices []string `json:",omitempty"`
	// TransferProtocol (optional) is the network protocol that the license service shall use to retrieve the license file
	// located at the LicenseFileURI.  Services should ignore this parameter if the URI provided in LicenseFileURI
	// contains a scheme.  If this parameter is not provided or supported, and if a transfer protocol is not
	// specified by a scheme contained within this URI, the service shall use HTTP to retrieve the file.
	TransferProtocol TransferProtocolType `json:",omitempty"`
	// Username (optional) is the user name to access the URI specified by the LicenseFileURI parameter.
	Username string `json:",omitempty"`
}

// Install will install one or more licenses from a remote file. The service may update an existing License resource.
func (licenseservice *LicenseService) Install(parameters *InstallLicenseParameters) error {
	if licenseservice.installTarget == "" {
		return errors.New("license install not supported by this service")
	}
	return licenseservice.Post(licenseservice.installTarget, parameters)
}

// Licenses gets the set of installed licenses.
func (licenseservice *LicenseService) Licenses() ([]*License, error) {
	return ListReferencedLicenses(licenseservice.GetClient(), licenseservice.licenses)
}

// Update commits updates to this object's properties to the running system.
func (licenseservice *LicenseService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LicenseService)
	original.UnmarshalJSON(licenseservice.rawData)

	readWriteFields := []string{
		"LicenseExpirationWarningDays",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(licenseservice).Elem()

	return licenseservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLicenseService will get a LicenseService instance from the service.
func GetLicenseService(c common.Client, uri string) (*LicenseService, error) {
	return common.GetObject[LicenseService](c, uri)
}

// ListReferencedLicenseServices gets the collection of LicenseService from
// a provided reference.
func ListReferencedLicenseServices(c common.Client, link string) ([]*LicenseService, error) {
	return common.GetCollectionObjects[LicenseService](c, link)
}
