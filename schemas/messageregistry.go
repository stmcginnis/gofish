//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #MessageRegistry.v1_7_0.MessageRegistry

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
)

const MessageIDSectionLength = 4

type ClearingType string

const (
	// SameOriginOfConditionClearingType shall indicate that a logged event is
	// cleared by a message if the 'OriginOfCondition' for both events are the
	// same.
	SameOriginOfConditionClearingType ClearingType = "SameOriginOfCondition"
)

type ParamType string

const (
	// stringParamType The argument is a string.
	StringParamType ParamType = "string"
	// numberParamType The argument is a number converted to a string.
	NumberParamType ParamType = "number"
)

// MessageRegistry shall represent a message registry for a Redfish
// implementation.
type MessageRegistry struct {
	Entity
	// Language shall contain an RFC5646-conformant language code.
	Language string
	// Messages shall contain the message keys contained in the message registry.
	// The message keys are the suffix of the 'MessageId' and shall be unique
	// within this message registry.
	Messages map[string]MessageRegistryMessage
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OwningEntity shall represent the publisher of this message registry.
	OwningEntity string
	// RegistryPrefix shall contain the Redfish Specification-defined prefix used
	// in forming and decoding 'MessageId' values that uniquely identifies all
	// messages that belong to this message registry.
	RegistryPrefix string
	// RegistryVersion shall contain the version of this message registry.
	RegistryVersion string
	// Release shall contain the version of the release bundle that first included
	// this revision of the message registry file. Message registry files from the
	// DMTF Redfish Forum shall use the release bundle version of DSP8011 that
	// first included the minor revision level of this message registry file,
	// ignoring errata releases, as the value of this property.
	//
	// Version added: v1.7.0
	Release string
}

// GetMessageRegistry will get a MessageRegistry instance from the service.
func GetMessageRegistry(c Client, uri string) (*MessageRegistry, error) {
	return GetObject[MessageRegistry](c, uri)
}

// ListReferencedMessageRegistries gets the collection of MessageRegistry from
// a provided reference.
func ListReferencedMessageRegistries(c Client, link string) ([]*MessageRegistry, error) {
	var result []*MessageRegistry
	links, err := GetCollection(c, link)
	if err != nil {
		return nil, err
	}

	// TODO: Look at what to do to make parallel
	for _, sLink := range links.ItemLinks {
		mrf, err := GetMessageRegistryFile(c, sLink)
		if err != nil {
			return nil, err
		}

		// get message registry from all location
		for _, location := range mrf.Location {
			mr, err := GetMessageRegistry(c, location.URI)
			if err != nil {
				return nil, err
			}
			// TODO: exclude attribute and privilege registries
			result = append(result, mr)
		}
	}

	return result, nil
}

// ListReferencedMessageRegistriesByLanguage gets the collection of MessageRegistry.
// language is the RFC5646-conformant language code for the message registry.
func ListReferencedMessageRegistriesByLanguage(c Client, link, language string) ([]*MessageRegistry, error) {
	language = strings.TrimSpace(language)
	if language == "" {
		return nil, fmt.Errorf("received empty language")
	}

	// TODO: Looks at what to do to make parallel.
	var result []*MessageRegistry
	links, err := GetCollection(c, link)
	if err != nil {
		return nil, err
	}

	for _, sLink := range links.ItemLinks {
		mrf, err := GetMessageRegistryFile(c, sLink)
		if err != nil {
			return nil, err
		}

		// get message registry by language
		for _, location := range mrf.Location {
			if location.Language == language {
				mr, err := GetMessageRegistry(c, location.URI)
				if err != nil {
					return nil, err
				}
				result = append(result, mr)
			}
		}
	}

	return result, nil
}

