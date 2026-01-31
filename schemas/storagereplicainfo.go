//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.5 - #StorageReplicaInfo.v1_4_0.StorageReplicaInfo

package schemas

import (
	"encoding/json"
)

// ConsistencyState is ConsistencyState enumeration literals may be used to
// describe the consistency type used by the source and its associated target
// group.
type ConsistencyState string

const (
	// ConsistentConsistencyState shall indicate that the source and target shall
	// be consistent.
	ConsistentConsistencyState ConsistencyState = "Consistent"
	// InconsistentConsistencyState shall indicate that the source and target are
	// not required to be consistent.
	InconsistentConsistencyState ConsistencyState = "Inconsistent"
)

// ConsistencyStatus is ConsistencyStatus enumeration literals may be used to
// indicate the current status of consistency. Consistency may have been
// disabled or may be experiencing an error condition.
type ConsistencyStatus string

const (
	// ConsistentConsistencyStatus shall indicate that the source and target are
	// consistent.
	ConsistentConsistencyStatus ConsistencyStatus = "Consistent"
	// InProgressConsistencyStatus shall indicate that the source and target are
	// becoming consistent.
	InProgressConsistencyStatus ConsistencyStatus = "InProgress"
	// DisabledConsistencyStatus shall indicate that the source and target have
	// consistency disabled.
	DisabledConsistencyStatus ConsistencyStatus = "Disabled"
	// InErrorConsistencyStatus shall indicate that the source and target are not
	// consistent.
	InErrorConsistencyStatus ConsistencyStatus = "InError"
)

// ReplicationConsistencyType enumeration literals indicate the
// consistency type used by the source and its associated target group.
type ReplicationConsistencyType string

const (
	// SequentiallyConsistentReplicationConsistencyType shall indicate that the source and
	// target shall be sequentially consistent.
	SequentiallyConsistentReplicationConsistencyType ReplicationConsistencyType = "SequentiallyConsistent"
)

// ReplicaFaultDomain is The enumeration literals may be used to specify the
// fault domain for the replication relationship.
type ReplicaFaultDomain string

const (
	// LocalReplicaFaultDomain shall indicate that the source and target replicas
	// are contained within a single fault domain.
	LocalReplicaFaultDomain ReplicaFaultDomain = "Local"
	// RemoteReplicaFaultDomain shall indicate that the source and target replicas
	// are in separate fault domains.
	RemoteReplicaFaultDomain ReplicaFaultDomain = "Remote"
)

// ReplicaPriority is The enumeration literals of the ReplicaPriority
// enumeration may be used to specify the priority of background copy engine I/O
// relative to host I/O operations during a sequential background copy
// operation.
type ReplicaPriority string

const (
	// LowReplicaPriority shall have a lower priority than host I/O.
	LowReplicaPriority ReplicaPriority = "Low"
	// SameReplicaPriority shall have the same priority as host I/O.
	SameReplicaPriority ReplicaPriority = "Same"
	// HighReplicaPriority shall have a higher priority than host I/O.
	HighReplicaPriority ReplicaPriority = "High"
	// UrgentReplicaPriority shall be performed as soon as possible.
	UrgentReplicaPriority ReplicaPriority = "Urgent"
)

// ReplicaProgressStatus is ReplicaProgressStatus enumeration literals may be
// used to describe the status of the session with respect to Replication
// activity.
type ReplicaProgressStatus string

