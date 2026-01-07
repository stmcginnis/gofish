//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.2 - #DataProtectionLineOfService.v1_3_0.DataProtectionLineOfService

package schemas

import (
	"encoding/json"
)

// DataProtectionLineOfService This service option describes a replica that
// protects data from loss. The requirements must be met collectively by the
// communication path and the replica.
type DataProtectionLineOfService struct {
	Entity
	// IsIsolated shall indicate that the replica is in a separate fault domain
	// from its source. The default value of this property is false.
	IsIsolated bool
	// MinLifetime shall be an ISO 8601 duration that specifies the minimum
	// required lifetime of the replica. Note: The maximum number of replicas can
	// be determined using this value together with the replicaSchedule.
	MinLifetime string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RecoveryGeographicObjective The value specifies the geographic scope of the
	// failure domain.
	RecoveryGeographicObjective FailureDomainScope
	// RecoveryPointObjectiveTime shall be an ISO 8601 duration that specifies the
	// maximum time over which source data may be lost on failure. In the case that
	// IsIsolated = false, failure of the domain is not a consideration.
	RecoveryPointObjectiveTime string
	// RecoveryTimeObjective shall be an enumeration that indicates the maximum
	// time required to access an alternate replica. In the case that IsIsolated =
	// false, failure of the domain is not a consideration.
	RecoveryTimeObjective RecoveryAccessScope
	// ReplicaAccessLocation shall be used if the data access location of the
	// replica is required to be at a specific location. Note 1: The location value
	// may be granular. Note 2: A value may be required for some regulatory
	// compliance.
	ReplicaAccessLocation Location
	// ReplicaClassOfService shall reference the class of service that defines the
	// required service levels of the replica.
	ReplicaClassOfService ClassOfService
	// ReplicaType shall conform to this value.
	ReplicaType ReplicaType
	// Schedule shall define the schedule.
	Schedule Schedule
	// createReplicasTarget is the URL to send CreateReplicas requests.
	createReplicasTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a DataProtectionLineOfService object from the raw JSON.
func (d *DataProtectionLineOfService) UnmarshalJSON(b []byte) error {
	type temp DataProtectionLineOfService
	type dActions struct {
		CreateReplicas ActionTarget `json:"#DataProtectionLineOfService.CreateReplicas"`
	}
	var tmp struct {
		temp
		Actions dActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DataProtectionLineOfService(tmp.temp)

	// Extract the links to other entities for later
	d.createReplicasTarget = tmp.Actions.CreateReplicas.Target

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DataProtectionLineOfService) Update() error {
	readWriteFields := []string{
		"IsIsolated",
		"MinLifetime",
		"RecoveryGeographicObjective",
		"RecoveryPointObjectiveTime",
		"RecoveryTimeObjective",
		"ReplicaClassOfService",
		"ReplicaType",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDataProtectionLineOfService will get a DataProtectionLineOfService instance from the service.
func GetDataProtectionLineOfService(c Client, uri string) (*DataProtectionLineOfService, error) {
	return GetObject[DataProtectionLineOfService](c, uri)
}

// ListReferencedDataProtectionLineOfServices gets the collection of DataProtectionLineOfService from
// a provided reference.
func ListReferencedDataProtectionLineOfServices(c Client, link string) ([]*DataProtectionLineOfService, error) {
	return GetCollectionObjects[DataProtectionLineOfService](c, link)
}

// This action shall create an on-demand replica that conforms to the bound
// DataProtectionLineOfService.
// replicaLineOfService - The value shall reference the data protection line of
// service this operation is bound to.
// replicaRequests - Each value shall reference a source resource and provide a
// name for the replica.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *DataProtectionLineOfService) CreateReplicas(replicaLineOfService string, replicaRequests []ReplicaRequest) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ReplicaLineOfService"] = replicaLineOfService
	payload["ReplicaRequests"] = replicaRequests
	resp, taskInfo, err := PostWithTask(d.client,
		d.createReplicasTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ReplicaRequest shall contain information about the ReplicaSource and the
// ReplicaName.
type ReplicaRequest struct {
	// ReplicaName shall be the names of the replica.
	//
	// Version added: v1.1.0
	ReplicaName string
	// ReplicaSource shall reference a resource to be replicated.
	//
	// Version added: v1.1.0
	replicaSource string
}

// UnmarshalJSON unmarshals a ReplicaRequest object from the raw JSON.
func (r *ReplicaRequest) UnmarshalJSON(b []byte) error {
	type temp ReplicaRequest
	var tmp struct {
		temp
		ReplicaSource Link `json:"ReplicaSource"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = ReplicaRequest(tmp.temp)

	// Extract the links to other entities for later
	r.replicaSource = tmp.ReplicaSource.String()

	return nil
}

// ReplicaSource gets the ReplicaSource linked resource.
func (r *ReplicaRequest) ReplicaSource(client Client) (*Entity, error) {
	if r.replicaSource == "" {
		return nil, nil
	}
	return GetObject[Entity](client, r.replicaSource)
}
