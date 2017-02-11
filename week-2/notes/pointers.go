// Main entry point!
package main

import (
	"fmt"
)

func withoutPointer() {
	a := 5

	// Doesn't really equal `a`, just makes a copy of `a`
	b := a

	// ..which is why this doesn't update `b`
	a = 7

	fmt.Println(a)
	// => 7

	fmt.Println(b)
	// => 5
}

func withPointers() {
	a := 5

	// pointer to memory address | "take a reference"
	b := &a

	a = 7

	fmt.Println(a)
	// => 7

	// will print memory address
	fmt.Println(b)
	// => 0xc42007c048

	// dereference | "get the value"
	fmt.Println(*b)
	// => 7
}

func main() {
	withoutPointer()
	withPointers()
}
