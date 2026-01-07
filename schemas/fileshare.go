//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.5 - #FileShare.v1_3_0.FileShare

package schemas

import (
	"encoding/json"
)

// QuotaType is The value shall indicate whether quotas are enabled and enforced
// by this file share. If QuotaType is present, a value of Soft shall mean that
// quotas are enabled but not enforced, and a value of Hard shall mean that
// quotas are enabled and enforced.
type QuotaType string

const (
	// SoftQuotaType shall indicate that quotas are enabled but not enforced.
	SoftQuotaType QuotaType = "Soft"
	// HardQuotaType shall indicate that quotas are enabled and enforced.
	HardQuotaType QuotaType = "Hard"
)

// FileShare shall be used to represent a shared set of files with a common
// directory structure.
type FileShare struct {
	Entity
	// CASupported shall indicate that Continuous Availability is supported.
	// Client/Server mediated recovery from network and server failure with
	// application transparency. This property shall be NULL unless the
	// FileSharingProtocols property includes SMB. The default value for this
	// property is false.
	CASupported bool
	// DefaultAccessCapabilities shall be an array containing entries for the
	// default access capabilities for the file share. Each entry shall specify a
	// default access privilege. The types of default access can include Read,
	// Write, and/or Execute.
	DefaultAccessCapabilities []StorageAccessCapability
	// EthernetInterfaces shall be a link to an EthernetInterfaceCollection with
	// members that provide access to the file share.
	ethernetInterfaces string
	// ExecuteSupport shall indicate whether Execute access is supported by the
	// file share. The default value for this property is false.
	ExecuteSupport bool
	// FileSharePath shall be a path (relative to the file system root) to the
	// exported file or directory on the file system where this file share is
	// hosted.
	FileSharePath string
	// FileShareQuotaType shall specify that quotas are not enforced, and a value
	// of Hard shall specify that writes shall fail if the space consumed would
	// exceed the value of the FileShareTotalQuotaBytes property.
	FileShareQuotaType QuotaType
	// FileShareRemainingQuotaBytes shall indicate the remaining number of bytes
	// that may be consumed by this file share.
	FileShareRemainingQuotaBytes *int `json:",omitempty"`
	// FileShareTotalQuotaBytes shall indicate the maximum number of bytes that may
	// be consumed by this file share.
	FileShareTotalQuotaBytes *int `json:",omitempty"`
	// FileSharingProtocols shall be an array containing entries for the file
	// sharing protocols supported by this file share. Each entry shall specify a
	// file sharing protocol supported by the file system.
	FileSharingProtocols []FileProtocol
	// LowSpaceWarningThresholdPercents shall be an array containing entries for
	// the percentages of file share capacity at which low space warning events are
	// be issued. A LOW_SPACE_THRESHOLD_WARNING event shall be triggered each time
	// the remaining file share capacity value becomes less than one of the values
	// in the array. The following shall be true: Across all CapacitySources
	// entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []*int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) -
	// SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100 represented as an integer
	// value.
	//
	// Version added: v1.1.0
	RemainingCapacityPercent *int `json:",omitempty"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on
	// the file share. This property shall be consistent with the state reflected
	// at the storage pool level.
	//
	// Version added: v1.3.0
	ReplicationEnabled bool
	// RootAccess shall indicate whether Root access is allowed by the file share.
	// The default value for this property is false.
	RootAccess bool
	// Status shall indicate the status of the file share.
	Status Status
	// WritePolicy shall define how writes are replicated to the shared source.
	WritePolicy ReplicaUpdateMode
	// classOfService is the URI for ClassOfService.
	classOfService string
	// fileSystem is the URI for FileSystem.
	fileSystem string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a FileShare object from the raw JSON.
func (f *FileShare) UnmarshalJSON(b []byte) error {
	type temp FileShare
	type fLinks struct {
		ClassOfService Link `json:"ClassOfService"`
		FileSystem     Link `json:"FileSystem"`
	}
	var tmp struct {
		temp
		Links              fLinks
		EthernetInterfaces Link `json:"EthernetInterfaces"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FileShare(tmp.temp)

	// Extract the links to other entities for later
	f.classOfService = tmp.Links.ClassOfService.String()
	f.fileSystem = tmp.Links.FileSystem.String()
	f.ethernetInterfaces = tmp.EthernetInterfaces.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *FileShare) Update() error {
	readWriteFields := []string{
		"CASupported",
		"FileShareQuotaType",
		"FileShareTotalQuotaBytes",
		"LowSpaceWarningThresholdPercents",
		"ReplicationEnabled",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetFileShare will get a FileShare instance from the service.
func GetFileShare(c Client, uri string) (*FileShare, error) {
	return GetObject[FileShare](c, uri)
}

// ListReferencedFileShares gets the collection of FileShare from
// a provided reference.
func ListReferencedFileShares(c Client, link string) ([]*FileShare, error) {
	return GetCollectionObjects[FileShare](c, link)
}

// ClassOfService gets the ClassOfService linked resource.
func (f *FileShare) ClassOfService() (*ClassOfService, error) {
	if f.classOfService == "" {
		return nil, nil
	}
	return GetObject[ClassOfService](f.client, f.classOfService)
}

// FileSystem gets the FileSystem linked resource.
func (f *FileShare) FileSystem() (*FileSystem, error) {
	if f.fileSystem == "" {
		return nil, nil
	}
	return GetObject[FileSystem](f.client, f.fileSystem)
}

// EthernetInterfaces gets the EthernetInterfaces collection.
func (f *FileShare) EthernetInterfaces() ([]*EthernetInterface, error) {
	if f.ethernetInterfaces == "" {
		return nil, nil
	}
	return GetCollectionObjects[EthernetInterface](f.client, f.ethernetInterfaces)
}
