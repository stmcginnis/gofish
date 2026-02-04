//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/LeakDetector.v1_6_0.json
// 2025.4 - #LeakDetector.v1_6_0.LeakDetector

package schemas

import (
	"encoding/json"
)

type DetectorState string

const (
	// OKDetectorState shall indicate that there is no leak detected and the
	// detector is operating normally.
	OKDetectorState DetectorState = "OK"
	// WarningDetectorState shall indicate that a warning-level leak is detected.
	WarningDetectorState DetectorState = "Warning"
	// CriticalDetectorState shall indicate that a critical-level leak is detected.
	CriticalDetectorState DetectorState = "Critical"
	// UnavailableDetectorState shall indicate that no valid data can be acquired
	// from the detector, due to a fault condition within the detector, a
	// disconnected cable, or the detector is disabled.
	UnavailableDetectorState DetectorState = "Unavailable"
	// AbsentDetectorState shall indicate that the implementation supports a leak
	// detector, but no leak detector is installed or configured.
	AbsentDetectorState DetectorState = "Absent"
)

type LeakDetectorType string

const (
	// MoistureLeakDetectorType is a moisture sensor.
	MoistureLeakDetectorType LeakDetectorType = "Moisture"
	// FloatSwitchLeakDetectorType is a float switch.
	FloatSwitchLeakDetectorType LeakDetectorType = "FloatSwitch"
)

type ReactionType string

const (
	// NoneReactionType shall indicate no reaction occurs when a leak is detected.
	NoneReactionType ReactionType = "None"
	// ForceOffReactionType shall indicate that the associated device, equipment,
	// or system monitored by this leak detector is forcefully shut down when a
	// leak is detected.
	ForceOffReactionType ReactionType = "ForceOff"
	// GracefulShutdownReactionType shall indicate that the associated device,
	// equipment, or system monitored by this leak detector is gracefully shut down
	// when a leak is detected.
	GracefulShutdownReactionType ReactionType = "GracefulShutdown"
)

// LeakDetector shall represent a state-based or digital-value leak detector for
// a Redfish implementation.
type LeakDetector struct {
	Entity
	// CriticalReactionType shall be performed when the 'DetectorState' property
	// changes to 'Critical'.
	//
	// Version added: v1.4.0
	CriticalReactionType ReactionType
	// CurrentAmps shall contain the measured current, in ampere units, for this
	// leak detector. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'.
	//
	// Version added: v1.6.0
	CurrentAmps SensorCurrentExcerpt
	// DetectorState shall contain the state of the leak detector. The value of
	// this property should equate to the value of 'Health' in 'Status' when the
	// detector is enabled and functional. If a fault occurs with the detector
	// itself, such as a short or a disconnected cable, the 'Conditions' property
	// in 'Status' should indicate the type of fault detected.
	DetectorState DetectorState
	// Enabled shall indicate whether the leak detector is enabled. The value
	// 'true' shall indicate the leak detector is enabled. The value 'false' shall
	// indicate the leak detector is disabled. When disabled, 'DetectorState' shall
	// contain 'Unavailable' and the leak detector shall not trigger events,
	// logging, or other functionality. This property allows a user to disable a
	// faulty leak detector or to otherwise remove it from use.
	//
	// Version added: v1.3.0
	Enabled bool
	// LeakDetectorType shall contain the reading type of the leak detection
	// sensor.
	LeakDetectorType LeakDetectorType
	// Location shall indicate the location information for this leak detector.
	Location Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the leak detector. This organization may be the entity from whom
	// the leak detector is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to
	// the leak detector.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the leak detector.
	PartNumber string
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this leak detector applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this leak detector applies. This property
	// generally differentiates multiple leak detectors within the same
	// 'PhysicalContext' instance.
	PhysicalSubContext PhysicalSubContext
	// ReactionDelaySeconds shall indicate the number of seconds to delay after the
	// 'DetectorState' changes before the selected reaction is executed. If the
	// 'DetectorState' returns to 'OK' prior to the delay value, the service shall
	// not perform the reaction.
	//
	// Version added: v1.4.0
	ReactionDelaySeconds int
	// SKU shall contain the stock-keeping unit number for this leak detector.
	SKU string
	// SensingFrequency shall contain the time interval between readings of the
	// physical leak detector.
	SensingFrequency *float64 `json:",omitempty"`
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the leak detector.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the leak detector.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.1.0
	UserLabel string
	// Voltage shall contain the measured voltage, in volt units, for this leak
	// detector. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'.
	//
	// Version added: v1.6.0
	Voltage SensorVoltageExcerpt
	// WarningReactionType shall be performed when the 'DetectorState' property
	// changes to 'Warning'.
	//
	// Version added: v1.4.0
	WarningReactionType ReactionType
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a LeakDetector object from the raw JSON.
func (l *LeakDetector) UnmarshalJSON(b []byte) error {
	type temp LeakDetector
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LeakDetector(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	l.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LeakDetector) Update() error {
	readWriteFields := []string{
		"CriticalReactionType",
		"Enabled",
		"ReactionDelaySeconds",
		"UserLabel",
		"WarningReactionType",
	}

	return l.UpdateFromRawData(l, l.RawData, readWriteFields)
}

// GetLeakDetector will get a LeakDetector instance from the service.
func GetLeakDetector(c Client, uri string) (*LeakDetector, error) {
	return GetObject[LeakDetector](c, uri)
}

// ListReferencedLeakDetectors gets the collection of LeakDetector from
// a provided reference.
func ListReferencedLeakDetectors(c Client, link string) ([]*LeakDetector, error) {
	return GetCollectionObjects[LeakDetector](c, link)
}

// LeakDetectorArrayExcerpt shall represent a state-based or digital-value leak
// detector for a Redfish implementation.
type LeakDetectorArrayExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// DetectorState shall contain the state of the leak detector. The value of
	// this property should equate to the value of 'Health' in 'Status' when the
	// detector is enabled and functional. If a fault occurs with the detector
	// itself, such as a short or a disconnected cable, the 'Conditions' property
	// in 'Status' should indicate the type of fault detected.
	DetectorState DetectorState
	// DeviceName shall contain the name of the device associated with this leak
	// detector. If the device is represented by a resource, the value shall
	// contain the value of the 'Name' property of the associated resource.
	//
	// Version added: v1.2.0
	DeviceName string
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this leak detector applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this leak detector applies. This property
	// generally differentiates multiple leak detectors within the same
	// 'PhysicalContext' instance.
	PhysicalSubContext PhysicalSubContext
}

// LeakDetectorExcerpt shall represent a state-based or digital-value leak
// detector for a Redfish implementation.
type LeakDetectorExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// DetectorState shall contain the state of the leak detector. The value of
	// this property should equate to the value of 'Health' in 'Status' when the
	// detector is enabled and functional. If a fault occurs with the detector
	// itself, such as a short or a disconnected cable, the 'Conditions' property
	// in 'Status' should indicate the type of fault detected.
	DetectorState DetectorState
}
