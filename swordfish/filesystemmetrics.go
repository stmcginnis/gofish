//
// SPDX-License-Identifier: BSD-3-Clause
//

//nolint:dupl
package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// FileSystemMetrics shall contain the usage and health statistics for a file system in a Redfish implementation.
type FileSystemMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IOStatistics shall represent IO statistics for this file system.
	IOStatistics IOStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetFileSystemMetrics will get a FileSystemMetrics instance from the service.
func GetFileSystemMetrics(c common.Client, uri string) (*FileSystemMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var filesystemmetrics FileSystemMetrics
	err = json.NewDecoder(resp.Body).Decode(&filesystemmetrics)
	if err != nil {
		return nil, err
	}

	filesystemmetrics.SetClient(c)
	return &filesystemmetrics, nil
}

// ListReferencedFileSystemMetricss gets the collection of FileSystemMetrics from
// a provided reference.
func ListReferencedFileSystemMetricss(c common.Client, link string) ([]*FileSystemMetrics, error) {
	var result []*FileSystemMetrics
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *FileSystemMetrics
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		filesystemmetrics, err := GetFileSystemMetrics(c, link)
		ch <- GetResult{Item: filesystemmetrics, Link: link, Error: err}
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
