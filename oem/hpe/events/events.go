//
// SPDX-License-Identifier: BSD-3-Clause
//

package events

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stmcginnis/gofish/common"
)

var (
	SubmitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"
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

// SubmitTestEvent sends event according to msgId and returns error
// more info https://hewlettpackard.github.io/iLOAmpPack-Redfish-API-Docs/#submitting-a-test-event
func SubmitTestEvent(client common.Client, eventID, msgID, eventType, message string) error {
	const condition = "/redfish/v1/Systems/1/"
	const severity = "OK"

	var messageArgs = []string{"NoAMS", "Busy", "Cached"}

	payload := PayloadType{
		EventID:           eventID,
		EventTimestamp:    time.Now().Format(time.RFC3339), // "2019-07-29T15:13:49Z",
		EventType:         eventType,                       // redfish.SupportedEventTypes["Alert"],
		Message:           message,
		MessageArgs:       messageArgs,
		MessageID:         msgID,
		OriginOfCondition: condition,
		Severity:          severity,
	}
	resp, err := client.Post(SubmitTestEventTarget, payload)

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
