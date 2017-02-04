package animals

// Anything considered a pet MUST have the included methods
// Interfaces are always a reference
type Pet interface {
  SetName(name string)
  GetName() string
  SetHobbies(hobbies []string)
  GetHobbies() []string
}
