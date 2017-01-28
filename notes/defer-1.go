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
  // the last defer will be the first one called
  defer one()
  defer two()
  three()
  // =>  "three"
  //     "two"
  //     "one"
}
