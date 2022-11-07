// Code generated by "stringer -type=LayerType"; DO NOT EDIT.

package pcore

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Matrix_-12]
	_ = x[STN_-13]
	_ = x[GP_-14]
	_ = x[Thal_-15]
	_ = x[PT_-16]
	_ = x[LayerTypeN-17]
}

const _LayerType_name = "Matrix_STN_GP_Thal_PT_LayerTypeN"

var _LayerType_index = [...]uint8{0, 7, 11, 14, 19, 22, 32}

func (i LayerType) String() string {
	i -= 12
	if i < 0 || i >= LayerType(len(_LayerType_index)-1) {
		return "LayerType(" + strconv.FormatInt(int64(i+12), 10) + ")"
	}
	return _LayerType_name[_LayerType_index[i]:_LayerType_index[i+1]]
}

func StringToLayerType(s string) (LayerType, error) {
	for i := 0; i < len(_LayerType_index)-1; i++ {
		if s == _LayerType_name[_LayerType_index[i]:_LayerType_index[i+1]] {
			return LayerType(i + 12), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: LayerType")
}