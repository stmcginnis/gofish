//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// AccelerationFunctionType is the type of acceleration provided.
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

// AccelerationFunction shall represent the acceleration function that a processor implements in a Redfish
// implementation. This can include functions such as audio processing, compression, encryption, packet inspection,
// packet switching, scheduling, or video processing.
type AccelerationFunction struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccelerationFunctionType shall contain the string that identifies the acceleration function type.
	AccelerationFunctionType AccelerationFunctionType
	// Description provides a description of this resource.
	Description string
	// FPGAReconfigurationSlots shall contain an array of the FPGA reconfiguration slot identifiers that this
	// acceleration function occupies.
	FPGAReconfigurationSlots []string `json:"FpgaReconfigurationSlots"`
	// Manufacturer shall contain a string that identifies the manufacturer of the acceleration function.
	Manufacturer string
	// PowerWatts shall contain the total acceleration function power consumption, in watt units.
	PowerWatts int
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// UUID shall contain a UUID for the acceleration function. RFC4122 describes methods that can create the value.
	// The value should be considered to be opaque. Client software should only treat the overall value as a UUID and
	// should not interpret any subfields within the UUID.
	UUID string
	// Version shall describe the acceleration function version.
	Version string
	// endpoints is a collection of URIs for connected endpoints.
	endpoints []string
	// EndpointsCount is the number of connected endpoints.
	EndpointsCount int
	// pcieFunctions is a collection of URIs to associated PCIe functions.
	pcieFunctions []string
	// PCIeFunctionsCount is the number of PCIe functions associated with this accelerator.
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a AccelerationFunction object from the raw JSON.
func (accelerationfunction *AccelerationFunction) UnmarshalJSON(b []byte) error {
	type temp AccelerationFunction
	type linkReference struct {
		Endpoints          common.Links
		EndpointsCount     int `json:"Endpoints@odata.count"`
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	}

	var t struct {
		temp
		Links linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*accelerationfunction = AccelerationFunction(t.temp)
	accelerationfunction.endpoints = t.Links.Endpoints.ToStrings()
	accelerationfunction.EndpointsCount = t.Links.EndpointsCount
	accelerationfunction.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	accelerationfunction.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	return nil
}

// GetAccelerationFunction will get a AccelerationFunction instance from the service.
func GetAccelerationFunction(c common.Client, uri string) (*AccelerationFunction, error) {
	return common.GetObject[AccelerationFunction](c, uri)
}

// ListReferencedAccelerationFunctions gets the collection of AccelerationFunction from
// a provided reference.
func ListReferencedAccelerationFunctions(c common.Client, link string) ([]*AccelerationFunction, error) {
	return common.GetCollectionObjects[AccelerationFunction](c, link)
}

// Endpoints gets the endpoints connected to this accelerator.
func (accelerationfunction *AccelerationFunction) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](accelerationfunction.GetClient(), accelerationfunction.endpoints)
}

// PCIeFunctions gets the PCIe functions associated with this accelerator.
func (accelerationfunction *AccelerationFunction) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](accelerationfunction.GetClient(), accelerationfunction.pcieFunctions)
}
