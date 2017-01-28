package main

import (
	"fmt"
)

func main() {
  // Similar to an array
  mystring := "Hello"

  // Slices mystring from position 0 through to position 2
  // Should print "He"
  fmt.Println(string(mystring[0:2]))

  // This slice selected everything including first to the end
  // Should print "ello"
  fmt.Println(string(mystring[1:]))

  // Prints the byte number
  // Should print 72
  fmt.Println(mystring[0])

  // string will turn a byte number back into a normal letter
  // Should print "H"
  fmt.Println(string(mystring[0]))
}
