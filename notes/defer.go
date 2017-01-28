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

  // Output should be:
  // "three"
  // "two"
  // "one"
}

// Real world example of why you would use it
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

func main() {
  openFile()
  defer doWork()
  closeFile()

  // Output should be:
  // "Open 1"
  // "Open 2"
  // "Deferred: Open 3"
  // "Close all"
  // "Panic: ARGH"
}
