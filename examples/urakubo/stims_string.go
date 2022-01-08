// Code generated by "stringer -type=Stims"; DO NOT EDIT.

package main

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Baseline-0]
	_ = x[CaTarg-1]
	_ = x[ClampCa1-2]
	_ = x[StimsN-3]
}

const _Stims_name = "BaselineCaTargClampCa1StimsN"

var _Stims_index = [...]uint8{0, 8, 14, 22, 28}

func (i Stims) String() string {
	if i < 0 || i >= Stims(len(_Stims_index)-1) {
		return "Stims(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Stims_name[_Stims_index[i]:_Stims_index[i+1]]
}

func (i *Stims) FromString(s string) error {
	for j := 0; j < len(_Stims_index)-1; j++ {
		if s == _Stims_name[_Stims_index[j]:_Stims_index[j+1]] {
			*i = Stims(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Stims")
}