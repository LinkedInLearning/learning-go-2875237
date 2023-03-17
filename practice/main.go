package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	var anInt int =5
	var aFloat float64=42
	sum:=float64(anInt) + aFloat
	fmt.Printf("Sum: %v, Type: %T\n", sum, sum)

	// reader :=bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
    // input, _:=reader.ReadString('\n')
	// fmt.Println("You entered:", input)

	// fmt.Print("Enter a number: ")
	// numInput, _ := reader.ReadString('\n')
	// aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput),64)
	// if err !=nil{
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println("Value of number: ", aFloat)
	// }


	
}
// const aCounst int =20
// func main() {
// 	var aString string = "This is GO!"
// 	fmt.Println(aString)
// 	fmt.Printf("The var type is : %T\n", aString)
// 	fmt.Printf("--------------------\n")

// 	var anInteger int =42
// 	fmt.Println(anInteger)

// 	var defaultInt int
// 	fmt.Println(defaultInt)
// 	fmt.Printf("--------------------\n")

// 	var anotherString string="This is another string"
// 	fmt.Println(anotherString)
// 	fmt.Printf("The variable is : %T\n", anotherString)
// 	fmt.Printf("--------------------\n")
// 	var anotherInt int=2002
// 	fmt.Println(anotherInt)
// 	fmt.Printf("%d The variable is : %T\n", anotherInt, anotherInt)
	
// 	fmt.Println(aCounst)
// 	fmt.Printf("The variable of constant is : %T\n", aCounst)


// }
