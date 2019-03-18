package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DataRequest(reqType string, params *UrlParams){

	if reqType == "http" {
		HttpRequest("http://localhost:9090/data", *params)
	}

}

func HttpRequest(baseUrl string, params UrlParams) []byte {

	client := &http.Client{}

	resp, err := client.Get("http://localhost:9090/data/100?format=json&count=100&length=100")
	if err != nil {
		log.Fatal("httprequest error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	err = resp.Body.Close()
	if err != nil {
		log.Fatal("error: " + err.Error())
	}

	return body

}