package functions

import "fmt"

func add(i, j int) int {
	fmt.Printf("Enter the numbers you want to add: %d+%d\n", i, j)
	return i + j
}

func Add() {
	result := add(5, 6)
	fmt.Println("Sum is: ", result)
}

