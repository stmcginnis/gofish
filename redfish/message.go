//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
)

// Message shall contain a message that the Redfish service returns, as described in the Redfish Specification.
type Message struct {
	// Message shall contain a human-readable message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted for the arguments in the message
	// when looked up in the message registry. It has the same semantics as the MessageArgs property in the Redfish
	// MessageRegistry schema. If the corresponding ParamType value contains 'number', the service shall convert the
	// number to a string representation of the number.
	MessageArgs []string
	// MessageID shall contain a MessageId, as defined in the 'MessageId format' clause of the Redfish Specification.
	MessageID string
	// MessageSeverity shall contain the severity of the message. Services can replace the value defined in the message
	// registry with a value more applicable to the implementation.
	MessageSeverity string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedProperties shall contain an array of RFC6901-defined JSON pointers indicating the properties described by
	// the message, if appropriate for the message.
	RelatedProperties []string
	// Resolution shall contain the resolution of the message. Services can replace the resolution defined in the
	// message registry with a more specific resolution in message payloads.
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the situation that caused the message.
	// This property shall not be present if the MessageSeverity or Severity properties contain 'OK'.
	ResolutionSteps []ResolutionStep
}
