package main

import (
	"github.com/gin-gonic/gin/json"
	"log"
)

const (
	Json = "JSON"
	Xml = "XML"
	ProtoBuf = "Protocol Buffers"
	FlatBuf = "FlatBuffers"
)

var Formats = []string{
	"JSON",
	"XML",
	"ProtoBuf",
	"FlatBuf",
	"Avro",
	"CBOR",

}

func FormatData(format string, data RawData) []byte {

	return FormatJson(data)

}

func FormatJson(data RawData) []byte {

	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return jsonString
}