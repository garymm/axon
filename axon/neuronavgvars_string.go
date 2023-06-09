// Code generated by "stringer -type=NeuronAvgVars"; DO NOT EDIT.

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
	_ = x[ActAvg-0]
	_ = x[AvgPct-1]
	_ = x[TrgAvg-2]
	_ = x[DTrgAvg-3]
	_ = x[AvgDif-4]
	_ = x[GeBase-5]
	_ = x[GiBase-6]
	_ = x[NeuronAvgVarsN-7]
}

const _NeuronAvgVars_name = "ActAvgAvgPctTrgAvgDTrgAvgAvgDifGeBaseGiBaseNeuronAvgVarsN"

var _NeuronAvgVars_index = [...]uint8{0, 6, 12, 18, 25, 31, 37, 43, 57}

func (i NeuronAvgVars) String() string {
	if i < 0 || i >= NeuronAvgVars(len(_NeuronAvgVars_index)-1) {
		return "NeuronAvgVars(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NeuronAvgVars_name[_NeuronAvgVars_index[i]:_NeuronAvgVars_index[i+1]]
}

func (i *NeuronAvgVars) FromString(s string) error {
	for j := 0; j < len(_NeuronAvgVars_index)-1; j++ {
		if s == _NeuronAvgVars_name[_NeuronAvgVars_index[j]:_NeuronAvgVars_index[j+1]] {
			*i = NeuronAvgVars(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: NeuronAvgVars")
}