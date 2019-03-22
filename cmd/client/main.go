package main

import (
	"flag"
)

type Options struct {
	offset int
	format string
	dataLength int
	arrayLength int
	requestCount int
	showDetails bool
	runAll bool
}

var Params Options

func main() {

	Params = parseOptions()

	benchmarks := RequestBenchmarks("http","http://localhost:9090/data/", &Params)

	benchmarks.DisplayBenchmarks("cmdline", Params.showDetails)

}

func parseOptions() Options {

	formatPtr := flag.String("f", "json", "Serialization format")
	dataLenPtr := flag.Int("d", 1000, "Data length (per array item)")
	arrayLenPtr := flag.Int("l", 1000, "Array length")
	offsetPtr := flag.Int("o", 1, "Offset Pointer")
	reqCountPtr := flag.Int("c", 10, "Number of requests to test")
	allOptionPtr := flag.Bool("a", false, "Benchmark all formats")
	detailsPtr := flag.Bool("details", false, "Display iteration details in benchmarks")

	flag.Parse()

	return Options  {
						format: *formatPtr,
						dataLength: *dataLenPtr,
						arrayLength: *arrayLenPtr,
						offset: *offsetPtr,
						runAll: *allOptionPtr,
						requestCount: *reqCountPtr,
						showDetails: *detailsPtr,
					}

}