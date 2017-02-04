package main

import (
  "fmt"
  "time"
)

type Person struct {
  Name string
  Age int
}

func calculateBirthYear(age int) int {
  currentYear := time.Now().Year()
  return (currentYear - age)
}

func main() {
  people := []Person{
    {"Agnes", 10},
    {"Maria", 15},
  }

  // OR
  // agnes := Person{"Agnes", 10}
  // maria := Person{"Maria", 15}
  // people := []Person{agnes, maria}

  for _, person := range people {
    birthYear := fmt.Sprint(calculateBirthYear(person.Age))
    fmt.Println(person.Name + " was born in the year " + birthYear)
  }
}
