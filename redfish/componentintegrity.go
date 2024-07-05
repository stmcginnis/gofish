//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ComponentIntegrityType string

const (
	// SPDMComponentIntegrityType shall indicate the integrity information is obtained through the Security Protocol
	// and Data Model (SPDM) protocol as defined in DMTF DSP0274.
	SPDMComponentIntegrityType ComponentIntegrityType = "SPDM"
	// TPMComponentIntegrityType shall indicate the integrity information is related to a Trusted Platform Module (TPM)
	// as defined by the Trusted Computing Group (TCG).
	TPMComponentIntegrityType ComponentIntegrityType = "TPM"
	// OEMComponentIntegrityType shall indicate the integrity information is OEM-specific and the OEM section may
	// include additional information.
	OEMComponentIntegrityType ComponentIntegrityType = "OEM"
)

type DMTFmeasurementTypes string

const (
	// ImmutableROMDMTFmeasurementTypes Immutable ROM.
	ImmutableROMDMTFmeasurementTypes DMTFmeasurementTypes = "ImmutableROM"
	// MutableFirmwareDMTFmeasurementTypes Mutable firmware or any mutable code.
	MutableFirmwareDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmware"
	// HardwareConfigurationDMTFmeasurementTypes Hardware configuration, such as straps.
	HardwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "HardwareConfiguration"
	// FirmwareConfigurationDMTFmeasurementTypes Firmware configuration, such as configurable firmware policy.
	FirmwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "FirmwareConfiguration"
	// MutableFirmwareVersionDMTFmeasurementTypes Mutable firmware version.
	MutableFirmwareVersionDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareVersion"
	// MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes Mutable firmware security version number.
	MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareSecurityVersionNumber"
	// MeasurementManifestDMTFmeasurementTypes Measurement Manifest.
	MeasurementManifestDMTFmeasurementTypes DMTFmeasurementTypes = "MeasurementManifest"
)

type MeasurementSpecification string

const (
	// DMTFMeasurementSpecification shall indicate the measurement specification is defined by DMTF in DSP0274.
	DMTFMeasurementSpecification MeasurementSpecification = "DMTF"
)

type SPDMmeasurementSummaryType string

const (
	// TCBSPDMmeasurementSummaryType The measurement summary covers the TCB.
	TCBSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "TCB"
	// AllSPDMmeasurementSummaryType The measurement summary covers all measurements in SPDM.
	AllSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "All"
)

type SecureSessionType string

const (
	// PlainSecureSessionType A plain text session without any protection.
	PlainSecureSessionType SecureSessionType = "Plain"
	// EncryptedAuthenticatedSecureSessionType An established session where both encryption and authentication are
	// protecting the communication.
	EncryptedAuthenticatedSecureSessionType SecureSessionType = "EncryptedAuthenticated"
	// AuthenticatedOnlySecureSessionType An established session where only authentication is protecting the
	// communication.
	AuthenticatedOnlySecureSessionType SecureSessionType = "AuthenticatedOnly"
)

type VerificationStatus string

const (
	// SuccessVerificationStatus Successful verification.
	SuccessVerificationStatus VerificationStatus = "Success"
	// FailedVerificationStatus Unsuccessful verification.
	FailedVerificationStatus VerificationStatus = "Failed"
)

// CommonAuthInfo shall contain common identity-related authentication information.
type CommonAuthInfo struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identity of the
	// component referenced by the TargetComponentURI property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a CommonAuthInfo object from the raw JSON.
