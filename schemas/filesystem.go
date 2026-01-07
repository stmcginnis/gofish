//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #FileSystem.v1_4_2.FileSystem

package schemas

import (
	"encoding/json"
)

// CharacterCodeSet is The values shall indicate the character code standards
// supported by the file system.
type CharacterCodeSet string

const (
	// ASCIICharacterCodeSet shall indicate that the ASCII character encoding is
	// supported by the file system.
	ASCIICharacterCodeSet CharacterCodeSet = "ASCII"
	// UnicodeCharacterCodeSet shall indicate that Unicode character encoding is
	// supported by the file system.
	UnicodeCharacterCodeSet CharacterCodeSet = "Unicode"
	// ISO2022CharacterCodeSet shall indicate that ISO-2022 character encoding is
	// supported by the file system.
	ISO2022CharacterCodeSet CharacterCodeSet = "ISO2022"
	// ISO88591CharacterCodeSet shall indicate that ISO-8859-1 character encoding
	// is supported by the file system.
	ISO88591CharacterCodeSet CharacterCodeSet = "ISO8859_1"
	// ExtendedUNIXCodeCharacterCodeSet shall indicate that Extended Unix Code
	// character encoding is supported by the file system.
	ExtendedUNIXCodeCharacterCodeSet CharacterCodeSet = "ExtendedUNIXCode"
	// UTF8CharacterCodeSet shall indicate that the UTF-8 character encoding is
	// supported by the file system.
	UTF8CharacterCodeSet CharacterCodeSet = "UTF_8"
	// UTF16CharacterCodeSet shall indicate that the UTF-16 character encoding is
	// supported by the file system.
	UTF16CharacterCodeSet CharacterCodeSet = "UTF_16"
	// UCS2CharacterCodeSet shall indicate that the UCS-2 character encoding is
	// supported by the file system.
	UCS2CharacterCodeSet CharacterCodeSet = "UCS_2"
)

// FileProtocol is The values shall indicate the file sharing protocols
// supported by the file system. At least one value shall be present.
type FileProtocol string

const (
	// NFSv3FileProtocol shall indicate that NFSv3, as defined in RFC 1813, is
	// supported by the file system.
	NFSv3FileProtocol FileProtocol = "NFSv3"
	// NFSv40FileProtocol shall indicate that NFSv4, as defined in RFC 7530, is
	// supported by the file system.
	NFSv40FileProtocol FileProtocol = "NFSv4_0"
	// NFSv41FileProtocol shall indicate that NFSv4.1, as defined in RFC 5661, is
	// supported by the file system.
	NFSv41FileProtocol FileProtocol = "NFSv4_1"
	// SMBv20FileProtocol shall indicate that Server Message Block version 2.0 is
	// supported by the file system.
	SMBv20FileProtocol FileProtocol = "SMBv2_0"
	// SMBv21FileProtocol shall indicate that Server Message Block version 2.1 is
	// supported by the file system.
	SMBv21FileProtocol FileProtocol = "SMBv2_1"
	// SMBv30FileProtocol shall indicate that Server Message Block version 3.0 is
	// supported by the file system.
	SMBv30FileProtocol FileProtocol = "SMBv3_0"
	// SMBv302FileProtocol shall indicate that Server Message Block version 3.0.2
	// is supported by the file system.
	SMBv302FileProtocol FileProtocol = "SMBv3_0_2"
	// SMBv311FileProtocol shall indicate that Server Message Block version 3.1.1
	// is supported by the file system.
	SMBv311FileProtocol FileProtocol = "SMBv3_1_1"
)

