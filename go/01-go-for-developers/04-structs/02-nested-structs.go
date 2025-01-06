package main

import (
	"fmt"
)

// Define the wheel struct
type wheel struct {
	radius   int
	material string
}

// Define the car struct with nested wheel structs
type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

func main() {
	// Initialize a car struct with nested wheel structs
	myCar := car{
		brand:   "Toyota",
		model:   "Corolla",
		doors:   4,
		mileage: 12000,
		frontWheel: wheel{
			radius:   16,
			material: "Alloy",
		},
		backWheel: wheel{
			radius:   16,
			material: "Alloy",
		},
	}

	// Access and modify nested struct fields
	myCar.frontWheel.radius = 17
	myCar.backWheel.material = "Steel"

	// Print car details, including nested struct fields
	fmt.Printf("Car Details: %+v\n", myCar)
	fmt.Printf("Front Wheel Details: %+v\n", myCar.frontWheel)
	fmt.Printf("Back Wheel Details: %+v\n", myCar.backWheel)
}

// Output:
// Car Details: {brand:Toyota model:Corolla doors:4 mileage:12000 frontWheel:{radius:17 material:Alloy} backWheel:{radius:16 material:Steel}}
// Front Wheel Details: {radius:17 material:Alloy}
// Back Wheel Details: {radius:16 material:Steel}
