package main

import (
	"fmt"
	"strconv"
	"time"

	"code.cloudfoundry.org/bytefmt"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// Benchmark stores individual benchmark data (per request)
type Benchmark struct {
	StartTime     time.Time
	EndTime       time.Time
	TotalExecTime time.Duration
}

// BenchAggregate stores an aggregate of all benchmarks per test
type BenchAggregate struct {
	Format          string
	RequestType     string
	Iterations      int
	DataEntryLength int
	DataArrayCount  int
	PayloadSize     uintptr
	TotalExecTime   time.Duration
	AvgExecTime     time.Duration
	Benchmarks      []Benchmark
}

// AggregateData totals all benchmark data and preps for output
func (benchmarks *BenchAggregate) AggregateData() {

	var totalDur time.Duration

	for _, bench := range benchmarks.Benchmarks {
		totalDur += bench.TotalExecTime
	}

	benchmarks.TotalExecTime = totalDur
	benchmarks.AvgExecTime = totalDur / time.Duration(benchmarks.Iterations)
}

//SingleBenchmark executes a request and measures the time
func (benchmarks *BenchAggregate) SingleBenchmark(reqType string, url string, params *Options, progress *pb.ProgressBar) {

	var bench = Benchmark{StartTime: time.Now()}
	var resp = RespData{}
	var dataSize uintptr

	// Executes the request
	if reqType == "http" {
		resp, dataSize = HttpRequest(url, params)
	}

	// Parses the response
	resp.ParseResponse(params.format)

	bench.EndTime = time.Now()
	bench.TotalExecTime = time.Since(bench.StartTime)
	benchmarks.PayloadSize = dataSize
	benchmarks.Benchmarks = append(benchmarks.Benchmarks, bench)

	// Increments progress bar
	progress.Increment()
}

// RequestBenchmarks builds benchmarks based on requests
func RequestBenchmarks(reqType string, url string, params *Options) BenchAggregate {

	fmt.Println("Running " + params.format + " benchmarks...")

	// Progress bar
	progress := pb.New(params.requestCount).SetMaxWidth(100).SetRefreshRate(10000)
	progress.Start()

	// Creates new benchmarks object
	benchmarks := BenchAggregate{Format: params.format,
		RequestType:     reqType,
		Iterations:      params.requestCount,
		DataArrayCount:  params.arrayLength,
		DataEntryLength: params.dataLength}

	threadCount = 0

	for i := 1; i <= params.requestCount; i++ {

		benchmarks.SingleBenchmark(reqType, url, params, progress)
	}

	benchmarks.AggregateData()

	progress.FinishPrint("Complete!")

	return benchmarks

}

// DisplayBenchmarks formats benchmarks for output display
func (benchmarks *BenchAggregate) DisplayBenchmarks(format string, showIterations bool) {

	dataEntryLength := strconv.Itoa(benchmarks.DataEntryLength)
	dataArrayCount := strconv.Itoa(benchmarks.DataArrayCount)

	if format == "cmdline" {
		fmt.Print("----------------------------------\n" +
			"BENCHMARK RESULTS \n" +
			"----------------------------------\n")
		fmt.Println("Request Type: " + benchmarks.RequestType)
		fmt.Println("Format: " + benchmarks.Format)
		fmt.Println("Data Array Items: " + dataArrayCount)
		fmt.Println("Data Entry Characters: " + dataEntryLength)
		fmt.Println("Payload Size: " + bytefmt.ByteSize(uint64(benchmarks.PayloadSize)))
		fmt.Println("Iterations: " + strconv.Itoa(benchmarks.Iterations))
		fmt.Println("Total Execution Time: " + benchmarks.TotalExecTime.String())
		fmt.Println("Average Execution Time: " + benchmarks.AvgExecTime.String())

		if showIterations == true {
			fmt.Print("----------------------------------\n" +
				"ITERATIONS \n" +
				"----------------------------------\n")
			for idx, bench := range benchmarks.Benchmarks {
				fmt.Println("Iteration #" + strconv.Itoa(idx+1) + ": " + bench.TotalExecTime.String())
			}
		}
	}
}
