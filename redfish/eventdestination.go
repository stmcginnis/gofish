//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// EventDestinationProtocol is the communication protocol of the event destination.
type EventDestinationProtocol string

const (

	// RedfishEventDestinationProtocol The destination follows the Redfish
	// specification for event notifications.
	RedfishEventDestinationProtocol EventDestinationProtocol = "Redfish"
)

// SubscriptionType is the type of subscription used.
type SubscriptionType string

const (

	// RedfishEventSubscriptionType The subscription follows the Redfish
	// specification for event notifications, which is done by a service
	// sending an HTTP POST to the subscriber's destination URI.
	RedfishEventSubscriptionType SubscriptionType = "RedfishEvent"
	// SSESubscriptionType The subscription follows the HTML5 Server-Sent
	// Event definition for event notifications.
	SSESubscriptionType SubscriptionType = "SSE"
)

// EventDestination is used to represent the target of an event
// subscription, including the types of events subscribed and context to
// provide to the target in the Event payload.
type EventDestination struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Context shall contain a client supplied context that will remain with the
	// connection through the connections lifetime.
	Context string
	// Description provides a description of this resource.
	Description string
	// Destination shall contain a URI to the destination where the events will
	// be sent.
	Destination string
	// EventFormatType shall indicate the the content types of the message that
	// this service will send to the EventDestination. If this property is not
	// present, the EventFormatType shall be assumed to be Event.
	EventFormatType EventFormatType
	// HTTPHeaders shall contain an object consisting of the names and values of
	// of HTTP header to be included with every event POST to the Event
	// Destination. This property shall be null or an empty array on a GET. An
	// empty array is the preferred return value on GET.
	HTTPHeaders []HTTPHeaderProperty `json:"HttpHeaders"`
	// MessageIDs shall specify an array of MessageIds that are the only
	// allowable values for the MessageId property within an EventRecord sent to
	// the subscriber. Events with MessageIds not contained in this array shall
	// not be sent to the subscriber. If this property is absent or the array is
	// empty, the service shall send Events with any MessageId to the subscriber.
	MessageIDs []string `json:"MessageIds"`
	// originResources shall specify an array of Resources, Resource Collections,
	// or Referenceable Members that are the only allowable values for the
	// OriginOfCondition property within an EventRecord sent to the subscriber.
	// Events originating from Resources, Resource Collections, or Referenceable
	// Members not contained in this array shall not be sent to the subscriber.
	// If this property is absent or the array is empty, the service shall send
	// Events originating from any Resource, Resource Collection, or
	// Referenceable Member to the subscriber.
	originResources []string
	// OriginResourcesCount is the number of OriginResources.
	OriginResourcesCount int `json:"OriginResources@odata.count"`
	// Protocol is used to indicate that the event type shall adhere to that
	// defined in the Redfish specification.
	Protocol EventDestinationProtocol
	// RegistryPrefixes is The value of this property is the array of the
	// Prefixes of the Message Registries that contain the MessageIds in the
	// Events that shall be sent to the EventDestination. If this property
	// is absent or the array is empty, the service shall send Events with
	// MessageIds from any Message Registry.
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of Resource Type values. When an
	// event is generated, if the OriginOfCondition's Resource Type matches a
	// value in this array, the event shall be sent to the event destination
	// (unless it would be filtered by other property conditions such as
	// RegistryPrefix). If this property is absent or the array is empty, the
	// service shall send Events from any Resource Type to the subscriber. The
	// value of this property shall be only the general namespace for the type
	// and not the versioned value. For example, it shall not be Task.v1_2_0.Task
	// and instead shall just be Task. To specify that a client is subscribing
	// for Metric Reports, the EventTypes property should include 'MetricReport'.
	ResourceTypes []string
	// SubordinateResources is When set to true and OriginResources is
	// specified, indicates the subscription shall be for events from the
	// OriginsResources specified and all subordinate resources. When set to
	// false and OriginResources is specified, indicates subscription shall
	// be for events only from the OriginResources. If OriginResources is
	// not specified, it has no relevance.
	SubordinateResources bool
	// SubscriptionType shall indicate the type of subscription for events. If
	// this property is not present, the SubscriptionType shall be assumed to be
	// RedfishEvent.
	SubscriptionType SubscriptionType
}

// UnmarshalJSON unmarshals a EventDestination object from the raw JSON.
func (eventdestination *EventDestination) UnmarshalJSON(b []byte) error {
	type temp EventDestination
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*eventdestination = EventDestination(t.temp)

	return nil
}

// GetEventDestination will get a EventDestination instance from the service.
func GetEventDestination(c common.Client, uri string) (*EventDestination, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventdestination EventDestination
	err = json.NewDecoder(resp.Body).Decode(&eventdestination)
	if err != nil {
		return nil, err
	}

	eventdestination.SetClient(c)
	return &eventdestination, nil
}

// ListReferencedEventDestinations gets the collection of EventDestination from
// a provided reference.
func ListReferencedEventDestinations(c common.Client, link string) ([]*EventDestination, error) {
	var result []*EventDestination
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, eventdestinationLink := range links.ItemLinks {
		eventdestination, err := GetEventDestination(c, eventdestinationLink)
		if err != nil {
			return result, err
		}
		result = append(result, eventdestination)
	}

	return result, nil
}

// HTTPHeaderProperty shall a names and value of an HTTP header to be included
// with every event POST to the Event Destination.
type HTTPHeaderProperty map[string][]string
