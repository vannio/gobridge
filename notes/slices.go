package main

import (
  "fmt"
)

func main() {
  // Slices are similar to arrays - both sequences of typed data.
  // A slice is a lightweight wrapper around an array which
  // usually represents a subset of that underlying array.

  // This string can be considered a sequence
  // "H", "e", "l", "l", "o"
  mystring := "Hello"

  first := mystring[0]
  // Prints the byte number of value at position 0 (the first character, "H")
  fmt.Println(first)
  // => 72

  // string() will turn a byte number back into a normal letter
  fmt.Println(string(first))
  // => "H"

  // Slice mystring from position 0 through to position 2
  fmt.Println(string(mystring[0:2]))
  // => "He"

  // This slice selected everything between position 1 ("e") to the end ("o")
  fmt.Println(string(mystring[1:]))
  // => "ello"
}
