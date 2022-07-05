//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"fmt"
	"github.com/stmcginnis/gofish/common"
	"net/http"
)

const (
	eventContext string = "root"
)

var (
	submitTestEventTarget = "/redfish/v1/EventService/Actions/EventService.SendTestEvent"
)

type (
	dellPayloadType struct {
		Destination string `json:"Destination"`
		EventTypes  string `json:"EventTypes"`
		Context     string `json:"Context"`
		Protocol    string `json:"Protocol"`
		MessageID   string `json:"MessageId"`
	}
)

// SendEventDell sends event according to msgId and returns error.
func SendEventDell(client common.Client, messageID, eType, protocol string) error {
	payload := dellPayloadType{
		Destination: submitTestEventTarget,
		EventTypes:  eType,
		Context:     eventContext,
		Protocol:    protocol,
		MessageID:   messageID,
	}
	resp, err := client.Post(submitTestEventTarget, payload)

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
