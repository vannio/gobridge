package main

import (
  "fmt"
)

// func main() {
//   defer one()
//
//   // the last defer will be the first one called
//   defer two()
//
//   three()
// }
//
// func one() {
//   fmt.Println("one")
// }
//
// func two() {
//   fmt.Println("two")
// }
//
// func three() {
//   fmt.Println("three")
// }

// Real world example of why you would use it
func main() {
  openFile()

  // the last defer will be the first one called
  defer doWork()

  closeFile()
}

func openFile() {
  fmt.Println("Open 1")
  defer fmt.Println("Deferred: Open 3")
  fmt.Println("Open 2")
}

func doWork() {
  panic("ARGH")
}

func closeFile() {
  fmt.Println("Close all")
}