// GetMessageRegistryByLanguage gets the message registry by language.
// registry is used to identify the correct Message Registry
// file (MessageRegistryFile.Registry) and it shall contain the
// Message Registry name and it major and minor versions, as defined
// by the Redfish Specification.
// language is the RFC5646-conformant language code for the message registry.
func GetMessageRegistryByLanguage(
	c Client,
	link string,
	registry string,
	language string,
) (*MessageRegistry, error) {
	registry = strings.TrimSpace(registry)
	if registry == "" {
		return nil, fmt.Errorf("received empty registry")
	}

	language = strings.TrimSpace(language)
	if language == "" {
		return nil, fmt.Errorf("received empty language")
	}

	// TODO: Look at what to do to make parallel
	links, err := GetCollection(c, link)
	if err != nil {
		return nil, err
	}

	for _, sLink := range links.ItemLinks {
		s, err := GetMessageRegistryFile(c, sLink)
		if err != nil {
			return nil, err
		}
		// search for the correct registry
		if s.Registry == registry {
			// search for the correct location
			for _, location := range s.Location {
				if location.Language == language {
					return GetMessageRegistry(c, location.URI)
				}
			}
		}
	}

	return nil, fmt.Errorf("message registry not found")
}

// GetMessageFromMessageRegistryByLanguage tries to find and get the message
// from the informed messageID.
// messageID is the key used to find the registry, version and message:
// Example of messageID: Alert.1.0.LanDisconnect
//
//   - The segment before the 1st period is the Registry Name (Registry Prefix): Alert
//   - The segment between the 1st and 2nd period is the major version: 1
//   - The segment between the 2nd and 3rd period is the minor version: 0
//   - The segment after the 3rd period is the Message Identifier in the Registry: LanDisconnect
//
// language is the RFC5646-conformant language code for the message registry.
// Example of language: en
func GetMessageFromMessageRegistryByLanguage(
	c Client,
	link string,
	messageID string,
	language string,
) (*MessageRegistryMessage, error) {
	messageID = strings.TrimSpace(messageID)
	if messageID == "" {
		return nil, fmt.Errorf("received empty messageID")
	}

	language = strings.TrimSpace(language)
	if language == "" {
		return nil, fmt.Errorf("received empty language")
	}

	// split messageID
	messageIDSplitted := strings.Split(messageID, ".")

	// validate messageID
	if len(messageIDSplitted) != MessageIDSectionLength {
		return nil, fmt.Errorf("received invalid messageID %s", messageID)
	}

	// get information from the messageID
	registryPrefix := messageIDSplitted[0]
	registryMajorVersion := messageIDSplitted[1]
	registryMinorVersion := messageIDSplitted[2]
	registryMajorMinorVersion := registryMajorVersion + "." + registryMinorVersion
	registryMessageKey := messageIDSplitted[3]

	allMessageRegistryByLanguage, err := ListReferencedMessageRegistriesByLanguage(c, link, language)
	if err != nil {
		return nil, err
	}
	for _, mr := range allMessageRegistryByLanguage {
		if mr.RegistryPrefix == registryPrefix &&
			strings.HasPrefix(mr.RegistryVersion, registryMajorMinorVersion) {
			if m, ok := mr.Messages[registryMessageKey]; ok {
				return &m, nil
			}
		}
	}

	return nil, fmt.Errorf("message not found")
}

// ClearingLogic shall contain the clearing logic associated with a message.
type ClearingLogic struct {
	// ClearsAll shall indicate whether all logged events containing messages from
	// this message registry are cleared when this message is received. If
	// conditional properties are present, such as the 'ClearsIf' property, the
	// specified conditions are required to clear the logged events.
	//
	// Version added: v1.2.0
	ClearsAll bool
	// ClearsIf shall contain the condition required to clear the logged events
	// specified by other properties in this object when this message is received.
	// If not present, no condition is checked prior to clearing logged events when
	// this message is received.
	//
	// Version added: v1.2.0
	ClearsIf ClearingType
	// ClearsMessage shall contain an array of message keys for logged events that
	// are cleared when this message is received. If conditional properties are
	// present, such as the 'ClearsIf' property, the specified conditions shall be
	// required to clear the logged events with these message keys. This property
	// shall contain message keys, without message registry names and versions, as
	// defined in the 'MessageId format' clause of the Redfish Specification. This
	// property shall not reference message keys in other message registries.
	//
	// Version added: v1.2.0
	ClearsMessage []string
}

