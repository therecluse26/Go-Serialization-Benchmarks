package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type RespData []byte

/**
 * Builds Benchmarks based on requests
 */
func RequestBenchmarks(reqType string, url string, params *UrlParams, reqCount int) BenchAggregate {

	dataCount, err := strconv.Atoi(params.arrayCount); if err != nil { log.Fatal("parameter parsing error: " + err.Error()) }
	dataLength, err := strconv.Atoi(params.length); if err != nil { log.Fatal("parameter parsing error: " + err.Error()) }

	benchmarks := BenchAggregate{Format: params.format, RequestType: reqType, Iterations:reqCount, DataArrayCount: dataCount, DataEntryLength: dataLength }

	for i := 1; i <= reqCount; i++ {
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
	}

	benchmarks.AggregateData()

	return benchmarks

}

func HttpRequest(baseUrl string, params *UrlParams) ([]byte, uintptr) {

	// Builds request object
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+params.id, nil)
	if err != nil {
		log.Fatal("httprequest error: " + err.Error())
	}

	// Builds query string
	q := req.URL.Query()
	q.Add("format", params.format)
	q.Add("count", params.arrayCount)
	q.Add("length", params.length)

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