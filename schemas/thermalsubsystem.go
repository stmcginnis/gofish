//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #ThermalSubsystem.v1_5_0.ThermalSubsystem

package schemas

import (
	"encoding/json"
)

// ThermalSubsystem shall represent a thermal subsystem for a Redfish
// implementation.
type ThermalSubsystem struct {
	Entity
	// CoolantConnectorRedundancy shall contain redundancy information for the set
	// of coolant connectors attached to this equipment. The values of the
	// 'RedundancyGroup' array shall reference resources of type
	// 'CoolantConnector'.
	//
	// Version added: v1.3.0
	CoolantConnectorRedundancy []RedundantGroup
	// CoolantConnectors shall contain a link to a resource collection of type
	// 'CoolantConnectorCollection' that contains the coolant connectors for this
	// equipment.
	//
	// Version added: v1.2.0
	coolantConnectors string
	// FanRedundancy shall contain redundancy information for the groups of fans in
	// this subsystem.
	FanRedundancy []RedundantGroup
	// Fans shall contain a link to a resource collection of type 'FanCollection'.
	fans string
	// FansFullSpeedOverrideEnable shall indicate whether the fans in this
	// equipment are overridden to operate at full speed.
	//
	// Version added: v1.5.0
	FansFullSpeedOverrideEnable bool
	// Filters shall contain a link to a resource collection of type
	// 'FilterCollection' that contains the filters for this equipment.
	//
	// Version added: v1.4.0
	filters string
	// Heaters shall contain a link to a resource collection of type
	// 'HeaterCollection'.
	//
	// Version added: v1.1.0
	heaters string
	// LeakDetection shall contain a link to a resource of type 'LeakDetection'.
	// This link should be used when the leak detection capabilities are tied to,
	// or are internal to, a particular 'Chassis'. For detection capabilities that
	// are tied to a 'CoolingUnit' resource, which may span multiple 'Chassis'
	// resources, populating the 'LeakDetection' resource under 'CoolingUnit' for
	// the relevant equipment is the preferred approach.
	//
	// Version added: v1.3.0
	leakDetection string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Pumps shall contain a link to a resource collection of type 'PumpCollection'
	// that contains details for the pumps included in this equipment.
	//
	// Version added: v1.3.0
	pumps string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// ThermalMetrics shall contain a link to a resource of type 'ThermalMetrics'.
	thermalMetrics string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ThermalSubsystem object from the raw JSON.
func (t *ThermalSubsystem) UnmarshalJSON(b []byte) error {
	type temp ThermalSubsystem
	var tmp struct {
		temp
		CoolantConnectors Link `json:"CoolantConnectors"`
		Fans              Link `json:"Fans"`
		Filters           Link `json:"Filters"`
		Heaters           Link `json:"Heaters"`
		LeakDetection     Link `json:"LeakDetection"`
		Pumps             Link `json:"Pumps"`
		ThermalMetrics    Link `json:"ThermalMetrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = ThermalSubsystem(tmp.temp)

	// Extract the links to other entities for later
	t.coolantConnectors = tmp.CoolantConnectors.String()
	t.fans = tmp.Fans.String()
	t.filters = tmp.Filters.String()
	t.heaters = tmp.Heaters.String()
	t.leakDetection = tmp.LeakDetection.String()
	t.pumps = tmp.Pumps.String()
	t.thermalMetrics = tmp.ThermalMetrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *ThermalSubsystem) Update() error {
	readWriteFields := []string{
		"FansFullSpeedOverrideEnable",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetThermalSubsystem will get a ThermalSubsystem instance from the service.
func GetThermalSubsystem(c Client, uri string) (*ThermalSubsystem, error) {
	return GetObject[ThermalSubsystem](c, uri)
}

// ListReferencedThermalSubsystems gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystems(c Client, link string) ([]*ThermalSubsystem, error) {
	return GetCollectionObjects[ThermalSubsystem](c, link)
}

// CoolantConnectors gets the CoolantConnectors collection.
func (t *ThermalSubsystem) CoolantConnectors() ([]*CoolantConnector, error) {
	if t.coolantConnectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolantConnector](t.client, t.coolantConnectors)
}

// Fans gets the Fans collection.
func (t *ThermalSubsystem) Fans() ([]*ThermalFan, error) {
	if t.fans == "" {
		return nil, nil
	}
	return GetCollectionObjects[ThermalFan](t.client, t.fans)
}

// Filters gets the Filters collection.
func (t *ThermalSubsystem) Filters() ([]*Filter, error) {
	if t.filters == "" {
		return nil, nil
	}
	return GetCollectionObjects[Filter](t.client, t.filters)
}

// Heaters gets the Heaters collection.
func (t *ThermalSubsystem) Heaters() ([]*Heater, error) {
	if t.heaters == "" {
		return nil, nil
	}
	return GetCollectionObjects[Heater](t.client, t.heaters)
}

// LeakDetection gets the LeakDetection linked resource.
func (t *ThermalSubsystem) LeakDetection() (*LeakDetection, error) {
	if t.leakDetection == "" {
		return nil, nil
	}
	return GetObject[LeakDetection](t.client, t.leakDetection)
}

// Pumps gets the Pumps collection.
func (t *ThermalSubsystem) Pumps() ([]*Pump, error) {
	if t.pumps == "" {
		return nil, nil
	}
	return GetCollectionObjects[Pump](t.client, t.pumps)
}

// ThermalMetrics gets the ThermalMetrics linked resource.
func (t *ThermalSubsystem) ThermalMetrics() (*ThermalMetrics, error) {
	if t.thermalMetrics == "" {
		return nil, nil
	}
	return GetObject[ThermalMetrics](t.client, t.thermalMetrics)
}
