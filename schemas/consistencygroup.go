//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.4a - #ConsistencyGroup.v1_1_1.ConsistencyGroup

package schemas

import (
	"encoding/json"
)

type ApplicationConsistencyMethod string

const (
	// HotStandbyApplicationConsistencyMethod Supports consistency method commonly
	// orchestrated using application-specific code.
	HotStandbyApplicationConsistencyMethod ApplicationConsistencyMethod = "HotStandby"
	// VASAApplicationConsistencyMethod Supports VMware consistency requirements,
	// such as for VASA and VVOLs.
	VASAApplicationConsistencyMethod ApplicationConsistencyMethod = "VASA"
	// VDIApplicationConsistencyMethod Supports Microsoft virtual backup device
	// interface (VDI).
	VDIApplicationConsistencyMethod ApplicationConsistencyMethod = "VDI"
	// VSSApplicationConsistencyMethod Supports Microsoft VSS.
	VSSApplicationConsistencyMethod ApplicationConsistencyMethod = "VSS"
	// OtherApplicationConsistencyMethod Supports consistency method orchestrated
	// using vendor-specific code.
	OtherApplicationConsistencyMethod ApplicationConsistencyMethod = "Other"
)

type ConsistencyType string

const (
	// CrashConsistentConsistencyType Requested operations are either triggered or
	// instituted without regard to pending IO.
	CrashConsistentConsistencyType ConsistencyType = "CrashConsistent"
	// ApplicationConsistentConsistencyType Orchestration exists to either flush or
	// halt pending IO to ensure operations occur in a transactionally consistent
	// manner.
	ApplicationConsistentConsistencyType ConsistencyType = "ApplicationConsistent"
)

