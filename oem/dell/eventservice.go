//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"fmt"
	"net/http"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const eventContext string = "root"

type PayloadType struct {
	Destination string                           `json:"Destination"`
	EventTypes  string                           `json:"EventTypes"`
	Context     string                           `json:"Context"`
	Protocol    redfish.EventDestinationProtocol `json:"Protocol"`
	MessageID   string                           `json:"MessageId"`
}

// EventService is the Dell-specific handler for the EventService instance.
type EventService struct {
	redfish.EventService
}

// SubmitTestEvent sends event according to msgId and returns error.
func (eventservice *EventService) SubmitTestEvent(messageID, eType string, protocol redfish.EventDestinationProtocol) error {
	payload := PayloadType{
		Destination: eventservice.SubmitTestEventTarget,
		EventTypes:  eType,
		Context:     eventContext,
		Protocol:    protocol,
		MessageID:   messageID,
	}

	resp, err := eventservice.GetClient().Post(eventservice.SubmitTestEventTarget, payload)
	defer common.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return fmt.Errorf("failed to post submitTestEvent due to: %w", err)
	}

	valid := map[int]bool{
		http.StatusNoContent: true,
		http.StatusCreated:   true}

	if !valid[resp.StatusCode] {
		return fmt.Errorf("on send event received response: %v due to: %s", resp.StatusCode, resp.Body)
	}

	return nil
}

// FromEventService converts a standard EventService object to the OEM implementation.
func FromEventService(eventservice *redfish.EventService) (*EventService, error) {
	es := &EventService{
		EventService: *eventservice,
	}

	// Extract any OEM-specific data from the raw JSON if needed
	// For now, Dell EventService doesn't have additional OEM fields beyond the standard ones

	es.SetClient(eventservice.GetClient())
	return es, nil
}
