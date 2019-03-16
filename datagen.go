package main

import (
	lorem2 "github.com/drhodes/golorem"
	"log"
	"strconv"
	"time"
)

type RawData struct {
	Id string
	Data string
	Timestamp time.Time
}

func BuildRawData(id string, loremIpsum string) RawData {
	offset, err := strconv.Atoi(id); if err != nil { log.Fatal(err) }
	data := RawData{id, loremIpsum[offset+100: offset+600], time.Now()}
	return data
}

func GenerateData () string {
	loremIpsum := lorem2.Paragraph(20000, 20000)
	return loremIpsum
}
