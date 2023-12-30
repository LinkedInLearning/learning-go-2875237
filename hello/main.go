package main

import (
	"fmt"
)

func main() {
	intPointer := 45.31
	// p is a pointer initialized with the address of intPointer
	var p = &intPointer
	fmt.Println("pointure", *p)
	// p1 is a pointer to the address of intPointer
	p1 := &intPointer
	fmt.Println("pointure 2", *p1)

	*p1 = *p1 / 31
	fmt.Println("pointure 1:", *p1)
	fmt.Println("pointure 2:", *p)
	fmt.Println("value :", intPointer)
}
