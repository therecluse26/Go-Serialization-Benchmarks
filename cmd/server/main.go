package main

import (
	"flag"
	"fmt"

	"github.com/therecluse26/Go-Serialization-Benchmarks/cmd/server/conf"
	"github.com/therecluse26/Go-Serialization-Benchmarks/cmd/server/modes"
)

// Options stores all CLI arguments
type Options struct {
	port       int
}

// Params instantiates options and stores their values
var Params Options

func init() {
	fmt.Println("Initializing Server...")
}

func main() {
	opts := parseOptions()
	InitializeServer(conf.Default.Mode, opts.port)
}

func InitializeServer(mode string, port int) {

	if mode == "http" {
		modes.StartHttpServer(port)
	}

}

func parseOptions() Options {

	portPtr := flag.Int("p", 9090, "Specifies which port to run the server on")

	flag.Parse()

	return Options{
		port:       *portPtr,
	}
}