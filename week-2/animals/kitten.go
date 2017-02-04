package animals // needs to be the same as folder(?)

type Kitten struct {
  Name string
}

// Must use k reference (hence the *)
// Otherwise it creates a copy of a kitten
// Using a * mutates the original, which is the intended behaviour here
func (k *Kitten) SetName(name string) {
  k.Name = name
}

func (k Kitten) GetName() string {
  return k.Name
}
