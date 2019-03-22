package modes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"../data"
	"../conf"
)

var Router mux.Router

func StartHttpServer(){

	Router.HandleFunc("/data/{id:[0-9]+}", HttpDataHandler).Queries("format", "{format}", "count", "{count:[0-9]+}", "length", "{length:[0-9]+}").Methods("GET")

	fmt.Println("Listening on port 9090")

	err := http.ListenAndServe(":9090", &Router)
	if err != nil {
		log.Fatal(err)
	}

}

func Wrapper(dat []byte, err error) []byte {
	if err != nil {
		log.Fatal(err)
	}
	return dat
}

func HttpDataHandler(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)

	id := params["id"]
	count, err := strconv.Atoi(r.FormValue("count"))
		if r.FormValue("count") == "" || err != nil { count = conf.Default.DataLength}

	strLen, err := strconv.Atoi(r.FormValue("length"))
		if r.FormValue("length") == "" || err != nil { strLen = conf.Default.DataLength}

	format := r.FormValue("format")
		if format == "" { format = conf.Default.Format }

	switch format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
	case "xml":
		w.Header().Set("Content-Type", "application/xml")
	case "protobuf":
		w.Header().Set("Content-Type", "application/octet-stream")
	default:
		w.Header().Set("Content-Type", "binary/octet-stream")
	}

	_, err = w.Write(
				Wrapper(
					data.FormatData(
						format,
						data.BuildRawData(id, data.LoremIpsum, count, strLen))))
	if err != nil {
		log.Fatal(err)
	}
}

func IpcSocket(){

}