//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// LineOfService This service option is the abstract base class for other ClassOfService and concrete lines of
// service.
type LineOfService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetLineOfService will get a LineOfService instance from the service.
func GetLineOfService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*LineOfService, error) {
	return GetLineOfServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetLineOfServiceWithContext will get a LineOfService instance from the service.
func GetLineOfServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*LineOfService, error) {
	return common.GetObjectWithContext[LineOfService](ctx, c, uri, queryOpts...)
}

// ListReferencedLineOfServices gets the collection of LineOfService from
// a provided reference.
func ListReferencedLineOfServices(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*LineOfService, error) {
	return ListReferencedLineOfServicesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedLineOfServicesWithContext gets the collection of LineOfService from
// a provided reference.
func ListReferencedLineOfServicesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*LineOfService, error) {
	return common.GetCollectionObjectsWithContext[LineOfService](ctx, c, link, queryOpts...)
}
