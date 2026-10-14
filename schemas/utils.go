//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"math"
	"strconv"
)

func parseString(val any) string {
	switch id := val.(type) {
	case string:
		return id
	case json.Number:
		return id.String()
	case int:
		return strconv.Itoa(id)
	case float32:
		return strconv.Itoa(int(id))
	case float64:
		return strconv.Itoa(int(id))
	}

	return ""
}

func toFloat32(val any) *float32 {
	if val == nil {
		return nil
	}

	var ret float32 = 0.0
	switch valu := val.(type) {
	case string:
		fl, err := strconv.ParseFloat(valu, 32)
		if err == nil {
			ret = float32(fl)
		}
	case int:
		ret = float32(valu)
	case float32:
		ret = float32(valu)
	case float64:
		conv := float32(valu)
		if math.IsInf(float64(conv), 1) {
			// Too big, return float32 max as a fallback
			ret = math.MaxFloat32
		} else if math.IsInf(float64(conv), 0) {
			// Too large negative
			ret = -math.MaxFloat32
		} else {
			ret = conv
		}
	}

	return &ret
}

func toInt(val any) *int {
	if val == nil {
		return nil
	}

	ret := int(*toFloat32(val))
	return &ret
}

// toFloat64 converts a JSON number or numeric string to a float64. It returns
// nil for a missing value or a string that is not a number.
func toFloat64(val any) *float64 {
	switch v := val.(type) {
	case float64:
		return &v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return &f
		}
	}
	return nil
}
