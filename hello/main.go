package main

import (
	"fmt"
)

func main() {
	fmt.Println("Group related values in structs")
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%v\n", poodle)
	fmt.Printf("Breed : %v\nweight : %v\n", poodle.Breed ,poodle.weight)
	poodle.weight = 9
	fmt.Printf("Breed : %v\nweight : %v\n", poodle.Breed ,poodle.weight)
}

// Dog in an explort struct
type Dog struct {
	Breed  string
	weight int
}
