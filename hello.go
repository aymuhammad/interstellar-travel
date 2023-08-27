package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var cars = [4]string{"Volvo", "BMW", "FORD", "MAZDA"}

	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(cars)
}
