// Code generated by "stringer -type=NeuronVars"; DO NOT EDIT.

package axon

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Spike-0]
	_ = x[Spiked-1]
	_ = x[Act-2]
	_ = x[ActInt-3]
	_ = x[ActM-4]
	_ = x[ActP-5]
	_ = x[Ext-6]
	_ = x[Target-7]
	_ = x[Ge-8]
	_ = x[Gi-9]
	_ = x[Gk-10]
	_ = x[Inet-11]
	_ = x[Vm-12]
	_ = x[VmDend-13]
	_ = x[ISI-14]
	_ = x[ISIAvg-15]
	_ = x[CaSpkP-16]
	_ = x[CaSpkD-17]
	_ = x[CaSyn-18]
	_ = x[CaSpkM-19]
	_ = x[CaSpkPM-20]
	_ = x[CaLrn-21]
	_ = x[NrnCaM-22]
	_ = x[NrnCaP-23]
	_ = x[NrnCaD-24]
	_ = x[CaDiff-25]
	_ = x[Attn-26]
	_ = x[RLRate-27]
	_ = x[SpkMaxCa-28]
	_ = x[SpkMax-29]
	_ = x[SpkPrv-30]
	_ = x[SpkSt1-31]
	_ = x[SpkSt2-32]
	_ = x[GeNoiseP-33]
	_ = x[GeNoise-34]
	_ = x[GiNoiseP-35]
	_ = x[GiNoise-36]
	_ = x[GeExt-37]
	_ = x[GeRaw-38]
	_ = x[GeSyn-39]
	_ = x[GiRaw-40]
	_ = x[GiSyn-41]
	_ = x[GeInt-42]
	_ = x[GeIntMax-43]
	_ = x[GiInt-44]
	_ = x[GModRaw-45]
	_ = x[GModSyn-46]
	_ = x[GMaintRaw-47]
	_ = x[GMaintSyn-48]
	_ = x[SSGi-49]
	_ = x[SSGiDend-50]
	_ = x[Gak-51]
	_ = x[MahpN-52]
	_ = x[SahpCa-53]
	_ = x[SahpN-54]
	_ = x[GknaMed-55]
	_ = x[GknaSlow-56]
	_ = x[GnmdaSyn-57]
	_ = x[Gnmda-58]
	_ = x[GnmdaMaint-59]
	_ = x[GnmdaLrn-60]
	_ = x[NmdaCa-61]
	_ = x[GgabaB-62]
	_ = x[GABAB-63]
	_ = x[GABABx-64]
	_ = x[Gvgcc-65]
	_ = x[VgccM-66]
	_ = x[VgccH-67]
	_ = x[VgccCa-68]
	_ = x[VgccCaInt-69]
	_ = x[SKCaIn-70]
	_ = x[SKCaR-71]
	_ = x[SKCaM-72]
	_ = x[Gsk-73]
	_ = x[Burst-74]
	_ = x[BurstPrv-75]
	_ = x[CtxtGe-76]
	_ = x[CtxtGeRaw-77]
	_ = x[CtxtGeOrig-78]
	_ = x[NrnFlags-79]
	_ = x[NeuronVarsN-80]
}

const _NeuronVars_name = "SpikeSpikedActActIntActMActPExtTargetGeGiGkInetVmVmDendISIISIAvgCaSpkPCaSpkDCaSynCaSpkMCaSpkPMCaLrnNrnCaMNrnCaPNrnCaDCaDiffAttnRLRateSpkMaxCaSpkMaxSpkPrvSpkSt1SpkSt2GeNoisePGeNoiseGiNoisePGiNoiseGeExtGeRawGeSynGiRawGiSynGeIntGeIntMaxGiIntGModRawGModSynGMaintRawGMaintSynSSGiSSGiDendGakMahpNSahpCaSahpNGknaMedGknaSlowGnmdaSynGnmdaGnmdaMaintGnmdaLrnNmdaCaGgabaBGABABGABABxGvgccVgccMVgccHVgccCaVgccCaIntSKCaInSKCaRSKCaMGskBurstBurstPrvCtxtGeCtxtGeRawCtxtGeOrigNrnFlagsNeuronVarsN"

var _NeuronVars_index = [...]uint16{0, 5, 11, 14, 20, 24, 28, 31, 37, 39, 41, 43, 47, 49, 55, 58, 64, 70, 76, 81, 87, 94, 99, 105, 111, 117, 123, 127, 133, 141, 147, 153, 159, 165, 173, 180, 188, 195, 200, 205, 210, 215, 220, 225, 233, 238, 245, 252, 261, 270, 274, 282, 285, 290, 296, 301, 308, 316, 324, 329, 339, 347, 353, 359, 364, 370, 375, 380, 385, 391, 400, 406, 411, 416, 419, 424, 432, 438, 447, 457, 465, 476}

