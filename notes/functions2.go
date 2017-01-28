package main

import (
	"fmt"
  "time"
)

// If there are multiple variables are of the same type,
// specifying type can be written shorthand
func sayHello(firstName, lastName, birthday string) {
  age := calculateAge(birthday)
  fmt.Println("Hi", firstName, lastName + ", you are " + age + " years old")
}

func calculateAge(birthday string) string {
  const hoursInDay = 24
  const daysinYear = 365.23
  const shortForm = "2006-Jan-02"

  t, _ := time.Parse(shortForm, birthday)

  hoursSinceBirthday := time.Since(t).Hours()
  ageInYears := hoursSinceBirthday / hoursInDay / daysinYear
  return fmt.Sprintf("%.1f", ageInYears)
}

func main() {
  sayHello("Van", "Le", "1990-Jul-24");
}
