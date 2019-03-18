package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var Modes = map[string][]string{
	"http": {"rest", "rpc", "websocket"},
	"local": {"socket", "pipe"},
}

var Router mux.Router


func InitializeServer(mode string){

	if mode == "http" {
		StartHttpServer()
	}

}


func StartHttpServer(){

	Router.HandleFunc("/data/{id:[0-9]+}", HttpDataHandler).Queries("format", "{format}", "count", "{count:[0-9]+}", "length", "{length:[0-9]+}").Methods("GET")

	fmt.Println("Listening on port 9090")
	err := http.ListenAndServe(":9090", &Router)
	if err != nil {
		log.Fatal(err)
	}

}

func HttpDataHandler(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)

	id := params["id"]
	count, err := strconv.Atoi(r.FormValue("count")); if r.FormValue("count") == "" || err != nil { count = Default.DataLength}
	strLen, err := strconv.Atoi(r.FormValue("length")); if r.FormValue("length") == "" || err != nil { strLen = Default.DataLength}
	format := r.FormValue("format"); if format == "" { format = Default.Format}

	data := BuildRawData(id, LoremIpsum, count, strLen)

	formatted := FormatData(format, data)

	if format == "json" {
		w.Header().Set("Content-Type", "application/json")
	} else if format == "protobuf" {
		w.Header().Set("Content-Type", "application/octet-stream")
	} else if format == "flatbuf" {
		w.Header().Set("Content-Type", "binary/octet-stream")

	}
	_, errr := w.Write(formatted)
	if errr != nil {
		log.Fatal(err)
	}
}

func IpcSocket(){

}