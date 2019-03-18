package main

var LoremIpsum string

func init(){
	LoremIpsum = GenerateData()
}

func main(){
	InitializeServer(Default.Mode)
}
