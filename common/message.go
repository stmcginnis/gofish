//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import "context"

type MessageExtendedInfo struct {
	// MessageID shall be a key into message registry as described in the
	// Redfish specification.
	MessageID string `json:"MessageId"`
	// Message shall contain an optional human readable message.
	Message string
	// MessageArgs shall contain the message substitution
	// arguments for the specific message referenced by the MessageID and
	// shall only be included if the MessageID is present.  Number and
	// integer type arguments shall be converted to strings.
	MessageArgs []string
	// RelatedProperties shall contain an array of JSON
	// Pointers indicating the properties described by the message, if
	// appropriate for the message.
	RelatedProperties []string `json:"RelatedProperties,omitempty"`
	// Severity is The value of this property shall be the severity of the
	// error, as defined in the Status section of the Redfish specification.
	Severity string
	// Resolution shall contain an override of the
	// Resolution of the message in message registry, if present.
	Resolution string
}

// Message is This type shall define a Message as described in the
// Redfish specification.
type Message struct {
	Entity
	MessageExtendedInfo
}

// GetMessage will get a Message instance from the service using the client's cached context.
func GetMessage(c Client, uri string, queryOpts ...QueryGroupOption) (*Message, error) {
	return GetMessageWithContext(ContextOf(c), c, uri, queryOpts...)
}

// GetMessageWithContext will get a Message instance from the service.
func GetMessageWithContext(ctx context.Context, c Client, uri string, queryOpts ...QueryGroupOption) (*Message, error) {
	return GetObjectWithContext[Message](ctx, c, uri, queryOpts...)
}

// ListReferencedMessages gets the collection of Message from
// a provided reference using the client's cached context.
func ListReferencedMessages(c Client, link string, queryOpts ...QueryGroupOption) ([]*Message, error) {
	return ListReferencedMessagesWithContext(ContextOf(c), c, link, queryOpts...)
}

// ListReferencedMessagesWithContext gets the collection of Message from
// a provided reference.
func ListReferencedMessagesWithContext(ctx context.Context, c Client, link string, queryOpts ...QueryGroupOption) ([]*Message, error) {
	return GetCollectionObjectsWithContext[Message](ctx, c, link, queryOpts...)
}
