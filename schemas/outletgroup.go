//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/OutletGroup.v1_2_0.json
// 2024.1 - #OutletGroup.v1_2_0.OutletGroup

package schemas

import (
	"encoding/json"
)

type OutletGroupType string

const (
	// HardwareDefinedOutletGroupType shall represent an outlet group that is
	// hardware-defined.
	HardwareDefinedOutletGroupType OutletGroupType = "HardwareDefined"
	// UserDefinedOutletGroupType shall represent an outlet group that is
	// user-defined.
	UserDefinedOutletGroupType OutletGroupType = "UserDefined"
)

// OutletGroup shall be used to represent an electrical outlet group for a
// Redfish implementation.
type OutletGroup struct {
	Entity
	// ConfigurationLocked shall indicate whether modification requests to this
	// resource are not permitted. If 'true', services shall reject modification
	// requests to other properties in this resource.
	//
	// Version added: v1.1.0
	ConfigurationLocked bool
	// CreatedBy shall contain the name of the person or application that created
	// this outlet group.
	CreatedBy string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// outlet group that represents the 'Total' 'ElectricalContext' sensor when
	// multiple energy sensors exist for this outlet group. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutletGroupType shall describe the type of this outlet group.
	//
	// Version added: v1.2.0
	OutletGroupType OutletGroupType
	// PowerControlLocked shall indicate whether requests to the 'PowerControl'
	// action are locked. If 'true', services shall reject requests to the
	// 'PowerControl' action.
	//
	// Version added: v1.1.0
	PowerControlLocked bool
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on
	// after a 'PowerControl' action to cycle power. The value '0' shall indicate
	// no delay to power on.
	PowerCycleDelaySeconds *float64 `json:",omitempty"`
	// PowerEnabled shall contain the power enable state of the outlet group. The
	// value 'true' shall indicate that the group can be powered on, and the value
	// 'false' shall indicate that the group cannot be powered.
	PowerEnabled bool
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off
	// after a 'PowerControl' action. The value '0' shall indicate no delay to
	// power off.
	PowerOffDelaySeconds *float64 `json:",omitempty"`
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up
	// after a power cycle or a 'PowerControl' action. The value '0' shall indicate
	// no delay to power up.
	PowerOnDelaySeconds *float64 `json:",omitempty"`
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power
	// on after a power fault. The value '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds *float64 `json:",omitempty"`
	// PowerRestorePolicy shall contain the desired 'PowerState' of the outlet
	// group when power is applied. The value 'LastState' shall return the outlet
	// group to the 'PowerState' it was in when power was lost.
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the outlet group.
	PowerState OutletPowerState
	// PowerStateInTransition shall indicate whether the 'PowerState' property will
	// undergo a transition between on and off states due to a configured delay.
	// The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the 'PowerState' property will
	// transition at the conclusion of a configured delay.
	//
	// Version added: v1.1.0
	PowerStateInTransition bool
	// PowerWatts shall contain the total power, in watt units, for this outlet
	// group that represents the 'Total' 'ElectricalContext' sensor when multiple
	// power sensors exist for this outlet group. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// Status shall contain any status or health properties of the resource.
	Status Status
	// powerControlTarget is the URL to send PowerControl requests.
	powerControlTarget string
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// outletGroups are the URIs for OutletGroups.
	outletGroups []string
	// outlets are the URIs for Outlets.
	outlets []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a OutletGroup object from the raw JSON.
func (o *OutletGroup) UnmarshalJSON(b []byte) error {
	type temp OutletGroup
	type oActions struct {
		PowerControl ActionTarget `json:"#OutletGroup.PowerControl"`
		ResetMetrics ActionTarget `json:"#OutletGroup.ResetMetrics"`
	}
	type oLinks struct {
		OutletGroups Links `json:"OutletGroups"`
		Outlets      Links `json:"Outlets"`
	}
	var tmp struct {
		temp
		Actions oActions
		Links   oLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = OutletGroup(tmp.temp)

	// Extract the links to other entities for later
	o.powerControlTarget = tmp.Actions.PowerControl.Target
	o.resetMetricsTarget = tmp.Actions.ResetMetrics.Target
	o.outletGroups = tmp.Links.OutletGroups.ToStrings()
	o.outlets = tmp.Links.Outlets.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	o.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (o *OutletGroup) Update() error {
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

	return o.UpdateFromRawData(o, o.RawData, readWriteFields)
}

// GetOutletGroup will get a OutletGroup instance from the service.
func GetOutletGroup(c Client, uri string) (*OutletGroup, error) {
	return GetObject[OutletGroup](c, uri)
}

// ListReferencedOutletGroups gets the collection of OutletGroup from
// a provided reference.
func ListReferencedOutletGroups(c Client, link string) ([]*OutletGroup, error) {
	return GetCollectionObjects[OutletGroup](c, link)
}

// This action shall control the power state of the outlet group.
// powerState - This parameter shall contain the desired power state of the
// outlet group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (o *OutletGroup) PowerControl(powerState OutletPowerState) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PowerState"] = powerState
	resp, taskInfo, err := PostWithTask(o.client,
		o.powerControlTarget, payload, o.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset any time intervals or counted values for this outlet
// group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (o *OutletGroup) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(o.client,
		o.resetMetricsTarget, payload, o.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// OutletGroups gets the OutletGroups linked resources.
func (o *OutletGroup) OutletGroups() ([]*OutletGroup, error) {
	return GetObjects[OutletGroup](o.client, o.outletGroups)
}

// Outlets gets the Outlets linked resources.
func (o *OutletGroup) Outlets() ([]*Outlet, error) {
	return GetObjects[Outlet](o.client, o.outlets)
}
