package data

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/google/flatbuffers/go"
	"github.com/therecluse26/Go-Serialization-Benchmarks/schemas"
	"zombiezen.com/go/capnproto2"
)

// Formats supported
var Formats = []string{
	"json",
	"xml",
	"protobuf",
	"flatbuffers",
	"capnproto",
	"messagepack",
	"avro",
	"cbor",
}

// FormatData invokes encoding functions by format
// and returns encoded byte slice
func FormatData(format string, data RawData, compress bool) ([]byte, error) {

	switch format {
	case "json":
		return FormatJSON(data, compress)
	/*case "xml":
		return FormatXML(data, compress)*/
	case "flatbuffers":
		return FormatFlatbuf(data, compress)
	case "protobuf":
		return FormatProtobuf(data, compress)
	case "capnproto":
		return FormatCapnProto(data, compress)
	default:
		return []byte{}, errors.New("no valid format given")
	}
}

// FormatXML encodes result set as XML
/*func FormatXML(data RawData, compress bool) ([]byte, error) {

	for i, d := range data.Data {
		d.MarshalXML(data, "", " ")
	}

	xmlData, err := xml.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(xmlData)

	return xmlData, err

} */

// FormatJSON encodes result set as JSON
func FormatJSON(data RawData, compress bool) ([]byte, error) {

	if compress == true {
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		dat, _ := json.Marshal(data)
		json.NewEncoder(gz).Encode(&dat)
		gz.Close()
		return []byte(dat), nil

	}

	return json.Marshal(data)
}

// FormatProtobuf encodes result set as Protocol Buffer
func FormatProtobuf(data RawData, compress bool) ([]byte, error) {

	if compress == true {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(data)
		gz.Close()
	} else {

		return proto.Marshal(&schemas.ProtobufLorem{
			Id:        &data.ID,
			Data:      data.Data,
			Timestamp: &data.Timestamp,
		})
	}

	return []byte{}, errors.New("no protobuf returned")
}

// FormatFlatbuf encodes result set as Flatbuffers
func FormatFlatbuf(data RawData, compress bool) ([]byte, error) {
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

	loremID := b.CreateByteString([]byte(data.ID))
	schemas.LoremFbStart(b)
	schemas.LoremFbAddId(b, loremID)
	schemas.LoremFbAddData(b, dataMap)
	schemas.LoremFbAddTimestamp(b, data.Timestamp)
	loremPosition := schemas.LoremFbEnd(b)

	b.Finish(loremPosition)

	return b.Bytes[b.Head():], nil

}

// FormatCapnProto encodes result set as Cap'n Proto
func FormatCapnProto(data RawData, compress bool) ([]byte, error) {
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
	if dataRoot.SetId(data.ID) != nil {
		log.Fatal(err)
	}

	for idx, tx := range data.Data {

		dataMap, err := schemas.NewRootLoremCpDataMap(seg)
		if err != nil {
			log.Fatal(err)
		}
		dataMap.SetIndex(idx)
		if dataMap.SetText(tx) != nil {
			log.Fatal(err)
		}

	}

	return msg.Marshal()
}
