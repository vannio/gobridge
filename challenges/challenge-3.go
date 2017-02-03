package main

import (
  "fmt"
)

func main() {
  // Create slice of numbers between 0 - 10
  nums := []int{0,1,2,3,4,5,6,7,8,9,10};

  // Iterate through numbers
  for _, num := range nums {
    // Print only when it's an even number
    if (num % 2 == 0) {
      fmt.Println(num);
    }
  }
}
