// Code generated by "stringer -type=PrjnGTypes"; DO NOT EDIT.

package axon

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExcitatoryG-0]
	_ = x[InhibitoryG-1]
	_ = x[ModulatoryG-2]
	_ = x[PrjnGTypesN-3]
}

const _PrjnGTypes_name = "ExcitatoryGInhibitoryGModulatoryGPrjnGTypesN"

var _PrjnGTypes_index = [...]uint8{0, 11, 22, 33, 44}

func (i PrjnGTypes) String() string {
	if i < 0 || i >= PrjnGTypes(len(_PrjnGTypes_index)-1) {
		return "PrjnGTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PrjnGTypes_name[_PrjnGTypes_index[i]:_PrjnGTypes_index[i+1]]
}

func (i *PrjnGTypes) FromString(s string) error {
	for j := 0; j < len(_PrjnGTypes_index)-1; j++ {
		if s == _PrjnGTypes_name[_PrjnGTypes_index[j]:_PrjnGTypes_index[j+1]] {
			*i = PrjnGTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: PrjnGTypes")
}
