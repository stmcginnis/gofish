//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0.5 - #ClassOfService.v1_2_0.ClassOfService

package schemas

import (
	"encoding/json"
)

// ClassOfService shall define a service option composed of one or more line of
// service entities. ITIL defines a service option as a choice of utility or
// warranty for a service.
type ClassOfService struct {
	Entity
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
	// DataProtectionLinesOfServiceCount
	DataProtectionLinesOfServiceCount int `json:"DataProtectionLinesOfService@odata.count"`
	// DataSecurityLinesOfService shall be a set of data security service options.
	//
	// Version added: v1.1.1
	dataSecurityLinesOfService []string
	// DataSecurityLinesOfServiceCount
	DataSecurityLinesOfServiceCount int `json:"DataSecurityLinesOfService@odata.count"`
	// DataStorageLinesOfService shall be a set of data protection service options.
	//
	// Version added: v1.1.1
	dataStorageLinesOfService []string
	// DataStorageLinesOfServiceCount
	DataStorageLinesOfServiceCount int `json:"DataStorageLinesOfService@odata.count"`
	// IOConnectivityLinesOfService shall be a set of IO connectivity service
	// options. Within a class of service, at most one IO connectivity service
	// option may be present for a value of AccessProtocol.
	//
	// Version added: v1.1.1
	iOConnectivityLinesOfService []string
	// IOConnectivityLinesOfServiceCount
	IOConnectivityLinesOfServiceCount int `json:"IOConnectivityLinesOfService@odata.count"`
	// IOPerformanceLinesOfService shall be a set of IO performance service
	// options.
	//
	// Version added: v1.1.1
	iOPerformanceLinesOfService []string
	// IOPerformanceLinesOfServiceCount
	IOPerformanceLinesOfServiceCount int `json:"IOPerformanceLinesOfService@odata.count"`
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ClassOfService object from the raw JSON.
func (c *ClassOfService) UnmarshalJSON(b []byte) error {
	type temp ClassOfService
	var tmp struct {
		temp
		DataProtectionLinesOfService Links `json:"DataProtectionLinesOfService"`
		DataSecurityLinesOfService   Links `json:"DataSecurityLinesOfService"`
		DataStorageLinesOfService    Links `json:"DataStorageLinesOfService"`
		IOConnectivityLinesOfService Links `json:"IOConnectivityLinesOfService"`
		IOPerformanceLinesOfService  Links `json:"IOPerformanceLinesOfService"`
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
	c.iOConnectivityLinesOfService = tmp.IOConnectivityLinesOfService.ToStrings()
	c.iOPerformanceLinesOfService = tmp.IOPerformanceLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ClassOfService) Update() error {
	readWriteFields := []string{
		"ClassOfServiceVersion",
		"DataProtectionLinesOfService",
		"DataSecurityLinesOfService",
		"DataStorageLinesOfService",
		"IOConnectivityLinesOfService",
		"IOPerformanceLinesOfService",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetClassOfService will get a ClassOfService instance from the service.
func GetClassOfService(c Client, uri string) (*ClassOfService, error) {
	return GetObject[ClassOfService](c, uri)
}

// ListReferencedClassOfServices gets the collection of ClassOfService from
// a provided reference.
func ListReferencedClassOfServices(c Client, link string) ([]*ClassOfService, error) {
	return GetCollectionObjects[ClassOfService](c, link)
}

// DataProtectionLinesOfService gets the DataProtectionLinesOfService linked resources.
func (c *ClassOfService) DataProtectionLinesOfService() ([]*DataProtectionLineOfService, error) {
	return GetObjects[DataProtectionLineOfService](c.client, c.dataProtectionLinesOfService)
}

// DataSecurityLinesOfService gets the DataSecurityLinesOfService linked resources.
func (c *ClassOfService) DataSecurityLinesOfService() ([]*DataSecurityLineOfService, error) {
	return GetObjects[DataSecurityLineOfService](c.client, c.dataSecurityLinesOfService)
}

// DataStorageLinesOfService gets the DataStorageLinesOfService linked resources.
func (c *ClassOfService) DataStorageLinesOfService() ([]*DataStorageLineOfService, error) {
	return GetObjects[DataStorageLineOfService](c.client, c.dataStorageLinesOfService)
}

// IOConnectivityLinesOfService gets the IOConnectivityLinesOfService linked resources.
func (c *ClassOfService) IOConnectivityLinesOfService() ([]*IOConnectivityLineOfService, error) {
	return GetObjects[IOConnectivityLineOfService](c.client, c.iOConnectivityLinesOfService)
}

// IOPerformanceLinesOfService gets the IOPerformanceLinesOfService linked resources.
func (c *ClassOfService) IOPerformanceLinesOfService() ([]*IOPerformanceLineOfService, error) {
	return GetObjects[IOPerformanceLineOfService](c.client, c.iOPerformanceLinesOfService)
}
