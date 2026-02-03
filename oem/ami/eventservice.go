//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/schemas"
)

// EventService is an AMI OEM instance of an EventService.
type EventService struct {
	schemas.EventService
	SecondarySMTP schemas.SMTP

	certificates string
}

// FromEventService converts a standard EventService object to the OEM implementation.
func FromEventService(eventService *schemas.EventService) (*EventService, error) {
	es := EventService{
		EventService: *eventService,
	}

	var t struct {
		Oem struct {
			AMI struct {
				Certificates  schemas.Link
				SecondarySMTP schemas.SMTP
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
func (es *EventService) Certificates() ([]*schemas.Certificate, error) {
	return schemas.ListReferencedCertificates(es.GetClient(), es.certificates)
}
