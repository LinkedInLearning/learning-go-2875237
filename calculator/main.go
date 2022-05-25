package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Value 1 : ")
	inp1, _ := reader.ReadString('\n')
	val1, err1 := strconv.ParseFloat(strings.TrimSpace(inp1), 64)
	if err1 == nil {

		fmt.Printf("Value 2 : ")
		inp2, _ := reader.ReadString('\n')
		val2, err2 := strconv.ParseFloat(strings.TrimSpace(inp2), 64)
		if err2 == nil {
			sum := val1 + val2
			fmt.Printf("The sum of %.2f and %.2f is %.2f", val1, val2, sum)
		} else {
			fmt.Println(err2)
			panic("Value 2 must be a number")
		}
	} else {
		fmt.Println(err1)
		panic("Value 1 must be a number")
	}

}
