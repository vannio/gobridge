package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// Pet : a subset of animals
type Pet struct {
	Name    string
	Hobbies []string
	Likes   int
}

// Pets : A type alias
type Pets []Pet

// Example of bubble sorting algorithm:
// c, b, a

// is c 'bigger than' b?
// b, c, a

// is c 'bigger than' a?
// b, a, c

// is c 'bigger than' b?
// a, b, c

func (a Pets) Len() int {
	return len(a)
}

func (a Pets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Pets) Less(i, j int) bool {
	return a[i].Name > a[j].Name
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
	pets := Pets{}

	// var pets Pets
	// is the same as:
	// var pets []Pet

	// We want Unmarshal to modify pet so we pass in a reference
	err = json.Unmarshal(data, &pets)

	// json.Unmarshal needs a slice of bytes (data) and a 'schema' in form of a blank interface
	// interface can be anything(?) - lowest form of object
	// Could be a value, could be a reference

	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(pets)

	// CHALLENGE
	// loop through pets and print each one on a different line
	for _, pet := range pets {
		log.Println(pet.Name, "likes", pet.Hobbies)
	}
}
