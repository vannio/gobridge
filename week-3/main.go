package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/vannio/gobridge/week-2/animals"
)

var pets []animals.Pet

// ListPets : lists all the pets instantiated in JSON format
func ListPets(res http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(pets)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(data)
}

// LikePet : handles liking a pet
func LikePet(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if len(name) == 0 {
		fmt.Fprint(res, "Specify a pet name to like!")
		return
	}

	for _, pet := range pets {
		if strings.ToLower(pet.GetName()) == strings.ToLower(name) {
			pet.AddLike()

			fmt.Fprint(res,
				"Ooh-la-la, you liked "+pet.GetName()+". ",
				pet.GetName()+" now has ",
				pluraliseLikesCount(pet.GetLikes()),
			)

			return
		}
	}

	fmt.Fprint(res, "Can't find pet with name "+name)
}

func pluraliseLikesCount(likes int) string {
	if likes > 1 {
		return fmt.Sprint(likes) + " likes!"
	}
	return fmt.Sprint(likes) + " like!"
}

func main() {
	pets = []animals.Pet{
		&animals.Kitten{
			Name: "Ms Tiggles",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		&animals.Kitten{
			Name: "Mr Tom",
			Hobbies: []string{
				"Chasing own tail",
				"Napping on cushions",
				"Eating",
			},
		},
		&animals.Dog{
			Name: "Fido",
			Hobbies: []string{
				"Barking",
				"Eating",
			},
		},
	}

	http.HandleFunc("/", ListPets)
	http.HandleFunc("/list", ListPets)
	http.HandleFunc("/like", LikePet)

	port := ":9000"

	fmt.Println("Listening for connections at localhost" + port)
	http.ListenAndServe(port, http.DefaultServeMux)
}
