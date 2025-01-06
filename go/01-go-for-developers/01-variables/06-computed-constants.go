package main

import (
	"fmt"
)

// Static constants
const myInt = 15
const pi = 3.14159

// Computed constants
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

// Complex computations that can occur at compile time
const radius = 5
const circumference = 2 * pi * radius

func main() {
	// Using static constants
	fmt.Println("Static constant myInt:", myInt)
	fmt.Println("Static constant pi:", pi)

	// Using computed constants
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Full Name:", fullName)

	// Using constants with arithmetic operations
	fmt.Println("Radius:", radius)
	fmt.Println("Circumference:", circumference)

	// Attempting to declare a constant that requires run-time computation would cause a compile error
	// Uncommenting the following line will cause an error
	// const currentTime = time.Now()
}
