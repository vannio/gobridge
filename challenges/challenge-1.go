package main

import (
  "fmt"
)

func main() {
  // TODO:
  // Print out strings for One, Four, Eight
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
      fmt.Println(i)
    }

    i++
  }
}
