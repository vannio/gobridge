package main

import (
  "fmt"
)

func main() {
  mystring := "Hello"

  // Assignment, evaluation, increment
  for i := 0; i < 5; i++ {
    // i is usually used as shorthand for 'index', aka position
    // mystring[i] will access the value of whatever is at position i
    //
    fmt.Println(i, string(mystring[i]))
  }

  // Alternative to while loops, since Go doesn't include it natively:
	// Assign an initial value to i
  i := 0
  // Code will keep executing until i is equal to 5
  for i < 5 {
    // prints the position (i) as well as the value at that position
    fmt.Println(i, string(mystring[i]))

    // i++ is the same as writing i = i + 1
    i++
    // If we don't change the value of i, i will _always_ be 0
    // so this will run indefinitely and probably crash your application
  }

  // For each byte in the string,
  // give back the position (index) and the value
  for i, val := range mystring {
    // fmt.Sprint() converts any interface (object) into a string
    fmt.Println(fmt.Sprint(i) + " " + string(val))
  }

  // Assign values to _ as a 'junk drawer' in case we don't need to use it
  // Otherwise compiler will error if we assign a value (but don't use it)
  for _, val := range mystring {
    const letterH = 72
    const letterE = 101

    switch val {
    case letterH:
      fmt.Println(string(72))
    case letterE:
      fmt.Println(string(101))
    default:
      fmt.Println("Not H, is " + string(val))
    }
  }
}
