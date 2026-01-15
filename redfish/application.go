//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #Application.v1_0_1.Application

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Application shall represent an application or service running on a computer
// system.
type Application struct {
	common.Entity
	// DestinationURIs shall contain an array of URIs to which this application
	// pushes data. This is typically for applications that act as logging or
	// metric agents that transmit data captured to remote servers.
	DestinationURIs []string
	// MetricsURIs shall contain an array of URIs that provide access to data or
	// other information in this application. This is typically for applications
	// that allow external users to perform requests to pull data from the
	// application.
	MetricsURIs []string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// StartTime shall indicate the date and time when the application started
	// running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Vendor shall contain the name of the company that provides this application.
	Vendor string
	// Version shall contain the version of this application.
	Version string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// softwareImage is the URI for SoftwareImage.
	softwareImage string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Application object from the raw JSON.
func (a *Application) UnmarshalJSON(b []byte) error {
	type temp Application
	type aActions struct {
		Reset common.ActionTarget `json:"#Application.Reset"`
	}
	type aLinks struct {
		SoftwareImage common.Link `json:"SoftwareImage"`
	}
	var tmp struct {
		temp
		Actions aActions
		Links   aLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = Application(tmp.temp)

	// Extract the links to other entities for later
	a.resetTarget = tmp.Actions.Reset.Target
	a.softwareImage = tmp.Links.SoftwareImage.String()

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *Application) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetApplication will get a Application instance from the service.
func GetApplication(c common.Client, uri string) (*Application, error) {
	return common.GetObject[Application](c, uri)
}

// ListReferencedApplications gets the collection of Application from
// a provided reference.
func ListReferencedApplications(c common.Client, link string) ([]*Application, error) {
	return common.GetCollectionObjects[Application](c, link)
}

// Reset shall reset the application.
// resetType - This parameter shall contain the type of reset.
// 'GracefulRestart' and 'ForceRestart' shall indicate requests to restart the
// application. 'GracefulShutdown' and 'ForceOff' shall indicate requests to
// stop or disable the application. 'On' and 'ForceOn' shall indicate requests
// to start or enable the application. The service can accept a request without
// the parameter and shall perform a 'GracefulRestart'.
func (a *Application) Reset(resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return a.Post(a.resetTarget, payload)
}

// SoftwareImage gets the SoftwareImage linked resource.
func (a *Application) SoftwareImage(client common.Client) (*SoftwareInventory, error) {
	if a.softwareImage == "" {
		return nil, nil
	}
	return common.GetObject[SoftwareInventory](client, a.softwareImage)
}
