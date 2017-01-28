package main

import (
  "fmt"
)

func main() {
  mystring := "Hello"

  // Here are 3 different ways to output:
  // => 0 H
  //    1 e
  //    2 l
  //    3 l
  //    4 o

  // 1 -
  // Assignment, evaluation, increment
  for i := 0; i < 5; i++ {
    // i is usually used as shorthand for 'index', aka position
    fmt.Println(i, string(mystring[i]))
  }

  // 2 -
  // Alternative to while loops, since Go doesn't include it natively:
  i := 0
  for i < 5 {
    fmt.Println(i, string(mystring[i]))

    // Same as i = i + 1
    i++
    // If we don't change the value of i, i will _always_ be 0
    // so this will run indefinitely and probably crash your application
  }

  // 3 -
  // For each byte in the string, give back the position (i) and the value
  for i, val := range mystring {
    // fmt.Sprint() converts any interface (object) into a string
    fmt.Println(fmt.Sprint(i) + " " + string(val))
  }

  // Assign values to _ as a 'junk drawer', if we don't need it
  // Otherwise compiler will throw a "declared and not used" error
  for i, val := range mystring {
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
