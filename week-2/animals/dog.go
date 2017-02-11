package animals

// Dog : Woof
type Dog struct {
	Name    string
	Hobbies []string
	Likes   int
}

// SetName : sets a pet's name
func (d *Dog) SetName(name string) {
	d.Name = name
}

// GetName : gets a pet's name
func (d *Dog) GetName() string {
	return d.Name
}

// SetHobbies : sets a pet's hobbies
func (d *Dog) SetHobbies(hobbies []string) {
	d.Hobbies = hobbies
}

// GetHobbies : gets a pet's hobbies
func (d *Dog) GetHobbies() []string {
	return d.Hobbies
}

// AddLike : increases the number of likes a pet has
func (d *Dog) AddLike() {
	d.Likes++
}

// GetLikes : returns the count of likes
func (d *Dog) GetLikes() int {
	return d.Likes
}

// Bark : dogs bark.
func (d *Dog) Bark() string {
	return "Woof!"
}
