// Code generated by "stringer -type=ValenceTypes"; DO NOT EDIT.

package axon

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Positive-0]
	_ = x[Negative-1]
	_ = x[ValenceTypesN-2]
}

const _ValenceTypes_name = "PositiveNegativeValenceTypesN"

var _ValenceTypes_index = [...]uint8{0, 8, 16, 29}

func (i ValenceTypes) String() string {
	if i < 0 || i >= ValenceTypes(len(_ValenceTypes_index)-1) {
		return "ValenceTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValenceTypes_name[_ValenceTypes_index[i]:_ValenceTypes_index[i+1]]
}

func (i *ValenceTypes) FromString(s string) error {
	for j := 0; j < len(_ValenceTypes_index)-1; j++ {
		if s == _ValenceTypes_name[_ValenceTypes_index[j]:_ValenceTypes_index[j+1]] {
			*i = ValenceTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ValenceTypes")
}

var _ValenceTypes_descMap = map[ValenceTypes]string{
	0: `Positive valence codes for outcomes aligned with drives / goals.`,
	1: `Negative valence codes for harmful or aversive outcomes.`,
	2: ``,
}

func (i ValenceTypes) Desc() string {
	if str, ok := _ValenceTypes_descMap[i]; ok {
		return str
	}
	return "ValenceTypes(" + strconv.FormatInt(int64(i), 10) + ")"
}
