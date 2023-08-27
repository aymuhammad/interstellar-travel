package main

import "fmt"

func main() {
	number1 := 23
	number2 := 12
	var result bool

	// eqaul to operator
	result = (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, result)

	// not equal to operator
	result = (number1 != number2) // ! is the NOT operator in GoLang
	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	// greater than operator
	result = (number1 > number2)
	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	// less than operator
	result = (number1 < number2)
	fmt.Printf("%d < %d returns %t \n", number1, number2, result)
}
