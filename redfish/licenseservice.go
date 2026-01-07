//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #LicenseService.v1_1_2.LicenseService

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type LicenseServiceTransferProtocolType string

const (
	// CIFSLicenseServiceTransferProtocolType Common Internet File System (CIFS).
	CIFSLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "CIFS"
	// FTPLicenseServiceTransferProtocolType File Transfer Protocol (FTP).
	FTPLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "FTP"
	// SFTPLicenseServiceTransferProtocolType SSH File Transfer Protocol (SFTP).
	SFTPLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "SFTP"
	// HTTPLicenseServiceTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "HTTP"
	// HTTPSLicenseServiceTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "HTTPS"
	// SCPLicenseServiceTransferProtocolType Secure Copy Protocol (SCP).
	SCPLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "SCP"
	// TFTPLicenseServiceTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "TFTP"
	// OEMLicenseServiceTransferProtocolType is a manufacturer-defined protocol.
	OEMLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "OEM"
	// NFSLicenseServiceTransferProtocolType Network File System (NFS).
	NFSLicenseServiceTransferProtocolType LicenseServiceTransferProtocolType = "NFS"
)

// LicenseService shall represent a license service and the properties that
// affect the service itself for a Redfish implementation.
type LicenseService struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// installTarget is the URL to send Install requests.
	installTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LicenseService object from the raw JSON.
func (l *LicenseService) UnmarshalJSON(b []byte) error {
	type temp LicenseService
	type lActions struct {
		Install common.ActionTarget `json:"#LicenseService.Install"`
	}
	var tmp struct {
		temp
		Actions  lActions
		Licenses common.Link `json:"licenses"`
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
	l.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LicenseService) Update() error {
	readWriteFields := []string{
		"LicenseExpirationWarningDays",
		"ServiceEnabled",
	}

	return l.UpdateFromRawData(l, l.rawData, readWriteFields)
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

// Install shall install one or more licenses from a remote file. The
// service may update an existing 'License' resource. The 'Location' header in
// the response shall contain the URI of the new or updated 'License' resource.
// authorizedDevices - This parameter shall contain an array of links to the
// devices to be authorized by the license. Clients can provide this parameter
// when installing a license to apply the license to specific devices. If not
// provided when installing a license, the service may determine the devices to
// which the license applies. This parameter shall not be present if the
// 'AuthorizationScope' property contains the value 'Service'.
// licenseFileURI - This parameter shall contain an RFC3986-defined URI that
// links to a file that the license service retrieves to install the license in
// that file. This URI should contain a scheme that describes the transfer
// protocol. If the 'TransferProtocol' parameter is absent or not supported,
// and a transfer protocol is not specified by a scheme contained within this
// URI, the service shall use HTTP to get the file.
// password - This parameter shall contain the password to access the URI
// specified by the 'LicenseFileURI' parameter.
// targetServices - This property shall contain an array of links to resources
// of type 'Manager' that represent the services where the license will be
// installed, such as remote Redfish services. This parameter shall only be
// present in aggregators when the 'AuthorizationScope' property contains
// 'Service' or 'Capacity'.
// transferProtocol - This parameter shall contain the network protocol that
// the license service shall use to retrieve the license file located at the
// 'LicenseFileURI'. Services should ignore this parameter if the URI provided
// in 'LicenseFileURI' contains a scheme. If this parameter is not provided or
// supported, and if a transfer protocol is not specified by a scheme contained
// within this URI, the service shall use HTTP to retrieve the file.
// username - This parameter shall contain the username to access the URI
// specified by the 'LicenseFileURI' parameter.
func (l *LicenseService) Install(authorizedDevices string, licenseFileURI string, password string, targetServices string, transferProtocol LicenseServiceTransferProtocolType, username string) error {
	payload := make(map[string]any)
	payload["AuthorizedDevices"] = authorizedDevices
	payload["LicenseFileURI"] = licenseFileURI
	payload["Password"] = password
	payload["TargetServices"] = targetServices
	payload["TransferProtocol"] = transferProtocol
	payload["Username"] = username
	return l.Post(l.installTarget, payload)
}

// Licenses gets the Licenses collection.
func (l *LicenseService) Licenses(client common.Client) ([]*License, error) {
	if l.licenses == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[License](client, l.licenses)
}
