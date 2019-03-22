package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type RespData []byte

var threadCount int

/**
 * HttpRequest executes an HTTP request (crazy, I know)
 */
func HttpRequest(baseUrl string, params *Options) ([]byte, uintptr) {

	// Builds request object
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+strconv.Itoa(params.offset), nil)
	if err != nil {
		log.Fatal("httprequest error: " + err.Error())
	}

	// Builds query string
	q := req.URL.Query()
	q.Add("format", params.format)
	q.Add("count", strconv.Itoa(params.arrayLength))
	q.Add("length", strconv.Itoa(params.dataLength))

	req.URL.RawQuery = q.Encode()

	// Retrieves response from server
	resp, err := client.Do(req)

	// Reads response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error: " + err.Error())
	}

	// Gets size of payload in bytes
	size := uintptr(len(body)) * reflect.TypeOf(body).Elem().Size()

	return body, size

}
