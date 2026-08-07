//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// ThermalSubsystem shall represent a thermal subsystem for a Redfish implementation.
type ThermalSubsystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
	// filters shall contain a link to a resource collection of type Filter
	filters string
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
		Filters           common.Link
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
	thermalsubsystem.filters = t.Filters.String()

	return nil
}

// CoolantConnectors gets the coolant connectors for this equipment.
func (thermalsubsystem *ThermalSubsystem) CoolantConnectors(queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return thermalsubsystem.CoolantConnectorsWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// CoolantConnectorsWithContext gets the coolant connectors for this equipment.
func (thermalsubsystem *ThermalSubsystem) CoolantConnectorsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectorsWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.coolantConnectors, queryOpts...)
}

// Fans gets the fans for this equipment.
func (thermalsubsystem *ThermalSubsystem) Fans(queryOpts ...common.QueryGroupOption) ([]*Fan, error) {
	return thermalsubsystem.FansWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// FansWithContext gets the fans for this equipment.
func (thermalsubsystem *ThermalSubsystem) FansWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Fan, error) {
	return ListReferencedFansWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.fans, queryOpts...)
}

// Heaters gets the heaters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) Heaters(queryOpts ...common.QueryGroupOption) ([]*Heater, error) {
	return thermalsubsystem.HeatersWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// HeatersWithContext gets the heaters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) HeatersWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Heater, error) {
	return ListReferencedHeatersWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.heaters, queryOpts...)
}

// LeakDetection gets the leak detection system within the ThermalSubsystem.
// This property has been deprecated in favor of LeakDetectors under the Chassis resource.
func (thermalsubsystem *ThermalSubsystem) LeakDetection(queryOpts ...common.QueryGroupOption) (*LeakDetection, error) {
	return thermalsubsystem.LeakDetectionWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// LeakDetectionWithContext gets the leak detection system within the ThermalSubsystem.
// This property has been deprecated in favor of LeakDetectors under the Chassis resource.
func (thermalsubsystem *ThermalSubsystem) LeakDetectionWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*LeakDetection, error) {
	if thermalsubsystem.leakDetection == "" {
		return nil, nil
	}

	return GetLeakDetectionWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.leakDetection, queryOpts...)
}

// Pumps gets the pumps for this equipment.
func (thermalsubsystem *ThermalSubsystem) Pumps(queryOpts ...common.QueryGroupOption) ([]*Pump, error) {
	return thermalsubsystem.PumpsWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// PumpsWithContext gets the pumps for this equipment.
func (thermalsubsystem *ThermalSubsystem) PumpsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Pump, error) {
	return ListReferencedPumpsWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.pumps, queryOpts...)
}

// Filters gets the filters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) Filters(queryOpts ...common.QueryGroupOption) ([]*Filter, error) {
	return thermalsubsystem.FiltersWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// FiltersWithContext gets the filters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) FiltersWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Filter, error) {
	return ListReferencedFiltersWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.filters, queryOpts...)
}

// ThermalMetrics gets the summary of thermal metrics for this subsystem.
func (thermalsubsystem *ThermalSubsystem) ThermalMetrics(queryOpts ...common.QueryGroupOption) (*ThermalMetrics, error) {
	return thermalsubsystem.ThermalMetricsWithContext(common.ContextOf(thermalsubsystem.GetClient()), queryOpts...)
}

// ThermalMetricsWithContext gets the summary of thermal metrics for this subsystem.
func (thermalsubsystem *ThermalSubsystem) ThermalMetricsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*ThermalMetrics, error) {
	if thermalsubsystem.thermalMetrics == "" {
		return nil, nil
	}
	return GetThermalMetricsWithContext(ctx, thermalsubsystem.GetClient(), thermalsubsystem.thermalMetrics, queryOpts...)
}

// GetThermalSubsystem will get a ThermalSubsystem instance from the service.
func GetThermalSubsystem(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*ThermalSubsystem, error) {
	return GetThermalSubsystemWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetThermalSubsystemWithContext will get a ThermalSubsystem instance from the service.
func GetThermalSubsystemWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*ThermalSubsystem, error) {
	return common.GetObjectWithContext[ThermalSubsystem](ctx, c, uri, queryOpts...)
}

// ListReferencedThermalSubsystems gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystems(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*ThermalSubsystem, error) {
	return ListReferencedThermalSubsystemsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedThermalSubsystemsWithContext gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystemsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*ThermalSubsystem, error) {
	return common.GetCollectionObjectsWithContext[ThermalSubsystem](ctx, c, link, queryOpts...)
}
