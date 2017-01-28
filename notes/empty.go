package main

import (
	"fmt"
  "time"
)

func main() {
  // print default "empty" values of types

  var structure struct{}
  fmt.Println(structure)

  var integer int
  fmt.Println(integer)

  var date time.Time
  fmt.Println(date)
}