const (
	// CompletedReplicaProgressStatus shall indicate that the request is completed.
	// Data flow is idle.
	CompletedReplicaProgressStatus ReplicaProgressStatus = "Completed"
	// DormantReplicaProgressStatus shall indicate that the data flow is inactive,
	// suspended or quiesced.
	DormantReplicaProgressStatus ReplicaProgressStatus = "Dormant"
	// InitializingReplicaProgressStatus shall indicate that replication is in the
	// process of establishing source/replica relationship and the data flow has
	// not started.
	InitializingReplicaProgressStatus ReplicaProgressStatus = "Initializing"
	// PreparingReplicaProgressStatus shall indicate that replication has
	// preparation in progress.
	PreparingReplicaProgressStatus ReplicaProgressStatus = "Preparing"
	// SynchronizingReplicaProgressStatus shall indicate that replication has
	// synchronization in progress.
	SynchronizingReplicaProgressStatus ReplicaProgressStatus = "Synchronizing"
	// ResyncingReplicaProgressStatus shall indicate that replication has
	// resynchronization in progress.
	ResyncingReplicaProgressStatus ReplicaProgressStatus = "Resyncing"
	// RestoringReplicaProgressStatus shall indicate that replication has a restore
	// in progress.
	RestoringReplicaProgressStatus ReplicaProgressStatus = "Restoring"
	// FracturingReplicaProgressStatus shall indicate that replication has a
	// fracture in progress.
	FracturingReplicaProgressStatus ReplicaProgressStatus = "Fracturing"
	// SplittingReplicaProgressStatus shall indicate that replication has a split
	// in progress.
	SplittingReplicaProgressStatus ReplicaProgressStatus = "Splitting"
	// FailingOverReplicaProgressStatus shall indicate that replication is in the
	// process of switching source and target.
	FailingOverReplicaProgressStatus ReplicaProgressStatus = "FailingOver"
	// FailingBackReplicaProgressStatus shall indicate that replication is undoing
	// the result of failover.
	FailingBackReplicaProgressStatus ReplicaProgressStatus = "FailingBack"
	// DetachingReplicaProgressStatus shall indicate that replication has a detach
	// in progress.
	DetachingReplicaProgressStatus ReplicaProgressStatus = "Detaching"
	// AbortingReplicaProgressStatus shall indicate that replication has an abort
	// in progress.
	AbortingReplicaProgressStatus ReplicaProgressStatus = "Aborting"
	// MixedReplicaProgressStatus shall indicate that replication status is mixed
	// across element pairs in a replication group. Generally, the individual
	// statuses need to be examined.
	MixedReplicaProgressStatus ReplicaProgressStatus = "Mixed"
	// SuspendingReplicaProgressStatus shall indicate that replication has a copy
	// operation in the process of being suspended.
	SuspendingReplicaProgressStatus ReplicaProgressStatus = "Suspending"
	// RequiresFractureReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be fractured before further copy operations can be issued.
	RequiresFractureReplicaProgressStatus ReplicaProgressStatus = "RequiresFracture"
	// RequiresResyncReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be resynced before further copy operations can be issued.
	RequiresResyncReplicaProgressStatus ReplicaProgressStatus = "RequiresResync"
	// RequiresActivateReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be activated before further copy operations can be issued.
	RequiresActivateReplicaProgressStatus ReplicaProgressStatus = "RequiresActivate"
	// PendingReplicaProgressStatus shall indicate that the flow of data has
	// stopped momentarily due to limited bandwidth or a busy system.
	PendingReplicaProgressStatus ReplicaProgressStatus = "Pending"
	// RequiresDetachReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be detached before further copy operations can be issued.
	RequiresDetachReplicaProgressStatus ReplicaProgressStatus = "RequiresDetach"
	// TerminatingReplicaProgressStatus shall indicate that the replication
	// relationship is in the process of terminating.
	TerminatingReplicaProgressStatus ReplicaProgressStatus = "Terminating"
	// RequiresSplitReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be split before further copy operations can be issued.
	RequiresSplitReplicaProgressStatus ReplicaProgressStatus = "RequiresSplit"
	// RequiresResumeReplicaProgressStatus shall indicate that the requested
	// operation has completed, however, the synchronization relationship needs to
	// be resumed before further copy operations can be issued.
	RequiresResumeReplicaProgressStatus ReplicaProgressStatus = "RequiresResume"
)

// ReplicaReadOnlyAccess is The enumeration literals may be used to specify
// whether the source, the target, or both elements are read-only to the host.
type ReplicaReadOnlyAccess string

const (
	// SourceElementReplicaReadOnlyAccess shall be read-only to the host.
	SourceElementReplicaReadOnlyAccess ReplicaReadOnlyAccess = "SourceElement"
	// ReplicaElementReplicaReadOnlyAccess shall be read-only to the host.
	ReplicaElementReplicaReadOnlyAccess ReplicaReadOnlyAccess = "ReplicaElement"
	// BothReplicaReadOnlyAccess shall be read only to the host.
	BothReplicaReadOnlyAccess ReplicaReadOnlyAccess = "Both"
)

// ReplicaRecoveryMode is The enumeration literals may be used to specify
// whether the copy operation continues after a broken link is restored.
type ReplicaRecoveryMode string

const (
	// AutomaticReplicaRecoveryMode shall resume automatically.
	AutomaticReplicaRecoveryMode ReplicaRecoveryMode = "Automatic"
	// ManualReplicaRecoveryMode shall be set to Suspended after the link is
	// restored. It is required to issue the Resume operation to continue.
	ManualReplicaRecoveryMode ReplicaRecoveryMode = "Manual"
)

