//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// MessageRegistryFileLocation is a location
// information for the Message Registry file.
type MessageRegistryFileLocation struct {
	Language string `json:"Language"`
	URI      string `json:"Uri"`
}

// MessageRegistryFile describes the Message Registry file locator Resource.
type MessageRegistryFile struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Languages is the RFC5646-conformant language codes for the
	// available Message Registries.
	Languages []string
	// Location is the location information for this Message Registry file.
	Location []MessageRegistryFileLocation
	// Registry shall contain the Message Registry name and it major and
	// minor versions, as defined by the Redfish Specification.
	Registry string
}

// GetMessageRegistryFile will get a MessageRegistryFile
// instance from the Redfish service.
func GetMessageRegistryFile(
	c common.Client,
	uri string,
) (*MessageRegistryFile, error) {
	var messageRegistryFile MessageRegistryFile
	return &messageRegistryFile, messageRegistryFile.Get(c, uri, &messageRegistryFile)
}

// ListReferencedMessageRegistryFiles gets the collection of MessageRegistryFile.
func ListReferencedMessageRegistryFiles(c common.Client, link string) ([]*MessageRegistryFile, error) { //nolint:dupl
	var result []*MessageRegistryFile
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *MessageRegistryFile
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		messageregistryfile, err := GetMessageRegistryFile(c, link)
		ch <- GetResult{Item: messageregistryfile, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
