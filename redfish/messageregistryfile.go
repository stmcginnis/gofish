//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2017.1 - #MessageRegistryFile.v1_1_5.MessageRegistryFile

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// MessageRegistryFile shall represent the registry file locator for a Redfish
// implementation.
type MessageRegistryFile struct {
	common.Entity
	// Languages This property contains a set of RFC5646-conformant language codes.
	Languages []string
	// Location shall contain the location information for this registry file.
	Location []MessageRegistryFileLocation
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Registry shall contain the registry name and it major and minor versions, as
	// defined by the Redfish Specification. This registry can be any type of
	// registry, such as message registry, privilege registry, or attribute
	// registry.
	Registry string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MessageRegistryFile object from the raw JSON.
func (m *MessageRegistryFile) UnmarshalJSON(b []byte) error {
	type temp MessageRegistryFile
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MessageRegistryFile(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	m.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MessageRegistryFile) Update() error {
	readWriteFields := []string{
		"Location",
	}

	return m.UpdateFromRawData(m, m.rawData, readWriteFields)
}

// GetMessageRegistryFile will get a MessageRegistryFile instance from the service.
func GetMessageRegistryFile(c common.Client, uri string) (*MessageRegistryFile, error) {
	return common.GetObject[MessageRegistryFile](c, uri)
}

// ListReferencedMessageRegistryFiles gets the collection of MessageRegistryFile from
// a provided reference.
func ListReferencedMessageRegistryFiles(c common.Client, link string) ([]*MessageRegistryFile, error) {
	return common.GetCollectionObjects[MessageRegistryFile](c, link)
}

// MessageRegistryFileLocation shall contain the location information for a registry file.
type MessageRegistryFileLocation struct {
	// ArchiveFile shall contain the file name of the individual registry file
	// within the archive file specified by the 'ArchiveURI' property. The file
	// name shall conform to the Redfish Specification-specified syntax.
	ArchiveFile string
	// ArchiveURI shall contain a URI that is colocated with the Redfish service
	// that specifies the location of the registry file, which can be retrieved
	// using the Redfish protocol and authentication methods. This property shall
	// be used for only ZIP or other archive files. The 'ArchiveFile' property
	// shall contain the file name of the individual registry file within the
	// archive file.
	ArchiveURI string
	// Language shall contain an RFC5646-conformant language code or 'default'.
	Language string
	// PublicationURI shall contain a URI not colocated with the Redfish service
	// that specifies the canonical location of the registry file. This property
	// shall be used for only individual registry files.
	PublicationURI string
	// URI shall contain a URI colocated with the Redfish service that specifies
	// the location of the registry file, which can be retrieved using the Redfish
	// protocol and authentication methods. This property shall be used for only
	// individual registry files. The file name portion of the URI shall conform to
	// Redfish Specification-specified syntax.
	URI string
}