// FileSystem shall be used to represent an instance of a hierarchical namespace
// of files.
type FileSystem struct {
	Entity
	// AccessCapabilities shall be an array containing entries for the supported IO
	// access capabilities. Each entry shall specify a current storage access
	// capability.
	AccessCapabilities []StorageAccessCapability
	// BlockSizeBytes shall be the block size of the file system in bytes.
	BlockSizeBytes *int `json:",omitempty"`
	// Capacity shall be the capacity allocated to the file system in bytes.
	Capacity Capacity
	// CapacitySources shall be an array containing entries for all the capacity
	// sources for the file system. Each entry shall provide capacity allocation
	// information from a named resource.
	CapacitySources []CapacitySource
	// CapacitySourcesCount
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// CasePreserved shall indicate that the case of file names is preserved by the
	// file system. A value of True shall indicate that case of file names shall be
	// preserved.
	CasePreserved bool
	// CaseSensitive shall indicate that case sensitive file names are supported by
	// the file system. A value of True shall indicate that file names are case
	// sensitive.
	CaseSensitive bool
	// CharacterCodeSet shall be an array containing entries for the character sets
	// or encodings supported by the file system. Each entry shall specify a
	// character set encoding supported by the file system.
	CharacterCodeSet []CharacterCodeSet
	// ClusterSizeBytes shall specify the minimum file allocation size imposed by
	// the file system. This minimum allocation size shall be the smallest amount
	// of storage allocated to a file by the file system. Under stress conditions,
	// the file system may allocate storage in amounts smaller than this value.
	ClusterSizeBytes *uint `json:",omitempty"`
	// ExportedShares shall be an array of exported file shares of this file
	// system. Each entry shall define an exported file share of this file system.
	exportedShares string
	// IOStatistics shall represent IO statistics for this FileSystem.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.4.0
	// This property is deprecated in favor of the IOStatistics property in
	// FileSystemMetrics.
	IOStatistics IOStatistics
	// Identifiers shall contain a list of all known durable names for this file
	// system.
	//
	// Version added: v1.1.1
	Identifiers []Identifier
	// ImportedShares shall be an array of imported file shares.
	//
	// Version added: v1.0.1
	// Defined in the spec as `"anyOf": []`
	ImportedShares []any
	// LowSpaceWarningThresholdPercents shall be an array containing entries for
	// the percentages of file system capacity at which low space warning events
	// are be issued. A LOW_SPACE_THRESHOLD_WARNING event shall be triggered each
	// time the remaining file system capacity value becomes less than one of the
	// values in the array. The following shall be true: Across all CapacitySources
	// entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []*int
	// MaxFileNameLengthBytes shall specify the maximum length of a file name
	// within the file system.
	MaxFileNameLengthBytes *uint `json:",omitempty"`
	// Metrics shall contain a link to a resource of type FileSystemMetrics that
	// specifies the metrics for this file system. IO metrics are reported in the
	// IOStatistics property.
	//
	// Version added: v1.4.0
	metrics string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RecoverableCapacitySourceCount The value is the number of available capacity
	// source resources currently available in the event that an equivalent
	// capacity source resource fails.
	//
	// Version added: v1.2.0
	RecoverableCapacitySourceCount *int `json:",omitempty"`
	// RemainingCapacity shall be the remaining capacity allocated to the file
	// system in bytes.
	//
	// Deprecated
	// This property is deprecated in favor of the Capacity property. The
	// RemainingCapacity can be computed from the values of that property, for each
	// of the sub groups xxx: Data, MetaData, and Snapshot. The RemainingCapacity
	// is xxx/ProvisionedBytes - xxx/ConsumedBytes.
	RemainingCapacity Capacity
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) -
	// SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100 represented as an integer
	// value.
	//
	// Version added: v1.1.0
	RemainingCapacityPercent *int `json:",omitempty"`
	// ReplicaInfo shall describe its replication attributes. This value shall not
	// be present if this file system is not a replica. A file system may be both a
	// source and a replica.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this
	// replica.
	//
	// Version added: v1.2.1
	ReplicaTargets []Entity
	// ReplicaTargetsCount
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on
	// the file system. This property shall be consistent with the state reflected
	// at the storage pool level.
	//
	// Version added: v1.3.0
	ReplicationEnabled bool
	// classOfService is the URI for ClassOfService.
	classOfService string
	// replicaCollection are the URIs for ReplicaCollection.
	replicaCollection []string
	// spareResourceSets are the URIs for SpareResourceSets.
	spareResourceSets []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a FileSystem object from the raw JSON.
func (f *FileSystem) UnmarshalJSON(b []byte) error {
	type temp FileSystem
	type fLinks struct {
		ClassOfService    Link  `json:"ClassOfService"`
		ReplicaCollection Links `json:"ReplicaCollection"`
		SpareResourceSets Links `json:"SpareResourceSets"`
	}
	var tmp struct {
		temp
		Links          fLinks
		ExportedShares Link `json:"ExportedShares"`
		Metrics        Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FileSystem(tmp.temp)

	// Extract the links to other entities for later
	f.classOfService = tmp.Links.ClassOfService.String()
	f.replicaCollection = tmp.Links.ReplicaCollection.ToStrings()
	f.spareResourceSets = tmp.Links.SpareResourceSets.ToStrings()
	f.exportedShares = tmp.ExportedShares.String()
	f.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *FileSystem) Update() error {
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
		"ReplicationEnabled",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetFileSystem will get a FileSystem instance from the service.
func GetFileSystem(c Client, uri string) (*FileSystem, error) {
	return GetObject[FileSystem](c, uri)
}

// ListReferencedFileSystems gets the collection of FileSystem from
// a provided reference.
func ListReferencedFileSystems(c Client, link string) ([]*FileSystem, error) {
	return GetCollectionObjects[FileSystem](c, link)
}

// ClassOfService gets the ClassOfService linked resource.
func (f *FileSystem) ClassOfService() (*ClassOfService, error) {
	if f.classOfService == "" {
		return nil, nil
	}
	return GetObject[ClassOfService](f.client, f.classOfService)
}

// ReplicaCollection gets the ReplicaCollection linked resources.
func (f *FileSystem) ReplicaCollection() ([]*FileSystem, error) {
	return GetObjects[FileSystem](f.client, f.replicaCollection)
}

// SpareResourceSets gets the SpareResourceSets linked resources.
func (f *FileSystem) SpareResourceSets() ([]*SpareResourceSet, error) {
	return GetObjects[SpareResourceSet](f.client, f.spareResourceSets)
}

// ExportedShares gets the ExportedShares collection.
func (f *FileSystem) ExportedShares() ([]*FileShare, error) {
	if f.exportedShares == "" {
		return nil, nil
	}
	return GetCollectionObjects[FileShare](f.client, f.exportedShares)
}

// Metrics gets the Metrics linked resource.
func (f *FileSystem) Metrics() (*FileSystemMetrics, error) {
	if f.metrics == "" {
		return nil, nil
	}
	return GetObject[FileSystemMetrics](f.client, f.metrics)
}
