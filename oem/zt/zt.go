//
// SPDX-License-Identifier: BSD-3-Clause
//

package zt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import (
	"github.com/stmcginnis/gofish/common"
)

var (
	SubmitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"
)

type (
	SubscriptionZtRequestType struct {
		Destination string            `json:"Destination"`
		HTTPHeaders map[string]string `json:"HttpHeaders,omitempty"`
		Protocol    string            `json:"Protocol,omitempty"`
		Context     string            `json:"Context,omitempty"`
	}

	// subscribeResponseType zt uses a unique subscription response.
	subscribeZtResponseType struct {
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
		common.Status        `json:"Status"`
		SubordinateResources bool   `json:"SubordinateResources"`
		SubscriptionType     string `json:"SubscriptionType"`
	}

	eventPayload struct {
		MessageID string `json:"MessageId"`
	}
)

func getSubscriptionUri(ztSubscribeResponse subscribeZtResponseType) string {
	return fmt.Sprintf("%s/%v", ztSubscribeResponse.ODataID, ztSubscribeResponse.ID)
}

func SubscribeZT(c common.Client, subscriptionUrl, eventsReceiverUrl, protocol, context string) (string, error) {
	z := &SubscriptionZtRequestType{
		Destination: eventsReceiverUrl,
		Protocol:    protocol,
		Context:     context,
	}
	resp, err := c.Post(subscriptionUrl, z)
	if err != nil {
		return "", fmt.Errorf("failed to POST subscribe request to redfish due to %w", err)
	}
	defer resp.Body.Close()

	var ztSubscribeResponse subscribeZtResponseType
	err = json.NewDecoder(resp.Body).Decode(&ztSubscribeResponse)
	if err != nil {
		return "", fmt.Errorf("failed to read response body from subscription request due to: %w", err)
	}

	subscriptionLink := getSubscriptionUri(ztSubscribeResponse)
	return subscriptionLink, nil
}

// SendEventZt sends event according to msgId and returns error.
func SendEventZt(client common.Client, msgID string) error {
	p := eventPayload{
		MessageID: msgID,
	}
	resp, err := client.Post(SubmitTestEventTarget, p)

	if err != nil {
		return fmt.Errorf("failed to send submitTestEvent in SendEventZt() due to: %w", err)
	}
	defer resp.Body.Close()

	valid := map[int]bool{http.StatusAccepted: true}

	if !valid[resp.StatusCode] {
		return fmt.Errorf("on send event received response: %v due to: %s", resp.StatusCode, resp.Body)
	}

	return nil
}
