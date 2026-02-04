//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/LineOfService.v1_1_0.json
// 1.2.1c - #LineOfService.v1_1_0.LineOfService

package schemas

import (
	"encoding/json"
)

// LineOfService This service option is the abstract base class for other
// ClassOfService and concrete lines of service.
type LineOfService struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetLineOfService will get a LineOfService instance from the service.
func GetLineOfService(c Client, uri string) (*LineOfService, error) {
	return GetObject[LineOfService](c, uri)
}

// ListReferencedLineOfServices gets the collection of LineOfService from
// a provided reference.
func ListReferencedLineOfServices(c Client, link string) ([]*LineOfService, error) {
	return GetCollectionObjects[LineOfService](c, link)
}
