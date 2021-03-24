//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// QuotaType shall indicate whether quotas are enabled and enforced by this file
// share. If QuotaType is present, a value of Soft shall mean that quotas are
// enabled but not enforced, and a value of Hard shall mean that quotas are
// enabled and enforced.
type QuotaType string

const (
	// SoftQuotaType shall indicate that quotas are enabled but not enforced.
	SoftQuotaType QuotaType = "Soft"
	// HardQuotaType shall indicate that quotas are enabled and enforced.
	HardQuotaType QuotaType = "Hard"
)

// FileShare is used to represent a shared set of files with a common directory
// structure.
type FileShare struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CASupported shall indicate that Continuous Availability is supported.
	// Client/Server mediated recovery from network and server failure with
	// application transparency. This property shall be NULL unless the
	// FileSharingProtocols property includes SMB. The default value for this
	// property is false.
	CASupported bool
	// DefaultAccessCapabilities shall be an array containing entries for the
	// default access capabilities for the file share. Each entry shall specify
	// a default access privilege. The types of default access can include Read,
	// Write, and/or Execute.
	DefaultAccessCapabilities []StorageAccessCapability
	// Description provides a description of this resource.
	Description string
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
	// FileShareQuotaType value of Soft shall specify that quotas are not
	// enforced, and a value of Hard shall specify that writes shall fail if the
	// space consumed would exceed the value of the FileShareTotalQuotaBytes
	// property.
	FileShareQuotaType QuotaType
	// FileShareRemainingQuotaBytes value of this property shall indicate the
	// remaining number of bytes that may be consumed by this file share.
	FileShareRemainingQuotaBytes int64
	// FileShareTotalQuotaBytes value of this property shall indicate the
	// maximum number of bytes that may be consumed by this file share.
	FileShareTotalQuotaBytes int64
	// FileSharingProtocols is This property shall be an array containing
	// entries for the file sharing protocols supported by this file share.
	// Each entry shall specify a file sharing protocol supported by the file
	// system.
	FileSharingProtocols []FileProtocol
	// LowSpaceWarningThresholdPercents is This property shall be an array
	// containing entries for the percentages of file share capacity at which
	// low space warning events are be issued. A LOW_SPACE_THRESHOLD_WARNING
	// event shall be triggered each time the remaining file share capacity
	// value becomes less than one of the values in the array. The following
	// shall be true: Across all CapacitySources entries, percent =
	// (SUM(AllocatedBytes) - SUM(ConsumedBytes))/SUM(AllocatedBytes)
	LowSpaceWarningThresholdPercents []int
	// RemainingCapacityPercent is If present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// RootAccess shall indicate whether Root
	// access is allowed by the file share. The default value for this
	// property is false.
	RootAccess bool
	// Status is This value of this property shall indicate the status of the
	// file share.
	Status common.Status
	// WritePolicy shall define how writes are
	// replicated to the shared source.
	WritePolicy ReplicaUpdateMode
	// classOfService shall be a link to the ClassOfService for this file share.
	classOfService string
	// fileSystem shall be a link to the file system containing the file share.
	fileSystem string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a FileShare object from the raw JSON.
func (fileshare *FileShare) UnmarshalJSON(b []byte) error {
	type temp FileShare
	type links struct {
		ClassOfService common.Link
		FileSystem     common.Link
	}
	var t struct {
		temp
		Links              links
		EthernetInterfaces common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*fileshare = FileShare(t.temp)
	fileshare.classOfService = string(t.Links.ClassOfService)
	fileshare.fileSystem = string(t.Links.FileSystem)
	fileshare.ethernetInterfaces = string(t.EthernetInterfaces)

	// This is a read/write object, so we need to save the raw object data for later
	fileshare.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fileshare *FileShare) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(FileShare)
	err := original.UnmarshalJSON(fileshare.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"CASupported",
		"FileShareQuotaType",
		"FileShareTotalQuotaBytes",
		"LowSpaceWarningThresholdPercents",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fileshare).Elem()

	return fileshare.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFileShare will get a FileShare instance from the service.
func GetFileShare(c common.Client, uri string) (*FileShare, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fileshare FileShare
	err = json.NewDecoder(resp.Body).Decode(&fileshare)
	if err != nil {
		return nil, err
	}

	fileshare.SetClient(c)
	return &fileshare, nil
}

// ListReferencedFileShares gets the collection of FileShare from a provided
// reference.
func ListReferencedFileShares(c common.Client, link string) ([]*FileShare, error) {
	var result []*FileShare
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, fileshareLink := range links.ItemLinks {
		fileshare, err := GetFileShare(c, fileshareLink)
		if err != nil {
			return result, err
		}
		result = append(result, fileshare)
	}

	return result, nil
}

// ClassOfService gets the file share's class of service.
func (fileshare *FileShare) ClassOfService() (*ClassOfService, error) {
	var result *ClassOfService
	if fileshare.classOfService == "" {
		return result, nil
	}
	return GetClassOfService(fileshare.Client, fileshare.classOfService)
}

// FileSystem gets the file share's associated file system.
func (fileshare *FileShare) FileSystem() (*FileSystem, error) {
	var result *FileSystem
	if fileshare.fileSystem == "" {
		return result, nil
	}
	return GetFileSystem(fileshare.Client, fileshare.fileSystem)
}

// EthernetInterfaces gets the EthernetInterfaces associated with this share.
func (fileshare *FileShare) EthernetInterfaces() ([]*redfish.EthernetInterface, error) {
	return redfish.ListReferencedEthernetInterfaces(fileshare.Client, fileshare.ethernetInterfaces)
}
