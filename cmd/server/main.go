package main

import (
	"./conf"
	"./modes"
	"fmt"
)

func init() {
	fmt.Println("Initializing Server...")
}

func main(){
	InitializeServer(conf.Default.Mode)
}

func InitializeServer(mode string){

	if mode == "http" {
		modes.StartHttpServer()
	}

}