//
// SPDX-License-Identifier: BSD-3-Clause
//

package zt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const SubscriptionURL = "/redfish/v1/EventService/Subscriptions"
const eventContext string = "root"

var (
	SubmitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"
)

type SubscriptionZtRequestType struct {
	Destination string            `json:"Destination"`
	HTTPHeaders map[string]string `json:"HttpHeaders,omitempty"`
	Protocol    string            `json:"Protocol,omitempty"`
	Context     string            `json:"Context,omitempty"`
}

// SubscribeResponseType zt uses a unique subscription response.
type SubscribeResponseType struct {
	ODataContext        string `json:"@odata.context"`
	ODataEtag           string `json:"@odata.etag"`
	ODataID             string `json:"@odata.id"`
	ODataType           string `json:"@odata.type"`
	Context             string `json:"Context"`
	DeliveryRetryPolicy string `json:"DeliveryRetryPolicy"`
	Description         string `json:"Description"`
	Destination         string `json:"Destination"`
	EventFormatType     string `json:"EventFormatType"`
	ID                  int    `json:"ID"`
	Name                string `json:"Name"`
	Protocol            string `json:"Protocol"`
	Status              struct {
		common.Status
	}
	SubordinateResources bool   `json:"SubordinateResources"`
	SubscriptionType     string `json:"SubscriptionType"`
}

type eventPayload struct {
	MessageID string `json:"MessageId"`
}

type EventService struct {
	redfish.EventService
}

func getSubscriptionURL(ztSubscribeResponse *SubscribeResponseType) string {
	return fmt.Sprintf("%s/%v", ztSubscribeResponse.ODataID, ztSubscribeResponse.ID)
}

// Subscribe to ZT systems redfish
// eventsReceiverURL is the http/s URL that will accept the events sent from redfish
// protocol is usually "redfish"
func (eventservice *EventService) Subscribe(eventsReceiverURL, protocol string) (string, error) {
	z := &SubscriptionZtRequestType{
		Destination: eventsReceiverURL,
		Protocol:    protocol,
		Context:     eventContext,
	}
	resp, err := eventservice.Client.Post(SubscriptionURL, z)

	if err != nil {
		return "", fmt.Errorf("failed to POST subscribe request to redfish due to %w", err)
	}
	defer resp.Body.Close()

	var ztSubscribeResponse SubscribeResponseType
	err = json.NewDecoder(resp.Body).Decode(&ztSubscribeResponse)
	if err != nil {
		return "", fmt.Errorf("failed to read response body from subscription request due to: %w", err)
	}

	subscriptionLink := getSubscriptionURL(&ztSubscribeResponse)
	return subscriptionLink, nil
}

// SubmitTestEvent sends event according to msgId and returns error.
func (eventservice *EventService) SubmitTestEvent(msgID string) error {
	p := eventPayload{
		MessageID: msgID,
	}
	resp, err := eventservice.Client.Post(SubmitTestEventTarget, p)

	if err != nil {
		return fmt.Errorf("failed to send submitTestEvent in SubmitTestEvent() due to: %w", err)
	}
	defer resp.Body.Close()

	valid := map[int]bool{http.StatusAccepted: true}

	if !valid[resp.StatusCode] {
		return fmt.Errorf("on send event received response: %v due to: %s", resp.StatusCode, resp.Body)
	}

	return nil
}

func FromEventService(eventservice *redfish.EventService) (EventService, error) {
	return EventService{*eventservice}, nil
}
