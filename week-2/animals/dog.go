package animals

type Dog struct {
	Name    string
	Hobbies []string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) SetHobbies(hobbies []string) {
	d.Hobbies = hobbies
}

func (d *Dog) GetHobbies() []string {
	return d.Hobbies
}

func (d *Dog) Bark() string {
	return "Woof!"
}
