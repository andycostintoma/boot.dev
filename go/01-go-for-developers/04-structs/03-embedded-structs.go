package main

import (
	"fmt"
)

// Define the car struct
type car struct {
	brand string
	model string
}

// Define the truck struct with an embedded car struct
type truck struct {
	car     // Embedding the car struct
	bedSize int
}

func main() {
	// Initialize a truck struct with embedded car struct
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			brand: "Toyota",
			model: "Tacoma",
		},
	}

	// Access truck fields, including embedded car fields
	fmt.Println("Truck Details:")
	fmt.Printf("Brand: %s\n", lanesTruck.brand) // Accessing the embedded field directly
	fmt.Printf("Model: %s\n", lanesTruck.model) // Accessing the embedded field directly
	fmt.Printf("Bed Size: %d\n", lanesTruck.bedSize)
}
