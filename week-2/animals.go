// Main entry point!
package main

// blank lines added (by convention) to separate types of packages
// eg. standard packages, custom packages
import (
  "fmt"

  "github.com/vannio/gobridge/week-2/animals"
)

func main() {
  kitty := animals.Kitten{} // namespace from package name
  kitty.SetName("Mr Tiggles")
  fmt.Println(kitty.GetName()) // => "Mr Tiggles"

  dog := animals.Dog{}
  fmt.Println(dog.Bark())
}
