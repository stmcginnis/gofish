//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

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
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a IOConnectivityLineOfService object from the raw JSON.
func (ioconnectivitylineofservice *IOConnectivityLineOfService) UnmarshalJSON(b []byte) error {
	type temp IOConnectivityLineOfService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ioconnectivitylineofservice = IOConnectivityLineOfService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	ioconnectivitylineofservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ioconnectivitylineofservice *IOConnectivityLineOfService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(IOConnectivityLineOfService)
	original.UnmarshalJSON(ioconnectivitylineofservice.rawData)

	readWriteFields := []string{
		"AccessProtocols",
		"MaxBytesPerSecond",
		"MaxIOPS",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ioconnectivitylineofservice).Elem()

	return ioconnectivitylineofservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetIOConnectivityLineOfService will get a IOConnectivityLineOfService instance from the service.
func GetIOConnectivityLineOfService(c common.Client, uri string) (*IOConnectivityLineOfService, error) {
	return common.GetObject[IOConnectivityLineOfService](c, uri)
}

// ListReferencedIOConnectivityLineOfServices gets the collection of IOConnectivityLineOfService from
// a provided reference.
func ListReferencedIOConnectivityLineOfServices(c common.Client, link string) ([]*IOConnectivityLineOfService, error) {
	return common.GetCollectionObjects[IOConnectivityLineOfService](c, link)
}
