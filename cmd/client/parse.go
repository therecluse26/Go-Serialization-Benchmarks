package main

import (
	"../../schemas"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"log"
)

type RawData struct {
	Id string
	Data map[int32]string
	Timestamp int64
}

func (resp *RespData) ParseResponse (format string) {

	var parsed RawData

	if format == "json" {
		err := json.Unmarshal(*resp, &parsed)
		if err != nil {
			log.Fatal("parsing error: " + err.Error())
		}

	} else if format == "flatbuffers" {
		schemas.GetRootAsLoremFb(*resp, 0)

	} else if format == "protobuf" {
		lorem := &schemas.ProtobufLorem{}
		err := proto.Unmarshal(*resp, lorem)
		if err != nil {
			log.Fatalln("parsing error: ", err)
		}

	}

}

