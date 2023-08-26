// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package axon

import "github.com/goki/ki/kit"

//go:generate stringer -type=GlobalVars

var KiT_GlobalVars = kit.Enums.AddEnum(GlobalVarsN, kit.NotBitFlag, nil)

func (ev GlobalVars) MarshalJSON() ([]byte, error)  { return kit.EnumMarshalJSON(ev) }
func (ev *GlobalVars) UnmarshalJSON(b []byte) error { return kit.EnumUnmarshalJSON(ev, b) }

//gosl: start globals

// GlobalVars are network-wide variables, such as neuromodulators, reward, drives, etc
// including the state for the PVLV phasic dopamine model.
type GlobalVars int32

const (
	/////////////////////////////////////////
	// Reward

	// GvRew is reward value -- this is set here in the Context struct, and the RL Rew layer grabs it from there -- must also set HasRew flag when rew is set -- otherwise is ignored.
	GvRew GlobalVars = iota

	// GvHasRew must be set to true when a reward is present -- otherwise Rew is ignored.  Also set during extinction by PVLV.  This drives ACh release in the PVLV model.
	GvHasRew

	// GvRewPred is reward prediction -- computed by a special reward prediction layer
	GvRewPred

	// GvPrevPred is previous time step reward prediction -- e.g., for TDPredLayer
	GvPrevPred

	/////////////////////////////////////////
	// NeuroMod neuromodulators

	// GvDA is dopamine -- represents reward prediction error, signaled as phasic increases or decreases in activity relative to a tonic baseline, which is represented by a value of 0.  Released by the VTA -- ventral tegmental area, or SNc -- substantia nigra pars compacta.
	GvDA

	// GvACh is acetylcholine -- activated by salient events, particularly at the onset of a reward / punishment outcome (US), or onset of a conditioned stimulus (CS).  Driven by BLA -> PPtg that detects changes in BLA activity, via LDTLayer type
	GvACh

	// NE is norepinepherine -- not yet in use
	GvNE

	// GvSer is serotonin -- not yet in use
	GvSer

	// GvAChRaw is raw ACh value used in updating global ACh value by LDTLayer
	GvAChRaw

	// GvNotMaint is activity of the PTNotMaintLayer -- drives top-down inhibition of LDT layer / ACh activity.
	GvNotMaint

	/////////////////////////////////////////
	// Effort & Urgency

	// GvEffortRaw is raw effort -- increments linearly upward for each additional effort step
	// This is also copied directly into NegUS[0] which tracks effort, but we maintain
	// a separate effort value to make it clearer.
	GvEffortRaw

	// GvEffortCurMax is current maximum raw effort level -- above this point, any current goal will be terminated during the GiveUp function, which also looks for accumulated disappointment.  See Max, MaxNovel, MaxPostDip for values depending on how the goal was triggered
	GvEffortCurMax

	// GvUrgency is the overall urgency activity level (normalized 0-1), computed from logistic function of GvUrgencyRaw
	GvUrgency

	// GvUrgencyRaw is raw effort for urgency -- increments linearly upward from effort increments per step
	GvUrgencyRaw

	/////////////////////////////////////////
	// VSMatrix gating and PVLV Rew flags

	// GvVSMatrixJustGated is VSMatrix just gated (to engage goal maintenance in PFC areas), set at end of plus phase -- this excludes any gating happening at time of US
	GvVSMatrixJustGated

	// GvVSMatrixHasGated is VSMatrix has gated since the last time HasRew was set (US outcome received or expected one failed to be received
	GvVSMatrixHasGated

	// HasRewPrev is state from the previous trial -- copied from HasRew in NewState -- used for updating Effort, Urgency at start of new trial
	GvHasRewPrev

	// HasPosUSPrev is state from the previous trial -- copied from HasPosUS in NewState -- used for updating Effort, Urgency at start of new trial
	GvHasPosUSPrev

	/////////////////////////////////////////
	// LHb lateral habenula component of the PVLV model -- does all US processing

	// computed LHb activity level that drives dipping / pausing of DA firing,
	// when VSPatch pos prediction > actual PV reward drive
	// or PVNeg > PVPos
	GvLHbDip

	// GvLHbBurst is computed LHb activity level that drives bursts of DA firing, when actual PV reward drive > VSPatch pos prediction
	GvLHbBurst

	// GvLHbPVDA is GvLHbBurst - GvLHbDip -- the LHb contribution to DA, reflecting PV and VSPatch (PVi), but not the CS (LV) contributions
	GvLHbPVDA

	// GvLHbDipSumCur is current sum of LHbDip over trials, which is reset when there is a PV value, an above-threshold PPTg value, or when it triggers reset
	GvLHbDipSumCur

	// GvLHbDipSum is copy of DipSum that is not reset -- used for driving negative dopamine dips on GiveUp trials
	GvLHbDipSum

	// GvLHbGiveUp is true if a reset was triggered from LHbDipSum > Reset Thr
	GvLHbGiveUp

	// GvLHbVSPatchPos is net shunting input from VSPatch (PosD1, named PVi in original PVLV)
	GvLHbVSPatchPos

	// GvLHbUSpos is total weighted positive valence primary value = sum of Weight * USpos * Drive
	GvLHbUSpos

	// GvLHbPVpos is positive valence primary value (normalized USpos) = (1 - 1/(1+LHb.PosGain * USpos))
	GvLHbPVpos

	// GvLHbUSneg is total weighted negative valence primary value = sum of Weight * USneg
	GvLHbUSneg

	// GvLHbPVpos is positive valence primary value (normalized USpos) = (1 - 1/(1+LHb.NegGain * USpos))
	GvLHbPVneg

	/////////////////////////////////////////
	// Amygdala CS / LV variables

	// GvCeMpos is positive valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLAPosAcqD1 - BLAPosExtD2|_+ positively rectified.  CeM sets Raw directly.  Note that a positive US onset even with no active Drive will be reflected here, enabling learning about unexpected outcomes
	GvCeMpos

	// GvCeMneg is negative valence central nucleus of the amygdala (CeM) LV (learned value) activity, reflecting |BLANegAcqD2 - BLANegExtD1|_+ positively rectified.  CeM sets Raw directly
	GvCeMneg

	/////////////////////////////////////////
	// VTA ventral tegmental area dopamine release

	// GvVtaDA is overall dopamine value reflecting all of the different inputs
	GvVtaDA

	/////////////////////////////////////////
	// USneg is negative valence US
	//   allocated for Nitems

	// GvUSneg is negative valence US outcomes -- NNegUSs of them
	GvUSneg

	///////////////////////////////////////////////////////////
	// Drives
	//   drives are allocated as a function of number of drives

	// GvDrives is current drive state -- updated with optional homeostatic exponential return to baseline values
	GvDrives

	// GvDrivesBase are baseline levels for each drive -- what they naturally trend toward in the absence of any input.  Set inactive drives to 0 baseline, active ones typically elevated baseline (0-1 range).
	GvBaseDrives

	// GvDrivesTau are time constants in ThetaCycle (trial) units for natural update toward Base values -- 0 values means no natural update.
	GvDrivesTau

	// GvDrivesUSDec are decrement factors for reducing drive value when Drive-US is consumed (multiply the US magnitude) -- these are positive valued numbers.
	GvDrivesUSDec

	///////////////////////////////////////////////////////////
	// USpos, VSPatch
	//   allocated as a function of number of drives

	// GvUSpos is current positive-valence drive-satisfying input(s) (unconditioned stimuli = US)
	GvUSpos

	// GvUSpos is current positive-valence drive-satisfying reward predicting VSPatch (PosD1) values
	GvVSPatch

	GlobalVarsN
)

//gosl: end globals
