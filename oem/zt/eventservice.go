//
// SPDX-License-Identifier: BSD-3-Clause
//

package zt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const eventContext string = "root"

type SubscriptionZtRequestType struct {
	Destination string                           `json:"Destination"`
	HTTPHeaders map[string]string                `json:"HttpHeaders,omitempty"`
	Protocol    redfish.EventDestinationProtocol `json:"Protocol,omitempty"`
	Context     string                           `json:"Context,omitempty"`
}

// SubscribeResponseType zt uses a unique subscription response.
type SubscribeResponseType struct {
	ODataContext         string `json:"@odata.context"`
	ODataEtag            string `json:"@odata.etag"`
	ODataID              string `json:"@odata.id"`
	ODataType            string `json:"@odata.type"`
	Context              string `json:"Context"`
	DeliveryRetryPolicy  string `json:"DeliveryRetryPolicy"`
	Description          string `json:"Description"`
	Destination          string `json:"Destination"`
	EventFormatType      string `json:"EventFormatType"`
	ID                   int    `json:"ID"`
	Name                 string `json:"Name"`
	Protocol             string `json:"Protocol"`
	Status               common.Status
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
func (eventservice *EventService) Subscribe(eventsReceiverURL string, protocol redfish.EventDestinationProtocol) (string, error) {
	z := &SubscriptionZtRequestType{
		Destination: eventsReceiverURL,
		Protocol:    protocol,
		Context:     eventContext,
	}

	resp, err := eventservice.GetClient().Post(eventservice.Subscriptions, z)
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
// ZT systems redfish current firmware limits the SubmitTestEvent() request rate.
// The first request will be OK, but if another request is sent with-in 5-15 sec, it will get a 400 error response.
// This issue is reported at https://bugzilla.redhat.com/show_bug.cgi?id=2094842
// To work around this limit, we retry sending the request until we get a good response.
func (eventservice *EventService) SubmitTestEvent(msgID string) error {
	const (
		retryInterval        = 5 * time.Second
		retryAttempts        = 6
		retryReportThreshold = 2
	)
	var err error
	var resp *http.Response

	p := eventPayload{
		MessageID: msgID,
	}

	for retryCounter := 0; retryCounter < retryAttempts; retryCounter++ {
		resp, err = eventservice.GetClient().Post(eventservice.SubmitTestEventTarget, p)
		if err == nil {
			if retryCounter > retryReportThreshold {
				log.Printf("Had to retry %v times to send SubmitTestEvent()", retryCounter)
			}

			break
		}

		if resp != nil {
			_ = resp.Body.Close()
		}

		time.Sleep(retryInterval)
	}
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	valid := map[int]bool{http.StatusAccepted: true}

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

	es.SetClient(eventservice.GetClient())
	return es, nil
}
