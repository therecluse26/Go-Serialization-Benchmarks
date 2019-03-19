package main

type UrlParams struct {
	id string
	format string
	arrayCount string
	length string
}

var params = &UrlParams{"1", "flatbuffers", "1000", "100"}

func main() {

	//fmt.Println(flag)

	reqCount := 1000

	benchmarks := RequestBenchmarks("http","http://localhost:9090/data/", params, reqCount)

	benchmarks.DisplayBenchmarks("cmdline", false)

}

func parseFlags() {

}