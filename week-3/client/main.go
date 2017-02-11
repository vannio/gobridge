package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Pet : a subset of animals
type Pet struct {
	Name    string
	Hobbies []string
	Likes   int
}

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

	// We want to read JSON and save it
	// Source data is an array/slice of pets
	pets := []Pet{}

	// We want Unmarshal to modify pet so we pass in a reference
	err = json.Unmarshal(data, &pets)

	if err != nil {
		log.Fatal(err)
	}

	// json.Unmarshal needs a slice of bytes (data) and a 'schema' in form of a blank interface
	// interface can be anything(?) - lowest form of object
	// Could be a value, could be a reference
	log.Println(pets)
}
