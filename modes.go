package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Modes = map[string][]string{
	"http": {"rest", "rpc", "websocket"},
	"local": {"socket", "pipe"},
}

var Router mux.Router


func SetTransMode(mode string, format string){

	if mode == "http" {
		StartHttpServer(format)
	}

}



func StartHttpServer(format string){


	/*router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Access-Control-Allow-Origin, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.Header().Set("Access-Control-Allow-Origin", "localhost:9090")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusNoContent)
		return
	})*/

	Router.Headers("fmt", format)

	Router.HandleFunc("/data/{id:[0-9]+}", HttpDataHandler).Methods("GET")

	fmt.Println("Listening on port 9090")
	err := http.ListenAndServe(":9090", &Router)
	if err != nil {
		log.Fatal(err)
	}


}

func HttpDataHandler(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)

	id := params["id"]

	format := r.Header.Get("fmt")

	data := BuildRawData(id, LoremIpsum)

	formatted := FormatData(format, data)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(formatted)
	if err != nil {
		log.Fatal(err)
	}
}

func IpcSocket(){

}