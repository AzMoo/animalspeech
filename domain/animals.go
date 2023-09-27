package domain

// Cat says meow
type Cat struct {
	Name string
}

func (Cat) Speak() string { return "meow" }

// Dog says woof
type Dog struct {
	Name string
}

func (Dog) Speak() string { return "woof" }
