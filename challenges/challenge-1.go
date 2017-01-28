package main

import (
  "fmt"
)

func main() {
  // Challenge:
  // Print out strings for "One", "Four", "Eight"
  // whilst looping through 0 to 9

  i := 0
  for i <= 9 {
    switch i {
    case 1:
      fmt.Println("One")
    case 4:
      fmt.Println("Four")
    case 8:
      fmt.Println("Eight")
    default:
      // If a value is neither 1, 4 or 8
      // print "Nope"
      fmt.Println("Nope")
    }

    i++
  }
}
