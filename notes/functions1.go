package main

import (
	"fmt"
)

// 'main' is a special function in Go
// It's executed when the application runs
func main() {
  // Saves both sets of values returned from a function
  // Similar to Javascript's ES6 destructuring assignment
  sum, success := add(1,3)

  if success {
    fmt.Println(sum, success)
  }
}

// Orders of functions do not matter to the compiler

// This function signature accepts 2 integers
// and returns an integer, hence specifying the return type (int)
// func add(x, y int) int {
//   return x + y
// }

// Return 2 values by specifying more than 1 return type
// Useful for error handling as Go doesn't really do this 
func add(x, y int) (int, bool) {
  return (x + y), true
}
