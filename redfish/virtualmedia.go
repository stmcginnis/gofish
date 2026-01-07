//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #VirtualMedia.v1_6_5.VirtualMedia

package redfish

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stmcginnis/gofish/common"
)

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
	// OnPowerOffEjectPolicy The virtual media ejection occurs during a system
	// power or reset event.
	OnPowerOffEjectPolicy EjectPolicy = "OnPowerOff"
	// SessionEjectPolicy The virtual media ejection occurs when a session is
	// terminated. The session might be outside the Redfish service.
	SessionEjectPolicy EjectPolicy = "Session"
	// TimedEjectPolicy The virtual media ejection occurs when a timer configured
	// by the 'EjectTimeout' property expires.
	TimedEjectPolicy EjectPolicy = "Timed"
	// AfterUseEjectPolicy The virtual media ejection occurs after the media is
	// used.
	AfterUseEjectPolicy EjectPolicy = "AfterUse"
	// PersistentEjectPolicy The virtual media mount information persists
	// indefinitely.
	PersistentEjectPolicy EjectPolicy = "Persistent"
)

type VirtualMediaType string

const (
	// CDVirtualMediaType is a CD-ROM format (ISO) image.
	CDVirtualMediaType VirtualMediaType = "CD"
	// FloppyVirtualMediaType is a floppy disk image.
	FloppyVirtualMediaType VirtualMediaType = "Floppy"
	// USBStickVirtualMediaType is an emulation of a USB storage device.
	USBStickVirtualMediaType VirtualMediaType = "USBStick"
	// DVDVirtualMediaType is a DVD-ROM format image.
	DVDVirtualMediaType VirtualMediaType = "DVD"
)

type TransferMethod string

const (
	// StreamTransferMethod Stream image file data from the source URI.
	StreamTransferMethod TransferMethod = "Stream"
	// UploadTransferMethod Upload the entire image file from the source URI to the
	// service.
	UploadTransferMethod TransferMethod = "Upload"
)

