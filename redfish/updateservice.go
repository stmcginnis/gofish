//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type SupportedUpdateImageFormatType string

const (
	// PLDMv10SupportedUpdateImageFormatType shall indicate an image that conforms to the v1.0 image format as defined
	// in DMTF DSP0267.
	PLDMv10SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_0"
	// PLDMv11SupportedUpdateImageFormatType shall indicate an image that conforms to the v1.1 image format as defined
	// in DMTF DSP0267.
	PLDMv11SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_1"
	// PLDMv12SupportedUpdateImageFormatType shall indicate an image that conforms to the v1.2 image format as defined
	// in DMTF DSP0267.
	PLDMv12SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_2"
	// PLDMv13SupportedUpdateImageFormatType shall indicate an image that conforms to the v1.3 image format as defined
	// in DMTF DSP0267.
	PLDMv13SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_3"
	// UEFICapsuleSupportedUpdateImageFormatType shall indicate an image that conforms to the UEFI capsule format as
	// defined in the UEFI Specification.
	UEFICapsuleSupportedUpdateImageFormatType SupportedUpdateImageFormatType = "UEFICapsule"
	// VendorDefinedSupportedUpdateImageFormatType shall indicate a vendor-defined format.
	VendorDefinedSupportedUpdateImageFormatType SupportedUpdateImageFormatType = "VendorDefined"
)

// HTTPPushURIApplyTime shall contain settings for when to apply HttpPushUri-provided software.
type HTTPPushURIApplyTime struct {
	// ApplyTime shall indicate the time when to apply the HttpPushUri-provided software update.
	ApplyTime ApplyTime
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property. This property shall be required if the
	// HttpPushUriApplyTime property value is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds int
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the HttpPushUri-
	// provided software as part of a maintenance window. This property shall be required if the HttpPushUriApplyTime
	// property value is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
}

// HTTPPushURIOptions shall contain settings and requirements of the service for HttpPushUri-provided software
// updates.
type HTTPPushURIOptions struct {
	// ForceUpdate shall indicate whether the service should bypass update policies when applying the HttpPushUri-
	// provided image, such as allowing a component to be downgraded. Services may contain update policies that are
	// never bypassed, such as minimum version enforcement. If this property is not present, it shall be assumed to be
	// 'false'.
	ForceUpdate bool
	// HttpPushUriApplyTime shall contain settings for when to apply HttpPushUri-provided firmware.
	HTTPPushURIApplyTime HTTPPushURIApplyTime
}

// UpdateParameters shall contain the update parameters when passing the update image when using the URI specified
// by the MultipartHTTPPushURI property to push a software image.
type UpdateParameters struct {
	// ForceUpdate shall indicate whether the service should bypass update policies when applying the provided image,
	// such as allowing a component to be downgraded. Services may contain update policies that are never bypassed,
	// such as minimum version enforcement. If the client does not provide this parameter, the service shall default
	// this value to 'false'.
	ForceUpdate bool
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Targets shall contain zero or more URIs that indicate where to apply the update image when using the URI
	// specified by the MultipartHttpPushUri property to push a software image. These targets should correspond to
	// software inventory instances or their related items. If this property is not present or contains no targets, the
	// service shall apply the software image to all applicable targets, as determined by the service. If the target
	// specifies a device resource, the software image file shall be applied to the specified device. If the target
	// specifies a resource collection, the software image shall be applied to each applicable member of the specified
	// collection. If the target resource specifies an Aggregate resource, the software image file shall be applied to
	// each applicable element of the specified aggregate. If the target resource specifies a ComputerSystem resource,
	// the software image file shall be applied to the applicable components within the specified computer system.
	Targets []string
}

