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

func main() {
	res, err := http.Get("http://localhost:9000/list")

	if err != nil {
		log.Fatal(err)
	}

	// Doesn't read from body, only saves a reference of response body
	body := res.Body

	// ReadAll expects an io.Reader argument
	// Luckily, res.Body is an io.ReadCloser interface,
	// which includes Reader:
	// type ReadCloser interface {
	//   Reader
	//   Closer
	// }
	data, err := ioutil.ReadAll(body)

	if err != nil {
		log.Fatal(err)
	}

	// We want to read JSON and save it
	// Source data is an array/slice of pets
	// `var pets Pets` is the same as `var pets []Pet`
	pets := Pets{}

	// We want Unmarshal to modify pet so we pass in a reference
	err = json.Unmarshal(data, &pets)

	if err != nil {
		log.Fatal(err)
	}

	// Sort method need the argument to have some methods attached
	// See lines 66 onwards
	sort.Sort(pets)

	// loop through pets and print each one on a different line
	for _, pet := range pets {
		log.Println(pet.Name, "likes", pet.Hobbies)
	}
}

// Pets : A type alias
type Pets []Pet

// Sort function needs these methods attached:
// Len to check the length of slice
func (a Pets) Len() int {
	return len(a)
}

// Less to compare 2 items
func (a Pets) Less(i, j int) bool {
	return a[i].Name > a[j].Name
}

// Swap to swap 2 items
func (a Pets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
