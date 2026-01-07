//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0.5 - #ClassOfService.v1_2_0.ClassOfService

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ClassOfService shall define a service option composed of one or more line of
// service entities. ITIL defines a service option as a choice of utility or
// warranty for a service.
type ClassOfService struct {
	common.Entity
	// ClassOfServiceVersion shall be in the form: M + '.' + N + '.' + U Where: M -
	// The major version (in numeric form). N - The minor version (in numeric
	// form). U - The update (e.g. errata or patch in numeric form).
	ClassOfServiceVersion string
	// DataProtectionLinesOfService shall be a set of data protection service
	// options. Within a class of service, one data protection service option shall
	// be present for each replication session.
	//
	// Version added: v1.1.1
	dataProtectionLinesOfService []string
	// DataProtectionLinesOfService@odata.count
	DataProtectionLinesOfServiceCount int `json:"DataProtectionLinesOfService@odata.count"`
	// DataSecurityLinesOfService shall be a set of data security service options.
	//
	// Version added: v1.1.1
	dataSecurityLinesOfService []string
	// DataSecurityLinesOfService@odata.count
	DataSecurityLinesOfServiceCount int `json:"DataSecurityLinesOfService@odata.count"`
	// DataStorageLinesOfService shall be a set of data protection service options.
	//
	// Version added: v1.1.1
	dataStorageLinesOfService []string
	// DataStorageLinesOfService@odata.count
	DataStorageLinesOfServiceCount int `json:"DataStorageLinesOfService@odata.count"`
	// IOConnectivityLinesOfService shall be a set of IO connectivity service
	// options. Within a class of service, at most one IO connectivity service
	// option may be present for a value of AccessProtocol.
	//
	// Version added: v1.1.1
	ioConnectivityLinesOfService []string
	// IOConnectivityLinesOfService@odata.count
	IOConnectivityLinesOfServiceCount int `json:"IOConnectivityLinesOfService@odata.count"`
	// IOPerformanceLinesOfService shall be a set of IO performance service
	// options.
	//
	// Version added: v1.1.1
	ioPerformanceLinesOfService []string
	// IOPerformanceLinesOfService@odata.count
	IOPerformanceLinesOfServiceCount int `json:"IOPerformanceLinesOfService@odata.count"`
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ClassOfService object from the raw JSON.
func (c *ClassOfService) UnmarshalJSON(b []byte) error {
	type temp ClassOfService
	var tmp struct {
		temp
		DataProtectionLinesOfService common.Links
		DataSecurityLinesOfService   common.Links
		DataStorageLinesOfService    common.Links
		IOConnectivityLinesOfService common.Links
		IOPerformanceLinesOfService  common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ClassOfService(tmp.temp)

	// Extract the links to other entities for later
	c.dataProtectionLinesOfService = tmp.DataProtectionLinesOfService.ToStrings()
	c.dataSecurityLinesOfService = tmp.DataSecurityLinesOfService.ToStrings()
	c.dataStorageLinesOfService = tmp.DataStorageLinesOfService.ToStrings()
	c.ioConnectivityLinesOfService = tmp.IOConnectivityLinesOfService.ToStrings()
	c.ioPerformanceLinesOfService = tmp.IOPerformanceLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ClassOfService) Update() error {
	readWriteFields := []string{
		"ClassOfServiceVersion",
		// "DataProtectionLinesOfService",
		// "DataProtectionLinesOfService@odata.count",
		// "DataSecurityLinesOfService",
		// "DataSecurityLinesOfService@odata.count",
		// "DataStorageLinesOfService",
		// "DataStorageLinesOfService@odata.count",
		// "IOConnectivityLinesOfService",
		// "IOConnectivityLinesOfService@odata.count",
		// "IOPerformanceLinesOfService",
		// "IOPerformanceLinesOfService@odata.count",
		"Identifier",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
}

// GetClassOfService will get a ClassOfService instance from the service.
func GetClassOfService(c common.Client, uri string) (*ClassOfService, error) {
	return common.GetObject[ClassOfService](c, uri)
}

// ListReferencedClassOfServices gets the collection of ClassOfService from
// a provided reference.
func ListReferencedClassOfServices(c common.Client, link string) ([]*ClassOfService, error) {
	return common.GetCollectionObjects[ClassOfService](c, link)
}

// DataProtectionLinesOfServices gets the DataProtectionLinesOfService that are
// part of this ClassOfService.
func (c *ClassOfService) DataProtectionLinesOfServices() ([]*DataProtectionLineOfService, error) {
	return common.GetObjects[DataProtectionLineOfService](c.GetClient(), c.dataProtectionLinesOfService)
}

// DataSecurityLinesOfServices gets the DataSecurityLinesOfService that are
// part of this ClassOfService.
func (c *ClassOfService) DataSecurityLinesOfServices() ([]*DataSecurityLineOfService, error) {
	return common.GetObjects[DataSecurityLineOfService](c.GetClient(), c.dataSecurityLinesOfService)
}

// DataStorageLinesOfServices gets the DataStorageLinesOfService that are
// part of this ClassOfService.
func (c *ClassOfService) DataStorageLinesOfServices() ([]*DataStorageLineOfService, error) {
	return common.GetObjects[DataStorageLineOfService](c.GetClient(), c.dataStorageLinesOfService)
}

// IOConnectivityLinesOfServices gets the IOConnectivityLinesOfService that are
// part of this ClassOfService.
func (c *ClassOfService) IOConnectivityLinesOfServices() ([]*IOConnectivityLineOfService, error) {
	return common.GetObjects[IOConnectivityLineOfService](c.GetClient(), c.ioConnectivityLinesOfService)
}

// IOPerformanceLinesOfServices gets the IOPerformanceLinesOfService that are
// part of this ClassOfService.
func (c *ClassOfService) IOPerformanceLinesOfServices() ([]*IOPerformanceLineOfService, error) {
	return common.GetObjects[IOPerformanceLineOfService](c.GetClient(), c.ioPerformanceLinesOfService)
}
