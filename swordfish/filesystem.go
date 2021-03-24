//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CharacterCodeSet shall indicate the character code standards supported by the
// file system.
type CharacterCodeSet string

const (
	// ASCIICharacterCodeSet shall indicate that the ASCII character encoding
	// is supported by the file system.
	ASCIICharacterCodeSet CharacterCodeSet = "ASCII"
	// UnicodeCharacterCodeSet shall indicate that Unicode character encoding
	// is supported by the file system.
	UnicodeCharacterCodeSet CharacterCodeSet = "Unicode"
	// ISO2022CharacterCodeSet shall indicate that ISO-2022 character
	// encoding is supported by the file system.
	ISO2022CharacterCodeSet CharacterCodeSet = "ISO2022"
	// ISO88591CharacterCodeSet shall indicate that ISO-8859-1 character
	// encoding is supported by the file system.
	ISO88591CharacterCodeSet CharacterCodeSet = "ISO8859_1"
	// ExtendedUNIXCodeCharacterCodeSet shall indicate that Extended Unix
	// Code character encoding is supported by the file system.
	ExtendedUNIXCodeCharacterCodeSet CharacterCodeSet = "ExtendedUNIXCode"
	// UTF8CharacterCodeSet shall indicate that the UTF-8 character encoding
	// is supported by the file system.
	UTF8CharacterCodeSet CharacterCodeSet = "UTF_8"
	// UTF16CharacterCodeSet shall indicate that the UTF-16 character
	// encoding is supported by the file system.
	UTF16CharacterCodeSet CharacterCodeSet = "UTF_16"
	// UCS2CharacterCodeSet shall indicate that the UCS-2 character encoding
	// is supported by the file system.
	UCS2CharacterCodeSet CharacterCodeSet = "UCS_2"
)

// FileProtocol shall indicate the file sharing protocols supported by the file
// system. At least one value shall be present.
type FileProtocol string

const (
	// NFSv3FileProtocol shall indicate that NFSv3, as defined in RFC 1813,
	// is supported by the file system.
	NFSv3FileProtocol FileProtocol = "NFSv3"
	// NFSv40FileProtocol shall indicate that NFSv4, as defined in RFC 7530,
	// is supported by the file system.
	NFSv40FileProtocol FileProtocol = "NFSv4_0"
	// NFSv41FileProtocol shall indicate that NFSv4.1, as defined in RFC
	// 5661, is supported by the file system.
	NFSv41FileProtocol FileProtocol = "NFSv4_1"
	// SMBv20FileProtocol shall indicate that Server Message Block version
	// 2.0 is supported by the file system.
	SMBv20FileProtocol FileProtocol = "SMBv2_0"
	// SMBv21FileProtocol shall indicate that Server Message Block version
	// 2.1 is supported by the file system.
	SMBv21FileProtocol FileProtocol = "SMBv2_1"
	// SMBv30FileProtocol shall indicate that Server Message Block version
	// 3.0 is supported by the file system.
	SMBv30FileProtocol FileProtocol = "SMBv3_0"
	// SMBv302FileProtocol shall indicate that Server Message Block version
	// 3.0.2 is supported by the file system.
	SMBv302FileProtocol FileProtocol = "SMBv3_0_2"
	// SMBv311FileProtocol shall indicate that Server Message Block version
	// 3.1.1 is supported by the file system.
	SMBv311FileProtocol FileProtocol = "SMBv3_1_1"
)

// FileSystemPersistenceType shall indicate the persistence characteristics of
// the file system.
type FileSystemPersistenceType string

const (
	// PersistentFileSystemPersistenceType shall indicate that the file
	// system is persistent, and shall be preserved through an orderly
	// shutdown.
	PersistentFileSystemPersistenceType FileSystemPersistenceType = "Persistent"
	// TemporaryFileSystemPersistenceType shall indicate that the file system
	// is non-persistent and may not survive a shutdown.
	TemporaryFileSystemPersistenceType FileSystemPersistenceType = "Temporary"
	// OtherFileSystemPersistenceType shall indicate that the persistence
	// type is known, but not defined by this standard. Use of this value is
	// not recommended.
	OtherFileSystemPersistenceType FileSystemPersistenceType = "Other"
)

