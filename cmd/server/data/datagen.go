package data

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	lorem2 "github.com/drhodes/golorem"
)

// RawData holds the request data structure
type RawData struct {
	ID        string
	Data      map[int32]string
	Timestamp int64
}

// LoremIpsum stores a huge block of Lorem Ipsum to memory
var LoremIpsum = GenerateData()

func BuildRawData(id string, loremIpsum string, count int, length int) RawData {

	offset, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	dataSlice := make(map[int32]string, count)

	for i := 0; i < count; i++ {
		random := rand.Intn(100)
		dataSlice[int32(i)] = loremIpsum[offset*random : offset*random+length]
	}

	data := RawData{id, dataSlice, time.Now().Unix()}
	return data
}

// GenerateData generates a huge block of Lorem Ipsum text
func GenerateData() string {
	return lorem2.Paragraph(200000, 200000)
}