// UpdateService is used to represent the update service offered by the redfish API
type UpdateService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the ImageURI property in
	// SimpleUpdate as part of TLS handshaking.
	clientCertificates []string
	// Description provides a description of this resource.
	// Description provides a description of this resource.
	Description string
	// FirmwareInventory points towards the firmware store endpoint
	firmwareInventory string
	// HTTPPushURI endpoint is used to push (POST) firmware updates
	HTTPPushURI string `json:"HttpPushUri"`
	// HTTPPushURIOptions shall contain options and requirements of the service for HttpPushUri-provided software
	// updates.
	HTTPPushURIOptions HTTPPushURIOptions
	// HTTPPushURIOptionsBusy shall indicate whether a client uses the HttpPushUriOptions properties for software
	// updates. When a client uses any HttpPushUriOptions properties for software updates, it should set this property
	// to 'true'. When a client no longer uses HttpPushUriOptions properties for software updates, it should set this
	// property to 'false'. This property can provide multiple clients a way to negotiate ownership of
	// HttpPushUriOptions properties. Clients can use this property to determine whether another client uses
	// HttpPushUriOptions properties for software updates. This property has no functional requirements for the
	// service.
	HTTPPushURIOptionsBusy bool
	// HTTPPushURITargets shall contain zero or more URIs that indicate where to apply the update image when using the
	// URI specified by the HttpPushUri property to push a software image. These targets should correspond to software
	// inventory instances or their related items. If this property is not present or contains no targets, the service
	// shall apply the software image to all applicable targets, as determined by the service. If the target specifies
	// a device resource, the software image file shall be applied to the specified device. If the target specifies a
	// resource collection, the software image shall be applied to each applicable member of the specified collection.
	// If the target resource specifies an Aggregate resource, the software image file shall be applied to each
	// applicable element of the specified aggregate. If the target resource specifies a ComputerSystem resource, the
	// software image file shall be applied to the applicable components within the specified computer system.
	HTTPPushURITargets []string
	// HTTPPushURITargetsBusy shall indicate whether any client has reserved the HttpPushUriTargets property for
	// firmware updates. A client should set this property to 'true' when it uses HttpPushUriTargets for firmware
	// updates. A client should set it to 'false' when it no longer uses HttpPushUriTargets for updates. The property
	// can provide multiple clients a way to negotiate ownership of HttpPushUriTargets and helps clients determine
	// whether another client is using HttpPushUriTargets to make firmware updates. This property has no functional
	// requirements for the service.
	HTTPPushURITargetsBusy bool
	// MaxImageSizeBytes shall indicate the maximum size of the software update image that clients can send to this
	// update service.
	MaxImageSizeBytes int
	// MultipartHTTPPushURI endpoint is used to perform a multipart push (POST) updates
	MultipartHTTPPushURI string `json:"MultiPartHttpPushUri"`
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PublicIdentitySSHKey shall contain a link to a resource of type Key that represents the public key that is used
	// with the SimpleUpdate action for the key-based authentication. This property shall not be present if a key-pair
	// is not available.
	publicIdentitySSHKey string
	// RemoteServerCertificates shall contain a link to a resource collection of type CertificateCollection that
	// represents the server certificates for the server referenced by the ImageURI property in SimpleUpdate. If
	// VerifyRemoteServerCertificate is 'true', services shall compare the certificates in this collection with the
	// certificate obtained during handshaking with the image server in order to verify the identity of the image
	// server prior to transferring the image. If the server cannot be verified, the service shall not send the
	// transfer request. If VerifyRemoteServerCertificate is 'false', the service shall not perform certificate
	// verification with certificates in this collection. Regardless of the contents of this collection, services may
	// perform additional verification based on other factors, such as the configuration of the SecurityPolicy
	// resource.
	remoteServerCertificates []string
	// RemoteServerSSHKeys shall contain a link to a resource collection of type KeyCollection that represents the
	// server SSH keys for the server referenced by the ImageURI property in SimpleUpdate. If VerifyRemoteServerSSHKey
	// is 'true', services shall compare the keys in this collection with the key obtained during handshaking with the
	// image server in order to verify the identity of the image server prior to transferring the image. If the server
	// cannot be verified, the service shall not send the transfer request. If VerifyRemoteServerSSHKey is 'false', the
	// service shall not perform key verification with keys in this collection.
	remoteServerSSHKeys []string
	// ServiceEnabled indicates whether this service isenabled.
	ServiceEnabled bool
	// SoftwareInventory points towards the firmware store endpoint
	softwareInventory string
	// Status describes the status and health of a resource and its children.
	Status common.Status
	// SupportedUpdateImageFormats shall contain the image format types supported by the service.
	SupportedUpdateImageFormats []SupportedUpdateImageFormatType
	// VerifyRemoteServerCertificate shall indicate whether the service will verify the certificate of the server
	// referenced by the ImageURI property in SimpleUpdate prior to sending the transfer request with the certificates
	// found in the collection referenced by the RemoteServerCertificates property. If this property is not supported
	// by the service, it shall be assumed to be 'false'. This property should default to 'false' in order to maintain
	// compatibility with older clients. Regardless of the value of this property, services may perform additional
	// verification based on other factors, such as the configuration of the SecurityPolicy resource.
	VerifyRemoteServerCertificate bool
	// VerifyRemoteServerSSHKey shall indicate whether the service will verify the SSH key of the server referenced by
	// the ImageURI property in SimpleUpdate prior to sending the transfer request with the keys found in the
	// collection referenced by the RemoteServerSSHKeys property. If this property is not supported by the service, it
	// shall be assumed to be 'false'. This property should default to 'false' in order to maintain compatibility with
	// older clients.
	VerifyRemoteServerSSHKey bool

	// TransferProtocol is an optional list of network protocols used by the UpdateService
	// to retrieve the software image file.
	TransferProtocol                 []string
	generateSSHIdentityKeyPairTarget string
	removeSSHIdentityKeyPairTarget   string
	simpleUpdateTarget               string
	startUpdateTarget                string

	// OemActions contains all the vendor specific actions. It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a UpdateService object from the raw JSON.
