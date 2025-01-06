package main

import (
	"errors"
	"fmt"
)

// Function with named return values
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, err
}

func main() {
	// Test the calculator function
	multiplication, division, err := calculator(6, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Multiplication: %d, Division: %d\n", multiplication, division)
	}

	// Test the calculator function with division by zero
	_, _, err = calculator(6, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: can't divide by zero
	}
}
