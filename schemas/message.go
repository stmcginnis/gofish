//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Message.v1_4_0.json
// 2026.2 - #Message.v1_4_0

package schemas

import (
	"encoding/json"
)

// Message shall contain a message that the Redfish service returns, as
// described in the Redfish Specification.
type Message struct {
	// Message shall contain a human-readable message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted
	// for the arguments in the message when looked up in the message registry. It
	// has the same semantics as the 'MessageArgs' property in the Redfish
	// 'MessageRegistry' schema. If the corresponding 'ParamType' value contains
	// 'number', the service shall convert the number to a string representation of
	// the number.
	MessageArgs []string
	// MessageID shall contain a 'MessageId', as defined in the 'MessageId format'
	// clause of the Redfish Specification.
	MessageID string `json:"MessageId"`
	// MessageSeverity shall contain the severity of the message. Services can
	// replace the value defined in the message registry with a value more
	// applicable to the implementation.
	//
	// Version added: v1.1.0
	MessageSeverity Health
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedProperties shall contain an array of RFC6901-defined JSON pointers
	// indicating the properties described by the message, if appropriate for the
	// message.
	RelatedProperties []string
	// Resolution shall contain the resolution of the message. Services can replace
	// the resolution defined in the message registry with a more specific
	// resolution in message payloads.
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the
	// situation that caused the message. This property shall not be present if the
	// 'MessageSeverity' or 'Severity' properties contain 'OK'.
	//
	// Version added: v1.2.0
	ResolutionSteps []ResolutionStep
	// Severity shall contain the severity of the message, as defined in the
	// 'Status' clause of the Redfish Specification. Services can replace the value
	// defined in the message registry with a value more applicable to the
	// implementation.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of 'MessageSeverity', which ties
	// the values to the enumerations defined for the 'Health' property within
	// 'Status'.
	Severity string
	// UserAuthenticationSource shall contain source of authentication for the
	// user, specified by the 'Username' property, associated with the message. For
	// local users, this property shall contain '/redfish/v1/AccountService'. For
	// users managed by external account providers, this property shall contain the
	// member of the 'ServiceAddresses' property that authenticated the user. This
	// should be used for messages that result from a user action.
	//
	// Version added: v1.3.0
	UserAuthenticationSource string
	// Username shall contain the username of the account associated with the
	// message. This should be used for messages that result from a user action.
	//
	// Version added: v1.3.0
	Username string
	// VendorCode shall contain a vendor-defined code or identifier equivalent to
	// this message or 'MessageId' value. This value allows clients to map a
	// 'MessageId' value to a legacy or vendor-specific value for the same
	// condition or error message.
	//
	// Version added: v1.4.0
	VendorCode string
}
