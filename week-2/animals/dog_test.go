package animals

import (
	"reflect"
	"testing"
)

func TestSetHobbiesSetsTheHobbies(t *testing.T) {
	// Setup
	hobbies := []string{"Barking"}
	dog := Dog{}

	// Do work
	dog.SetHobbies(hobbies)

	// Assert
	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
		t.Fail()
	}
}

func TestBarkGetsWoofed(t *testing.T) {
	dog := Dog{Name: "Spot"}

	if dog.Bark() != "Woof!" {
		t.Fail()
	}
}
