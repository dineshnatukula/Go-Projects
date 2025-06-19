package backend

import "fmt"

// Define the interface
type Speaker interface {
	Speak() string
}

// Struct types
type Dog struct{}
type Cat struct{}

// Implement Speak for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Implement Speak for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts the interface
func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func InterfaceDemo() {
	d := Dog{}
	c := Cat{}
	MakeItSpeak(d) // Woof!
	MakeItSpeak(c) // Meow!
}
