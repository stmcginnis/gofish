//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/AccelerationFunction.v1_0_5.json
// 2018.3 - #AccelerationFunction.v1_0_5.AccelerationFunction

package schemas

import (
	"encoding/json"
)

type AccelerationFunctionType string

const (
	// EncryptionAccelerationFunctionType is an encryption function.
	EncryptionAccelerationFunctionType AccelerationFunctionType = "Encryption"
	// CompressionAccelerationFunctionType is a compression function.
	CompressionAccelerationFunctionType AccelerationFunctionType = "Compression"
	// PacketInspectionAccelerationFunctionType is a packet inspection function.
	PacketInspectionAccelerationFunctionType AccelerationFunctionType = "PacketInspection"
	// PacketSwitchAccelerationFunctionType is a packet switch function.
	PacketSwitchAccelerationFunctionType AccelerationFunctionType = "PacketSwitch"
	// SchedulerAccelerationFunctionType is a scheduler function.
	SchedulerAccelerationFunctionType AccelerationFunctionType = "Scheduler"
	// AudioProcessingAccelerationFunctionType is an audio processing function.
	AudioProcessingAccelerationFunctionType AccelerationFunctionType = "AudioProcessing"
	// VideoProcessingAccelerationFunctionType is a video processing function.
	VideoProcessingAccelerationFunctionType AccelerationFunctionType = "VideoProcessing"
	// OEMAccelerationFunctionType is an OEM-defined acceleration function.
	OEMAccelerationFunctionType AccelerationFunctionType = "OEM"
)

// AccelerationFunction shall represent the acceleration function that a
// processor implements in a Redfish implementation. This can include functions
// such as audio processing, compression, encryption, packet inspection, packet
// switching, scheduling, or video processing.
type AccelerationFunction struct {
	Entity
	// AccelerationFunctionType shall contain the string that identifies the
	// acceleration function type.
	AccelerationFunctionType AccelerationFunctionType
	// FPGAReconfigurationSlots shall contain an array of the FPGA reconfiguration
	// slot identifiers that this acceleration function occupies.
	FPGAReconfigurationSlots []string `json:"FpgaReconfigurationSlots"`
	// Manufacturer shall contain a string that identifies the manufacturer of the
	// acceleration function.
	Manufacturer string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the total acceleration function power consumption,
	// in watt units.
	PowerWatts int
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UUID shall contain a UUID for the acceleration function. RFC4122 describes
	// methods that can create the value. The value should be considered to be
	// opaque. Client software should only treat the overall value as a UUID and
	// should not interpret any subfields within the UUID.
	UUID string
	// Version shall describe the acceleration function version.
	Version string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
}

// UnmarshalJSON unmarshals a AccelerationFunction object from the raw JSON.
func (a *AccelerationFunction) UnmarshalJSON(b []byte) error {
	type temp AccelerationFunction
	type aLinks struct {
		Endpoints     Links `json:"Endpoints"`
		PCIeFunctions Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Links aLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AccelerationFunction(tmp.temp)

	// Extract the links to other entities for later
	a.endpoints = tmp.Links.Endpoints.ToStrings()
	a.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()

	return nil
}

// GetAccelerationFunction will get a AccelerationFunction instance from the service.
func GetAccelerationFunction(c Client, uri string) (*AccelerationFunction, error) {
	return GetObject[AccelerationFunction](c, uri)
}

// ListReferencedAccelerationFunctions gets the collection of AccelerationFunction from
// a provided reference.
func ListReferencedAccelerationFunctions(c Client, link string) ([]*AccelerationFunction, error) {
	return GetCollectionObjects[AccelerationFunction](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (a *AccelerationFunction) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](a.client, a.endpoints)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (a *AccelerationFunction) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](a.client, a.pCIeFunctions)
}
