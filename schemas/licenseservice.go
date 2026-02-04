//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/LicenseService.v1_1_2.json
// 2022.3 - #LicenseService.v1_1_2.LicenseService

package schemas

import (
	"encoding/json"
)

// LicenseService shall represent a license service and the properties that
// affect the service itself for a Redfish implementation.
type LicenseService struct {
	Entity
	// LicenseExpirationWarningDays shall contain the number of days prior to a
	// license expiration that the service shall send the 'DaysBeforeExpiration'
	// message from the License Message Registry at least once. A value of zero
	// shall indicate that no warning messages are sent prior to license
	// expiration.
	LicenseExpirationWarningDays *int `json:",omitempty"`
	// Licenses shall contain a link to a resource collection of type
	// 'LicenseCollection'. When installing a license with a 'POST' operation to
	// this collection, the service may update an existing License resource instead
	// of creating a new resource. In these cases, the service shall respond with
	// the HTTP '200 OK' status code or HTTP '204 No Content' status code and the
	// 'Location' header in the response shall contain the URI of the updated
	// License resource.
	licenses string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// installTarget is the URL to send Install requests.
	installTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a LicenseService object from the raw JSON.
func (l *LicenseService) UnmarshalJSON(b []byte) error {
	type temp LicenseService
	type lActions struct {
		Install ActionTarget `json:"#LicenseService.Install"`
	}
	var tmp struct {
		temp
		Actions  lActions
		Licenses Link `json:"Licenses"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LicenseService(tmp.temp)

	// Extract the links to other entities for later
	l.installTarget = tmp.Actions.Install.Target
	l.licenses = tmp.Licenses.String()

	// This is a read/write object, so we need to save the raw object data for later
	l.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LicenseService) Update() error {
	readWriteFields := []string{
		"LicenseExpirationWarningDays",
		"ServiceEnabled",
	}

	return l.UpdateFromRawData(l, l.RawData, readWriteFields)
}

// GetLicenseService will get a LicenseService instance from the service.
func GetLicenseService(c Client, uri string) (*LicenseService, error) {
	return GetObject[LicenseService](c, uri)
}

// ListReferencedLicenseServices gets the collection of LicenseService from
// a provided reference.
func ListReferencedLicenseServices(c Client, link string) ([]*LicenseService, error) {
	return GetCollectionObjects[LicenseService](c, link)
}

// LicenseServiceInstallParameters holds the parameters for the Install action.
type LicenseServiceInstallParameters struct {
	// AuthorizedDevices shall contain an array of links to the devices to be
	// authorized by the license. Clients can provide this parameter when
	// installing a license to apply the license to specific devices. If not
	// provided when installing a license, the service may determine the devices to
	// which the license applies. This parameter shall not be present if the
	// 'AuthorizationScope' property contains the value 'Service'.
	AuthorizedDevices []string `json:"AuthorizedDevices,omitempty"`
	// LicenseFileURI shall contain an RFC3986-defined URI that links to a file
	// that the license service retrieves to install the license in that file. This
	// URI should contain a scheme that describes the transfer protocol. If the
	// 'TransferProtocol' parameter is absent or not supported, and a transfer
	// protocol is not specified by a scheme contained within this URI, the service
	// shall use HTTP to get the file.
	LicenseFileURI string `json:"LicenseFileURI,omitempty"`
	// Password shall contain the password to access the URI specified by the
	// 'LicenseFileURI' parameter.
	Password string `json:"Password,omitempty"`
	// TargetServices shall contain an array of links to resources of type
	// 'Manager' that represent the services where the license will be installed,
	// such as remote Redfish services. This parameter shall only be present in
	// aggregators when the 'AuthorizationScope' property contains 'Service' or
	// 'Capacity'.
	TargetServices []string `json:"TargetServices,omitempty"`
	// TransferProtocol shall contain the network protocol that the license service
	// shall use to retrieve the license file located at the 'LicenseFileURI'.
	// Services should ignore this parameter if the URI provided in
	// 'LicenseFileURI' contains a scheme. If this parameter is not provided or
	// supported, and if a transfer protocol is not specified by a scheme contained
	// within this URI, the service shall use HTTP to retrieve the file.
	TransferProtocol TransferProtocolType `json:"TransferProtocol,omitempty"`
	// Username shall contain the username to access the URI specified by the
	// 'LicenseFileURI' parameter.
	Username string `json:"Username,omitempty"`
}

// This action shall install one or more licenses from a remote file. The
// service may update an existing 'License' resource. The 'Location' header in
// the response shall contain the URI of the new or updated 'License' resource.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (l *LicenseService) Install(params *LicenseServiceInstallParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(l.client,
		l.installTarget, params, l.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Licenses gets the Licenses collection.
func (l *LicenseService) Licenses() ([]*License, error) {
	if l.licenses == "" {
		return nil, nil
	}
	return GetCollectionObjects[License](l.client, l.licenses)
}
