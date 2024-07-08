//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// Message is This type shall define a Message as described in the
// Redfish specification.
type Message struct {
	Entity

	// Message shall contain an optional human readable message.
	Message string
	// MessageArgs shall contain the message substitution
	// arguments for the specific message referenced by the MessageID and
	// shall only be included if the MessageID is present.  Number and
	// integer type arguments shall be converted to strings.
	MessageArgs []string
	// MessageID shall be a key into message registry as described in the
	// Redfish specification.
	MessageID string `json:"MessageId"`
	// RelatedProperties shall contain an array of JSON
	// Pointers indicating the properties described by the message, if
	// appropriate for the message.
	RelatedProperties []string
	// Resolution shall contain an override of the
	// Resolution of the message in message registry, if present.
	Resolution string
	// Severity is The value of this property shall be the severity of the
	// error, as defined in the Status section of the Redfish specification.
	Severity string
}

// GetMessage will get a Message instance from the service.
func GetMessage(c Client, uri string) (*Message, error) {
	return GetObject[Message](c, uri)
}

// ListReferencedMessages gets the collection of Message from
// a provided reference.
func ListReferencedMessages(c Client, link string) ([]*Message, error) {
	return GetCollectionObjects[Message](c, link)
}
