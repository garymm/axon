// Code generated by "stringer -type=MatrixPatch"; DO NOT EDIT.

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
	_ = x[Matrix-0]
	_ = x[Patch-1]
	_ = x[MatrixPatchN-2]
}

const _MatrixPatch_name = "MatrixPatchMatrixPatchN"

var _MatrixPatch_index = [...]uint8{0, 6, 11, 23}

func (i MatrixPatch) String() string {
	if i < 0 || i >= MatrixPatch(len(_MatrixPatch_index)-1) {
		return "MatrixPatch(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MatrixPatch_name[_MatrixPatch_index[i]:_MatrixPatch_index[i+1]]
}

func (i *MatrixPatch) FromString(s string) error {
	for j := 0; j < len(_MatrixPatch_index)-1; j++ {
		if s == _MatrixPatch_name[_MatrixPatch_index[j]:_MatrixPatch_index[j+1]] {
			*i = MatrixPatch(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: MatrixPatch")
}