package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int{5, 4, 3, 2, 1}
	fmt.Println(numbers)

	fmt.Println("Number of colors: ", len(colors))
	fmt.Println("Number of numbers: ", len(numbers))
}
