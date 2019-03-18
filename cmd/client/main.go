package main

type UrlParams struct {
	id int
	format string
	count int
	length int
}

var params = &UrlParams{10, "json", 100, 100}

func main() {

	DataRequest("http", params)

}