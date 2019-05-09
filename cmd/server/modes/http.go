package modes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/therecluse26/Go-Serialization-Benchmarks/cmd/server/conf"
	"github.com/therecluse26/Go-Serialization-Benchmarks/cmd/server/data"
)

var Router mux.Router

func StartHttpServer(port int) {

	var portStr = strconv.Itoa(port)

	Router.HandleFunc("/data/{id:[0-9]+}", HttpDataHandler).Queries("format", "{format}", "count", "{count:[0-9]+}", "length", "{length:[0-9]+}").Methods("GET")

	fmt.Println("Listening on port "+portStr)

	err := http.ListenAndServe(":"+portStr, &Router)
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

func HttpDataHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]
	count, err := strconv.Atoi(r.FormValue("count"))
	if r.FormValue("count") == "" || err != nil {
		count = conf.Default.DataLength
	}

	strLen, err := strconv.Atoi(r.FormValue("length"))
	if r.FormValue("length") == "" || err != nil {
		strLen = conf.Default.DataLength
	}

	format := r.FormValue("format")
	if format == "" {
		format = conf.Default.Format
	}

	compress, err := strconv.ParseBool(r.FormValue("compress"))
	if r.FormValue("compress") == "" || err != nil {
		compress = conf.Default.Compression
	}

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

	if compress == true {
		w.Header().Set("Content-Encoding", "gzip")
	}

	_, err = w.Write(
		Wrapper(
			data.FormatData(
				format,
				data.BuildRawData(id, data.LoremIpsum, count, strLen), compress)))
	if err != nil {
		//log.Fatal(err)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
	}
}

func IpcSocket() {

}
