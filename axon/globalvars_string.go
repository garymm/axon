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
	_ = x[GvHadRew-4]
	_ = x[GvDA-5]
	_ = x[GvACh-6]
	_ = x[GvNE-7]
	_ = x[GvSer-8]
	_ = x[GvAChRaw-9]
	_ = x[GvNotMaint-10]
	_ = x[GvVSMatrixJustGated-11]
	_ = x[GvVSMatrixHasGated-12]
	_ = x[GvCuriosityPoolGated-13]
	_ = x[GvTime-14]
	_ = x[GvEffort-15]
	_ = x[GvUrgencyRaw-16]
	_ = x[GvUrgency-17]
	_ = x[GvHasPosUS-18]
	_ = x[GvHadPosUS-19]
	_ = x[GvNegUSOutcome-20]
	_ = x[GvHadNegUSOutcome-21]
	_ = x[GvPVposSum-22]
	_ = x[GvPVpos-23]
	_ = x[GvPVnegSum-24]
	_ = x[GvPVneg-25]
	_ = x[GvPVposEst-26]
	_ = x[GvPVposEstSum-27]
	_ = x[GvPVposEstDisc-28]
	_ = x[GvGiveUpDiff-29]
	_ = x[GvGiveUpProb-30]
	_ = x[GvGiveUp-31]
	_ = x[GvGaveUp-32]
	_ = x[GvVSPatchPos-33]
	_ = x[GvVSPatchPosPrev-34]
	_ = x[GvVSPatchPosSum-35]
	_ = x[GvLHbDip-36]
	_ = x[GvLHbBurst-37]
	_ = x[GvLHbPVDA-38]
	_ = x[GvCeMpos-39]
	_ = x[GvCeMneg-40]
	_ = x[GvVtaDA-41]
	_ = x[GvUSneg-42]
	_ = x[GvUSnegRaw-43]
	_ = x[GvDrives-44]
	_ = x[GvUSpos-45]
	_ = x[GvVSPatch-46]
	_ = x[GvVSPatchPrev-47]
	_ = x[GvOFCposUSPTMaint-48]
	_ = x[GvVSMatrixPoolGated-49]
	_ = x[GlobalVarsN-50]
}

const _GlobalVars_name = "GvRewGvHasRewGvRewPredGvPrevPredGvHadRewGvDAGvAChGvNEGvSerGvAChRawGvNotMaintGvVSMatrixJustGatedGvVSMatrixHasGatedGvCuriosityPoolGatedGvTimeGvEffortGvUrgencyRawGvUrgencyGvHasPosUSGvHadPosUSGvNegUSOutcomeGvHadNegUSOutcomeGvPVposSumGvPVposGvPVnegSumGvPVnegGvPVposEstGvPVposEstSumGvPVposEstDiscGvGiveUpDiffGvGiveUpProbGvGiveUpGvGaveUpGvVSPatchPosGvVSPatchPosPrevGvVSPatchPosSumGvLHbDipGvLHbBurstGvLHbPVDAGvCeMposGvCeMnegGvVtaDAGvUSnegGvUSnegRawGvDrivesGvUSposGvVSPatchGvVSPatchPrevGvOFCposUSPTMaintGvVSMatrixPoolGatedGlobalVarsN"

