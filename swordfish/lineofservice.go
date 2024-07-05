//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// LineOfService This service option is the abstract base class for other ClassOfService and concrete lines of
// service.
type LineOfService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetLineOfService will get a LineOfService instance from the service.
func GetLineOfService(c common.Client, uri string) (*LineOfService, error) {
	return common.GetObject[LineOfService](c, uri)
}

// ListReferencedLineOfServices gets the collection of LineOfService from
// a provided reference.
func ListReferencedLineOfServices(c common.Client, link string) ([]*LineOfService, error) {
	return common.GetCollectionObjects[LineOfService](c, link)
}
