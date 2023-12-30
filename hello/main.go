package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Green", "red", "blue"} // now [] it's a slice , i can add new items
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Print(colors, "\n")

	numbers := make([]int, 5) //([]int, 5(number of items) , 5 (capacity that caps the number of items))
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	// sort slice by adding the sort package
	sort.Ints(numbers)
	fmt.Println(numbers)

}
