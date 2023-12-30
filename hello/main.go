package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I lunch this App At :", n)

	t := time.Date(2023, time.December, 30, 7, 11, 0, 0, time.UTC)
	fmt.Println("I lunch this App At :", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sat Dec 30 07:11:00 2023")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
