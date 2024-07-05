//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ThermalSubsystem shall represent a thermal subsystem for a Redfish implementation.
type ThermalSubsystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CoolantConnectorRedundancy shall contain redundancy information for the set of coolant connectors attached to
	// this equipment. The values of the RedundancyGroup array shall reference resources of type CoolantConnector.
	CoolantConnectorRedundancy []RedundantGroup
	// CoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that contains
	// the coolant connectors for this equipment.
	coolantConnectors string
	// Description provides a description of this resource.
	Description string
	// FanRedundancy shall contain redundancy information for the groups of fans in this subsystem.
	FanRedundancy []RedundantGroup
	// Fans shall contain a link to a resource collection of type FanCollection.
	fans string
	// Heaters shall contain a link to a resource collection of type HeaterCollection.
	heaters string
	// LeakDetection shall contain a link to a resource collection of type LeakDetection.
	leakDetection string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Pumps shall contain a link to a resource collection of type PumpCollection that contains details for the pumps
	// included in this equipment.
	pumps string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ThermalMetrics shall contain a link to a resource collection of type ThermalMetrics.
	thermalMetrics string
}

// UnmarshalJSON unmarshals a ThermalSubsystem object from the raw JSON.
func (thermalsubsystem *ThermalSubsystem) UnmarshalJSON(b []byte) error {
	type temp ThermalSubsystem
	var t struct {
		temp
		CoolantConnectors common.Link
		Fans              common.Link
		Heaters           common.Link
		LeakDetection     common.Link
		Pumps             common.Link
		ThermalMetrics    common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalsubsystem = ThermalSubsystem(t.temp)

	// Extract the links to other entities for later
	thermalsubsystem.coolantConnectors = t.CoolantConnectors.String()
	thermalsubsystem.fans = t.Fans.String()
	thermalsubsystem.heaters = t.Heaters.String()
	thermalsubsystem.leakDetection = t.LeakDetection.String()
	thermalsubsystem.pumps = t.Pumps.String()
	thermalsubsystem.thermalMetrics = t.ThermalMetrics.String()

	return nil
}

// CoolantConnectors gets the coolant connectors for this equipment.
func (thermalsubsystem *ThermalSubsystem) CoolantConnectors() ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectors(thermalsubsystem.GetClient(), thermalsubsystem.coolantConnectors)
}

// Fans gets the fans for this equipment.
func (thermalsubsystem *ThermalSubsystem) Fans() ([]*Fan, error) {
	return ListReferencedFans(thermalsubsystem.GetClient(), thermalsubsystem.fans)
}

// Heaters gets the heaters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) Heaters() ([]*Heater, error) {
	return ListReferencedHeaters(thermalsubsystem.GetClient(), thermalsubsystem.heaters)
}

// LeakDetection gets the leak detection system within this chassis.
func (thermalsubsystem *ThermalSubsystem) LeakDetection() ([]*LeakDetection, error) {
	return ListReferencedLeakDetections(thermalsubsystem.GetClient(), thermalsubsystem.leakDetection)
}

// Pumps gets the pumps for this equipment.
func (thermalsubsystem *ThermalSubsystem) Pumps() ([]*Pump, error) {
	return ListReferencedPumps(thermalsubsystem.GetClient(), thermalsubsystem.pumps)
}

// ThermalMetrics gets the summary of thermal metrics for this subsystem.
func (thermalsubsystem *ThermalSubsystem) ThermalMetrics() (*ThermalMetrics, error) {
	if thermalsubsystem.thermalMetrics == "" {
		return nil, nil
	}
	return GetThermalMetrics(thermalsubsystem.GetClient(), thermalsubsystem.thermalMetrics)
}

// GetThermalSubsystem will get a ThermalSubsystem instance from the service.
func GetThermalSubsystem(c common.Client, uri string) (*ThermalSubsystem, error) {
	return common.GetObject[ThermalSubsystem](c, uri)
}

// ListReferencedThermalSubsystems gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystems(c common.Client, link string) ([]*ThermalSubsystem, error) {
	return common.GetCollectionObjects[ThermalSubsystem](c, link)
}
