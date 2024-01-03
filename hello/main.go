package main

import (
	"fmt"
)

func main() {
	fmt.Println("Define functions as methods of custom types")
	poodle := Dog{"Poodle", 10, "woof!"}
	fmt.Println(poodle)
	fmt.Printf("%v\n", poodle)
	fmt.Printf("Breed : %v\nweight : %v\n", poodle.Breed, poodle.weight)
	poodle.weight = 9
	fmt.Printf("Breed : %v\nweight : %v\n", poodle.Breed, poodle.weight)

	poodle.Speak()
	poodle.sound = "Arf !"
	poodle.Speak()
	poodle.SpeakThreeTimes()
}

// Dog in an explort struct
type Dog struct {
	Breed  string
	weight int
	sound  string
}

// exported fumction as method with upper case caracter in the name
func (d Dog) Speak() {
	fmt.Println(d.sound)
}

func (d Dog) SpeakThreeTimes() {
	d.sound = fmt.Sprintf("%v %v %v", d.sound, d.sound, d.sound)
	fmt.Println(d.sound)
}
//The ability to create custom methods for your own types,
// makes Go behave more like a fully Object Oriented language.
