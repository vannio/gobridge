package main

import (
  "fmt"
)

func main() {
  // TODO:
  // Print out strings for One, Four, Eight
  // whilst iterating over an array of numbers
  
  numbers := []int{0,1,2,3,4,5,6,7,8,9}

  for _, val := range numbers {
    switch val {
    case 1:
      fmt.Println("One")
    case 4:
      fmt.Println("Four")
    case 8:
      fmt.Println("Eight")
    default:
      fmt.Println(val)
    }
  }
}
