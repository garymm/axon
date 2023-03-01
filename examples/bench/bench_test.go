package bench

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"testing"

	"github.com/emer/axon/axon"
	"github.com/emer/emergent/etime"
	"github.com/emer/etable/etable"
	"github.com/goki/gi/gi"
)

func init() {
	// must lock main thread for gpu!
	runtime.LockOSThread()
}

const (
	convergenceTestEpochs = 10
	defaultNumEpochs      = 350
)

var gpu = flag.Bool("gpu", false, "whether to run gpu or not")
var maxProcs = flag.Int("maxProcs", 0, "GOMAXPROCS value to set -- 0 = use current default")
var threadsNeuron = flag.Int("thrNeuron", 0, "number of goroutines to launch for NeuronFun")
var threadsSendSpike = flag.Int("thrSendSpike", 0, "number of goroutines to launch for SendSpike")
var threadsSynCa = flag.Int("thrSynCa", 0, "number of goroutines to launch for SynCa")
var numEpochs = flag.Int("epochs", defaultNumEpochs, "number of epochs to run")
var numPats = flag.Int("pats", 10, "number of patterns per epoch")
var numUnits = flag.Int("units", 100, "number of units per layer -- uses NxN where N = sqrt(units)")
var verbose = flag.Bool("verbose", true, "if false, only report the final time")
var writeStats = flag.Bool("writestats", false, "whether to write network stats to a CSV file")

func BenchmarkBenchNetFull(b *testing.B) {
	if *maxProcs > 0 {
		runtime.GOMAXPROCS(*maxProcs)
	}

	if *verbose {
		fmt.Printf("Running bench with: %d neuronThreads, %d sendSpikeThreads, %d synCaThreads, %d epochs, %d pats, %d units\n", *threadsNeuron, *threadsSendSpike, *threadsSynCa, *numEpochs, *numPats, *numUnits)
	}

	rand.Seed(42)

	net := &axon.Network{}
	ConfigNet(net, *threadsNeuron, *threadsSendSpike, *threadsSynCa, *numUnits, *verbose)
	if *verbose {
		log.Println(net.SizeReport())
	}

	pats := &etable.Table{}
	ConfigPats(pats, *numPats, *numUnits)

	epcLog := &etable.Table{}
	ConfigEpcLog(epcLog)

	TrainNet(net, pats, epcLog, *numEpochs, *verbose, *gpu)

	if *writeStats {
		filename := fmt.Sprintf("bench_%d_units.csv", *numUnits)
		err := epcLog.SaveCSV(gi.FileName(filename), ',', etable.Headers)
		if err != nil {
			b.Log(err)
		}
	}

	if *numEpochs < defaultNumEpochs {
		if *verbose {
			b.Logf("skipping convergence test because numEpochs < %d", defaultNumEpochs)
		}
		return
	}
	corSimSum := 0.0
	for epoch := *numEpochs - convergenceTestEpochs; epoch < *numEpochs; epoch++ {
		corSimSum += epcLog.CellFloat("CorSim", epoch)
		if math.IsNaN(corSimSum) {
			b.Errorf("CorSim for epoch %d is NaN", epoch)
		}
	}
	corSimAvg := corSimSum / float64(convergenceTestEpochs)
	if corSimAvg < 0.90 {
		b.Errorf("average of CorSim for last %d epochs too low. Want %v, got %v", convergenceTestEpochs, 0.95, corSimAvg)
	}
}

