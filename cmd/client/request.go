package main

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type RespData []byte

var ThreadCount int

/**
 * Builds Benchmarks based on requests
 */
func RequestBenchmarks(reqType string, url string, params *Options) BenchAggregate {

	fmt.Println("Running " + params.format + " benchmarks...")

	// Progress bar
	progress := pb.New(params.requestCount).SetMaxWidth(100).SetRefreshRate(10000)
	progress.Start()

	// Creates new benchmarks object
	benchmarks := BenchAggregate{Format: params.format, RequestType: reqType, Iterations:  params.requestCount, DataArrayCount: params.arrayLength, DataEntryLength: params.dataLength }

	ThreadCount = 0

	for i := 1; i <= params.requestCount; i++ {

		benchmarks.SingleBenchmark(reqType, url, params, progress)

		/*go func(idx int, aggregate *BenchAggregate, req string, u string, p *Options, pr *pb.ProgressBar) {
			fmt.Println(idx)
			aggregate.SingleBenchmark(req, u, p, idx, pr)
		}(i, &benchmarks, reqType, url, params, progress)*/

	}

	benchmarks.AggregateData()

	progress.FinishPrint("Complete!")

	return benchmarks

}

func (benchmarks *BenchAggregate) SingleBenchmark(reqType string, url string, params *Options, progress *pb.ProgressBar){

	var bench = Benchmark{StartTime: time.Now()}
	var resp = RespData{}
	var dataSize uintptr
	if reqType == "http" {
		resp, dataSize = HttpRequest(url, params)
	}
	resp.ParseResponse(params.format)
	bench.EndTime = time.Now()
	bench.TotalExecTime = time.Since(bench.StartTime)
	benchmarks.PayloadSize = dataSize
	benchmarks.Benchmarks = append(benchmarks.Benchmarks, bench)

	// Increments progress bar
	progress.Increment()
}


func HttpRequest(baseUrl string, params *Options) ([]byte, uintptr) {

	// Builds request object
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl + strconv.Itoa(params.offset), nil)
	if err != nil {
		log.Fatal("httprequest error: " + err.Error())
	}

	// Builds query string
	q := req.URL.Query()
	q.Add("format", params.format)
	q.Add("count", strconv.Itoa(params.arrayLength))
	q.Add("length", strconv.Itoa(params.dataLength))

	req.URL.RawQuery = q.Encode()

	// Retrieves response from server
	resp, err := client.Do(req)

	// Reads response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error: " + err.Error())
	}

	// Gets size of payload in bytes
	size := uintptr(len(body)) * reflect.TypeOf(body).Elem().Size()

	return body, size

}