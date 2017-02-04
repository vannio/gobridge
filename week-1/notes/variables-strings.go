package main

import (
  "fmt"
)

func main() {
  // Declare empty variable and specify type
  var hello string

  // Assign a value to the empty variable
  hello = "Hello"
  
  // Declare variable and assign value on one line
  target := "World"

  // Or like this..
  // var target = "World"

  fmt.Println(hello, target)
  // => "Hello World"

  // Assign a new value in existing variable
  target = "Van"

  fmt.Println(hello, target)
  // => "Hello Van"

  // target = 1
  // Cannot assign an integer to 'target' because it's the wrong type
  // This is strict, unlike in Ruby or Javascript

  const myConstant = "Constant"
  fmt.Println(myConstant)
  // => "Constant"

  // myConstant = "Test"
  // Cannot assign a new value to a constant
  // Constants are immutable data (doesn't/cannot change)
}
