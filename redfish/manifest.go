//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type Expand string

const (
	// NoneExpand shall indicate that references in the manifest response will not be expanded.
	NoneExpand Expand = "None"
	// AllExpand shall indicate that all subordinate references in the manifest response will be expanded.
	AllExpand Expand = "All"
	// RelevantExpand shall indicate that relevant subordinate references in the manifest response will be expanded.
	RelevantExpand Expand = "Relevant"
)

type StanzaType string

const (
	// ComposeSystemStanzaType shall indicate a stanza that describes the specific, constrained, or mixed resources
	// required to compose a computer system. The resource blocks assigned to the computer system shall be moved to the
	// active pool. The Request property of the stanza shall contain a resource of type ComputerSystem that represents
	// the composition request. The Response property of the stanza shall contain a resource of type ComputerSystem
	// that represents the composed system or a Redfish Specification-defined error response.
	ComposeSystemStanzaType StanzaType = "ComposeSystem"
	// DecomposeSystemStanzaType shall indicate a stanza that references a computer system to decompose and return the
	// resource blocks to the free pool that are no longer contributing to composed resources. The Request property of
	// the stanza shall be a Redfish Specification-defined reference object containing a reference to the resource of
	// type ComputerSystem to decompose. The Response property of the stanza shall contain a resource of type
	// ComputerSystem that represents the decomposed system or a Redfish Specification-defined error response.
	DecomposeSystemStanzaType StanzaType = "DecomposeSystem"
	// ComposeResourceStanzaType shall indicate a stanza that describes a composed resource block. The resource blocks
	// assigned to the composed resource block shall be moved to the active pool. The Request property of the stanza
	// shall contain a resource of type ResourceBlock that represents the composition request. The Response property of
	// the stanza shall contain a resource of type ResourceBlock that represents the composed resource block or a
	// Redfish Specification-defined error response.
	ComposeResourceStanzaType StanzaType = "ComposeResource"
	// DecomposeResourceStanzaType shall indicate a stanza that references a composed resource block to decompose and
	// return the resource blocks to the free pool that are no longer contributing to composed resources. The Request
	// property of the stanza shall be a reference object as defined by the 'Reference properties' clause of the
	// Redfish Specification containing a reference to the resource of type ResourceBlock to decompose. The Response
	// property of the stanza shall contain a resource of type ResourceBlock that represents the decomposed resource
	// block or a Redfish Specification-defined error response.
	DecomposeResourceStanzaType StanzaType = "DecomposeResource"
	// OEMStanzaType shall indicate a stanza that describes an OEM-specific request. The OEMStanzaType property shall
	// contain the specific OEM stanza type.
	OEMStanzaType StanzaType = "OEM"
	// RegisterResourceBlockStanzaType shall indicate a stanza that references a resource to create a resource block
	// that references the resource and add it to the free pool. The Request property of the stanza shall contain a
	// resource of type ResourceBlock that represents the registration request. The Response property of the stanza
	// shall contain a resource of type ResourceBlock that represents the composed system or a Redfish Specification-
	// defined error response.
	RegisterResourceBlockStanzaType StanzaType = "RegisterResourceBlock"
)

// Manifest shall describe a manifest containing a set of requests to be fulfilled.
type Manifest struct {
	// The schema doesn't define this as a full Entity object, but it shouldn't hurt.
	common.Entity
	// Description provides a description of this resource.
	Description string
	// Expand shall contain the expansion control for references in manifest responses.
	Expand Expand
	// Stanzas shall contain an array of stanzas that describe the requests specified by this manifest.
	Stanzas []Stanza
	// Timestamp shall contain the date and time when the manifest was created.
	Timestamp string
}

// GetManifest will get a Manifest instance from the service.
func GetManifest(c common.Client, uri string) (*Manifest, error) {
	return common.GetObject[Manifest](c, uri)
}

// ListReferencedManifests gets the collection of Manifest from
// a provided reference.
func ListReferencedManifests(c common.Client, link string) ([]*Manifest, error) {
	return common.GetCollectionObjects[Manifest](c, link)
}

// Stanza shall contain properties that describe a request to be fulfilled within a manifest.
type Stanza struct {
	// OEMStanzaType shall contain the OEM-defined type of stanza. This property shall be present if StanzaType is
	// 'OEM'.
	OEMStanzaType string
	// Request shall contain the request details for the stanza, and the contents vary depending on the value of the
	// StanzaType property.
	Request json.RawMessage
	// Response shall contain the response details for the stanza, and the contents vary depending on the value of the
	// StanzaType property.
	Response json.RawMessage
	// StanzaID shall contain the identifier of the stanza.
	StanzaID string `json:"StanzaId"`
	// StanzaType shall contain the type of stanza.
	StanzaType StanzaType
}
