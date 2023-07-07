// Code generated by "stringer -type=GlobalVars"; DO NOT EDIT.

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
	_ = x[GvRew-0]
	_ = x[GvHasRew-1]
	_ = x[GvRewPred-2]
	_ = x[GvPrevPred-3]
	_ = x[GvDA-4]
	_ = x[GvACh-5]
	_ = x[GvNE-6]
	_ = x[GvSer-7]
	_ = x[GvAChRaw-8]
	_ = x[GvNotMaint-9]
	_ = x[GvEffortRaw-10]
	_ = x[GvEffortDisc-11]
	_ = x[GvEffortCurMax-12]
	_ = x[GvUrgency-13]
	_ = x[GvUrgencyRaw-14]
	_ = x[GvVSMatrixJustGated-15]
	_ = x[GvVSMatrixHasGated-16]
	_ = x[GvHasRewPrev-17]
	_ = x[GvHasPosUSPrev-18]
	_ = x[GvLHbDip-19]
	_ = x[GvLHbBurst-20]
	_ = x[GvLHbDipSumCur-21]
	_ = x[GvLHbDipSum-22]
	_ = x[GvLHbGiveUp-23]
	_ = x[GvLHbPos-24]
	_ = x[GvLHbNeg-25]
	_ = x[GvVtaDA-26]
	_ = x[GvVtaUSpos-27]
	_ = x[GvVtaPVpos-28]
	_ = x[GvVtaPVneg-29]
	_ = x[GvVtaCeMpos-30]
	_ = x[GvVtaCeMneg-31]
	_ = x[GvVtaLHbDip-32]
	_ = x[GvVtaLHbBurst-33]
	_ = x[GvVtaVSPatchPos-34]
	_ = x[GvUSneg-35]
	_ = x[GvDrives-36]
	_ = x[GvBaseDrives-37]
	_ = x[GvDrivesTau-38]
	_ = x[GvDrivesUSDec-39]
	_ = x[GvUSpos-40]
	_ = x[GvVSPatch-41]
	_ = x[GlobalVarsN-42]
}

const _GlobalVars_name = "GvRewGvHasRewGvRewPredGvPrevPredGvDAGvAChGvNEGvSerGvAChRawGvNotMaintGvEffortRawGvEffortDiscGvEffortCurMaxGvUrgencyGvUrgencyRawGvVSMatrixJustGatedGvVSMatrixHasGatedGvHasRewPrevGvHasPosUSPrevGvLHbDipGvLHbBurstGvLHbDipSumCurGvLHbDipSumGvLHbGiveUpGvLHbPosGvLHbNegGvVtaDAGvVtaUSposGvVtaPVposGvVtaPVnegGvVtaCeMposGvVtaCeMnegGvVtaLHbDipGvVtaLHbBurstGvVtaVSPatchPosGvUSnegGvDrivesGvBaseDrivesGvDrivesTauGvDrivesUSDecGvUSposGvVSPatchGlobalVarsN"

var _GlobalVars_index = [...]uint16{0, 5, 13, 22, 32, 36, 41, 45, 50, 58, 68, 79, 91, 105, 114, 126, 145, 163, 175, 189, 197, 207, 221, 232, 243, 251, 259, 266, 276, 286, 296, 307, 318, 329, 342, 357, 364, 372, 384, 395, 408, 415, 424, 435}

func (i GlobalVars) String() string {
	if i < 0 || i >= GlobalVars(len(_GlobalVars_index)-1) {
		return "GlobalVars(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GlobalVars_name[_GlobalVars_index[i]:_GlobalVars_index[i+1]]
}

func (i *GlobalVars) FromString(s string) error {
	for j := 0; j < len(_GlobalVars_index)-1; j++ {
		if s == _GlobalVars_name[_GlobalVars_index[j]:_GlobalVars_index[j+1]] {
			*i = GlobalVars(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: GlobalVars")
}
