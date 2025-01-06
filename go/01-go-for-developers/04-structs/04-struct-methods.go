package main

import (
	"fmt"
)

// Define the rect struct
type rect struct {
	width  int
	height int
}

// Method to calculate the area of the rectangle
// `r` is the receiver of type `rect`
func (r rect) area() int {
	return r.width * r.height
}

// Method to calculate the perimeter of the rectangle
// `r` is the receiver of type `rect`
func (r rect) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	// Initialize a rect struct
	r := rect{
		width:  5,
		height: 10,
	}

	// Call the methods and print the results
	fmt.Printf("Rectangle: %+v\n", r)
	fmt.Printf("Area: %d\n", r.area())
	fmt.Printf("Perimeter: %d\n", r.perimeter())
}
