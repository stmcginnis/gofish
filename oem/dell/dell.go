//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"fmt"
	"net/http"

	"github.com/stmcginnis/gofish/redfish"
)

const eventContext string = "root"

var SubmitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"

type (
	PayloadType struct {
		Destination string `json:"Destination"`
		EventTypes  string `json:"EventTypes"`
		Context     string `json:"Context"`
		Protocol    string `json:"Protocol"`
		MessageID   string `json:"MessageId"`
	}
)

type EventService struct {
	redfish.EventService
}

// SubmitTestEvent sends event according to msgId and returns error.
func (eventservice *EventService) SubmitTestEvent(messageID, eType, protocol string) error {
	payload := PayloadType{
		Destination: SubmitTestEventTarget,
		EventTypes:  eType,
		Context:     eventContext,
		Protocol:    protocol,
		MessageID:   messageID,
	}
	resp, err := eventservice.Client.Post(SubmitTestEventTarget, payload)

	if err != nil {
		return fmt.Errorf("failed to post submitTestEvent due to: %w", err)
	}
	defer resp.Body.Close()

	valid := map[int]bool{
		http.StatusNoContent: true,
		http.StatusCreated:   true}

	if !valid[resp.StatusCode] {
		return fmt.Errorf("on send event received response: %v due to: %s", resp.StatusCode, resp.Body)
	}

	return nil
}

func FromEventService(eventservice *redfish.EventService) (EventService, error) {
	return EventService{*eventservice}, nil
}