var _GlobalVars_index = [...]uint16{0, 5, 13, 22, 32, 40, 44, 49, 53, 58, 66, 76, 95, 113, 133, 139, 147, 159, 168, 178, 188, 202, 219, 229, 236, 246, 253, 263, 276, 290, 302, 314, 322, 330, 342, 358, 373, 381, 391, 400, 408, 416, 423, 430, 440, 448, 455, 464, 477, 494, 513, 524}

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
	0:  `Rew is reward value -- this is set here in the Context struct, and the RL Rew layer grabs it from there -- must also set HasRew flag when rew is set -- otherwise is ignored.`,
	1:  `HasRew must be set to true when a reward is present -- otherwise Rew is ignored. Also set when PVLV BOA model gives up. This drives ACh release in the PVLV model.`,
	2:  `RewPred is reward prediction -- computed by a special reward prediction layer`,
	3:  `PrevPred is previous time step reward prediction -- e.g., for TDPredLayer`,
	4:  `HadRew is HasRew state from the previous trial -- copied from HasRew in NewState -- used for updating Effort, Urgency at start of new trial`,
	5:  `DA is dopamine -- represents reward prediction error, signaled as phasic increases or decreases in activity relative to a tonic baseline, which is represented by a value of 0. Released by the VTA -- ventral tegmental area, or SNc -- substantia nigra pars compacta.`,
	6:  `ACh is acetylcholine -- activated by salient events, particularly at the onset of a reward / punishment outcome (US), or onset of a conditioned stimulus (CS). Driven by BLA -&gt; PPtg that detects changes in BLA activity, via LDTLayer type`,
	7:  `NE is norepinepherine -- not yet in use`,
	8:  `Ser is serotonin -- not yet in use`,
	9:  `AChRaw is raw ACh value used in updating global ACh value by LDTLayer`,
	10: `NotMaint is activity of the PTNotMaintLayer -- drives top-down inhibition of LDT layer / ACh activity.`,
	11: `VSMatrixJustGated is VSMatrix just gated (to engage goal maintenance in PFC areas), set at end of plus phase -- this excludes any gating happening at time of US`,
	12: `VSMatrixHasGated is VSMatrix has gated since the last time HasRew was set (US outcome received or expected one failed to be received`,
	13: `CuriosityPoolGated is true if VSMatrixJustGated and the first pool representing the curiosity / novelty drive gated -- this can change the giving up Effort.Max parameter.`,
	14: `Time is raw time counter, incrementing upward during goal engaged window. This is also copied directly into NegUS[0] which tracks time, but we maintain a separate effort value to make it clearer.`,
	15: `Effort is raw effort counter -- incrementing upward for each effort step during goal engaged window. This is also copied directly into NegUS[1] which tracks effort, but we maintain a separate effort value to make it clearer.`,
	16: `UrgencyRaw is raw effort for urgency -- incrementing upward from effort increments per step when _not_ goal engaged`,
	17: `Urgency is the overall urgency activity level (normalized 0-1), computed from logistic function of GvUrgencyRaw`,
	18: `HasPosUS indicates has positive US on this trial -- drives goal accomplishment logic and gating.`,
	19: `HadPosUS is state from the previous trial (copied from HasPosUS in NewState).`,
	20: `NegUSOutcome indicates that a strong negative US stimulus was experienced, driving phasic ACh, VSMatrix gating to reset current goal engaged plan (if any), and phasic dopamine based on the outcome.`,
	21: `HadNegUSOutcome is state from the previous trial (copied from NegUSOutcome in NewState)`,
	22: `PVposSum is total weighted positive valence primary value = sum of Weight * USpos * Drive`,
	23: `PVpos is normalized positive valence primary value = (1 - 1/(1+PVposGain * PVposSum))`,
	24: `PVnegSum is total weighted negative valence primary value = sum of Weight * USneg`,
	25: `PVpos is normalized negative valence primary value = (1 - 1/(1+PVnegGain * PVnegSum))`,
	26: `PVposEst is the estimated PVpos value based on OFCposUSPT and VSMatrix gating`,
	27: `PVposEstSum is the sum that goes into computing estimated PVpos value based on OFCposUSPT and VSMatrix gating`,
	28: `PVposEstDisc is the discounted version of PVposEst, subtracting VSPatchPosSum, which represents the accumulated expectation of PVpos to this point.`,
	29: `GiveUpDiff is the difference: PVposEstDisc - PVneg representing the expected positive outcome up to this point. When this turns negative, the chance of giving up goes up proportionally, as a logistic function of this difference.`,
	30: `GiveUpProb is the probability from the logistic function of GiveUpDiff`,
	31: `GiveUp is true if a reset was triggered probabilistically based on GiveUpProb`,
	32: `GaveUp is copy of GiveUp from previous trial`,
	33: `VSPatchPos is net shunting input from VSPatch (PosD1, named PVi in original PVLV) computed as the Max of US-specific VSPatch saved values. This is also stored as GvRewPred.`,
	34: `VSPatchPosPrev is the previous-trial version of VSPatchPos -- for adjusting the VSPatchThr threshold`,
	35: `VSPatchPosSum is the sum of VSPatchPos over goal engaged trials, representing the integrated prediction that the US is going to occur`,
	36: `computed LHb activity level that drives dipping / pausing of DA firing, when VSPatch pos prediction &gt; actual PV reward drive or PVneg &gt; PVpos`,
	37: `LHbBurst is computed LHb activity level that drives bursts of DA firing, when actual PV reward drive &gt; VSPatch pos prediction`,
	38: `LHbPVDA is GvLHbBurst - GvLHbDip -- the LHb contribution to DA, reflecting PV and VSPatch (PVi), but not the CS (LV) contributions`,
	39: `CeMpos is positive valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLAPosAcqD1 - BLAPosExtD2|_+ positively rectified. CeM sets Raw directly. Note that a positive US onset even with no active Drive will be reflected here, enabling learning about unexpected outcomes`,
	40: `CeMneg is negative valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLANegAcqD2 - BLANegExtD1|_+ positively rectified. CeM sets Raw directly`,
	41: `VtaDA is overall dopamine value reflecting all of the different inputs`,
	42: `USneg are negative valence US outcomes -- normalized version of raw, NNegUSs of them`,
	43: `USnegRaw are raw, linearly incremented negative valence US outcomes, this value is also integrated together with all US vals for PVneg`,
	44: `Drives is current drive state -- updated with optional homeostatic exponential return to baseline values`,
	45: `USpos is current positive-valence drive-satisfying input(s) (unconditioned stimuli = US)`,
	46: `VSPatch is current reward predicting VSPatch (PosD1) values`,
	47: `VSPatch is previous reward predicting VSPatch (PosD1) values`,
	48: `OFCposUSPTMaint is activity level of given OFCposUSPT maintenance pool used in anticipating potential USpos outcome value`,
	49: `VSMatrixPoolGated indicates whether given VSMatrix pool gated this is reset after last goal accomplished -- records gating since then.`,
	50: ``,
}

func (i GlobalVars) Desc() string {
	if str, ok := _GlobalVars_descMap[i]; ok {
		return str
	}
	return "GlobalVars(" + strconv.FormatInt(int64(i), 10) + ")"
}
