package main

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/therecluse26/Go-Serialization-Benchmarks/schemas"
	capnp "zombiezen.com/go/capnproto2"
)

type RawData struct {
	ID        string
	Data      map[int32]string
	Timestamp int64
}

var Parsed RawData

/**
 * ParseResponse parses the response from the server
 * Important for overall benchmark to include both encoding/decoding
 */
func (resp *RespData) ParseResponse(format string, compress bool) {

	switch format {

	case "json":
		Parsed = resp.ParseJson(compress)
	case "xml":
		Parsed = resp.ParseXML()
	case "flatbuffers":
		schemas.GetRootAsLoremFb(*resp, 0)
	case "protobuf":
		Parsed = resp.ParseProtoBuf()
	case "capnproto":
		Parsed = resp.ParseCapnProto()
	}

}

func (data *RespData) ParseJson(compress bool) RawData {
	var parsed RawData
	//fmt.Println(*data)
	var reader io.Reader

	if compress == true {

		gz, err := gzip.NewReader(reader)
		_, err = gz.Read(*data)
		if err != nil {
			log.Fatal("error with gzip reader: " + err.Error())
		}
		defer gz.Close()
		reader = gz

		/*var buf bytes.Buffer
		err := gunzipWrite(&buf, *data)
		if err != nil {
			log.Fatal(err)
		}
		_, err = buf.Write(*data)*/

		//fmt.Println(*data)
	}

	err := json.Unmarshal(*data, &parsed)
	if err != nil {
		log.Fatal("parsing error: " + err.Error())
	}
	return parsed
}

func (data *RespData) ParseXML() RawData {
	var parsed RawData
	err := xml.Unmarshal(*data, &parsed)
	if err != nil {
		log.Fatal("parsing error: " + err.Error())
	}
	return parsed
}

func (data *RespData) ParseCapnProto() RawData {

	_, err := capnp.Unmarshal(*data)
	if err != nil {
		log.Fatal(err)
	}

	//dat, _ :=result.Arena.Data(0)

	return RawData{}
}

func (data *RespData) ParseProtoBuf() RawData {

	lorem := &schemas.ProtobufLorem{}
	err := proto.Unmarshal(*data, lorem)
	if err != nil {
		log.Fatalln("parsing error: ", err)
	}
	return RawData{*lorem.Id, lorem.Data, *lorem.Timestamp}
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	/*gr, err := gzip.NewReader(bytes.NewBuffer(data))
	gr.Header("asdf", "asdf")
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)*/
	return nil
}
