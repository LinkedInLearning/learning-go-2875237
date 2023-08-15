package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter Test: ")
	input, _:=reader.ReadString('\n')
	fmt.Println("You Entered:", input)
}
