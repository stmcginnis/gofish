//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// IOConnectivityLineOfService is used to specify the characteristics of
// storage connectivity.
type IOConnectivityLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessProtocols shall specify the Access protocol for this service
	// option. NOTE: If multiple protocols are specified,  the corresponding
	// MaxSupportedIOPS governs the max achieved across all protocol uses. This
	// may be less than the sum of the individual max values, which may be
	// specified by individual Line of Service entries.
	AccessProtocols []common.Protocol
	// Description provides a description of this resource.
	Description string
	// MaxBytesPerSecond shall be the maximum bytes per second that a connection
	// can support.
	MaxBytesPerSecond int64
	// MaxIOPS shall be the maximum IOs per second that the connection shall
	// allow for the selected access protocol.
	MaxIOPS int
}

// GetIOConnectivityLineOfService will get a IOConnectivityLineOfService instance from the service.
func GetIOConnectivityLineOfService(c common.Client, uri string) (*IOConnectivityLineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ioconnectivitylineofservice IOConnectivityLineOfService
	err = json.NewDecoder(resp.Body).Decode(&ioconnectivitylineofservice)
	if err != nil {
		return nil, err
	}

	ioconnectivitylineofservice.SetClient(c)
	return &ioconnectivitylineofservice, nil
}

// ListReferencedIOConnectivityLineOfServices gets the collection of IOConnectivityLineOfService from
// a provided reference.
func ListReferencedIOConnectivityLineOfServices(c common.Client, link string) ([]*IOConnectivityLineOfService, error) { //nolint:dupl
	var result []*IOConnectivityLineOfService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *IOConnectivityLineOfService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		ioconnectivitylineofservice, err := GetIOConnectivityLineOfService(c, link)
		ch <- GetResult{Item: ioconnectivitylineofservice, Link: link, Error: err}
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