// Run just the threading benchmarks with `go test -bench=".*Thread.*" .`
func benchmarkNeuronFunMultiThread(numThread, numUnits int, b *testing.B) {
	// this benchmark constructs the network just like `bench.go`, but without
	// setting up the projections (not needed for benching NeuronFun) -> Test setup is much quicker.
	net := &axon.Network{}
	net.InitName(net, "BenchNet")

	squn := int(math.Sqrt(float64(numUnits)))
	shp := []int{squn, squn}

	net.AddLayer("Input", shp, axon.InputLayer)
	net.AddLayer("Hidden1", shp, axon.SuperLayer)
	net.AddLayer("Hidden2", shp, axon.SuperLayer)
	net.AddLayer("Hidden3", shp, axon.SuperLayer)
	net.AddLayer("Output", shp, axon.TargetLayer)

	net.RecFunTimes = true

	// builds with default threads
	if err := net.Build(); err != nil {
		panic(err)
	}
	net.Defaults()
	if _, err := net.ApplyParams(ParamSets[0].Sheets["Network"], false); err != nil {
		panic(err)
	}

	// override defaults: neurons, sendSpike, synCa, learn
	if err := net.Threads.Set(numThread, 1, 1); err != nil {
		b.Error(err)
	}

	net.InitWts()
	ctx := axon.NewContext()

	// reset timer to avoid counting setup time
	b.ResetTimer()

	// timing seems to correspond well to the real benchmark, where we run the whole network
	// For the real benchmark: Look at the profile generated by run_bench.sh, find out how much time is spent in
	// NeuronFun and divide that by (epochs * pats * quarters * cycles)
	for i := 0; i < b.N; i++ {
		ctx.NewState(etime.Train)
		net.NeuronFun(func(ly axon.AxonLayer, ni uint32, nrn *axon.Neuron) { ly.CycleNeuron(ctx, ni, nrn) },
			"CycleNeuron")
	}
}

const (
	smallNumUnits = 2048       // 5 * 2048 * 80 * 4B = 3MB (should fit in the cache)
	hugeNumUnits  = 256 * 2048 // 5 * 256 * 2048 * 80 * 4B = 786MB (should not fit in the cache)
)

// Get a profile with `go test -bench=".*Thread.*" . -test.cpuprofile=neuronFun_T1.prof`
func BenchmarkNeuronFun1ThreadsSmall(b *testing.B) {
	benchmarkNeuronFunMultiThread(1, smallNumUnits, b)
}

func BenchmarkNeuronFun2ThreadsSmall(b *testing.B) {
	benchmarkNeuronFunMultiThread(2, smallNumUnits, b)
}
func BenchmarkNeuronFun4ThreadsSmall(b *testing.B) {
	benchmarkNeuronFunMultiThread(4, smallNumUnits, b)
}
func BenchmarkNeuronFun8ThreadsSmall(b *testing.B) {
	benchmarkNeuronFunMultiThread(8, smallNumUnits, b)
}

func BenchmarkNeuronFun1ThreadsBig(b *testing.B) {
	benchmarkNeuronFunMultiThread(1, hugeNumUnits, b)
}

func BenchmarkNeuronFun2ThreadsBig(b *testing.B) {
	benchmarkNeuronFunMultiThread(2, hugeNumUnits, b)
}

func BenchmarkNeuronFun4ThreadsBig(b *testing.B) {
	benchmarkNeuronFunMultiThread(4, hugeNumUnits, b)
}

func BenchmarkNeuronFun8ThreadsBig(b *testing.B) {
	benchmarkNeuronFunMultiThread(8, hugeNumUnits, b)
}

// store to global to avoid compiler optimization
var fp32Result float32

// Benchmark the cost of doing a type assert on a layer
func BenchmarkLayerTypeAssert(b *testing.B) {
	net := &axon.Network{}
	ConfigNet(net, 1, 1, 1, 2048, false)
	tmp := float32(0.0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// there'll be 5 layers total
		for _, ly := range net.Layers {
			lyr := ly.(axon.AxonLayer)
			// do something with lyr st the compiler cannot optimize it away
			tmp += lyr.AsAxon().Neurons[0].Spike
		}
	}
	// avoid compiler optimization
	fp32Result = tmp
}

// Benchmark cost of not doing the type assertion, for comparison
func BenchmarkLayerTypeAssertBaseline(b *testing.B) {
	net := &axon.Network{}
	ConfigNet(net, 1, 1, 1, 2048, false)
	tmp := float32(0.0)
	layers := make([]axon.AxonLayer, len(net.Layers))
	// pre-convert all the layers to AxonLayer
	for i, ly := range net.Layers {
		layers[i] = ly.(axon.AxonLayer)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, ly := range layers {
			tmp += ly.AsAxon().Neurons[0].Spike
		}
	}
	// avoid compiler optimization
	fp32Result = tmp
}