// MessageRegistryMessage shall represent how a message is defined within a message registry.
type MessageRegistryMessage struct {
	// ArgDescriptions shall contain an ordered array of text describing each
	// argument used as substitution in the message.
	//
	// Version added: v1.3.0
	ArgDescriptions []string
	// ArgLongDescriptions shall contain an ordered array of normative language for
	// each argument used as substitution in the message.
	//
	// Version added: v1.3.0
	ArgLongDescriptions []string
	// ClearingLogic shall contain the clearing logic associated with this message.
	// Clearing in this context deasserts the event rather than removes the event
	// from a log.
	//
	// Version added: v1.2.0
	ClearingLogic ClearingLogic
	// Deprecated shall indicate that a message is deprecated. The value of the
	// string should explain the deprecation, including reference to a new message
	// or messages to be used. The message can be supported in new and existing
	// implementations, but usage in new implementations is discouraged. Deprecated
	// messages are likely to be removed in a future major version of the message
	// registry. The 'ReplacedBy' property may be used to provide a reference to a
	// replacement message definition.
	//
	// Version added: v1.5.0
	Deprecated string
	// Description provides a description of this resource.
	Description string
	// LongDescription shall contain the normative language that describes this
	// message's usage in a Redfish implementation.
	//
	// Version added: v1.3.0
	LongDescription string
	// MapsToGeneralMessages shall indicate that this message maps to general or
	// less-specific messages that duplicates information about the condition that
	// generated this message. Services may issue the referenced messages along
	// with this message to provide consistency for clients. The array shall
	// contain 'MessageRegistryPrefix.MessageKey' formatted values that describe
	// the message registry and message key used to identify the messages.
	//
	// Version added: v1.6.0
	MapsToGeneralMessages []string
	// Message shall contain the message to display. If a %integer is included in
	// part of the string, it shall represent a string substitution for any
	// 'MessageArgs' that accompany the message, in order.
	Message string
	// MessageSeverity shall contain the severity of the message. Services can
	// replace the severity defined in the message registry with a value more
	// applicable to the implementation in message payloads and event payloads.
	//
	// Version added: v1.4.0
	MessageSeverity string
	// NumberOfArgs shall contain the number of arguments that are substituted for
	// the locations marked with %<integer> in the message.
	NumberOfArgs uint
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ParamTypes shall contain an ordered array of the data types of the values in
	// 'MessageArgs', prior to their conversion to strings for inclusion in a
	// message.
	ParamTypes []ParamType
	// ReplacedBy shall contain the message registry and message key, in the
	// 'MessageRegistryPrefix.MessageKey' format, that identifies the message that
	// replaces this message. This property may be used to indicate replacement for
	// a deprecated message, including cases where a standardized version replaces
	// an OEM-created message.
	//
	// Version added: v1.6.0
	ReplacedBy string
	// Resolution shall contain the resolution of the message. Services can replace
	// the resolution defined in the message registry with a more specific
	// resolution in message payloads.
	Resolution string
	// Severity shall contain the severity of the condition resulting in the
	// message, as defined in the 'Status' clause of the Redfish Specification.
	// Services can replace the severity defined in the message registry with a
	// value more applicable to the implementation in message payloads and event
	// payloads.
	//
	// Deprecated: v1.4.0
	// This property has been deprecated in favor of 'MessageSeverity', which ties
	// the values to the enumerations defined for the 'Health' property within
	// 'Status'.
	Severity string
	// VersionAdded shall contain the version of the message registry when the
	// message was added. This property shall not appear for messages created at
	// version '1.0.0' of a message registry.
	//
	// Version added: v1.5.0
	VersionAdded string
	// VersionDeprecated shall contain the version of the registry when the message
	// was deprecated. This property shall not appear if the message has not been
	// deprecated.
	//
	// Version added: v1.5.0
	VersionDeprecated string
}
