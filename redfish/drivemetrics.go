//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// DriveMetrics shall contain the usage and health statistics for a drive in a Redfish implementation.
type DriveMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BadBlockCount shall contain the total number of bad blocks reported by the drive.
	BadBlockCount int
	// CorrectableIOReadErrorCount shall contain the number of correctable read errors for the lifetime of the drive.
	CorrectableIOReadErrorCount int
	// CorrectableIOWriteErrorCount shall contain the number of correctable write errors for the lifetime of the drive.
	CorrectableIOWriteErrorCount int
	// Description provides a description of this resource.
	Description string
	// NVMeSMART shall contain the NVMe SMART metrics for the drive as defined by the NVMe SMART/Health Information log
	// page. This property shall not be present if the service represents NVMe controllers in the drive as
	// StorageController resources.
	NVMeSMART NVMeSMARTMetrics
	// NativeCommandQueueDepth shall contain the current depth of the Native Command Queue as defined by the SATA
	// Specification.
	NativeCommandQueueDepth int
	// PowerOnHours shall contain the number of power-on hours for the lifetime of the drive.
	PowerOnHours float64
	// ReadIOKiBytes shall contain the total number of kibibytes read from the time of last reset or wrap.
	ReadIOKiBytes int
	// UncorrectableIOReadErrorCount shall contain the number of uncorrectable read errors for the lifetime of the
	// drive.
	UncorrectableIOReadErrorCount int
	// UncorrectableIOWriteErrorCount shall contain the number of uncorrectable write errors for the lifetime of the
	// drive.
	UncorrectableIOWriteErrorCount int
	// WriteIOKiBytes shall contain the total number of kibibytes written from the time of last reset or wrap.
	WriteIOKiBytes int
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
