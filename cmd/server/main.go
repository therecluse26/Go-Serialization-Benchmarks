package main

import (
	"./conf"
	"./modes"
)

func main(){
	InitializeServer(conf.Default.Mode)
}

func InitializeServer(mode string){

	if mode == "http" {
		modes.StartHttpServer()
	}

}