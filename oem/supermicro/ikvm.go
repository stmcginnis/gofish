//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type IKVMInterface string

const (
	JavaIKVMInterface IKVMInterface = "JAVA plug-in"
	HTMLIKVMInterface IKVMInterface = "HTML 5"
)

type IKVM struct {
	common.Entity

	OdataType        string        `json:"@odata.type"`
	CurrentInterface IKVMInterface `json:"Current interface"`
	URI              string        `json:"URI"`

	rawData []byte
}

func (ikvm *IKVM) UnmarshalJSON(b []byte) error {
	type temp IKVM

	var rtn temp
	err := json.Unmarshal(b, &rtn)
	if err != nil {
		return err
	}

	*ikvm = IKVM(rtn)
	ikvm.rawData = b

	return nil
}

func (ikvm *IKVM) Update() error {
	original := new(IKVM)
	err := original.UnmarshalJSON(ikvm.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Current interface",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ikvm).Elem()

	return ikvm.Entity.Update(originalElement, currentElement, readWriteFields)
}

func GetIKVM(c common.Client, uri string) (*IKVM, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ikvm IKVM
	err = json.NewDecoder(resp.Body).Decode(&ikvm)
	if err != nil {
		return nil, err
	}

	ikvm.SetClient(c)
	return &ikvm, nil
}
