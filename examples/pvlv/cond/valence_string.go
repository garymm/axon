// Code generated by "stringer -type=Valence"; DO NOT EDIT.

package cond

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Pos-0]
	_ = x[Neg-1]
	_ = x[ValenceN-2]
}

const _Valence_name = "PosNegValenceN"

var _Valence_index = [...]uint8{0, 3, 6, 14}

func (i Valence) String() string {
	if i < 0 || i >= Valence(len(_Valence_index)-1) {
		return "Valence(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Valence_name[_Valence_index[i]:_Valence_index[i+1]]
}

func (i *Valence) FromString(s string) error {
	for j := 0; j < len(_Valence_index)-1; j++ {
		if s == _Valence_name[_Valence_index[j]:_Valence_index[j+1]] {
			*i = Valence(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Valence")
}