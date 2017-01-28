package main

import (
  "fmt"
  "time"
)

func main() {
  // print default "empty" values of different types

  var structure struct{}
  fmt.Println(structure)
  // Output: {}

  var integer int
  fmt.Println(integer)
  // Output: 0

  var date time.Time
  fmt.Println(date)
  // Output: 0001-01-01 00:00:00 +0000 UTC
}
