//
// SPDX-License-Identifier: BSD-3-Clause
//

package oem

import (
	"fmt"
	"strings"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/oem/dell"
	"github.com/stmcginnis/gofish/redfish"
)

const (
	DellRedfishOem = "Dell"
	Dell           = "Dell"
	HpeRedfishOem  = "Hpe"
	Hpe            = "Hpe"
	ZTRedfishOem   = "Ami"
	ZT             = "zt"
)

type EventService struct {
	redfish.EventService
}

// GetRedfishVendor queries the redfish root for the OEM field and returns the vendor.
func GetRedfishVendor(es *redfish.EventService) (string, error) {
	serviceRoot, err := gofish.ServiceRoot(es.Client)
	if err != nil {
		return "", fmt.Errorf("failed to GET redfish service root due to: %w", err)
	}

	oem := string(serviceRoot.Oem)

	switch {
	case strings.Contains(oem, DellRedfishOem):
		return Dell, nil
	case strings.Contains(oem, HpeRedfishOem):
		return Hpe, nil
	case strings.Contains(oem, ZTRedfishOem):
		return ZT, nil
	default:
		return "", fmt.Errorf("failed to match vendor from output: %v", oem)
	}
}

// GetEventServiceByVendor returns a vendor specific event service
func GetEventServiceByVendor(es *redfish.EventService) (*dell.EventService, error) {
	vendor, err := GetRedfishVendor(es)
	if err != nil {
		return nil, fmt.Errorf("failed to get event service by vendor dur to: %w", err)
	}
	switch vendor {
	case Dell:
		return dell.FromEventService(es)
	default:
		return nil, fmt.Errorf("failed to get event service by vendor for vendor: %v", vendor)
	}
}
