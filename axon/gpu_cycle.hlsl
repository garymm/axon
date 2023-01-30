// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// performs the CycleNeuron function on all neurons

#include "context.hlsl"
#include "layerparams.hlsl"
#include "prjnparams.hlsl"

// note: binding is var, set

// Set 0: uniforms -- these are constant
[[vk::binding(0, 0)]] uniform LayerParams Layers[]; // [Layer]
[[vk::binding(1, 0)]] uniform PrjnParams Prjns[]; // [Layer][RecvPrjns]

// Set 1: main network structs and vals
[[vk::binding(0, 1)]] StructuredBuffer<Context> Ctxt; // [0]
[[vk::binding(1, 1)]] RWStructuredBuffer<Neuron> Neurons; // [Layer][Neuron]
[[vk::binding(2, 1)]] RWStructuredBuffer<Pool> Pools; // [Layer][Pools]
[[vk::binding(3, 1)]] RWStructuredBuffer<LayerVals> LayVals; // [Layer]
// [[vk::binding(4, 1)]] RWStructuredBuffer<Synapse> Synapses;  // [Layer][RecvPrjns][RecvNeurons][Syns]
// [[vk::binding(5, 1)]] RWStructuredBuffer<float> GBuf;  // [Layer][RecvPrjns][RecvNeurons][MaxDel+1]
// [[vk::binding(6, 1)]] RWStructuredBuffer<float> GSyns;  // [Layer][RecvPrjns][RecvNeurons]

// Set 2: prjn, synapse level indexes and buffer values
// [[vk::binding(0, 2)]] StructuredBuffer<StartN> RecvCon; // [Layer][RecvPrjns][RecvNeurons]

void PulvinarDriver(in LayerParams ly, in LayerParams dly, in Pool lpl, uint ni, uint nin, out float drvGe, out float nonDrvPct) {
	float drvMax = lpl.AvgMax.CaSpkP.Cycle.Max;
	nonDrvPct = ly.Pulv.NonDrivePct(drvMax); // how much non-driver to keep
	uint pnin = ni + ly.Idxs.NeurSt;
	drvGe = ly.Pulv.DriveGe(Neurons[pnin].Burst);
}

// GInteg integrates conductances G over time (Ge, NMDA, etc).
// calls NeuronGatherSpikes, GFmRawSyn, GiInteg
void GInteg(in Context ctx, in LayerParams ly, uint ni, uint nin, inout Neuron nrn, in Pool pl, in LayerVals vals, inout uint2 randctr) {
	float drvGe = 0;
	float nonDrvPct = 0;
	if (ly.LayType == PulvinarLayer) {
		PulvinarDriver(ly, Layers[ly.Pulv.DriveLayIdx], Pools[Layers[ly.Pulv.DriveLayIdx].Idxs.PoolSt], ni, nin, drvGe, nonDrvPct);
	}

	float saveVal = ly.SpecialPreGs(ctx, ni, nrn, drvGe, nonDrvPct, randctr);
	
	ly.GFmRawSyn(ctx, ni, nrn, randctr);
	ly.GiInteg(ctx, ni, nrn, pl, vals);
	ly.GNeuroMod(ctx, ni, nrn, vals);
	
	ly.SpecialPostGs(ctx, ni, nrn, randctr, 0);
}

void CycleNeuron2(in Context ctx, in LayerParams ly, uint nin, inout Neuron nrn, in Pool pl, in Pool lpl, in LayerVals vals) {
	uint ni = nin - ly.Idxs.NeurSt; // layer-based as in Go
	uint2 randctr = ctx.RandCtr.Uint2();
	
	GInteg(ctx, ly, ni, nin, nrn, pl, vals, randctr);
	ly.SpikeFmG(ctx, ni, nrn);
	ly.PostSpikeSpecial(ctx, ni, nrn, pl, lpl, vals);
	ly.PostSpike(ctx, ni, nrn, pl, vals);
}

void CycleNeuron(in Context ctx, uint ni, inout Neuron nrn) {
	CycleNeuron2(ctx, Layers[nrn.LayIdx], ni, nrn, Pools[nrn.SubPoolN], Pools[Layers[nrn.LayIdx].Idxs.PoolSt], LayVals[nrn.LayIdx]);
}

[numthreads(64, 1, 1)]
void main(uint3 idx : SV_DispatchThreadID) {  // over Recv Neurons
	uint ns;
	uint st;
	Neurons.GetDimensions(ns, st);
	if (idx.x < ns) {
		CycleNeuron(Ctxt[0], idx.x, Neurons[idx.x]);
	}
}
