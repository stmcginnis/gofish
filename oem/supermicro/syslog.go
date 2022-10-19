//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type Syslog struct {
	common.Entity

	OdataType        string `json:"@odata.type"`
	EnableSyslog     bool   `json:"EnableSyslog"`
	SyslogServer     string `json:"SyslogServer"`
	SyslogPortNumber int    `json:"SyslogPortNumber"`

	rawData []byte
}

func (syslog *Syslog) UnmarshalJSON(b []byte) error {
	type temp Syslog

	var rtn temp
	err := json.Unmarshal(b, &rtn)
	if err != nil {
		return err
	}

	*syslog = Syslog(rtn)
	syslog.rawData = b

	return nil
}

func (syslog *Syslog) Update() error {
	original := new(Syslog)
	err := original.UnmarshalJSON(syslog.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"EnableSyslog",
		"SyslogServer",
		"SyslogPortNumber",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(syslog).Elem()

	return syslog.Entity.Update(originalElement, currentElement, readWriteFields)
}

func GetSyslog(c common.Client, uri string) (*Syslog, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var syslog Syslog
	err = json.NewDecoder(resp.Body).Decode(&syslog)
	if err != nil {
		return nil, err
	}

	syslog.SetClient(c)
	return &syslog, nil
}