// VirtualMedia shall represent a virtual media service for a Redfish
// implementation.
type VirtualMedia struct {
	common.Entity
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the server certificates for the
	// server referenced by the 'Image' property. If 'VerifyCertificate' is 'true',
	// services shall compare the certificates in this collection with the
	// certificate obtained during handshaking with the image server in order to
	// verify the identity of the image server prior to completing the remote media
	// connection. If the server cannot be verified, the service shall not complete
	// the remote media connection. If 'VerifyCertificate' is 'false', the service
	// shall not perform certificate verification with certificates in this
	// collection. Regardless of the contents of this collection, services may
	// perform additional verification based on other factors, such as the
	// configuration of the 'SecurityPolicy' resource.
	//
	// Version added: v1.4.0
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates
	// that are provided to the server referenced by the 'Image' property as part
	// of TLS handshaking.
	//
	// Version added: v1.5.0
	clientCertificates string
	// ConnectedVia shall contain the current connection method from a client to
	// the virtual media that this resource represents.
	ConnectedVia ConnectedVia
	// EjectPolicy shall contain the ejection policy for the virtual media.
	//
	// Version added: v1.6.0
	EjectPolicy EjectPolicy
	// EjectTimeout shall indicate the amount of time before virtual media is
	// automatically ejected when 'EjectPolicy' contains 'Timed'.
	//
	// Version added: v1.6.0
	EjectTimeout string
	// Image shall contain the URI of the media attached to the virtual media. This
	// value may specify an absolute URI to remote media or a relative URI to media
	// local to the implementation. A service may allow a relative URI to reference
	// a SoftwareInventory resource. The value 'null' shall indicate no image
	// connection. If the URI specified by this property is not accessible,
	// services should populate the 'Conditions' property in 'Status' with messages
	// that describe the condition.
	Image string
	// ImageName shall contain the name of the image.
	ImageName string
	// Inserted shall indicate whether the media is mounted and visible to the host
	// system.
	Inserted *bool
	// MediaTypes shall contain an array of the supported media types for this
	// connection.
	MediaTypes []VirtualMediaType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password to access the URI specified by the
	// 'Image' property. The value shall be 'null' in responses.
	//
	// Version added: v1.3.0
	Password string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.4.0
	Status common.Status
	// TransferMethod shall describe how the image transfer occurs.
	//
	// Version added: v1.3.0
	TransferMethod TransferMethod
	// TransferProtocolType shall contain network protocol to use with the URI
	// specified by the 'Image' property.
	//
	// Version added: v1.3.0
	TransferProtocolType TransferProtocolType
	// UserName shall contain the username to access the URI specified by the
	// 'Image' property.
	//
	// Version added: v1.3.0
	UserName string
	// VerifyCertificate shall indicate whether the service will verify the
	// certificate of the server referenced by the 'Image' property prior to
	// completing the remote media connection with the certificates found in the
	// collection referenced by the 'Certificates' property. If this property is
	// not supported by the service, it shall be assumed to be 'false'. This
	// property should default to 'false' in order to maintain compatibility with
	// older clients. Regardless of the value of this property, services may
	// perform additional verification based on other factors, such as the
	// configuration of the SecurityPolicy resource.
	//
	// Version added: v1.4.0
	VerifyCertificate bool
	// WriteProtected shall indicate whether the remote device media prevents
	// writing to that media.
	WriteProtected bool
	// ejectMedia is the URL to send EjectMedia requests.
	ejectMedia common.ActionTarget
	// SupportsMediaEject indicates if ejecting virtual media is supported.
	SupportsMediaEject bool

	// insertMedia is the URL to send InsertMedia requests.
	insertMedia common.ActionTarget
	// SupportsMediaInsert indicates if inserting virtual media is supported.
	SupportsMediaInsert bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (v *VirtualMedia) UnmarshalJSON(b []byte) error {
	type temp VirtualMedia
	type vActions struct {
		EjectMedia  common.ActionTarget `json:"#VirtualMedia.EjectMedia"`
		InsertMedia common.ActionTarget `json:"#VirtualMedia.InsertMedia"`
	}
	var tmp struct {
		temp
		Actions            vActions
		Certificates       common.Link `json:"certificates"`
		ClientCertificates common.Link `json:"clientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VirtualMedia(tmp.temp)

	// Extract the links to other entities for later
	v.ejectMedia = tmp.Actions.EjectMedia
	v.insertMedia = tmp.Actions.InsertMedia
	v.SupportsMediaEject = v.ejectMedia.Target != ""
	v.SupportsMediaInsert = v.insertMedia.Target != ""
	v.certificates = tmp.Certificates.String()
	v.clientCertificates = tmp.ClientCertificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	v.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *VirtualMedia) Update() error {
	readWriteFields := []string{
		"EjectPolicy",
		"EjectTimeout",
		"Image",
		"Inserted",
		"Password",
		"Status",
		"TransferMethod",
		"TransferProtocolType",
		"UserName",
		"VerifyCertificate",
		"WriteProtected",
	}

	return v.UpdateFromRawData(v, v.rawData, readWriteFields)
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

// EjectMediaActionInfo provides the ActionInfo, if supported, for an EjectMedia Action
func (v *VirtualMedia) EjectMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia EjectMedia ActionInfo not supported by this service")
	}

	return common.GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// EjectMedia shall detach the remote media from the virtual media. At the
// completion of the operation, inserted shall be set to 'false' and the image
// name shall be cleared.
func (v *VirtualMedia) EjectMedia() error {
	if !v.SupportsMediaEject {
		return errors.New("redfish service does not support VirtualMedia.EjectMedia calls")
	}

	return v.Post(v.ejectMedia.Target, nil)
}

// VirtualMediaConfig is used to pass config data when inserting media.
type VirtualMediaConfig struct {
	Image                string                `json:"Image"`
	Inserted             *bool                 `json:"Inserted,omitempty"`
	MediaType            *VirtualMediaType     `json:"MediaType,omitempty"`
	Password             *string               `json:"Password,omitempty"`
	TransferMethod       *TransferMethod       `json:"TransferMethod,omitempty"`
	TransferProtocolType *TransferProtocolType `json:"TransferProtocolType,omitempty"`
	UserName             *string               `json:"UserName,omitempty"`
	WriteProtected       *bool                 `json:"WriteProtected,omitempty"`
}

// InsertMediaActionInfo provides the ActionInfo, if supported, for an InsertMedia Action
func (v *VirtualMedia) InsertMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia InsertMedia ActionInfo not supported by this service")
	}

	return common.GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// InsertMedia sends a request to insert virtual media.
func (v *VirtualMedia) InsertMedia(image string, inserted, writeProtected bool) (*Task, error) {
	return v.InsertMediaConfig(VirtualMediaConfig{
		Image:          image,
		Inserted:       &inserted,
		WriteProtected: &writeProtected,
	})
}

// InsertMediaConfig sends a request to insert virtual media using the VirtualMediaConfig struct.
func (v *VirtualMedia) InsertMediaConfig(config VirtualMediaConfig) (*Task, error) {
	if !v.SupportsMediaInsert {
		return nil, errors.New("redfish service does not support VirtualMedia.InsertMedia calls")
	}

	resp, err := v.PostWithResponse(v.insertMedia.Target, config)
	defer common.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusAccepted {
		return nil, nil
	}

	if location := resp.Header["Location"]; len(location) > 0 {
		return GetTask(v.GetClient(), location[0])
	}

	return nil, nil
}

// Certificates gets the Certificates collection.
func (v *VirtualMedia) Certificates(client common.Client) ([]*Certificate, error) {
	if v.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, v.certificates)
}

// ClientCertificates gets the ClientCertificates collection.
func (v *VirtualMedia) ClientCertificates(client common.Client) ([]*Certificate, error) {
	if v.clientCertificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, v.clientCertificates)
}
