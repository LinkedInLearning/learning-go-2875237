package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Evaluate expressions with switch statements")
	rand.Seed(time.Now().Unix())

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Monday"
		fallthrough
	case 2:
		result = "Tuesday"
		//fallthrough // make the control flow less predictable
	default:
		result = "it's another day"
	}
	fmt.Println(result)
}
