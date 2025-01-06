package main

import (
	"fmt"
)

// Function with no arguments and no return value
func greet() {
	fmt.Println("Hello, World!")
}

// Function with one argument and no return value
func printNumber(num int) {
	fmt.Printf("The number is: %d\n", num)
}

// Function with two arguments and a return value
func add(x int, y int) int {
	return x + y
}

// Function with two arguments of the same type, showcasing a more concise parameter list
func subtract(x, y int) int {
	return x - y
}

// Function with a string argument and returning a string
func personalizeGreet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	// Calling function with no arguments
	greet()

	// Calling function with one argument
	printNumber(5)

	// Calling function with two arguments and using the return value
	sum := add(3, 4)
	fmt.Printf("3 + 4 = %d\n", sum)

	// Calling function with two arguments of the same type and using the return value
	difference := subtract(10, 5)
	fmt.Printf("10 - 5 = %d\n", difference)

	// Calling function with a string argument and using the return value
	message := personalizeGreet("Alice")
	fmt.Println(message)
}
