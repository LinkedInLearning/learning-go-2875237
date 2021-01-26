package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	roundedSum := math.Round(floatSum)
	fmt.Println("Rounded sum:", roundedSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(f1)
	b2.SetFloat64(f2)
	b3.SetFloat64(f3)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)

}
