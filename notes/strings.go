package main

import (
	"fmt"
)

func main() {
  // Declare variable and assign value on one line
  hello := "Hello"

  // Or like this..
  // var hello = "Hello"

  // Declare empty variable and specify type
  var target string

  // Check the length of the 'target' string
  // Should be 0 since it's not assigned a value
  // Refer to empty.go for other examples
  fmt.Println(len(target))

  // Assign a value to the empty variable
  target = "World"

  // Print "Hello World"
  fmt.Println(hello, target)

  // Assign a new value in existing variable
  target = "Van"

  // Should now print "Hello Van"
  fmt.Println(hello, target)

  // Cannot assign a new value because it's the wrong type
  // This is strict, unlike in Ruby or Javascript
  // Attempting to assign an integer into a string variable should throw error
  // target = 1

  // Immutable data (doesn't/cannot change)
  const myConstant = "Constant"
  fmt.Println(myConstant)

  // Cannot assign a new value to a constant
  // It'll just throw an error! Boo
  // myConstant = "Test"
}
