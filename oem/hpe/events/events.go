package events

import (
	"fmt"
	"github.com/stmcginnis/gofish/common"
	"net/http"
	"time"
)

var (
	submitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"
)

type hpePayloadType struct {
	EventID           string `json:"EventID"`
	EventTimestamp    string `json:"EventTimestamp"`
	EventType         string `json:"EventType"`
	Message           string
	MessageArgs       []string
	MessageID         string `json:"MessageId"`
	OriginOfCondition string
	Severity          string
}

// SendEventHP sends event according to msgId and returns error
// more info https://hewlettpackard.github.io/iLOAmpPack-Redfish-API-Docs/#submitting-a-test-event
func SendEventHP(client common.Client, msgID string) error {
	payload := hpePayloadType{
		EventID:           "TestEventId",
		EventTimestamp:    time.Now().Format(time.RFC3339), // "2019-07-29T15:13:49Z",
		EventType:         "Alert",                         // redfish.SupportedEventTypes["Alert"],
		Message:           "Test Event",
		MessageArgs:       []string{"NoAMS", "Busy", "Cached"},
		MessageID:         msgID,
		OriginOfCondition: "/redfish/v1/Systems/1/",
		Severity:          "OK",
	}
	resp, err := client.Post(submitTestEventTarget, payload)

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
