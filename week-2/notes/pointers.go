// Main entry point!
package main

import (
  "fmt"
)

func withoutPointer() {
  a := 5
  b := a // Doesn't _really_ equal `a`. `b` is a copy of `a`
  a = 7  // ..which is why this doesn't update `b`

  fmt.Println(a) // => 7
  fmt.Println(b) // => 5
}

func withPointers() {
  a := 5
  b := &a // pointer to memory address | "take a reference"
  a = 7

  fmt.Println(a)
  fmt.Println(b)  // will print memory address
  fmt.Println(*b) // dereference | "get the value"
}

func main() {
  withoutPointer()
  withPointers()
}
