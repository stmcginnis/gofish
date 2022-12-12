//
// SPDX-License-Identifier: BSD-3-Clause
//

package huawei

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/redfish"
)

type VolumeOem struct {
	Huawei struct {
		State              string `json:"State"`
		VolumeName         string `json:"VolumeName"`
		RaidControllerID   int    `json:"RaidControllerID"`
		VolumeRaidLevel    string `json:"VolumeRaidLevel"`
		DefaultReadPolicy  string `json:"DefaultReadPolicy"`
		DefaultWritePolicy string `json:"DefaultWritePolicy"`
		DefaultCachePolicy string `json:"DefaultCachePolicy"`
		ConsistencyCheck   bool   `json:"ConsistencyCheck"`
		SpanNumber         int    `json:"SpanNumber"`
		NumDrivePerSpan    int    `json:"NumDrivePerSpan"`
		Spans              []struct {
			SpanName string `json:"SpanName"`
			Drives   []struct {
				OdataID string `json:"@odata.id"`
			} `json:"Drives"`
		} `json:"Spans"`
		CurrentReadPolicy         string        `json:"CurrentReadPolicy"`
		CurrentWritePolicy        string        `json:"CurrentWritePolicy"`
		CurrentCachePolicy        string        `json:"CurrentCachePolicy"`
		AccessPolicy              string        `json:"AccessPolicy"`
		BootEnable                bool          `json:"BootEnable"`
		BGIEnable                 bool          `json:"BGIEnable"`
		SSDCachecadeVolume        bool          `json:"SSDCachecadeVolume"`
		SSDCachingEnable          bool          `json:"SSDCachingEnable"`
		AssociatedCacheCadeVolume []interface{} `json:"AssociatedCacheCadeVolume"`
		DriveCachePolicy          string        `json:"DriveCachePolicy"`
		OSDriveName               interface{}   `json:"OSDriveName"`
		InitializationMode        string        `json:"InitializationMode"`
	} `json:"Huawei"`
}

type Volumes struct {
	redfish.Volume
	Oem VolumeOem
}

func FromVolume(volume *redfish.Volume) (Volumes, error) {
	oem := VolumeOem{}

	_ = json.Unmarshal(volume.Oem, &oem)

	return Volumes{*volume, oem}, nil
}
