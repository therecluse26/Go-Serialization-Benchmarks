package main

import (
	"encoding/json"
	"encoding/xml"
	"log"

	"../../schemas"
	"github.com/golang/protobuf/proto"
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
func (resp *RespData) ParseResponse(format string) {

	switch format {

	case "json":
		Parsed = resp.ParseJson()

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

func (data *RespData) ParseJson() RawData {
	var parsed RawData
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
