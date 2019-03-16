package main

import (
	"fmt"
	"os"
)

var Mode = "http"
var Format = "json"
var LoremIpsum = GenerateData()

func main(){

	for i := range os.Args {
		fmt.Println(i)
	}

	SetTransMode(Mode, Format)


}
