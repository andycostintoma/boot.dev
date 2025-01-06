package main

import (
	"fmt"
)

// Define the car struct
type car struct {
	brand   string
	model   string
	doors   int
	mileage int
}

func main() {
	// Initialize a car struct using a struct literal
	myCar := car{
		brand:   "Toyota",
		model:   "Corolla",
		doors:   4,
		mileage: 12000,
	}

	// Print initial car details
	fmt.Printf("Initial Car: %+v\n", myCar)

	// Modify struct fields
	myCar.mileage += 500

	// Print the updated details
	fmt.Printf("Updated Car: %+v\n", myCar)
}

// Output:
// Initial Car: {brand:Toyota model:Corolla doors:4 mileage:12000}
// Updated Car: {brand:Toyota model:Corolla doors:4 mileage:12500}
