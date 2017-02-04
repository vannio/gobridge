package main

import (
  "fmt"
  "strings"
  "net/http"
  "encoding/json"

  "github.com/vannio/gobridge/week-2/animals"
)

// Global scope because its accessed within multiple functions
var pets []animals.Pet

func ListPets(res http.ResponseWriter, req *http.Request) {
  data, err := json.Marshal(pets)

  // Deal with error immediately
  // Better quality code than using try-catch
  if err != nil {
    res.WriteHeader(http.StatusInternalServerError)
    return
  }

  res.Write(data)
}

func LikePet(res http.ResponseWriter, req *http.Request) {
  name := req.URL.Query().Get("name")

  if (len(name) == 0) {
    fmt.Fprint(res, "Specify a pet name to like!")
    return
  }

  for _, pet := range pets {
    // Case insensitive name-matching
    if (strings.ToLower(pet.GetName()) == strings.ToLower(name)) {
      pet.AddLike()

      fmt.Fprint(res,
        "Ooh-la-la, you liked " + pet.GetName() + ". ",
        pet.GetName() + " now has ",
        pluraliseLikesCount(pet.GetLikes()),
      )

      // Quit iterating as soon as it matches a name
      // and finishes what it needs to do
      return
    }
  }

  // If there have been no matches, the `return` above won't be called
  // So the program should continue to this line
  fmt.Fprint(res, "Can't find pet with name " + name)
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

  http.HandleFunc("/", ListPets) // The root route
  http.HandleFunc("/list", ListPets)
  http.HandleFunc("/like", LikePet)

  fmt.Println("Listening for connections on port", 9000)
  http.ListenAndServe(":9000", http.DefaultServeMux)
}
