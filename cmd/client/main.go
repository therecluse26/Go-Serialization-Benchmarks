package main

import (
	"flag"
)

// Options stores all CLI arguments
type Options struct {
	offset       int
	format       string
	dataLength   int
	arrayLength  int
	requestCount int
	showDetails  bool
	runAll       bool
	inputFile    string
	outputFile   string
	url          string
	compress     bool
}

// Params instantiates options and stores their values
var Params Options

func main() {
	// Snags options from the command line
	Params = parseOptions()

	// Sends requests to the server component
	benchmarks := RequestBenchmarks("http", Params.url, &Params)

	// Displays the benchmark results in the command line
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
	inputFilePtr := flag.String("in", "", "Path of an input file to serialize")
	outputFilePtr := flag.String("out", "", "Output file path of serialized input file")
	urlPtr := flag.String("u", "http://localhost:9090/data/", "Benchmark server url")
	compressPtr := flag.Bool("compress", false, "Compress the data packets for transmission")

	flag.Parse()

	validFmt := false
	for _, f := range Formats {
		if f == *formatPtr {
			validFmt = true
		}
	}
	if validFmt == false {
		panic("invalid format")
	}

	return Options{
		format:       *formatPtr,
		dataLength:   *dataLenPtr,
		arrayLength:  *arrayLenPtr,
		offset:       *offsetPtr,
		runAll:       *allOptionPtr,
		requestCount: *reqCountPtr,
		showDetails:  *detailsPtr,
		inputFile:    *inputFilePtr,
		outputFile:   *outputFilePtr,
		url:          *urlPtr,
		compress:     *compressPtr,
	}

}
