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
	_ = x[GvEffortCurMax-11]
	_ = x[GvUrgency-12]
	_ = x[GvUrgencyRaw-13]
	_ = x[GvVSMatrixJustGated-14]
	_ = x[GvVSMatrixHasGated-15]
	_ = x[GvHasRewPrev-16]
	_ = x[GvHasPosUSPrev-17]
	_ = x[GvLHbDip-18]
	_ = x[GvLHbBurst-19]
	_ = x[GvLHbPVDA-20]
	_ = x[GvLHbDipSumCur-21]
	_ = x[GvLHbDipSum-22]
	_ = x[GvLHbGiveUp-23]
	_ = x[GvLHbVSPatchPos-24]
	_ = x[GvLHbUSpos-25]
	_ = x[GvLHbPVpos-26]
	_ = x[GvLHbUSneg-27]
	_ = x[GvLHbPVneg-28]
	_ = x[GvCeMpos-29]
	_ = x[GvCeMneg-30]
	_ = x[GvVtaDA-31]
	_ = x[GvUSneg-32]
	_ = x[GvDrives-33]
	_ = x[GvBaseDrives-34]
	_ = x[GvDrivesTau-35]
	_ = x[GvDrivesUSDec-36]
	_ = x[GvUSpos-37]
	_ = x[GvVSPatch-38]
	_ = x[GlobalVarsN-39]
}

const _GlobalVars_name = "GvRewGvHasRewGvRewPredGvPrevPredGvDAGvAChGvNEGvSerGvAChRawGvNotMaintGvEffortRawGvEffortCurMaxGvUrgencyGvUrgencyRawGvVSMatrixJustGatedGvVSMatrixHasGatedGvHasRewPrevGvHasPosUSPrevGvLHbDipGvLHbBurstGvLHbPVDAGvLHbDipSumCurGvLHbDipSumGvLHbGiveUpGvLHbVSPatchPosGvLHbUSposGvLHbPVposGvLHbUSnegGvLHbPVnegGvCeMposGvCeMnegGvVtaDAGvUSnegGvDrivesGvBaseDrivesGvDrivesTauGvDrivesUSDecGvUSposGvVSPatchGlobalVarsN"

var _GlobalVars_index = [...]uint16{0, 5, 13, 22, 32, 36, 41, 45, 50, 58, 68, 79, 93, 102, 114, 133, 151, 163, 177, 185, 195, 204, 218, 229, 240, 255, 265, 275, 285, 295, 303, 311, 318, 325, 333, 345, 356, 369, 376, 385, 396}

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

