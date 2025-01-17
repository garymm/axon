// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fffb

// Bg has parameters for a slower, low level of background inhibition
// based on main FFFB computed inhibition.
type Bg struct {

	// enable adaptive layer inhibition gain as stored in layer GiCur value
	On bool `desc:"enable adaptive layer inhibition gain as stored in layer GiCur value"`

	// [def: .1] [viewif: On=true] level of inhibition as proporition of FFFB Gi value -- will need to reduce FFFB level to compensate for this additional source of inhibition
	Gi float32 `def:".1" viewif:"On=true" desc:"level of inhibition as proporition of FFFB Gi value -- will need to reduce FFFB level to compensate for this additional source of inhibition"`

	// [def: 10] [viewif: On=true] time constant for integrating background inhibition (tau is roughly how long it takes for value to change significantly -- 1.4x the half-life)
	Tau float32 `def:"10" viewif:"On=true" desc:"time constant for integrating background inhibition (tau is roughly how long it takes for value to change significantly -- 1.4x the half-life)"`

	// [view: -] rate = 1 / tau
	Dt float32 `inactive:"+" view:"-" json:"-" xml:"-" desc:"rate = 1 / tau"`
}

func (bg *Bg) Update() {
	bg.Dt = 1 / bg.Tau
}

func (bg *Bg) Defaults() {
	bg.On = false
	bg.Gi = 0.1
	bg.Tau = 10
	bg.Update()
}

// GiBg updates the gi background value from Gi FFFB computed value
func (bg *Bg) GiBg(gibg *float32, gi float32) {
	if !bg.On {
		*gibg = 0
		return
	}
	*gibg += bg.Dt * (bg.Gi*gi - *gibg)
}
