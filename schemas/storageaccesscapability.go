//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// StorageAccessCapability is StorageAccessCapability enumeration literals may
// be used to describe abilities to read or write storage.
type StorageAccessCapability string

const (
	// ReadStorageAccessCapability shall indicate that the storage may be read.
	ReadStorageAccessCapability StorageAccessCapability = "Read"
	// WriteStorageAccessCapability shall indicate that the storage may be written
	// multiple times.
	WriteStorageAccessCapability StorageAccessCapability = "Write"
	// WriteOnceStorageAccessCapability shall indicate that the storage may be
	// written only once.
	WriteOnceStorageAccessCapability StorageAccessCapability = "WriteOnce"
	// AppendStorageAccessCapability shall indicate that the storage may be written
	// only to append.
	AppendStorageAccessCapability StorageAccessCapability = "Append"
	// StreamingStorageAccessCapability shall indicate that the storage may be read
	// sequentially.
	StreamingStorageAccessCapability StorageAccessCapability = "Streaming"
	// ExecuteStorageAccessCapability shall indicate that Execute access is allowed
	// by the file share.
	ExecuteStorageAccessCapability StorageAccessCapability = "Execute" // from datastorageloscapabilities
)
