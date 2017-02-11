package main

import (
	"fmt"
)

// Real world example of why you would use defer

func openFile() {
	fmt.Println("Open 1")

	// Nested defer
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

	// =>  "Open 1"
	//     "Open 2"
	//     "Deferred: Open 3"
	//     "Close all"
	//     "Panic: ARGH"
}
