package Printing

import (
	"fmt"
)
 const aCounst int =20
func main() {
	var aString string = "This is GO!"
	fmt.Println(aString)
	fmt.Printf("The var type is : %T\n", aString)
	fmt.Printf("--------------------\n")

	var anInteger int =42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("--------------------\n")

	var anotherString string="This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable is : %T\n", anotherString)
	fmt.Printf("--------------------\n")
	var anotherInt int=2002
	fmt.Println(anotherInt)
	fmt.Printf("%d The variable is : %T\n", anotherInt, anotherInt)
	
	fmt.Println(aCounst)
	fmt.Printf("The variable of constant is : %T\n", aCounst)


}
