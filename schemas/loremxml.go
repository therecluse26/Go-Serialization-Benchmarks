package schemas

import (
	"encoding/xml"
	"strconv"
)

// LoremMap is a map[int32]string.
type LoremMap map[int32]string

// Lorem is the main LoremIpsum struct
type Lorem struct {
	ID        string   `xml:"ID,omitempty"`
	Data      LoremMap `xml:"Data,omitempty"`
	Timestamp int64    `xml:"Timestamp,omitempty"`
}

// MarshalXML marshals into XML.
func (s LoremMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", strconv.Itoa(int(key))}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}
