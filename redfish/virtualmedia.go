//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ConnectedVia are the ways the media may be connected.
type ConnectedVia string

const (

	// NotConnectedConnectedVia No current connection.
	NotConnectedConnectedVia ConnectedVia = "NotConnected"
	// URIConnectedVia Connected to a URI location.
	URIConnectedVia ConnectedVia = "URI"
	// AppletConnectedVia Connected to a client application.
	AppletConnectedVia ConnectedVia = "Applet"
	// OemConnectedVia Connected through an OEM-defined method.
	OemConnectedVia ConnectedVia = "Oem"
)

type EjectPolicy string

const (
	// OnPowerOffEjectPolicy The virtual media ejection occurs during a system power or reset event.
	OnPowerOffEjectPolicy EjectPolicy = "OnPowerOff"
	// SessionEjectPolicy The virtual media ejection occurs when a session is terminated. The session might be outside
	// the Redfish service.
	SessionEjectPolicy EjectPolicy = "Session"
	// TimedEjectPolicy The virtual media ejection occurs when a timer configured by the EjectTimeout property expires.
	TimedEjectPolicy EjectPolicy = "Timed"
	// AfterUseEjectPolicy The virtual media ejection occurs after the media is used.
	AfterUseEjectPolicy EjectPolicy = "AfterUse"
	// PersistentEjectPolicy The virtual media mount information persists indefinitely.
	PersistentEjectPolicy EjectPolicy = "Persistent"
)

// VirtualMediaType is the type of media.
type VirtualMediaType string

const (

	// CDMediaType A CD-ROM format (ISO) image.
	CDMediaType VirtualMediaType = "CD"
	// FloppyMediaType A floppy disk image.
	FloppyMediaType VirtualMediaType = "Floppy"
	// USBStickMediaType An emulation of a USB storage device.
	USBStickMediaType VirtualMediaType = "USBStick"
	// DVDMediaType A DVD-ROM format image.
	DVDMediaType VirtualMediaType = "DVD"
)

// TransferMethod is how the data is transferred.
type TransferMethod string

const (

	// StreamTransferMethod Stream image file data from the source URI.
	StreamTransferMethod TransferMethod = "Stream"
	// UploadTransferMethod Upload the entire image file from the source URI
	// to the service.
	UploadTransferMethod TransferMethod = "Upload"
)

// TransferProtocolType is the protocol used to transfer.
type TransferProtocolType string

const (

	// CIFSTransferProtocolType Common Internet File System (CIFS).
	CIFSTransferProtocolType TransferProtocolType = "CIFS"
	// FTPTransferProtocolType File Transfer Protocol (FTP).
	FTPTransferProtocolType TransferProtocolType = "FTP"
	// SFTPTransferProtocolType Secure File Transfer Protocol (SFTP).
	SFTPTransferProtocolType TransferProtocolType = "SFTP"
	// HTTPTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPTransferProtocolType TransferProtocolType = "HTTP"
	// HTTPSTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSTransferProtocolType TransferProtocolType = "HTTPS"
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType A manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
)

// VirtualMedia shall represent a virtual media service for a Redfish
// implementation.
type VirtualMedia struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represents the
	// server certificates for the server referenced by the Image property. If VerifyCertificate is 'true', services
	// shall compare the certificates in this collection with the certificate obtained during handshaking with the
	// image server in order to verify the identity of the image server prior to completing the remote media
	// connection. If the server cannot be verified, the service shall not complete the remote media connection. If
	// VerifyCertificate is 'false', the service shall not perform certificate verification with certificates in this
	// collection. Regardless of the contents of this collection, services may perform additional verification based on
	// other factors, such as the configuration of the SecurityPolicy resource.
	certificates []string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the Image property as part of TLS
	// handshaking.
	clientCertificates []string
	// ConnectedVia shall contain the current connection method from a client to the virtual media that this Resource
	// represents.
	ConnectedVia ConnectedVia
	// Description provides a description of this resource.
	Description string
	// EjectPolicy shall contain the ejection policy for the virtual media.
	EjectPolicy EjectPolicy
	// EjectTimeout shall indicate the amount of time before virtual media is automatically ejected when EjectPolicy
	// contains 'Timed'.
	EjectTimeout string
	// Image shall contain an URI. A null value indicated
	// no image connection.
	Image string
	// ImageName shall contain the name of the image.
	ImageName string
	// Inserted shall indicate whether media is present in the virtual media device.
	Inserted bool
	// MediaTypes shall be the supported media types for this connection.
	MediaTypes []VirtualMediaType
	// Password shall represent the password to access the
	// Image parameter-specified URI. The value shall be null in responses.
	Password string
	// TransferMethod shall describe how the image transfer
	// occurs.
	TransferMethod TransferMethod
	// TransferProtocolType shall represent the network
	// protocol to use with the specified image URI.
	TransferProtocolType TransferProtocolType
	// UserName shall represent the user name to access the
	// Image parameter-specified URI.
	UserName string
	// VerifyCertificate shall indicate whether the service will verify the certificate of the server referenced by the
	// Image property prior to completing the remote media connection with the certificates found in the collection
	// referenced by the Certificates property. If this property is not supported by the service, it shall be assumed
	// to be 'false'. This property should default to 'false' in order to maintain compatibility with older clients.
	// Regardless of the value of this property, services may perform additional verification based on other factors,
	// such as the configuration of the SecurityPolicy resource.
	VerifyCertificate bool
	// WriteProtected shall indicate whether the remote
	// device media prevents writing to that media.
	WriteProtected bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// ejectMediaTarget is the URL to send EjectMedia requests.
	ejectMediaTarget string
	// insertMediaTarget is the URL to send InsertMedia requests.
	insertMediaTarget string
	// SupportsMediaEject indicates if this implementation supports ejecting
	// virtual media or not (added in schema 1.2.0).
	SupportsMediaEject bool
	// SupportsMediaInsert indicates if this implementation supports inserting
	// virtual media or not (added in schema 1.2.0).
	SupportsMediaInsert bool
}

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (virtualmedia *VirtualMedia) UnmarshalJSON(b []byte) error {
	type temp VirtualMedia
	type actions struct {
		EjectMedia  common.ActionTarget `json:"#VirtualMedia.EjectMedia"`
		InsertMedia common.ActionTarget `json:"#VirtualMedia.InsertMedia"`
	}
	var t struct {
		temp
		Actions           actions
		Certificates      common.LinksCollection
		ClientCertificate common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualmedia = VirtualMedia(t.temp)

	// Extract the links to other entities for later
	virtualmedia.certificates = t.Certificates.ToStrings()
	virtualmedia.clientCertificates = t.ClientCertificate.ToStrings()

	virtualmedia.ejectMediaTarget = t.Actions.EjectMedia.Target
	virtualmedia.insertMediaTarget = t.Actions.InsertMedia.Target

	virtualmedia.SupportsMediaEject = (virtualmedia.ejectMediaTarget != "")
	virtualmedia.SupportsMediaInsert = (virtualmedia.insertMediaTarget != "")

	// This is a read/write object, so we need to save the raw object data for later
	virtualmedia.rawData = b

	return nil
}

// Certificates gets the the server certificates for the server referenced by the
// Image property. If VerifyCertificate is `true`, services shall compare the
// certificates in this collection with the certificate obtained during handshaking
// with the image server in order to verify the identity of the image server prior
// to completing the remote media connection. If the server cannot be verified,
// the service shall not complete the remote media connection. If VerifyCertificate
// is `false`, the service shall not perform certificate verification with
// certificates in this collection. Regardless of the contents of this collection,
// services may perform additional verification based on other factors, such as the
// configuration of the SecurityPolicy resource.
func (virtualmedia *VirtualMedia) Certificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](virtualmedia.GetClient(), virtualmedia.certificates)
}

