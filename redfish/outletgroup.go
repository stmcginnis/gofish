//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// OutletGroup shall be used to represent an electrical outlet group for a Redfish implementation.
type OutletGroup struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConfigurationLocked shall indicate whether modification requests to this resource are not permitted. If 'true',
	// services shall reject modification requests to other properties in this resource.
	ConfigurationLocked bool
	// CreatedBy shall contain the name of the person or application that created this outlet group.
	CreatedBy string
	// Description provides a description of this resource.
	Description string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this outlet group that represents the
	// 'Total' ElectricalContext sensor when multiple energy sensors exist for this outlet group. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerControlLocked shall indicate whether requests to the PowerControl action are locked. If 'true', services
	// shall reject requests to the PowerControl action.
	PowerControlLocked bool
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a PowerControl action to
	// cycle power. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerEnabled shall contain the power enable state of the outlet group. The value 'true' shall indicate that the
	// group can be powered on, and the value 'false' shall indicate that the group cannot be powered.
	PowerEnabled bool
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off after a PowerControl action. The
	// value '0' shall indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up after a power cycle or a PowerControl
	// action. The value '0' shall indicate no delay to power up.
	PowerOnDelaySeconds float64
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power on after a power fault. The value
	// '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds float64
	// PowerRestorePolicy shall contain the desired PowerState of the outlet group when power is applied. The value
	// 'LastState' shall return the outlet group to the PowerState it was in when power was lost.
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the outlet group.
	PowerState PowerState
	// PowerStateInTransition shall indicate whether the PowerState property will undergo a transition between on and
	// off states due to a configured delay. The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the PowerState property will transition at the conclusion of a configured
	// delay.
	PowerStateInTransition bool
	// PowerWatts shall contain the total power, in watt units, for this outlet group that represents the 'Total'
	// ElectricalContext sensor when multiple power sensors exist for this outlet group. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	outlets []string
	// OutletCounts is the number of outlets in this group.
	OutletsCount int

	powerControlTarget string
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a OutletGroup object from the raw JSON.
func (outletgroup *OutletGroup) UnmarshalJSON(b []byte) error {
	type temp OutletGroup
	type Actions struct {
		PowerControl common.ActionTarget `json:"#OutletGroup.PowerControl"`
		ResetMetrics common.ActionTarget `json:"#OutletGroup.ResetMetrics"`
	}
	type Links struct {
		Outlets      common.Links
		OutletsCount int `json:"Outlets@odata.count"`
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

	*outletgroup = OutletGroup(t.temp)

	// Extract the links to other entities for later
	outletgroup.outlets = t.Links.Outlets.ToStrings()
	outletgroup.OutletsCount = t.Links.OutletsCount

	outletgroup.powerControlTarget = t.Actions.PowerControl.Target
	outletgroup.resetMetricsTarget = t.Actions.ResetMetrics.Target

	// This is a read/write object, so we need to save the raw object data for later
	outletgroup.rawData = b

	return nil
}

// PowerControl controls the power state of the outlet group.
func (outletgroup *OutletGroup) PowerControl(powerState PowerState) error {
	params := struct {
		PowerState PowerState
	}{
		PowerState: powerState,
	}
	return outletgroup.Post(outletgroup.powerControlTarget, params)
}

// ResetMetrics resets metrics related to this outlet group.
func (outletgroup *OutletGroup) ResetMetrics() error {
	return outletgroup.Post(outletgroup.resetMetricsTarget, nil)
}

// Outlets get the outlets that are in this outlet group.
func (outletgroup *OutletGroup) Outlets() ([]*Outlet, error) {
	return common.GetObjects[Outlet](outletgroup.GetClient(), outletgroup.outlets)
}

// Update commits updates to this object's properties to the running system.
func (outletgroup *OutletGroup) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(OutletGroup)
	original.UnmarshalJSON(outletgroup.rawData)

	readWriteFields := []string{
		"ConfigurationLocked",
		"CreatedBy",
		"PowerControlLocked",
		"PowerCycleDelaySeconds",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestoreDelaySeconds",
		"PowerRestorePolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(outletgroup).Elem()

	return outletgroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetOutletGroup will get a OutletGroup instance from the service.
func GetOutletGroup(c common.Client, uri string) (*OutletGroup, error) {
	return common.GetObject[OutletGroup](c, uri)
}

// ListReferencedOutletGroups gets the collection of OutletGroup from
// a provided reference.
func ListReferencedOutletGroups(c common.Client, link string) ([]*OutletGroup, error) {
	return common.GetCollectionObjects[OutletGroup](c, link)
}
