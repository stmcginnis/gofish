//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/UpdateServiceCapabilities.v1_0_0.json
// 2025.2 - #UpdateServiceCapabilities.v1_0_0.UpdateServiceCapabilities

package schemas

import (
	"encoding/json"
)

// UpdateServiceCapabilities shall represent the update service capabilities for
// a Redfish implementation.
type UpdateServiceCapabilities struct {
	Entity
	// AllowableStaging shall contain an array of the allowable URIs that a client
	// can specify in the 'Targets' parameter for 'SimpleUpdate' or
	// 'MultipartHttpPushUri' update requests with 'Stage' set to 'true'.
	AllowableStaging []string
	// AllowableTargets shall contain an array of the allowable URIs that a client
	// can specify in the 'Targets' parameter for 'SimpleUpdate' or
	// 'MultipartHttpPushUri' update requests.
	AllowableTargets []string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetUpdateServiceCapabilities will get a UpdateServiceCapabilities instance from the service.
func GetUpdateServiceCapabilities(c Client, uri string) (*UpdateServiceCapabilities, error) {
	return GetObject[UpdateServiceCapabilities](c, uri)
}

// ListReferencedUpdateServiceCapabilitiess gets the collection of UpdateServiceCapabilities from
// a provided reference.
func ListReferencedUpdateServiceCapabilitiess(c Client, link string) ([]*UpdateServiceCapabilities, error) {
	return GetCollectionObjects[UpdateServiceCapabilities](c, link)
}
