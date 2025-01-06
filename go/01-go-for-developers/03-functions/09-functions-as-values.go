package main

import (
	"fmt"
)

// Define the add function
func add(x, y int) int {
	return x + y
}

// Define the mul function
func mul(x, y int) int {
	return x * y
}

// Define the aggregate function that accepts another function as an argument
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}

func main() {
	// Use the aggregate function with the add function
	sum := aggregate(2, 3, 4, add)
	fmt.Printf("Sum: %d\n", sum) // Output: Sum: 9

	// Use the aggregate function with the mul function
	product := aggregate(2, 3, 4, mul)
	fmt.Printf("Product: %d\n", product) // Output: Product: 24
}
