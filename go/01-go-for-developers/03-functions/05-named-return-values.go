package main

import (
	"fmt"
)

// Function using named return values and a naked return
func getCoords() (x, y int) {
	// x and y are initialized to their zero values (0 for int)
	x = 10
	y = 20

	// Naked return: automatically returns x and y
	return
}

// Function with explicitly returned values
func getDimensions() (width, height int) {
	// width and height are initialized to their zero values (0 for int)
	width = 100
	height = 50

	// Explicitly returning width and height
	return width, height
}

func main() {
	// Using getCoords function with named return values and naked return
	x, y := getCoords()
	fmt.Printf("Coordinates: x = %d, y = %d\n", x, y) // Output: Coordinates: x = 10, y = 20

	// Using getDimensions function with explicit return values
	width, height := getDimensions()
	fmt.Printf("Dimensions: width = %d, height = %d\n", width, height) // Output: Dimensions: width = 100, height = 50
}
