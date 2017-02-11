package animals // needs to be the same as folder(?)

// Kitten : Meow
type Kitten struct {
	Name    string
	Hobbies []string
	Likes   int
}

// Must use k reference (hence the *)
// Otherwise it creates a copy of a kitten
// Using a * mutates the original, which is the intended behaviour here

// SetName : sets a pet's name
func (k *Kitten) SetName(name string) {
	k.Name = name
}

// GetName : gets a pet's name
func (k *Kitten) GetName() string {
	return k.Name
}

// SetHobbies : sets a pet's hobbies
func (k *Kitten) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

// GetHobbies : gets a pet's hobbies
func (k *Kitten) GetHobbies() []string {
	return k.Hobbies
}

// AddLike : increases the number of likes a pet has
func (k *Kitten) AddLike() {
	k.Likes++
}

// GetLikes : returns the count of likes
func (k *Kitten) GetLikes() int {
	return k.Likes
}
