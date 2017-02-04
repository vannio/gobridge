package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/vannio/gobridge/week-2/animals"
)

// Global scope because its used in both ListKittens() and main()
var kittens []animals.Kitten

func ListKittens(res http.ResponseWriter, req *http.Request) {
  data, err := json.Marshal(kittens)

  // Deal with error immediately
  // Better quality code than using try catch
  if err != nil {
    res.WriteHeader(http.StatusInternalServerError)
    return
  }

  res.Write(data)
}

func main() {
  kittens = []animals.Kitten{
    animals.Kitten{
      Name: "Ms Tiggles",
      Hobbies: []string{
        "Playing with wool",
        "Eating",
      },
    },
    animals.Kitten{
      Name: "Mr Tom",
      Hobbies: []string{
        "Chasing own tail",
        "Napping on cushions",
        "Eating",
      },
    },
  }

  http.HandleFunc("/list", ListKittens)

  fmt.Println("Listening for connections on port", 9000)
  http.ListenAndServe(":9000", http.DefaultServeMux)
}
