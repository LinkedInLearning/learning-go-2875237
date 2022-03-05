package main

import "fmt"

const aConstant int = 64

func main() {

	var aString string = "Wellcome to Go!"
	fmt.Println(aString)
	fmt.Printf("The variable type is: %T\n", aString)

	var anInteger uint8 = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is: %T\n", anotherString)

	var anotherint = 52
	fmt.Println(anotherint)
	fmt.Printf("The variable type is: %T\n", anotherint)

	myString := "A new string!"  // THis only works for variable inside functions.
	fmt.Println(myString)
	fmt.Printf("The variable type is: %T\n", myString)

	fmt.Println(aConstant)
	fmt.Printf("The variable type is: %T\n", aConstant)
}
