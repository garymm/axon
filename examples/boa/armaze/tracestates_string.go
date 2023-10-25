// Code generated by "stringer -type=TraceStates"; DO NOT EDIT.

package armaze

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TrSearching-0]
	_ = x[TrDeciding-1]
	_ = x[TrJustEngaged-2]
	_ = x[TrApproaching-3]
	_ = x[TrConsuming-4]
	_ = x[TrRewarded-5]
	_ = x[TrGiveUp-6]
	_ = x[TrBumping-7]
	_ = x[TraceStatesN-8]
}

const _TraceStates_name = "TrSearchingTrDecidingTrJustEngagedTrApproachingTrConsumingTrRewardedTrGiveUpTrBumpingTraceStatesN"

var _TraceStates_index = [...]uint8{0, 11, 21, 34, 47, 58, 68, 76, 85, 97}

func (i TraceStates) String() string {
	if i < 0 || i >= TraceStates(len(_TraceStates_index)-1) {
		return "TraceStates(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TraceStates_name[_TraceStates_index[i]:_TraceStates_index[i+1]]
}

func (i *TraceStates) FromString(s string) error {
	for j := 0; j < len(_TraceStates_index)-1; j++ {
		if s == _TraceStates_name[_TraceStates_index[j]:_TraceStates_index[j+1]] {
			*i = TraceStates(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: TraceStates")
}

var _TraceStates_descMap = map[TraceStates]string{
	0: `Searching is not yet goal engaged, looking for a goal`,
	1: `Deciding is having some partial gating but not in time for action`,
	2: `JustEngaged means just decided to engage in a goal`,
	3: `Approaching is goal engaged, approaching the goal`,
	4: `Consuming is consuming the US, first step (prior to getting reward, step1)`,
	5: `Rewarded is just received reward from a US`,
	6: `GiveUp is when goal is abandoned`,
	7: `Bumping is bumping into a wall`,
	8: ``,
}

func (i TraceStates) Desc() string {
	if str, ok := _TraceStates_descMap[i]; ok {
		return str
	}
	return "TraceStates(" + strconv.FormatInt(int64(i), 10) + ")"
}