package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program conditional logic")
	theAns := 42
	var result string
	if theAns < 0 {
		result = "The result is less than 0"
	} else if theAns == 0 {
		result = "The result null "
	} else {
		result = "The result is Greater than 0"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "The result is less than 0"
	} else if x == 0 {
		result = "The result null "
	} else {
		result = "The result is Greater than 0"
	}
	fmt.Println(result)

}
