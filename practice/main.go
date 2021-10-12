package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I recorded the video at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Println("The type of parseTime is %T\n", parsedTime)
}
