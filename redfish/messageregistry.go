//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

const MessageIDSectionLength = 4

// MessageRegistryMessage is a message contained in the message registry.
type MessageRegistryMessage struct {
	// Description is a short description of how and when to use this message.
	Description string
	// Message is the actual message.
	// This property shall contain the message to display.  If a %integer is
	// included in part of the string, it shall represent a string substitution
	// for any MessageArgs that accompany the message, in order.
	Message string
	// Severity property shall contain the severity of the condition resulting in
	// the message, as defined in the Status clause of the Redfish Specification.
	// This property has been deprecated in favor of MessageSeverity, which ties
	// the values to the enumerations defined for the Health property within Status.
	Severity string
	// MessageSeverity is the severity of the message. This property shall contain
	// the severity of the message.
	MessageSeverity string
	// NumberOfArgs is the number of arguments in the message.
	// This property shall contain the number of arguments that are substituted
	// for the locations marked with %<integer> in the message.
	NumberOfArgs int
	// ParamTypes are the MessageArg types, in order, for the message.
	ParamTypes []string
	// Resolution is used to provide suggestions on how to resolve the situation
	// that caused the error.
	Resolution string
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem interface{}
}

// MessageRegistry schema describes all message registries.
// It represents the properties for the message registries themselves.
type MessageRegistry struct {
	common.Entity

	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// RegistryPrefix is the single-word prefix that is used in forming and decoding MessageIds.
	RegistryPrefix string
	// RegistryVersion is the message registry version in the middle portion of a MessageId.
	RegistryVersion string
	// Language is the RFC5646-conformant language code for the message registry.
	Language string
	// OwningEntity ins the organization or company that publishes this message registry.
	OwningEntity string
	// Messages are the messages for the message registry.
	Messages map[string]MessageRegistryMessage
}

// GetMessageRegistry will get a MessageRegistry instance from the Redfish service.
func GetMessageRegistry(
	c common.Client,
	uri string,
) (*MessageRegistry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t MessageRegistry
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedMessageRegistries gets the collection of MessageRegistry.
func ListReferencedMessageRegistries(
	c common.Client,
	link string,
) ([]*MessageRegistry, error) {
	var result []*MessageRegistry
	links, err := common.GetCollection(c, link)
	if err != nil {
		return nil, err
	}

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
			result = append(result, mr)
		}
	}

	return result, nil
}

// ListReferencedMessageRegistriesByLanguage gets the collection of MessageRegistry.
// language is the RFC5646-conformant language code for the message registry.
func ListReferencedMessageRegistriesByLanguage(
	c common.Client,
	link string,
	language string,
) ([]*MessageRegistry, error) {
	language = strings.TrimSpace(language)
	if language == "" {
		return nil, fmt.Errorf("received empty language")
	}

	var result []*MessageRegistry
	links, err := common.GetCollection(c, link)
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
	c common.Client,
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

	links, err := common.GetCollection(c, link)
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
//  - The segment before the 1st period is the Registry Name (Registry Prefix): Alert
//  - The segment between the 1st and 2nd period is the major version: 1
//  - The segment between the 2nd and 3rd period is the minor version: 0
//  - The segment after the 3rd period is the Message Identifier in the Registry: LanDisconnect
// language is the RFC5646-conformant language code for the message registry.
// Example of language: en
func GetMessageFromMessageRegistryByLanguage(
	c common.Client,
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
