package main

import (
	"fmt"
)

// conversions function that accepts another function as an argument
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

// Named function for doubling a value
func double(a int) int {
	return a + a
}

func main() {
	// Using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	fmt.Printf("Using named function - newX: %d, newY: %d, newZ: %d\n", newX, newY, newZ)
	// Output: Using named function - newX: 2, newY: 4, newZ: 6

	// Using an anonymous function for squaring a value
	newX, newY, newZ = conversions(func(a int) int {
		return a * a
	}, 1, 2, 3)
	fmt.Printf("Using anonymous function - newX: %d, newY: %d, newZ: %d\n", newX, newY, newZ)
	// Output: Using anonymous function - newX: 1, newY: 4, newZ: 9
}
