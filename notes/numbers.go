package main

import (
  "fmt"
  "reflect"
)

func main() {
  // Whole numbers are integer types
  const magicNumber = 3
  fmt.Println(reflect.TypeOf(magicNumber))
  // => int

  // Float types are numbers with decimal places.
  // Only use floats when necessary as they use more memory than integers
  const half = 0.5
  fmt.Println(reflect.TypeOf(half))
  // => float64
}
