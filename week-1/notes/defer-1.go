package main

import (
  "fmt"
)

func one() {
  fmt.Println("one")
}

func two() {
  fmt.Println("two")
}

func three() {
  fmt.Println("three")
}

func main() {
  // defer is used to change the order of functions called
  
  // This will run last
  defer one()
  
  // The last defer is the first one called
  defer two()
  
  // This code is not deferred, so will run first!
  three()
  
  // =>  "three"
  //     "two"
  //     "one"
}
