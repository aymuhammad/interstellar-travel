package main

import "fmt"

func main() {
	var name, state string
	var age, salary int

	fmt.Println("Enter your name and state: ")
	fmt.Scanln(&name, &state)

	fmt.Println("Enter your age and salary: ")
	fmt.Scanln(&age, &salary)

	fmt.Printf("Name: %s\nState: %t", name, state)
	// fmt.Printf("Age: %s\nSlary: %d", age, salary)
}
