//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// LineOfService This service option is the abstract base class for other ClassOfService and concrete lines of
// service.
type LineOfService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetLineOfService will get a LineOfService instance from the service.
func GetLineOfService(c common.Client, uri string) (*LineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lineofservice LineOfService
	err = json.NewDecoder(resp.Body).Decode(&lineofservice)
	if err != nil {
		return nil, err
	}

	lineofservice.SetClient(c)
	return &lineofservice, nil
}

// ListReferencedLineOfServices gets the collection of LineOfService from
// a provided reference.
func ListReferencedLineOfServices(c common.Client, link string) ([]*LineOfService, error) { //nolint:dupl
	if link == "" {
		return nil, nil
	}

	type GetResult struct {
		Item  *LineOfService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		lineofservice, err := GetLineOfService(c, link)
		ch <- GetResult{Item: lineofservice, Link: link, Error: err}
	}

	var links []string
	var err error
	go func() {
		links, err = common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	// Save unordered results into link-to-LineOfService helper map.
	unorderedResults := map[string]*LineOfService{}
	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			unorderedResults[r.Link] = r.Item
		}
	}

	if !collectionError.Empty() {
		return nil, collectionError
	}
	// Build the final ordered slice based on the original order from the links list.
	results := make([]*LineOfService, len(links))
	for i, link := range links {
		results[i] = unorderedResults[link]
	}

	return results, nil
}
