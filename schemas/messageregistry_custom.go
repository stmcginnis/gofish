//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"fmt"
	"strings"
)

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

const (
	MessageIDSectionLength = 4
)
