// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package axon

//gosl: start layervals

// ActAvgVals are long-running-average activation levels stored in the LayerVals,
// for monitoring and adapting inhibition and possibly scaling parameters.
// All of these integrate over NData within a network, so are the same across them.
type ActAvgVals struct {

	// running-average minus-phase activity integrated at Dt.LongAvgTau -- used for adapting inhibition relative to target level
	ActMAvg float32 `inactive:"+" desc:"running-average minus-phase activity integrated at Dt.LongAvgTau -- used for adapting inhibition relative to target level"`

	// running-average plus-phase activity integrated at Dt.LongAvgTau
	ActPAvg float32 `inactive:"+" desc:"running-average plus-phase activity integrated at Dt.LongAvgTau"`

	// running-average max of minus-phase Ge value across the layer integrated at Dt.LongAvgTau
	AvgMaxGeM float32 `inactive:"+" desc:"running-average max of minus-phase Ge value across the layer integrated at Dt.LongAvgTau"`

	// running-average max of minus-phase Gi value across the layer integrated at Dt.LongAvgTau
	AvgMaxGiM float32 `inactive:"+" desc:"running-average max of minus-phase Gi value across the layer integrated at Dt.LongAvgTau"`

	// multiplier on inhibition -- adapted to maintain target activity level
	GiMult float32 `inactive:"+" desc:"multiplier on inhibition -- adapted to maintain target activity level"`

	// adaptive threshold -- only used for specialized layers, e.g., VSPatch
	AdaptThr float32 `inactive:"+" desc:"adaptive threshold -- only used for specialized layers, e.g., VSPatch"`

	pad, pad1 float32
}

func (lv *ActAvgVals) Init() {
	lv.ActMAvg = 0 // will be set to ly.Params.InhibActAvg.Nominal in InitWts
	lv.ActPAvg = 0 // will be set to ly.Params.InhibActAvg.Nominal in InitWts
	lv.AvgMaxGeM = 1
	lv.AvgMaxGiM = 1
	lv.GiMult = 1
	lv.AdaptThr = 0 // will be initialized per-user type
}

// CorSimStats holds correlation similarity (centered cosine aka normalized dot product)
// statistics at the layer level
type CorSimStats struct {

	// correlation (centered cosine aka normalized dot product) activation difference between ActP and ActM on this alpha-cycle for this layer -- computed by CorSimFmActs called by PlusPhase
	Cor float32 `inactive:"+" desc:"correlation (centered cosine aka normalized dot product) activation difference between ActP and ActM on this alpha-cycle for this layer -- computed by CorSimFmActs called by PlusPhase"`

	// running average of correlation similarity between ActP and ActM -- computed with CorSim.Tau time constant in PlusPhase
	Avg float32 `inactive:"+" desc:"running average of correlation similarity between ActP and ActM -- computed with CorSim.Tau time constant in PlusPhase"`

	// running variance of correlation similarity between ActP and ActM -- computed with CorSim.Tau time constant in PlusPhase
	Var float32 `inactive:"+" desc:"running variance of correlation similarity between ActP and ActM -- computed with CorSim.Tau time constant in PlusPhase"`

	pad float32
}

func (cd *CorSimStats) Init() {
	cd.Cor = 0
	cd.Avg = 0
	cd.Var = 0
}

// LaySpecialVals holds special values used to communicate to other layers
// based on neural values, used for special algorithms such as RL where
// some of the computation is done algorithmically.
type LaySpecialVals struct {

	// one value
	V1 float32 `inactive:"+" desc:"one value"`

	// one value
	V2 float32 `inactive:"+" desc:"one value"`

	// one value
	V3 float32 `inactive:"+" desc:"one value"`

	// one value
	V4 float32 `inactive:"+" desc:"one value"`
}

func (lv *LaySpecialVals) Init() {
	lv.V1 = 0
	lv.V2 = 0
	lv.V3 = 0
	lv.V4 = 0
}

// LayerVals holds extra layer state that is updated per layer.
// It is sync'd down from the GPU to the CPU after every Cycle.
type LayerVals struct {

	// [view: -] layer index for these vals
	LayIdx uint32 `view:"-" desc:"layer index for these vals"`

	// [view: -] data index for these vals
	DataIdx uint32 `view:"-" desc:"data index for these vals"`

	// reaction time for this layer in cycles, which is -1 until the Max CaSpkP level (after MaxCycStart) exceeds the Act.Attn.RTThr threshold
	RT  float32 `inactive:"-" desc:"reaction time for this layer in cycles, which is -1 until the Max CaSpkP level (after MaxCycStart) exceeds the Act.Attn.RTThr threshold"`
	pad uint32

	// [view: inline] running-average activation levels used for adaptive inhibition, and other adapting values
	ActAvg ActAvgVals `view:"inline" desc:"running-average activation levels used for adaptive inhibition, and other adapting values"`

	// correlation (centered cosine aka normalized dot product) similarity between ActM, ActP states
	CorSim CorSimStats `desc:"correlation (centered cosine aka normalized dot product) similarity between ActM, ActP states"`

	// [view: inline] special values used to communicate to other layers based on neural values computed on the GPU -- special cross-layer computations happen CPU-side and are sent back into the network via Context on the next cycle -- used for special algorithms such as RL / DA etc
	Special LaySpecialVals `view:"inline" desc:"special values used to communicate to other layers based on neural values computed on the GPU -- special cross-layer computations happen CPU-side and are sent back into the network via Context on the next cycle -- used for special algorithms such as RL / DA etc"`
}

func (lv *LayerVals) Init() {
	lv.ActAvg.Init()
	lv.CorSim.Init()
	lv.Special.Init()
}

//gosl: end layervals
