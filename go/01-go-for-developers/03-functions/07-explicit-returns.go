package main

import (
	"errors"
	"fmt"
)

// Function with named return values and explicit return
func getCoords(useDefault bool) (x, y int) {
	if useDefault {
		// Explicit return, overriding the named return values
		return 5, 6
	}

	// Named return values assigned but using explicit return
	x = 10
	y = 20
	return x, y // explicit return using the names
}

// Function with named return values and naked return
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero") // explicit return with error
	}
	mul = a * b
	div = a / b
	return // naked return, returns mul, div, and err
}

func main() {
	// Using the getCoords function with default values
	x, y := getCoords(true)
	fmt.Printf("Coordinates with default values: x = %d, y = %d\n", x, y) // Output: Coordinates with default values: x = 5, y = 6

	// Using the getCoords function without default values
	x, y = getCoords(false)
	fmt.Printf("Coordinates without default values: x = %d, y = %d\n", x, y) // Output: Coordinates without default values: x = 10, y = 20

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
