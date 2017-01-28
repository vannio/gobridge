package main

import (
	"fmt"
  "time"
)

// Type does not _need_ to start with a capital letter
// but may make you want to gouge your eyes out later
type Kitten struct {
  firstname string
  Surname string
  Birthday string
}

// These functions belongs to the Kitten type
func (k Kitten) PrintName() {
  // Print the name of this particular kitten
  fmt.Println(k.firstname, k.Surname)
}

func (k Kitten) PrintAge() {
  age := calculateAge(k.Birthday)
  fmt.Println(age)
}

func main() {
  // Instantiate a structure
  kitty := Kitten{}

  // Assign some values
  kitty.firstname = "Mr"
  kitty.Surname = "Tiggles"
  kitty.Birthday = "2016-Jan-01"

  // Alternative methods to initialising structures:
  // Sets values in the order that they're declared in the type setup
  // kitty := Kitten{"Mr", "Tiggles"}

  // kitty := Kitten{
  //   Firstname: "Mr",
  //   Surname: "Tiggles",
  // }

  // Call functions that belongs to the Kitten type
  kitty.PrintName()
  kitty.PrintAge()
}

// Reused this from functions2.go file
func calculateAge(birthday string) string {
  const hoursInDay = 24
  const daysinYear = 365.23
  const shortForm = "2006-Jan-02"

  t, _ := time.Parse(shortForm, birthday)

  hoursSinceBirthday := time.Since(t).Hours()
  ageInYears := hoursSinceBirthday / hoursInDay / daysinYear
  return fmt.Sprintf("%.1f", ageInYears)
}
