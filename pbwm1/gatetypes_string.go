// Code generated by "stringer -type=GateTypes"; DO NOT EDIT.

package pbwm

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Maint-0]
	_ = x[Out-1]
	_ = x[MaintOut-2]
	_ = x[GateTypesN-3]
}

const _GateTypes_name = "MaintOutMaintOutGateTypesN"

var _GateTypes_index = [...]uint8{0, 5, 8, 16, 26}

func (i GateTypes) String() string {
	if i < 0 || i >= GateTypes(len(_GateTypes_index)-1) {
		return "GateTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GateTypes_name[_GateTypes_index[i]:_GateTypes_index[i+1]]
}

func (i *GateTypes) FromString(s string) error {
	for j := 0; j < len(_GateTypes_index)-1; j++ {
		if s == _GateTypes_name[_GateTypes_index[j]:_GateTypes_index[j+1]] {
			*i = GateTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: GateTypes")
}