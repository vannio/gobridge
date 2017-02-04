package main

import (
  "fmt"
)

// Orders of functions declared do not matter to the compiler

// 'main' is a special function in Go
// It's executed when the application runs
func main() {
  // Assign both sets of values returned from a function
  // This is similar to Javascript's ES6 destructuring assignment
  sum, success := add(1,3)

  // Can also be written as: if success == true
  if success {
    fmt.Println(sum, success)
    // => 4 true
  }
}

// This function signature accepts 2 integers and returns an integer,
// hence specifying the return type (int).
// func add(x, y int) int {
//   return x + y
// }

// Return 2 values by specifying more than 1 return type
// Useful for error handling as Go doesn't really do this
func add(x, y int) (int, bool) {
  return (x + y), true
}
