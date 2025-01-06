package main

import (
	"fmt"
)

// Car struct definition
type Car struct {
	color string
}

// Method with a pointer receiver
func (c *Car) setColorWithPointer(color string) {
	c.color = color
}

// Method with a value receiver
func (c Car) setColorWithoutPointer(color string) {
	c.color = color
}

func main() {
	// Example with pointer receiver
	car1 := Car{color: "white"}
	fmt.Println("Initial color of car1:", car1.color)
	car1.setColorWithPointer("blue")
	fmt.Println("Color of car1 after setColorWithPointer:", car1.color)

	// Example with value receiver
	car2 := Car{color: "white"}
	fmt.Println("\nInitial color of car2:", car2.color)
	// pointer receiver don't require a pointer to be used when calling the method
	car2.setColorWithoutPointer("blue")
	fmt.Println("Color of car2 after setColorWithoutPointer:", car2.color)
}