func (i NeuronVars) String() string {
	if i < 0 || i >= NeuronVars(len(_NeuronVars_index)-1) {
		return "NeuronVars(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NeuronVars_name[_NeuronVars_index[i]:_NeuronVars_index[i+1]]
}

func (i *NeuronVars) FromString(s string) error {
	for j := 0; j < len(_NeuronVars_index)-1; j++ {
		if s == _NeuronVars_name[_NeuronVars_index[j]:_NeuronVars_index[j+1]] {
			*i = NeuronVars(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: NeuronVars")
}

var _NeuronVars_descMap = map[NeuronVars]string{
	0:  `Spike is whether neuron has spiked or not on this cycle (0 or 1)`,
	1:  `Spiked is 1 if neuron has spiked within the last 10 cycles (msecs), corresponding to a nominal max spiking rate of 100 Hz, 0 otherwise -- useful for visualization and computing activity levels in terms of average spiked levels.`,
	2:  `Act is rate-coded activation value reflecting instantaneous estimated rate of spiking, based on 1 / ISIAvg. This drives feedback inhibition in the FFFB function (todo: this will change when better inhibition is implemented), and is integrated over time for ActInt which is then used for performance statistics and layer average activations, etc. Should not be used for learning or other computations.`,
	3:  `ActInt is integrated running-average activation value computed from Act with time constant Act.Dt.IntTau, to produce a longer-term integrated value reflecting the overall activation state across the ThetaCycle time scale, as the overall response of network to current input state -- this is copied to ActM and ActP at the ends of the minus and plus phases, respectively, and used in computing performance-level statistics (which are typically based on ActM). Should not be used for learning or other computations.`,
	4:  `ActM is ActInt activation state at end of third quarter, representing the posterior-cortical minus phase activation -- used for statistics and monitoring network performance. Should not be used for learning or other computations.`,
	5:  `ActP is ActInt activation state at end of fourth quarter, representing the posterior-cortical plus_phase activation -- used for statistics and monitoring network performance. Should not be used for learning or other computations.`,
	6:  `Ext is external input: drives activation of unit from outside influences (e.g., sensory input)`,
	7:  `Target is the target value: drives learning to produce this activation value`,
	8:  `Ge is total excitatory conductance, including all forms of excitation (e.g., NMDA) -- does *not* include Gbar.E`,
	9:  `Gi is total inhibitory synaptic conductance -- the net inhibitory input to the neuron -- does *not* include Gbar.I`,
	10: `Gk is total potassium conductance, typically reflecting sodium-gated potassium currents involved in adaptation effects -- does *not* include Gbar.K`,
	11: `Inet is net current produced by all channels -- drives update of Vm`,
	12: `Vm is membrane potential -- integrates Inet current over time`,
	13: `VmDend is dendritic membrane potential -- has a slower time constant, is not subject to the VmR reset after spiking`,
	14: `ISI is current inter-spike-interval -- counts up since last spike. Starts at -1 when initialized.`,
	15: `ISIAvg is average inter-spike-interval -- average time interval between spikes, integrated with ISITau rate constant (relatively fast) to capture something close to an instantaneous spiking rate. Starts at -1 when initialized, and goes to -2 after first spike, and is only valid after the second spike post-initialization.`,
	16: `CaSpkP is continuous cascaded integration of CaSpkM at PTau time constant (typically 40), representing neuron-level purely spiking version of plus, LTP direction of weight change and capturing the function of CaMKII in the Kinase learning rule. Used for specialized learning and computational functions, statistics, instead of Act.`,
	17: `CaSpkD is continuous cascaded integration CaSpkP at DTau time constant (typically 40), representing neuron-level purely spiking version of minus, LTD direction of weight change and capturing the function of DAPK1 in the Kinase learning rule. Used for specialized learning and computational functions, statistics, instead of Act.`,
	18: `CaSyn is spike-driven calcium trace for synapse-level Ca-driven learning: exponential integration of SpikeG * Spike at SynTau time constant (typically 30). Synapses integrate send.CaSyn * recv.CaSyn across M, P, D time integrals for the synaptic trace driving credit assignment in learning. Time constant reflects binding time of Glu to NMDA and Ca buffering postsynaptically, and determines time window where pre * post spiking must overlap to drive learning.`,
	19: `CaSpkM is spike-driven calcium trace used as a neuron-level proxy for synpatic credit assignment factor based on continuous time-integrated spiking: exponential integration of SpikeG * Spike at MTau time constant (typically 5). Simulates a calmodulin (CaM) like signal at the most abstract level.`,
	20: `CaSpkPM is minus-phase snapshot of the CaSpkP value -- similar to ActM but using a more directly spike-integrated value.`,
	21: `CaLrn is recv neuron calcium signal used to drive temporal error difference component of standard learning rule, combining NMDA (NmdaCa) and spiking-driven VGCC (VgccCaInt) calcium sources (vs. CaSpk* which only reflects spiking component). This is integrated into CaM, CaP, CaD, and temporal derivative is CaP - CaD (CaMKII - DAPK1). This approximates the backprop error derivative on net input, but VGCC component adds a proportion of recv activation delta as well -- a balance of both works best. The synaptic-level trace multiplier provides the credit assignment factor, reflecting coincident activity and potentially integrated over longer multi-trial timescales.`,
	22: `NrnCaM is integrated CaLrn at MTau timescale (typically 5), simulating a calmodulin (CaM) like signal, which then drives CaP, CaD for delta signal driving error-driven learning.`,
	23: `NrnCaP is cascaded integration of CaM at PTau time constant (typically 40), representing the plus, LTP direction of weight change and capturing the function of CaMKII in the Kinase learning rule.`,
	24: `NrnCaD is cascaded integratoin of CaP at DTau time constant (typically 40), representing the minus, LTD direction of weight change and capturing the function of DAPK1 in the Kinase learning rule.`,
	25: `CaDiff is difference between CaP - CaD -- this is the error signal that drives error-driven learning.`,
	26: `Attn is Attentional modulation factor, which can be set by special layers such as the TRC -- multiplies Ge`,
	27: `RLRate is recv-unit based learning rate multiplier, reflecting the sigmoid derivative computed from the CaSpkD of recv unit, and the normalized difference CaSpkP - CaSpkD / MAX(CaSpkP - CaSpkD).`,
	28: `SpkMaxCa is Ca integrated like CaSpkP but only starting at MaxCycStart cycle, to prevent inclusion of carryover spiking from prior theta cycle trial -- the PTau time constant otherwise results in significant carryover. This is the input to SpkMax`,
	29: `SpkMax is maximum CaSpkP across one theta cycle time window (max of SpkMaxCa) -- used for specialized algorithms that have more phasic behavior within a single trial, e.g., BG Matrix layer gating. Also useful for visualization of peak activity of neurons.`,
	30: `SpkPrv is final CaSpkD activation state at end of previous theta cycle. used for specialized learning mechanisms that operate on delayed sending activations.`,
	31: `SpkSt1 is the activation state at specific time point within current state processing window (e.g., 50 msec for beta cycle within standard theta cycle), as saved by SpkSt1() function. Used for example in hippocampus for CA3, CA1 learning`,
	32: `SpkSt2 is the activation state at specific time point within current state processing window (e.g., 100 msec for beta cycle within standard theta cycle), as saved by SpkSt2() function. Used for example in hippocampus for CA3, CA1 learning`,
	33: `GeNoiseP is accumulating poisson probability factor for driving excitatory noise spiking -- multiply times uniform random deviate at each time step, until it gets below the target threshold based on lambda.`,
	34: `GeNoise is integrated noise excitatory conductance, added into Ge`,
	35: `GiNoiseP is accumulating poisson probability factor for driving inhibitory noise spiking -- multiply times uniform random deviate at each time step, until it gets below the target threshold based on lambda.`,
	36: `GiNoise is integrated noise inhibotyr conductance, added into Gi`,
	37: `GeExt is extra excitatory conductance added to Ge -- from Ext input, GeCtxt etc`,
	38: `GeRaw is raw excitatory conductance (net input) received from senders = current raw spiking drive`,
	39: `GeSyn is time-integrated total excitatory synaptic conductance, with an instantaneous rise time from each spike (in GeRaw) and exponential decay with Dt.GeTau, aggregated over projections -- does *not* include Gbar.E`,
	40: `GiRaw is raw inhibitory conductance (net input) received from senders = current raw spiking drive`,
	41: `GiSyn is time-integrated total inhibitory synaptic conductance, with an instantaneous rise time from each spike (in GiRaw) and exponential decay with Dt.GiTau, aggregated over projections -- does *not* include Gbar.I. This is added with computed FFFB inhibition to get the full inhibition in Gi`,
	42: `GeInt is integrated running-average activation value computed from Ge with time constant Act.Dt.IntTau, to produce a longer-term integrated value reflecting the overall Ge level across the ThetaCycle time scale (Ge itself fluctuates considerably) -- useful for stats to set strength of connections etc to get neurons into right range of overall excitatory drive`,
	43: `GeIntMax is maximum GeInt value across one theta cycle time window.`,
	44: `GiInt is integrated running-average activation value computed from GiSyn with time constant Act.Dt.IntTau, to produce a longer-term integrated value reflecting the overall synaptic Gi level across the ThetaCycle time scale (Gi itself fluctuates considerably) -- useful for stats to set strength of connections etc to get neurons into right range of overall inhibitory drive`,
	45: `GModRaw is raw modulatory conductance, received from GType = ModulatoryG projections`,
	46: `GModSyn is syn integrated modulatory conductance, received from GType = ModulatoryG projections`,
	47: `GMaintRaw is raw maintenance conductance, received from GType = MaintG projections`,
	48: `GMaintSyn is syn integrated maintenance conductance, integrated using MaintNMDA params.`,
	49: `SSGi is SST+ somatostatin positive slow spiking inhibition`,
	50: `SSGiDend is amount of SST+ somatostatin positive slow spiking inhibition applied to dendritic Vm (VmDend)`,
	51: `Gak is conductance of A-type K potassium channels`,
	52: `MahpN is accumulating voltage-gated gating value for the medium time scale AHP`,
	53: `SahpCa is slowly accumulating calcium value that drives the slow AHP`,
	54: `SahpN is sAHP gating value`,
	55: `GknaMed is conductance of sodium-gated potassium channel (KNa) medium dynamics (Slick) -- produces accommodation / adaptation of firing`,
	56: `GknaSlow is conductance of sodium-gated potassium channel (KNa) slow dynamics (Slack) -- produces accommodation / adaptation of firing`,
	57: `GnmdaSyn is integrated NMDA recv synaptic current -- adds GeRaw and decays with time constant`,
	58: `Gnmda is net postsynaptic (recv) NMDA conductance, after Mg V-gating and Gbar -- added directly to Ge as it has the same reversal potential`,
	59: `GnmdaMaint is net postsynaptic maintenance NMDA conductance, computed from GMaintSyn and GMaintRaw, after Mg V-gating and Gbar -- added directly to Ge as it has the same reversal potential`,
	60: `GnmdaLrn is learning version of integrated NMDA recv synaptic current -- adds GeRaw and decays with time constant -- drives NmdaCa that then drives CaM for learning`,
	61: `NmdaCa is NMDA calcium computed from GnmdaLrn, drives learning via CaM`,
	62: `GgabaB is net GABA-B conductance, after Vm gating and Gbar + Gbase -- applies to Gk, not Gi, for GIRK, with .1 reversal potential.`,
	63: `GABAB is GABA-B / GIRK activation -- time-integrated value with rise and decay time constants`,
	64: `GABABx is GABA-B / GIRK internal drive variable -- gets the raw activation and decays`,
	65: `Gvgcc is conductance (via Ca) for VGCC voltage gated calcium channels`,
	66: `VgccM is activation gate of VGCC channels`,
	67: `VgccH inactivation gate of VGCC channels`,
	68: `VgccCa is instantaneous VGCC calcium flux -- can be driven by spiking or directly from Gvgcc`,
	69: `VgccCaInt time-integrated VGCC calcium flux -- this is actually what drives learning`,
	70: `SKCaIn is intracellular calcium store level, available to be released with spiking as SKCaR, which can bind to SKCa receptors and drive K current. replenishment is a function of spiking activity being below a threshold`,
	71: `SKCaR released amount of intracellular calcium, from SKCaIn, as a function of spiking events. this can bind to SKCa channels and drive K currents.`,
	72: `SKCaM is Calcium-gated potassium channel gating factor, driven by SKCaR via a Hill equation as in chans.SKPCaParams.`,
	73: `Gsk is Calcium-gated potassium channel conductance as a function of Gbar * SKCaM.`,
	74: `Burst is 5IB bursting activation value, computed by thresholding regular CaSpkP value in Super superficial layers`,
	75: `BurstPrv is previous Burst bursting activation from prior time step -- used for context-based learning`,
	76: `CtxtGe is context (temporally delayed) excitatory conductance, driven by deep bursting at end of the plus phase, for CT layers.`,
	77: `CtxtGeRaw is raw update of context (temporally delayed) excitatory conductance, driven by deep bursting at end of the plus phase, for CT layers.`,
	78: `CtxtGeOrig is original CtxtGe value prior to any decay factor -- updates at end of plus phase.`,
	79: `NrnFlags are bit flags for binary state variables, which are converted to / from uint32. These need to be in Vars because they can be differential per data (for ext inputs) and are writable (indexes are read only).`,
	80: ``,
}

func (i NeuronVars) Desc() string {
	if str, ok := _NeuronVars_descMap[i]; ok {
		return str
	}
	return "NeuronVars(" + strconv.FormatInt(int64(i), 10) + ")"
}
