//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Application.v1_0_1.json
// 2023.2 - #Application.v1_0_1.Application

package schemas

import (
	"encoding/json"
)

// Application shall represent an application or service running on a computer
// system.
type Application struct {
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// StartTime shall indicate the date and time when the application started
	// running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Vendor shall contain the name of the company that provides this application.
	Vendor string
	// Version shall contain the version of this application.
	Version string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// softwareImage is the URI for SoftwareImage.
	softwareImage string
}

// UnmarshalJSON unmarshals a Application object from the raw JSON.
func (a *Application) UnmarshalJSON(b []byte) error {
	type temp Application
	type aActions struct {
		Reset ActionTarget `json:"#Application.Reset"`
	}
	type aLinks struct {
		SoftwareImage Link `json:"SoftwareImage"`
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

	return nil
}

// GetApplication will get a Application instance from the service.
func GetApplication(c Client, uri string) (*Application, error) {
	return GetObject[Application](c, uri)
}

// ListReferencedApplications gets the collection of Application from
// a provided reference.
func ListReferencedApplications(c Client, link string) ([]*Application, error) {
	return GetCollectionObjects[Application](c, link)
}

// This action shall reset the application.
// resetType - This parameter shall contain the type of reset.
// 'GracefulRestart' and 'ForceRestart' shall indicate requests to restart the
// application. 'GracefulShutdown' and 'ForceOff' shall indicate requests to
// stop or disable the application. 'On' and 'ForceOn' shall indicate requests
// to start or enable the application. The service can accept a request without
// the parameter and shall perform a 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *Application) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(a.client,
		a.resetTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// SoftwareImage gets the SoftwareImage linked resource.
func (a *Application) SoftwareImage() (*SoftwareInventory, error) {
	if a.softwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](a.client, a.softwareImage)
}