func (commonauthinfo *CommonAuthInfo) UnmarshalJSON(b []byte) error {
	type temp CommonAuthInfo
	var t struct {
		temp
		ComponentCertificate common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*commonauthinfo = CommonAuthInfo(t.temp)

	// Extract the links to other entities for later
	commonauthinfo.componentCertificate = t.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the identity of the component.
func (commonauthinfo *CommonAuthInfo) ComponentCertificate(c common.Client) (*Certificate, error) {
	if commonauthinfo.componentCertificate == "" {
		return nil, nil
	}

	return GetCertificate(c, commonauthinfo.componentCertificate)
}

// CommunicationInfo shall contain information about communication between two components.
type CommunicationInfo struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components. The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// ComponentIntegrity shall represent critical and pertinent security information about a specific device, system,
// software element, or other managed entity.
type ComponentIntegrity struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ComponentIntegrityEnabled shall indicate whether security protocols are enabled for the component. If
	// ComponentIntegrityType contains 'SPDM', a value of 'false' shall prohibit the SPDM Requester from using SPDM to
	// communicate with the component identified by the TargetComponentURI property. If ComponentIntegrityType contains
	// 'TPM', a value of 'false' shall disable the TPM component identified by the TargetComponentURI property
	// entirely. If 'false', services shall not provide the TPM and SPDM properties in response payloads for this
	// resource. If 'false', services shall reject action requests to this resource. If 'true', services shall allow
	// security protocols with the component identified by the TargetComponentURI property.
	ComponentIntegrityEnabled bool
	// ComponentIntegrityType shall contain the underlying security technology providing integrity information for the
	// component.
	ComponentIntegrityType ComponentIntegrityType
	// ComponentIntegrityTypeVersion shall contain the version of the security technology indicated by the
	// ComponentIntegrityType property. If the service has not established secure communication with the device or if
	// security protocols are disabled, this property shall contain an empty string. If ComponentIntegrityType contains
	// 'SPDM', this property shall contain the negotiated or selected SPDM protocol and shall follow the regular
	// expression pattern '^\d+\.\d+\.\d+$'. If ComponentIntegrityType contains 'TPM', this property shall contain the
	// version of the TPM.
	ComponentIntegrityTypeVersion string
	// Description provides a description of this resource.
	Description string
	// LastUpdated shall contain the date and time when information for the component was last updated.
	LastUpdated string
	// SPDM shall contain integrity information about the SPDM Responder identified by the TargetComponentURI property
	// as reported by an SPDM Requester. This property shall be present if ComponentIntegrityType contains 'SPDM' and
	// if 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property shall be absent.
	SPDM SPDMinfo
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TPM shall contain integrity information about the Trusted Platform Module (TPM) identified by the
	// TargetComponentURI property. This property shall be present if ComponentIntegrityType contains 'TPM' and if
	// 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property shall be absent.
	TPM TPMinfo
	// TargetComponentURI shall contain a link to the resource whose integrity information is reported in this
	// resource. If ComponentIntegrityType contains 'SPDM', this property shall contain a URI to the resource that
	// represents the SPDM Responder. If ComponentIntegrityType contains 'TPM', this property shall contain a URI with
	// RFC6901-defined JSON fragment notation to a member of the TrustedModules array in a ComputerSystem resource that
	// represents the TPM or a resource of type TrustedComponent that represents the TPM.
	TargetComponentURI string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData                         []byte
	spdmGetSignedMeasurementsTarget string
	tpmGetSignedMeasurementsTarget  string
	componentsProtected             []string
	// ComponentsProtectedCount is the number of resources protected by the component identified by TargetComponentURI.
	ComponentsProtectedCount int
}

// UnmarshalJSON unmarshals a ComponentIntegrity object from the raw JSON.
func (componentintegrity *ComponentIntegrity) UnmarshalJSON(b []byte) error {
	type temp ComponentIntegrity
	type Actions struct {
		SPDMGetSignedMeasurements common.ActionTarget `json:"#ComponentIntegrity.SPDMGetSignedMeasurements"`
		TPMGetSignedMeasurements  common.ActionTarget `json:"#ComponentIntegrity.TPMGetSignedMeasurements"`
	}
	type Links struct {
		// ComponentsProtected shall contain an array of links to resources that the component identified by the
		// TargetComponentURI property provides integrity protection. This property shall not contain the value of the
		// TargetComponentURI property.
		ComponentsProtected common.Links
		// ComponentsProtected@odata.count
		ComponentsProtectedCount int `json:"ComponentsProtected@odata.count"`
	}
	var t struct {
		temp
		Actions Actions
		Links   Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*componentintegrity = ComponentIntegrity(t.temp)

	// Extract the links to other entities for later
	componentintegrity.spdmGetSignedMeasurementsTarget = t.Actions.SPDMGetSignedMeasurements.Target
	componentintegrity.tpmGetSignedMeasurementsTarget = t.Actions.TPMGetSignedMeasurements.Target
	componentintegrity.componentsProtected = t.Links.ComponentsProtected.ToStrings()
	componentintegrity.ComponentsProtectedCount = t.Links.ComponentsProtectedCount

	// This is a read/write object, so we need to save the raw object data for later
	componentintegrity.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (componentintegrity *ComponentIntegrity) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ComponentIntegrity)
	original.UnmarshalJSON(componentintegrity.rawData)

	readWriteFields := []string{
		"ComponentIntegrityEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(componentintegrity).Elem()

	return componentintegrity.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetComponentIntegrity will get a ComponentIntegrity instance from the service.
func GetComponentIntegrity(c common.Client, uri string) (*ComponentIntegrity, error) {
	return common.GetObject[ComponentIntegrity](c, uri)
}

// ListReferencedComponentIntegritys gets the collection of ComponentIntegrity from
// a provided reference.
func ListReferencedComponentIntegritys(c common.Client, link string) ([]*ComponentIntegrity, error) {
	return common.GetCollectionObjects[ComponentIntegrity](c, link)
}

// SPDMGetSignedMeasurementsRequest contains the parameters for the SPDMGetSignedMeasurements action.
type SPDMGetSignedMeasurementsRequest struct {
	// MeasurementIndices is an array of indices that identify the measurement blocks to sign.
	MeasurementIndices []int
	// Nonce is a 32-byte hex-encoded string that is signed with the measurements. The value should be unique.
	Nonce string
	// SlotID is the slot identifier for the certificate containing the private key to generate the signature over the measurements.
	SlotID int `json:"SlotId"`
}

// SPDMGetSignedMeasurementsResponse shall contain the SPDM signed measurements from an SPDM Responder.
type SPDMGetSignedMeasurementsResponse struct {
	// Certificate shall contain a link to a resource of type Certificate that represents the certificate corresponding
	// to the SPDM slot identifier that can be used to validate the signature. This property shall not be present if
	// the SlotId parameter contains the value '15'.
	certificate string
	// HashingAlgorithm shall contain the hashing algorithm negotiated between the SPDM Requester and the SPDM
	// Responder. The allowable values for this property shall be the hash algorithm names found in the 'BaseHashAlgo'
	// field of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this
	// property shall contain the value 'OEM'.
	HashingAlgorithm string
	// PublicKey shall contain a Privacy Enhanced Mail (PEM)-encoded public key, as defined in section 13 of RFC7468,
	// that can be used to validate the signature. This property shall only be present when the SPDM Requester was pre-
	// provisioned with the SPDM Responder's public key and the SlotId parameter contains the value '15'.
	PublicKey string
	// SignedMeasurements shall contain the cryptographic signed statement over the given nonce and measurement blocks
	// corresponding to the requested measurement indices. If the SPDM version is 1.2, this value shall be a
	// concatenation of SPDM 'VCA' and 'GET_MEASUREMENTS' requests and responses exchanged between the SPDM Requester
	// and the SPDM Responder. If SPDM version is 1.0 or 1.1, this value shall be a concatenation of SPDM
	// 'GET_MEASUREMENTS' requests and responses exchanged between the SPDM Requester and the SPDM Responder. The last
	// 'MEASUREMENTS' response shall contain a signature generated over the 'L2' string by the SPDM Responder.
	SignedMeasurements string
	// SigningAlgorithm shall contain the asymmetric signing algorithm negotiated between the SPDM Requester and the
	// SPDM Responder. The allowable values for this property shall be the asymmetric key signature algorithm names
	// found in the 'BaseAsymAlgo' field of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is
	// an extended algorithm, this property shall contain the value 'OEM'.
	SigningAlgorithm string
	// Version shall contain the SPDM version negotiated between the SPDM Requester and the SPDM Responder to generate
	// the cryptographic signed statement. For example, '1.0', '1.1', or '1.2'.
	Version string
	client  common.Client
}

// UnmarshalJSON unmarshals a SPDMGetSignedMeasurementsResponse object from the raw JSON.
func (spdmgetsignedmeasurementsresponse *SPDMGetSignedMeasurementsResponse) UnmarshalJSON(b []byte) error {
	type temp SPDMGetSignedMeasurementsResponse
	var t struct {
		temp
		Certificate common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmgetsignedmeasurementsresponse = SPDMGetSignedMeasurementsResponse(t.temp)

	// Extract the links to other entities for later
	spdmgetsignedmeasurementsresponse.certificate = t.Certificate.String()

	return nil
}

// Certificate gets the certificate corresponding to the SPDM slot identifier that can be used to validate the signature.
func (spdmgetsignedmeasurementsresponse *SPDMGetSignedMeasurementsResponse) Certificate() (*Certificate, error) {
	if spdmgetsignedmeasurementsresponse.certificate == "" {
		return nil, nil
	}

	return GetCertificate(spdmgetsignedmeasurementsresponse.client, spdmgetsignedmeasurementsresponse.certificate)
}

func (spdmgetsignedmeasurementsresponse *SPDMGetSignedMeasurementsResponse) setClient(c common.Client) {
	spdmgetsignedmeasurementsresponse.client = c
}

// SPDMGetSignedMeasurements generates an SPDM cryptographic signed statement over the given nonce and measurements of the SPDM Responder.
func (componentintegrity *ComponentIntegrity) SPDMGetSignedMeasurements(request *SPDMGetSignedMeasurementsRequest) (*SPDMGetSignedMeasurementsResponse, error) {
	resp, err := componentintegrity.PostWithResponse(componentintegrity.spdmGetSignedMeasurementsTarget, request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response SPDMGetSignedMeasurementsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	response.setClient(componentintegrity.GetClient())
	return &response, nil
}

// TPMGetSignedMeasurements generates a TPM cryptographic signed statement over the given nonce and PCRs of the TPM for TPM 2.0 devices.
func (componentintegrity *ComponentIntegrity) TPMGetSignedMeasurements(request *TPMGetSignedMeasurementsRequest) (*TPMGetSignedMeasurementsResponse, error) {
	resp, err := componentintegrity.PostWithResponse(componentintegrity.tpmGetSignedMeasurementsTarget, request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response TPMGetSignedMeasurementsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SPDMcommunication shall contain information about communication between two components.
type SPDMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components. The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// SPDMidentity shall contain identity authentication information about the SPDM Requester and SPDM Responder.
type SPDMidentity struct {
	// RequesterAuthentication shall contain authentication information of the identity of the SPDM Requester.
	RequesterAuthentication SPDMrequesterAuth
	// ResponderAuthentication shall contain authentication information of the identity of the SPDM Responder.
	ResponderAuthentication SPDMresponderAuth
}

// SPDMinfo shall contain integrity information about an SPDM Responder as reported by an SPDM Requester.
type SPDMinfo struct {
	// ComponentCommunication shall contain information about communication between the SPDM Requester and SPDM
	// Responder.
	ComponentCommunication SPDMcommunication
	// IdentityAuthentication shall contain identity authentication information about the SPDM Requester and SPDM
	// Responder.
	IdentityAuthentication SPDMidentity
	// MeasurementSet shall contain measurement information for the SPDM Responder.
	MeasurementSet SPDMmeasurementSet
	// Requester shall contain a link to the resource representing the SPDM Responder that is reporting the integrity
	// of the SPDM Responder identified by the TargetComponentURI property.
	requester string
}

// UnmarshalJSON unmarshals a SPDMinfo object from the raw JSON.
func (spdminfo *SPDMinfo) UnmarshalJSON(b []byte) error {
	type temp SPDMinfo
	var t struct {
		temp
		Requestor common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdminfo = SPDMinfo(t.temp)

	// Extract the links to other entities for later
	spdminfo.requester = t.Requestor.String()

	return nil
}

// SPDMmeasurementSet shall contain SPDM Responder measurement information.
type SPDMmeasurementSet struct {
	// MeasurementSpecification shall contain the measurement specification negotiated between the SPDM Requester and
	// SPDM Responder.
	MeasurementSpecification MeasurementSpecification
	// MeasurementSummary shall contain the Base64-encoded measurement summary using the hash algorithm indicated by
	// the MeasurementSummaryHashAlgorithm property.
	MeasurementSummary string
	// MeasurementSummaryHashAlgorithm shall contain the hash algorithm used to compute the measurement summary. The
	// allowable values for this property shall be the hash algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this property
	// shall contain the value 'OEM'.
	MeasurementSummaryHashAlgorithm string
	// MeasurementSummaryType shall contain the type of measurement summary.
	MeasurementSummaryType SPDMmeasurementSummaryType
	// Measurements shall contain measurements from an SPDM Responder.
	Measurements []SPDMsingleMeasurement
}

// SPDMrequesterAuth shall contain authentication information of the identity of the SPDM Requester.
type SPDMrequesterAuth struct {
	// ProvidedCertificate shall contain a link to a resource of type Certificate that represents the identity of the
	// SPDM Requester provided in mutual authentication.
	providedCertificate string
}

// UnmarshalJSON unmarshals a SPDMrequesterAuth object from the raw JSON.
func (spdmrequesterauth *SPDMrequesterAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMrequesterAuth
	var t struct {
		temp
		ProvidedCertificate common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmrequesterauth = SPDMrequesterAuth(t.temp)

	// Extract the links to other entities for later
	spdmrequesterauth.providedCertificate = t.ProvidedCertificate.String()

	return nil
}

// ProvidedCertificate gets the identity of the SPDM Requester provided in mutual authentication.
func (spdmrequesterauth *SPDMrequesterAuth) ProvidedCertificate(c common.Client) (*Certificate, error) {
	if spdmrequesterauth.providedCertificate == "" {
		return nil, nil
	}

	return GetCertificate(c, spdmrequesterauth.providedCertificate)
}

// SPDMresponderAuth shall contain common identity-related authentication information.
type SPDMresponderAuth struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identity of the
	// component referenced by the TargetComponentURI property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a SPDMresponderAuth object from the raw JSON.
func (spdmresponderauth *SPDMresponderAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMresponderAuth
	var t struct {
		temp
		ComponentCertificate common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmresponderauth = SPDMresponderAuth(t.temp)

	// Extract the links to other entities for later
	spdmresponderauth.componentCertificate = t.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the identity of the component referenced by the TargetComponentURI property.
func (spdmresponderauth *SPDMresponderAuth) ComponentCertificate(c common.Client) (*Certificate, error) {
	if spdmresponderauth.componentCertificate == "" {
		return nil, nil
	}

	return GetCertificate(c, spdmresponderauth.componentCertificate)
}

// SPDMsingleMeasurement shall contain a single SPDM measurement for an SPDM Responder.
type SPDMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the measurement was last updated.
	LastUpdated string
	// Measurement shall contain the Base64-encoded measurement using the hash algorithm indicated by the
	// MeasurementHashAlgorithm property. This property shall not contain a raw bit stream as a measurement. If the
	// SPDM Responder provides a raw bit stream, the SPDM Requester may apply a hash algorithm to the raw bit stream in
	// order to report the measurement.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute the measurement. The allowable values
	// for this property shall be the hash algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this property
	// shall contain the value 'OEM'. This property shall not be present if MeasurementSpecification does not contain
	// 'DMTF'.
	MeasurementHashAlgorithm string
	// MeasurementIndex shall contain the index of the measurement.
	MeasurementIndex int
	// MeasurementType shall contain the type or characteristics of the data that this measurement represents. This
	// property shall not be present if MeasurementSpecification does not contain 'DMTF'.
	MeasurementType DMTFmeasurementTypes
	// PartofSummaryHash shall indicate if this measurement is part of the measurement summary in the
	// MeasurementSummary property. If this property is not present, it shall be assumed to be 'false'.
	PartofSummaryHash bool
	// SecurityVersionNumber shall contain an 8-byte hex-encoded string of the security version number the measurement
	// represents. This property shall only be present if MeasurementType contains the value
	// 'MutableFirmwareSecurityVersionNumber'.
	SecurityVersionNumber string
}

// SingleSessionInfo shall contain information about a single communication channel or session between two
// components.
type SingleSessionInfo struct {
	// SessionID shall contain the unique identifier for the active session or communication channel between two
	// components.
	SessionID int `json:"SessionId"`
	// SessionType shall contain the type of session or communication channel between two components.
	SessionType SecureSessionType
}

// TPMGetSignedMeasurementsRequest contains the properties to pass in to the TPMGetSignedMeasurements action.
type TPMGetSignedMeasurementsRequest struct {
	// CertificateODataID is the ODataID for the certificate that represents the TPM attestation key.
	CertificateODataID string `json:"Certificate"`
	// Nonce is a set of bytes as a hex-encoded string that is signed with the measurements. The value should be unique.
	Nonce string `json:",omitempty"`
	// PCRSelection is the Base64-encoded representation of the 'TPML_PCR_SELECTION' object, as defined by the
	// Trusted Platform Module Library Specification, that identifies the PCRs to sign. The service shall send this
	// value to the TPM in the 'PCRselect' parameter of the 'TPM2_Quote' command defined in the Trusted Platform Module
	// Library Specification.
	PCRSelection string
	// Scheme is the Base64-encoded representation of the 'TPMT_SIG_SCHEME' object, as defined in the Trusted Platform
	// Module Library Specification, that identifies the signing scheme to use for the TPM attestation key. The service
	// shall send this value to the TPM in the 'inScheme' parameter of the 'TPM2_Quote' command defined in the Trusted
	// Platform Module Library Specification.
	Scheme string
}

// TPMGetSignedMeasurementsResponse shall contain the TPM signed PCR measurements from a TPM.
type TPMGetSignedMeasurementsResponse struct {
	// SignedMeasurements shall contain a Base64-encoded cryptographic signed statement generated by the signer. This
	// value shall be the concatenation of the 'quoted' and 'signature' response values of the 'TPM2_Quote' command
	// defined in the Trusted Platform Module Library Specification.
	SignedMeasurements string
}

// TPMauth shall contain common identity-related authentication information.
type TPMauth struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identity of the
	// component referenced by the TargetComponentURI property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a TPMauth object from the raw JSON.
func (tpmauth *TPMauth) UnmarshalJSON(b []byte) error {
	type temp TPMauth
	var t struct {
		temp
		ComponentCertificate common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmauth = TPMauth(t.temp)

	// Extract the links to other entities for later
	tpmauth.componentCertificate = t.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the identity of the component referenced by the TargetComponentURI property.
func (tpmauth *TPMauth) ComponentCertificate(c common.Client) (*Certificate, error) {
	if tpmauth.componentCertificate == "" {
		return nil, nil
	}

	return GetCertificate(c, tpmauth.componentCertificate)
}

// TPMcommunication shall contain information about communication between two components.
type TPMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components. The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// TPMinfo shall contain integrity information about a Trusted Platform Module (TPM).
type TPMinfo struct {
	// ComponentCommunication shall contain information about communication with the TPM.
	ComponentCommunication TPMcommunication
	// IdentityAuthentication shall contain identity authentication information about the TPM.
	IdentityAuthentication TPMauth
	// MeasurementSet shall contain measurement information from the TPM.
	MeasurementSet TPMmeasurementSet
	// NonceSizeBytesMaximum shall contain the maximum number of bytes that can be specified in the Nonce parameter of
	// the TPMGetSignedMeasurements action.
	NonceSizeBytesMaximum int
}

// TPMmeasurementSet shall contain Trusted Computing Group TPM measurement information.
type TPMmeasurementSet struct {
	// Measurements shall contain measurements from a TPM.
	Measurements []TPMsingleMeasurement
}

// TPMsingleMeasurement shall contain a single Trusted Computing Group TPM measurement.
type TPMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the measurement was last updated.
	LastUpdated string
	// Measurement shall contain the Base64-encoded PCR digest using the hashing algorithm indicated by
	// MeasurementHashAlgorithm property.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute the measurement. The allowable values
	// for this property shall be the strings in the 'Algorithm Name' field of the 'TPM_ALG_ID Constants' table within
	// the 'Trusted Computing Group Algorithm Registry'.
	MeasurementHashAlgorithm string
	// PCR shall contain the Platform Configuration Register (PCR) bank of the measurement.
	PCR int
}
