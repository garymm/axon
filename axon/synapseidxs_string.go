// Code generated by "stringer -type=SynapseIdxs"; DO NOT EDIT.

package axon

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SynRecvIdx-0]
	_ = x[SynSendIdx-1]
	_ = x[SynPrjnIdx-2]
	_ = x[SynapseIdxsN-3]
}

const _SynapseIdxs_name = "SynRecvIdxSynSendIdxSynPrjnIdxSynapseIdxsN"

var _SynapseIdxs_index = [...]uint8{0, 10, 20, 30, 42}

func (i SynapseIdxs) String() string {
	if i < 0 || i >= SynapseIdxs(len(_SynapseIdxs_index)-1) {
		return "SynapseIdxs(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SynapseIdxs_name[_SynapseIdxs_index[i]:_SynapseIdxs_index[i+1]]
}

func (i *SynapseIdxs) FromString(s string) error {
	for j := 0; j < len(_SynapseIdxs_index)-1; j++ {
		if s == _SynapseIdxs_name[_SynapseIdxs_index[j]:_SynapseIdxs_index[j+1]] {
			*i = SynapseIdxs(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: SynapseIdxs")
}

var _SynapseIdxs_descMap = map[SynapseIdxs]string{
	0: `SynRecvIdx is receiving neuron index in network&#39;s global list of neurons`,
	1: `SynSendIdx is sending neuron index in network&#39;s global list of neurons`,
	2: `SynPrjnIdx is projection index in global list of projections organized as [Layers][RecvPrjns]`,
	3: ``,
}

func (i SynapseIdxs) Desc() string {
	if str, ok := _SynapseIdxs_descMap[i]; ok {
		return str
	}
	return "SynapseIdxs(" + strconv.FormatInt(int64(i), 10) + ")"
}
