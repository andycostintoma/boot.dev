package main

import (
	"fmt"
)

// getPoint function returns two integers
func getPoint() (x int, y int) {
	return 3, 4
}

// getCircle function returns two values: the center point and the radius
func getCircle() (cx int, cy int, radius int) {
	return 1, 2, 5
}

func main() {
	// Ignoring the second return value from getPoint
	x, _ := getPoint()
	fmt.Println("Only using x from getPoint:", x) // Output: Only using x from getPoint: 3

	// Ignoring the first and second return values from getCircle, only using the radius
	_, _, radius := getCircle()
	fmt.Println("Only using radius from getCircle:", radius) // Output: Only using radius from getCircle: 5

	// Using all return values from getCircle
	cx, cy, r := getCircle()
	fmt.Println("Using all return values from getCircle:", cx, cy, r) // Output: Using all return values from getCircle: 1 2 5
}
