package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("A simple calculator")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Value 1: ")
	n1, _ := reader.ReadString('\n')
	f1, err := strconv.ParseFloat(strings.TrimSpace(n1), 64)

	fmt.Printf("Enter Value 2: ")
	n2, _ := reader.ReadString('\n')
	f2, err := strconv.ParseFloat(strings.TrimSpace(n2), 64)
	if err != nil {
		fmt.Print(err)
	} else {

		sum := f1 + f2
		fmt.Println("Sum is ", sum)
	}
}
