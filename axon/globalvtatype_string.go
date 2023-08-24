// Code generated by "stringer -type=GlobalVTAType"; DO NOT EDIT.

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
	_ = x[GvVtaRaw-0]
	_ = x[GvVtaVals-1]
	_ = x[GvVtaPrev-2]
	_ = x[GlobalVTATypeN-3]
}

const _GlobalVTAType_name = "GvVtaRawGvVtaValsGvVtaPrevGlobalVTATypeN"

var _GlobalVTAType_index = [...]uint8{0, 8, 17, 26, 40}

func (i GlobalVTAType) String() string {
	if i < 0 || i >= GlobalVTAType(len(_GlobalVTAType_index)-1) {
		return "GlobalVTAType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GlobalVTAType_name[_GlobalVTAType_index[i]:_GlobalVTAType_index[i+1]]
}

func (i *GlobalVTAType) FromString(s string) error {
	for j := 0; j < len(_GlobalVTAType_index)-1; j++ {
		if s == _GlobalVTAType_name[_GlobalVTAType_index[j]:_GlobalVTAType_index[j+1]] {
			*i = GlobalVTAType(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: GlobalVTAType")
}

var _GlobalVTAType_descMap = map[GlobalVTAType]string{
	0: `GvVtaRaw are raw VTA values -- inputs to the computation`,
	1: `GvVtaVals are computed current VTA values`,
	2: `GvVtaPrev are previous computed values -- to avoid a data race`,
	3: ``,
}

func (i GlobalVTAType) Desc() string {
	if str, ok := _GlobalVTAType_descMap[i]; ok {
		return str
	}
	return "GlobalVTAType(" + strconv.FormatInt(int64(i), 10) + ")"
}
