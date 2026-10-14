//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "fmt"

func (a *ActionInfo) GetParamValues(name string, dataType ParameterTypes) ([]string, error) {
	for idx := range a.Parameters {
		param := &a.Parameters[idx]
		if param.Name != name || (param.DataType != "" && param.DataType != dataType) {
			continue
		}

		return param.AllowableValues, nil
	}
	return nil, fmt.Errorf("failed to find supported values of type %s", name)
}