// ReplicaRole is The enumeration literals may be used to specify whether the
// resource is a source of replication or the target of replication.
type ReplicaRole string

const (
	// SourceReplicaRole shall indicate a source element.
	SourceReplicaRole ReplicaRole = "Source"
	// TargetReplicaRole shall indicate target element.
	TargetReplicaRole ReplicaRole = "Target"
)

// ReplicaState is ReplicaState enumeration literals may be used to describe the
// state of the relationship with respect to Replication activity.
type ReplicaState string

const (
	// InitializedReplicaState shall indicate that the link to enable replication
	// is established and source/replica elements are associated, but the data flow
	// has not started.
	InitializedReplicaState ReplicaState = "Initialized"
	// UnsynchronizedReplicaState shall indicate that not all the source element
	// data has been copied to the target element.
	UnsynchronizedReplicaState ReplicaState = "Unsynchronized"
	// SynchronizedReplicaState shall indicate that for Mirror, Snapshot, or Clone
	// replication, the target represents a copy of the source.
	SynchronizedReplicaState ReplicaState = "Synchronized"
	// BrokenReplicaState shall indicate that the relationship is non-functional
	// due to errors in the source, the target, the path between the two or space
	// constraints.
	BrokenReplicaState ReplicaState = "Broken"
	// FracturedReplicaState shall indicate that the Target is split from the
	// source. The target may not be consistent.
	FracturedReplicaState ReplicaState = "Fractured"
	// SplitReplicaState shall indicate that the target element was gracefully (or
	// systematically) split from its source element -- consistency shall be
	// guaranteed.
	SplitReplicaState ReplicaState = "Split"
	// InactiveReplicaState shall indicate that data flow has stopped, writes to
	// source element shall not be sent to target element.
	InactiveReplicaState ReplicaState = "Inactive"
	// SuspendedReplicaState shall indicate that the data flow between the source
	// and target elements has stopped. Writes to source element shall be held
	// until the relationship is Resumed.
	SuspendedReplicaState ReplicaState = "Suspended"
	// FailedoverReplicaState shall indicate that the reads and writes are sent to
	// the target element. The source element may not be reachable.
	FailedoverReplicaState ReplicaState = "Failedover"
	// PreparedReplicaState shall indicate that initialization is completed,
	// however, the data flow has not started.
	PreparedReplicaState ReplicaState = "Prepared"
	// AbortedReplicaState shall indicate that the copy operation is aborted with
	// the Abort operation. The Resync Replica operation can be used to restart the
	// copy operation.
	AbortedReplicaState ReplicaState = "Aborted"
	// SkewedReplicaState shall indicate that the target has been modified and is
	// no longer synchronized with the source element or the point-in-time view.
	SkewedReplicaState ReplicaState = "Skewed"
	// MixedReplicaState shall indicate the ReplicaState of GroupSynchronized. The
	// value indicates the StorageSynchronized relationships of the elements in the
	// group have different ReplicaState values.
	MixedReplicaState ReplicaState = "Mixed"
	// PartitionedReplicaState shall indicate that the state of replication
	// relationship can not be determined, for example, due to a connection
	// problem.
	PartitionedReplicaState ReplicaState = "Partitioned"
	// InvalidReplicaState shall indicate that the storage server is unable to
	// determine the state of the replication relationship, for example, after the
	// connection is restored; however, either source or target elements have an
	// unknown status.
	InvalidReplicaState ReplicaState = "Invalid"
	// RestoredReplicaState shall indicate that the source element was restored
	// from the target element.
	RestoredReplicaState ReplicaState = "Restored"
)

// ReplicaType is The enumeration literals may be used to specify the intended
// outcome of the replication.
type ReplicaType string

const (
	// MirrorReplicaType shall indicate that replication shall create and maintain
	// a copy of the source.
	MirrorReplicaType ReplicaType = "Mirror"
	// SnapshotReplicaType shall indicate that replication shall create a point in
	// time, virtual copy of the source.
	SnapshotReplicaType ReplicaType = "Snapshot"
	// CloneReplicaType shall indicate that replication shall create a point in
	// time, full copy the source.
	CloneReplicaType ReplicaType = "Clone"
	// TokenizedCloneReplicaType shall indicate that replication shall create a
	// token based clone.
	TokenizedCloneReplicaType ReplicaType = "TokenizedClone"
)