// ConsistencyGroup is a collection of volumes grouped together to ensure write
// order consistency across all those volumes. A management operation on a
// consistency group, such as configuring replication properties, applies to all
// the volumes within the consistency group.
type ConsistencyGroup struct {
	Entity
	// ConsistencyMethod shall set the consistency method used by this group.
	ConsistencyMethod ApplicationConsistencyMethod
	// ConsistencyType shall set the consistency type used by this group.
	ConsistencyType ConsistencyType
	// IsConsistent shall be set to true when the consistency group is in a
	// consistent state.
	IsConsistent bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas
	// that are sourced by this replica. Remote indicates that the replica is
	// managed by a separate Swordfish service instance.
	//
	// Version added: v1.1.0
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replication relationship between this storage
	// group and a corresponding source storage group.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this
	// replica.
	ReplicaTargets []Entity
	// ReplicaTargetsCount
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status shall contain the status of the ConsistencyGroup.
	Status Status
	// Volumes is an array of references to volumes managed by this storage group.
	volumes []string
	// VolumesCount
	VolumesCount int `json:"Volumes@odata.count"`
	// assignReplicaTargetTarget is the URL to send AssignReplicaTarget requests.
	assignReplicaTargetTarget string
	// createReplicaTargetTarget is the URL to send CreateReplicaTarget requests.
	createReplicaTargetTarget string
	// removeReplicaRelationshipTarget is the URL to send RemoveReplicaRelationship requests.
	removeReplicaRelationshipTarget string
	// resumeReplicationTarget is the URL to send ResumeReplication requests.
	resumeReplicationTarget string
	// reverseReplicationRelationshipTarget is the URL to send ReverseReplicationRelationship requests.
	reverseReplicationRelationshipTarget string
	// splitReplicationTarget is the URL to send SplitReplication requests.
	splitReplicationTarget string
	// suspendReplicationTarget is the URL to send SuspendReplication requests.
	suspendReplicationTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ConsistencyGroup object from the raw JSON.
func (c *ConsistencyGroup) UnmarshalJSON(b []byte) error {
	type temp ConsistencyGroup
	type cActions struct {
		AssignReplicaTarget            ActionTarget `json:"#ConsistencyGroup.AssignReplicaTarget"`
		CreateReplicaTarget            ActionTarget `json:"#ConsistencyGroup.CreateReplicaTarget"`
		RemoveReplicaRelationship      ActionTarget `json:"#ConsistencyGroup.RemoveReplicaRelationship"`
		ResumeReplication              ActionTarget `json:"#ConsistencyGroup.ResumeReplication"`
		ReverseReplicationRelationship ActionTarget `json:"#ConsistencyGroup.ReverseReplicationRelationship"`
		SplitReplication               ActionTarget `json:"#ConsistencyGroup.SplitReplication"`
		SuspendReplication             ActionTarget `json:"#ConsistencyGroup.SuspendReplication"`
	}
	var tmp struct {
		temp
		Actions cActions
		Volumes Links `json:"Volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ConsistencyGroup(tmp.temp)

	// Extract the links to other entities for later
	c.assignReplicaTargetTarget = tmp.Actions.AssignReplicaTarget.Target
	c.createReplicaTargetTarget = tmp.Actions.CreateReplicaTarget.Target
	c.removeReplicaRelationshipTarget = tmp.Actions.RemoveReplicaRelationship.Target
	c.resumeReplicationTarget = tmp.Actions.ResumeReplication.Target
	c.reverseReplicationRelationshipTarget = tmp.Actions.ReverseReplicationRelationship.Target
	c.splitReplicationTarget = tmp.Actions.SplitReplication.Target
	c.suspendReplicationTarget = tmp.Actions.SuspendReplication.Target
	c.volumes = tmp.Volumes.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ConsistencyGroup) Update() error {
	readWriteFields := []string{
		"ConsistencyMethod",
		"ConsistencyType",
		"Volumes",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetConsistencyGroup will get a ConsistencyGroup instance from the service.
func GetConsistencyGroup(c Client, uri string) (*ConsistencyGroup, error) {
	return GetObject[ConsistencyGroup](c, uri)
}

// ListReferencedConsistencyGroups gets the collection of ConsistencyGroup from
// a provided reference.
func ListReferencedConsistencyGroups(c Client, link string) ([]*ConsistencyGroup, error) {
	return GetCollectionObjects[ConsistencyGroup](c, link)
}

// This action shall be used to establish a replication relationship by
// assigning an existing consistency group to serve as a target replica for an
// existing source consistency group.
// replicaType - This parameter shall contain the type of replica relationship
// to be created.
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) AssignReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.assignReplicaTargetTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ConsistencyGroupCreateReplicaTargetParameters holds the parameters for the CreateReplicaTarget action.
type ConsistencyGroupCreateReplicaTargetParameters struct {
	// ConsistencyGroupName shall contain the Name for the target consistency
	// group.
	ConsistencyGroupName string `json:"ConsistencyGroupName,omitempty"`
	// ReplicaType shall contain the type of replica relationship to be created.
	ReplicaType ReplicaType `json:"ReplicaType,omitempty"`
	// ReplicaUpdateMode shall specify the replica update mode.
	ReplicaUpdateMode ReplicaUpdateMode `json:"ReplicaUpdateMode,omitempty"`
	// TargetStoragePool shall contain the Uri to the existing StoragePool in which
	// to create the target consistency group.
	TargetStoragePool string `json:"TargetStoragePool,omitempty"`
}

// This action shall be used to create a new consistency group resource to
// provide expanded data protection through a replica relationship with the
// specified source consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) CreateReplicaTarget(params *ConsistencyGroupCreateReplicaTargetParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(c.client,
		c.createReplicaTargetTarget, params, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall be used to disable data synchronization between a source
// and target consistency group, remove the replication relationship, and
// optionally delete the target consistency group.
// deleteTargetConsistencyGroup - This parameter shall indicate whether or not
// to delete the target consistency group as part of the operation. If not
// specified, the system should use its default behavior.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) RemoveReplicaRelationship(deleteTargetConsistencyGroup bool, targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["DeleteTargetConsistencyGroup"] = deleteTargetConsistencyGroup
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.removeReplicaRelationshipTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall be used to resume the active data synchronization between
// a source and target consistency group, without otherwise altering the
// replication relationship.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) ResumeReplication(targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.resumeReplicationTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall be used to reverse the replication relationship between a
// source and target consistency group.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) ReverseReplicationRelationship(targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.reverseReplicationRelationshipTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall be used to split the replication relationship and suspend
// data synchronization between a source and target consistency group.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) SplitReplication(targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.splitReplicationTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall be used to suspend active data synchronization between a
// source and target consistency group, without otherwise altering the
// replication relationship.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ConsistencyGroup) SuspendReplication(targetConsistencyGroup string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	resp, taskInfo, err := PostWithTask(c.client,
		c.suspendReplicationTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Volumes gets the Volumes linked resources.
func (c *ConsistencyGroup) Volumes() ([]*Volume, error) {
	return GetObjects[Volume](c.client, c.volumes)
}
