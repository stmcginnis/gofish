//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #DriveMetrics.v1_3_0.DriveMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// DriveMetrics shall contain the usage and health statistics for a drive in a
// Redfish implementation.
type DriveMetrics struct {
	common.Entity
	// BadBlockCount shall contain the total number of bad blocks reported by the
	// drive.
	BadBlockCount *int `json:",omitempty"`
	// CorrectableIOReadErrorCount shall contain the number of correctable read
	// errors for the lifetime of the drive.
	CorrectableIOReadErrorCount *int `json:",omitempty"`
	// CorrectableIOWriteErrorCount shall contain the number of correctable write
	// errors for the lifetime of the drive.
	CorrectableIOWriteErrorCount *int `json:",omitempty"`
	// LifetimeStartDateTime shall contain the date and time when the drive started
	// accumulating data for properties that contain lifetime data, such as
	// 'UncorrectableIOReadErrorCount'. This might contain the same value as the
	// production date of the drive.
	//
	// Version added: v1.3.0
	LifetimeStartDateTime string
	// NVMeSMART shall contain the NVMe SMART metrics for the drive as defined by
	// the NVMe SMART/Health Information log page. This property shall not be
	// present if the service represents NVMe controllers in the drive as
	// StorageController resources.
	NVMeSMART NVMeSMARTMetrics
	// NativeCommandQueueDepth shall contain the current depth of the Native
	// Command Queue as defined by the SATA Specification.
	//
	// Version added: v1.1.0
	NativeCommandQueueDepth *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerOnHours shall contain the number of power-on hours for the lifetime of
	// the drive.
	PowerOnHours *float64 `json:",omitempty"`
	// ReadIOKiBytes shall contain the total number of kibibytes read from the time
	// of last reset or wrap.
	//
	// Version added: v1.2.0
	ReadIOKiBytes *int `json:",omitempty"`
	// UncorrectableIOReadErrorCount shall contain the number of uncorrectable read
	// errors for the lifetime of the drive.
	UncorrectableIOReadErrorCount *int `json:",omitempty"`
	// UncorrectableIOWriteErrorCount shall contain the number of uncorrectable
	// write errors for the lifetime of the drive.
	UncorrectableIOWriteErrorCount *int `json:",omitempty"`
	// WriteIOKiBytes shall contain the total number of kibibytes written from the
	// time of last reset or wrap.
	//
	// Version added: v1.2.0
	WriteIOKiBytes *int `json:",omitempty"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a DriveMetrics object from the raw JSON.
func (d *DriveMetrics) UnmarshalJSON(b []byte) error {
	type temp DriveMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DriveMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	d.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *DriveMetrics) Update() error {
	readWriteFields := []string{
		"NVMeSMART",
	}

	return d.UpdateFromRawData(d, d.rawData, readWriteFields)
}

// GetDriveMetrics will get a DriveMetrics instance from the service.
func GetDriveMetrics(c common.Client, uri string) (*DriveMetrics, error) {
	return common.GetObject[DriveMetrics](c, uri)
}

// ListReferencedDriveMetricss gets the collection of DriveMetrics from
// a provided reference.
func ListReferencedDriveMetricss(c common.Client, link string) ([]*DriveMetrics, error) {
	return common.GetCollectionObjects[DriveMetrics](c, link)
}
