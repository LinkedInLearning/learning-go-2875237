package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Evaluate expressions with switch statements")
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	default:
		result = "it's another day"
	}
	fmt.Println(result)
}
