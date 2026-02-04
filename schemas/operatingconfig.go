//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/OperatingConfig.v1_0_4.json
// 2020.2 - #OperatingConfig.v1_0_4.OperatingConfig

package schemas

import (
	"encoding/json"
)

// OperatingConfig shall represent an operational configuration for a processor
// in the Redfish Specification.
type OperatingConfig struct {
	Entity
	// BaseSpeedMHz shall contain the base (nominal) clock speed of the processor
	// in MHz.
	BaseSpeedMHz *uint `json:",omitempty"`
	// BaseSpeedPrioritySettings shall contain an array of objects that specify the
	// clock speed for sets of cores when the configuration is operational.
	BaseSpeedPrioritySettings []BaseSpeedPrioritySettings
	// MaxJunctionTemperatureCelsius shall contain the maximum temperature of the
	// junction in degree Celsius units.
	MaxJunctionTemperatureCelsius *int `json:",omitempty"`
	// MaxSpeedMHz shall contain the maximum clock speed to which the processor can
	// be configured in MHz.
	MaxSpeedMHz *uint `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// TDPWatts shall contain the thermal design point of the processor in watt
	// units.
	TDPWatts *uint `json:",omitempty"`
	// TotalAvailableCoreCount shall contain the number of cores in the processor
	// that can be configured.
	TotalAvailableCoreCount *uint `json:",omitempty"`
	// TurboProfile shall contain an array of objects that specify the turbo
	// profile for a set of active cores.
	TurboProfile []TurboProfileDatapoint
}

// GetOperatingConfig will get a OperatingConfig instance from the service.
func GetOperatingConfig(c Client, uri string) (*OperatingConfig, error) {
	return GetObject[OperatingConfig](c, uri)
}

// ListReferencedOperatingConfigs gets the collection of OperatingConfig from
// a provided reference.
func ListReferencedOperatingConfigs(c Client, link string) ([]*OperatingConfig, error) {
	return GetCollectionObjects[OperatingConfig](c, link)
}

// BaseSpeedPrioritySettings shall specify the clock speed for a set of cores.
type BaseSpeedPrioritySettings struct {
	// BaseSpeedMHz shall contain the clock speed to configure the set of cores in
	// MHz.
	BaseSpeedMHz *uint `json:",omitempty"`
	// CoreCount shall contain the number of cores to configure with the speed
	// specified by the 'BaseSpeedMHz' property. The sum of all 'CoreCount'
	// properties shall equal the value of the 'TotalAvailableCoreCount' property.
	CoreCount *uint `json:",omitempty"`
	// CoreIDs shall contain an array identifying the cores to configure with the
	// speed specified by the 'BaseSpeedMHz' property. The length of the array
	// shall equal the value of the 'CoreCount' property.
	CoreIDs []*int
}

// TurboProfileDatapoint shall specify the turbo profile for a set of active
// cores.
type TurboProfileDatapoint struct {
	// ActiveCoreCount shall contain the number of cores to be configured with the
	// maximum turbo clock speed. The value shall be less than or equal to the
	// 'TotalAvailableCoreCount' property.
	ActiveCoreCount *uint `json:",omitempty"`
	// MaxSpeedMHz shall contain the maximum turbo clock speed that correspond to
	// the number of active cores in MHz.
	MaxSpeedMHz *uint `json:",omitempty"`
}
