// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kinase

// SynParams has rate constants for averaging over activations
// at different time scales, to produce the running average activation
// values that then drive learning.
type SynParams struct {
	Rule     Rules   `desc:"which learning rule to use"`
	OptInteg bool    `desc:"use the optimized spike-only integration of cascaded CaM, CaP, CaD values -- iterates cascaded updates between spikes."`
	SpikeG   float32 `def:"10,36" desc:"spiking gain factor for all algos -- only alters the overall range of values, keeping them in roughly the unit scale."`
	MTau     float32 `def:"10" min:"1" desc:"CaM mean running-average time constant in cycles, reflecting calmodulin and total Ca influx biologically, which should be milliseconds typically (tau is roughly how long it takes for value to change significantly -- 1.4x the half-life). This provides a pre-integration step before integrating into the CaP value"`
	PTau     float32 `def:"40" min:"1" desc:"LTP Ca-driven factor time constant in cycles, reflecting CaMKII dynamics biologically, should be milliseconds typically (tau is roughly how long it takes for value to change significantly -- 1.4x the half-life). Continuously updates based on current CaM value, resulting in faster tracking of plus-phase signals."`
	DTau     float32 `def:"40" min:"1" desc:"LTD Ca-driven factor time constant in cycles, reflecting DAPK1 dynamics biologically, should be milliseconds typically (tau is roughly how long it takes for value to change significantly -- 1.4x the half-life).  Continuously updates based on current CaP value, resulting in slower integration that still reflects earlier minus-phase signals."`
	DScale   float32 `def:"1,0.93,1.05" desc:"scaling factor on CaD as it enters into the learning rule, to compensate for systematic decrease in activity over the course of a theta cycle"`

	MDt float32 `view:"-" json:"-" xml:"-" inactive:"+" desc:"rate = 1 / tau"`
	PDt float32 `view:"-" json:"-" xml:"-" inactive:"+" desc:"rate = 1 / tau"`
	DDt float32 `view:"-" json:"-" xml:"-" inactive:"+" desc:"rate = 1 / tau"`
}

func (kp *SynParams) Defaults() {
	kp.Rule = SynSpkCa
	kp.OptInteg = false
	kp.SpikeG = 36
	kp.MTau = 10 // was 40 -- 10 matches Urakubo NMDA
	kp.PTau = 40 // 10
	kp.DTau = 40
	kp.DScale = 1 // 0.93, 1.05
	kp.Update()
}

func (kp *SynParams) Update() {
	kp.MDt = 1 / kp.MTau
	kp.PDt = 1 / kp.PTau
	kp.DDt = 1 / kp.DTau
}

// FmSpike computes updates from current spike value, for
// continuously-updating mode
func (kp *SynParams) FmSpike(spike float32, caM, caP, caD *float32) {
	*caM += kp.MDt * (kp.SpikeG*spike - *caM)
	*caP += kp.PDt * (*caM - *caP)
	*caD += kp.DDt * (*caP - *caD)
}

// FmCa computes updates from current calcium level, for
// continuously-updating mode
func (kp *SynParams) FmCa(ca float32, caM, caP, caD *float32) {
	*caM += kp.MDt * (ca - *caM)
	*caP += kp.PDt * (*caM - *caP)
	*caD += kp.DDt * (*caP - *caD)
}

// DWt computes the weight change from CaP, CaD values
func (kp *SynParams) DWt(caP, caD float32) float32 {
	return caP - kp.DScale*caD
}

// ISIFmTime returns the inter spike interval from current time
// and last spike time, which is -1 if never spiked
// (in which case return ISI is -1)
func (kp *SynParams) ISIFmTime(ctime, stime int32) int {
	if stime < 0 {
		return -1
	}
	return int(ctime - stime)
}

// CurCa returns the current Ca* values, dealing with updating for
// optimized spike-time update versions.
// ctime is current time in msec, and stime is last spike time (-1 if never)
func (kp *SynParams) CurCa(ctime, stime int32, caM, caP, caD float32) (cCaM, cCaP, cCaD float32) {
	// isi := kp.ISIFmTime(ctime, stime)
	// if !kp.OptInteg || isi < 0 {
	return caM, caP, caD
	// }
	// cCaM = caM * mat32.FastExp(-float32(isi)/(kp.MTau-0.5)) // 0.5 factor makes it fit perfectly..
	// cCaP = kp.PFmLastSpike(caP, caM, isi)
	// cCaD = kp.DFmLastSpike(caD, caP, caM, isi)
	// return
}

// FuntCaFmSpike updates the function-table-based Ca values based on a spike
// having occured (spk > 0) -- can pass a Ca value for spk instead of 1
// for NMDA-based case.
// ctime is current time in msec, and stime is last spike time (-1 if never)
// func (kp *SynParams) FuntCaFmSpike(ctime int32, stime *int32, spk float32, caM, caP, caD *float32) {
// 	isi := kp.ISIFmTime(ctime, *stime)
// 	var mprv float32
//
// 	// get old before cam update, for previous isi
// 	if isi >= 0 {
// 		*caD = kp.DFmLastSpike(*caD, *caP, *caM, isi) // update in reverse order
// 		*caP = kp.PFmLastSpike(*caP, *caM, isi)
// 		mprv = *caM * mat32.FastExp(-float32(isi)/(kp.MTau-0.5))
// 	}
// 	*caM = mprv + kp.MDt*(kp.SpikeG*spk-mprv)
// 	*stime = ctime
// }
