//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type NTP struct {
	common.Entity

	OdataType          string `json:"@odata.type"`
	NTPEnable          bool   `json:"NTPEnable"`
	PrimaryNTPServer   string `json:"PrimaryNTPServer"`
	SecondaryNTPServer string `json:"SecondaryNTPServer"`
	DaylightSavingTime bool   `json:"DaylightSavingTime"`

	rawData []byte
}

func (ntp *NTP) UnmarshalJSON(b []byte) error {
	type temp NTP

	var rtn temp
	err := json.Unmarshal(b, &rtn)
	if err != nil {
		return err
	}

	*ntp = NTP(rtn)

	ntp.rawData = b

	return nil
}

func (ntp *NTP) Update() error {
	original := new(NTP)
	err := original.UnmarshalJSON(ntp.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"NTPEnable",
		"PrimaryNTPServer",
		"SecondaryNTPServer",
		"DaylightSavingTime",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ntp).Elem()

	return ntp.Entity.Update(originalElement, currentElement, readWriteFields)
}

func GetNTP(c common.Client, uri string) (*NTP, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ntp NTP
	err = json.NewDecoder(resp.Body).Decode(&ntp)
	if err != nil {
		return nil, err
	}

	ntp.SetClient(c)
	return &ntp, nil
}
