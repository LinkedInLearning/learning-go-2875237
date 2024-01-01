package main

import (
	"fmt"
)

func main() {
	fmt.Println("Define and call functions")
	doSomthing()
	sum := addValue(4, 5)
	fmt.Println("The sum equal to :", sum)

	multiSum, multiCount := getSumAllValues(4, 5, 6)
	fmt.Println("Sum of multiple values", multiSum)
	fmt.Println("Count of items", multiCount)

}
func doSomthing() {
	fmt.Println("Doing somthing")
}

// in go i could declare the type of values just once if it's the same in function parameteres
func addValue(value1, value2 int) int {
	return value1 + value2
}

/*this function seems the same as the function before
func addValue(value1 int ,value2 int) int {
	return value1 + value2
}*/

// but you can also return multiple values from a single function for Arbitrary numbers
func getSumAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
