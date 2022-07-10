//
// SPDX-License-Identifier: BSD-3-Clause
//

package hpe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stmcginnis/gofish/redfish"
)

type PayloadType struct {
	EventID           string `json:"EventID"`
	EventTimestamp    string `json:"EventTimestamp"`
	EventType         string `json:"EventType"`
	Message           string
	MessageArgs       []string
	MessageID         string `json:"MessageId"`
	OriginOfCondition string
	Severity          string
}

// EventService is the Hpe-specific handler for the EventService instance.
type EventService struct {
	redfish.EventService
}

// SubmitTestEvent sends event according to msgId and returns error
// more info https://hewlettpackard.github.io/iLOAmpPack-Redfish-API-Docs/#submitting-a-test-event
func (eventservice *EventService) SubmitTestEvent(eventID, messageID, eType, message string) error {
	const condition = "/redfish/v1/Systems/1/"
	const severity = "OK"

	var messageArgs = []string{"NoAMS", "Busy", "Cached"}

	payload := PayloadType{
		EventID:           eventID,
		EventTimestamp:    time.Now().Format(time.RFC3339), // "2019-07-29T15:13:49Z",
		EventType:         eType,                           // redfish.SupportedEventTypes["Alert"],
		Message:           message,
		MessageArgs:       messageArgs,
		MessageID:         messageID,
		OriginOfCondition: condition,
		Severity:          severity,
	}
	resp, err := eventservice.Client.Post(eventservice.SubmitTestEventTarget, payload)

	if err != nil {
		return fmt.Errorf("failed to send submitTestEvent due to: %w", err)
	}
	defer resp.Body.Close()

	valid := map[int]bool{
		http.StatusOK:        true,
		http.StatusNoContent: true,
		http.StatusCreated:   true}

	if !valid[resp.StatusCode] {
		return fmt.Errorf("on send event received response: %v due to: %s", resp.StatusCode, resp.Body)
	}

	return nil
}

// FromEventService converts a standard EventService object to the OEM implementation.
func FromEventService(eventservice *redfish.EventService) (*EventService, error) {
	es := &EventService{}
	err := json.Unmarshal(eventservice.RawData, es)
	if err != nil {
		return nil, err
	}

	es.SetClient(eventservice.Client)
	return es, nil
}