// ClientCertificates gets the client identity certificates that are provided to
// the server referenced by the Image property as part of TLS handshaking.
func (virtualmedia *VirtualMedia) ClientCertificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](virtualmedia.GetClient(), virtualmedia.clientCertificates)
}

// Update commits updates to this object's properties to the running system.
func (virtualmedia *VirtualMedia) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(VirtualMedia)
	err := original.UnmarshalJSON(virtualmedia.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"EjectPolicy",
		"EjectTimeout",
		"Image",
		"Inserted",
		"MediaType",
		"Password",
		"TransferMethod",
		"TransferProtocolType",
		"UserName",
		"WriteProtected",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(virtualmedia).Elem()

	return virtualmedia.Entity.Update(originalElement, currentElement, readWriteFields)
}

// EjectMedia sends a request to eject the media.
func (virtualmedia *VirtualMedia) EjectMedia() error {
	if !virtualmedia.SupportsMediaEject {
		return errors.New("redfish service does not support VirtualMedia.EjectMedia calls")
	}

	return virtualmedia.Post(virtualmedia.ejectMediaTarget, struct{}{})
}

// InsertMedia sends a request to insert virtual media.
func (virtualmedia *VirtualMedia) InsertMedia(image string, inserted, writeProtected bool) error {
	if !virtualmedia.SupportsMediaInsert {
		return errors.New("redfish service does not support VirtualMedia.InsertMedia calls")
	}

	t := struct {
		Image          string
		Inserted       bool
		WriteProtected bool
	}{
		Image:          image,
		Inserted:       inserted,
		WriteProtected: writeProtected,
	}

	return virtualmedia.Post(virtualmedia.insertMediaTarget, t)
}

// VirtualMediaConfig is an struct used to pass config data to build the HTTP body when inserting media
type VirtualMediaConfig struct {
	Image                string
	Inserted             bool                 `json:",omitempty"`
	MediaType            VirtualMediaType     `json:",omitempty"`
	Password             string               `json:",omitempty"`
	TransferMethod       TransferMethod       `json:",omitempty"`
	TransferProtocolType TransferProtocolType `json:",omitempty"`
	UserName             string               `json:",omitempty"`
	WriteProtected       bool                 `json:",omitempty"`
}

// InsertMediaConfig sends a request to insert virtual media using the VirtualMediaConfig struct
func (virtualmedia *VirtualMedia) InsertMediaConfig(config VirtualMediaConfig) error { //nolint
	if !virtualmedia.SupportsMediaInsert {
		return errors.New("redfish service does not support VirtualMedia.InsertMedia calls")
	}

	return virtualmedia.Post(virtualmedia.insertMediaTarget, config)
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c common.Client, uri string) (*VirtualMedia, error) {
	return common.GetObject[VirtualMedia](c, uri)
}

// ListReferencedVirtualMedias gets the collection of VirtualMedia from
// a provided reference.
func ListReferencedVirtualMedias(c common.Client, link string) ([]*VirtualMedia, error) {
	return common.GetCollectionObjects[VirtualMedia](c, link)
}
