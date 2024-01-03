package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 22, 45
	intSum := i1 + i2 + i3
	fmt.Println("The integer Sum :", intSum)

	f1, f2, f3 := 12.2, 22.4, 45.5
	floatSum := f1 + f2 + f3
	fmt.Println("The integer Sum :", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The new integer Sum :", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference : %.2f\n", circumference) // print f change  the output format
}
