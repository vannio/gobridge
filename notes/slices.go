package main

import (
  "fmt"
)

func main() {
  // Slices are similar to arrays - both sequences of typed data.
  // A slice is a lightweight wrapper around an array which
  // usually represents a subset of that underlying array.

  // This string can be considered a sequence:
  mystring := "Hello"
  // "H", "e", "l", "l", "o"

  first := mystring[0]
  // Prints the byte number of value at position 0 (the first character, "H")
  // Should print 72
  fmt.Println(first)

  // string() will turn a byte number back into a normal letter
  // Should print "H"
  fmt.Println(string(first))

  // Slice mystring from position 0 through to position 2
  // Should print "He"
  fmt.Println(string(mystring[0:2]))

  // This slice selected everything between position 1 ("e") to the end ("o")
  // Should print "ello"
  fmt.Println(string(mystring[1:]))
}
