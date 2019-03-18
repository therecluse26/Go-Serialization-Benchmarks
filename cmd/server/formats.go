package main

import (
	"./schemas"
	"github.com/gin-gonic/gin/json"
	"github.com/golang/protobuf/proto"
	"github.com/google/flatbuffers/go"
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
	"CapnProto",
	"Avro",
	"CBOR",

}

func FormatData(format string, data RawData) []byte {

	if format == "json" {
		return FormatJson(data)
	} else if format == "flatbuf" {
		return FormatFlatbuf(data)
	} else if format == "protobuf" {
		return FormatProtobuf(data)
	}

	return []byte{}

}

func FormatJson(data RawData) []byte {

	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return jsonString
}

func FormatFlatbuf(data RawData) []byte {

	b := &flatbuffers.Builder{}
	b.Reset()

	count := len(data.Data)

	// Initializes data map
	for idx, tx := range data.Data {
		txData := b.CreateString(tx)
		schemas.LoremFbDataMapStart(b)
		schemas.LoremFbDataMapAddIndex(b, idx)
		schemas.LoremFbDataMapAddText(b, txData)
		schemas.LoremFbDataMapEnd(b)
	}

	schemas.LoremFbStartDataVector(b, count)
	dataMap := b.EndVector(count)

	loremId := b.CreateByteString([]byte(data.Id))
	schemas.LoremFbStart(b)
	schemas.LoremFbAddId(b, loremId)
	schemas.LoremFbAddData(b, dataMap)
	schemas.LoremFbAddTimestamp(b, data.Timestamp)
	loremPosition := schemas.LoremFbEnd(b)

	b.Finish(loremPosition)

	return b.Bytes[b.Head():]

}

func FormatProtobuf(data RawData) []byte {

	protoData := &schemas.ProtobufLorem{
		Id: &data.Id,
		Data: data.Data,
		Timestamp: &data.Timestamp,
	}

	protoFmt, err := proto.Marshal(protoData)
	if err != nil {
		log.Fatal(err)
	}

	return protoFmt

}