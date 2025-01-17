// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package armaze

import "github.com/goki/ki/kit"

// Paradigms is a list of experimental paradigms that
// govern the configuration and updating of environment
// state over time and the appropriate evaluation criteria.
type Paradigms int

const (
	// Approach is a basic case where one Drive (chosen at random each trial) is fully active and others are at InactiveDrives levels -- goal is to approach the CS associated with the Drive-satisfying US, and avoid negative any negative USs.  USs are always placed in same Arms (NArms must be >= NUSs -- any additional Arms are filled at random with additional US copies)
	Approach Paradigms = iota

	ParadigmsN
)

//go:generate stringer -type=Paradigms

var KiT_Paradigms = kit.Enums.AddEnum(ParadigmsN, kit.NotBitFlag, nil)

///////////////////////////////////////////////
// Approach

// ConfigApproach does initial config for Approach paradigm
func (ev *Env) ConfigApproach() {
	if ev.Config.NArms < ev.Config.NUSs {
		ev.Config.NArms = ev.Config.NUSs
	}
	if ev.Config.NCSs < ev.Config.NUSs {
		ev.Config.NCSs = ev.Config.NUSs
	}
}

// StartApproach does new start state setting for Approach
// Selects a new TrgDrive at random, sets that to 1,
// others to inactive levels
func (ev *Env) StartApproach() {
	ev.TrgDrive = ev.Rand.Intn(ev.Config.NDrives, -1)
	for i := range ev.Drives {
		if i == ev.TrgDrive {
			ev.Drives[i] = 1
		} else {
			ev.Drives[i] = ev.InactiveVal()
		}
	}
}

func (ev *Env) StepApproach() {

}
