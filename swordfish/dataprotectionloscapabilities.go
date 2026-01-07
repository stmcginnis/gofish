//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.1c - #DataProtectionLoSCapabilities.v1_2_0.DataProtectionLoSCapabilities

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// FailureDomainScope is The enumeration literals of this enumeration shall
// represent a geographic scope in which all components within that scope have
// similar vulnerabilities.
type FailureDomainScope string

const (
	// ServerFailureDomainScope Components of a CPU/memory complex that share the
	// same infrastructure.
	ServerFailureDomainScope FailureDomainScope = "Server"
	// RackFailureDomainScope is a container within a datacenter that provides
	// communication, power, or cooling to a set of components.
	RackFailureDomainScope FailureDomainScope = "Rack"
	// RackGroupFailureDomainScope is a set of racks that may share common
	// communication, power, or cooling.
	RackGroupFailureDomainScope FailureDomainScope = "RackGroup"
	// RowFailureDomainScope is a set of adjacent racks or rackgroups that may
	// share common communication, power, or cooling.
	RowFailureDomainScope FailureDomainScope = "Row"
	// DatacenterFailureDomainScope is a facility that provides communication,
	// power, or cooling infrastructure to a co-located set of servers, networking
	// and storage.
	DatacenterFailureDomainScope FailureDomainScope = "Datacenter"
	// RegionFailureDomainScope is a set of resources that are required to be
	// either geographically or politically isolated from resources not in the
	// resources.
	RegionFailureDomainScope FailureDomainScope = "Region"
)

// RecoveryAccessScope is The enumeration literals shall represent the relative
// time required to make a replica available as a source.
type RecoveryAccessScope string

const (
	// OnlineActiveRecoveryAccessScope shall be instantaneous.
	OnlineActiveRecoveryAccessScope RecoveryAccessScope = "OnlineActive"
	// OnlinePassiveRecoveryAccessScope shall be consistent with switching access
	// to a different path the same front-end interconnect. A restore step shall
	// not be required.
	OnlinePassiveRecoveryAccessScope RecoveryAccessScope = "OnlinePassive"
	// NearlineRecoveryAccessScope shall be consistent with switching access to a
	// different path through a different front-end interconnection infrastructure.
	// Some inconsistency may occur. A restore step may be required before recovery
	// can commence.
	NearlineRecoveryAccessScope RecoveryAccessScope = "Nearline"
	// OfflineRecoveryAccessScope Access to a replica may take a significant amount
	// of time. No direct connection to the replica is assumed. Some inconsistency
	// loss may occur. A restore step is likely to be required.
	OfflineRecoveryAccessScope RecoveryAccessScope = "Offline"
)

// DataProtectionLoSCapabilities shall be met collectively by the communication
// path and the replica. There should be one instance associated to a class of
// service for each replica. Each replica independently should have a class of
// service that describes its characteristics.
type DataProtectionLoSCapabilities struct {
	common.Entity
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
	// SupportedLinesOfService shall contain known and supported
	// DataProtectionLinesOfService.
	supportedLinesOfService []string
	// SupportedLinesOfService@odata.count
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedMinLifetimes shall be an ISO 8601 duration that specifies the
	// minimum lifetime required for the replica.
	SupportedMinLifetimes []string
	// SupportedRecoveryGeographicObjectives shall specify a supported failure
	// domain.
	SupportedRecoveryGeographicObjectives []FailureDomainScope
	// SupportedRecoveryPointObjectiveTimes shall specify a supported ISO 8601 time
	// interval defining the maximum source information that may be lost on
	// failure. In the case that IsIsolated = false, failure of the domain is not a
	// consideration.
	SupportedRecoveryPointObjectiveTimes []string
	// SupportedRecoveryTimeObjectives shall specify an enumerated value that
	// indicates a supported expectation for the time required to access an
	// alternate replica. In the case that IsIsolated = false, failure of the
	// domain is not a consideration.
	SupportedRecoveryTimeObjectives []RecoveryAccessScope
	// SupportedReplicaTypes shall specify a supported replica type.
	SupportedReplicaTypes []ReplicaType
	// SupportsIsolated shall indicate that allocating a replica in a separate
	// fault domain is supported. The default value for this property is false.
	SupportsIsolated bool
	// SupportedReplicaOptionsCount is the number of supported replica options.
	SupportedReplicaOptionsCount int
	// supportedReplicaOptions are the URIs for SupportedReplicaOptions.
	supportedReplicaOptions []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a DataProtectionLoSCapabilities object from the raw JSON.
func (d *DataProtectionLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp DataProtectionLoSCapabilities
	type dLinks struct {
		SupportedReplicaOptions      common.Links `json:"SupportedReplicaOptions"`
		SupportedReplicaOptionsCount int          `json:"SupportedReplicaOptions@odata.count"`
	}
	var tmp struct {
		temp
		Links                   dLinks
		SupportedLinesOfService common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataProtectionLoSCapabilities(tmp.temp)

	// Extract the links to other entities for later
	d.supportedReplicaOptions = tmp.Links.SupportedReplicaOptions.ToStrings()
	d.SupportedReplicaOptionsCount = tmp.Links.SupportedReplicaOptionsCount
	d.supportedLinesOfService = tmp.SupportedLinesOfService.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	d.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataProtectionLoSCapabilities) Update() error {
	readWriteFields := []string{
		"Identifier",
		"SupportedLinesOfService",
		"SupportedLinesOfService@odata.count",
		"SupportedMinLifetimes",
		"SupportedRecoveryGeographicObjectives",
		"SupportedRecoveryPointObjectiveTimes",
		"SupportedRecoveryTimeObjectives",
		"SupportedReplicaTypes",
		"SupportsIsolated",
	}

	return d.UpdateFromRawData(d, d.rawData, readWriteFields)
}

// GetDataProtectionLoSCapabilities will get a DataProtectionLoSCapabilities instance from the service.
func GetDataProtectionLoSCapabilities(c common.Client, uri string) (*DataProtectionLoSCapabilities, error) {
	return common.GetObject[DataProtectionLoSCapabilities](c, uri)
}

// ListReferencedDataProtectionLoSCapabilities gets the collection of DataProtectionLoSCapabilities from
// a provided reference.
func ListReferencedDataProtectionLoSCapabilities(c common.Client, link string) ([]*DataProtectionLoSCapabilities, error) {
	return common.GetCollectionObjects[DataProtectionLoSCapabilities](c, link)
}

// SupportedReplicaOptions gets the SupportedReplicaOptions linked resources.
func (d *DataProtectionLoSCapabilities) SupportedReplicaOptions(client common.Client) ([]*ClassOfService, error) {
	return common.GetObjects[ClassOfService](client, d.supportedReplicaOptions)
}

// SupportedLinesOfService gets the supported lines of service.
func (d *DataProtectionLoSCapabilities) SupportedLinesOfService() ([]*DataProtectionLineOfService, error) {
	return common.GetObjects[DataProtectionLineOfService](
		d.GetClient(),
		d.supportedLinesOfService)
}
