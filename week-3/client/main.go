package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:9000/list")

	if err != nil {
		log.Fatal(err)
	}

	// Doesn't read from body, only saves a reference of response body
	a := res.Body
	// a is interface io.ReadCloser:
	// type ReadCloser interface {
	//   Reader
	//   Closer
	// }

	data, err := ioutil.ReadAll(a)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
}
