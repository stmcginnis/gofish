//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Application shall represent an application or service running on a computer system.
type Application struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// DestinationURIs shall contain an array of URIs to which this application pushes data. This is typically for
	// applications that act as logging or metric agents that transmit data captured to remote servers.
	DestinationURIs []string
	// MetricsURIs shall contain an array of URIs that provide access to data or other information in this application.
	// This is typically for applications that allow external users to perform requests to pull data from the
	// application.
	MetricsURIs []string
	// StartTime shall indicate the date and time when the application started running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Vendor shall contain the name of the company that provides this application.
	Vendor string
	// Version shall contain the version of this application.
	Version string
	// resetTarget is the internal URL to send reset actions to.
	resetTarget string
	// softwareImage links a resource of type SoftwareInventory that represents the software image from which this application runs.
	softwareImage string
}

// UnmarshalJSON unmarshals a Application object from the raw JSON.
func (application *Application) UnmarshalJSON(b []byte) error {
	type temp Application
	type Actions struct {
		ApplicationReset common.ActionTarget `json:"#Application.Reset"`
	}
	type Links struct {
		// SoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the software image
		// from which this application runs.
		SoftwareImage string
	}
	var t struct {
		temp
		Actions Actions
		Links   Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*application = Application(t.temp)

	// Extract the links to other entities for later
	application.resetTarget = t.Actions.ApplicationReset.Target
	application.softwareImage = t.Links.SoftwareImage

	return nil
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

// SoftwareImage returns a `SoftwareInventory“ that represents the software image from which this application runs.
func (application *Application) SoftwareImage() (*SoftwareInventory, error) {
	return GetSoftwareInventory(application.GetClient(), application.softwareImage)
}

// Reset resets the application.
//
// `ResetType` is the type of reset.
// `GracefulRestart` and `ForceRestart` shall indicate requests to restart the application.
// `GracefulShutdown` and `ForceOff` shall indicate requests to stop or disable the application.
// `On` and `ForceOn` shall indicate requests to start or enable the application.
func (application *Application) Reset(resetType ResetType) error {
	t := struct {
		ResetType ResetType
	}{ResetType: resetType}

	return application.Post(application.resetTarget, t)
}