// ReplicaUpdateMode is The enumeration literals may be used to specify whether
// the target elements will be updated synchronously or asynchronously.
type ReplicaUpdateMode string

const (
	// ActiveReplicaUpdateMode shall indicate Active-Active (i.e. bidirectional)
	// synchronous updates.
	ActiveReplicaUpdateMode ReplicaUpdateMode = "Active"
	// SynchronousReplicaUpdateMode shall indicate Synchronous updates.
	SynchronousReplicaUpdateMode ReplicaUpdateMode = "Synchronous"
	// AsynchronousReplicaUpdateMode shall indicate Asynchronous updates.
	AsynchronousReplicaUpdateMode ReplicaUpdateMode = "Asynchronous"
	// AdaptiveReplicaUpdateMode shall indicate that an implementation may switch
	// between synchronous and asynchronous modes.
	AdaptiveReplicaUpdateMode ReplicaUpdateMode = "Adaptive"
)

// UndiscoveredElement is The enumeration literals may be used to specify
// whether the source, the target, or both elements involved in a copy operation
// are undiscovered. An element shall be considered undiscovered if its object
// model is not known to the service performing the copy operation.
type UndiscoveredElement string

const (
	// SourceElementUndiscoveredElement shall indicate that the source element is
	// undiscovered.
	SourceElementUndiscoveredElement UndiscoveredElement = "SourceElement"
	// ReplicaElementUndiscoveredElement shall indicate that the replica element is
	// undiscovered.
	ReplicaElementUndiscoveredElement UndiscoveredElement = "ReplicaElement"
)

// StorageReplicaInfo shall define the characteristics of a replica.
type StorageReplicaInfo struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetStorageReplicaInfo will get a StorageReplicaInfo instance from the service.
func GetStorageReplicaInfo(c Client, uri string) (*StorageReplicaInfo, error) {
	return GetObject[StorageReplicaInfo](c, uri)
}

// ListReferencedStorageReplicaInfos gets the collection of StorageReplicaInfo from
// a provided reference.
func ListReferencedStorageReplicaInfos(c Client, link string) ([]*StorageReplicaInfo, error) {
	return GetCollectionObjects[StorageReplicaInfo](c, link)
}

