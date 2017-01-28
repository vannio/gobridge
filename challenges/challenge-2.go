package main

import (
  "fmt"
)

func main() {
  // Challenge:
  // Print out strings for "One", "Four", "Eight"
  // whilst looping over a slice of sequenced numbers

  numbers := []int{0,1,2,3,4,5,6,7,8,9}

  for _, value := range numbers {
    switch value {
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
  }
}