// FileSystem is used to represent an instance of a hierarchical namespace of
// files.
type FileSystem struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessCapabilities shall be an array containing entries for the supported
	// IO access capabilities. Each entry shall specify a current storage access
	// capability.
	AccessCapabilities []StorageAccessCapability
	// BlockSizeBytes shall be the block size of the file system in bytes.
	BlockSizeBytes int64
	// Capacity shall be the capacity allocated to the file system in bytes.
	Capacity Capacity
	// CapacitySources shall be an array containing entries for all the capacity
	// sources for the file system. Each entry shall provide capacity allocation
	// information from a named resource.
	CapacitySources []CapacitySource
	// CapacitySourcesCount is the number of CapacitySource entries.
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// CasePreserved shall indicate that the case of file names is preserved by
	// the file system. A value of True shall indicate that case of file names
	// shall be preserved.
	CasePreserved bool
	// CaseSensitive shall indicate that case sensitive file names are supported
	// by the file system. A value of True shall indicate that file names are
	// case sensitive.
	CaseSensitive bool
	// CharacterCodeSet shall be an array containing entries for the character
	// sets or encodings supported by the file system. Each entry shall specify
	// a character set encoding supported by the file system.
	CharacterCodeSet []CharacterCodeSet
	// ClusterSizeBytes shall specify the minimum file allocation size imposed
	// by the file system. This minimum allocation size shall be the smallest
	// amount of storage allocated to a file by the file system. Under stress
	// conditions, the file system may allocate storage in amounts smaller than
	// this value.
	ClusterSizeBytes int
	// Description provides a description of this resource.
	Description string
	// Identifiers shall contain a list of all known durable names for this file
	// system.
	Identifiers []common.Identifier
	// IOStatistics shall represent IO statistics for this FileSystem.
	IOStatistics IOStatistics
	// LowSpaceWarningThresholdPercents shall be an array containing entries for
	// the percentages of file system capacity at which low space warning events
	// are be issued. A LOW_SPACE_THRESHOLD_WARNING event shall be triggered
	// each time the remaining file system capacity value becomes less than one
	// of the values in the array. The following shall be true: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes)
	LowSpaceWarningThresholdPercents []int
	// MaxFileNameLengthBytes shall specify the maximum length of a file name
	// within the file system.
	MaxFileNameLengthBytes int64
	// RecoverableCapacitySourceCount is the number of available capacity source
	// resources currently available in the event that an equivalent capacity
	// source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent is If present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// ReplicaInfo if this file system is a replica, this value shall
	// describe its replication attributes. This value shall not be present
	// if this file system is not a replica. A file system may be both a
	// source and a replica.
	ReplicaInfo ReplicaInfo
	// ReplicaTargetCount is the number of replica targets.
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// exportedShares shall be an array of exported file shares of this file
	// system. Each entry shall define an exported file share of this file
	// system.
	exportedShares string // FileShareCollection
	// importedShares shall be an array of imported file shares.
	// importedShares []string // ImportedShare
	// ReplicaTargets shall reference the target replicas that
	// are sourced by this replica.
	replicaTargets []string
	classOfService string
	// ReplicaCollectionCount is the number of replica collections.
	ReplicaCollectionCount int
	spareResourceSets      []string
	// SpareResourceSetsCount is the number of spare resource sets.
	SpareResourceSetsCount int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a FileSystem object from the raw JSON.
func (filesystem *FileSystem) UnmarshalJSON(b []byte) error {
	type temp FileSystem
	type links struct {
		// ClassOfService shall be a link to the ClassOfService for this file
		// system.
		ClassOfService common.Link
		// ReplicaCollectionCount is the number of replica collections.
		ReplicaCollectionCount int `json:"ReplicaCollection@odata.count"`
		// SpareResourceSets each referenced SpareResourceSet shall contain
		// resources that may be utilized to replace the capacity provided by a
		// failed resource having a compatible type.
		SpareResourceSets common.Links
		// SpareResourceSetsCount is the number of spare resource sets.
		SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	}
	var t struct {
		temp
		ExportedShares common.Link
		ReplicaTargets common.Links
		Links          links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*filesystem = FileSystem(t.temp)

	// Extract the links to other entities for later
	filesystem.exportedShares = string(t.ExportedShares)
	filesystem.replicaTargets = t.ReplicaTargets.ToStrings()
	filesystem.classOfService = string(t.Links.ClassOfService)
	filesystem.ReplicaCollectionCount = t.Links.ReplicaCollectionCount
	filesystem.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	filesystem.SpareResourceSetsCount = t.Links.SpareResourceSetsCount

	// This is a read/write object, so we need to save the raw object data for later
	filesystem.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (filesystem *FileSystem) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(FileSystem)
	err := original.UnmarshalJSON(filesystem.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccessCapabilities",
		"CapacitySources",
		"CasePreserved",
		"CaseSensitive",
		"CharacterCodeSet",
		"ClusterSizeBytes",
		"ExportedShares",
		"LowSpaceWarningThresholdPercents",
		"MaxFileNameLengthBytes",
		"RecoverableCapacitySourceCount",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(filesystem).Elem()

	return filesystem.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFileSystem will get a FileSystem instance from the service.
func GetFileSystem(c common.Client, uri string) (*FileSystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var filesystem FileSystem
	err = json.NewDecoder(resp.Body).Decode(&filesystem)
	if err != nil {
		return nil, err
	}

	filesystem.SetClient(c)
	return &filesystem, nil
}

// ListReferencedFileSystems gets the collection of FileSystem from
// a provided reference.
func ListReferencedFileSystems(c common.Client, link string) ([]*FileSystem, error) {
	var result []*FileSystem
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, filesystemLink := range links.ItemLinks {
		filesystem, err := GetFileSystem(c, filesystemLink)
		if err != nil {
			return result, err
		}
		result = append(result, filesystem)
	}

	return result, nil
}

// ExportedShares gets the exported file shares for this file system.
func (filesystem *FileSystem) ExportedShares() ([]*FileShare, error) {
	return ListReferencedFileShares(filesystem.Client, filesystem.exportedShares)
}

// ClassOfService gets the filesystem's class of service.
func (filesystem *FileSystem) ClassOfService() (*ClassOfService, error) {
	var result *ClassOfService
	if filesystem.classOfService == "" {
		return result, nil
	}
	return GetClassOfService(filesystem.Client, filesystem.classOfService)
}

// SpareResourceSets gets the spare resource sets used for this filesystem.
func (filesystem *FileSystem) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet
	for _, rsLink := range filesystem.spareResourceSets {
		rs, err := GetSpareResourceSet(filesystem.Client, rsLink)
		if err != nil {
			return result, err
		}
		result = append(result, rs)
	}

	return result, nil
}