// ReplicaInfo shall define the characteristics of a replica.
type ReplicaInfo struct {
	// ConsistencyEnabled shall be enabled across the source and its associated
	// target replica(s). The default value for this property is false.
	ConsistencyEnabled bool
	// ConsistencyState shall indicate the current state of consistency.
	ConsistencyState ConsistencyState
	// ConsistencyStatus shall specify the current status of consistency.
	// Consistency may have been disabled or is experiencing an error condition.
	ConsistencyStatus ConsistencyStatus
	// ConsistencyType shall indicate the consistency type used by the source and
	// its associated target group.
	ConsistencyType ReplicationConsistencyType
	// DataProtectionLineOfService shall be a pointer to the data protection line
	// of service that describes this replica.
	//
	// Version added: v1.1.0
	dataProtectionLineOfService string
	// FailedCopyStopsHostIO shall stop receiving data to the source element if
	// copying to a remote element fails. The default value for this property is
	// false.
	FailedCopyStopsHostIO bool
	// PercentSynced shall be an average of the PercentSynced across all members of
	// the group.
	PercentSynced *int `json:",omitempty"`
	// RemoteSourceReplica shall describe the fault domain (local or remote) of the
	// replica relationship.
	//
	// Version added: v1.4.0
	RemoteSourceReplica string
	// Replica shall reference the resource that is the source of this replica.
	replica string
	// ReplicaFaultDomain shall describe the fault domain (local or remote) of the
	// replica relationship.
	//
	// Version added: v1.3.0
	ReplicaFaultDomain ReplicaFaultDomain
	// ReplicaPriority shall specify the priority of background copy engine I/O to
	// be managed relative to host I/O operations during a sequential background
	// copy operation.
	ReplicaPriority ReplicaPriority
	// ReplicaProgressStatus shall specify the status of the session with respect
	// to Replication activity.
	ReplicaProgressStatus ReplicaProgressStatus
	// ReplicaReadOnlyAccess shall specify whether the source, the target, or both
	// elements are read only to the host.
	ReplicaReadOnlyAccess ReplicaReadOnlyAccess
	// ReplicaRecoveryMode shall specify whether the copy operation continues after
	// a broken link is restored.
	ReplicaRecoveryMode ReplicaRecoveryMode
	// ReplicaRole shall represent the source or target role of this replica as
	// known to the containing resource.
	//
	// Deprecated
	// ReplicaInfo is only used within a replica target. The Replica property here
	// addresses the source resource. A TargetReplicas property in each source
	// resource describes the replica targets of a source.
	ReplicaRole ReplicaRole
	// ReplicaSkewBytes shall be switched to synchronous.
	ReplicaSkewBytes *int `json:",omitempty"`
	// ReplicaState shall specify the state of the relationship with respect to
	// Replication activity.
	ReplicaState ReplicaState
	// ReplicaType shall describe the intended outcome of the replication.
	ReplicaType ReplicaType
	// ReplicaUpdateMode shall specify whether the target elements will be updated
	// synchronously or asynchronously.
	ReplicaUpdateMode ReplicaUpdateMode
	// RequestedReplicaState shall be represented by ReplicaState. When
	// RequestedState reaches the requested state, this property shall be null.
	RequestedReplicaState ReplicaState
	// SourceReplica shall contain the URI to the source replica when located on a
	// different Swordfish service instance.
	//
	// Version added: v1.2.0
	sourceReplica string
	// SyncMaintained shall be maintained. The default value for this property is
	// false.
	SyncMaintained bool
	// UndiscoveredElement shall specify whether the source, the target, or both
	// elements involved in a copy operation are undiscovered. An element is
	// considered undiscovered if its object model is not known to the service
	// performing the copy operation.
	UndiscoveredElement UndiscoveredElement
	// WhenActivated shall be an ISO 8601 conformant time of day that specifies
	// when the point-in-time copy was taken or when the replication relationship
	// is activated, reactivated, resumed or re-established. This property shall be
	// null if the implementation is not capable of providing this information.
	WhenActivated string
	// WhenDeactivated shall be an ISO 8601 conformant time of day that specifies
	// when the replication relationship is deactivated. Do not instantiate this
	// property if implementation is not capable of providing this information.
	WhenDeactivated string
	// WhenEstablished shall be an ISO 8601 conformant time of day that specifies
	// when the replication relationship is established. Do not instantiate this
	// property if implementation is not capable of providing this information.
	WhenEstablished string
	// WhenSuspended shall be an ISO 8601 conformant time of day that specifies
	// when the replication relationship is suspended. Do not instantiate this
	// property if implementation is not capable of providing this information.
	WhenSuspended string
	// WhenSynced shall be an ISO 8601 conformant time of day that specifies when
	// the elements were synchronized.
	WhenSynced string
	// WhenSynchronized shall be an ISO 8601 conformant time of day that specifies
	// when the replication relationship is synchronized. Do not instantiate this
	// property if implementation is not capable of providing this information.
	WhenSynchronized string
}

// UnmarshalJSON unmarshals a ReplicaInfo object from the raw JSON.
func (r *ReplicaInfo) UnmarshalJSON(b []byte) error {
	type temp ReplicaInfo
	var tmp struct {
		temp
		DataProtectionLineOfService Link `json:"DataProtectionLineOfService"`
		Replica                     Link `json:"Replica"`
		SourceReplica               Link `json:"SourceReplica"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = ReplicaInfo(tmp.temp)

	// Extract the links to other entities for later
	r.dataProtectionLineOfService = tmp.DataProtectionLineOfService.String()
	r.replica = tmp.Replica.String()
	r.sourceReplica = tmp.SourceReplica.String()

	return nil
}

// DataProtectionLineOfService gets the DataProtectionLineOfService linked resource.
func (r *ReplicaInfo) DataProtectionLineOfService(client Client) (*DataProtectionLineOfService, error) {
	if r.dataProtectionLineOfService == "" {
		return nil, nil
	}
	return GetObject[DataProtectionLineOfService](client, r.dataProtectionLineOfService)
}

// Replica gets the Replica linked resource.
func (r *ReplicaInfo) Replica(client Client) (*Entity, error) {
	if r.replica == "" {
		return nil, nil
	}
	return GetObject[Entity](client, r.replica)
}

// SourceReplica gets the SourceReplica linked resource.
func (r *ReplicaInfo) SourceReplica(client Client) (*Entity, error) {
	if r.sourceReplica == "" {
		return nil, nil
	}
	return GetObject[Entity](client, r.sourceReplica)
}
