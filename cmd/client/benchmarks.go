package main

import (
	"fmt"
	"strconv"
	"time"
	"code.cloudfoundry.org/bytefmt"
)

// Individual benchmark (per request)
type Benchmark struct {
	StartTime time.Time
	EndTime time.Time
	TotalExecTime time.Duration
}

// Aggregate of all benchmarks per test
type BenchAggregate struct {
	Format string
	RequestType string
	Iterations int
	DataEntryLength int
	DataArrayCount int
	PayloadSize uintptr
	TotalExecTime time.Duration
	AvgExecTime time.Duration
	Benchmarks []Benchmark
}

/**
 * Aggregates all benchmark data
 */
func (benchmarks *BenchAggregate) AggregateData() {

	var totalDur time.Duration

	for _, bench := range benchmarks.Benchmarks {
		totalDur += bench.TotalExecTime
	}

	benchmarks.TotalExecTime = totalDur
	benchmarks.AvgExecTime = totalDur / time.Duration(benchmarks.Iterations)
}

/**
 * Formats Benchmarks for output display
 */
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