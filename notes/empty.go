package main

import (
  "fmt"
  "time"
)

func main() {
  // Default "empty" values of different types

  var structure struct{}
  fmt.Println(structure)
  // => {}

  var integer int
  fmt.Println(integer)
  // => 0

  var date time.Time
  fmt.Println(date)
  // => 0001-01-01 00:00:00 +0000 UTC

  // Length of empty string
  var string string
  fmt.Println(len(string))
  // => 0
}
