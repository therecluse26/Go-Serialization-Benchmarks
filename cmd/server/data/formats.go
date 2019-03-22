package data

import (
	"../../../schemas"
	"encoding/xml"
	"errors"
	"github.com/gin-gonic/gin/json"
	"github.com/golang/protobuf/proto"
	"github.com/google/flatbuffers/go"
	"log"
	"zombiezen.com/go/capnproto2"
)

var Formats = []string{
	"JSON",
	"XML",
	"ProtoBuf",
	"FlatBuffers",
	"CapnProto",
	"MessagePack",
	"Avro",
	"CBOR",
}

/**
 * FormatData invokes encoding functions by format
 * and returns encoded byte slice
 */
func FormatData(format string, data RawData) ([]byte, error) {

	switch format {
	case "json":
		return FormatJson(data)
	case "xml":
		return FormatXML(data)
	case "flatbuffers":
		return FormatFlatbuf(data)
	case "protobuf":
		return FormatProtobuf(data)
	case "capnproto":
		return FormatCapnProto(data)
	default:
		return []byte{}, errors.New("no vali format given")
	}
}

/**
 * FormatXML encodes result set as XML
 */
func FormatXML(data RawData) ([]byte, error) {

	xmlData, err := xml.Marshal(data)
	if err != nil {
		panic("HEY")
	}

	return xmlData, err

}

/**
 * FormatJson encodes result set as JSON
 */
func FormatJson(data RawData) ([]byte, error) {
	return json.Marshal(data)
}

/**
 * FormatProtobuf encodes result set as Protocol Buffer
 */
func FormatProtobuf(data RawData) ([]byte, error) {
	return proto.Marshal(&schemas.ProtobufLorem{
		Id: &data.Id,
		Data: data.Data,
		Timestamp: &data.Timestamp,
	})
}

/**
 * FormatFlatbuf encodes result set as Flatbuffers
 */
func FormatFlatbuf(data RawData) ([]byte, error) {
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

	return b.Bytes[b.Head():], nil

}

/**
 * FormatCapnProto encodes result set as Cap'n Proto
 */
func FormatCapnProto(data RawData) ([]byte, error) {
	// New empty message for structs
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		log.Fatal(err)
	}

	// New root element
	dataRoot, err := schemas.NewRootLoremCp(seg)
	if err != nil {
		log.Fatal(err)
	}

	dataRoot.SetTimestamp(data.Timestamp)
	if dataRoot.SetId(data.Id) != nil { log.Fatal(err) }

	for idx, tx := range data.Data {

		dataMap, err := schemas.NewRootLoremCpDataMap(seg)
		if err != nil {
			log.Fatal(err)
		}
		dataMap.SetIndex(idx)
		if dataMap.SetText(tx) != nil { log.Fatal(err) }

	}

	return msg.Marshal()
}