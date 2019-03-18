package main

import (
	lorem2 "github.com/drhodes/golorem"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type RawData struct {
	Id string
	Data map[int32]string
	Timestamp int64
}

func BuildRawData(id string, loremIpsum string, count int, length int) RawData {

	offset, err := strconv.Atoi(id); if err != nil { log.Fatal(err) }

	dataSlice := make(map[int32]string, count)

	for i := 0; i < count; i++ {
		random := rand.Intn(100)
		dataSlice[int32(i)] = loremIpsum[offset*random: offset*random+length]
	}

	data := RawData{id, dataSlice, time.Now().Unix()}
	return data
}

func GenerateData () string {
	loremIpsum := lorem2.Paragraph(200000, 200000)
	return loremIpsum
}