func (updateService *UpdateService) UnmarshalJSON(b []byte) error {
	type temp UpdateService
	type actions struct {
		GenerateSSHIdentityKeyPair common.ActionTarget `json:"#UpdateService.GenerateSSHIdentityKeyPair"`
		RemoveSSHIdentityKeyPair   common.ActionTarget `json:"#UpdateService.RemoveSSHIdentityKeyPair"`
		SimpleUpdate               struct {
			AllowableValues []string `json:"TransferProtocol@Redfish.AllowableValues"`
			Target          string
		} `json:"#UpdateService.SimpleUpdate"`
		StartUpdate common.ActionTarget `json:"#UpdateService.StartUpdate"`

		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		Actions                  actions
		ClientCertificates       common.LinksCollection
		FirmwareInventory        common.Link
		PublicIdentitySSHKey     common.Link
		RemoteServerCertificates common.LinksCollection
		RemoteServerSSHKeys      common.LinksCollection
		SoftwareInventory        common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*updateService = UpdateService(t.temp)
	updateService.clientCertificates = t.ClientCertificates.ToStrings()
	updateService.firmwareInventory = t.FirmwareInventory.String()
	updateService.publicIdentitySSHKey = t.PublicIdentitySSHKey.String()
	updateService.remoteServerCertificates = t.RemoteServerCertificates.ToStrings()
	updateService.remoteServerSSHKeys = t.RemoteServerSSHKeys.ToStrings()
	updateService.softwareInventory = t.SoftwareInventory.String()

	updateService.generateSSHIdentityKeyPairTarget = t.Actions.GenerateSSHIdentityKeyPair.Target
	updateService.removeSSHIdentityKeyPairTarget = t.Actions.RemoveSSHIdentityKeyPair.Target
	updateService.startUpdateTarget = t.Actions.StartUpdate.Target
	updateService.simpleUpdateTarget = t.Actions.SimpleUpdate.Target
	updateService.TransferProtocol = t.Actions.SimpleUpdate.AllowableValues
	updateService.OemActions = t.Actions.Oem

	updateService.RawData = b

	return nil
}

// GenerateSSHIdentityKeyPair generates a new SSH identity key-pair to be used with the
// UpdateService resource. The generated public key is stored in the Key resource
// referenced by the PublicIdentitySSHKey property. Any existing key-pair is deleted
// and replaced by the new key-pair.
func (updateService *UpdateService) GenerateSSHIdentityKeyPair(curve ECDSACurveType, keyLength int, keyType SSHAlgoKeyType) error {
	t := struct {
		Curve     ECDSACurveType
		KeyLength int
		KeyType   SSHAlgoKeyType
	}{
		Curve:     curve,
		KeyLength: keyLength,
		KeyType:   keyType,
	}
	return updateService.Post(updateService.generateSSHIdentityKeyPairTarget, t)
}

// RemoveSSHIdentityKeyPair removes the SSH identity key-pair used with the UpdateService resource.
func (updateService *UpdateService) RemoveSSHIdentityKeyPair() error {
	return updateService.Post(updateService.generateSSHIdentityKeyPairTarget, nil)
}

// SimpleUpdateParameters contains the parameters for the SimpleUpdate action.
type SimpleUpdateParameters struct {
	// ForceUpdate is an indication of whether the service should bypass update policies when
	// applying the provided image. The default is `false`."
	ForceUpdate bool `json:",omitempty"`
	// ImageURI is the URI of the software image to install.
	ImageURI string
	// Password to access the URI specified by the ImageURI parameter.
	Passord string `json:",omitempty"`
	// Targets shall contain zero or more URIs that indicate where to apply the update image.
	// These targets should correspond to software inventory instances or their related items.
	// If this parameter is not present or contains no targets, the service shall apply the
	// software image to all applicable targets, as determined by the service.
	// If the target specifies a device resource, the software image file shall be applied to
	// the specified device. If the target specifies a resource collection, the software image
	// shall be applied to each applicable member of the specified collection.
	// If the target resource specifies an Aggregate resource, the software image file shall be
	// applied to each applicable element of the specified aggregate. If the target resource
	// specifies a ComputerSystem resource, the software image file shall be applied to the
	// applicable components within the specified computer system.
	Targets []string `json:",omitempty"`
	// TransferProtocol shall contain the network protocol that the update service shall use to
	// retrieve the software image located at the ImageURI. Services should ignore this
	// parameter if the URI provided in ImageURI contains a scheme. If this parameter is not
	// provided (or supported), and a transfer protocol is not specified by a scheme contained
	// within this URI, the service shall use HTTP to retrieve the image.
	TransferProtocol TransferProtocolType `json:",omitempty"`
	// Username shall represent the user name to access the URI specified by the ImageURI parameter.
	Username string `json:",omitempty"`
}

// SimpleUpdate will update installed software components using a software image file
// located at an ImageURI parameter-specified URI.
func (updateService *UpdateService) SimpleUpdate(parameters *SimpleUpdateParameters) error {
	return updateService.Post(updateService.simpleUpdateTarget, parameters)
}

// StartUpdate starts updating all images that have been previously invoked using an
// OperationApplyTime value of `OnStartUpdateRequest`.
func (updateService *UpdateService) StartUpdate() error {
	return updateService.Post(updateService.startUpdateTarget, nil)
}

// FirmwareInventories gets the collection of firmware inventories of this update service
func (updateService *UpdateService) FirmwareInventories() ([]*SoftwareInventory, error) {
	return ListReferencedSoftwareInventories(updateService.GetClient(), updateService.firmwareInventory)
}

// PublicIdentitySSHKey get the public key that is used with the SimpleUpdate action
// for the key-based authentication. The GenerateSSHIdentityKeyPair and RemoveSSHIdentityKeyPair
// are used to update the key for the SimpleUpdate action.
// This property shall not be present if a key-pair is not available.
func (updateService *UpdateService) PublicIdentitySSHKey() (*Key, error) {
	if updateService.publicIdentitySSHKey == "" {
		return nil, nil
	}
	return GetKey(updateService.GetClient(), updateService.publicIdentitySSHKey)
}

// RemoteServerCertificates gets the server certificates for the server referenced by the
// ImageURI property in SimpleUpdate.  If VerifyRemoteServerCertificate is `true`, services
// shall compare the certificates in this collection with the certificate obtained during
// handshaking with the image server in order to verify the identity of the image server
// prior to transferring the image.  If the server cannot be verified, the service shall not
// send the transfer request. If VerifyRemoteServerCertificate is `false`, the service shall
// not perform certificate verification with certificates in this collection. Regardless of
// the contents of this collection, services may perform additional verification based on
// other factors, such as the configuration of the SecurityPolicy resource.
func (updateService *UpdateService) RemoteServerCertificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](updateService.GetClient(), updateService.remoteServerCertificates)
}

// RemoteServerSSHKeys gets the server SSH keys for the server referenced by the ImageURI
// property in SimpleUpdate. If VerifyRemoteServerSSHKey is `true`, services shall compare
// the keys in this collection with the key obtained during handshaking with the image
// server in order to verify the identity of the image server prior to transferring the
// image. If the server cannot be verified, the service shall not send the transfer request.
// If VerifyRemoteServerSSHKey is `false`, the service shall not perform key verification
// with keys in this collection.
func (updateService *UpdateService) RemoteServerSSHKeys() ([]*Key, error) {
	return common.GetObjects[Key](updateService.GetClient(), updateService.remoteServerSSHKeys)
}

// SoftwareInventories gets the collection of software inventories of this update service
func (updateService *UpdateService) SoftwareInventories() ([]*SoftwareInventory, error) {
	return ListReferencedSoftwareInventories(updateService.GetClient(), updateService.softwareInventory)
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(c common.Client, uri string) (*UpdateService, error) {
	return common.GetObject[UpdateService](c, uri)
}
