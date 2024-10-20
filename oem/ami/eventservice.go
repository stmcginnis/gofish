//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// EventService is an AMI OEM instance of an EventService.
type EventService struct {
	redfish.EventService
	SecondarySMTP redfish.SMTP

	certificates string
}

// FromEventService converts a standard EventService object to the OEM implementation.
func FromEventService(eventService *redfish.EventService) (*EventService, error) {
	es := EventService{
		EventService: *eventService,
	}

	var t struct {
		Oem struct {
			AMI struct {
				Certificates  common.Link
				SecondarySMTP redfish.SMTP
			} `json:"Ami"`
		} `json:"Oem"`
	}

	err := json.Unmarshal(eventService.RawData, &t)
	if err != nil {
		return nil, err
	}

	es.SecondarySMTP = t.Oem.AMI.SecondarySMTP
	es.certificates = t.Oem.AMI.Certificates.String()
	es.SetClient(eventService.GetClient())

	return &es, nil
}

// Certificates will get the Certificates for this EventService.
func (es *EventService) Certificates() ([]*redfish.Certificate, error) {
	return redfish.ListReferencedCertificates(es.GetClient(), es.certificates)
}