var _GlobalVars_descMap = map[GlobalVars]string{
	0:  `GvRew is reward value -- this is set here in the Context struct, and the RL Rew layer grabs it from there -- must also set HasRew flag when rew is set -- otherwise is ignored.`,
	1:  `GvHasRew must be set to true when a reward is present -- otherwise Rew is ignored. Also set during extinction by PVLV. This drives ACh release in the PVLV model.`,
	2:  `GvRewPred is reward prediction -- computed by a special reward prediction layer`,
	3:  `GvPrevPred is previous time step reward prediction -- e.g., for TDPredLayer`,
	4:  `GvDA is dopamine -- represents reward prediction error, signaled as phasic increases or decreases in activity relative to a tonic baseline, which is represented by a value of 0. Released by the VTA -- ventral tegmental area, or SNc -- substantia nigra pars compacta.`,
	5:  `GvACh is acetylcholine -- activated by salient events, particularly at the onset of a reward / punishment outcome (US), or onset of a conditioned stimulus (CS). Driven by BLA -&gt; PPtg that detects changes in BLA activity, via LDTLayer type`,
	6:  `NE is norepinepherine -- not yet in use`,
	7:  `GvSer is serotonin -- not yet in use`,
	8:  `GvAChRaw is raw ACh value used in updating global ACh value by LDTLayer`,
	9:  `GvNotMaint is activity of the PTNotMaintLayer -- drives top-down inhibition of LDT layer / ACh activity.`,
	10: `GvEffortRaw is raw effort -- increments linearly upward for each additional effort step This is also copied directly into NegUS[0] which tracks effort, but we maintain a separate effort value to make it clearer.`,
	11: `GvEffortCurMax is current maximum raw effort level -- above this point, any current goal will be terminated during the GiveUp function, which also looks for accumulated disappointment. See Max, MaxNovel, MaxPostDip for values depending on how the goal was triggered`,
	12: `GvUrgency is the overall urgency activity level (normalized 0-1), computed from logistic function of GvUrgencyRaw`,
	13: `GvUrgencyRaw is raw effort for urgency -- increments linearly upward from effort increments per step`,
	14: `GvVSMatrixJustGated is VSMatrix just gated (to engage goal maintenance in PFC areas), set at end of plus phase -- this excludes any gating happening at time of US`,
	15: `GvVSMatrixHasGated is VSMatrix has gated since the last time HasRew was set (US outcome received or expected one failed to be received`,
	16: `HasRewPrev is state from the previous trial -- copied from HasRew in NewState -- used for updating Effort, Urgency at start of new trial`,
	17: `HasPosUSPrev is state from the previous trial -- copied from HasPosUS in NewState -- used for updating Effort, Urgency at start of new trial`,
	18: `computed LHb activity level that drives dipping / pausing of DA firing, when VSPatch pos prediction &gt; actual PV reward drive or PVNeg &gt; PVPos`,
	19: `GvLHbBurst is computed LHb activity level that drives bursts of DA firing, when actual PV reward drive &gt; VSPatch pos prediction`,
	20: `GvLHbPVDA is GvLHbBurst - GvLHbDip -- the LHb contribution to DA, reflecting PV and VSPatch (PVi), but not the CS (LV) contributions`,
	21: `GvLHbDipSumCur is current sum of LHbDip over trials, which is reset when there is a PV value, an above-threshold PPTg value, or when it triggers reset`,
	22: `GvLHbDipSum is copy of DipSum that is not reset -- used for driving negative dopamine dips on GiveUp trials`,
	23: `GvLHbGiveUp is true if a reset was triggered from LHbDipSum &gt; Reset Thr`,
	24: `GvLHbVSPatchPos is net shunting input from VSPatch (PosD1, named PVi in original PVLV)`,
	25: `GvLHbUSpos is total weighted positive valence primary value = sum of Weight * USpos * Drive`,
	26: `GvLHbPVpos is positive valence primary value (normalized USpos) = (1 - 1/(1+LHb.PosGain * USpos))`,
	27: `GvLHbUSneg is total weighted negative valence primary value = sum of Weight * USneg`,
	28: `GvLHbPVpos is positive valence primary value (normalized USpos) = (1 - 1/(1+LHb.NegGain * USpos))`,
	29: `GvCeMpos is positive valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLAPosAcqD1 - BLAPosExtD2|_+ positively rectified. CeM sets Raw directly. Note that a positive US onset even with no active Drive will be reflected here, enabling learning about unexpected outcomes`,
	30: `GvCeMneg is negative valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLANegAcqD2 - BLANegExtD1|_+ positively rectified. CeM sets Raw directly`,
	31: `GvVtaDA is overall dopamine value reflecting all of the different inputs`,
	32: `GvUSneg is negative valence US outcomes -- NNegUSs of them`,
	33: `GvDrives is current drive state -- updated with optional homeostatic exponential return to baseline values`,
	34: `GvDrivesBase are baseline levels for each drive -- what they naturally trend toward in the absence of any input. Set inactive drives to 0 baseline, active ones typically elevated baseline (0-1 range).`,
	35: `GvDrivesTau are time constants in ThetaCycle (trial) units for natural update toward Base values -- 0 values means no natural update.`,
	36: `GvDrivesUSDec are decrement factors for reducing drive value when Drive-US is consumed (multiply the US magnitude) -- these are positive valued numbers.`,
	37: `GvUSpos is current positive-valence drive-satisfying input(s) (unconditioned stimuli = US)`,
	38: `GvUSpos is current positive-valence drive-satisfying reward predicting VSPatch (PosD1) values`,
	39: ``,
}

func (i GlobalVars) Desc() string {
	if str, ok := _GlobalVars_descMap[i]; ok {
		return str
	}
	return "GlobalVars(" + strconv.FormatInt(int64(i), 10) + ")"
}
